// Code generated by protoc-gen-go.
// source: github.com/safeguardproperties/gocontentservice/pkg/contentservice/contentservice.proto
// DO NOT EDIT!

/*
Package contentservice is a generated protocol buffer package.

It is generated from these files:
	github.com/safeguardproperties/gocontentservice/pkg/contentservice/contentservice.proto

It has these top-level messages:
	ImageIdsRequest
	OrderNumberRequest
	ImageIds
	ListResponse
	ImageDetail
*/
package contentservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ImageIdsRequest sent to the server
type ImageIdsRequest struct {
	ImageIds *ImageIds `protobuf:"bytes,1,opt,name=imageIds" json:"imageIds,omitempty"`
}

func (m *ImageIdsRequest) Reset()                    { *m = ImageIdsRequest{} }
func (m *ImageIdsRequest) String() string            { return proto.CompactTextString(m) }
func (*ImageIdsRequest) ProtoMessage()               {}
func (*ImageIdsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ImageIdsRequest) GetImageIds() *ImageIds {
	if m != nil {
		return m.ImageIds
	}
	return nil
}

// OrderNumberRequest sent to the server
type OrderNumberRequest struct {
	OrderNumber int64 `protobuf:"varint,1,opt,name=orderNumber" json:"orderNumber,omitempty"`
}

func (m *OrderNumberRequest) Reset()                    { *m = OrderNumberRequest{} }
func (m *OrderNumberRequest) String() string            { return proto.CompactTextString(m) }
func (*OrderNumberRequest) ProtoMessage()               {}
func (*OrderNumberRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OrderNumberRequest) GetOrderNumber() int64 {
	if m != nil {
		return m.OrderNumber
	}
	return 0
}

type ImageIds struct {
	ImageIds []int64 `protobuf:"varint,1,rep,packed,name=imageIds" json:"imageIds,omitempty"`
}

func (m *ImageIds) Reset()                    { *m = ImageIds{} }
func (m *ImageIds) String() string            { return proto.CompactTextString(m) }
func (*ImageIds) ProtoMessage()               {}
func (*ImageIds) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ImageIds) GetImageIds() []int64 {
	if m != nil {
		return m.ImageIds
	}
	return nil
}

// ListResponse sent from the server
type ListResponse struct {
	ImageDetails []*ImageDetail `protobuf:"bytes,1,rep,name=imageDetails" json:"imageDetails,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListResponse) GetImageDetails() []*ImageDetail {
	if m != nil {
		return m.ImageDetails
	}
	return nil
}

type ImageDetail struct {
	ImageId                int64                      `protobuf:"varint,1,opt,name=imageId" json:"imageId,omitempty"`
	ImageFileName          string                     `protobuf:"bytes,2,opt,name=imageFileName" json:"imageFileName,omitempty"`
	ScanDate               *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=scanDate" json:"scanDate,omitempty"`
	ImageUTCDate           *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=imageUTCDate" json:"imageUTCDate,omitempty"`
	ImageTakenDate         *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=imageTakenDate" json:"imageTakenDate,omitempty"`
	DateCreated            *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=dateCreated" json:"dateCreated,omitempty"`
	OrderNumber            int64                      `protobuf:"varint,8,opt,name=orderNumber" json:"orderNumber,omitempty"`
	Archived               string                     `protobuf:"bytes,9,opt,name=archived" json:"archived,omitempty"`
	Category               string                     `protobuf:"bytes,10,opt,name=category" json:"category,omitempty"`
	ContractorId           int64                      `protobuf:"varint,11,opt,name=contractorId" json:"contractorId,omitempty"`
	DateModified           *google_protobuf.Timestamp `protobuf:"bytes,12,opt,name=dateModified" json:"dateModified,omitempty"`
	DeptCode               string                     `protobuf:"bytes,13,opt,name=deptCode" json:"deptCode,omitempty"`
	DescPrefix             string                     `protobuf:"bytes,14,opt,name=descPrefix" json:"descPrefix,omitempty"`
	DescText               string                     `protobuf:"bytes,15,opt,name=descText" json:"descText,omitempty"`
	FileSize               int32                      `protobuf:"varint,16,opt,name=fileSize" json:"fileSize,omitempty"`
	ImageHeight            int32                      `protobuf:"varint,17,opt,name=imageHeight" json:"imageHeight,omitempty"`
	ImageWidth             int32                      `protobuf:"varint,18,opt,name=imageWidth" json:"imageWidth,omitempty"`
	ImageType              int32                      `protobuf:"varint,19,opt,name=imageType" json:"imageType,omitempty"`
	ImageRotated           bool                       `protobuf:"varint,20,opt,name=imageRotated" json:"imageRotated,omitempty"`
	ReleaseDate            *google_protobuf.Timestamp `protobuf:"bytes,21,opt,name=releaseDate" json:"releaseDate,omitempty"`
	ThumbnailSize          int32                      `protobuf:"varint,22,opt,name=thumbnailSize" json:"thumbnailSize,omitempty"`
	MimeType               string                     `protobuf:"bytes,23,opt,name=mimeType" json:"mimeType,omitempty"`
	WebFileName            string                     `protobuf:"bytes,24,opt,name=webFileName" json:"webFileName,omitempty"`
	GeneratedImageFilePath string                     `protobuf:"bytes,25,opt,name=generatedImageFilePath" json:"generatedImageFilePath,omitempty"`
	Guid                   string                     `protobuf:"bytes,26,opt,name=guid" json:"guid,omitempty"`
}

