// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue503.proto

package issue503

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/crx666/protobuf/gogoproto"
	proto "github.com/crx666/protobuf/proto"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Foo struct {
	Num1                 []int64  `protobuf:"varint,1,rep,packed,name=num1,proto3" json:"num1,omitempty"`
	Num2                 []int32  `protobuf:"varint,2,rep,packed,name=num2,proto3" json:"num2,omitempty"`
	Str1                 []string `protobuf:"bytes,3,rep,name=str1,proto3" json:"str1,omitempty"`
	Dat1                 [][]byte `protobuf:"bytes,4,rep,name=dat1,proto3" json:"dat1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8886d01c7a4b3fa8, []int{0}
}
func (m *Foo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(m, src)
}
func (m *Foo) XXX_Size() int {
	return m.Size()
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

func (m *Foo) GetNum1() []int64 {
	if m != nil {
		return m.Num1
	}
	return nil
}

func (m *Foo) GetNum2() []int32 {
	if m != nil {
		return m.Num2
	}
	return nil
}

func (m *Foo) GetStr1() []string {
	if m != nil {
		return m.Str1
	}
	return nil
}

func (m *Foo) GetDat1() [][]byte {
	if m != nil {
		return m.Dat1
	}
	return nil
}

func init() {
	proto.RegisterType((*Foo)(nil), "issue503.Foo")
}

func init() { proto.RegisterFile("issue503.proto", fileDescriptor_8886d01c7a4b3fa8) }

var fileDescriptor_8886d01c7a4b3fa8 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x35, 0x30, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x74,
	0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5,
	0xc1, 0x0a, 0x92, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x68, 0x54, 0x0a, 0xe5, 0x62,
	0x76, 0xcb, 0xcf, 0x17, 0x12, 0xe2, 0x62, 0xc9, 0x2b, 0xcd, 0x35, 0x94, 0x60, 0x54, 0x60, 0xd6,
	0x60, 0x0e, 0x02, 0xb3, 0xa1, 0x62, 0x46, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0xac, 0x60, 0x31, 0x23,
	0x90, 0x58, 0x71, 0x49, 0x91, 0xa1, 0x04, 0xb3, 0x02, 0xb3, 0x06, 0x67, 0x10, 0x98, 0x0d, 0x12,
	0x4b, 0x49, 0x2c, 0x31, 0x94, 0x60, 0x51, 0x60, 0xd6, 0xe0, 0x09, 0x02, 0xb3, 0x9d, 0x24, 0x1e,
	0x3c, 0x94, 0x63, 0xfc, 0xf1, 0x50, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x13, 0x8f, 0xe4, 0x18,
	0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x31, 0x89, 0x0d, 0x6c, 0xaf, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xe2, 0xf0, 0x54, 0xb4, 0xc2, 0x00, 0x00, 0x00,
}

func (this *Foo) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Foo)
	if !ok {
		that2, ok := that.(Foo)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Foo")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Foo but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Foo but is not nil && this == nil")
	}
	if len(this.Num1) != len(that1.Num1) {
		return fmt.Errorf("Num1 this(%v) Not Equal that(%v)", len(this.Num1), len(that1.Num1))
	}
	for i := range this.Num1 {
		if this.Num1[i] != that1.Num1[i] {
			return fmt.Errorf("Num1 this[%v](%v) Not Equal that[%v](%v)", i, this.Num1[i], i, that1.Num1[i])
		}
	}
	if len(this.Num2) != len(that1.Num2) {
		return fmt.Errorf("Num2 this(%v) Not Equal that(%v)", len(this.Num2), len(that1.Num2))
	}
	for i := range this.Num2 {
		if this.Num2[i] != that1.Num2[i] {
			return fmt.Errorf("Num2 this[%v](%v) Not Equal that[%v](%v)", i, this.Num2[i], i, that1.Num2[i])
		}
	}
	if len(this.Str1) != len(that1.Str1) {
		return fmt.Errorf("Str1 this(%v) Not Equal that(%v)", len(this.Str1), len(that1.Str1))
	}
	for i := range this.Str1 {
		if this.Str1[i] != that1.Str1[i] {
			return fmt.Errorf("Str1 this[%v](%v) Not Equal that[%v](%v)", i, this.Str1[i], i, that1.Str1[i])
		}
	}
	if len(this.Dat1) != len(that1.Dat1) {
		return fmt.Errorf("Dat1 this(%v) Not Equal that(%v)", len(this.Dat1), len(that1.Dat1))
	}
	for i := range this.Dat1 {
		if !bytes.Equal(this.Dat1[i], that1.Dat1[i]) {
			return fmt.Errorf("Dat1 this[%v](%v) Not Equal that[%v](%v)", i, this.Dat1[i], i, that1.Dat1[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *Foo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Foo)
	if !ok {
		that2, ok := that.(Foo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Num1) != len(that1.Num1) {
		return false
	}
	for i := range this.Num1 {
		if this.Num1[i] != that1.Num1[i] {
			return false
		}
	}
	if len(this.Num2) != len(that1.Num2) {
		return false
	}
	for i := range this.Num2 {
		if this.Num2[i] != that1.Num2[i] {
			return false
		}
	}
	if len(this.Str1) != len(that1.Str1) {
		return false
	}
	for i := range this.Str1 {
		if this.Str1[i] != that1.Str1[i] {
			return false
		}
	}
	if len(this.Dat1) != len(that1.Dat1) {
		return false
	}
	for i := range this.Dat1 {
		if !bytes.Equal(this.Dat1[i], that1.Dat1[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *Foo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Foo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Num1) > 0 {
		dAtA2 := make([]byte, len(m.Num1)*10)
		var j1 int
		for _, num1 := range m.Num1 {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0xa
		i++
		i = encodeVarintIssue503(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if len(m.Num2) > 0 {
		dAtA4 := make([]byte, len(m.Num2)*10)
		var j3 int
		for _, num1 := range m.Num2 {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintIssue503(dAtA, i, uint64(j3))
		i += copy(dAtA[i:], dAtA4[:j3])
	}
	if len(m.Str1) > 0 {
		for _, s := range m.Str1 {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Dat1) > 0 {
		for _, b := range m.Dat1 {
			dAtA[i] = 0x22
			i++
			i = encodeVarintIssue503(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintIssue503(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedFoo(r randyIssue503, easy bool) *Foo {
	this := &Foo{}
	v1 := r.Intn(10)
	this.Num1 = make([]int64, v1)
	for i := 0; i < v1; i++ {
		this.Num1[i] = int64(r.Int63())
		if r.Intn(2) == 0 {
			this.Num1[i] *= -1
		}
	}
	v2 := r.Intn(10)
	this.Num2 = make([]int32, v2)
	for i := 0; i < v2; i++ {
		this.Num2[i] = int32(r.Int31())
		if r.Intn(2) == 0 {
			this.Num2[i] *= -1
		}
	}
	v3 := r.Intn(10)
	this.Str1 = make([]string, v3)
	for i := 0; i < v3; i++ {
		this.Str1[i] = string(randStringIssue503(r))
	}
	v4 := r.Intn(10)
	this.Dat1 = make([][]byte, v4)
	for i := 0; i < v4; i++ {
		v5 := r.Intn(100)
		this.Dat1[i] = make([]byte, v5)
		for j := 0; j < v5; j++ {
			this.Dat1[i][j] = byte(r.Intn(256))
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedIssue503(r, 5)
	}
	return this
}

type randyIssue503 interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneIssue503(r randyIssue503) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringIssue503(r randyIssue503) string {
	v6 := r.Intn(100)
	tmps := make([]rune, v6)
	for i := 0; i < v6; i++ {
		tmps[i] = randUTF8RuneIssue503(r)
	}
	return string(tmps)
}
func randUnrecognizedIssue503(r randyIssue503, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldIssue503(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldIssue503(dAtA []byte, r randyIssue503, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateIssue503(dAtA, uint64(key))
		v7 := r.Int63()
		if r.Intn(2) == 0 {
			v7 *= -1
		}
		dAtA = encodeVarintPopulateIssue503(dAtA, uint64(v7))
	case 1:
		dAtA = encodeVarintPopulateIssue503(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateIssue503(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateIssue503(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateIssue503(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateIssue503(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Foo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Num1) > 0 {
		l = 0
		for _, e := range m.Num1 {
			l += sovIssue503(uint64(e))
		}
		n += 1 + sovIssue503(uint64(l)) + l
	}
	if len(m.Num2) > 0 {
		l = 0
		for _, e := range m.Num2 {
			l += sovIssue503(uint64(e))
		}
		n += 1 + sovIssue503(uint64(l)) + l
	}
	if len(m.Str1) > 0 {
		for _, s := range m.Str1 {
			l = len(s)
			n += 1 + l + sovIssue503(uint64(l))
		}
	}
	if len(m.Dat1) > 0 {
		for _, b := range m.Dat1 {
			l = len(b)
			n += 1 + l + sovIssue503(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIssue503(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozIssue503(x uint64) (n int) {
	return sovIssue503(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Foo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIssue503
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Foo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Foo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowIssue503
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Num1 = append(m.Num1, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowIssue503
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthIssue503
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthIssue503
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Num1) == 0 {
					m.Num1 = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowIssue503
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Num1 = append(m.Num1, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Num1", wireType)
			}
		case 2:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowIssue503
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Num2 = append(m.Num2, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowIssue503
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthIssue503
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthIssue503
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Num2) == 0 {
					m.Num2 = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowIssue503
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Num2 = append(m.Num2, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Num2", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Str1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue503
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIssue503
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIssue503
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Str1 = append(m.Str1, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dat1", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue503
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIssue503
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIssue503
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dat1 = append(m.Dat1, make([]byte, postIndex-iNdEx))
			copy(m.Dat1[len(m.Dat1)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIssue503(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIssue503
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIssue503
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
func skipIssue503(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIssue503
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
					return 0, ErrIntOverflowIssue503
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowIssue503
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
				return 0, ErrInvalidLengthIssue503
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthIssue503
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIssue503
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipIssue503(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthIssue503
				}
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
	ErrInvalidLengthIssue503 = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIssue503   = fmt.Errorf("proto: integer overflow")
)
