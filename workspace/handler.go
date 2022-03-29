// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package workspace

import (
	"context"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/schema"
)

type FrameworkExt struct {
	Include           []schema.PackageName
	FrameworkSpecific *anypb.Any
}

type ServerInputs struct {
	Services []*schema.GrpcExportService
}

// XXX we're injection Location in these, which allows for loading arbitrary files for the workspace;
// Ideally we'd pass a PackageLoader instead.
type FrameworkHandler interface {
	ParseNode(context.Context, Location, *FrameworkExt) error
	PreParseServer(context.Context, Location, *FrameworkExt) error
	PostParseServer(context.Context, *Sealed) error
	InjectService(Location, *schema.Node, *CueService) error
}

type CueService struct {
	ProtoTypename string `json:"protoTypename"`
	GoPackage     string `json:"goPackage"`
}

var (
	FrameworkHandlers = map[schema.Node_Framework]FrameworkHandler{}
)

func RegisterFrameworkHandler(framework schema.Node_Framework, handler FrameworkHandler) {
	FrameworkHandlers[framework] = handler
}

func GetExtension(extensions []*anypb.Any, msg proto.Message) bool {
	for _, ext := range extensions {
		if ext.MessageIs(msg) {
			if ext.UnmarshalTo(msg) == nil {
				return true
			}
		}
	}

	return false
}

func MustExtension(extensions []*anypb.Any, msg proto.Message) error {
	if !GetExtension(extensions, msg) {
		return fnerrors.InternalError("didn't find required extension: %s", msg.ProtoReflect().Descriptor().FullName())
	}

	return nil
}