func (m *ImageDetail) Reset()                    { *m = ImageDetail{} }
func (m *ImageDetail) String() string            { return proto.CompactTextString(m) }
func (*ImageDetail) ProtoMessage()               {}
func (*ImageDetail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ImageDetail) GetImageId() int64 {
	if m != nil {
		return m.ImageId
	}
	return 0
}

func (m *ImageDetail) GetImageFileName() string {
	if m != nil {
		return m.ImageFileName
	}
	return ""
}

func (m *ImageDetail) GetScanDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.ScanDate
	}
	return nil
}

func (m *ImageDetail) GetImageUTCDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.ImageUTCDate
	}
	return nil
}

func (m *ImageDetail) GetImageTakenDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.ImageTakenDate
	}
	return nil
}

func (m *ImageDetail) GetDateCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.DateCreated
	}
	return nil
}

func (m *ImageDetail) GetOrderNumber() int64 {
	if m != nil {
		return m.OrderNumber
	}
	return 0
}

func (m *ImageDetail) GetArchived() string {
	if m != nil {
		return m.Archived
	}
	return ""
}

func (m *ImageDetail) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ImageDetail) GetContractorId() int64 {
	if m != nil {
		return m.ContractorId
	}
	return 0
}

func (m *ImageDetail) GetDateModified() *google_protobuf.Timestamp {
	if m != nil {
		return m.DateModified
	}
	return nil
}

func (m *ImageDetail) GetDeptCode() string {
	if m != nil {
		return m.DeptCode
	}
	return ""
}

func (m *ImageDetail) GetDescPrefix() string {
	if m != nil {
		return m.DescPrefix
	}
	return ""
}

func (m *ImageDetail) GetDescText() string {
	if m != nil {
		return m.DescText
	}
	return ""
}

func (m *ImageDetail) GetFileSize() int32 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *ImageDetail) GetImageHeight() int32 {
	if m != nil {
		return m.ImageHeight
	}
	return 0
}

func (m *ImageDetail) GetImageWidth() int32 {
	if m != nil {
		return m.ImageWidth
	}
	return 0
}

func (m *ImageDetail) GetImageType() int32 {
	if m != nil {
		return m.ImageType
	}
	return 0
}

func (m *ImageDetail) GetImageRotated() bool {
	if m != nil {
		return m.ImageRotated
	}
	return false
}

func (m *ImageDetail) GetReleaseDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.ReleaseDate
	}
	return nil
}

func (m *ImageDetail) GetThumbnailSize() int32 {
	if m != nil {
		return m.ThumbnailSize
	}
	return 0
}

