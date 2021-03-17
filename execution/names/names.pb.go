// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: names.proto

package names

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	github_com_klye-dev_hsc-main_crypto "github.com/KLYE-Dev/HSC-MAIN/crypto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// NameReg provides a global key value store based on Name, Data pairs that are subject to expiry and ownership by an
// account.
type Entry struct {
	// registered name for the entry
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// address that created the entry
	Owner github_com_klye-dev_hsc-main_crypto.Address `protobuf:"bytes,2,opt,name=Owner,proto3,customtype=github.com/KLYE-Dev/HSC-MAIN/crypto.Address" json:"Owner"`
	// data to store under this name
	Data string `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	// block at which this entry expires
	Expires              uint64   `protobuf:"varint,4,opt,name=Expires,proto3" json:"Expires,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entry) Reset()      { *m = Entry{} }
func (*Entry) ProtoMessage() {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4268625867c617c, []int{0}
}
func (m *Entry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return m.Size()
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Entry) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Entry) GetExpires() uint64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func (*Entry) XXX_MessageName() string {
	return "names.Entry"
}
func init() {
	proto.RegisterType((*Entry)(nil), "names.Entry")
	golang_proto.RegisterType((*Entry)(nil), "names.Entry")
}

func init() { proto.RegisterFile("names.proto", fileDescriptor_f4268625867c617c) }
func init() { golang_proto.RegisterFile("names.proto", fileDescriptor_f4268625867c617c) }

var fileDescriptor_f4268625867c617c = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x4b, 0xcc, 0x4d,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xa4, 0x44, 0xd2, 0xf3, 0xd3,
	0xf3, 0xc1, 0x22, 0xfa, 0x20, 0x16, 0x44, 0x52, 0x69, 0x36, 0x23, 0x17, 0xab, 0x6b, 0x5e, 0x49,
	0x51, 0xa5, 0x90, 0x10, 0x17, 0x8b, 0x5f, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x98, 0x2d, 0xe4, 0xc5, 0xc5, 0xea, 0x5f, 0x9e, 0x97, 0x5a, 0x24, 0xc1, 0xa4, 0xc0, 0xa8,
	0xc1, 0xe3, 0x64, 0x72, 0xe2, 0x9e, 0x3c, 0xc3, 0xad, 0x7b, 0xf2, 0x3a, 0xe9, 0x99, 0x25, 0x19,
	0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x19, 0x95, 0x05, 0xa9, 0x45, 0x39, 0xa9, 0x29, 0xe9,
	0xa9, 0x45, 0xfa, 0x49, 0xa5, 0x45, 0x45, 0xf9, 0xe5, 0xfa, 0xc9, 0x45, 0x95, 0x05, 0x25, 0xf9,
	0x7a, 0x8e, 0x29, 0x29, 0x45, 0xa9, 0xc5, 0xc5, 0x41, 0x10, 0x23, 0x40, 0xe6, 0xbb, 0x24, 0x96,
	0x24, 0x4a, 0x30, 0x43, 0xcc, 0x07, 0xb1, 0x85, 0x24, 0xb8, 0xd8, 0x5d, 0x2b, 0x0a, 0x32, 0x8b,
	0x52, 0x8b, 0x25, 0x58, 0x14, 0x18, 0x35, 0x58, 0x82, 0x60, 0x5c, 0x2b, 0x96, 0x19, 0x0b, 0xe4,
	0x19, 0x9c, 0xdc, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc6, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x0f, 0x3c, 0x96, 0x63, 0x3c, 0xf1, 0x58, 0x8e, 0x31, 0x4a, 0x17, 0xbf,
	0x13, 0x52, 0x2b, 0x52, 0x93, 0x4b, 0x4b, 0x32, 0xf3, 0xf3, 0xf4, 0xc1, 0x9e, 0x4f, 0x62, 0x03,
	0xfb, 0xd6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x78, 0xc9, 0x3e, 0x19, 0x01, 0x00, 0x00,
}

func (m *Entry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Entry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Entry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Expires != 0 {
		i = encodeVarintNames(dAtA, i, uint64(m.Expires))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintNames(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Owner.Size()
		i -= size
		if _, err := m.Owner.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintNames(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintNames(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNames(dAtA []byte, offset int, v uint64) int {
	offset -= sovNames(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Entry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNames(uint64(l))
	}
	l = m.Owner.Size()
	n += 1 + l + sovNames(uint64(l))
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovNames(uint64(l))
	}
	if m.Expires != 0 {
		n += 1 + sovNames(uint64(m.Expires))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovNames(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNames(x uint64) (n int) {
	return sovNames(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Entry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNames
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
			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNames
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
				return ErrInvalidLengthNames
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNames
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNames
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
				return ErrInvalidLengthNames
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNames
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Owner.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNames
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
				return ErrInvalidLengthNames
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNames
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expires", wireType)
			}
			m.Expires = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNames
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expires |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNames(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNames
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
func skipNames(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNames
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
					return 0, ErrIntOverflowNames
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
					return 0, ErrIntOverflowNames
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
				return 0, ErrInvalidLengthNames
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNames
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNames
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNames        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNames          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNames = fmt.Errorf("proto: unexpected end of group")
)
