// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/download.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Download struct {
	Id         int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     int64      `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DocumentId int64      `protobuf:"varint,3,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	Ip         string     `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	IsPay      bool       `protobuf:"varint,5,opt,name=is_pay,json=isPay,proto3" json:"is_pay,omitempty"`
	CreatedAt  *time.Time `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at,omitempty"`
	UpdatedAt  *time.Time `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3,stdtime" json:"updated_at,omitempty"`
	Title      string     `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`
	Ext        string     `protobuf:"bytes,9,opt,name=ext,proto3" json:"ext,omitempty"`
	Score      int32      `protobuf:"varint,10,opt,name=score,proto3" json:"score,omitempty"`
	Size_      int64      `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
	Pages      int32      `protobuf:"varint,12,opt,name=pages,proto3" json:"pages,omitempty"`
}

func (m *Download) Reset()         { *m = Download{} }
func (m *Download) String() string { return proto.CompactTextString(m) }
func (*Download) ProtoMessage()    {}
func (*Download) Descriptor() ([]byte, []int) {
	return fileDescriptor_abfc1e2562b8c071, []int{0}
}
func (m *Download) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Download) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Download.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Download) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Download.Merge(m, src)
}
func (m *Download) XXX_Size() int {
	return m.Size()
}
func (m *Download) XXX_DiscardUnknown() {
	xxx_messageInfo_Download.DiscardUnknown(m)
}

var xxx_messageInfo_Download proto.InternalMessageInfo

func (m *Download) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Download) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Download) GetDocumentId() int64 {
	if m != nil {
		return m.DocumentId
	}
	return 0
}

func (m *Download) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Download) GetIsPay() bool {
	if m != nil {
		return m.IsPay
	}
	return false
}

func (m *Download) GetCreatedAt() *time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Download) GetUpdatedAt() *time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Download) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Download) GetExt() string {
	if m != nil {
		return m.Ext
	}
	return ""
}

func (m *Download) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Download) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *Download) GetPages() int32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

type ListDownloadRequest struct {
	Page  int64    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size_ int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Wd    string   `protobuf:"bytes,3,opt,name=wd,proto3" json:"wd,omitempty"`
	Field []string `protobuf:"bytes,4,rep,name=field,proto3" json:"field,omitempty"`
	Order string   `protobuf:"bytes,5,opt,name=order,proto3" json:"order,omitempty"`
}

func (m *ListDownloadRequest) Reset()         { *m = ListDownloadRequest{} }
func (m *ListDownloadRequest) String() string { return proto.CompactTextString(m) }
func (*ListDownloadRequest) ProtoMessage()    {}
func (*ListDownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abfc1e2562b8c071, []int{1}
}
func (m *ListDownloadRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListDownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListDownloadRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListDownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDownloadRequest.Merge(m, src)
}
func (m *ListDownloadRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListDownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDownloadRequest proto.InternalMessageInfo

func (m *ListDownloadRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListDownloadRequest) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *ListDownloadRequest) GetWd() string {
	if m != nil {
		return m.Wd
	}
	return ""
}

func (m *ListDownloadRequest) GetField() []string {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *ListDownloadRequest) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

type ListDownloadReply struct {
	Total    int64       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Download []*Download `protobuf:"bytes,2,rep,name=download,proto3" json:"download,omitempty"`
}

func (m *ListDownloadReply) Reset()         { *m = ListDownloadReply{} }
func (m *ListDownloadReply) String() string { return proto.CompactTextString(m) }
func (*ListDownloadReply) ProtoMessage()    {}
func (*ListDownloadReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_abfc1e2562b8c071, []int{2}
}
func (m *ListDownloadReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListDownloadReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListDownloadReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListDownloadReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDownloadReply.Merge(m, src)
}
func (m *ListDownloadReply) XXX_Size() int {
	return m.Size()
}
func (m *ListDownloadReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDownloadReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListDownloadReply proto.InternalMessageInfo

func (m *ListDownloadReply) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListDownloadReply) GetDownload() []*Download {
	if m != nil {
		return m.Download
	}
	return nil
}

func init() {
	proto.RegisterType((*Download)(nil), "api.v1.Download")
	proto.RegisterType((*ListDownloadRequest)(nil), "api.v1.ListDownloadRequest")
	proto.RegisterType((*ListDownloadReply)(nil), "api.v1.ListDownloadReply")
}

func init() { proto.RegisterFile("api/v1/download.proto", fileDescriptor_abfc1e2562b8c071) }

var fileDescriptor_abfc1e2562b8c071 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x5d, 0xa7, 0x6d, 0xb6, 0x99, 0xc0, 0x6a, 0xd7, 0xec, 0x0a, 0xab, 0x87, 0x34, 0xea, 0x29,
	0x07, 0x94, 0xa8, 0x45, 0x9c, 0x38, 0xa0, 0xae, 0xb8, 0x54, 0xe2, 0x50, 0x45, 0x48, 0x48, 0x5c,
	0x2a, 0x6f, 0xed, 0x8d, 0x2c, 0xa5, 0x6b, 0x93, 0x38, 0x2d, 0xe5, 0x2b, 0xf6, 0xb3, 0x38, 0xee,
	0x91, 0x1b, 0xa8, 0x15, 0xff, 0x81, 0x6c, 0x27, 0x15, 0xdc, 0xb8, 0xcd, 0x7b, 0x33, 0xf3, 0xc6,
	0xf3, 0x3c, 0x70, 0x43, 0x95, 0xc8, 0xb6, 0xd3, 0x8c, 0xc9, 0xdd, 0x43, 0x29, 0x29, 0x4b, 0x55,
	0x25, 0xb5, 0xc4, 0x3e, 0x55, 0x22, 0xdd, 0x4e, 0x47, 0xe3, 0x42, 0xca, 0xa2, 0xe4, 0x99, 0x65,
	0xef, 0x9a, 0xfb, 0x4c, 0x8b, 0x0d, 0xaf, 0x35, 0xdd, 0x28, 0x57, 0x38, 0xba, 0x2e, 0x64, 0x21,
	0x6d, 0x98, 0x99, 0xc8, 0xb1, 0x93, 0xdf, 0x1e, 0x0c, 0xdf, 0xb7, 0x8a, 0xf8, 0x02, 0x3c, 0xc1,
	0x08, 0x8a, 0x51, 0xd2, 0xcb, 0x3d, 0xc1, 0xf0, 0x4b, 0x38, 0x6f, 0x6a, 0x5e, 0xad, 0x04, 0x23,
	0x9e, 0x25, 0x7d, 0x03, 0x17, 0x0c, 0x8f, 0x21, 0x64, 0x72, 0xdd, 0x6c, 0xf8, 0x83, 0x36, 0xc9,
	0x9e, 0x4d, 0x42, 0x47, 0x2d, 0x9c, 0x92, 0x22, 0xfd, 0x18, 0x25, 0x41, 0xee, 0x09, 0x85, 0x6f,
	0xc0, 0x17, 0xf5, 0x4a, 0xd1, 0x3d, 0x19, 0xc4, 0x28, 0x19, 0xe6, 0x03, 0x51, 0x2f, 0xe9, 0x1e,
	0xbf, 0x03, 0x58, 0x57, 0x9c, 0x6a, 0xce, 0x56, 0x54, 0x13, 0x3f, 0x46, 0x49, 0x38, 0x1b, 0xa5,
	0x6e, 0x93, 0xb4, 0xdb, 0x24, 0xfd, 0xd8, 0x6d, 0x72, 0xdb, 0x7f, 0xfc, 0x39, 0x46, 0x79, 0xd0,
	0xf6, 0xcc, 0xb5, 0x11, 0x68, 0x14, 0xeb, 0x04, 0xce, 0xff, 0x57, 0xa0, 0xed, 0x99, 0x6b, 0x7c,
	0x0d, 0x03, 0x2d, 0x74, 0xc9, 0xc9, 0xd0, 0xbe, 0xd5, 0x01, 0x7c, 0x09, 0x3d, 0xfe, 0x55, 0x93,
	0xc0, 0x72, 0x26, 0x34, 0x75, 0xf5, 0x5a, 0x56, 0x9c, 0x40, 0x8c, 0x92, 0x41, 0xee, 0x00, 0xc6,
	0xd0, 0xaf, 0xc5, 0x37, 0x4e, 0x42, 0x6b, 0x80, 0x8d, 0x4d, 0xa5, 0xa2, 0x05, 0xaf, 0xc9, 0x33,
	0x57, 0x69, 0xc1, 0x64, 0x0f, 0x2f, 0x3e, 0x88, 0x5a, 0x77, 0x56, 0xe7, 0xfc, 0x4b, 0xc3, 0x6b,
	0x6d, 0x04, 0x4c, 0xbe, 0xf5, 0xdc, 0xc6, 0x27, 0x51, 0xef, 0x2f, 0xd1, 0x0b, 0xf0, 0x76, 0xce,
	0xe7, 0x20, 0xf7, 0x76, 0xcc, 0x0c, 0xb9, 0x17, 0xbc, 0x64, 0xa4, 0x1f, 0xf7, 0xcc, 0xb3, 0x2d,
	0x30, 0xac, 0xac, 0x18, 0xaf, 0xac, 0xc9, 0x41, 0xee, 0xc0, 0xe4, 0x13, 0x5c, 0xfd, 0x3b, 0x5a,
	0x95, 0x7b, 0xbb, 0xb7, 0xd4, 0xb4, 0x6c, 0x27, 0x3b, 0x80, 0x5f, 0xc1, 0xb0, 0x3b, 0x2f, 0xe2,
	0xc5, 0xbd, 0x24, 0x9c, 0x5d, 0xa6, 0xee, 0xbe, 0xd2, 0x53, 0xfb, 0xa9, 0x62, 0xf6, 0x1c, 0xc2,
	0x8e, 0x9d, 0x2f, 0x17, 0xb7, 0x6f, 0xbe, 0x1f, 0x22, 0xf4, 0x74, 0x88, 0xd0, 0xaf, 0x43, 0x84,
	0x1e, 0x8f, 0xd1, 0xd9, 0xd3, 0x31, 0x3a, 0xfb, 0x71, 0x8c, 0xce, 0xa0, 0xbd, 0xd1, 0x25, 0xfa,
	0x7c, 0xb5, 0x91, 0x15, 0x67, 0x72, 0x9d, 0xb9, 0x63, 0x7e, 0xbb, 0x9d, 0xde, 0xf9, 0xf6, 0x9b,
	0x5e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x06, 0xd6, 0x15, 0x9f, 0xe0, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DownloadAPIClient is the client API for DownloadAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DownloadAPIClient interface {
}

type downloadAPIClient struct {
	cc *grpc.ClientConn
}

func NewDownloadAPIClient(cc *grpc.ClientConn) DownloadAPIClient {
	return &downloadAPIClient{cc}
}

// DownloadAPIServer is the server API for DownloadAPI service.
type DownloadAPIServer interface {
}

// UnimplementedDownloadAPIServer can be embedded to have forward compatible implementations.
type UnimplementedDownloadAPIServer struct {
}

func RegisterDownloadAPIServer(s *grpc.Server, srv DownloadAPIServer) {
	s.RegisterService(&_DownloadAPI_serviceDesc, srv)
}

var _DownloadAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.DownloadAPI",
	HandlerType: (*DownloadAPIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "api/v1/download.proto",
}

func (m *Download) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Download) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Download) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pages != 0 {
		i = encodeVarintDownload(dAtA, i, uint64(m.Pages))
		i--
		dAtA[i] = 0x60
	}
	if m.Size_ != 0 {
		i = encodeVarintDownload(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x58
	}
	if m.Score != 0 {
		i = encodeVarintDownload(dAtA, i, uint64(m.Score))
		i--
		dAtA[i] = 0x50
	}
	if len(m.Ext) > 0 {
		i -= len(m.Ext)
		copy(dAtA[i:], m.Ext)
		i = encodeVarintDownload(dAtA, i, uint64(len(m.Ext)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintDownload(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x42
	}
	if m.UpdatedAt != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.UpdatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.UpdatedAt):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintDownload(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x3a
	}
	if m.CreatedAt != nil {
		n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedAt):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintDownload(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x32
	}
	if m.IsPay {
		i--
		if m.IsPay {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Ip) > 0 {
		i -= len(m.Ip)
		copy(dAtA[i:], m.Ip)
		i = encodeVarintDownload(dAtA, i, uint64(len(m.Ip)))
		i--
		dAtA[i] = 0x22
	}
	if m.DocumentId != 0 {
		i = encodeVarintDownload(dAtA, i, uint64(m.DocumentId))
		i--
		dAtA[i] = 0x18
	}
	if m.UserId != 0 {
		i = encodeVarintDownload(dAtA, i, uint64(m.UserId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintDownload(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ListDownloadRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListDownloadRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListDownloadRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Order) > 0 {
		i -= len(m.Order)
		copy(dAtA[i:], m.Order)
		i = encodeVarintDownload(dAtA, i, uint64(len(m.Order)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Field) > 0 {
		for iNdEx := len(m.Field) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Field[iNdEx])
			copy(dAtA[i:], m.Field[iNdEx])
			i = encodeVarintDownload(dAtA, i, uint64(len(m.Field[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Wd) > 0 {
		i -= len(m.Wd)
		copy(dAtA[i:], m.Wd)
		i = encodeVarintDownload(dAtA, i, uint64(len(m.Wd)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Size_ != 0 {
		i = encodeVarintDownload(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x10
	}
	if m.Page != 0 {
		i = encodeVarintDownload(dAtA, i, uint64(m.Page))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ListDownloadReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListDownloadReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListDownloadReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Download) > 0 {
		for iNdEx := len(m.Download) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Download[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDownload(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Total != 0 {
		i = encodeVarintDownload(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDownload(dAtA []byte, offset int, v uint64) int {
	offset -= sovDownload(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Download) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovDownload(uint64(m.Id))
	}
	if m.UserId != 0 {
		n += 1 + sovDownload(uint64(m.UserId))
	}
	if m.DocumentId != 0 {
		n += 1 + sovDownload(uint64(m.DocumentId))
	}
	l = len(m.Ip)
	if l > 0 {
		n += 1 + l + sovDownload(uint64(l))
	}
	if m.IsPay {
		n += 2
	}
	if m.CreatedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedAt)
		n += 1 + l + sovDownload(uint64(l))
	}
	if m.UpdatedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.UpdatedAt)
		n += 1 + l + sovDownload(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovDownload(uint64(l))
	}
	l = len(m.Ext)
	if l > 0 {
		n += 1 + l + sovDownload(uint64(l))
	}
	if m.Score != 0 {
		n += 1 + sovDownload(uint64(m.Score))
	}
	if m.Size_ != 0 {
		n += 1 + sovDownload(uint64(m.Size_))
	}
	if m.Pages != 0 {
		n += 1 + sovDownload(uint64(m.Pages))
	}
	return n
}

func (m *ListDownloadRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Page != 0 {
		n += 1 + sovDownload(uint64(m.Page))
	}
	if m.Size_ != 0 {
		n += 1 + sovDownload(uint64(m.Size_))
	}
	l = len(m.Wd)
	if l > 0 {
		n += 1 + l + sovDownload(uint64(l))
	}
	if len(m.Field) > 0 {
		for _, s := range m.Field {
			l = len(s)
			n += 1 + l + sovDownload(uint64(l))
		}
	}
	l = len(m.Order)
	if l > 0 {
		n += 1 + l + sovDownload(uint64(l))
	}
	return n
}

func (m *ListDownloadReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Total != 0 {
		n += 1 + sovDownload(uint64(m.Total))
	}
	if len(m.Download) > 0 {
		for _, e := range m.Download {
			l = e.Size()
			n += 1 + l + sovDownload(uint64(l))
		}
	}
	return n
}

func sovDownload(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDownload(x uint64) (n int) {
	return sovDownload(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Download) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDownload
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Download: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Download: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DocumentId", wireType)
			}
			m.DocumentId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DocumentId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDownload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDownload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ip = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPay", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsPay = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDownload
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDownload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedAt == nil {
				m.CreatedAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDownload
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDownload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UpdatedAt == nil {
				m.UpdatedAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDownload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDownload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ext", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDownload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDownload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ext = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			m.Score = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Score |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pages", wireType)
			}
			m.Pages = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pages |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDownload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDownload
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListDownloadRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDownload
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListDownloadRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListDownloadRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Page", wireType)
			}
			m.Page = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Page |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wd", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDownload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDownload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Wd = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDownload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDownload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field = append(m.Field, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Order", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDownload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDownload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Order = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDownload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDownload
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListDownloadReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDownload
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListDownloadReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListDownloadReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Download", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDownload
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDownload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Download = append(m.Download, &Download{})
			if err := m.Download[len(m.Download)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDownload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDownload
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDownload(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDownload
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDownload
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDownload
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDownload
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDownload
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDownload        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDownload          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDownload = fmt.Errorf("proto: unexpected end of group")
)