func (m *ImageDetail) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *ImageDetail) GetWebFileName() string {
	if m != nil {
		return m.WebFileName
	}
	return ""
}

func (m *ImageDetail) GetGeneratedImageFilePath() string {
	if m != nil {
		return m.GeneratedImageFilePath
	}
	return ""
}

func (m *ImageDetail) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func init() {
	proto.RegisterType((*ImageIdsRequest)(nil), "contentservice.ImageIdsRequest")
	proto.RegisterType((*OrderNumberRequest)(nil), "contentservice.OrderNumberRequest")
	proto.RegisterType((*ImageIds)(nil), "contentservice.ImageIds")
	proto.RegisterType((*ListResponse)(nil), "contentservice.ListResponse")
	proto.RegisterType((*ImageDetail)(nil), "contentservice.ImageDetail")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ContentService service

type ContentServiceClient interface {
	// Returns ImageDetail records for the requested imageids
	ListByImageIds(ctx context.Context, in *ImageIdsRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Returns Imagedetail records for the requested ordernumber
	ListByOrderNumber(ctx context.Context, in *OrderNumberRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type contentServiceClient struct {
	cc *grpc.ClientConn
}

func NewContentServiceClient(cc *grpc.ClientConn) ContentServiceClient {
	return &contentServiceClient{cc}
}

func (c *contentServiceClient) ListByImageIds(ctx context.Context, in *ImageIdsRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/contentservice.ContentService/ListByImageIds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) ListByOrderNumber(ctx context.Context, in *OrderNumberRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/contentservice.ContentService/ListByOrderNumber", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ContentService service

type ContentServiceServer interface {
	// Returns ImageDetail records for the requested imageids
	ListByImageIds(context.Context, *ImageIdsRequest) (*ListResponse, error)
	// Returns Imagedetail records for the requested ordernumber
	ListByOrderNumber(context.Context, *OrderNumberRequest) (*ListResponse, error)
}

func RegisterContentServiceServer(s *grpc.Server, srv ContentServiceServer) {
	s.RegisterService(&_ContentService_serviceDesc, srv)
}

func _ContentService_ListByImageIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).ListByImageIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentservice.ContentService/ListByImageIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).ListByImageIds(ctx, req.(*ImageIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_ListByOrderNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).ListByOrderNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentservice.ContentService/ListByOrderNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).ListByOrderNumber(ctx, req.(*OrderNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "contentservice.ContentService",
	HandlerType: (*ContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListByImageIds",
			Handler:    _ContentService_ListByImageIds_Handler,
		},
		{
			MethodName: "ListByOrderNumber",
			Handler:    _ContentService_ListByOrderNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/safeguardproperties/gocontentservice/pkg/contentservice/contentservice.proto",
}

func init() {
	proto.RegisterFile("github.com/safeguardproperties/gocontentservice/pkg/contentservice/contentservice.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 671 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0xfd, 0xf5, 0xb7, 0x7f, 0x9d, 0xdb, 0x75, 0xcc, 0xc0, 0x30, 0x65, 0x82, 0xaa, 0x42, 0xa8,
	0x4f, 0x8d, 0x34, 0xd0, 0x5e, 0x40, 0x20, 0xad, 0x13, 0x50, 0x09, 0xb6, 0x91, 0x15, 0xed, 0xd9,
	0x89, 0x6f, 0x53, 0x6b, 0x49, 0x1c, 0x6c, 0x67, 0xac, 0x7c, 0x36, 0xbe, 0x11, 0x5f, 0x02, 0xd9,
	0x6e, 0xb2, 0xa4, 0x63, 0xea, 0x5b, 0xce, 0xb9, 0xe7, 0xfe, 0xe9, 0x3d, 0xf5, 0x45, 0x97, 0x11,
	0xd7, 0xb3, 0x3c, 0x18, 0x86, 0x22, 0xf1, 0x14, 0x9d, 0x42, 0x94, 0x53, 0xc9, 0x32, 0x29, 0x32,
	0x90, 0x9a, 0x83, 0xf2, 0x22, 0x11, 0x8a, 0x54, 0x43, 0xaa, 0x15, 0xc8, 0x6b, 0x1e, 0x82, 0x97,
	0x5d, 0x45, 0xde, 0x12, 0x55, 0x87, 0xc3, 0x4c, 0x0a, 0x2d, 0x70, 0xa7, 0xce, 0x76, 0xdf, 0x56,
	0x1a, 0x45, 0x22, 0xa6, 0x69, 0xe4, 0x59, 0x61, 0x90, 0x4f, 0xbd, 0x4c, 0xcf, 0x33, 0x50, 0x9e,
	0xe6, 0x09, 0x28, 0x4d, 0x93, 0xec, 0xf6, 0xcb, 0x15, 0xeb, 0x7f, 0x42, 0xbb, 0xe3, 0x84, 0x46,
	0x30, 0x66, 0xca, 0x87, 0x1f, 0x39, 0x28, 0x8d, 0xdf, 0xa0, 0x26, 0x5f, 0x50, 0xa4, 0xd1, 0x6b,
	0x0c, 0x5a, 0x87, 0x64, 0xb8, 0x34, 0x48, 0x99, 0x52, 0x2a, 0xfb, 0x47, 0x08, 0x9f, 0x49, 0x06,
	0xf2, 0x34, 0x4f, 0x02, 0x90, 0x45, 0xad, 0x1e, 0x6a, 0x89, 0x5b, 0xd6, 0x96, 0x5b, 0xf3, 0xab,
	0x54, 0xff, 0x15, 0x6a, 0x16, 0xd5, 0x70, 0xb7, 0xd6, 0x79, 0x6d, 0xb0, 0x56, 0xa9, 0x7f, 0x86,
	0xda, 0x5f, 0xb8, 0xd2, 0x3e, 0xa8, 0x4c, 0xa4, 0x0a, 0xf0, 0x07, 0xd4, 0xb6, 0xb1, 0x13, 0xd0,
	0x94, 0xc7, 0x4e, 0xdf, 0x3a, 0x7c, 0xf6, 0xcf, 0x49, 0x9d, 0xc6, 0xaf, 0x25, 0xf4, 0xff, 0x6c,
	0xa1, 0x56, 0x25, 0x8a, 0x09, 0xda, 0x5a, 0x34, 0x5b, 0x8c, 0x59, 0x40, 0xfc, 0x12, 0xed, 0xd8,
	0xcf, 0x8f, 0x3c, 0x86, 0x53, 0x9a, 0x00, 0xf9, 0xbf, 0xd7, 0x18, 0x6c, 0xfb, 0x75, 0x12, 0x1f,
	0xa1, 0xa6, 0x0a, 0x69, 0x7a, 0x42, 0x35, 0x90, 0x75, 0xbb, 0xb6, 0xee, 0x30, 0x12, 0x22, 0x8a,
	0x17, 0xbe, 0x05, 0xf9, 0x74, 0x38, 0x29, 0xb6, 0xef, 0x97, 0x5a, 0xfc, 0x7e, 0xf1, 0x43, 0xbe,
	0x4f, 0x46, 0x36, 0x77, 0x63, 0x65, 0x6e, 0x4d, 0x8f, 0x8f, 0x51, 0xc7, 0xe2, 0x09, 0xbd, 0x02,
	0xd7, 0x7d, 0x73, 0x65, 0x85, 0xa5, 0x0c, 0xfc, 0x0e, 0xb5, 0x18, 0xd5, 0x30, 0x92, 0x40, 0x35,
	0x30, 0xb2, 0xb5, 0xb2, 0x40, 0x55, 0xbe, 0x6c, 0x72, 0xf3, 0x8e, 0xc9, 0xc6, 0x58, 0x2a, 0xc3,
	0x19, 0xbf, 0x06, 0x46, 0xb6, 0xed, 0xf2, 0x4a, 0x6c, 0x62, 0x21, 0xd5, 0x10, 0x09, 0x39, 0x27,
	0xc8, 0xc5, 0x0a, 0x8c, 0xfb, 0xa8, 0x6d, 0xfc, 0x94, 0x34, 0xd4, 0x42, 0x8e, 0x19, 0x69, 0xd9,
	0xd2, 0x35, 0xce, 0xec, 0xcf, 0x0c, 0xf3, 0x55, 0x30, 0x3e, 0xe5, 0xc0, 0x48, 0x7b, 0xf5, 0xfe,
	0xaa, 0x7a, 0xd3, 0x9f, 0x41, 0xa6, 0x47, 0x82, 0x01, 0xd9, 0x71, 0xfd, 0x0b, 0x8c, 0x9f, 0x23,
	0xc4, 0x40, 0x85, 0xe7, 0x12, 0xa6, 0xfc, 0x86, 0x74, 0x6c, 0xb4, 0xc2, 0xb8, 0x5c, 0x15, 0x4e,
	0xe0, 0x46, 0x93, 0xdd, 0x22, 0xd7, 0x61, 0x13, 0x9b, 0xf2, 0x18, 0x2e, 0xf8, 0x2f, 0x20, 0x0f,
	0x7a, 0x8d, 0xc1, 0x86, 0x5f, 0x62, 0xb3, 0x31, 0xeb, 0xc0, 0x67, 0xe0, 0xd1, 0x4c, 0x93, 0x3d,
	0x1b, 0xae, 0x52, 0xa6, 0xb3, 0x85, 0x97, 0x9c, 0xe9, 0x19, 0xc1, 0x56, 0x50, 0x61, 0xf0, 0x01,
	0xda, 0x76, 0x1e, 0xce, 0x33, 0x20, 0x0f, 0x6d, 0xf8, 0x96, 0x30, 0x7b, 0xb3, 0xc0, 0x17, 0xda,
	0x1a, 0xfa, 0xa8, 0xd7, 0x18, 0x34, 0xfd, 0x1a, 0x67, 0x3c, 0x97, 0x10, 0x03, 0x55, 0x60, 0xff,
	0x34, 0x8f, 0x57, 0x7b, 0x5e, 0x91, 0x9b, 0x37, 0xa1, 0x67, 0x79, 0x12, 0xa4, 0x94, 0xc7, 0xf6,
	0x27, 0xee, 0xdb, 0x19, 0xea, 0xa4, 0xd9, 0x41, 0xc2, 0x13, 0x37, 0xe4, 0x13, 0xb7, 0x9f, 0x02,
	0x9b, 0x1d, 0xfc, 0x84, 0xa0, 0x7c, 0x53, 0xc4, 0x86, 0xab, 0x14, 0x3e, 0x42, 0xfb, 0x11, 0xa4,
	0x20, 0xcd, 0xb8, 0xe3, 0xe2, 0xad, 0x9d, 0x53, 0x3d, 0x23, 0x4f, 0xad, 0xf8, 0x9e, 0x28, 0xc6,
	0x68, 0x3d, 0xca, 0x39, 0x23, 0x5d, 0xab, 0xb2, 0xdf, 0x87, 0xbf, 0x1b, 0xa8, 0x33, 0x72, 0xa7,
	0xe1, 0xc2, 0x9d, 0x06, 0xfc, 0x0d, 0x75, 0xcc, 0x45, 0x39, 0x9e, 0x97, 0xf7, 0xe7, 0xc5, 0xbd,
	0x77, 0xce, 0x9d, 0xb3, 0xee, 0xc1, 0xb2, 0xa0, 0x7a, 0x92, 0xfa, 0xff, 0xe1, 0x4b, 0xb4, 0xe7,
	0x4a, 0x56, 0x4e, 0x21, 0xee, 0x2f, 0x27, 0xdd, 0xbd, 0x93, 0xab, 0x0a, 0x07, 0x9b, 0xd6, 0x8f,
	0xd7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x83, 0x12, 0x91, 0x63, 0x55, 0x06, 0x00, 0x00,
}