// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: service.proto

package pb

import (
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

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa6, 0x03, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x52, 0x50, 0x43, 0x12, 0x23, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a,
	0x0b, 0x2e, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x1d,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x03, 0x2e, 0x49,
	0x44, 0x1a, 0x08, 0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x22, 0x00, 0x12, 0x22, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x03, 0x2e, 0x49, 0x44, 0x1a, 0x09, 0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x22,
	0x00, 0x12, 0x23, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x73, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x73, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x15, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x12, 0x44, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x67, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x15, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x12, 0x15, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x16, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x10, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x67, 0x61, 0x72, 0x61, 0x70, 0x68, 0x73, 0x12, 0x15,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42,
	0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*ID)(nil),                    // 0: ID
	(*SearchRequestMessage)(nil),  // 1: SearchRequestMessage
	(*Regulation)(nil),            // 2: Regulation
	(*Chapter)(nil),               // 3: Chapter
	(*Chapters)(nil),              // 4: Chapters
	(*Paragraphs)(nil),            // 5: Paragraphs
	(*SearchResponseMessage)(nil), // 6: SearchResponseMessage
}
var file_service_proto_depIdxs = []int32{
	0, // 0: RegulationGRPC.GetRegulation:input_type -> ID
	0, // 1: RegulationGRPC.GetChapter:input_type -> ID
	0, // 2: RegulationGRPC.GetAllChapters:input_type -> ID
	0, // 3: RegulationGRPC.GetParagraphs:input_type -> ID
	1, // 4: RegulationGRPC.Search:input_type -> SearchRequestMessage
	1, // 5: RegulationGRPC.SearchRegulations:input_type -> SearchRequestMessage
	1, // 6: RegulationGRPC.SearchChapters:input_type -> SearchRequestMessage
	1, // 7: RegulationGRPC.SearchPargaraphs:input_type -> SearchRequestMessage
	2, // 8: RegulationGRPC.GetRegulation:output_type -> Regulation
	3, // 9: RegulationGRPC.GetChapter:output_type -> Chapter
	4, // 10: RegulationGRPC.GetAllChapters:output_type -> Chapters
	5, // 11: RegulationGRPC.GetParagraphs:output_type -> Paragraphs
	6, // 12: RegulationGRPC.Search:output_type -> SearchResponseMessage
	6, // 13: RegulationGRPC.SearchRegulations:output_type -> SearchResponseMessage
	6, // 14: RegulationGRPC.SearchChapters:output_type -> SearchResponseMessage
	6, // 15: RegulationGRPC.SearchPargaraphs:output_type -> SearchResponseMessage
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_messages_proto_init()
	file_search_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}