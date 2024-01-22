// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethermint/bridge/v1/withdraw.proto

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

// Withdraw represents a record of withdraw operation.
type Withdraw struct {
	// tx_hash is the hash of the transaction.
	TxHash string `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	// from of withdraw tx.
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	// to of withdraw tx.
	To string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	// coin_type of withdraw tx.
	CoinType string `protobuf:"bytes,4,opt,name=coin_type,json=coinType,proto3" json:"coin_type,omitempty"`
	// value of withdraw tx.
	Value uint64 `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`
	// data of withdraw tx.
	Data string `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	// status of the withdraw record.
	// status can be one of the following: pending,signed,completed,failed.
	Status string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	// signatures of the withdraw record.
	Signatures map[string]string `protobuf:"bytes,8,rep,name=signatures,proto3" json:"signatures,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// creator is the sender of message.
	Creator string `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Withdraw) Reset()         { *m = Withdraw{} }
func (m *Withdraw) String() string { return proto.CompactTextString(m) }
func (*Withdraw) ProtoMessage()    {}
func (*Withdraw) Descriptor() ([]byte, []int) {
	return fileDescriptor_04502cae42d2afa0, []int{0}
}
func (m *Withdraw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Withdraw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Withdraw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Withdraw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Withdraw.Merge(m, src)
}
func (m *Withdraw) XXX_Size() int {
	return m.Size()
}
func (m *Withdraw) XXX_DiscardUnknown() {
	xxx_messageInfo_Withdraw.DiscardUnknown(m)
}

var xxx_messageInfo_Withdraw proto.InternalMessageInfo

func (m *Withdraw) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *Withdraw) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Withdraw) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Withdraw) GetCoinType() string {
	if m != nil {
		return m.CoinType
	}
	return ""
}

func (m *Withdraw) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Withdraw) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Withdraw) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Withdraw) GetSignatures() map[string]string {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *Withdraw) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Withdraw)(nil), "ethermint.bridge.v1.Withdraw")
	proto.RegisterMapType((map[string]string)(nil), "ethermint.bridge.v1.Withdraw.SignaturesEntry")
}

func init() {
	proto.RegisterFile("ethermint/bridge/v1/withdraw.proto", fileDescriptor_04502cae42d2afa0)
}

var fileDescriptor_04502cae42d2afa0 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xcd, 0x4a, 0xc3, 0x40,
	0x18, 0xec, 0xa6, 0x6d, 0xda, 0x7e, 0x82, 0xca, 0x2a, 0xba, 0x28, 0x84, 0xd2, 0x53, 0x3d, 0xb8,
	0xa1, 0x7a, 0x11, 0xc1, 0x8b, 0x22, 0x78, 0xf1, 0x52, 0x05, 0xc1, 0x4b, 0xd9, 0xb6, 0x6b, 0x13,
	0x34, 0xd9, 0xb2, 0xfb, 0x25, 0x6d, 0xde, 0xc2, 0xc7, 0xf2, 0x24, 0x3d, 0x7a, 0x94, 0xf6, 0x45,
	0x24, 0xd9, 0xfe, 0x88, 0x78, 0xfb, 0x66, 0x98, 0x4c, 0x76, 0x66, 0xa0, 0x25, 0x31, 0x90, 0x3a,
	0x0a, 0x63, 0xf4, 0xfb, 0x3a, 0x1c, 0x8e, 0xa4, 0x9f, 0x76, 0xfc, 0x49, 0x88, 0xc1, 0x50, 0x8b,
	0x09, 0x1f, 0x6b, 0x85, 0x8a, 0xee, 0xad, 0x35, 0xdc, 0x6a, 0x78, 0xda, 0x69, 0x7d, 0x3a, 0x50,
	0x7f, 0x5a, 0xea, 0xe8, 0x21, 0xd4, 0x70, 0xda, 0x0b, 0x84, 0x09, 0x18, 0x69, 0x92, 0x76, 0xa3,
	0xeb, 0xe2, 0xf4, 0x4e, 0x98, 0x80, 0x52, 0xa8, 0xbc, 0x68, 0x15, 0x31, 0xa7, 0x60, 0x8b, 0x9b,
	0x6e, 0x83, 0x83, 0x8a, 0x95, 0x0b, 0xc6, 0x41, 0x45, 0x8f, 0xa1, 0x31, 0x50, 0x61, 0xdc, 0xc3,
	0x6c, 0x2c, 0x59, 0xa5, 0xa0, 0xeb, 0x39, 0xf1, 0x98, 0x8d, 0x25, 0xdd, 0x87, 0x6a, 0x2a, 0xde,
	0x12, 0xc9, 0xaa, 0x4d, 0xd2, 0xae, 0x74, 0x2d, 0xc8, 0x6d, 0x87, 0x02, 0x05, 0x73, 0xad, 0x6d,
	0x7e, 0xd3, 0x03, 0x70, 0x0d, 0x0a, 0x4c, 0x0c, 0xab, 0xd9, 0x27, 0x58, 0x44, 0xef, 0x01, 0x4c,
	0x38, 0x8a, 0x05, 0x26, 0x5a, 0x1a, 0x56, 0x6f, 0x96, 0xdb, 0x5b, 0x67, 0xa7, 0xfc, 0x9f, 0x48,
	0x7c, 0x15, 0x87, 0x3f, 0xac, 0xf5, 0xb7, 0x31, 0xea, 0xac, 0xfb, 0xcb, 0x80, 0x32, 0xa8, 0x0d,
	0xb4, 0x14, 0xa8, 0x34, 0x6b, 0x14, 0xff, 0x59, 0xc1, 0xa3, 0x2b, 0xd8, 0xf9, 0xf3, 0x21, 0xdd,
	0x85, 0xf2, 0xab, 0xcc, 0x96, 0x9d, 0xe4, 0xe7, 0x26, 0x8f, 0x6d, 0xc4, 0x82, 0x4b, 0xe7, 0x82,
	0x5c, 0xdf, 0x7c, 0xcc, 0x3d, 0x32, 0x9b, 0x7b, 0xe4, 0x7b, 0xee, 0x91, 0xf7, 0x85, 0x57, 0x9a,
	0x2d, 0xbc, 0xd2, 0xd7, 0xc2, 0x2b, 0x3d, 0x9f, 0x8c, 0x42, 0x0c, 0x92, 0x3e, 0x1f, 0xa8, 0xc8,
	0x97, 0x69, 0xa4, 0x8c, 0xbf, 0x19, 0x6d, 0xba, 0x9a, 0x2d, 0x6f, 0xcf, 0xf4, 0xdd, 0x62, 0xb1,
	0xf3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xcf, 0x17, 0xed, 0xd7, 0x01, 0x00, 0x00,
}

