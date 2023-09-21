// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gossip.proto

package gossip

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	cache "github.com/matrixorigin/matrixone/pkg/pb/cache"
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

type Operation int32

const (
	Operation_Unknown Operation = 0
	Operation_Set     Operation = 1
	Operation_Delete  Operation = 2
)

var Operation_name = map[int32]string{
	0: "Unknown",
	1: "Set",
	2: "Delete",
}

var Operation_value = map[string]int32{
	"Unknown": 0,
	"Set":     1,
	"Delete":  2,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}

func (Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_878fa4887b90140c, []int{0}
}

type CacheKeyItem struct {
	Operation            Operation      `protobuf:"varint,1,opt,name=Operation,proto3,enum=gossip.Operation" json:"Operation,omitempty"`
	CacheKey             cache.CacheKey `protobuf:"bytes,2,opt,name=CacheKey,proto3" json:"CacheKey"`
	TargetAddress        string         `protobuf:"bytes,3,opt,name=TargetAddress,proto3" json:"TargetAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CacheKeyItem) Reset()         { *m = CacheKeyItem{} }
func (m *CacheKeyItem) String() string { return proto.CompactTextString(m) }
func (*CacheKeyItem) ProtoMessage()    {}
func (*CacheKeyItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_878fa4887b90140c, []int{0}
}
func (m *CacheKeyItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CacheKeyItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CacheKeyItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CacheKeyItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheKeyItem.Merge(m, src)
}
func (m *CacheKeyItem) XXX_Size() int {
	return m.Size()
}
func (m *CacheKeyItem) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheKeyItem.DiscardUnknown(m)
}

var xxx_messageInfo_CacheKeyItem proto.InternalMessageInfo

func (m *CacheKeyItem) GetOperation() Operation {
	if m != nil {
		return m.Operation
	}
	return Operation_Unknown
}

func (m *CacheKeyItem) GetCacheKey() cache.CacheKey {
	if m != nil {
		return m.CacheKey
	}
	return cache.CacheKey{}
}

func (m *CacheKeyItem) GetTargetAddress() string {
	if m != nil {
		return m.TargetAddress
	}
	return ""
}

func init() {
	proto.RegisterEnum("gossip.Operation", Operation_name, Operation_value)
	proto.RegisterType((*CacheKeyItem)(nil), "gossip.CacheKeyItem")
}

func init() { proto.RegisterFile("gossip.proto", fileDescriptor_878fa4887b90140c) }

var fileDescriptor_878fa4887b90140c = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcf, 0x2f, 0x2e,
	0xce, 0x2c, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0xa4, 0x74, 0xd3, 0x33,
	0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0xd2,
	0x49, 0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0xb4, 0x49, 0x71, 0x27, 0x27, 0x26, 0x67,
	0xa4, 0x42, 0x38, 0x4a, 0x33, 0x18, 0xb9, 0x78, 0x9c, 0x41, 0x7c, 0xef, 0xd4, 0x4a, 0xcf, 0x92,
	0xd4, 0x5c, 0x21, 0x7d, 0x2e, 0x4e, 0xff, 0x82, 0xd4, 0xa2, 0xc4, 0x92, 0xcc, 0xfc, 0x3c, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x41, 0x3d, 0xa8, 0xb5, 0x70, 0x89, 0x20, 0x84, 0x1a, 0x21,
	0x43, 0x2e, 0x0e, 0x98, 0x01, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xfc, 0x7a, 0x10, 0x1b,
	0x60, 0xc2, 0x4e, 0x2c, 0x27, 0xee, 0xc9, 0x33, 0x04, 0xc1, 0x95, 0x09, 0xa9, 0x70, 0xf1, 0x86,
	0x24, 0x16, 0xa5, 0xa7, 0x96, 0x38, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b, 0x30, 0x2b, 0x30,
	0x6a, 0x70, 0x06, 0xa1, 0x0a, 0x6a, 0xe9, 0x22, 0xb9, 0x44, 0x88, 0x9b, 0x8b, 0x3d, 0x34, 0x2f,
	0x3b, 0x2f, 0xbf, 0x3c, 0x4f, 0x80, 0x41, 0x88, 0x9d, 0x8b, 0x39, 0x38, 0xb5, 0x44, 0x80, 0x51,
	0x88, 0x8b, 0x8b, 0xcd, 0x25, 0x35, 0x27, 0xb5, 0x24, 0x55, 0x80, 0xc9, 0xc9, 0xf6, 0xc4, 0x23,
	0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x8c, 0xd2, 0x47, 0x0a, 0x93, 0xdc,
	0xc4, 0x92, 0xa2, 0xcc, 0x8a, 0xfc, 0xa2, 0xcc, 0xf4, 0xcc, 0x3c, 0x18, 0x27, 0x2f, 0x55, 0xbf,
	0x20, 0x3b, 0x5d, 0xbf, 0x20, 0x49, 0x1f, 0xe2, 0xb7, 0x24, 0x36, 0x70, 0x78, 0x18, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xd2, 0xbf, 0x3f, 0xcf, 0x63, 0x01, 0x00, 0x00,
}

func (m *CacheKeyItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CacheKeyItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CacheKeyItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.TargetAddress) > 0 {
		i -= len(m.TargetAddress)
		copy(dAtA[i:], m.TargetAddress)
		i = encodeVarintGossip(dAtA, i, uint64(len(m.TargetAddress)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.CacheKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGossip(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Operation != 0 {
		i = encodeVarintGossip(dAtA, i, uint64(m.Operation))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGossip(dAtA []byte, offset int, v uint64) int {
	offset -= sovGossip(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CacheKeyItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Operation != 0 {
		n += 1 + sovGossip(uint64(m.Operation))
	}
	l = m.CacheKey.Size()
	n += 1 + l + sovGossip(uint64(l))
	l = len(m.TargetAddress)
	if l > 0 {
		n += 1 + l + sovGossip(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGossip(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGossip(x uint64) (n int) {
	return sovGossip(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CacheKeyItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGossip
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
			return fmt.Errorf("proto: CacheKeyItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CacheKeyItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operation", wireType)
			}
			m.Operation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGossip
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Operation |= Operation(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacheKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGossip
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
				return ErrInvalidLengthGossip
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGossip
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CacheKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGossip
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
				return ErrInvalidLengthGossip
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGossip
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGossip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGossip
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
func skipGossip(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGossip
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
					return 0, ErrIntOverflowGossip
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
					return 0, ErrIntOverflowGossip
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
				return 0, ErrInvalidLengthGossip
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGossip
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGossip
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGossip        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGossip          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGossip = fmt.Errorf("proto: unexpected end of group")
)