// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: schema/integration.proto

package schema

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NodejsBuild_NodePkgMgr int32

const (
	NodejsBuild_PKG_MGR_UNKNOWN NodejsBuild_NodePkgMgr = 0
	NodejsBuild_NPM             NodejsBuild_NodePkgMgr = 1
	NodejsBuild_YARN            NodejsBuild_NodePkgMgr = 2
	NodejsBuild_YARN3           NodejsBuild_NodePkgMgr = 4
	NodejsBuild_PNPM            NodejsBuild_NodePkgMgr = 3
)

// Enum value maps for NodejsBuild_NodePkgMgr.
var (
	NodejsBuild_NodePkgMgr_name = map[int32]string{
		0: "PKG_MGR_UNKNOWN",
		1: "NPM",
		2: "YARN",
		4: "YARN3",
		3: "PNPM",
	}
	NodejsBuild_NodePkgMgr_value = map[string]int32{
		"PKG_MGR_UNKNOWN": 0,
		"NPM":             1,
		"YARN":            2,
		"YARN3":           4,
		"PNPM":            3,
	}
)

func (x NodejsBuild_NodePkgMgr) Enum() *NodejsBuild_NodePkgMgr {
	p := new(NodejsBuild_NodePkgMgr)
	*p = x
	return p
}

func (x NodejsBuild_NodePkgMgr) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodejsBuild_NodePkgMgr) Descriptor() protoreflect.EnumDescriptor {
	return file_schema_integration_proto_enumTypes[0].Descriptor()
}

func (NodejsBuild_NodePkgMgr) Type() protoreflect.EnumType {
	return &file_schema_integration_proto_enumTypes[0]
}

func (x NodejsBuild_NodePkgMgr) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodejsBuild_NodePkgMgr.Descriptor instead.
func (NodejsBuild_NodePkgMgr) EnumDescriptor() ([]byte, []int) {
	return file_schema_integration_proto_rawDescGZIP(), []int{4, 0}
}

// For servers or tests.
type Integration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains the integration-specific configuration, see below.
	Data *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Integration) Reset() {
	*x = Integration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_integration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Integration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Integration) ProtoMessage() {}

func (x *Integration) ProtoReflect() protoreflect.Message {
	mi := &file_schema_integration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Integration.ProtoReflect.Descriptor instead.
func (*Integration) Descriptor() ([]byte, []int) {
	return file_schema_integration_proto_rawDescGZIP(), []int{0}
}

func (x *Integration) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type DockerfileIntegration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Dockerfile to use.
	Src string `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`
	// If set, this config is used to run the container instead of the one from the image.
	// Args and env come from the server definition.
	WorkingDir string   `protobuf:"bytes,2,opt,name=working_dir,json=workingDir,proto3" json:"working_dir,omitempty"`
	Command    []string `protobuf:"bytes,3,rep,name=command,proto3" json:"command,omitempty"`
}

func (x *DockerfileIntegration) Reset() {
	*x = DockerfileIntegration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_integration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerfileIntegration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerfileIntegration) ProtoMessage() {}

func (x *DockerfileIntegration) ProtoReflect() protoreflect.Message {
	mi := &file_schema_integration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerfileIntegration.ProtoReflect.Descriptor instead.
func (*DockerfileIntegration) Descriptor() ([]byte, []int) {
	return file_schema_integration_proto_rawDescGZIP(), []int{1}
}

func (x *DockerfileIntegration) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *DockerfileIntegration) GetWorkingDir() string {
	if x != nil {
		return x.WorkingDir
	}
	return ""
}

func (x *DockerfileIntegration) GetCommand() []string {
	if x != nil {
		return x.Command
	}
	return nil
}

type ShellScriptIntegration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entrypoint string `protobuf:"bytes,1,opt,name=entrypoint,proto3" json:"entrypoint,omitempty"`
	// Additional packages to install in the base image. By default, bash and curl are installed.
	RequiredPackages []string `protobuf:"bytes,2,rep,name=required_packages,json=requiredPackages,proto3" json:"required_packages,omitempty"`
}

func (x *ShellScriptIntegration) Reset() {
	*x = ShellScriptIntegration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_integration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShellScriptIntegration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellScriptIntegration) ProtoMessage() {}

