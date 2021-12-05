// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/bandoracle/v1beta1/gov.proto

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

type UpdateAdminProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	Address     string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
}

func (m *UpdateAdminProposal) Reset()         { *m = UpdateAdminProposal{} }
func (m *UpdateAdminProposal) String() string { return proto.CompactTextString(m) }
func (*UpdateAdminProposal) ProtoMessage()    {}
func (*UpdateAdminProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e186a48d46b50075, []int{0}
}
func (m *UpdateAdminProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateAdminProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateAdminProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateAdminProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAdminProposal.Merge(m, src)
}
func (m *UpdateAdminProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdateAdminProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAdminProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAdminProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UpdateAdminProposal)(nil), "comdex.bandoracle.v1beta1.UpdateAdminProposal")
}

func init() {
	proto.RegisterFile("comdex/bandoracle/v1beta1/gov.proto", fileDescriptor_e186a48d46b50075)
}

var fileDescriptor_e186a48d46b50075 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xd0, 0x41, 0x4a, 0xc4, 0x30,
	0x18, 0x05, 0xe0, 0x46, 0x51, 0xb1, 0x8a, 0x48, 0x14, 0xa9, 0x2e, 0x52, 0xa9, 0x20, 0x2e, 0xb4,
	0x61, 0xd0, 0x85, 0xb8, 0x73, 0x4e, 0x20, 0x05, 0x5d, 0xb8, 0x4b, 0x9b, 0x4c, 0x0d, 0xb4, 0xfd,
	0x4b, 0x12, 0x07, 0xe7, 0x16, 0x5e, 0x42, 0xf0, 0x28, 0xb3, 0x9c, 0xa5, 0xab, 0xa2, 0xed, 0x0d,
	0x7a, 0x02, 0x31, 0xa9, 0x38, 0xee, 0x92, 0xff, 0x7d, 0x6f, 0xf3, 0xfc, 0x93, 0x0c, 0x4a, 0x2e,
	0x5e, 0x68, 0xca, 0x2a, 0x0e, 0x8a, 0x65, 0x85, 0xa0, 0xd3, 0x51, 0x2a, 0x0c, 0x1b, 0xd1, 0x1c,
	0xa6, 0x71, 0xad, 0xc0, 0x00, 0x3e, 0x74, 0x28, 0xfe, 0x43, 0xf1, 0x80, 0x8e, 0xf6, 0x73, 0xc8,
	0xc1, 0x2a, 0xfa, 0xf3, 0x72, 0x85, 0xe8, 0x0d, 0xf9, 0x7b, 0xf7, 0x35, 0x67, 0x46, 0xdc, 0xf2,
	0x52, 0x56, 0x77, 0x0a, 0x6a, 0xd0, 0xac, 0xc0, 0xa7, 0xfe, 0x9a, 0x91, 0xa6, 0x10, 0x01, 0x3a,
	0x46, 0x67, 0x9b, 0xe3, 0xdd, 0xbe, 0x09, 0xb7, 0x67, 0xac, 0x2c, 0x6e, 0x22, 0x7b, 0x8e, 0x12,
	0x17, 0xe3, 0x6b, 0x7f, 0x8b, 0x0b, 0x9d, 0x29, 0x59, 0x1b, 0x09, 0x55, 0xb0, 0x62, 0xf5, 0x41,
	0xdf, 0x84, 0xd8, 0xe9, 0xa5, 0x30, 0x4a, 0x96, 0x29, 0x3e, 0xf7, 0x37, 0x18, 0xe7, 0x4a, 0x68,
	0x1d, 0xac, 0xda, 0x16, 0xee, 0x9b, 0x70, 0xc7, 0xb5, 0x86, 0x20, 0x4a, 0x7e, 0xc9, 0xf8, 0x61,
	0xfe, 0x45, 0xbc, 0xf7, 0x96, 0x78, 0xf3, 0x96, 0xa0, 0x45, 0x4b, 0xd0, 0x67, 0x4b, 0xd0, 0x6b,
	0x47, 0xbc, 0x45, 0x47, 0xbc, 0x8f, 0x8e, 0x78, 0x8f, 0x57, 0xb9, 0x34, 0x4f, 0xcf, 0x69, 0x9c,
	0x41, 0x49, 0xdd, 0x0a, 0x17, 0x30, 0x99, 0xc8, 0x4c, 0xb2, 0x62, 0xf8, 0xd3, 0x7f, 0xe3, 0x99,
	0x59, 0x2d, 0x74, 0xba, 0x6e, 0x67, 0xb8, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x04, 0x3b, 0xac,
	0x9e, 0x5e, 0x01, 0x00, 0x00,
}

func (m *UpdateAdminProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateAdminProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateAdminProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UpdateAdminProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UpdateAdminProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: UpdateAdminProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateAdminProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
