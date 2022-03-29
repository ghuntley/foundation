// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package protos

import (
	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func LookupDescriptorProto(src *FileDescriptorSetAndDeps, typename string) (*dpb.FileDescriptorProto, *dpb.DescriptorProto) {
	t := protoreflect.FullName(typename)
	parent := t.Parent()
	name := t.Name()

	for _, p := range src.File {
		if p.GetPackage() != string(parent) {
			continue
		}

		for _, msg := range p.GetMessageType() {
			if msg.GetName() != string(name) {
				continue
			}

			return p, msg
		}
	}

	return nil, nil
}

func LookupEnumDescriptorProto(src *FileDescriptorSetAndDeps, typename string) *dpb.EnumDescriptorProto {
	if len(typename) == 0 {
		return nil
	}
	if typename[0] == '.' {
		typename = typename[1:]
	}

	t := protoreflect.FullName(typename)
	parent := t.Parent()
	name := t.Name()

	for _, p := range src.File {
		if p.GetPackage() != string(parent) {
			continue
		}

		for _, msg := range p.GetEnumType() {
			if msg.GetName() != string(name) {
				continue
			}

			return msg
		}
	}

	return nil
}