func (x *ShellScriptIntegration) ProtoReflect() protoreflect.Message {
	mi := &file_schema_integration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShellScriptIntegration.ProtoReflect.Descriptor instead.
func (*ShellScriptIntegration) Descriptor() ([]byte, []int) {
	return file_schema_integration_proto_rawDescGZIP(), []int{2}
}

func (x *ShellScriptIntegration) GetEntrypoint() string {
	if x != nil {
		return x.Entrypoint
	}
	return ""
}

func (x *ShellScriptIntegration) GetRequiredPackages() []string {
	if x != nil {
		return x.RequiredPackages
	}
	return nil
}

type GoIntegration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pkg string `protobuf:"bytes,1,opt,name=pkg,proto3" json:"pkg,omitempty"`
}

func (x *GoIntegration) Reset() {
	*x = GoIntegration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_integration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoIntegration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoIntegration) ProtoMessage() {}

func (x *GoIntegration) ProtoReflect() protoreflect.Message {
	mi := &file_schema_integration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoIntegration.ProtoReflect.Descriptor instead.
func (*GoIntegration) Descriptor() ([]byte, []int) {
	return file_schema_integration_proto_rawDescGZIP(), []int{3}
}

func (x *GoIntegration) GetPkg() string {
	if x != nil {
		return x.Pkg
	}
	return ""
}

// Shared between "integration" and "imageFrom" for now.
type NodejsBuild struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path to `package.json`, relative to the Namespace package. Default is "."
	Pkg string `protobuf:"bytes,1,opt,name=pkg,proto3" json:"pkg,omitempty"`
	// Detected Node.js package manager.
	NodePkgMgr NodejsBuild_NodePkgMgr `protobuf:"varint,2,opt,name=node_pkg_mgr,json=nodePkgMgr,proto3,enum=foundation.schema.NodejsBuild_NodePkgMgr" json:"node_pkg_mgr,omitempty"`
	// For Web builds. It is here because building the image for Web is done by the Node.js builder.
	// This is a temporary internal API. Will be replaced with a nodejs-independent way in the future.
	InternalDoNotUseBackend []*NodejsBuild_Backend `protobuf:"bytes,5,rep,name=internal_do_not_use_backend,json=internalDoNotUseBackend,proto3" json:"internal_do_not_use_backend,omitempty"`
	// Entry point for the container.
	RunScript string `protobuf:"bytes,6,opt,name=run_script,json=runScript,proto3" json:"run_script,omitempty"`
	// Configuration for prod/test builds.
	Prod *NodejsBuild_Prod `protobuf:"bytes,7,opt,name=prod,proto3" json:"prod,omitempty"`
}

func (x *NodejsBuild) Reset() {
	*x = NodejsBuild{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_integration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodejsBuild) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodejsBuild) ProtoMessage() {}

