// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package api

import (
	"context"

	"google.golang.org/protobuf/proto"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/pkggraph"
)

func SetServerBinary(pkg *pkggraph.Package, buildPlan *schema.LayeredImageBuildPlan, commands []string) error {
	pkg.Binaries = append(pkg.Binaries, &schema.Binary{
		Name:      pkg.Server.Name,
		BuildPlan: buildPlan,
		Config: &schema.BinaryConfig{
			Command: commands,
		},
	})

	return SetServerBinaryRef(pkg, schema.MakePackageRef(pkg.Location.PackageName, pkg.Server.Name))
}

func SetServerBinaryRef(pkg *pkggraph.Package, binaryRef *schema.PackageRef) error {
	if pkg.Server.Self == nil {
		pkg.Server.Self = &schema.ServerFragment{
			MainContainer: &schema.Container{},
		}
	} else if pkg.Server.Self.MainContainer == nil {
		pkg.Server.Self.MainContainer = &schema.Container{}
	}

	if pkg.Server.Self.MainContainer.BinaryRef != nil {
		// TODO: add a more meaningful error message
		return fnerrors.NewWithLocation(pkg.Location, "server binary is set multiple times")
	}

	pkg.Server.Self.MainContainer.BinaryRef = binaryRef

	return nil
}

func GenerateBinaryAndAddToPackage(ctx context.Context, env *schema.Environment, pl pkggraph.PackageLoader, pkg *pkggraph.Package, binaryName string, data proto.Message) (*schema.PackageRef, error) {
	binary, err := GenerateBinary(ctx, env, pl, pkg.Location, binaryName, data)
	if err != nil {
		return nil, err
	}

	pkg.Binaries = append(pkg.Binaries, binary)

	return schema.MakePackageRef(pkg.Location.PackageName, binaryName), nil
}

// Simply creates a binary for a test
type DefaultBinaryTestIntegration[ServerData proto.Message] struct{}

func (DefaultBinaryTestIntegration[ServerData]) ApplyToTest(ctx context.Context, env *schema.Environment, pl pkggraph.PackageLoader, pkg *pkggraph.Package, test *schema.Test, data ServerData) error {
	ref, err := GenerateBinaryAndAddToPackage(ctx, env, pl, pkg, test.Name, data)
	if err != nil {
		return err
	}

	test.Driver = ref

	return nil

}
