// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: google/devtools/cloudtrace/v2/tracing.proto

package tracepb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message for the `BatchWriteSpans` method.
type BatchWriteSpansRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the project where the spans belong. The format is
	// `projects/[PROJECT_ID]`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. A list of new spans. The span names must not match existing
	// spans, otherwise the results are undefined.
	Spans []*Span `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
}

func (x *BatchWriteSpansRequest) Reset() {
	*x = BatchWriteSpansRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_cloudtrace_v2_tracing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchWriteSpansRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchWriteSpansRequest) ProtoMessage() {}

func (x *BatchWriteSpansRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_cloudtrace_v2_tracing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchWriteSpansRequest.ProtoReflect.Descriptor instead.
func (*BatchWriteSpansRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_cloudtrace_v2_tracing_proto_rawDescGZIP(), []int{0}
}

func (x *BatchWriteSpansRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BatchWriteSpansRequest) GetSpans() []*Span {
	if x != nil {
		return x.Spans
	}
	return nil
}

var File_google_devtools_cloudtrace_v2_tracing_proto protoreflect.FileDescriptor

var file_google_devtools_cloudtrace_v2_tracing_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x32, 0x2f,
	0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x16, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x47, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x33, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x73,
	0x70, 0x61, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x32, 0xba, 0x03, 0x0a, 0x0c,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa1, 0x01, 0x0a,
	0x0f, 0x42, 0x61, 0x74, 0x63, 0x68, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x73,
	0x12, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x32,
	0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x3f, 0xda, 0x41, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x76, 0x32, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x73, 0x3a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x12, 0x89, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x12,
	0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x53, 0x70, 0x61, 0x6e, 0x1a, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65,
	0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2b, 0x3a, 0x01, 0x2a, 0x22, 0x26, 0x2f, 0x76, 0x32, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x73, 0x2f, 0x2a, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x1a, 0x7a, 0xca, 0x41,
	0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x5b, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2c, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x42, 0xaf, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x32, 0x42, 0x0c,
	0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x70, 0x62, 0x3b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x70, 0x62, 0xaa,
	0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x5c, 0x56, 0x32, 0xea,
	0x02, 0x18, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x54, 0x72, 0x61, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_devtools_cloudtrace_v2_tracing_proto_rawDescOnce sync.Once
	file_google_devtools_cloudtrace_v2_tracing_proto_rawDescData = file_google_devtools_cloudtrace_v2_tracing_proto_rawDesc
)

func file_google_devtools_cloudtrace_v2_tracing_proto_rawDescGZIP() []byte {
	file_google_devtools_cloudtrace_v2_tracing_proto_rawDescOnce.Do(func() {
		file_google_devtools_cloudtrace_v2_tracing_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_cloudtrace_v2_tracing_proto_rawDescData)
	})
	return file_google_devtools_cloudtrace_v2_tracing_proto_rawDescData
}

var file_google_devtools_cloudtrace_v2_tracing_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_devtools_cloudtrace_v2_tracing_proto_goTypes = []interface{}{
	(*BatchWriteSpansRequest)(nil), // 0: google.devtools.cloudtrace.v2.BatchWriteSpansRequest
	(*Span)(nil),                   // 1: google.devtools.cloudtrace.v2.Span
	(*emptypb.Empty)(nil),          // 2: google.protobuf.Empty
}
var file_google_devtools_cloudtrace_v2_tracing_proto_depIdxs = []int32{
	1, // 0: google.devtools.cloudtrace.v2.BatchWriteSpansRequest.spans:type_name -> google.devtools.cloudtrace.v2.Span
	0, // 1: google.devtools.cloudtrace.v2.TraceService.BatchWriteSpans:input_type -> google.devtools.cloudtrace.v2.BatchWriteSpansRequest
	1, // 2: google.devtools.cloudtrace.v2.TraceService.CreateSpan:input_type -> google.devtools.cloudtrace.v2.Span
	2, // 3: google.devtools.cloudtrace.v2.TraceService.BatchWriteSpans:output_type -> google.protobuf.Empty
	1, // 4: google.devtools.cloudtrace.v2.TraceService.CreateSpan:output_type -> google.devtools.cloudtrace.v2.Span
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_devtools_cloudtrace_v2_tracing_proto_init() }
func file_google_devtools_cloudtrace_v2_tracing_proto_init() {
	if File_google_devtools_cloudtrace_v2_tracing_proto != nil {
		return
	}
	file_google_devtools_cloudtrace_v2_trace_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_cloudtrace_v2_tracing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchWriteSpansRequest); i {
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
			RawDescriptor: file_google_devtools_cloudtrace_v2_tracing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_devtools_cloudtrace_v2_tracing_proto_goTypes,
		DependencyIndexes: file_google_devtools_cloudtrace_v2_tracing_proto_depIdxs,
		MessageInfos:      file_google_devtools_cloudtrace_v2_tracing_proto_msgTypes,
	}.Build()
	File_google_devtools_cloudtrace_v2_tracing_proto = out.File
	file_google_devtools_cloudtrace_v2_tracing_proto_rawDesc = nil
	file_google_devtools_cloudtrace_v2_tracing_proto_goTypes = nil
	file_google_devtools_cloudtrace_v2_tracing_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TraceServiceClient is the client API for TraceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceServiceClient interface {
	// Batch writes new spans to new or existing traces. You cannot update
	// existing spans.
	BatchWriteSpans(ctx context.Context, in *BatchWriteSpansRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Creates a new span.
	CreateSpan(ctx context.Context, in *Span, opts ...grpc.CallOption) (*Span, error)
}

type traceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTraceServiceClient(cc grpc.ClientConnInterface) TraceServiceClient {
	return &traceServiceClient{cc}
}

func (c *traceServiceClient) BatchWriteSpans(ctx context.Context, in *BatchWriteSpansRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/google.devtools.cloudtrace.v2.TraceService/BatchWriteSpans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceServiceClient) CreateSpan(ctx context.Context, in *Span, opts ...grpc.CallOption) (*Span, error) {
	out := new(Span)
	err := c.cc.Invoke(ctx, "/google.devtools.cloudtrace.v2.TraceService/CreateSpan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TraceServiceServer is the server API for TraceService service.
type TraceServiceServer interface {
	// Batch writes new spans to new or existing traces. You cannot update
	// existing spans.
	BatchWriteSpans(context.Context, *BatchWriteSpansRequest) (*emptypb.Empty, error)
	// Creates a new span.
	CreateSpan(context.Context, *Span) (*Span, error)
}

// UnimplementedTraceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTraceServiceServer struct {
}

func (*UnimplementedTraceServiceServer) BatchWriteSpans(context.Context, *BatchWriteSpansRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchWriteSpans not implemented")
}
func (*UnimplementedTraceServiceServer) CreateSpan(context.Context, *Span) (*Span, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSpan not implemented")
}

func RegisterTraceServiceServer(s *grpc.Server, srv TraceServiceServer) {
	s.RegisterService(&_TraceService_serviceDesc, srv)
}

func _TraceService_BatchWriteSpans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchWriteSpansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceServiceServer).BatchWriteSpans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.cloudtrace.v2.TraceService/BatchWriteSpans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceServiceServer).BatchWriteSpans(ctx, req.(*BatchWriteSpansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceService_CreateSpan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Span)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceServiceServer).CreateSpan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.cloudtrace.v2.TraceService/CreateSpan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceServiceServer).CreateSpan(ctx, req.(*Span))
	}
	return interceptor(ctx, in, info, handler)
}

var _TraceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.cloudtrace.v2.TraceService",
	HandlerType: (*TraceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BatchWriteSpans",
			Handler:    _TraceService_BatchWriteSpans_Handler,
		},
		{
			MethodName: "CreateSpan",
			Handler:    _TraceService_CreateSpan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/cloudtrace/v2/tracing.proto",
}
