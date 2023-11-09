// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.11
// source: app/harmoni/api/mq/v1/like/like.proto

package v1

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

type LikeType int32

const (
	LikeType_LikeNo      LikeType = 0
	LikeType_LikePost    LikeType = 1
	LikeType_LikeComment LikeType = 2
	LikeType_LikeUser    LikeType = 3
)

// Enum value maps for LikeType.
var (
	LikeType_name = map[int32]string{
		0: "LikeNo",
		1: "LikePost",
		2: "LikeComment",
		3: "LikeUser",
	}
	LikeType_value = map[string]int32{
		"LikeNo":      0,
		"LikePost":    1,
		"LikeComment": 2,
		"LikeUser":    3,
	}
)

func (x LikeType) Enum() *LikeType {
	p := new(LikeType)
	*p = x
	return p
}

func (x LikeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LikeType) Descriptor() protoreflect.EnumDescriptor {
	return file_app_harmoni_api_mq_v1_like_like_proto_enumTypes[0].Descriptor()
}

func (LikeType) Type() protoreflect.EnumType {
	return &file_app_harmoni_api_mq_v1_like_like_proto_enumTypes[0]
}

func (x LikeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LikeType.Descriptor instead.
func (LikeType) EnumDescriptor() ([]byte, []int) {
	return file_app_harmoni_api_mq_v1_like_like_proto_rawDescGZIP(), []int{0}
}

type BaseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LikeType LikeType `protobuf:"varint,1,opt,name=likeType,proto3,enum=mq.v1.like.LikeType" json:"likeType,omitempty"`
}

func (x *BaseMessage) Reset() {
	*x = BaseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_harmoni_api_mq_v1_like_like_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseMessage) ProtoMessage() {}

func (x *BaseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_app_harmoni_api_mq_v1_like_like_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseMessage.ProtoReflect.Descriptor instead.
func (*BaseMessage) Descriptor() ([]byte, []int) {
	return file_app_harmoni_api_mq_v1_like_like_proto_rawDescGZIP(), []int{0}
}

func (x *BaseMessage) GetLikeType() LikeType {
	if x != nil {
		return x.LikeType
	}
	return LikeType_LikeNo
}

type LikeCreatedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseMessage  *BaseMessage           `protobuf:"bytes,1,opt,name=baseMessage,proto3" json:"baseMessage,omitempty"`
	UserID       int64                  `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	TargetUserID int64                  `protobuf:"varint,3,opt,name=targetUserID,proto3" json:"targetUserID,omitempty"`
	LikingID     int64                  `protobuf:"varint,4,opt,name=likingID,proto3" json:"likingID,omitempty"`
	IsCancel     bool                   `protobuf:"varint,5,opt,name=isCancel,proto3" json:"isCancel,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *LikeCreatedMessage) Reset() {
	*x = LikeCreatedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_harmoni_api_mq_v1_like_like_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeCreatedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeCreatedMessage) ProtoMessage() {}

func (x *LikeCreatedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_app_harmoni_api_mq_v1_like_like_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeCreatedMessage.ProtoReflect.Descriptor instead.
func (*LikeCreatedMessage) Descriptor() ([]byte, []int) {
	return file_app_harmoni_api_mq_v1_like_like_proto_rawDescGZIP(), []int{1}
}

func (x *LikeCreatedMessage) GetBaseMessage() *BaseMessage {
	if x != nil {
		return x.BaseMessage
	}
	return nil
}

func (x *LikeCreatedMessage) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *LikeCreatedMessage) GetTargetUserID() int64 {
	if x != nil {
		return x.TargetUserID
	}
	return 0
}

func (x *LikeCreatedMessage) GetLikingID() int64 {
	if x != nil {
		return x.LikingID
	}
	return 0
}

func (x *LikeCreatedMessage) GetIsCancel() bool {
	if x != nil {
		return x.IsCancel
	}
	return false
}

func (x *LikeCreatedMessage) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_app_harmoni_api_mq_v1_like_like_proto protoreflect.FileDescriptor

var file_app_harmoni_api_mq_v1_like_like_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x70, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x69, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x71, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x2f, 0x6c, 0x69, 0x6b,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x6c,
	0x69, 0x6b, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x6c, 0x69, 0x6b, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x6c, 0x69,
	0x6b, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6c, 0x69, 0x6b,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0xfd, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x0b,
	0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x22, 0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x69, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x73, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x38, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x43, 0x0a, 0x08, 0x4c, 0x69, 0x6b, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x69, 0x6b, 0x65, 0x4e, 0x6f, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x4c, 0x69, 0x6b, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4c,
	0x69, 0x6b, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08,
	0x4c, 0x69, 0x6b, 0x65, 0x55, 0x73, 0x65, 0x72, 0x10, 0x03, 0x42, 0x27, 0x5a, 0x25, 0x68, 0x61,
	0x72, 0x6d, 0x6f, 0x6e, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e,
	0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x71, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6b, 0x65,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_harmoni_api_mq_v1_like_like_proto_rawDescOnce sync.Once
	file_app_harmoni_api_mq_v1_like_like_proto_rawDescData = file_app_harmoni_api_mq_v1_like_like_proto_rawDesc
)

func file_app_harmoni_api_mq_v1_like_like_proto_rawDescGZIP() []byte {
	file_app_harmoni_api_mq_v1_like_like_proto_rawDescOnce.Do(func() {
		file_app_harmoni_api_mq_v1_like_like_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_harmoni_api_mq_v1_like_like_proto_rawDescData)
	})
	return file_app_harmoni_api_mq_v1_like_like_proto_rawDescData
}

var file_app_harmoni_api_mq_v1_like_like_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_app_harmoni_api_mq_v1_like_like_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_harmoni_api_mq_v1_like_like_proto_goTypes = []interface{}{
	(LikeType)(0),                 // 0: mq.v1.like.LikeType
	(*BaseMessage)(nil),           // 1: mq.v1.like.BaseMessage
	(*LikeCreatedMessage)(nil),    // 2: mq.v1.like.LikeCreatedMessage
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_app_harmoni_api_mq_v1_like_like_proto_depIdxs = []int32{
	0, // 0: mq.v1.like.BaseMessage.likeType:type_name -> mq.v1.like.LikeType
	1, // 1: mq.v1.like.LikeCreatedMessage.baseMessage:type_name -> mq.v1.like.BaseMessage
	3, // 2: mq.v1.like.LikeCreatedMessage.CreatedAt:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_app_harmoni_api_mq_v1_like_like_proto_init() }
func file_app_harmoni_api_mq_v1_like_like_proto_init() {
	if File_app_harmoni_api_mq_v1_like_like_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_harmoni_api_mq_v1_like_like_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseMessage); i {
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
		file_app_harmoni_api_mq_v1_like_like_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeCreatedMessage); i {
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
			RawDescriptor: file_app_harmoni_api_mq_v1_like_like_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_harmoni_api_mq_v1_like_like_proto_goTypes,
		DependencyIndexes: file_app_harmoni_api_mq_v1_like_like_proto_depIdxs,
		EnumInfos:         file_app_harmoni_api_mq_v1_like_like_proto_enumTypes,
		MessageInfos:      file_app_harmoni_api_mq_v1_like_like_proto_msgTypes,
	}.Build()
	File_app_harmoni_api_mq_v1_like_like_proto = out.File
	file_app_harmoni_api_mq_v1_like_like_proto_rawDesc = nil
	file_app_harmoni_api_mq_v1_like_like_proto_goTypes = nil
	file_app_harmoni_api_mq_v1_like_like_proto_depIdxs = nil
}