func (x *NodejsBuild) ProtoReflect() protoreflect.Message {
	mi := &file_schema_integration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodejsBuild.ProtoReflect.Descriptor instead.
func (*NodejsBuild) Descriptor() ([]byte, []int) {
	return file_schema_integration_proto_rawDescGZIP(), []int{4}
}

func (x *NodejsBuild) GetPkg() string {
	if x != nil {
		return x.Pkg
	}
	return ""
}

func (x *NodejsBuild) GetNodePkgMgr() NodejsBuild_NodePkgMgr {
	if x != nil {
		return x.NodePkgMgr
	}
	return NodejsBuild_PKG_MGR_UNKNOWN
}

func (x *NodejsBuild) GetInternalDoNotUseBackend() []*NodejsBuild_Backend {
	if x != nil {
		return x.InternalDoNotUseBackend
	}
	return nil
}

func (x *NodejsBuild) GetRunScript() string {
	if x != nil {
		return x.RunScript
	}
	return ""
}

func (x *NodejsBuild) GetProd() *NodejsBuild_Prod {
	if x != nil {
		return x.Prod
	}
	return nil
}

type WebIntegration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodejs *NodejsBuild `protobuf:"bytes,1,opt,name=nodejs,proto3" json:"nodejs,omitempty"`
	// The service that corresponds to this web integration.
	// Needed to get the port for prod serving.
	Service string `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *WebIntegration) Reset() {
	*x = WebIntegration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_integration_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebIntegration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebIntegration) ProtoMessage() {}

func (x *WebIntegration) ProtoReflect() protoreflect.Message {
	mi := &file_schema_integration_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebIntegration.ProtoReflect.Descriptor instead.
func (*WebIntegration) Descriptor() ([]byte, []int) {
	return file_schema_integration_proto_rawDescGZIP(), []int{5}
}

func (x *WebIntegration) GetNodejs() *NodejsBuild {
	if x != nil {
		return x.Nodejs
	}
	return nil
}

func (x *WebIntegration) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type WebBuild struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodejs *NodejsBuild `protobuf:"bytes,1,opt,name=nodejs,proto3" json:"nodejs,omitempty"`
	// Passed to nginx
	Port int32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *WebBuild) Reset() {
	*x = WebBuild{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_integration_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebBuild) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebBuild) ProtoMessage() {}

func (x *WebBuild) ProtoReflect() protoreflect.Message {
	mi := &file_schema_integration_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebBuild.ProtoReflect.Descriptor instead.
func (*WebBuild) Descriptor() ([]byte, []int) {
	return file_schema_integration_proto_rawDescGZIP(), []int{6}
}

func (x *WebBuild) GetNodejs() *NodejsBuild {
	if x != nil {
		return x.Nodejs
	}
	return nil
}

func (x *WebBuild) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type NodejsBuild_Prod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If set, this script from package.json is executed.
	BuildScript string `protobuf:"bytes,1,opt,name=build_script,json=buildScript,proto3" json:"build_script,omitempty"`
	// Relative path within the build image to copy to the prod image (to the same path).
	BuildOutDir string `protobuf:"bytes,2,opt,name=build_out_dir,json=buildOutDir,proto3" json:"build_out_dir,omitempty"`
	// If true, the "install" package manager command is executed in the prod image, too.
	// Dev dependencies are not installed in this case.
	InstallDeps bool `protobuf:"varint,3,opt,name=install_deps,json=installDeps,proto3" json:"install_deps,omitempty"`
}

func (x *NodejsBuild_Prod) Reset() {
	*x = NodejsBuild_Prod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_integration_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodejsBuild_Prod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodejsBuild_Prod) ProtoMessage() {}

func (x *NodejsBuild_Prod) ProtoReflect() protoreflect.Message {
	mi := &file_schema_integration_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodejsBuild_Prod.ProtoReflect.Descriptor instead.
func (*NodejsBuild_Prod) Descriptor() ([]byte, []int) {
	return file_schema_integration_proto_rawDescGZIP(), []int{4, 0}
}

func (x *NodejsBuild_Prod) GetBuildScript() string {
	if x != nil {
		return x.BuildScript
	}
	return ""
}

func (x *NodejsBuild_Prod) GetBuildOutDir() string {
	if x != nil {
		return x.BuildOutDir
	}
	return ""
}

func (x *NodejsBuild_Prod) GetInstallDeps() bool {
	if x != nil {
		return x.InstallDeps
	}
	return false
}

type NodejsBuild_Backend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the backend for this particular Web build, e.g. "api".
	Name    string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Service *PackageRef `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// For internal use. Needed to connect to transcoded gRPC endpoints.
	Manager string `protobuf:"bytes,3,opt,name=manager,proto3" json:"manager,omitempty"`
}

func (x *NodejsBuild_Backend) Reset() {
	*x = NodejsBuild_Backend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_integration_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodejsBuild_Backend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodejsBuild_Backend) ProtoMessage() {}

func (x *NodejsBuild_Backend) ProtoReflect() protoreflect.Message {
	mi := &file_schema_integration_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodejsBuild_Backend.ProtoReflect.Descriptor instead.
func (*NodejsBuild_Backend) Descriptor() ([]byte, []int) {
	return file_schema_integration_proto_rawDescGZIP(), []int{4, 1}
}

func (x *NodejsBuild_Backend) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodejsBuild_Backend) GetService() *PackageRef {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *NodejsBuild_Backend) GetManager() string {
	if x != nil {
		return x.Manager
	}
	return ""
}

var File_schema_integration_proto protoreflect.FileDescriptor

