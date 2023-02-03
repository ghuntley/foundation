// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package buildkit

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/docker/cli/cli/config"
	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/session"
	"github.com/moby/buildkit/session/auth"
	"github.com/moby/buildkit/session/auth/authprovider"
	"github.com/moby/buildkit/session/secrets/secretsprovider"
	"github.com/moby/buildkit/session/sshforward/sshprovider"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"namespacelabs.dev/foundation/framework/rpcerrors"
	"namespacelabs.dev/foundation/framework/rpcerrors/multierr"
	"namespacelabs.dev/foundation/framework/secrets"
	"namespacelabs.dev/foundation/internal/artifacts/oci"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/parsing/devhost"
	"namespacelabs.dev/foundation/internal/workspace/dirs"
	"namespacelabs.dev/foundation/schema"
)

type FrontendRequest struct {
	Def            *llb.Definition
	OriginalState  *llb.State
	Frontend       string
	FrontendOpt    map[string]string
	FrontendInputs map[string]llb.State
	Secrets        []*schema.PackageRef
}

func MakeLocalExcludes(src LocalContents) []string {
	excludePatterns := []string{}
	excludePatterns = append(excludePatterns, dirs.BasePatternsToExclude...)
	excludePatterns = append(excludePatterns, devhost.HostOnlyFiles()...)
	excludePatterns = append(excludePatterns, src.ExcludePatterns...)

	return excludePatterns
}

func MakeLocalState(src LocalContents) llb.State {
	return llb.Local(src.Abs(),
		llb.WithCustomName(fmt.Sprintf("Workspace %s (from %s)", src.Path, src.Module.ModuleName())),
		llb.SharedKeyHint(src.Abs()),
		llb.LocalUniqueID(src.Abs()),
		llb.ExcludePatterns(MakeLocalExcludes(src)))
}

func prepareSession(ctx context.Context, keychain oci.Keychain, src secrets.GroundedSecrets, secrets []*schema.PackageRef) ([]session.Attachable, error) {
	var fs []secretsprovider.Source

	for _, def := range strings.Split(BuildkitSecrets, ";") {
		if def == "" {
			continue
		}

		parts := strings.Split(def, ":")
		if len(parts) != 3 {
			return nil, fnerrors.BadInputError("bad secret definition, expected {name}:env|file:{value}")
		}

		src := secretsprovider.Source{
			ID: parts[0],
		}

		switch parts[1] {
		case "env":
			src.Env = parts[2]
		case "file":
			src.FilePath = parts[2]
		default:
			return nil, fnerrors.BadInputError("expected env or file, got %q", parts[1])
		}

		fs = append(fs, src)
	}

	store, err := secretsprovider.NewStore(fs)
	if err != nil {
		return nil, err
	}

	secretValues := map[string][]byte{}
	if len(secrets) > 0 {
		if src == nil {
			return nil, fnerrors.InternalError("secrets specified, but secret source missing")
		}

		var errs []error
		for _, sec := range secrets {
			result, err := src.Get(ctx, sec)
			if err != nil {
				errs = append(errs, err)
				continue
			}

			if result.Value == nil {
				return nil, fnerrors.New("can't use secret %q, no value available (it's generated)", sec.Canonical())
			}

			secretValues[sec.Canonical()] = result.Value
		}

		if err := multierr.New(errs...); err != nil {
			return nil, err
		}
	}

	attachables := []session.Attachable{
		secretsprovider.NewSecretProvider(secretSource{store, secretValues}),
	}

	if ForwardKeychain {
		if keychain != nil {
			attachables = append(attachables, keychainWrapper{ctx: ctx, errorLogger: console.Output(ctx, "buildkit-auth"), keychain: keychain})
		}
	} else {
		dockerConfig := config.LoadDefaultConfigFile(console.Stderr(ctx))
		attachables = append(attachables, authprovider.NewDockerAuthProvider(dockerConfig))
	}

	// XXX make this configurable; eg at the devhost side.
	if os.Getenv("SSH_AUTH_SOCK") != "" {
		ssh, err := sshprovider.NewSSHAgentProvider([]sshprovider.AgentConfig{{ID: SSHAgentProviderID}})
		if err != nil {
			return nil, err
		}

		attachables = append(attachables, ssh)
	}

	return attachables, nil
}

type keychainWrapper struct {
	ctx         context.Context // Solve's parent context.
	errorLogger io.Writer
	keychain    oci.Keychain
}

func (kw keychainWrapper) Register(server *grpc.Server) {
	auth.RegisterAuthServer(server, kw)
}

func (kw keychainWrapper) Credentials(ctx context.Context, req *auth.CredentialsRequest) (*auth.CredentialsResponse, error) {
	response, err := kw.credentials(ctx, req.Host)

	if err == nil {
		fmt.Fprintf(console.Debug(kw.ctx), "[buildkit] AuthServer.Credentials %q --> %q\n", req.Host, response.Username)
	} else {
		fmt.Fprintf(console.Debug(kw.ctx), "[buildkit] AuthServer.Credentials %q: failed: %v\n", req.Host, err)

	}

	return response, err
}

func (kw keychainWrapper) credentials(ctx context.Context, host string) (*auth.CredentialsResponse, error) {
	// The parent context, not the incoming context is used, as the parent
	// context has an ActionSink attached (etc) while the incoming context is
	// managed by buildkit.
	authn, err := kw.keychain.Resolve(kw.ctx, resourceWrapper{host})
	if err != nil {
		return nil, err
	}

	if authn == nil {
		return &auth.CredentialsResponse{}, nil
	}

	authz, err := authn.Authorization()
	if err != nil {
		return nil, err
	}

	if authz.IdentityToken != "" || authz.RegistryToken != "" {
		fmt.Fprintf(kw.errorLogger, "%s: authentication type mismatch, got token expected username/secret", host)
		return nil, rpcerrors.Errorf(codes.InvalidArgument, "expected username/secret got token")
	}

	return &auth.CredentialsResponse{Username: authz.Username, Secret: authz.Password}, nil
}

func (kw keychainWrapper) FetchToken(ctx context.Context, req *auth.FetchTokenRequest) (*auth.FetchTokenResponse, error) {
	fmt.Fprintf(kw.errorLogger, "AuthServer.FetchToken %s\n", asJson(req))
	return nil, rpcerrors.Errorf(codes.Unimplemented, "unimplemented")
}

func (kw keychainWrapper) GetTokenAuthority(ctx context.Context, req *auth.GetTokenAuthorityRequest) (*auth.GetTokenAuthorityResponse, error) {
	fmt.Fprintf(kw.errorLogger, "AuthServer.GetTokenAuthority %s\n", asJson(req))
	return nil, rpcerrors.Errorf(codes.Unimplemented, "unimplemented")
}

func (kw keychainWrapper) VerifyTokenAuthority(ctx context.Context, req *auth.VerifyTokenAuthorityRequest) (*auth.VerifyTokenAuthorityResponse, error) {
	fmt.Fprintf(kw.errorLogger, "AuthServer.VerifyTokenAuthority %s\n", asJson(req))
	return nil, rpcerrors.Errorf(codes.Unimplemented, "unimplemented")
}

type resourceWrapper struct {
	host string
}

func (rw resourceWrapper) String() string      { return rw.host }
func (rw resourceWrapper) RegistryStr() string { return rw.host }

func asJson(msg any) string {
	str, _ := json.Marshal(msg)
	return string(str)
}
