// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/locker/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type GenesisState struct {
	Lockers                          []Locker                           `protobuf:"bytes,1,rep,name=lockers,proto3" json:"lockers" yaml:"lockers"`
	LockerProductAssetMapping        []LockerProductAssetMapping        `protobuf:"bytes,2,rep,name=lockerProductAssetMapping,proto3" json:"lockerProductAssetMapping" yaml:"lockerProductAssetMapping"`
	LockerTotalRewardsByAssetAppWise []LockerTotalRewardsByAssetAppWise `protobuf:"bytes,3,rep,name=lockerTotalRewardsByAssetAppWise,proto3" json:"lockerTotalRewardsByAssetAppWise" yaml:"lockerTotalRewardsByAssetAppWise"`
	LockerLookupTable                []LockerLookupTableData            `protobuf:"bytes,4,rep,name=lockerLookupTable,proto3" json:"lockerLookupTable" yaml:"lockerLookupTable"`
	UserLockerAssetMapping           []UserAppAssetLockerMapping        `protobuf:"bytes,5,rep,name=userLockerAssetMapping,proto3" json:"userLockerAssetMapping" yaml:"userLockerAssetMapping"`
	Params                           Params                             `protobuf:"bytes,6,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bd3c61b4d1357a2, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "fury.locker.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("fury/locker/v1beta1/genesis.proto", fileDescriptor_4bd3c61b4d1357a2)
}

var fileDescriptor_4bd3c61b4d1357a2 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x8a, 0xd4, 0x30,
	0x1c, 0xc6, 0x1b, 0x77, 0x9d, 0x85, 0x28, 0x82, 0x45, 0xa5, 0x2e, 0x6c, 0x5a, 0x23, 0xe2, 0x1e,
	0xb4, 0xdd, 0xd5, 0x83, 0xa0, 0xa7, 0x29, 0x82, 0x97, 0x15, 0x96, 0xb8, 0x22, 0xec, 0x2d, 0xed,
	0x64, 0x6b, 0xd9, 0xce, 0x26, 0x34, 0xa9, 0x3a, 0xe0, 0x0b, 0x78, 0xd2, 0xa3, 0x8f, 0xe0, 0x03,
	0xf8, 0x10, 0x73, 0xdc, 0xa3, 0xa7, 0xa2, 0x9d, 0x37, 0xd8, 0x27, 0x90, 0x26, 0x19, 0xa9, 0xd8,
	0xd6, 0x5b, 0xd3, 0xfc, 0xfe, 0xdf, 0xf7, 0x83, 0x24, 0xf0, 0x6e, 0xca, 0xe7, 0x33, 0xf6, 0x21,
	0x2a, 0x78, 0x7a, 0xca, 0xca, 0xe8, 0xdd, 0x7e, 0xc2, 0x14, 0xdd, 0x8f, 0x32, 0x76, 0xc6, 0x64,
	0x2e, 0x43, 0x51, 0x72, 0xc5, 0xdd, 0x9b, 0x06, 0x0a, 0x0d, 0x14, 0x5a, 0x68, 0xfb, 0x46, 0xc6,
	0x33, 0xae, 0x89, 0xa8, 0xfd, 0x32, 0xf0, 0x36, 0xee, 0x4f, 0xb4, 0xb3, 0xa3, 0x8c, 0xa0, 0x25,
	0x9d, 0xdb, 0x52, 0xfc, 0x69, 0x02, 0xaf, 0xbe, 0x30, 0x1a, 0xaf, 0x14, 0x55, 0xcc, 0x3d, 0x86,
	0x5b, 0x86, 0x97, 0x1e, 0x08, 0x36, 0x76, 0xaf, 0x3c, 0xda, 0x09, 0x7b, 0xbd, 0xc2, 0x03, 0xbd,
	0x8c, 0xef, 0x2c, 0x6b, 0xdf, 0x69, 0x6a, 0x7f, 0xcb, 0xac, 0xe5, 0x45, 0xed, 0x5f, 0x5b, 0xd0,
	0x79, 0xf1, 0x14, 0xdb, 0x18, 0x4c, 0xd6, 0x81, 0xee, 0x57, 0x00, 0x6f, 0x9b, 0xef, 0xc3, 0x92,
	0xcf, 0xaa, 0x54, 0x4d, 0xa5, 0x64, 0xea, 0x25, 0x15, 0x22, 0x3f, 0xcb, 0xbc, 0x4b, 0xba, 0x6e,
	0x6f, 0xb4, 0xae, 0x67, 0x2e, 0xde, 0x6d, 0x0d, 0x2e, 0x6a, 0x3f, 0xe8, 0xd6, 0xf6, 0x80, 0x98,
	0x0c, 0x97, 0xbb, 0xdf, 0x01, 0x0c, 0xcc, 0xee, 0x11, 0x57, 0xb4, 0x20, 0xec, 0x3d, 0x2d, 0x67,
	0x32, 0x5e, 0x68, 0x68, 0x2a, 0xc4, 0x9b, 0x5c, 0x32, 0x6f, 0x43, 0x1b, 0x3e, 0x19, 0x35, 0x1c,
	0x1e, 0x8f, 0x23, 0x2b, 0x7a, 0xbf, 0x2b, 0x3a, 0xcc, 0x63, 0xf2, 0x5f, 0x23, 0xf7, 0x23, 0xbc,
	0x6e, 0x98, 0x03, 0xce, 0x4f, 0x2b, 0x71, 0x44, 0x93, 0x82, 0x79, 0x9b, 0x5a, 0xf3, 0xc1, 0xa8,
	0x66, 0x87, 0x7f, 0x4e, 0x15, 0x8d, 0x03, 0xeb, 0xe6, 0x75, 0xdd, 0x3a, 0x10, 0x26, 0xff, 0x16,
	0xb9, 0x9f, 0x01, 0xbc, 0x55, 0xc9, 0xf6, 0x5f, 0xbb, 0xf3, 0xd7, 0x61, 0x5e, 0x1e, 0x3d, 0xcc,
	0xd7, 0x92, 0x95, 0x53, 0x21, 0xf4, 0x84, 0x19, 0x5e, 0x1f, 0xe6, 0x3d, 0xeb, 0xb1, 0x63, 0x3c,
	0xfa, 0xd3, 0x31, 0x19, 0xa8, 0x75, 0x9f, 0xc1, 0x89, 0xb9, 0xde, 0xde, 0x24, 0x00, 0x23, 0x97,
	0xf7, 0x50, 0x43, 0xf1, 0x66, 0xdb, 0x46, 0xec, 0x48, 0x4c, 0x96, 0xbf, 0x90, 0xf3, 0xad, 0x41,
	0xce, 0xb2, 0x41, 0xe0, 0xbc, 0x41, 0xe0, 0x67, 0x83, 0xc0, 0x97, 0x15, 0x72, 0xce, 0x57, 0xc8,
	0xf9, 0xb1, 0x42, 0xce, 0xf1, 0x5e, 0x96, 0xab, 0xb7, 0x55, 0xd2, 0x86, 0x46, 0x26, 0xf8, 0x21,
	0x3f, 0x39, 0xc9, 0xd3, 0x9c, 0x16, 0x76, 0x1d, 0xfd, 0x79, 0x6e, 0x6a, 0x21, 0x98, 0x4c, 0x26,
	0xfa, 0x99, 0x3d, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x1c, 0x09, 0x50, 0x02, 0x04, 0x00,
	0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.UserLockerAssetMapping) > 0 {
		for iNdEx := len(m.UserLockerAssetMapping) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserLockerAssetMapping[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.LockerLookupTable) > 0 {
		for iNdEx := len(m.LockerLookupTable) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LockerLookupTable[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.LockerTotalRewardsByAssetAppWise) > 0 {
		for iNdEx := len(m.LockerTotalRewardsByAssetAppWise) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LockerTotalRewardsByAssetAppWise[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.LockerProductAssetMapping) > 0 {
		for iNdEx := len(m.LockerProductAssetMapping) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LockerProductAssetMapping[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Lockers) > 0 {
		for iNdEx := len(m.Lockers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Lockers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Lockers) > 0 {
		for _, e := range m.Lockers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LockerProductAssetMapping) > 0 {
		for _, e := range m.LockerProductAssetMapping {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LockerTotalRewardsByAssetAppWise) > 0 {
		for _, e := range m.LockerTotalRewardsByAssetAppWise {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LockerLookupTable) > 0 {
		for _, e := range m.LockerLookupTable {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UserLockerAssetMapping) > 0 {
		for _, e := range m.UserLockerAssetMapping {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lockers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lockers = append(m.Lockers, Locker{})
			if err := m.Lockers[len(m.Lockers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockerProductAssetMapping", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockerProductAssetMapping = append(m.LockerProductAssetMapping, LockerProductAssetMapping{})
			if err := m.LockerProductAssetMapping[len(m.LockerProductAssetMapping)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockerTotalRewardsByAssetAppWise", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockerTotalRewardsByAssetAppWise = append(m.LockerTotalRewardsByAssetAppWise, LockerTotalRewardsByAssetAppWise{})
			if err := m.LockerTotalRewardsByAssetAppWise[len(m.LockerTotalRewardsByAssetAppWise)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockerLookupTable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockerLookupTable = append(m.LockerLookupTable, LockerLookupTableData{})
			if err := m.LockerLookupTable[len(m.LockerLookupTable)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserLockerAssetMapping", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserLockerAssetMapping = append(m.UserLockerAssetMapping, UserAppAssetLockerMapping{})
			if err := m.UserLockerAssetMapping[len(m.UserLockerAssetMapping)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
