// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mars/incentives/v1beta1/proposals.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// CreateIncentivesScheduleProposal defines a governance proposal for starting a new incentives schedule
type CreateIncentivesScheduleProposal struct {
	// Title is the title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Description is the description of the proposal
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// StartTime is the UNIX timestamp of which this incentives schedule shall begin
	StartTime time.Time `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	// EndTime is the UNIX timestamp of which this incentives schedule shall finish
	EndTime time.Time `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time" yaml:"end_time"`
	// Amount is the total amount of coins that shall be released to stakers throughout the span of this incentives schedule
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *CreateIncentivesScheduleProposal) Reset()      { *m = CreateIncentivesScheduleProposal{} }
func (*CreateIncentivesScheduleProposal) ProtoMessage() {}
func (*CreateIncentivesScheduleProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_4413b31d9752c467, []int{0}
}
func (m *CreateIncentivesScheduleProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateIncentivesScheduleProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateIncentivesScheduleProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateIncentivesScheduleProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateIncentivesScheduleProposal.Merge(m, src)
}
func (m *CreateIncentivesScheduleProposal) XXX_Size() int {
	return m.Size()
}
func (m *CreateIncentivesScheduleProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateIncentivesScheduleProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CreateIncentivesScheduleProposal proto.InternalMessageInfo

// TerminateIncentivesSchedulesProposal defines a governance proposal for pre-mature ending of one or more incentives schedules
type TerminateIncentivesSchedulesProposal struct {
	// Title is the title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Description is the description of the proposal
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Ids is the array of identifiers of the incentives schedules which are to be terminated
	Ids []uint64 `protobuf:"varint,3,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (m *TerminateIncentivesSchedulesProposal) Reset()      { *m = TerminateIncentivesSchedulesProposal{} }
func (*TerminateIncentivesSchedulesProposal) ProtoMessage() {}
func (*TerminateIncentivesSchedulesProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_4413b31d9752c467, []int{1}
}
func (m *TerminateIncentivesSchedulesProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TerminateIncentivesSchedulesProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TerminateIncentivesSchedulesProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TerminateIncentivesSchedulesProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminateIncentivesSchedulesProposal.Merge(m, src)
}
func (m *TerminateIncentivesSchedulesProposal) XXX_Size() int {
	return m.Size()
}
func (m *TerminateIncentivesSchedulesProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminateIncentivesSchedulesProposal.DiscardUnknown(m)
}

var xxx_messageInfo_TerminateIncentivesSchedulesProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateIncentivesScheduleProposal)(nil), "mars.incentives.v1beta1.CreateIncentivesScheduleProposal")
	proto.RegisterType((*TerminateIncentivesSchedulesProposal)(nil), "mars.incentives.v1beta1.TerminateIncentivesSchedulesProposal")
}

func init() {
	proto.RegisterFile("mars/incentives/v1beta1/proposals.proto", fileDescriptor_4413b31d9752c467)
}

var fileDescriptor_4413b31d9752c467 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xb1, 0x6f, 0xd4, 0x30,
	0x14, 0xc6, 0x93, 0xa6, 0x2d, 0xd4, 0x87, 0x04, 0x44, 0x95, 0x08, 0x87, 0x48, 0xa2, 0x13, 0x12,
	0x59, 0x6a, 0xd3, 0xb2, 0x75, 0xbc, 0x4e, 0xdd, 0x50, 0xe8, 0x80, 0x58, 0x90, 0x93, 0x3c, 0x72,
	0x16, 0x89, 0x5f, 0x14, 0x3b, 0x15, 0x5d, 0x98, 0x19, 0x3b, 0x32, 0xde, 0xcc, 0x5f, 0xd2, 0xb1,
	0x23, 0x53, 0x8b, 0xee, 0x16, 0x66, 0x56, 0x16, 0x14, 0x3b, 0xd7, 0x1e, 0x12, 0x12, 0x03, 0x53,
	0x9e, 0x5f, 0xbe, 0xef, 0xf7, 0x9e, 0xdf, 0x33, 0x79, 0x5e, 0xf3, 0x56, 0x31, 0x21, 0x73, 0x90,
	0x5a, 0x9c, 0x82, 0x62, 0xa7, 0xfb, 0x19, 0x68, 0xbe, 0xcf, 0x9a, 0x16, 0x1b, 0x54, 0xbc, 0x52,
	0xb4, 0x69, 0x51, 0xa3, 0xff, 0xa8, 0x17, 0xd2, 0x5b, 0x21, 0x1d, 0x84, 0xe3, 0x30, 0x47, 0x55,
	0xa3, 0x62, 0x19, 0x57, 0x70, 0xe3, 0xce, 0x51, 0x48, 0x6b, 0x1c, 0xef, 0x96, 0x58, 0xa2, 0x09,
	0x59, 0x1f, 0x0d, 0xd9, 0xa8, 0x44, 0x2c, 0x2b, 0x60, 0xe6, 0x94, 0x75, 0xef, 0x99, 0x16, 0x35,
	0x28, 0xcd, 0xeb, 0xc6, 0x0a, 0x26, 0xbf, 0x36, 0x48, 0x7c, 0xd4, 0x02, 0xd7, 0x70, 0x7c, 0x53,
	0xf3, 0x75, 0x3e, 0x83, 0xa2, 0xab, 0xe0, 0xd5, 0xd0, 0x9b, 0xbf, 0x4b, 0xb6, 0xb4, 0xd0, 0x15,
	0x04, 0x6e, 0xec, 0x26, 0x3b, 0xa9, 0x3d, 0xf8, 0x31, 0x19, 0x15, 0xa0, 0xf2, 0x56, 0x34, 0x5a,
	0xa0, 0x0c, 0x36, 0xcc, 0xbf, 0xf5, 0x94, 0xff, 0x86, 0x10, 0xa5, 0x79, 0xab, 0xdf, 0xf5, 0x55,
	0x03, 0x2f, 0x76, 0x93, 0xd1, 0xc1, 0x98, 0xda, 0x96, 0xe8, 0xaa, 0x25, 0x7a, 0xb2, 0x6a, 0x69,
	0xfa, 0xf4, 0xe2, 0x2a, 0x72, 0x7e, 0x5e, 0x45, 0x0f, 0xcf, 0x78, 0x5d, 0x1d, 0x4e, 0x6e, 0xbd,
	0x93, 0xf3, 0xeb, 0xc8, 0x4d, 0x77, 0x4c, 0xa2, 0x97, 0xfb, 0x29, 0xb9, 0x0b, 0xb2, 0xb0, 0xdc,
	0xcd, 0x7f, 0x72, 0x9f, 0x0c, 0xdc, 0xfb, 0x96, 0xbb, 0x72, 0x5a, 0xea, 0x1d, 0x90, 0x85, 0x61,
	0xe6, 0x64, 0x9b, 0xd7, 0xd8, 0x49, 0x1d, 0x6c, 0xc5, 0x5e, 0x32, 0x3a, 0x78, 0x4c, 0xed, 0xc8,
	0x69, 0x3f, 0xf2, 0xd5, 0x1e, 0xe8, 0x11, 0x0a, 0x39, 0x7d, 0xd1, 0x03, 0xbf, 0x5e, 0x47, 0x49,
	0x29, 0xf4, 0xac, 0xcb, 0x68, 0x8e, 0x35, 0x1b, 0xf6, 0x63, 0x3f, 0x7b, 0xaa, 0xf8, 0xc0, 0xf4,
	0x59, 0x03, 0xca, 0x18, 0x54, 0x3a, 0xa0, 0x0f, 0xef, 0x7d, 0x9e, 0x47, 0xce, 0x97, 0x79, 0xe4,
	0xfc, 0x98, 0x47, 0xce, 0xe4, 0x13, 0x79, 0x76, 0x02, 0x6d, 0x2d, 0xe4, 0x5f, 0xe7, 0xaf, 0xfe,
	0x7b, 0x01, 0x0f, 0x88, 0x27, 0x0a, 0x15, 0x78, 0xb1, 0x97, 0x6c, 0xa6, 0x7d, 0xf8, 0x67, 0xfd,
	0xe9, 0xf1, 0xc5, 0x22, 0x74, 0x2f, 0x17, 0xa1, 0xfb, 0x7d, 0x11, 0xba, 0xe7, 0xcb, 0xd0, 0xb9,
	0x5c, 0x86, 0xce, 0xb7, 0x65, 0xe8, 0xbc, 0x65, 0x6b, 0x37, 0xeb, 0x9f, 0xe4, 0x9e, 0x19, 0x6b,
	0x8e, 0x15, 0x9b, 0x75, 0x19, 0xfb, 0xb8, 0xfe, 0x94, 0xcd, 0x35, 0xb3, 0x6d, 0x23, 0x78, 0xf9,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x2c, 0x14, 0xd8, 0xea, 0x02, 0x00, 0x00,
}

func (m *CreateIncentivesScheduleProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateIncentivesScheduleProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateIncentivesScheduleProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposals(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintProposals(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintProposals(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TerminateIncentivesSchedulesProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TerminateIncentivesSchedulesProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TerminateIncentivesSchedulesProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ids) > 0 {
		dAtA4 := make([]byte, len(m.Ids)*10)
		var j3 int
		for _, num := range m.Ids {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintProposals(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposals(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposals(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreateIncentivesScheduleProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovProposals(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovProposals(uint64(l))
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovProposals(uint64(l))
		}
	}
	return n
}

func (m *TerminateIncentivesSchedulesProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	if len(m.Ids) > 0 {
		l = 0
		for _, e := range m.Ids {
			l += sovProposals(uint64(e))
		}
		n += 1 + sovProposals(uint64(l)) + l
	}
	return n
}

func sovProposals(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposals(x uint64) (n int) {
	return sovProposals(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreateIncentivesScheduleProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposals
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
			return fmt.Errorf("proto: CreateIncentivesScheduleProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateIncentivesScheduleProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
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
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposals(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposals
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
func (m *TerminateIncentivesSchedulesProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposals
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
			return fmt.Errorf("proto: TerminateIncentivesSchedulesProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TerminateIncentivesSchedulesProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
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
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProposals
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Ids = append(m.Ids, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProposals
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthProposals
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthProposals
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
				if elementCount != 0 && len(m.Ids) == 0 {
					m.Ids = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProposals
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Ids = append(m.Ids, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Ids", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProposals(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposals
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
func skipProposals(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposals
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
					return 0, ErrIntOverflowProposals
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
					return 0, ErrIntOverflowProposals
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
				return 0, ErrInvalidLengthProposals
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposals
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposals
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposals        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposals          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposals = fmt.Errorf("proto: unexpected end of group")
)
