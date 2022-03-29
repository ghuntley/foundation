// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
	"tailscale.com/util/multierr"
)

type AnyResolver interface {
	protoregistry.ExtensionTypeResolver
	protoregistry.MessageTypeResolver
}

func AsResolver(pr *protoregistry.Files) (AnyResolver, error) {
	ptypes := &protoregistry.Types{}

	var errs []error
	pr.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		for i := 0; i < fd.Extensions().Len(); i++ {
			if err := ptypes.RegisterExtension(dynamicpb.NewExtensionType(fd.Extensions().Get(i))); err != nil {
				errs = append(errs, err)
			}
		}

		for i := 0; i < fd.Enums().Len(); i++ {
			if err := ptypes.RegisterEnum(dynamicpb.NewEnumType(fd.Enums().Get(i))); err != nil {
				errs = append(errs, err)
			}
		}

		for i := 0; i < fd.Messages().Len(); i++ {
			if err := ptypes.RegisterMessage(dynamicpb.NewMessageType(fd.Messages().Get(i))); err != nil {
				errs = append(errs, err)
			}
		}

		return true
	})

	return ptypes, multierr.New(errs...)
}