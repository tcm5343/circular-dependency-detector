// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: protos/filestream.proto

package protos

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

type AlloyProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *AlloyProperties) Reset() {
	*x = AlloyProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_filestream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlloyProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlloyProperties) ProtoMessage() {}

func (x *AlloyProperties) ProtoReflect() protoreflect.Message {
	mi := &file_protos_filestream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlloyProperties.ProtoReflect.Descriptor instead.
func (*AlloyProperties) Descriptor() ([]byte, []int) {
	return file_protos_filestream_proto_rawDescGZIP(), []int{0}
}

func (x *AlloyProperties) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

type FileChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content  []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Sequence int32  `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *FileChunk) Reset() {
	*x = FileChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_filestream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunk) ProtoMessage() {}

func (x *FileChunk) ProtoReflect() protoreflect.Message {
	mi := &file_protos_filestream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunk.ProtoReflect.Descriptor instead.
func (*FileChunk) Descriptor() ([]byte, []int) {
	return file_protos_filestream_proto_rawDescGZIP(), []int{1}
}

func (x *FileChunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *FileChunk) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

type FileStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//
	//	*FileStreamRequest_Props
	//	*FileStreamRequest_Chunk
	Request isFileStreamRequest_Request `protobuf_oneof:"request"`
}

func (x *FileStreamRequest) Reset() {
	*x = FileStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_filestream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileStreamRequest) ProtoMessage() {}

func (x *FileStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_filestream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileStreamRequest.ProtoReflect.Descriptor instead.
func (*FileStreamRequest) Descriptor() ([]byte, []int) {
	return file_protos_filestream_proto_rawDescGZIP(), []int{2}
}

func (m *FileStreamRequest) GetRequest() isFileStreamRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *FileStreamRequest) GetProps() *AlloyProperties {
	if x, ok := x.GetRequest().(*FileStreamRequest_Props); ok {
		return x.Props
	}
	return nil
}

func (x *FileStreamRequest) GetChunk() *FileChunk {
	if x, ok := x.GetRequest().(*FileStreamRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isFileStreamRequest_Request interface {
	isFileStreamRequest_Request()
}

type FileStreamRequest_Props struct {
	Props *AlloyProperties `protobuf:"bytes,1,opt,name=props,proto3,oneof"`
}

type FileStreamRequest_Chunk struct {
	Chunk *FileChunk `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*FileStreamRequest_Props) isFileStreamRequest_Request() {}

func (*FileStreamRequest_Chunk) isFileStreamRequest_Request() {}

type AnalysisResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AnalysisResult) Reset() {
	*x = AnalysisResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_filestream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalysisResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalysisResult) ProtoMessage() {}

func (x *AnalysisResult) ProtoReflect() protoreflect.Message {
	mi := &file_protos_filestream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalysisResult.ProtoReflect.Descriptor instead.
func (*AnalysisResult) Descriptor() ([]byte, []int) {
	return file_protos_filestream_proto_rawDescGZIP(), []int{3}
}

func (x *AnalysisResult) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_protos_filestream_proto protoreflect.FileDescriptor

var file_protos_filestream_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x2b, 0x0a, 0x0f, 0x41, 0x6c, 0x6c, 0x6f, 0x79, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x22, 0x41, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x70,
	0x72, 0x6f, 0x70, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x79, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x48, 0x00, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73,
	0x12, 0x2d, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x48, 0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42,
	0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x0e, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x32, 0x5f, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x51, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x6e, 0x64, 0x41,
	0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x12, 0x1d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x28, 0x01, 0x30, 0x01, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6d, 0x35, 0x33, 0x34, 0x33, 0x2f, 0x63, 0x69, 0x72, 0x63,
	0x75, 0x6c, 0x61, 0x72, 0x2d, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x2d,
	0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_filestream_proto_rawDescOnce sync.Once
	file_protos_filestream_proto_rawDescData = file_protos_filestream_proto_rawDesc
)

func file_protos_filestream_proto_rawDescGZIP() []byte {
	file_protos_filestream_proto_rawDescOnce.Do(func() {
		file_protos_filestream_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_filestream_proto_rawDescData)
	})
	return file_protos_filestream_proto_rawDescData
}

var file_protos_filestream_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_filestream_proto_goTypes = []any{
	(*AlloyProperties)(nil),   // 0: filestream.AlloyProperties
	(*FileChunk)(nil),         // 1: filestream.FileChunk
	(*FileStreamRequest)(nil), // 2: filestream.FileStreamRequest
	(*AnalysisResult)(nil),    // 3: filestream.AnalysisResult
}
var file_protos_filestream_proto_depIdxs = []int32{
	0, // 0: filestream.FileStreamRequest.props:type_name -> filestream.AlloyProperties
	1, // 1: filestream.FileStreamRequest.chunk:type_name -> filestream.FileChunk
	2, // 2: filestream.FileStream.UploadAndAnalyze:input_type -> filestream.FileStreamRequest
	3, // 3: filestream.FileStream.UploadAndAnalyze:output_type -> filestream.AnalysisResult
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_filestream_proto_init() }
func file_protos_filestream_proto_init() {
	if File_protos_filestream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_filestream_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AlloyProperties); i {
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
		file_protos_filestream_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FileChunk); i {
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
		file_protos_filestream_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*FileStreamRequest); i {
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
		file_protos_filestream_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AnalysisResult); i {
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
	file_protos_filestream_proto_msgTypes[2].OneofWrappers = []any{
		(*FileStreamRequest_Props)(nil),
		(*FileStreamRequest_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_filestream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_filestream_proto_goTypes,
		DependencyIndexes: file_protos_filestream_proto_depIdxs,
		MessageInfos:      file_protos_filestream_proto_msgTypes,
	}.Build()
	File_protos_filestream_proto = out.File
	file_protos_filestream_proto_rawDesc = nil
	file_protos_filestream_proto_goTypes = nil
	file_protos_filestream_proto_depIdxs = nil
}
