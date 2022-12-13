// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: orchestration/proto/service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	schema "namespacelabs.dev/foundation/schema"
	orchestration "namespacelabs.dev/foundation/schema/orchestration"
	protolog "namespacelabs.dev/foundation/std/tasks/protolog"
	configuration "namespacelabs.dev/foundation/universe/aws/configuration"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeployRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plan           *schema.DeployPlan           `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	Aws            *configuration.Configuration `protobuf:"bytes,4,opt,name=aws,proto3" json:"aws,omitempty"`
	Auth           *InternalUserAuth            `protobuf:"bytes,5,opt,name=auth,proto3" json:"auth,omitempty"` // Time-limited Namespace session.
	SerializedAuth []byte                       `protobuf:"bytes,6,opt,name=serialized_auth,json=serializedAuth,proto3" json:"serialized_auth,omitempty"`
}

func (x *DeployRequest) Reset() {
	*x = DeployRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestration_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployRequest) ProtoMessage() {}

func (x *DeployRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestration_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployRequest.ProtoReflect.Descriptor instead.
func (*DeployRequest) Descriptor() ([]byte, []int) {
	return file_orchestration_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *DeployRequest) GetPlan() *schema.DeployPlan {
	if x != nil {
		return x.Plan
	}
	return nil
}

func (x *DeployRequest) GetAws() *configuration.Configuration {
	if x != nil {
		return x.Aws
	}
	return nil
}

func (x *DeployRequest) GetAuth() *InternalUserAuth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *DeployRequest) GetSerializedAuth() []byte {
	if x != nil {
		return x.SerializedAuth
	}
	return nil
}

type InternalUserAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Org      string `protobuf:"bytes,2,opt,name=org,proto3" json:"org,omitempty"`
	Opaque   []byte `protobuf:"bytes,3,opt,name=opaque,proto3" json:"opaque,omitempty"`
}

func (x *InternalUserAuth) Reset() {
	*x = InternalUserAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestration_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalUserAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalUserAuth) ProtoMessage() {}

func (x *InternalUserAuth) ProtoReflect() protoreflect.Message {
	mi := &file_orchestration_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalUserAuth.ProtoReflect.Descriptor instead.
func (*InternalUserAuth) Descriptor() ([]byte, []int) {
	return file_orchestration_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *InternalUserAuth) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *InternalUserAuth) GetOrg() string {
	if x != nil {
		return x.Org
	}
	return ""
}

func (x *InternalUserAuth) GetOpaque() []byte {
	if x != nil {
		return x.Opaque
	}
	return nil
}

type DeployResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // Deployment to follow
}

func (x *DeployResponse) Reset() {
	*x = DeployResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestration_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployResponse) ProtoMessage() {}

func (x *DeployResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestration_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployResponse.ProtoReflect.Descriptor instead.
func (*DeployResponse) Descriptor() ([]byte, []int) {
	return file_orchestration_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeployResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeploymentStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // Deployment to follow
	LogLevel int32  `protobuf:"varint,2,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
}

func (x *DeploymentStatusRequest) Reset() {
	*x = DeploymentStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestration_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentStatusRequest) ProtoMessage() {}

func (x *DeploymentStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestration_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentStatusRequest.ProtoReflect.Descriptor instead.
func (*DeploymentStatusRequest) Descriptor() ([]byte, []int) {
	return file_orchestration_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeploymentStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeploymentStatusRequest) GetLogLevel() int32 {
	if x != nil {
		return x.LogLevel
	}
	return 0
}

type DeploymentStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *orchestration.Event `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	Log   *protolog.Log        `protobuf:"bytes,4,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *DeploymentStatusResponse) Reset() {
	*x = DeploymentStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestration_proto_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentStatusResponse) ProtoMessage() {}

