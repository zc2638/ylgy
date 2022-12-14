// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: pb/ylgy.proto

package pb

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

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameType     uint32  `protobuf:"varint,1,opt,name=gameType,proto3" json:"gameType,omitempty"`
	MapId        uint32  `protobuf:"varint,2,opt,name=mapId,proto3" json:"mapId,omitempty"`
	MapSeed      uint32  `protobuf:"varint,3,opt,name=mapSeed,proto3" json:"mapSeed,omitempty"`
	StepInfoList []*Step `protobuf:"bytes,4,rep,name=stepInfoList,proto3" json:"stepInfoList,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_ylgy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_pb_ylgy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_pb_ylgy_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetGameType() uint32 {
	if x != nil {
		return x.GameType
	}
	return 0
}

func (x *Data) GetMapId() uint32 {
	if x != nil {
		return x.MapId
	}
	return 0
}

func (x *Data) GetMapSeed() uint32 {
	if x != nil {
		return x.MapSeed
	}
	return 0
}

func (x *Data) GetStepInfoList() []*Step {
	if x != nil {
		return x.StepInfoList
	}
	return nil
}

type Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChessIndex uint32 `protobuf:"varint,1,opt,name=chessIndex,proto3" json:"chessIndex,omitempty"`
	TimeTag    uint32 `protobuf:"varint,2,opt,name=timeTag,proto3" json:"timeTag,omitempty"`
}

func (x *Step) Reset() {
	*x = Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_ylgy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Step) ProtoMessage() {}

func (x *Step) ProtoReflect() protoreflect.Message {
	mi := &file_pb_ylgy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Step.ProtoReflect.Descriptor instead.
func (*Step) Descriptor() ([]byte, []int) {
	return file_pb_ylgy_proto_rawDescGZIP(), []int{1}
}

func (x *Step) GetChessIndex() uint32 {
	if x != nil {
		return x.ChessIndex
	}
	return 0
}

func (x *Step) GetTimeTag() uint32 {
	if x != nil {
		return x.TimeTag
	}
	return 0
}

var File_pb_ylgy_proto protoreflect.FileDescriptor

var file_pb_ylgy_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x79, 0x6c, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x80, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08,
	0x67, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x67, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x70, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d, 0x61, 0x70, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x61, 0x70, 0x53, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x6d, 0x61, 0x70, 0x53, 0x65, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x0c, 0x73, 0x74, 0x65, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x0c, 0x73, 0x74, 0x65, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x40, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x68, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x63, 0x68, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x67, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_ylgy_proto_rawDescOnce sync.Once
	file_pb_ylgy_proto_rawDescData = file_pb_ylgy_proto_rawDesc
)

func file_pb_ylgy_proto_rawDescGZIP() []byte {
	file_pb_ylgy_proto_rawDescOnce.Do(func() {
		file_pb_ylgy_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_ylgy_proto_rawDescData)
	})
	return file_pb_ylgy_proto_rawDescData
}

var file_pb_ylgy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_ylgy_proto_goTypes = []interface{}{
	(*Data)(nil), // 0: pb.Data
	(*Step)(nil), // 1: pb.Step
}
var file_pb_ylgy_proto_depIdxs = []int32{
	1, // 0: pb.Data.stepInfoList:type_name -> pb.Step
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_ylgy_proto_init() }
func file_pb_ylgy_proto_init() {
	if File_pb_ylgy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_ylgy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_pb_ylgy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Step); i {
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
			RawDescriptor: file_pb_ylgy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_ylgy_proto_goTypes,
		DependencyIndexes: file_pb_ylgy_proto_depIdxs,
		MessageInfos:      file_pb_ylgy_proto_msgTypes,
	}.Build()
	File_pb_ylgy_proto = out.File
	file_pb_ylgy_proto_rawDesc = nil
	file_pb_ylgy_proto_goTypes = nil
	file_pb_ylgy_proto_depIdxs = nil
}
