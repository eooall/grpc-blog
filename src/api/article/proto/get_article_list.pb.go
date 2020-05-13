// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: get_article_list.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ArticleIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleType string `protobuf:"bytes,1,opt,name=articleType,proto3" json:"articleType,omitempty"`  //类型
	PageCurrent int64  `protobuf:"varint,2,opt,name=pageCurrent,proto3" json:"pageCurrent,omitempty"` //当前页
	Search      string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`            //搜索
}

func (x *ArticleIn) Reset() {
	*x = ArticleIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_article_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleIn) ProtoMessage() {}

func (x *ArticleIn) ProtoReflect() protoreflect.Message {
	mi := &file_get_article_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleIn.ProtoReflect.Descriptor instead.
func (*ArticleIn) Descriptor() ([]byte, []int) {
	return file_get_article_list_proto_rawDescGZIP(), []int{0}
}

func (x *ArticleIn) GetArticleType() string {
	if x != nil {
		return x.ArticleType
	}
	return ""
}

func (x *ArticleIn) GetPageCurrent() int64 {
	if x != nil {
		return x.PageCurrent
	}
	return 0
}

func (x *ArticleIn) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type ArticleListOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *ArticleListData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ArticleListOut) Reset() {
	*x = ArticleListOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_article_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleListOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleListOut) ProtoMessage() {}

func (x *ArticleListOut) ProtoReflect() protoreflect.Message {
	mi := &file_get_article_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleListOut.ProtoReflect.Descriptor instead.
func (*ArticleListOut) Descriptor() ([]byte, []int) {
	return file_get_article_list_proto_rawDescGZIP(), []int{1}
}

func (x *ArticleListOut) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ArticleListOut) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ArticleListOut) GetData() *ArticleListData {
	if x != nil {
		return x.Data
	}
	return nil
}

//文章列表定义
type ArticleListData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageCount int64          `protobuf:"varint,1,opt,name=pageCount,proto3" json:"pageCount,omitempty"` // 文章总页数
	PageLimit int64          `protobuf:"varint,2,opt,name=pageLimit,proto3" json:"pageLimit,omitempty"`
	Article   []*ArticleData `protobuf:"bytes,3,rep,name=article,proto3" json:"article,omitempty"`
}

func (x *ArticleListData) Reset() {
	*x = ArticleListData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_article_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleListData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleListData) ProtoMessage() {}

func (x *ArticleListData) ProtoReflect() protoreflect.Message {
	mi := &file_get_article_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleListData.ProtoReflect.Descriptor instead.
func (*ArticleListData) Descriptor() ([]byte, []int) {
	return file_get_article_list_proto_rawDescGZIP(), []int{2}
}

func (x *ArticleListData) GetPageCount() int64 {
	if x != nil {
		return x.PageCount
	}
	return 0
}

func (x *ArticleListData) GetPageLimit() int64 {
	if x != nil {
		return x.PageLimit
	}
	return 0
}

func (x *ArticleListData) GetArticle() []*ArticleData {
	if x != nil {
		return x.Article
	}
	return nil
}

type ArticleData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`              //文章名
	Account    string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`        //作者
	Time       string `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`              //创建时间
	CommentNum int64  `protobuf:"varint,4,opt,name=commentNum,proto3" json:"commentNum,omitempty"` //评论数量
	ArticleId  int64  `protobuf:"varint,6,opt,name=articleId,proto3" json:"articleId,omitempty"`   //文章id
}

func (x *ArticleData) Reset() {
	*x = ArticleData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_article_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleData) ProtoMessage() {}

func (x *ArticleData) ProtoReflect() protoreflect.Message {
	mi := &file_get_article_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleData.ProtoReflect.Descriptor instead.
func (*ArticleData) Descriptor() ([]byte, []int) {
	return file_get_article_list_proto_rawDescGZIP(), []int{3}
}

func (x *ArticleData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ArticleData) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *ArticleData) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *ArticleData) GetCommentNum() int64 {
	if x != nil {
		return x.CommentNum
	}
	return 0
}

func (x *ArticleData) GetArticleId() int64 {
	if x != nil {
		return x.ArticleId
	}
	return 0
}

var File_get_article_list_proto protoreflect.FileDescriptor

var file_get_article_list_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x09, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61,
	0x67, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x70, 0x61, 0x67, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x22, 0x64, 0x0a, 0x0e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x75, 0x0a, 0x0f, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x22, 0x8d, 0x01, 0x0a, 0x0b, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49,
	0x64, 0x32, 0x47, 0x0a, 0x0b, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x38, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x0a, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x49, 0x6e, 0x1a, 0x0f, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x75, 0x74, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x08, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_get_article_list_proto_rawDescOnce sync.Once
	file_get_article_list_proto_rawDescData = file_get_article_list_proto_rawDesc
)

func file_get_article_list_proto_rawDescGZIP() []byte {
	file_get_article_list_proto_rawDescOnce.Do(func() {
		file_get_article_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_get_article_list_proto_rawDescData)
	})
	return file_get_article_list_proto_rawDescData
}

var file_get_article_list_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_get_article_list_proto_goTypes = []interface{}{
	(*ArticleIn)(nil),       // 0: ArticleIn
	(*ArticleListOut)(nil),  // 1: ArticleListOut
	(*ArticleListData)(nil), // 2: ArticleListData
	(*ArticleData)(nil),     // 3: ArticleData
}
var file_get_article_list_proto_depIdxs = []int32{
	2, // 0: ArticleListOut.data:type_name -> ArticleListData
	3, // 1: ArticleListData.article:type_name -> ArticleData
	0, // 2: ArticleList.Call:input_type -> ArticleIn
	1, // 3: ArticleList.Call:output_type -> ArticleListOut
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_get_article_list_proto_init() }
func file_get_article_list_proto_init() {
	if File_get_article_list_proto != nil {
		return
	}
	file_annotations_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_get_article_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleIn); i {
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
		file_get_article_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleListOut); i {
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
		file_get_article_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleListData); i {
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
		file_get_article_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleData); i {
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
			RawDescriptor: file_get_article_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_get_article_list_proto_goTypes,
		DependencyIndexes: file_get_article_list_proto_depIdxs,
		MessageInfos:      file_get_article_list_proto_msgTypes,
	}.Build()
	File_get_article_list_proto = out.File
	file_get_article_list_proto_rawDesc = nil
	file_get_article_list_proto_goTypes = nil
	file_get_article_list_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ArticleListClient is the client API for ArticleList service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArticleListClient interface {
	Call(ctx context.Context, in *ArticleIn, opts ...grpc.CallOption) (*ArticleListOut, error)
}

type articleListClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleListClient(cc grpc.ClientConnInterface) ArticleListClient {
	return &articleListClient{cc}
}

func (c *articleListClient) Call(ctx context.Context, in *ArticleIn, opts ...grpc.CallOption) (*ArticleListOut, error) {
	out := new(ArticleListOut)
	err := c.cc.Invoke(ctx, "/ArticleList/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleListServer is the server API for ArticleList service.
type ArticleListServer interface {
	Call(context.Context, *ArticleIn) (*ArticleListOut, error)
}

// UnimplementedArticleListServer can be embedded to have forward compatible implementations.
type UnimplementedArticleListServer struct {
}

func (*UnimplementedArticleListServer) Call(context.Context, *ArticleIn) (*ArticleListOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}

func RegisterArticleListServer(s *grpc.Server, srv ArticleListServer) {
	s.RegisterService(&_ArticleList_serviceDesc, srv)
}

func _ArticleList_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleListServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ArticleList/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleListServer).Call(ctx, req.(*ArticleIn))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArticleList_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ArticleList",
	HandlerType: (*ArticleListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _ArticleList_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "get_article_list.proto",
}