func (x *DeploymentStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestration_proto_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentStatusResponse.ProtoReflect.Descriptor instead.
func (*DeploymentStatusResponse) Descriptor() ([]byte, []int) {
	return file_orchestration_proto_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeploymentStatusResponse) GetEvent() *orchestration.Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *DeploymentStatusResponse) GetLog() *protolog.Log {
	if x != nil {
		return x.Log
	}
	return nil
}

type GetOrchestratorVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkipCache bool `protobuf:"varint,1,opt,name=skip_cache,json=skipCache,proto3" json:"skip_cache,omitempty"`
}

func (x *GetOrchestratorVersionRequest) Reset() {
	*x = GetOrchestratorVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestration_proto_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrchestratorVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrchestratorVersionRequest) ProtoMessage() {}

func (x *GetOrchestratorVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestration_proto_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrchestratorVersionRequest.ProtoReflect.Descriptor instead.
func (*GetOrchestratorVersionRequest) Descriptor() ([]byte, []int) {
	return file_orchestration_proto_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetOrchestratorVersionRequest) GetSkipCache() bool {
	if x != nil {
		return x.SkipCache
	}
	return false
}

type GetOrchestratorVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequiresUpdate bool `protobuf:"varint,5,opt,name=requires_update,json=requiresUpdate,proto3" json:"requires_update,omitempty"`
}

func (x *GetOrchestratorVersionResponse) Reset() {
	*x = GetOrchestratorVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestration_proto_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrchestratorVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrchestratorVersionResponse) ProtoMessage() {}

func (x *GetOrchestratorVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestration_proto_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrchestratorVersionResponse.ProtoReflect.Descriptor instead.
func (*GetOrchestratorVersionResponse) Descriptor() ([]byte, []int) {
	return file_orchestration_proto_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetOrchestratorVersionResponse) GetRequiresUpdate() bool {
	if x != nil {
		return x.RequiresUpdate
	}
	return false
}

var File_orchestration_proto_service_proto protoreflect.FileDescriptor

var file_orchestration_proto_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x73, 0x6c, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x24, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x2f, 0x61, 0x77, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x61, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x73, 0x74, 0x64, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x6c, 0x6f, 0x67, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01,
	0x0a, 0x0d, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x31, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70, 0x6c,
	0x61, 0x6e, 0x12, 0x46, 0x0a, 0x03, 0x61, 0x77, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x34, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x75, 0x6e, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x61, 0x77, 0x73, 0x12, 0x37, 0x0a, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x73, 0x6c, 0x2e, 0x6f,
	0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61,
	0x75, 0x74, 0x68, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x64, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x75, 0x74, 0x68, 0x4a, 0x04, 0x08, 0x02,
	0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x58, 0x0a, 0x10, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x72, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x72, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70,
	0x61, 0x71, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71,
	0x75, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x17, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0xa0, 0x01, 0x0a,
	0x18, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63,
	0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x03,
	0x6c, 0x6f, 0x67, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22,
	0x3e, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x6b, 0x69, 0x70, 0x43, 0x61, 0x63, 0x68, 0x65, 0x22,
	0x4f, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x05,
	0x32, 0xd3, 0x02, 0x0a, 0x14, 0x4f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x12, 0x20, 0x2e, 0x6e, 0x73, 0x6c, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x73, 0x6c, 0x2e, 0x6f, 0x72, 0x63, 0x68,
	0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x10, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x2e, 0x6e,
	0x73, 0x6c, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6e, 0x73, 0x6c, 0x2e, 0x6f,
	0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x7d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x30, 0x2e, 0x6e, 0x73, 0x6c, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6e, 0x73, 0x6c, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x63, 0x68, 0x65,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_orchestration_proto_service_proto_rawDescOnce sync.Once
	file_orchestration_proto_service_proto_rawDescData = file_orchestration_proto_service_proto_rawDesc
)

func file_orchestration_proto_service_proto_rawDescGZIP() []byte {
	file_orchestration_proto_service_proto_rawDescOnce.Do(func() {
		file_orchestration_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_orchestration_proto_service_proto_rawDescData)
	})
	return file_orchestration_proto_service_proto_rawDescData
}

