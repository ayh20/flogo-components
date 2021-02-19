// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: tagDataPoint.proto

package f1telemetry2019proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TagDataPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParameterID   int32                  `protobuf:"varint,1,opt,name=parameterID,proto3" json:"parameterID,omitempty"`
	ParameterName string                 `protobuf:"bytes,2,opt,name=parameterName,proto3" json:"parameterName,omitempty"`
	Value         float64                `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	Captured      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=captured,proto3" json:"captured,omitempty"`
}

func (x *TagDataPoint) Reset() {
	*x = TagDataPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tagDataPoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagDataPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagDataPoint) ProtoMessage() {}

func (x *TagDataPoint) ProtoReflect() protoreflect.Message {
	mi := &file_tagDataPoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagDataPoint.ProtoReflect.Descriptor instead.
func (*TagDataPoint) Descriptor() ([]byte, []int) {
	return file_tagDataPoint_proto_rawDescGZIP(), []int{0}
}

func (x *TagDataPoint) GetParameterID() int32 {
	if x != nil {
		return x.ParameterID
	}
	return 0
}

func (x *TagDataPoint) GetParameterName() string {
	if x != nil {
		return x.ParameterName
	}
	return ""
}

func (x *TagDataPoint) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *TagDataPoint) GetCaptured() *timestamppb.Timestamp {
	if x != nil {
		return x.Captured
	}
	return nil
}

var File_tagDataPoint_proto protoreflect.FileDescriptor

var file_tagDataPoint_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x61, 0x67, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x74,
	0x61, 0x67, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x01, 0x0a, 0x0c, 0x54, 0x61, 0x67, 0x44,
	0x61, 0x74, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x64, 0x42, 0x08,
	0x5a, 0x06, 0x2e, 0x3b, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tagDataPoint_proto_rawDescOnce sync.Once
	file_tagDataPoint_proto_rawDescData = file_tagDataPoint_proto_rawDesc
)

func file_tagDataPoint_proto_rawDescGZIP() []byte {
	file_tagDataPoint_proto_rawDescOnce.Do(func() {
		file_tagDataPoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_tagDataPoint_proto_rawDescData)
	})
	return file_tagDataPoint_proto_rawDescData
}

var file_tagDataPoint_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tagDataPoint_proto_goTypes = []interface{}{
	(*TagDataPoint)(nil),          // 0: protobuf.tagdata.TagDataPoint
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_tagDataPoint_proto_depIdxs = []int32{
	1, // 0: protobuf.tagdata.TagDataPoint.captured:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tagDataPoint_proto_init() }
func file_tagDataPoint_proto_init() {
	if File_tagDataPoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tagDataPoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagDataPoint); i {
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
			RawDescriptor: file_tagDataPoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tagDataPoint_proto_goTypes,
		DependencyIndexes: file_tagDataPoint_proto_depIdxs,
		MessageInfos:      file_tagDataPoint_proto_msgTypes,
	}.Build()
	File_tagDataPoint_proto = out.File
	file_tagDataPoint_proto_rawDesc = nil
	file_tagDataPoint_proto_goTypes = nil
	file_tagDataPoint_proto_depIdxs = nil
}
