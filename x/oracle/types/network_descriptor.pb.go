// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sifnode/oracle/v1/network_descriptor.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

// NetworkDescriptor is a unique identifier for all chains that Sifchain
// enables.
type NetworkDescriptor int32

const (
	// Not currently in use
	NetworkDescriptor_NETWORK_DESCRIPTOR_UNSPECIFIED NetworkDescriptor = 0
	// https://ethereum.org
	NetworkDescriptor_NETWORK_DESCRIPTOR_ETHEREUM NetworkDescriptor = 1
	// Bitcoin mainnet
	NetworkDescriptor_NETWORK_DESCRIPTOR_BITCOIN NetworkDescriptor = 2
	// https://github.com/ethereum/ropsten
	NetworkDescriptor_NETWORK_DESCRIPTOR_ETHEREUM_TESTNET_ROPSTEN NetworkDescriptor = 3
	// https://www.binance.org
	NetworkDescriptor_NETWORK_DESCRIPTOR_BINANCE_SMART_CHAIN NetworkDescriptor = 56
	// https://testnet.binance.org/
	NetworkDescriptor_NETWORK_DESCRIPTOR_BINANCE_SMART_CHAIN_TESTNET NetworkDescriptor = 97
	// Ganache local testnet
	NetworkDescriptor_NETWORK_DESCRIPTOR_GANACHE NetworkDescriptor = 5777
	// Hardhat local testnet
	NetworkDescriptor_NETWORK_DESCRIPTOR_HARDHAT NetworkDescriptor = 31337
)

var NetworkDescriptor_name = map[int32]string{
	0:     "NETWORK_DESCRIPTOR_UNSPECIFIED",
	1:     "NETWORK_DESCRIPTOR_ETHEREUM",
	2:     "NETWORK_DESCRIPTOR_BITCOIN",
	3:     "NETWORK_DESCRIPTOR_ETHEREUM_TESTNET_ROPSTEN",
	56:    "NETWORK_DESCRIPTOR_BINANCE_SMART_CHAIN",
	97:    "NETWORK_DESCRIPTOR_BINANCE_SMART_CHAIN_TESTNET",
	5777:  "NETWORK_DESCRIPTOR_GANACHE",
	31337: "NETWORK_DESCRIPTOR_HARDHAT",
}

var NetworkDescriptor_value = map[string]int32{
	"NETWORK_DESCRIPTOR_UNSPECIFIED":                 0,
	"NETWORK_DESCRIPTOR_ETHEREUM":                    1,
	"NETWORK_DESCRIPTOR_BITCOIN":                     2,
	"NETWORK_DESCRIPTOR_ETHEREUM_TESTNET_ROPSTEN":    3,
	"NETWORK_DESCRIPTOR_BINANCE_SMART_CHAIN":         56,
	"NETWORK_DESCRIPTOR_BINANCE_SMART_CHAIN_TESTNET": 97,
	"NETWORK_DESCRIPTOR_GANACHE":                     5777,
	"NETWORK_DESCRIPTOR_HARDHAT":                     31337,
}

func (x NetworkDescriptor) String() string {
	return proto.EnumName(NetworkDescriptor_name, int32(x))
}

func (NetworkDescriptor) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_721e8ae3af4d5f0a, []int{0}
}

func init() {
	proto.RegisterEnum("sifnode.oracle.v1.NetworkDescriptor", NetworkDescriptor_name, NetworkDescriptor_value)
}

func init() {
	proto.RegisterFile("sifnode/oracle/v1/network_descriptor.proto", fileDescriptor_721e8ae3af4d5f0a)
}

