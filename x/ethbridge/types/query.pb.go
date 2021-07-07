// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sifnode/ethbridge/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/Sifchain/sifnode/x/oracle/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// QueryEthProphecyRequest payload for EthProphecy rpc query
type QueryEthProphecyRequest struct {
	NetworkDescriptor types.NetworkDescriptor `protobuf:"varint,1,opt,name=network_descriptor,json=networkDescriptor,proto3,enum=sifnode.oracle.v1.NetworkDescriptor" json:"network_descriptor,omitempty"`
	// bridge_contract_address is an EthereumAddress
	BridgeContractAddress string `protobuf:"bytes,2,opt,name=bridge_contract_address,json=bridgeContractAddress,proto3" json:"bridge_contract_address,omitempty" yaml:"bridge_registry_contract_address"`
	Nonce                 int64  `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Symbol                string `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// token_contract_address is an EthereumAddress
	TokenContractAddress string `protobuf:"bytes,5,opt,name=token_contract_address,json=tokenContractAddress,proto3" json:"token_contract_address,omitempty"`
	// ethereum_sender is an EthereumAddress
	EthereumSender string `protobuf:"bytes,6,opt,name=ethereum_sender,json=ethereumSender,proto3" json:"ethereum_sender,omitempty"`
}

func (m *QueryEthProphecyRequest) Reset()         { *m = QueryEthProphecyRequest{} }
func (m *QueryEthProphecyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEthProphecyRequest) ProtoMessage()    {}
func (*QueryEthProphecyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7077edcf9f792b78, []int{0}
}
func (m *QueryEthProphecyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEthProphecyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEthProphecyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEthProphecyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEthProphecyRequest.Merge(m, src)
}
func (m *QueryEthProphecyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEthProphecyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEthProphecyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEthProphecyRequest proto.InternalMessageInfo

func (m *QueryEthProphecyRequest) GetNetworkDescriptor() types.NetworkDescriptor {
	if m != nil {
		return m.NetworkDescriptor
	}
	return types.NetworkDescriptor_NETWORK_DESCRIPTOR_UNSPECIFIED
}

func (m *QueryEthProphecyRequest) GetBridgeContractAddress() string {
	if m != nil {
		return m.BridgeContractAddress
	}
	return ""
}

func (m *QueryEthProphecyRequest) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *QueryEthProphecyRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *QueryEthProphecyRequest) GetTokenContractAddress() string {
	if m != nil {
		return m.TokenContractAddress
	}
	return ""
}

func (m *QueryEthProphecyRequest) GetEthereumSender() string {
	if m != nil {
		return m.EthereumSender
	}
	return ""
}

