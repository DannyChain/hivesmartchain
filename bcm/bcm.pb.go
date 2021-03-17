// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bcm.proto

package bcm

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_klye-dev_hsc-main_binary "github.com/klye-dev/hivesmartchain/binary"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SyncInfo struct {
	LatestBlockHeight uint64                                        `protobuf:"varint,1,opt,name=LatestBlockHeight,proto3" json:"LatestBlockHeight,omitempty"`
	LatestBlockHash   github_com_klye-dev_hsc-main_binary.HexBytes `protobuf:"bytes,2,opt,name=LatestBlockHash,proto3,customtype=github.com/klye-dev/hivesmartchain/binary.HexBytes" json:"LatestBlockHash"`
	LatestAppHash     github_com_klye-dev_hsc-main_binary.HexBytes `protobuf:"bytes,3,opt,name=LatestAppHash,proto3,customtype=github.com/klye-dev/hivesmartchain/binary.HexBytes" json:"LatestAppHash"`
	// Timestamp of block as set by the block proposer
	LatestBlockTime time.Time `protobuf:"bytes,4,opt,name=LatestBlockTime,proto3,stdtime" json:"LatestBlockTime"`
	// Time at which we committed the last block
	LatestBlockSeenTime time.Time `protobuf:"bytes,5,opt,name=LatestBlockSeenTime,proto3,stdtime" json:"LatestBlockSeenTime"`
	// Time elapsed since last commit
	LatestBlockDuration  time.Duration `protobuf:"bytes,6,opt,name=LatestBlockDuration,proto3,stdduration" json:"LatestBlockDuration"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SyncInfo) Reset()         { *m = SyncInfo{} }
func (m *SyncInfo) String() string { return proto.CompactTextString(m) }
func (*SyncInfo) ProtoMessage()    {}
func (*SyncInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c9ff3e1ca1cc0f1, []int{0}
}
func (m *SyncInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SyncInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SyncInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncInfo.Merge(m, src)
}
func (m *SyncInfo) XXX_Size() int {
	return m.Size()
}
func (m *SyncInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SyncInfo proto.InternalMessageInfo

func (m *SyncInfo) GetLatestBlockHeight() uint64 {
	if m != nil {
		return m.LatestBlockHeight
	}
	return 0
}

func (m *SyncInfo) GetLatestBlockTime() time.Time {
	if m != nil {
		return m.LatestBlockTime
	}
	return time.Time{}
}

func (m *SyncInfo) GetLatestBlockSeenTime() time.Time {
	if m != nil {
		return m.LatestBlockSeenTime
	}
	return time.Time{}
}

func (m *SyncInfo) GetLatestBlockDuration() time.Duration {
	if m != nil {
		return m.LatestBlockDuration
	}
	return 0
}

func (*SyncInfo) XXX_MessageName() string {
	return "bcm.SyncInfo"
}

type PersistedState struct {
	AppHashAfterLastBlock github_com_klye-dev_hsc-main_binary.HexBytes `protobuf:"bytes,1,opt,name=AppHashAfterLastBlock,proto3,customtype=github.com/klye-dev/hivesmartchain/binary.HexBytes" json:"AppHashAfterLastBlock"`
	LastBlockTime         time.Time                                     `protobuf:"bytes,2,opt,name=LastBlockTime,proto3,stdtime" json:"LastBlockTime"`
	LastBlockHeight       uint64                                        `protobuf:"varint,3,opt,name=LastBlockHeight,proto3" json:"LastBlockHeight,omitempty"`
	GenesisHash           github_com_klye-dev_hsc-main_binary.HexBytes `protobuf:"bytes,4,opt,name=GenesisHash,proto3,customtype=github.com/klye-dev/hivesmartchain/binary.HexBytes" json:"GenesisHash"`
	XXX_NoUnkeyedLiteral  struct{}                                      `json:"-"`
	XXX_unrecognized      []byte                                        `json:"-"`
	XXX_sizecache         int32                                         `json:"-"`
}

func (m *PersistedState) Reset()         { *m = PersistedState{} }
func (m *PersistedState) String() string { return proto.CompactTextString(m) }
func (*PersistedState) ProtoMessage()    {}
func (*PersistedState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c9ff3e1ca1cc0f1, []int{1}
}
func (m *PersistedState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PersistedState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *PersistedState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersistedState.Merge(m, src)
}
func (m *PersistedState) XXX_Size() int {
	return m.Size()
}
func (m *PersistedState) XXX_DiscardUnknown() {
	xxx_messageInfo_PersistedState.DiscardUnknown(m)
}

var xxx_messageInfo_PersistedState proto.InternalMessageInfo

func (m *PersistedState) GetLastBlockTime() time.Time {
	if m != nil {
		return m.LastBlockTime
	}
	return time.Time{}
}

func (m *PersistedState) GetLastBlockHeight() uint64 {
	if m != nil {
		return m.LastBlockHeight
	}
	return 0
}

func (*PersistedState) XXX_MessageName() string {
	return "bcm.PersistedState"
}
func init() {
	proto.RegisterType((*SyncInfo)(nil), "bcm.SyncInfo")
	golang_proto.RegisterType((*SyncInfo)(nil), "bcm.SyncInfo")
	proto.RegisterType((*PersistedState)(nil), "bcm.PersistedState")
	golang_proto.RegisterType((*PersistedState)(nil), "bcm.PersistedState")
}

func init() { proto.RegisterFile("bcm.proto", fileDescriptor_0c9ff3e1ca1cc0f1) }
func init() { golang_proto.RegisterFile("bcm.proto", fileDescriptor_0c9ff3e1ca1cc0f1) }

var fileDescriptor_0c9ff3e1ca1cc0f1 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4f, 0x8f, 0x93, 0x40,
	0x18, 0xc6, 0x9d, 0x52, 0x37, 0xeb, 0xac, 0x7f, 0xe2, 0xa8, 0x09, 0xf6, 0x00, 0x75, 0x4f, 0x1c,
	0x14, 0x12, 0x8d, 0x27, 0x4f, 0x4b, 0x4c, 0x5c, 0xcd, 0xc6, 0x98, 0x76, 0xd5, 0x44, 0x0f, 0x06,
	0xe8, 0x5b, 0x98, 0x6c, 0x61, 0xc8, 0xcc, 0x10, 0xe5, 0x5b, 0x78, 0xf4, 0xe3, 0x78, 0xec, 0xc1,
	0x83, 0x47, 0xe3, 0xa1, 0x1a, 0xfa, 0x19, 0xbc, 0x1b, 0x06, 0x48, 0x80, 0x36, 0x31, 0xb5, 0x37,
	0xe6, 0x7d, 0xde, 0xf7, 0xc7, 0xcc, 0xf3, 0x3e, 0xf8, 0x8a, 0x1f, 0xc4, 0x76, 0xca, 0x99, 0x64,
	0x44, 0xf3, 0x83, 0x78, 0x74, 0x3b, 0x64, 0x21, 0x53, 0x67, 0xa7, 0xfc, 0xaa, 0xa4, 0x91, 0x19,
	0x32, 0x16, 0x2e, 0xc0, 0x51, 0x27, 0x3f, 0x9b, 0x3b, 0x92, 0xc6, 0x20, 0xa4, 0x17, 0xa7, 0x75,
	0x83, 0xd1, 0x6f, 0x98, 0x65, 0xdc, 0x93, 0x94, 0x25, 0x95, 0x7e, 0xfc, 0x47, 0xc3, 0x87, 0xd3,
	0x3c, 0x09, 0x9e, 0x27, 0x73, 0x46, 0xee, 0xe3, 0x9b, 0x67, 0x9e, 0x04, 0x21, 0xdd, 0x05, 0x0b,
	0x2e, 0x4e, 0x81, 0x86, 0x91, 0xd4, 0xd1, 0x18, 0x59, 0xc3, 0xc9, 0xa6, 0x40, 0x3e, 0xe0, 0x1b,
	0xed, 0xa2, 0x27, 0x22, 0x7d, 0x30, 0x46, 0xd6, 0x55, 0xf7, 0xf1, 0x72, 0x65, 0x5e, 0xfa, 0xb9,
	0x32, 0x1f, 0x84, 0x54, 0x46, 0x99, 0x6f, 0x07, 0x2c, 0x76, 0xa2, 0x3c, 0x05, 0xbe, 0x80, 0x59,
	0x08, 0xdc, 0xf1, 0x33, 0xce, 0xd9, 0x47, 0xc7, 0xa7, 0x89, 0xc7, 0x73, 0xfb, 0x14, 0x3e, 0xb9,
	0xb9, 0x04, 0x31, 0xe9, 0xd3, 0xc8, 0x7b, 0x7c, 0xad, 0x2a, 0x9d, 0xa4, 0xa9, 0xc2, 0x6b, 0xfb,
	0xe0, 0xbb, 0x2c, 0xf2, 0xb2, 0x73, 0xfb, 0x73, 0x1a, 0x83, 0x3e, 0x1c, 0x23, 0xeb, 0xe8, 0xe1,
	0xc8, 0xae, 0x2c, 0xb3, 0x1b, 0xcb, 0xec, 0xf3, 0xc6, 0x53, 0xf7, 0xb0, 0xfc, 0xf5, 0xe7, 0x5f,
	0x26, 0x9a, 0xf4, 0x87, 0xc9, 0x1b, 0x7c, 0xab, 0x55, 0x9a, 0x02, 0x24, 0x8a, 0x79, 0x79, 0x07,
	0xe6, 0x36, 0x00, 0x79, 0xdd, 0xe1, 0x3e, 0xad, 0xb7, 0xa7, 0x1f, 0x28, 0xee, 0xdd, 0x0d, 0x6e,
	0xd3, 0x50, 0x61, 0xbf, 0xf4, 0xb1, 0x8d, 0x7c, 0xfc, 0x6d, 0x80, 0xaf, 0xbf, 0x02, 0x2e, 0xa8,
	0x90, 0x30, 0x9b, 0x4a, 0x4f, 0x02, 0xb9, 0xc0, 0x77, 0x6a, 0x73, 0x4e, 0xe6, 0x12, 0xf8, 0x99,
	0x57, 0xcf, 0xa8, 0x04, 0xfc, 0xb7, 0xed, 0xdb, 0x99, 0xe4, 0x45, 0xb9, 0xdb, 0xb6, 0xf9, 0x83,
	0x1d, 0x8c, 0xea, 0x8e, 0x12, 0xab, 0x5c, 0x65, 0x37, 0xb4, 0x9a, 0x0a, 0x6d, 0xbf, 0x4c, 0xde,
	0xe2, 0xa3, 0x67, 0x90, 0x80, 0xa0, 0x42, 0xe5, 0x69, 0xb8, 0xcf, 0xc3, 0xda, 0x24, 0xf7, 0xc9,
	0xb2, 0x30, 0xd0, 0xf7, 0xc2, 0x40, 0x3f, 0x0a, 0x03, 0xfd, 0x2e, 0x0c, 0xf4, 0x75, 0x6d, 0xa0,
	0xe5, 0xda, 0x40, 0xef, 0xee, 0xfd, 0x83, 0x1a, 0xc4, 0xfe, 0x81, 0x7a, 0xec, 0xa3, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x6e, 0x6b, 0x06, 0x93, 0xf3, 0x03, 0x00, 0x00,
}

func (m *SyncInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SyncInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.LatestBlockDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.LatestBlockDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintBcm(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LatestBlockSeenTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LatestBlockSeenTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintBcm(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LatestBlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LatestBlockTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintBcm(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x22
	{
		size := m.LatestAppHash.Size()
		i -= size
		if _, err := m.LatestAppHash.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBcm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.LatestBlockHash.Size()
		i -= size
		if _, err := m.LatestBlockHash.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBcm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.LatestBlockHeight != 0 {
		i = encodeVarintBcm(dAtA, i, uint64(m.LatestBlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PersistedState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PersistedState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PersistedState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	{
		size := m.GenesisHash.Size()
		i -= size
		if _, err := m.GenesisHash.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBcm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.LastBlockHeight != 0 {
		i = encodeVarintBcm(dAtA, i, uint64(m.LastBlockHeight))
		i--
		dAtA[i] = 0x18
	}
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastBlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastBlockTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintBcm(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x12
	{
		size := m.AppHashAfterLastBlock.Size()
		i -= size
		if _, err := m.AppHashAfterLastBlock.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBcm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintBcm(dAtA []byte, offset int, v uint64) int {
	offset -= sovBcm(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SyncInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LatestBlockHeight != 0 {
		n += 1 + sovBcm(uint64(m.LatestBlockHeight))
	}
	l = m.LatestBlockHash.Size()
	n += 1 + l + sovBcm(uint64(l))
	l = m.LatestAppHash.Size()
	n += 1 + l + sovBcm(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LatestBlockTime)
	n += 1 + l + sovBcm(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LatestBlockSeenTime)
	n += 1 + l + sovBcm(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.LatestBlockDuration)
	n += 1 + l + sovBcm(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PersistedState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.AppHashAfterLastBlock.Size()
	n += 1 + l + sovBcm(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastBlockTime)
	n += 1 + l + sovBcm(uint64(l))
	if m.LastBlockHeight != 0 {
		n += 1 + sovBcm(uint64(m.LastBlockHeight))
	}
	l = m.GenesisHash.Size()
	n += 1 + l + sovBcm(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBcm(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBcm(x uint64) (n int) {
	return sovBcm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SyncInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBcm
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
			return fmt.Errorf("proto: SyncInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestBlockHeight", wireType)
			}
			m.LatestBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestBlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LatestBlockHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestAppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LatestAppHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestBlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBcm
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
				return ErrInvalidLengthBcm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LatestBlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestBlockSeenTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBcm
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
				return ErrInvalidLengthBcm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LatestBlockSeenTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestBlockDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBcm
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
				return ErrInvalidLengthBcm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.LatestBlockDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBcm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBcm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PersistedState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBcm
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
			return fmt.Errorf("proto: PersistedState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PersistedState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppHashAfterLastBlock", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AppHashAfterLastBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBcm
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
				return ErrInvalidLengthBcm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastBlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockHeight", wireType)
			}
			m.LastBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GenesisHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBcm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBcm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBcm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBcm
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
					return 0, ErrIntOverflowBcm
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
					return 0, ErrIntOverflowBcm
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
				return 0, ErrInvalidLengthBcm
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBcm
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBcm
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBcm        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBcm          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBcm = fmt.Errorf("proto: unexpected end of group")
)