var fileDescriptor_721e8ae3af4d5f0a = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd1, 0xb1, 0x4a, 0xc3, 0x40,
	0x1c, 0xc7, 0xf1, 0xa4, 0x82, 0xc3, 0x4d, 0xd7, 0xc3, 0xa9, 0xc2, 0x55, 0x1c, 0x1c, 0x22, 0xe6,
	0xa8, 0x2e, 0xae, 0xd7, 0xe4, 0xaf, 0x09, 0xd2, 0x4b, 0xb9, 0x5c, 0x11, 0x5c, 0x8e, 0x36, 0x4d,
	0xdb, 0xa0, 0xf6, 0x4a, 0x1a, 0xab, 0xbe, 0x85, 0xbe, 0x95, 0x83, 0x43, 0x47, 0x47, 0x69, 0x27,
	0x77, 0x1f, 0x40, 0x6c, 0xad, 0x53, 0x10, 0xb7, 0x83, 0xfb, 0xf0, 0xbd, 0x83, 0x1f, 0x72, 0xa6,
	0xd9, 0x60, 0x6c, 0xfa, 0x29, 0x33, 0x79, 0x37, 0xb9, 0x49, 0xd9, 0xac, 0xc1, 0xc6, 0x69, 0x71,
	0x6f, 0xf2, 0x6b, 0xdd, 0x4f, 0xa7, 0x49, 0x9e, 0x4d, 0x0a, 0x93, 0xbb, 0x93, 0xdc, 0x14, 0x86,
	0x54, 0x7f, 0xac, 0xbb, 0xb6, 0xee, 0xac, 0x51, 0xdb, 0x19, 0x9a, 0xa1, 0x59, 0xdd, 0xb2, 0xef,
	0xd3, 0x1a, 0x3a, 0xaf, 0x15, 0x54, 0x15, 0xeb, 0x8a, 0xff, 0x1b, 0x21, 0xfb, 0x88, 0x0a, 0x50,
	0x97, 0x91, 0xbc, 0xd0, 0x3e, 0xc4, 0x9e, 0x0c, 0xdb, 0x2a, 0x92, 0xba, 0x23, 0xe2, 0x36, 0x78,
	0xe1, 0x59, 0x08, 0x3e, 0xb6, 0x48, 0x1d, 0xed, 0x96, 0x18, 0x50, 0x01, 0x48, 0xe8, 0xb4, 0xb0,
	0x4d, 0x28, 0xaa, 0x95, 0x80, 0x66, 0xa8, 0xbc, 0x28, 0x14, 0xb8, 0x42, 0x18, 0x3a, 0xfc, 0x23,
	0xa0, 0x15, 0xc4, 0x4a, 0x80, 0xd2, 0x32, 0x6a, 0xc7, 0x0a, 0x04, 0xde, 0x22, 0x0e, 0x3a, 0x28,
	0x0d, 0x0a, 0x2e, 0x3c, 0xd0, 0x71, 0x8b, 0x4b, 0xa5, 0xbd, 0x80, 0x87, 0x02, 0x9f, 0x92, 0x63,
	0xe4, 0xfe, 0xcf, 0x6e, 0xde, 0xc1, 0x5d, 0x52, 0x2f, 0xfd, 0xf0, 0x39, 0x17, 0xdc, 0x0b, 0x00,
	0x3f, 0x1f, 0x91, 0xbd, 0x52, 0x10, 0x70, 0xe9, 0x07, 0x5c, 0xe1, 0x8f, 0x4f, 0xbb, 0xe9, 0xbf,
	0x2c, 0xa8, 0x3d, 0x5f, 0x50, 0xfb, 0x7d, 0x41, 0xed, 0xa7, 0x25, 0xb5, 0xe6, 0x4b, 0x6a, 0xbd,
	0x2d, 0xa9, 0x75, 0xe5, 0x0c, 0xb3, 0x62, 0x74, 0xd7, 0x73, 0x13, 0x73, 0xcb, 0xe2, 0x6c, 0x90,
	0x8c, 0xba, 0xd9, 0x98, 0x6d, 0x16, 0x7d, 0xd8, 0x6c, 0x5a, 0x3c, 0x4e, 0xd2, 0x69, 0x6f, 0x7b,
	0xb5, 0xcd, 0xc9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x7b, 0x4b, 0xd2, 0xf2, 0x01, 0x00,
	0x00,
}