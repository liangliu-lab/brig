// Code generated by protoc-gen-gogo.
// source: protocol.proto
// DO NOT EDIT!

/*
	Package proto is a generated protocol buffer package.

	It is generated from these files:
		protocol.proto

	It has these top-level messages:
		Request
		Response
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RequestType int32

const (
	RequestType_INVALID     RequestType = 0
	RequestType_QUIT        RequestType = 1
	RequestType_ERROR       RequestType = 2
	RequestType_CLONE       RequestType = 3
	RequestType_UPDATE_FILE RequestType = 4
)

var RequestType_name = map[int32]string{
	0: "INVALID",
	1: "QUIT",
	2: "ERROR",
	3: "CLONE",
	4: "UPDATE_FILE",
}
var RequestType_value = map[string]int32{
	"INVALID":     0,
	"QUIT":        1,
	"ERROR":       2,
	"CLONE":       3,
	"UPDATE_FILE": 4,
}

func (x RequestType) Enum() *RequestType {
	p := new(RequestType)
	*p = x
	return p
}
func (x RequestType) String() string {
	return proto1.EnumName(RequestType_name, int32(x))
}
func (x *RequestType) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(RequestType_value, data, "RequestType")
	if err != nil {
		return err
	}
	*x = RequestType(value)
	return nil
}

type Request struct {
	Type             *RequestType `protobuf:"varint,1,req,name=type,enum=transfer.protocol.RequestType" json:"type,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto1.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetType() RequestType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return RequestType_INVALID
}

type Response struct {
	Type             *RequestType `protobuf:"varint,1,req,name=type,enum=transfer.protocol.RequestType" json:"type,omitempty"`
	Data             []byte       `protobuf:"bytes,2,req,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto1.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetType() RequestType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return RequestType_INVALID
}

func (m *Response) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto1.RegisterType((*Request)(nil), "transfer.protocol.Request")
	proto1.RegisterType((*Response)(nil), "transfer.protocol.Response")
	proto1.RegisterEnum("transfer.protocol.RequestType", RequestType_name, RequestType_value)
}
func (m *Request) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Request) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0x8
		i++
		i = encodeVarintProtocol(data, i, uint64(*m.Type))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Response) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Response) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0x8
		i++
		i = encodeVarintProtocol(data, i, uint64(*m.Type))
	}
	if m.Data == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0x12
		i++
		i = encodeVarintProtocol(data, i, uint64(len(m.Data)))
		i += copy(data[i:], m.Data)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Protocol(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Protocol(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintProtocol(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Request) Size() (n int) {
	var l int
	_ = l
	if m.Type != nil {
		n += 1 + sovProtocol(uint64(*m.Type))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Response) Size() (n int) {
	var l int
	_ = l
	if m.Type != nil {
		n += 1 + sovProtocol(uint64(*m.Type))
	}
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProtocol(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProtocol(x uint64) (n int) {
	return sovProtocol(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var v RequestType
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (RequestType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Type = &v
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var v RequestType
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (RequestType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Type = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProtocol(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthProtocol
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProtocol
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipProtocol(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthProtocol = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtocol   = fmt.Errorf("proto: integer overflow")
)
