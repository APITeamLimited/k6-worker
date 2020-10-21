// Protocol Buffers - Google's data interchange format
// Copyright 2008 Google Inc.  All rights reserved.
// https://developers.google.com/protocol-buffers/
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/source_context.proto

package sourcecontextpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

// `SourceContext` represents information about the source of a
// protobuf element, like the file in which it is defined.
type SourceContext struct ***REMOVED***
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path-qualified name of the .proto file that contained the associated
	// protobuf element.  For example: `"google/protobuf/source_context.proto"`.
	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
***REMOVED***

func (x *SourceContext) Reset() ***REMOVED***
	*x = SourceContext***REMOVED******REMOVED***
	if protoimpl.UnsafeEnabled ***REMOVED***
		mi := &file_google_protobuf_source_context_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	***REMOVED***
***REMOVED***

func (x *SourceContext) String() string ***REMOVED***
	return protoimpl.X.MessageStringOf(x)
***REMOVED***

func (*SourceContext) ProtoMessage() ***REMOVED******REMOVED***

func (x *SourceContext) ProtoReflect() protoreflect.Message ***REMOVED***
	mi := &file_google_protobuf_source_context_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil ***REMOVED***
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil ***REMOVED***
			ms.StoreMessageInfo(mi)
		***REMOVED***
		return ms
	***REMOVED***
	return mi.MessageOf(x)
***REMOVED***

// Deprecated: Use SourceContext.ProtoReflect.Descriptor instead.
func (*SourceContext) Descriptor() ([]byte, []int) ***REMOVED***
	return file_google_protobuf_source_context_proto_rawDescGZIP(), []int***REMOVED***0***REMOVED***
***REMOVED***

func (x *SourceContext) GetFileName() string ***REMOVED***
	if x != nil ***REMOVED***
		return x.FileName
	***REMOVED***
	return ""
***REMOVED***

var File_google_protobuf_source_context_proto protoreflect.FileDescriptor

var file_google_protobuf_source_context_proto_rawDesc = []byte***REMOVED***
	0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x2c, 0x0a, 0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x95, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x42, 0x12, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x3b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0xa2, 0x02, 0x03, 0x47, 0x50, 0x42, 0xaa, 0x02, 0x1e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x57,
	0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
***REMOVED***

var (
	file_google_protobuf_source_context_proto_rawDescOnce sync.Once
	file_google_protobuf_source_context_proto_rawDescData = file_google_protobuf_source_context_proto_rawDesc
)

func file_google_protobuf_source_context_proto_rawDescGZIP() []byte ***REMOVED***
	file_google_protobuf_source_context_proto_rawDescOnce.Do(func() ***REMOVED***
		file_google_protobuf_source_context_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_protobuf_source_context_proto_rawDescData)
	***REMOVED***)
	return file_google_protobuf_source_context_proto_rawDescData
***REMOVED***

var file_google_protobuf_source_context_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_protobuf_source_context_proto_goTypes = []interface***REMOVED******REMOVED******REMOVED***
	(*SourceContext)(nil), // 0: google.protobuf.SourceContext
***REMOVED***
var file_google_protobuf_source_context_proto_depIdxs = []int32***REMOVED***
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
***REMOVED***

func init() ***REMOVED*** file_google_protobuf_source_context_proto_init() ***REMOVED***
func file_google_protobuf_source_context_proto_init() ***REMOVED***
	if File_google_protobuf_source_context_proto != nil ***REMOVED***
		return
	***REMOVED***
	if !protoimpl.UnsafeEnabled ***REMOVED***
		file_google_protobuf_source_context_proto_msgTypes[0].Exporter = func(v interface***REMOVED******REMOVED***, i int) interface***REMOVED******REMOVED*** ***REMOVED***
			switch v := v.(*SourceContext); i ***REMOVED***
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			***REMOVED***
		***REMOVED***
	***REMOVED***
	type x struct***REMOVED******REMOVED***
	out := protoimpl.TypeBuilder***REMOVED***
		File: protoimpl.DescBuilder***REMOVED***
			GoPackagePath: reflect.TypeOf(x***REMOVED******REMOVED***).PkgPath(),
			RawDescriptor: file_google_protobuf_source_context_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		***REMOVED***,
		GoTypes:           file_google_protobuf_source_context_proto_goTypes,
		DependencyIndexes: file_google_protobuf_source_context_proto_depIdxs,
		MessageInfos:      file_google_protobuf_source_context_proto_msgTypes,
	***REMOVED***.Build()
	File_google_protobuf_source_context_proto = out.File
	file_google_protobuf_source_context_proto_rawDesc = nil
	file_google_protobuf_source_context_proto_goTypes = nil
	file_google_protobuf_source_context_proto_depIdxs = nil
***REMOVED***