var file_orchestration_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_orchestration_proto_service_proto_goTypes = []interface{}{
	(*DeployRequest)(nil),                  // 0: nsl.orchestration.DeployRequest
	(*InternalUserAuth)(nil),               // 1: nsl.orchestration.InternalUserAuth
	(*DeployResponse)(nil),                 // 2: nsl.orchestration.DeployResponse
	(*DeploymentStatusRequest)(nil),        // 3: nsl.orchestration.DeploymentStatusRequest
	(*DeploymentStatusResponse)(nil),       // 4: nsl.orchestration.DeploymentStatusResponse
	(*GetOrchestratorVersionRequest)(nil),  // 5: nsl.orchestration.GetOrchestratorVersionRequest
	(*GetOrchestratorVersionResponse)(nil), // 6: nsl.orchestration.GetOrchestratorVersionResponse
	(*schema.DeployPlan)(nil),              // 7: foundation.schema.DeployPlan
	(*configuration.Configuration)(nil),    // 8: foundation.universe.aws.configuration.Configuration
	(*orchestration.Event)(nil),            // 9: foundation.schema.orchestration.Event
	(*protolog.Log)(nil),                   // 10: foundation.workspace.tasks.protolog.Log
}
var file_orchestration_proto_service_proto_depIdxs = []int32{
	7,  // 0: nsl.orchestration.DeployRequest.plan:type_name -> foundation.schema.DeployPlan
	8,  // 1: nsl.orchestration.DeployRequest.aws:type_name -> foundation.universe.aws.configuration.Configuration
	1,  // 2: nsl.orchestration.DeployRequest.auth:type_name -> nsl.orchestration.InternalUserAuth
	9,  // 3: nsl.orchestration.DeploymentStatusResponse.event:type_name -> foundation.schema.orchestration.Event
	10, // 4: nsl.orchestration.DeploymentStatusResponse.log:type_name -> foundation.workspace.tasks.protolog.Log
	0,  // 5: nsl.orchestration.OrchestrationService.Deploy:input_type -> nsl.orchestration.DeployRequest
	3,  // 6: nsl.orchestration.OrchestrationService.DeploymentStatus:input_type -> nsl.orchestration.DeploymentStatusRequest
	5,  // 7: nsl.orchestration.OrchestrationService.GetOrchestratorVersion:input_type -> nsl.orchestration.GetOrchestratorVersionRequest
	2,  // 8: nsl.orchestration.OrchestrationService.Deploy:output_type -> nsl.orchestration.DeployResponse
	4,  // 9: nsl.orchestration.OrchestrationService.DeploymentStatus:output_type -> nsl.orchestration.DeploymentStatusResponse
	6,  // 10: nsl.orchestration.OrchestrationService.GetOrchestratorVersion:output_type -> nsl.orchestration.GetOrchestratorVersionResponse
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_orchestration_proto_service_proto_init() }
func file_orchestration_proto_service_proto_init() {
	if File_orchestration_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orchestration_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployRequest); i {
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
		file_orchestration_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalUserAuth); i {
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
		file_orchestration_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployResponse); i {
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
		file_orchestration_proto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentStatusRequest); i {
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
		file_orchestration_proto_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentStatusResponse); i {
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
		file_orchestration_proto_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrchestratorVersionRequest); i {
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
		file_orchestration_proto_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrchestratorVersionResponse); i {
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
			RawDescriptor: file_orchestration_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orchestration_proto_service_proto_goTypes,
		DependencyIndexes: file_orchestration_proto_service_proto_depIdxs,
		MessageInfos:      file_orchestration_proto_service_proto_msgTypes,
	}.Build()
	File_orchestration_proto_service_proto = out.File
	file_orchestration_proto_service_proto_rawDesc = nil
	file_orchestration_proto_service_proto_goTypes = nil
	file_orchestration_proto_service_proto_depIdxs = nil
}