func (m *Withdraw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Withdraw) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Withdraw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintWithdraw(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Signatures) > 0 {
		for k := range m.Signatures {
			v := m.Signatures[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintWithdraw(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintWithdraw(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintWithdraw(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintWithdraw(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintWithdraw(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x32
	}
	if m.Value != 0 {
		i = encodeVarintWithdraw(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x28
	}
	if len(m.CoinType) > 0 {
		i -= len(m.CoinType)
		copy(dAtA[i:], m.CoinType)
		i = encodeVarintWithdraw(dAtA, i, uint64(len(m.CoinType)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintWithdraw(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintWithdraw(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintWithdraw(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWithdraw(dAtA []byte, offset int, v uint64) int {
	offset -= sovWithdraw(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Withdraw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovWithdraw(uint64(l))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovWithdraw(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovWithdraw(uint64(l))
	}
	l = len(m.CoinType)
	if l > 0 {
		n += 1 + l + sovWithdraw(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovWithdraw(uint64(m.Value))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovWithdraw(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovWithdraw(uint64(l))
	}
	if len(m.Signatures) > 0 {
		for k, v := range m.Signatures {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovWithdraw(uint64(len(k))) + 1 + len(v) + sovWithdraw(uint64(len(v)))
			n += mapEntrySize + 1 + sovWithdraw(uint64(mapEntrySize))
		}
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovWithdraw(uint64(l))
	}
	return n
}

func sovWithdraw(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWithdraw(x uint64) (n int) {
	return sovWithdraw(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Withdraw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWithdraw
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
			return fmt.Errorf("proto: Withdraw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Withdraw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoinType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Signatures == nil {
				m.Signatures = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowWithdraw
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWithdraw
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthWithdraw
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthWithdraw
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWithdraw
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthWithdraw
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthWithdraw
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipWithdraw(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthWithdraw
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Signatures[mapkey] = mapvalue
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWithdraw
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
				return ErrInvalidLengthWithdraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWithdraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWithdraw(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWithdraw
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
func skipWithdraw(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWithdraw
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
					return 0, ErrIntOverflowWithdraw
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
					return 0, ErrIntOverflowWithdraw
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
				return 0, ErrInvalidLengthWithdraw
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWithdraw
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWithdraw
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWithdraw        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWithdraw          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWithdraw = fmt.Errorf("proto: unexpected end of group")
)
