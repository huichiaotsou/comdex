// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/bandoracle/v1beta1/gold_price.proto

package types

import (
	fmt "fmt"
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

type GoldPriceCallData struct {
	Multiplier uint64 `protobuf:"varint,1,opt,name=multiplier,proto3" json:"multiplier,omitempty"`
}

func (m *GoldPriceCallData) Reset()         { *m = GoldPriceCallData{} }
func (m *GoldPriceCallData) String() string { return proto.CompactTextString(m) }
func (*GoldPriceCallData) ProtoMessage()    {}
func (*GoldPriceCallData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a57bb37bded58289, []int{0}
}
func (m *GoldPriceCallData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoldPriceCallData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoldPriceCallData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoldPriceCallData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoldPriceCallData.Merge(m, src)
}
func (m *GoldPriceCallData) XXX_Size() int {
	return m.Size()
}
func (m *GoldPriceCallData) XXX_DiscardUnknown() {
	xxx_messageInfo_GoldPriceCallData.DiscardUnknown(m)
}

var xxx_messageInfo_GoldPriceCallData proto.InternalMessageInfo

func (m *GoldPriceCallData) GetMultiplier() uint64 {
	if m != nil {
		return m.Multiplier
	}
	return 0
}

type GoldPriceResult struct {
	Price uint64 `protobuf:"varint,1,opt,name=price,proto3" json:"price,omitempty"`
}

func (m *GoldPriceResult) Reset()         { *m = GoldPriceResult{} }
func (m *GoldPriceResult) String() string { return proto.CompactTextString(m) }
func (*GoldPriceResult) ProtoMessage()    {}
func (*GoldPriceResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a57bb37bded58289, []int{1}
}
func (m *GoldPriceResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoldPriceResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoldPriceResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoldPriceResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoldPriceResult.Merge(m, src)
}
func (m *GoldPriceResult) XXX_Size() int {
	return m.Size()
}
func (m *GoldPriceResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GoldPriceResult.DiscardUnknown(m)
}

var xxx_messageInfo_GoldPriceResult proto.InternalMessageInfo

func (m *GoldPriceResult) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto.RegisterType((*GoldPriceCallData)(nil), "comdex.bandoracle.v1beta1.GoldPriceCallData")
	proto.RegisterType((*GoldPriceResult)(nil), "comdex.bandoracle.v1beta1.GoldPriceResult")
}

func init() {
	proto.RegisterFile("comdex/bandoracle/v1beta1/gold_price.proto", fileDescriptor_a57bb37bded58289)
}

var fileDescriptor_a57bb37bded58289 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4a, 0xce, 0xcf, 0x4d,
	0x49, 0xad, 0xd0, 0x4f, 0x4a, 0xcc, 0x4b, 0xc9, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x2f, 0x33,
	0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x4f, 0xcf, 0xcf, 0x49, 0x89, 0x2f, 0x28, 0xca, 0x4c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x84, 0xa8, 0xd5, 0x43, 0xa8, 0xd5, 0x83, 0xaa,
	0x55, 0x32, 0xe6, 0x12, 0x74, 0xcf, 0xcf, 0x49, 0x09, 0x00, 0xa9, 0x76, 0x4e, 0xcc, 0xc9, 0x71,
	0x49, 0x2c, 0x49, 0x14, 0x92, 0xe3, 0xe2, 0xca, 0x2d, 0xcd, 0x29, 0xc9, 0x2c, 0xc8, 0xc9, 0x4c,
	0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0x42, 0x12, 0x51, 0x52, 0xe7, 0xe2, 0x87, 0x6b,
	0x0a, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x11, 0x12, 0xe1, 0x62, 0x05, 0xdb, 0x08, 0x55, 0x0d, 0xe1,
	0x38, 0xf9, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13,
	0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x49, 0x7a, 0x66,
	0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xc4, 0x75, 0xba, 0xf9, 0x69, 0x69, 0x99,
	0xc9, 0x99, 0x89, 0x39, 0x50, 0xbe, 0x3e, 0x8a, 0xdf, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8,
	0xc0, 0xfe, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x82, 0x1f, 0x09, 0xfd, 0x00, 0x00,
	0x00,
}

func (m *GoldPriceCallData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoldPriceCallData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoldPriceCallData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Multiplier != 0 {
		i = encodeVarintGoldPrice(dAtA, i, uint64(m.Multiplier))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GoldPriceResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoldPriceResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoldPriceResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Price != 0 {
		i = encodeVarintGoldPrice(dAtA, i, uint64(m.Price))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGoldPrice(dAtA []byte, offset int, v uint64) int {
	offset -= sovGoldPrice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GoldPriceCallData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Multiplier != 0 {
		n += 1 + sovGoldPrice(uint64(m.Multiplier))
	}
	return n
}

func (m *GoldPriceResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Price != 0 {
		n += 1 + sovGoldPrice(uint64(m.Price))
	}
	return n
}

func sovGoldPrice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGoldPrice(x uint64) (n int) {
	return sovGoldPrice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GoldPriceCallData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGoldPrice
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
			return fmt.Errorf("proto: GoldPriceCallData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoldPriceCallData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multiplier", wireType)
			}
			m.Multiplier = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoldPrice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Multiplier |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGoldPrice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGoldPrice
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
func (m *GoldPriceResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGoldPrice
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
			return fmt.Errorf("proto: GoldPriceResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoldPriceResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			m.Price = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoldPrice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Price |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGoldPrice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGoldPrice
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
func skipGoldPrice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGoldPrice
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
					return 0, ErrIntOverflowGoldPrice
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
					return 0, ErrIntOverflowGoldPrice
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
				return 0, ErrInvalidLengthGoldPrice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGoldPrice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGoldPrice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGoldPrice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGoldPrice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGoldPrice = fmt.Errorf("proto: unexpected end of group")
)
