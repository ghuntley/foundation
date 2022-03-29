// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package nodejs

import (
	"context"
	"errors"
	"io"
	"reflect"

	"google.golang.org/protobuf/proto"
	"namespacelabs.dev/foundation/build"
	"namespacelabs.dev/foundation/internal/engine/ops"
	"namespacelabs.dev/foundation/internal/engine/ops/defs"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/fnfs"
	"namespacelabs.dev/foundation/internal/production"
	"namespacelabs.dev/foundation/languages"
	"namespacelabs.dev/foundation/provision"
	"namespacelabs.dev/foundation/runtime"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/workspace"
	"namespacelabs.dev/foundation/workspace/source/protos"
)

func Register() {
	languages.Register(schema.Node_NODEJS_GRPC, impl{})
	languages.Register(schema.Node_NODEJS, plainImpl{})

	ops.Register(&OpGenServer{}, generator{})
}

type generator struct{}

func (generator) Run(ctx context.Context, env ops.Environment, _ *schema.Definition, msg proto.Message) (*ops.DispatcherResult, error) {
	wenv, ok := env.(workspace.Packages)
	if !ok {
		return nil, errors.New("workspace.Packages required")
	}

	switch x := msg.(type) {
	case *OpGenServer:
		loc, err := wenv.Resolve(ctx, schema.PackageName(x.Server.PackageName))
		if err != nil {
			return nil, err
		}

		return nil, fnfs.WriteWorkspaceFile(ctx, loc.Module.ReadWriteFS(), loc.Rel("main.fn.ts"), func(w io.Writer) error {
			f, err := lib.Open("main.ts")
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(w, f)
			return err
		})

	default:
		return nil, fnerrors.InternalError("unsupported type: %s", reflect.TypeOf(x).String())
	}
}

type impl struct {
	languages.MaybeGenerate
	languages.MaybeTidy
	languages.NoDev
}

func (impl) PrepareBuild(ctx context.Context, _ languages.Endpoints, server provision.Server, isFocus bool) (build.Spec, error) {
	return buildNodeJS{loc: server.Location, isFocus: isFocus}, nil
}

func (impl) PrepareRun(ctx context.Context, t provision.Server, run *runtime.ServerRunOpts) error {
	run.Command = []string{"node", "main.fn.js"}
	run.WorkingDir = "/app"
	run.ReadOnlyFilesystem = true
	run.RunAs = production.NonRootRunAsWithID(production.NonRootUserID)
	return nil
}

func (impl) TidyServer(ctx context.Context, loc workspace.Location, server *schema.Server) error {
	packages := []string{
		"typescript@4.5.4",
		"ts-node@10.4.x",
		"@grpc/grpc-js@1.5.1",
		"yargs@16.x",
	}

	devPackages := []string{
		"@types/yargs@16.x",
		"@types/node@16.x",
		"@tsconfig/node16@1.0.2",
	}

	if err := RunYarn(ctx, loc, append([]string{"add"}, packages...)); err != nil {
		return err
	}

	if err := RunYarn(ctx, loc, append([]string{"add", "-D"}, devPackages...)); err != nil {
		return err
	}

	return nil
}

func (impl) GenerateServer(pkg *workspace.Package, nodes []*schema.Node) ([]*schema.Definition, error) {
	var dl defs.DefList
	dl.Add("Generate Typescript server dependencies", &OpGenServer{Server: pkg.Server}, pkg.PackageName())
	return dl.Serialize()
}

func (impl) ParseNode(ctx context.Context, loc workspace.Location, ext *workspace.FrameworkExt) error {
	return nil
}

func (impl) PreParseServer(ctx context.Context, loc workspace.Location, ext *workspace.FrameworkExt) error {
	ext.Include = append(ext.Include, "namespacelabs.dev/foundation/std/nodejs/grpc")
	return nil
}

func (impl) PostParseServer(ctx context.Context, _ *workspace.Sealed) error {
	return nil
}

func (impl) InjectService(loc workspace.Location, node *schema.Node, svc *workspace.CueService) error {
	return nil
}

func (impl) GenerateNode(pkg *workspace.Package, nodes []*schema.Node) ([]*schema.Definition, error) {
	var dl defs.DefList

	var list []*protos.FileDescriptorSetAndDeps
	for _, dl := range pkg.Provides {
		list = append(list, dl)
	}
	for _, svc := range pkg.Services {
		list = append(list, svc)
	}

	// TODO(#348): enable when we figure out where to store the codegen.
	// if len(list) > 0 {
	// 	dl.Add("Generate Javascript/Typescript proto sources", &source.OpProtoGen{
	// 		PackageName:         pkg.PackageName().String(),
	// 		GenerateHttpGateway: pkg.Node().ExportServicesAsHttp,
	// 		Protos:              protos.Merge(list...),
	// 		Framework:           source.OpProtoGen_TYPESCRIPT,
	// 	})
	// }

	return dl.Serialize()
}

type plainImpl struct {
	languages.MaybeGenerate
	languages.MaybeTidy
	languages.NoDev
}

func (plainImpl) PrepareBuild(ctx context.Context, _ languages.Endpoints, server provision.Server, isFocus bool) (build.Spec, error) {
	return buildNodeJS{loc: server.Location, isFocus: isFocus}, nil
}

func (plainImpl) PrepareRun(ctx context.Context, t provision.Server, run *runtime.ServerRunOpts) error {
	run.Command = []string{"yarn", "serve"}
	run.WorkingDir = "/app"
	run.Env = map[string]string{
		"NODE_ENV": nodeEnv(t.Env()),
	}
	return nil
}

func (plainImpl) ParseNode(ctx context.Context, loc workspace.Location, ext *workspace.FrameworkExt) error {
	return nil
}

func (plainImpl) PreParseServer(ctx context.Context, loc workspace.Location, ext *workspace.FrameworkExt) error {
	return nil
}

func (plainImpl) PostParseServer(ctx context.Context, sealed *workspace.Sealed) error {
	sealed.Proto.Server.StaticPort = []*schema.Endpoint_Port{{Name: "http-port", ContainerPort: 8080}}
	return nil
}

func (plainImpl) InjectService(loc workspace.Location, node *schema.Node, svc *workspace.CueService) error {
	return nil
}