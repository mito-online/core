// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: common/v1/openapi.proto

package commonv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_common_v1_openapi_proto protoreflect.FileDescriptor

var file_common_v1_openapi_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x42, 0xff, 0x05, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x92, 0x41, 0xe9, 0x04, 0x12, 0x77, 0x0a,
	0x11, 0x4d, 0x69, 0x74, 0x6f, 0x20, 0x43, 0x6f, 0x72, 0x65, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x1f, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x20, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x22, 0x3c, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x12, 0x23, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x74, 0x69, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x69, 0x74, 0x6f, 0x2d, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x1a,
	0x10, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x40, 0x6d, 0x69, 0x74, 0x6f, 0x2e, 0x63, 0x6f,
	0x6d, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x33, 0x0a,
	0x03, 0x34, 0x30, 0x30, 0x12, 0x2c, 0x0a, 0x0c, 0x42, 0x61, 0x64, 0x20, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x12, 0x1c, 0x0a, 0x1a, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x32, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x2b, 0x0a, 0x0d, 0x55, 0x6e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x2e, 0x12, 0x1a, 0x0a, 0x18, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x2f, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x28, 0x0a,
	0x0a, 0x4e, 0x6f, 0x74, 0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x12, 0x1a, 0x0a, 0x18, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0xf5, 0x01, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12,
	0xed, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x2e, 0x12, 0x1a, 0x0a, 0x18, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0xbf, 0x01,
	0x0a, 0x10, 0x58, 0x2d, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d,
	0x49, 0x64, 0x12, 0xaa, 0x01, 0x0a, 0x2b, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x20, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x12, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x32, 0x26, 0x22, 0x32, 0x34, 0x33, 0x38, 0x61, 0x63, 0x33, 0x63, 0x2d, 0x33, 0x37, 0x65, 0x62,
	0x2d, 0x34, 0x39, 0x30, 0x32, 0x2d, 0x61, 0x64, 0x65, 0x66, 0x2d, 0x65, 0x64, 0x31, 0x36, 0x62,
	0x34, 0x34, 0x33, 0x31, 0x30, 0x33, 0x30, 0x22, 0x6a, 0x45, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x41,
	0x2d, 0x46, 0x5d, 0x7b, 0x38, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b,
	0x34, 0x7d, 0x2d, 0x34, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x33, 0x7d, 0x2d,
	0x5b, 0x38, 0x39, 0x41, 0x42, 0x5d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x33,
	0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x31, 0x32, 0x7d, 0x24, 0x5a,
	0x23, 0x0a, 0x21, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x13, 0x08, 0x02, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x02, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_common_v1_openapi_proto_goTypes = []interface{}{}
var file_common_v1_openapi_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_v1_openapi_proto_init() }
func file_common_v1_openapi_proto_init() {
	if File_common_v1_openapi_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_v1_openapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_v1_openapi_proto_goTypes,
		DependencyIndexes: file_common_v1_openapi_proto_depIdxs,
	}.Build()
	File_common_v1_openapi_proto = out.File
	file_common_v1_openapi_proto_rawDesc = nil
	file_common_v1_openapi_proto_goTypes = nil
	file_common_v1_openapi_proto_depIdxs = nil
}