// QueryEthProphecyResponse payload for EthProphecy rpc query
type QueryEthProphecyResponse struct {
	Id     string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status *types.Status     `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Claims []*EthBridgeClaim `protobuf:"bytes,3,rep,name=claims,proto3" json:"claims,omitempty"`
}

func (m *QueryEthProphecyResponse) Reset()         { *m = QueryEthProphecyResponse{} }
func (m *QueryEthProphecyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEthProphecyResponse) ProtoMessage()    {}
func (*QueryEthProphecyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7077edcf9f792b78, []int{1}
}
func (m *QueryEthProphecyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEthProphecyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEthProphecyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEthProphecyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEthProphecyResponse.Merge(m, src)
}
func (m *QueryEthProphecyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEthProphecyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEthProphecyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEthProphecyResponse proto.InternalMessageInfo

func (m *QueryEthProphecyResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *QueryEthProphecyResponse) GetStatus() *types.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *QueryEthProphecyResponse) GetClaims() []*EthBridgeClaim {
	if m != nil {
		return m.Claims
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryEthProphecyRequest)(nil), "sifnode.ethbridge.v1.QueryEthProphecyRequest")
	proto.RegisterType((*QueryEthProphecyResponse)(nil), "sifnode.ethbridge.v1.QueryEthProphecyResponse")
}

func init() { proto.RegisterFile("sifnode/ethbridge/v1/query.proto", fileDescriptor_7077edcf9f792b78) }

var fileDescriptor_7077edcf9f792b78 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0x96, 0x46, 0x9a, 0x2b, 0x15, 0x61, 0x95, 0x2d, 0x54, 0x22, 0x44, 0x11, 0xd2,
	0x2a, 0xd0, 0x12, 0xb5, 0x70, 0x42, 0x5c, 0x28, 0x4c, 0xdc, 0x10, 0xa4, 0x37, 0x2e, 0x55, 0xea,
	0xbc, 0x25, 0xd6, 0x1a, 0x3b, 0xb3, 0x9d, 0x8e, 0x7c, 0x0b, 0xee, 0x7c, 0x21, 0x8e, 0x3b, 0x72,
	0x42, 0x53, 0xfb, 0x0d, 0xf8, 0x04, 0xa8, 0x76, 0x3a, 0x55, 0x4b, 0x91, 0xb8, 0xc5, 0xef, 0xfd,
	0x9e, 0xdf, 0x7b, 0xff, 0xfc, 0x8d, 0x3c, 0x49, 0x2f, 0x18, 0x4f, 0x20, 0x04, 0x95, 0x2d, 0x04,
	0x4d, 0x52, 0x08, 0x57, 0xe3, 0xf0, 0xaa, 0x04, 0x51, 0x05, 0x85, 0xe0, 0x8a, 0xe3, 0x41, 0x4d,
	0x04, 0x77, 0x44, 0xb0, 0x1a, 0x0f, 0x07, 0x29, 0x4f, 0xb9, 0x06, 0xc2, 0xed, 0x97, 0x61, 0x87,
	0x87, 0x6f, 0x53, 0x55, 0x01, 0xb2, 0x26, 0x9e, 0xee, 0x08, 0x2e, 0x62, 0xb2, 0x6c, 0xa4, 0x5f,
	0x34, 0xd3, 0x0c, 0xd4, 0x35, 0x17, 0x97, 0xf3, 0x04, 0x24, 0x11, 0xb4, 0x50, 0x5c, 0x18, 0xd6,
	0xbf, 0x6d, 0xa3, 0x93, 0x2f, 0xdb, 0x41, 0xcf, 0x55, 0xf6, 0x59, 0xf0, 0x22, 0x03, 0x52, 0x45,
	0x70, 0x55, 0x82, 0x54, 0x78, 0x86, 0x70, 0xb3, 0xce, 0xb1, 0x3c, 0x6b, 0xd4, 0x9f, 0x3c, 0x0f,
	0x76, 0x1b, 0x99, 0x26, 0xc1, 0x6a, 0x1c, 0x7c, 0x32, 0xf0, 0x87, 0x3b, 0x36, 0x7a, 0xc4, 0xee,
	0x87, 0x30, 0x41, 0x27, 0x66, 0xa9, 0x39, 0xe1, 0x4c, 0x89, 0x98, 0xa8, 0x79, 0x9c, 0x24, 0x02,
	0xa4, 0x74, 0xda, 0x9e, 0x35, 0x3a, 0x9a, 0xbe, 0xfc, 0xf3, 0xfb, 0xd9, 0x69, 0x15, 0xe7, 0xcb,
	0x37, 0x7e, 0x0d, 0x0a, 0x48, 0xa9, 0x54, 0xa2, 0x6a, 0x54, 0xf8, 0xd1, 0x63, 0x83, 0xbc, 0xaf,
	0x13, 0xef, 0x4c, 0x1c, 0x0f, 0x50, 0x97, 0x71, 0x46, 0xc0, 0xe9, 0x78, 0xd6, 0xa8, 0x13, 0x99,
	0x03, 0x3e, 0x46, 0xb6, 0xac, 0xf2, 0x05, 0x5f, 0x3a, 0x0f, 0xb6, 0x9d, 0xa2, 0xfa, 0x84, 0x5f,
	0xa3, 0x63, 0xc5, 0x2f, 0x81, 0x35, 0x27, 0xea, 0x6a, 0x6e, 0xa0, 0xb3, 0xf7, 0x7b, 0x9c, 0xa2,
	0x87, 0xa0, 0x32, 0x10, 0x50, 0xe6, 0x73, 0x09, 0x2c, 0x01, 0xe1, 0xd8, 0x1a, 0xef, 0xef, 0xc2,
	0x33, 0x1d, 0xf5, 0x7f, 0x58, 0xc8, 0x69, 0x4a, 0x2c, 0x0b, 0xce, 0x24, 0xe0, 0x3e, 0x6a, 0xd3,
	0x44, 0x6b, 0x7a, 0x14, 0xb5, 0x69, 0x82, 0xc7, 0xc8, 0x96, 0x2a, 0x56, 0xa5, 0x51, 0xa3, 0x37,
	0x79, 0x72, 0x40, 0xe7, 0x99, 0x06, 0xa2, 0x1a, 0xc4, 0x6f, 0x91, 0x4d, 0x96, 0x31, 0xcd, 0xa5,
	0xd3, 0xf1, 0x3a, 0xa3, 0xde, 0xde, 0xaf, 0xd9, 0x37, 0x5b, 0x70, 0xae, 0xb2, 0xa9, 0x11, 0x6b,
	0x0b, 0x47, 0x75, 0xcd, 0xe4, 0x1a, 0x75, 0xf5, 0x70, 0x98, 0xa1, 0xde, 0xde, 0x80, 0xf8, 0xec,
	0xf0, 0x2d, 0xff, 0xf0, 0xca, 0x30, 0xf8, 0x5f, 0xdc, 0xec, 0xed, 0xb7, 0xa6, 0x1f, 0x7f, 0xae,
	0x5d, 0xeb, 0x66, 0xed, 0x5a, 0xb7, 0x6b, 0xd7, 0xfa, 0xbe, 0x71, 0x5b, 0x37, 0x1b, 0xb7, 0xf5,
	0x6b, 0xe3, 0xb6, 0xbe, 0x9e, 0xa5, 0x54, 0x65, 0xe5, 0x22, 0x20, 0x3c, 0x0f, 0x67, 0xf4, 0x82,
	0x64, 0x31, 0x65, 0xe1, 0xce, 0xd3, 0xdf, 0xf6, 0x9e, 0x85, 0x36, 0xfd, 0xc2, 0xd6, 0x4e, 0x7e,
	0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xaa, 0x36, 0x72, 0x86, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// EthProphecy queries an EthProphecy
	EthProphecy(ctx context.Context, in *QueryEthProphecyRequest, opts ...grpc.CallOption) (*QueryEthProphecyResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) EthProphecy(ctx context.Context, in *QueryEthProphecyRequest, opts ...grpc.CallOption) (*QueryEthProphecyResponse, error) {
	out := new(QueryEthProphecyResponse)
	err := c.cc.Invoke(ctx, "/sifnode.ethbridge.v1.Query/EthProphecy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// EthProphecy queries an EthProphecy
	EthProphecy(context.Context, *QueryEthProphecyRequest) (*QueryEthProphecyResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) EthProphecy(ctx context.Context, req *QueryEthProphecyRequest) (*QueryEthProphecyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EthProphecy not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_EthProphecy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEthProphecyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EthProphecy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sifnode.ethbridge.v1.Query/EthProphecy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EthProphecy(ctx, req.(*QueryEthProphecyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sifnode.ethbridge.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EthProphecy",
			Handler:    _Query_EthProphecy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sifnode/ethbridge/v1/query.proto",
}

func (m *QueryEthProphecyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEthProphecyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEthProphecyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EthereumSender) > 0 {
		i -= len(m.EthereumSender)
		copy(dAtA[i:], m.EthereumSender)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.EthereumSender)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.TokenContractAddress) > 0 {
		i -= len(m.TokenContractAddress)
		copy(dAtA[i:], m.TokenContractAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.TokenContractAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x22
	}
	if m.Nonce != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x18
	}
	if len(m.BridgeContractAddress) > 0 {
		i -= len(m.BridgeContractAddress)
		copy(dAtA[i:], m.BridgeContractAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BridgeContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.NetworkDescriptor != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.NetworkDescriptor))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryEthProphecyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEthProphecyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEthProphecyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Claims) > 0 {
		for iNdEx := len(m.Claims) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Claims[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Status != nil {
		{
			size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryEthProphecyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NetworkDescriptor != 0 {
		n += 1 + sovQuery(uint64(m.NetworkDescriptor))
	}
	l = len(m.BridgeContractAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovQuery(uint64(m.Nonce))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.TokenContractAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.EthereumSender)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryEthProphecyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	if len(m.Claims) > 0 {
		for _, e := range m.Claims {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryEthProphecyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEthProphecyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEthProphecyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkDescriptor", wireType)
			}
			m.NetworkDescriptor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NetworkDescriptor |= types.NetworkDescriptor(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BridgeContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthereumSender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthereumSender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryEthProphecyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEthProphecyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEthProphecyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &types.Status{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claims", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Claims = append(m.Claims, &EthBridgeClaim{})
			if err := m.Claims[len(m.Claims)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)