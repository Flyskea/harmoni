// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: api/common/object/v1/object.proto

package v1

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

type ObjectType int32

const (
	ObjectType_OBJECT_TYPE_UNSPECIFIED ObjectType = 0
	ObjectType_OBJECT_TYPE_USER        ObjectType = 1
	ObjectType_OBJECT_TYPE_POST        ObjectType = 2
	ObjectType_OBJECT_TYPE_COMMENT     ObjectType = 3
	ObjectType_OBJECT_TYPE_TAG         ObjectType = 4
)

// Enum value maps for ObjectType.
var (
	ObjectType_name = map[int32]string{
		0: "OBJECT_TYPE_UNSPECIFIED",
		1: "OBJECT_TYPE_USER",
		2: "OBJECT_TYPE_POST",
		3: "OBJECT_TYPE_COMMENT",
		4: "OBJECT_TYPE_TAG",
	}
	ObjectType_value = map[string]int32{
		"OBJECT_TYPE_UNSPECIFIED": 0,
		"OBJECT_TYPE_USER":        1,
		"OBJECT_TYPE_POST":        2,
		"OBJECT_TYPE_COMMENT":     3,
		"OBJECT_TYPE_TAG":         4,
	}
)

func (x ObjectType) Enum() *ObjectType {
	p := new(ObjectType)
	*p = x
	return p
}

func (x ObjectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjectType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_common_object_v1_object_proto_enumTypes[0].Descriptor()
}

func (ObjectType) Type() protoreflect.EnumType {
	return &file_api_common_object_v1_object_proto_enumTypes[0]
}

func (x ObjectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjectType.Descriptor instead.
func (ObjectType) EnumDescriptor() ([]byte, []int) {
	return file_api_common_object_v1_object_proto_rawDescGZIP(), []int{0}
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type ObjectType `protobuf:"varint,2,opt,name=type,proto3,enum=object.ObjectType" json:"type,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_object_v1_object_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_object_v1_object_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_api_common_object_v1_object_proto_rawDescGZIP(), []int{0}
}

func (x *Object) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Object) GetType() ObjectType {
	if x != nil {
		return x.Type
	}
	return ObjectType_OBJECT_TYPE_UNSPECIFIED
}

var File_api_common_object_v1_object_proto protoreflect.FileDescriptor

var file_api_common_object_v1_object_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x40, 0x0a, 0x06, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x83, 0x01,
	0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17,
	0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4f, 0x42, 0x4a,
	0x45, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x12,
	0x14, 0x0a, 0x10, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50,
	0x4f, 0x53, 0x54, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x13,
	0x0a, 0x0f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x41,
	0x47, 0x10, 0x04, 0x42, 0x21, 0x5a, 0x1f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_common_object_v1_object_proto_rawDescOnce sync.Once
	file_api_common_object_v1_object_proto_rawDescData = file_api_common_object_v1_object_proto_rawDesc
)

func file_api_common_object_v1_object_proto_rawDescGZIP() []byte {
	file_api_common_object_v1_object_proto_rawDescOnce.Do(func() {
		file_api_common_object_v1_object_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_common_object_v1_object_proto_rawDescData)
	})
	return file_api_common_object_v1_object_proto_rawDescData
}

var file_api_common_object_v1_object_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_common_object_v1_object_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_common_object_v1_object_proto_goTypes = []interface{}{
	(ObjectType)(0), // 0: object.ObjectType
	(*Object)(nil),  // 1: object.Object
}
var file_api_common_object_v1_object_proto_depIdxs = []int32{
	0, // 0: object.Object.type:type_name -> object.ObjectType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_common_object_v1_object_proto_init() }
func file_api_common_object_v1_object_proto_init() {
	if File_api_common_object_v1_object_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_common_object_v1_object_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
			RawDescriptor: file_api_common_object_v1_object_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_common_object_v1_object_proto_goTypes,
		DependencyIndexes: file_api_common_object_v1_object_proto_depIdxs,
		EnumInfos:         file_api_common_object_v1_object_proto_enumTypes,
		MessageInfos:      file_api_common_object_v1_object_proto_msgTypes,
	}.Build()
	File_api_common_object_v1_object_proto = out.File
	file_api_common_object_v1_object_proto_rawDesc = nil
	file_api_common_object_v1_object_proto_goTypes = nil
	file_api_common_object_v1_object_proto_depIdxs = nil
}
