package f1telemetry2019proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NameValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *NameValue) Reset() {
	*x = NameValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nameValue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameValue) ProtoMessage() {}

func (x *NameValue) ProtoReflect() protoreflect.Message {
	mi := &file_nameValue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameValue.ProtoReflect.Descriptor instead.
func (*NameValue) Descriptor() ([]byte, []int) {
	return file_nameValue_proto_rawDescGZIP(), []int{0}
}

func (x *NameValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NameValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_nameValue_proto protoreflect.FileDescriptor

var file_nameValue_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6e, 0x61, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x22, 0x35, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6d,
	0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nameValue_proto_rawDescOnce sync.Once
	file_nameValue_proto_rawDescData = file_nameValue_proto_rawDesc
)

func file_nameValue_proto_rawDescGZIP() []byte {
	file_nameValue_proto_rawDescOnce.Do(func() {
		file_nameValue_proto_rawDescData = protoimpl.X.CompressGZIP(file_nameValue_proto_rawDescData)
	})
	return file_nameValue_proto_rawDescData
}

var file_nameValue_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_nameValue_proto_goTypes = []interface{}{
	(*NameValue)(nil), // 0: protobuf.common.NameValue
}
var file_nameValue_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nameValue_proto_init() }
func file_nameValue_proto_init() {
	if File_nameValue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nameValue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameValue); i {
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
			RawDescriptor: file_nameValue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nameValue_proto_goTypes,
		DependencyIndexes: file_nameValue_proto_depIdxs,
		MessageInfos:      file_nameValue_proto_msgTypes,
	}.Build()
	File_nameValue_proto = out.File
	file_nameValue_proto_rawDesc = nil
	file_nameValue_proto_goTypes = nil
	file_nameValue_proto_depIdxs = nil
}