var file_schema_integration_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37,
	0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x64, 0x0a, 0x15, 0x44, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x72, 0x63, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x44, 0x69, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x65, 0x0a,
	0x16, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x73, 0x22, 0x21, 0x0a, 0x0d, 0x47, 0x6f, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6b, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x70, 0x6b, 0x67, 0x22, 0xe5, 0x04, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65,
	0x6a, 0x73, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6b, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x6b, 0x67, 0x12, 0x4b, 0x0a, 0x0c, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x70, 0x6b, 0x67, 0x5f, 0x6d, 0x67, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x29, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x6a, 0x73, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6b, 0x67, 0x4d, 0x67, 0x72, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65,
	0x50, 0x6b, 0x67, 0x4d, 0x67, 0x72, 0x12, 0x64, 0x0a, 0x1b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x6a, 0x73, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x42, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x52, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x4e,
	0x6f, 0x74, 0x55, 0x73, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x75, 0x6e, 0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x72, 0x75, 0x6e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x70,
	0x72, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x6a, 0x73, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x04,
	0x70, 0x72, 0x6f, 0x64, 0x1a, 0x70, 0x0a, 0x04, 0x50, 0x72, 0x6f, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12,
	0x22, 0x0a, 0x0d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x64, 0x69, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x4f, 0x75, 0x74,
	0x44, 0x69, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f, 0x64,
	0x65, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x44, 0x65, 0x70, 0x73, 0x1a, 0x70, 0x0a, 0x07, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x66, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x49, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65,
	0x50, 0x6b, 0x67, 0x4d, 0x67, 0x72, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4b, 0x47, 0x5f, 0x4d, 0x47,
	0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4e,
	0x50, 0x4d, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x59, 0x41, 0x52, 0x4e, 0x10, 0x02, 0x12, 0x09,
	0x0a, 0x05, 0x59, 0x41, 0x52, 0x4e, 0x33, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4e, 0x50,
	0x4d, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x22,
	0x62, 0x0a, 0x0e, 0x57, 0x65, 0x62, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x36, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x6a, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x6a, 0x73, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x6a, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x56, 0x0a, 0x08, 0x57, 0x65, 0x62, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12,
	0x36, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x6a, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x6a, 0x73, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x6a, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x25, 0x5a, 0x23, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76,
	0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_integration_proto_rawDescOnce sync.Once
	file_schema_integration_proto_rawDescData = file_schema_integration_proto_rawDesc
)

func file_schema_integration_proto_rawDescGZIP() []byte {
	file_schema_integration_proto_rawDescOnce.Do(func() {
		file_schema_integration_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_integration_proto_rawDescData)
	})
	return file_schema_integration_proto_rawDescData
}

var file_schema_integration_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_schema_integration_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_schema_integration_proto_goTypes = []interface{}{
	(NodejsBuild_NodePkgMgr)(0),    // 0: foundation.schema.NodejsBuild.NodePkgMgr
	(*Integration)(nil),            // 1: foundation.schema.Integration
	(*DockerfileIntegration)(nil),  // 2: foundation.schema.DockerfileIntegration
	(*ShellScriptIntegration)(nil), // 3: foundation.schema.ShellScriptIntegration
	(*GoIntegration)(nil),          // 4: foundation.schema.GoIntegration
	(*NodejsBuild)(nil),            // 5: foundation.schema.NodejsBuild
	(*WebIntegration)(nil),         // 6: foundation.schema.WebIntegration
	(*WebBuild)(nil),               // 7: foundation.schema.WebBuild
	(*NodejsBuild_Prod)(nil),       // 8: foundation.schema.NodejsBuild.Prod
	(*NodejsBuild_Backend)(nil),    // 9: foundation.schema.NodejsBuild.Backend
	(*anypb.Any)(nil),              // 10: google.protobuf.Any
	(*PackageRef)(nil),             // 11: foundation.schema.PackageRef
}
var file_schema_integration_proto_depIdxs = []int32{
	10, // 0: foundation.schema.Integration.data:type_name -> google.protobuf.Any
	0,  // 1: foundation.schema.NodejsBuild.node_pkg_mgr:type_name -> foundation.schema.NodejsBuild.NodePkgMgr
	9,  // 2: foundation.schema.NodejsBuild.internal_do_not_use_backend:type_name -> foundation.schema.NodejsBuild.Backend
	8,  // 3: foundation.schema.NodejsBuild.prod:type_name -> foundation.schema.NodejsBuild.Prod
	5,  // 4: foundation.schema.WebIntegration.nodejs:type_name -> foundation.schema.NodejsBuild
	5,  // 5: foundation.schema.WebBuild.nodejs:type_name -> foundation.schema.NodejsBuild
	11, // 6: foundation.schema.NodejsBuild.Backend.service:type_name -> foundation.schema.PackageRef
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_schema_integration_proto_init() }
func file_schema_integration_proto_init() {
	if File_schema_integration_proto != nil {
		return
	}
	file_schema_package_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_schema_integration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Integration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_integration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DockerfileIntegration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_integration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShellScriptIntegration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_integration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoIntegration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_integration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodejsBuild); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_integration_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebIntegration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_integration_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebBuild); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_integration_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodejsBuild_Prod); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_integration_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodejsBuild_Backend); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_integration_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_integration_proto_goTypes,
		DependencyIndexes: file_schema_integration_proto_depIdxs,
		EnumInfos:         file_schema_integration_proto_enumTypes,
		MessageInfos:      file_schema_integration_proto_msgTypes,
	}.Build()
	File_schema_integration_proto = out.File
	file_schema_integration_proto_rawDesc = nil
	file_schema_integration_proto_goTypes = nil
	file_schema_integration_proto_depIdxs = nil
}
