// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: internal/runtime/kubernetes/networking/ingress/op.proto

package ingress

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OpMapAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fdqn string `protobuf:"bytes,1,opt,name=fdqn,proto3" json:"fdqn,omitempty"`
	// If specified, will map to the IP address of the LoadBalancer resolved.
	IngressNs   string `protobuf:"bytes,2,opt,name=ingress_ns,json=ingressNs,proto3" json:"ingress_ns,omitempty"`
	IngressName string `protobuf:"bytes,3,opt,name=ingress_name,json=ingressName,proto3" json:"ingress_name,omitempty"`
	// If specified, will map to the specified target.
	CnameTarget string `protobuf:"bytes,4,opt,name=cname_target,json=cnameTarget,proto3" json:"cname_target,omitempty"`
}

func (x *OpMapAddress) Reset() {
	*x = OpMapAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_runtime_kubernetes_networking_ingress_op_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpMapAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpMapAddress) ProtoMessage() {}

func (x *OpMapAddress) ProtoReflect() protoreflect.Message {
	mi := &file_internal_runtime_kubernetes_networking_ingress_op_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpMapAddress.ProtoReflect.Descriptor instead.
func (*OpMapAddress) Descriptor() ([]byte, []int) {
	return file_internal_runtime_kubernetes_networking_ingress_op_proto_rawDescGZIP(), []int{0}
}

func (x *OpMapAddress) GetFdqn() string {
	if x != nil {
		return x.Fdqn
	}
	return ""
}

func (x *OpMapAddress) GetIngressNs() string {
	if x != nil {
		return x.IngressNs
	}
	return ""
}

func (x *OpMapAddress) GetIngressName() string {
	if x != nil {
		return x.IngressName
	}
	return ""
}

func (x *OpMapAddress) GetCnameTarget() string {
	if x != nil {
		return x.CnameTarget
	}
	return ""
}

type OpCleanupMigration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *OpCleanupMigration) Reset() {
	*x = OpCleanupMigration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_runtime_kubernetes_networking_ingress_op_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpCleanupMigration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpCleanupMigration) ProtoMessage() {}

func (x *OpCleanupMigration) ProtoReflect() protoreflect.Message {
	mi := &file_internal_runtime_kubernetes_networking_ingress_op_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpCleanupMigration.ProtoReflect.Descriptor instead.
func (*OpCleanupMigration) Descriptor() ([]byte, []int) {
	return file_internal_runtime_kubernetes_networking_ingress_op_proto_rawDescGZIP(), []int{1}
}

func (x *OpCleanupMigration) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_internal_runtime_kubernetes_networking_ingress_op_proto protoreflect.FileDescriptor

var file_internal_runtime_kubernetes_networking_ingress_op_proto_rawDesc = []byte{
	0x0a, 0x37, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x2f, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x30, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x0c,
	0x4f, 0x70, 0x4d, 0x61, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x64, 0x71, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x64, 0x71, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4e, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6e, 0x61, 0x6d, 0x65, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x32, 0x0a, 0x12, 0x4f, 0x70, 0x43, 0x6c, 0x65, 0x61, 0x6e,
	0x75, 0x70, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x4d, 0x5a, 0x4b, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_runtime_kubernetes_networking_ingress_op_proto_rawDescOnce sync.Once
	file_internal_runtime_kubernetes_networking_ingress_op_proto_rawDescData = file_internal_runtime_kubernetes_networking_ingress_op_proto_rawDesc
)

func file_internal_runtime_kubernetes_networking_ingress_op_proto_rawDescGZIP() []byte {
	file_internal_runtime_kubernetes_networking_ingress_op_proto_rawDescOnce.Do(func() {
		file_internal_runtime_kubernetes_networking_ingress_op_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_runtime_kubernetes_networking_ingress_op_proto_rawDescData)
	})
	return file_internal_runtime_kubernetes_networking_ingress_op_proto_rawDescData
}

var file_internal_runtime_kubernetes_networking_ingress_op_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_runtime_kubernetes_networking_ingress_op_proto_goTypes = []interface{}{
	(*OpMapAddress)(nil),       // 0: foundation.runtime.kubernetes.networking.ingress.OpMapAddress
	(*OpCleanupMigration)(nil), // 1: foundation.runtime.kubernetes.networking.ingress.OpCleanupMigration
}
var file_internal_runtime_kubernetes_networking_ingress_op_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_runtime_kubernetes_networking_ingress_op_proto_init() }
func file_internal_runtime_kubernetes_networking_ingress_op_proto_init() {
	if File_internal_runtime_kubernetes_networking_ingress_op_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_runtime_kubernetes_networking_ingress_op_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpMapAddress); i {
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
		file_internal_runtime_kubernetes_networking_ingress_op_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpCleanupMigration); i {
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
			RawDescriptor: file_internal_runtime_kubernetes_networking_ingress_op_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_runtime_kubernetes_networking_ingress_op_proto_goTypes,
		DependencyIndexes: file_internal_runtime_kubernetes_networking_ingress_op_proto_depIdxs,
		MessageInfos:      file_internal_runtime_kubernetes_networking_ingress_op_proto_msgTypes,
	}.Build()
	File_internal_runtime_kubernetes_networking_ingress_op_proto = out.File
	file_internal_runtime_kubernetes_networking_ingress_op_proto_rawDesc = nil
	file_internal_runtime_kubernetes_networking_ingress_op_proto_goTypes = nil
	file_internal_runtime_kubernetes_networking_ingress_op_proto_depIdxs = nil
}
