// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.2
// source: contents/v1/comment.proto

package contents_grpc

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

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId  int32                  `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Field     string                 `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
	UserId    int32                  `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name      string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	ContentId int32                  `protobuf:"varint,6,opt,name=content_id,json=contentId,proto3" json:"content_id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contents_v1_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_contents_v1_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_contents_v1_comment_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comment) GetParentId() int32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Comment) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Comment) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Comment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Comment) GetContentId() int32 {
	if x != nil {
		return x.ContentId
	}
	return 0
}

func (x *Comment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Comment) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CommentCreateUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentId  int32  `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Field     string `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
	UserId    int32  `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name      string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	ContentId int32  `protobuf:"varint,6,opt,name=content_id,json=contentId,proto3" json:"content_id,omitempty"`
}

func (x *CommentCreateUpdate) Reset() {
	*x = CommentCreateUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contents_v1_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentCreateUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentCreateUpdate) ProtoMessage() {}

func (x *CommentCreateUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_contents_v1_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentCreateUpdate.ProtoReflect.Descriptor instead.
func (*CommentCreateUpdate) Descriptor() ([]byte, []int) {
	return file_contents_v1_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CommentCreateUpdate) GetParentId() int32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *CommentCreateUpdate) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *CommentCreateUpdate) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CommentCreateUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommentCreateUpdate) GetContentId() int32 {
	if x != nil {
		return x.ContentId
	}
	return 0
}

var File_contents_v1_comment_proto protoreflect.FileDescriptor

var file_contents_v1_comment_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x02, 0x0a, 0x07, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x94, 0x01, 0x0a, 0x13, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x76, 0x70, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x61, 0x74, 0x61, 0x6d, 0x61, 0x74, 0x61,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d,
	0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contents_v1_comment_proto_rawDescOnce sync.Once
	file_contents_v1_comment_proto_rawDescData = file_contents_v1_comment_proto_rawDesc
)

func file_contents_v1_comment_proto_rawDescGZIP() []byte {
	file_contents_v1_comment_proto_rawDescOnce.Do(func() {
		file_contents_v1_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_contents_v1_comment_proto_rawDescData)
	})
	return file_contents_v1_comment_proto_rawDescData
}

var file_contents_v1_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_contents_v1_comment_proto_goTypes = []interface{}{
	(*Comment)(nil),               // 0: contents.v1.Comment
	(*CommentCreateUpdate)(nil),   // 1: contents.v1.CommentCreateUpdate
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_contents_v1_comment_proto_depIdxs = []int32{
	2, // 0: contents.v1.Comment.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: contents.v1.Comment.updated_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_contents_v1_comment_proto_init() }
func file_contents_v1_comment_proto_init() {
	if File_contents_v1_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contents_v1_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_contents_v1_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentCreateUpdate); i {
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
			RawDescriptor: file_contents_v1_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contents_v1_comment_proto_goTypes,
		DependencyIndexes: file_contents_v1_comment_proto_depIdxs,
		MessageInfos:      file_contents_v1_comment_proto_msgTypes,
	}.Build()
	File_contents_v1_comment_proto = out.File
	file_contents_v1_comment_proto_rawDesc = nil
	file_contents_v1_comment_proto_goTypes = nil
	file_contents_v1_comment_proto_depIdxs = nil
}
