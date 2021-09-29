// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/vault/v1beta1/querier.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type VaultInfo struct {
	Id                     uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PairID                 uint64                                 `protobuf:"varint,2,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty" yaml:"pair_id"`
	Owner                  string                                 `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	Collateral             types.Coin                             `protobuf:"bytes,4,opt,name=collateral,proto3" json:"collateral" yaml:"collateral"`
	Debt                   types.Coin                             `protobuf:"bytes,5,opt,name=debt,proto3" json:"debt" yaml:"debt"`
	CollateralizationRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=collateralization_ratio,json=collateralizationRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"collateralization_ratio" yaml:"collateralization_ratio"`
}

func (m *VaultInfo) Reset()         { *m = VaultInfo{} }
func (m *VaultInfo) String() string { return proto.CompactTextString(m) }
func (*VaultInfo) ProtoMessage()    {}
func (*VaultInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3453303216b6f983, []int{0}
}
func (m *VaultInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VaultInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VaultInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultInfo.Merge(m, src)
}
func (m *VaultInfo) XXX_Size() int {
	return m.Size()
}
func (m *VaultInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VaultInfo proto.InternalMessageInfo

type QueryVaultRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
}

func (m *QueryVaultRequest) Reset()         { *m = QueryVaultRequest{} }
func (m *QueryVaultRequest) String() string { return proto.CompactTextString(m) }
func (*QueryVaultRequest) ProtoMessage()    {}
func (*QueryVaultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3453303216b6f983, []int{1}
}
func (m *QueryVaultRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVaultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVaultRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVaultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVaultRequest.Merge(m, src)
}
func (m *QueryVaultRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryVaultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVaultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVaultRequest proto.InternalMessageInfo

type QueryVaultResponse struct {
	VaultInfo VaultInfo `protobuf:"bytes,1,opt,name=vaultInfo,proto3" json:"vaultInfo" yaml:"vaultInfo"`
}

func (m *QueryVaultResponse) Reset()         { *m = QueryVaultResponse{} }
func (m *QueryVaultResponse) String() string { return proto.CompactTextString(m) }
func (*QueryVaultResponse) ProtoMessage()    {}
func (*QueryVaultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3453303216b6f983, []int{2}
}
func (m *QueryVaultResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVaultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVaultResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVaultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVaultResponse.Merge(m, src)
}
func (m *QueryVaultResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryVaultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVaultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVaultResponse proto.InternalMessageInfo

type QueryVaultsRequest struct {
	Owner      string             `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty" yaml:"pagination"`
}

func (m *QueryVaultsRequest) Reset()         { *m = QueryVaultsRequest{} }
func (m *QueryVaultsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryVaultsRequest) ProtoMessage()    {}
func (*QueryVaultsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3453303216b6f983, []int{3}
}
func (m *QueryVaultsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVaultsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVaultsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVaultsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVaultsRequest.Merge(m, src)
}
func (m *QueryVaultsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryVaultsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVaultsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVaultsRequest proto.InternalMessageInfo

type QueryVaultsResponse struct {
	VaultsInfo []VaultInfo         `protobuf:"bytes,1,rep,name=vaultsInfo,proto3" json:"vaultsInfo" yaml:"vaultsInfo"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty" yaml:"pagination"`
}

func (m *QueryVaultsResponse) Reset()         { *m = QueryVaultsResponse{} }
func (m *QueryVaultsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryVaultsResponse) ProtoMessage()    {}
func (*QueryVaultsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3453303216b6f983, []int{4}
}
func (m *QueryVaultsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVaultsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVaultsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVaultsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVaultsResponse.Merge(m, src)
}
func (m *QueryVaultsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryVaultsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVaultsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVaultsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VaultInfo)(nil), "comdex.vault.v1beta1.VaultInfo")
	proto.RegisterType((*QueryVaultRequest)(nil), "comdex.vault.v1beta1.QueryVaultRequest")
	proto.RegisterType((*QueryVaultResponse)(nil), "comdex.vault.v1beta1.QueryVaultResponse")
	proto.RegisterType((*QueryVaultsRequest)(nil), "comdex.vault.v1beta1.QueryVaultsRequest")
	proto.RegisterType((*QueryVaultsResponse)(nil), "comdex.vault.v1beta1.QueryVaultsResponse")
}

func init() {
	proto.RegisterFile("comdex/vault/v1beta1/querier.proto", fileDescriptor_3453303216b6f983)
}

var fileDescriptor_3453303216b6f983 = []byte{
	// 700 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x31, 0x6f, 0xd3, 0x40,
	0x18, 0xcd, 0xa5, 0x6d, 0x50, 0x2e, 0xa5, 0x6a, 0xaf, 0x05, 0xd2, 0xa8, 0xd8, 0xe5, 0x80, 0x34,
	0x45, 0xd4, 0x56, 0x83, 0x58, 0x18, 0x4d, 0x97, 0x6e, 0xad, 0x41, 0x1d, 0x90, 0xa0, 0xba, 0xd8,
	0xd7, 0xf4, 0x84, 0xe3, 0x4b, 0x6d, 0x27, 0xd0, 0xa2, 0x2e, 0x6c, 0x30, 0x55, 0x62, 0x64, 0x61,
	0x64, 0xe2, 0x77, 0x74, 0xac, 0xc4, 0x82, 0x18, 0x0c, 0xa4, 0xfc, 0x82, 0xfc, 0x02, 0xe4, 0x3b,
	0x3b, 0x76, 0x68, 0xd4, 0x74, 0xf2, 0xe9, 0xee, 0x7b, 0xef, 0x7b, 0xef, 0xbb, 0x77, 0x86, 0xd8,
	0xe2, 0x2d, 0x9b, 0xbe, 0xd5, 0xbb, 0xa4, 0xe3, 0x04, 0x7a, 0x77, 0xbd, 0x41, 0x03, 0xb2, 0xae,
	0x1f, 0x74, 0xa8, 0xc7, 0xa8, 0xa7, 0xb5, 0x3d, 0x1e, 0x70, 0xb4, 0x20, 0x6b, 0x34, 0x51, 0xa3,
	0xc5, 0x35, 0x95, 0x07, 0x16, 0xf7, 0x5b, 0xdc, 0xd7, 0x1b, 0xc4, 0xa7, 0x02, 0x70, 0x38, 0x80,
	0xb7, 0x49, 0x93, 0xb9, 0x24, 0x60, 0xdc, 0x95, 0x0c, 0x95, 0x85, 0x26, 0x6f, 0x72, 0xb1, 0xd4,
	0xa3, 0x55, 0xbc, 0xbb, 0xd4, 0xe4, 0xbc, 0xe9, 0x50, 0x9d, 0xb4, 0x99, 0x4e, 0x5c, 0x97, 0x07,
	0x02, 0xe2, 0xc7, 0xa7, 0x4a, 0x96, 0x3f, 0x61, 0xb6, 0x38, 0x8b, 0x39, 0xf1, 0x97, 0x09, 0x58,
	0xdc, 0x89, 0x14, 0x6d, 0xba, 0x7b, 0x1c, 0xcd, 0xc0, 0x3c, 0xb3, 0xcb, 0x60, 0x19, 0xd4, 0x26,
	0xcd, 0x3c, 0xb3, 0xd1, 0x63, 0x78, 0xad, 0x4d, 0x98, 0xb7, 0xcb, 0xec, 0x72, 0x3e, 0xda, 0x34,
	0x96, 0x7a, 0xa1, 0x5a, 0xd8, 0x22, 0xcc, 0xdb, 0xdc, 0xe8, 0x87, 0xea, 0xcc, 0x21, 0x69, 0x39,
	0x4f, 0x70, 0x5c, 0x82, 0xcd, 0x42, 0xb4, 0xda, 0xb4, 0x51, 0x15, 0x4e, 0xf1, 0x37, 0x2e, 0xf5,
	0xca, 0x13, 0xcb, 0xa0, 0x56, 0x34, 0x66, 0xfb, 0xa1, 0x3a, 0x2d, 0x4b, 0xc5, 0x36, 0x36, 0xe5,
	0x31, 0x7a, 0x0e, 0xa1, 0xc5, 0x1d, 0x87, 0x04, 0xd4, 0x23, 0x4e, 0x79, 0x72, 0x19, 0xd4, 0x4a,
	0xf5, 0x45, 0x4d, 0x2a, 0xd6, 0x22, 0xc5, 0xc9, 0x98, 0xb4, 0xa7, 0x9c, 0xb9, 0xc6, 0xe2, 0x69,
	0xa8, 0xe6, 0xfa, 0xa1, 0x3a, 0x27, 0xb9, 0x52, 0x28, 0x36, 0x33, 0x3c, 0xc8, 0x80, 0x93, 0x36,
	0x6d, 0x04, 0xe5, 0xa9, 0x71, 0x7c, 0xf3, 0x31, 0x5f, 0x49, 0xf2, 0x45, 0x20, 0x6c, 0x0a, 0x2c,
	0xfa, 0x00, 0xe0, 0xad, 0x94, 0x92, 0x1d, 0x89, 0x99, 0xee, 0x7a, 0xd1, 0xa7, 0x5c, 0x10, 0xa6,
	0xb6, 0x22, 0xf0, 0xcf, 0x50, 0xad, 0x36, 0x59, 0xb0, 0xdf, 0x69, 0x68, 0x16, 0x6f, 0xe9, 0xf1,
	0xac, 0xe5, 0x67, 0xcd, 0xb7, 0x5f, 0xeb, 0xc1, 0x61, 0x9b, 0xfa, 0xda, 0x06, 0xb5, 0xfa, 0xa1,
	0xaa, 0xfc, 0x2f, 0x7b, 0x88, 0x16, 0x9b, 0x37, 0x2f, 0x9c, 0x98, 0xe2, 0xa0, 0x0e, 0xe7, 0xb6,
	0xa3, 0x60, 0x88, 0x6b, 0x32, 0xe9, 0x41, 0x87, 0xfa, 0x01, 0xba, 0x9d, 0xde, 0x94, 0x71, 0xbd,
	0x1f, 0xaa, 0x45, 0x49, 0x1e, 0xdd, 0x42, 0x9e, 0xd9, 0xf8, 0x08, 0xa2, 0x2c, 0xc6, 0x6f, 0x73,
	0xd7, 0xa7, 0xc8, 0x86, 0xc5, 0x6e, 0x72, 0xd7, 0x02, 0x5b, 0xaa, 0xab, 0xda, 0xa8, 0x58, 0x6a,
	0x83, 0x48, 0x18, 0xf7, 0x23, 0x9f, 0xbd, 0x50, 0x4d, 0x53, 0xd2, 0x0f, 0xd5, 0x59, 0xd9, 0x6d,
	0x40, 0x86, 0xcd, 0x94, 0x18, 0x7f, 0x06, 0xd9, 0xe6, 0x7e, 0xa2, 0x78, 0x10, 0x0a, 0x70, 0x79,
	0x28, 0x5e, 0x42, 0x98, 0x26, 0x5f, 0xc4, 0xae, 0x54, 0xaf, 0x0e, 0x5d, 0xa2, 0x78, 0x26, 0x03,
	0xa9, 0x5b, 0xa4, 0x49, 0xe3, 0x1e, 0xc6, 0x8d, 0x34, 0x1d, 0x29, 0x07, 0x36, 0x33, 0x84, 0xf8,
	0x17, 0x80, 0xf3, 0x43, 0xea, 0xe2, 0xd9, 0xec, 0x43, 0x28, 0x2c, 0xf8, 0xf1, 0x70, 0x26, 0xae,
	0x32, 0x9c, 0x95, 0x78, 0x38, 0x70, 0x67, 0x00, 0x4d, 0x15, 0xa4, 0x74, 0xd8, 0xcc, 0x70, 0xa3,
	0x57, 0x23, 0x0c, 0xae, 0x8c, 0x35, 0x28, 0x65, 0x5e, 0xc1, 0x61, 0xfd, 0x5b, 0x1e, 0x4e, 0x0b,
	0x87, 0xcf, 0xa8, 0xd7, 0x65, 0x16, 0x45, 0x1f, 0x01, 0x84, 0xa9, 0x65, 0xb4, 0x32, 0xda, 0xd5,
	0x85, 0x8c, 0x55, 0x6a, 0xe3, 0x0b, 0xa5, 0x2a, 0xbc, 0xfa, 0xfe, 0xfb, 0xdf, 0x4f, 0xf9, 0xbb,
	0xe8, 0x8e, 0x3e, 0xf2, 0x47, 0x28, 0xcd, 0xeb, 0xef, 0x98, 0x7d, 0x8c, 0x4e, 0x00, 0x2c, 0x65,
	0xe6, 0x8f, 0xc6, 0x36, 0x49, 0x02, 0x54, 0x59, 0xbd, 0x42, 0x65, 0xac, 0xe7, 0xa1, 0xd0, 0x53,
	0x45, 0xf7, 0x2e, 0xd7, 0x23, 0x02, 0x77, 0x6c, 0x6c, 0x9f, 0xfe, 0x51, 0x72, 0x5f, 0x7b, 0x4a,
	0xee, 0xb4, 0xa7, 0x80, 0xb3, 0x9e, 0x02, 0x7e, 0xf7, 0x14, 0x70, 0x72, 0xae, 0xe4, 0xce, 0xce,
	0x95, 0xdc, 0x8f, 0x73, 0x25, 0xf7, 0x42, 0x1f, 0x7a, 0xe4, 0x11, 0xe3, 0x1a, 0xdf, 0xdb, 0x63,
	0x16, 0x23, 0x4e, 0xd2, 0x21, 0xe9, 0x21, 0x5e, 0x7c, 0xa3, 0x20, 0xfe, 0xae, 0x8f, 0xfe, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x20, 0x07, 0x04, 0x73, 0x19, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	QueryVault(ctx context.Context, in *QueryVaultRequest, opts ...grpc.CallOption) (*QueryVaultResponse, error)
	QueryVaults(ctx context.Context, in *QueryVaultsRequest, opts ...grpc.CallOption) (*QueryVaultsResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) QueryVault(ctx context.Context, in *QueryVaultRequest, opts ...grpc.CallOption) (*QueryVaultResponse, error) {
	out := new(QueryVaultResponse)
	err := c.cc.Invoke(ctx, "/comdex.vault.v1beta1.QueryService/QueryVault", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) QueryVaults(ctx context.Context, in *QueryVaultsRequest, opts ...grpc.CallOption) (*QueryVaultsResponse, error) {
	out := new(QueryVaultsResponse)
	err := c.cc.Invoke(ctx, "/comdex.vault.v1beta1.QueryService/QueryVaults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	QueryVault(context.Context, *QueryVaultRequest) (*QueryVaultResponse, error)
	QueryVaults(context.Context, *QueryVaultsRequest) (*QueryVaultsResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) QueryVault(ctx context.Context, req *QueryVaultRequest) (*QueryVaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryVault not implemented")
}
func (*UnimplementedQueryServiceServer) QueryVaults(ctx context.Context, req *QueryVaultsRequest) (*QueryVaultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryVaults not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_QueryVault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QueryVault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comdex.vault.v1beta1.QueryService/QueryVault",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QueryVault(ctx, req.(*QueryVaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_QueryVaults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVaultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QueryVaults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comdex.vault.v1beta1.QueryService/QueryVaults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QueryVaults(ctx, req.(*QueryVaultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comdex.vault.v1beta1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryVault",
			Handler:    _QueryService_QueryVault_Handler,
		},
		{
			MethodName: "QueryVaults",
			Handler:    _QueryService_QueryVaults_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comdex/vault/v1beta1/querier.proto",
}

func (m *VaultInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CollateralizationRatio.Size()
		i -= size
		if _, err := m.CollateralizationRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuerier(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.Debt.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuerier(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Collateral.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuerier(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintQuerier(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PairID != 0 {
		i = encodeVarintQuerier(dAtA, i, uint64(m.PairID))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintQuerier(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryVaultRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVaultRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVaultRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuerier(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryVaultResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVaultResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVaultResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.VaultInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuerier(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryVaultsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVaultsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVaultsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuerier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintQuerier(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryVaultsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVaultsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVaultsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuerier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.VaultsInfo) > 0 {
		for iNdEx := len(m.VaultsInfo) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VaultsInfo[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuerier(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuerier(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuerier(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VaultInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuerier(uint64(m.Id))
	}
	if m.PairID != 0 {
		n += 1 + sovQuerier(uint64(m.PairID))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovQuerier(uint64(l))
	}
	l = m.Collateral.Size()
	n += 1 + l + sovQuerier(uint64(l))
	l = m.Debt.Size()
	n += 1 + l + sovQuerier(uint64(l))
	l = m.CollateralizationRatio.Size()
	n += 1 + l + sovQuerier(uint64(l))
	return n
}

func (m *QueryVaultRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuerier(uint64(m.Id))
	}
	return n
}

func (m *QueryVaultResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.VaultInfo.Size()
	n += 1 + l + sovQuerier(uint64(l))
	return n
}

func (m *QueryVaultsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovQuerier(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QueryVaultsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.VaultsInfo) > 0 {
		for _, e := range m.VaultsInfo {
			l = e.Size()
			n += 1 + l + sovQuerier(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func sovQuerier(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuerier(x uint64) (n int) {
	return sovQuerier(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VaultInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: VaultInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairID", wireType)
			}
			m.PairID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PairID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collateral", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Collateral.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Debt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Debt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralizationRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollateralizationRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryVaultRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryVaultRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVaultRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryVaultResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryVaultResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVaultResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VaultInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VaultInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryVaultsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryVaultsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVaultsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryVaultsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryVaultsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVaultsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VaultsInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VaultsInfo = append(m.VaultsInfo, VaultInfo{})
			if err := m.VaultsInfo[len(m.VaultsInfo)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func skipQuerier(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuerier
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
					return 0, ErrIntOverflowQuerier
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
					return 0, ErrIntOverflowQuerier
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
				return 0, ErrInvalidLengthQuerier
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuerier
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuerier
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuerier        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuerier          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuerier = fmt.Errorf("proto: unexpected end of group")
)