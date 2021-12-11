// Code generated by protoc-gen-go. DO NOT EDIT.
// source: DataFormats.proto

package bk_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LedgerMetadataFormat_State int32

const (
	LedgerMetadataFormat_OPEN        LedgerMetadataFormat_State = 1
	LedgerMetadataFormat_IN_RECOVERY LedgerMetadataFormat_State = 2
	LedgerMetadataFormat_CLOSED      LedgerMetadataFormat_State = 3
)

var LedgerMetadataFormat_State_name = map[int32]string{
	1: "OPEN",
	2: "IN_RECOVERY",
	3: "CLOSED",
}

var LedgerMetadataFormat_State_value = map[string]int32{
	"OPEN":        1,
	"IN_RECOVERY": 2,
	"CLOSED":      3,
}

func (x LedgerMetadataFormat_State) Enum() *LedgerMetadataFormat_State {
	p := new(LedgerMetadataFormat_State)
	*p = x
	return p
}

func (x LedgerMetadataFormat_State) String() string {
	return proto.EnumName(LedgerMetadataFormat_State_name, int32(x))
}

func (x *LedgerMetadataFormat_State) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LedgerMetadataFormat_State_value, data, "LedgerMetadataFormat_State")
	if err != nil {
		return err
	}
	*x = LedgerMetadataFormat_State(value)
	return nil
}

func (LedgerMetadataFormat_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_38ce640af007ee02, []int{0, 0}
}

type LedgerMetadataFormat_DigestType int32

const (
	LedgerMetadataFormat_CRC32  LedgerMetadataFormat_DigestType = 1
	LedgerMetadataFormat_HMAC   LedgerMetadataFormat_DigestType = 2
	LedgerMetadataFormat_CRC32C LedgerMetadataFormat_DigestType = 3
	LedgerMetadataFormat_DUMMY  LedgerMetadataFormat_DigestType = 4
)

var LedgerMetadataFormat_DigestType_name = map[int32]string{
	1: "CRC32",
	2: "HMAC",
	3: "CRC32C",
	4: "DUMMY",
}

var LedgerMetadataFormat_DigestType_value = map[string]int32{
	"CRC32":  1,
	"HMAC":   2,
	"CRC32C": 3,
	"DUMMY":  4,
}

func (x LedgerMetadataFormat_DigestType) Enum() *LedgerMetadataFormat_DigestType {
	p := new(LedgerMetadataFormat_DigestType)
	*p = x
	return p
}

func (x LedgerMetadataFormat_DigestType) String() string {
	return proto.EnumName(LedgerMetadataFormat_DigestType_name, int32(x))
}

func (x *LedgerMetadataFormat_DigestType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LedgerMetadataFormat_DigestType_value, data, "LedgerMetadataFormat_DigestType")
	if err != nil {
		return err
	}
	*x = LedgerMetadataFormat_DigestType(value)
	return nil
}

func (LedgerMetadataFormat_DigestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_38ce640af007ee02, []int{0, 1}
}

//*
// Metadata format for storing ledger information
type LedgerMetadataFormat struct {
	QuorumSize           *int32                                   `protobuf:"varint,1,req,name=quorumSize" json:"quorumSize,omitempty"`
	EnsembleSize         *int32                                   `protobuf:"varint,2,req,name=ensembleSize" json:"ensembleSize,omitempty"`
	Length               *int64                                   `protobuf:"varint,3,req,name=length" json:"length,omitempty"`
	LastEntryId          *int64                                   `protobuf:"varint,4,opt,name=lastEntryId" json:"lastEntryId,omitempty"`
	State                *LedgerMetadataFormat_State              `protobuf:"varint,5,req,name=state,enum=bk.proto.LedgerMetadataFormat_State,def=1" json:"state,omitempty"`
	Segment              []*LedgerMetadataFormat_Segment          `protobuf:"bytes,6,rep,name=segment" json:"segment,omitempty"`
	DigestType           *LedgerMetadataFormat_DigestType         `protobuf:"varint,7,opt,name=digestType,enum=bk.proto.LedgerMetadataFormat_DigestType" json:"digestType,omitempty"`
	Password             []byte                                   `protobuf:"bytes,8,opt,name=password" json:"password,omitempty"`
	AckQuorumSize        *int32                                   `protobuf:"varint,9,opt,name=ackQuorumSize" json:"ackQuorumSize,omitempty"`
	Ctime                *int64                                   `protobuf:"varint,10,opt,name=ctime" json:"ctime,omitempty"`
	CustomMetadata       []*LedgerMetadataFormatCMetadataMapEntry `protobuf:"bytes,11,rep,name=customMetadata" json:"customMetadata,omitempty"`
	CToken               *int64                                   `protobuf:"varint,12,opt,name=cToken" json:"cToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *LedgerMetadataFormat) Reset()         { *m = LedgerMetadataFormat{} }
func (m *LedgerMetadataFormat) String() string { return proto.CompactTextString(m) }
func (*LedgerMetadataFormat) ProtoMessage()    {}
func (*LedgerMetadataFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_38ce640af007ee02, []int{0}
}

func (m *LedgerMetadataFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LedgerMetadataFormat.Unmarshal(m, b)
}
func (m *LedgerMetadataFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LedgerMetadataFormat.Marshal(b, m, deterministic)
}
func (m *LedgerMetadataFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LedgerMetadataFormat.Merge(m, src)
}
func (m *LedgerMetadataFormat) XXX_Size() int {
	return xxx_messageInfo_LedgerMetadataFormat.Size(m)
}
func (m *LedgerMetadataFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_LedgerMetadataFormat.DiscardUnknown(m)
}

var xxx_messageInfo_LedgerMetadataFormat proto.InternalMessageInfo

const Default_LedgerMetadataFormat_State LedgerMetadataFormat_State = LedgerMetadataFormat_OPEN

func (m *LedgerMetadataFormat) GetQuorumSize() int32 {
	if m != nil && m.QuorumSize != nil {
		return *m.QuorumSize
	}
	return 0
}

func (m *LedgerMetadataFormat) GetEnsembleSize() int32 {
	if m != nil && m.EnsembleSize != nil {
		return *m.EnsembleSize
	}
	return 0
}

func (m *LedgerMetadataFormat) GetLength() int64 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *LedgerMetadataFormat) GetLastEntryId() int64 {
	if m != nil && m.LastEntryId != nil {
		return *m.LastEntryId
	}
	return 0
}

func (m *LedgerMetadataFormat) GetState() LedgerMetadataFormat_State {
	if m != nil && m.State != nil {
		return *m.State
	}
	return Default_LedgerMetadataFormat_State
}

func (m *LedgerMetadataFormat) GetSegment() []*LedgerMetadataFormat_Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (m *LedgerMetadataFormat) GetDigestType() LedgerMetadataFormat_DigestType {
	if m != nil && m.DigestType != nil {
		return *m.DigestType
	}
	return LedgerMetadataFormat_CRC32
}

func (m *LedgerMetadataFormat) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *LedgerMetadataFormat) GetAckQuorumSize() int32 {
	if m != nil && m.AckQuorumSize != nil {
		return *m.AckQuorumSize
	}
	return 0
}

func (m *LedgerMetadataFormat) GetCtime() int64 {
	if m != nil && m.Ctime != nil {
		return *m.Ctime
	}
	return 0
}

func (m *LedgerMetadataFormat) GetCustomMetadata() []*LedgerMetadataFormatCMetadataMapEntry {
	if m != nil {
		return m.CustomMetadata
	}
	return nil
}

func (m *LedgerMetadataFormat) GetCToken() int64 {
	if m != nil && m.CToken != nil {
		return *m.CToken
	}
	return 0
}

type LedgerMetadataFormat_Segment struct {
	EnsembleMember       []string `protobuf:"bytes,1,rep,name=ensembleMember" json:"ensembleMember,omitempty"`
	FirstEntryId         *int64   `protobuf:"varint,2,req,name=firstEntryId" json:"firstEntryId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LedgerMetadataFormat_Segment) Reset()         { *m = LedgerMetadataFormat_Segment{} }
func (m *LedgerMetadataFormat_Segment) String() string { return proto.CompactTextString(m) }
func (*LedgerMetadataFormat_Segment) ProtoMessage()    {}
func (*LedgerMetadataFormat_Segment) Descriptor() ([]byte, []int) {
	return fileDescriptor_38ce640af007ee02, []int{0, 0}
}

func (m *LedgerMetadataFormat_Segment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LedgerMetadataFormat_Segment.Unmarshal(m, b)
}
func (m *LedgerMetadataFormat_Segment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LedgerMetadataFormat_Segment.Marshal(b, m, deterministic)
}
func (m *LedgerMetadataFormat_Segment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LedgerMetadataFormat_Segment.Merge(m, src)
}
func (m *LedgerMetadataFormat_Segment) XXX_Size() int {
	return xxx_messageInfo_LedgerMetadataFormat_Segment.Size(m)
}
func (m *LedgerMetadataFormat_Segment) XXX_DiscardUnknown() {
	xxx_messageInfo_LedgerMetadataFormat_Segment.DiscardUnknown(m)
}

var xxx_messageInfo_LedgerMetadataFormat_Segment proto.InternalMessageInfo

func (m *LedgerMetadataFormat_Segment) GetEnsembleMember() []string {
	if m != nil {
		return m.EnsembleMember
	}
	return nil
}

func (m *LedgerMetadataFormat_Segment) GetFirstEntryId() int64 {
	if m != nil && m.FirstEntryId != nil {
		return *m.FirstEntryId
	}
	return 0
}

type LedgerMetadataFormatCMetadataMapEntry struct {
	Key                  *string  `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LedgerMetadataFormatCMetadataMapEntry) Reset()         { *m = LedgerMetadataFormatCMetadataMapEntry{} }
func (m *LedgerMetadataFormatCMetadataMapEntry) String() string { return proto.CompactTextString(m) }
func (*LedgerMetadataFormatCMetadataMapEntry) ProtoMessage()    {}
func (*LedgerMetadataFormatCMetadataMapEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_38ce640af007ee02, []int{0, 1}
}

func (m *LedgerMetadataFormatCMetadataMapEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LedgerMetadataFormatCMetadataMapEntry.Unmarshal(m, b)
}
func (m *LedgerMetadataFormatCMetadataMapEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LedgerMetadataFormatCMetadataMapEntry.Marshal(b, m, deterministic)
}
func (m *LedgerMetadataFormatCMetadataMapEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LedgerMetadataFormatCMetadataMapEntry.Merge(m, src)
}
func (m *LedgerMetadataFormatCMetadataMapEntry) XXX_Size() int {
	return xxx_messageInfo_LedgerMetadataFormatCMetadataMapEntry.Size(m)
}
func (m *LedgerMetadataFormatCMetadataMapEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LedgerMetadataFormatCMetadataMapEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LedgerMetadataFormatCMetadataMapEntry proto.InternalMessageInfo

func (m *LedgerMetadataFormatCMetadataMapEntry) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *LedgerMetadataFormatCMetadataMapEntry) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type LedgerRereplicationLayoutFormat struct {
	Type                 *string  `protobuf:"bytes,1,req,name=type" json:"type,omitempty"`
	Version              *int32   `protobuf:"varint,2,req,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LedgerRereplicationLayoutFormat) Reset()         { *m = LedgerRereplicationLayoutFormat{} }
func (m *LedgerRereplicationLayoutFormat) String() string { return proto.CompactTextString(m) }
func (*LedgerRereplicationLayoutFormat) ProtoMessage()    {}
func (*LedgerRereplicationLayoutFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_38ce640af007ee02, []int{1}
}

func (m *LedgerRereplicationLayoutFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LedgerRereplicationLayoutFormat.Unmarshal(m, b)
}
func (m *LedgerRereplicationLayoutFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LedgerRereplicationLayoutFormat.Marshal(b, m, deterministic)
}
func (m *LedgerRereplicationLayoutFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LedgerRereplicationLayoutFormat.Merge(m, src)
}
func (m *LedgerRereplicationLayoutFormat) XXX_Size() int {
	return xxx_messageInfo_LedgerRereplicationLayoutFormat.Size(m)
}
func (m *LedgerRereplicationLayoutFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_LedgerRereplicationLayoutFormat.DiscardUnknown(m)
}

var xxx_messageInfo_LedgerRereplicationLayoutFormat proto.InternalMessageInfo

func (m *LedgerRereplicationLayoutFormat) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *LedgerRereplicationLayoutFormat) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

type UnderreplicatedLedgerFormat struct {
	Replica              []string `protobuf:"bytes,1,rep,name=replica" json:"replica,omitempty"`
	Ctime                *int64   `protobuf:"varint,2,opt,name=ctime" json:"ctime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnderreplicatedLedgerFormat) Reset()         { *m = UnderreplicatedLedgerFormat{} }
func (m *UnderreplicatedLedgerFormat) String() string { return proto.CompactTextString(m) }
func (*UnderreplicatedLedgerFormat) ProtoMessage()    {}
func (*UnderreplicatedLedgerFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_38ce640af007ee02, []int{2}
}

func (m *UnderreplicatedLedgerFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnderreplicatedLedgerFormat.Unmarshal(m, b)
}
func (m *UnderreplicatedLedgerFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnderreplicatedLedgerFormat.Marshal(b, m, deterministic)
}
func (m *UnderreplicatedLedgerFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnderreplicatedLedgerFormat.Merge(m, src)
}
func (m *UnderreplicatedLedgerFormat) XXX_Size() int {
	return xxx_messageInfo_UnderreplicatedLedgerFormat.Size(m)
}
func (m *UnderreplicatedLedgerFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_UnderreplicatedLedgerFormat.DiscardUnknown(m)
}

var xxx_messageInfo_UnderreplicatedLedgerFormat proto.InternalMessageInfo

func (m *UnderreplicatedLedgerFormat) GetReplica() []string {
	if m != nil {
		return m.Replica
	}
	return nil
}

func (m *UnderreplicatedLedgerFormat) GetCtime() int64 {
	if m != nil && m.Ctime != nil {
		return *m.Ctime
	}
	return 0
}

//*
// Cookie format for storing cookie information
type CookieFormat struct {
	BookieHost           *string  `protobuf:"bytes,1,req,name=bookieHost" json:"bookieHost,omitempty"`
	JournalDir           *string  `protobuf:"bytes,2,req,name=journalDir" json:"journalDir,omitempty"`
	LedgerDirs           *string  `protobuf:"bytes,3,req,name=ledgerDirs" json:"ledgerDirs,omitempty"`
	InstanceId           *string  `protobuf:"bytes,4,opt,name=instanceId" json:"instanceId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CookieFormat) Reset()         { *m = CookieFormat{} }
func (m *CookieFormat) String() string { return proto.CompactTextString(m) }
func (*CookieFormat) ProtoMessage()    {}
func (*CookieFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_38ce640af007ee02, []int{3}
}

func (m *CookieFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CookieFormat.Unmarshal(m, b)
}
func (m *CookieFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CookieFormat.Marshal(b, m, deterministic)
}
func (m *CookieFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CookieFormat.Merge(m, src)
}
func (m *CookieFormat) XXX_Size() int {
	return xxx_messageInfo_CookieFormat.Size(m)
}
func (m *CookieFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_CookieFormat.DiscardUnknown(m)
}

var xxx_messageInfo_CookieFormat proto.InternalMessageInfo

func (m *CookieFormat) GetBookieHost() string {
	if m != nil && m.BookieHost != nil {
		return *m.BookieHost
	}
	return ""
}

func (m *CookieFormat) GetJournalDir() string {
	if m != nil && m.JournalDir != nil {
		return *m.JournalDir
	}
	return ""
}

func (m *CookieFormat) GetLedgerDirs() string {
	if m != nil && m.LedgerDirs != nil {
		return *m.LedgerDirs
	}
	return ""
}

func (m *CookieFormat) GetInstanceId() string {
	if m != nil && m.InstanceId != nil {
		return *m.InstanceId
	}
	return ""
}

//*
// Debug information for locks
type LockDataFormat struct {
	BookieId             *string  `protobuf:"bytes,1,opt,name=bookieId" json:"bookieId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockDataFormat) Reset()         { *m = LockDataFormat{} }
func (m *LockDataFormat) String() string { return proto.CompactTextString(m) }
func (*LockDataFormat) ProtoMessage()    {}
func (*LockDataFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_38ce640af007ee02, []int{4}
}

func (m *LockDataFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockDataFormat.Unmarshal(m, b)
}
func (m *LockDataFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockDataFormat.Marshal(b, m, deterministic)
}
func (m *LockDataFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockDataFormat.Merge(m, src)
}
func (m *LockDataFormat) XXX_Size() int {
	return xxx_messageInfo_LockDataFormat.Size(m)
}
func (m *LockDataFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_LockDataFormat.DiscardUnknown(m)
}

var xxx_messageInfo_LockDataFormat proto.InternalMessageInfo

func (m *LockDataFormat) GetBookieId() string {
	if m != nil && m.BookieId != nil {
		return *m.BookieId
	}
	return ""
}

//*
// Debug information for auditor votes
type AuditorVoteFormat struct {
	BookieId             *string  `protobuf:"bytes,1,opt,name=bookieId" json:"bookieId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuditorVoteFormat) Reset()         { *m = AuditorVoteFormat{} }
func (m *AuditorVoteFormat) String() string { return proto.CompactTextString(m) }
func (*AuditorVoteFormat) ProtoMessage()    {}
func (*AuditorVoteFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_38ce640af007ee02, []int{5}
}

func (m *AuditorVoteFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditorVoteFormat.Unmarshal(m, b)
}
func (m *AuditorVoteFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditorVoteFormat.Marshal(b, m, deterministic)
}
func (m *AuditorVoteFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditorVoteFormat.Merge(m, src)
}
func (m *AuditorVoteFormat) XXX_Size() int {
	return xxx_messageInfo_AuditorVoteFormat.Size(m)
}
func (m *AuditorVoteFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditorVoteFormat.DiscardUnknown(m)
}

var xxx_messageInfo_AuditorVoteFormat proto.InternalMessageInfo

func (m *AuditorVoteFormat) GetBookieId() string {
	if m != nil && m.BookieId != nil {
		return *m.BookieId
	}
	return ""
}

//*
// information of checkAllLedgers execution
type CheckAllLedgersFormat struct {
	CheckAllLedgersCTime *int64   `protobuf:"varint,1,opt,name=checkAllLedgersCTime" json:"checkAllLedgersCTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckAllLedgersFormat) Reset()         { *m = CheckAllLedgersFormat{} }
func (m *CheckAllLedgersFormat) String() string { return proto.CompactTextString(m) }
func (*CheckAllLedgersFormat) ProtoMessage()    {}
func (*CheckAllLedgersFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_38ce640af007ee02, []int{6}
}

func (m *CheckAllLedgersFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAllLedgersFormat.Unmarshal(m, b)
}
func (m *CheckAllLedgersFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAllLedgersFormat.Marshal(b, m, deterministic)
}
func (m *CheckAllLedgersFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAllLedgersFormat.Merge(m, src)
}
func (m *CheckAllLedgersFormat) XXX_Size() int {
	return xxx_messageInfo_CheckAllLedgersFormat.Size(m)
}
func (m *CheckAllLedgersFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAllLedgersFormat.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAllLedgersFormat proto.InternalMessageInfo

func (m *CheckAllLedgersFormat) GetCheckAllLedgersCTime() int64 {
	if m != nil && m.CheckAllLedgersCTime != nil {
		return *m.CheckAllLedgersCTime
	}
	return 0
}

//*
// information of PlacementPolicyCheck execution
type PlacementPolicyCheckFormat struct {
	PlacementPolicyCheckCTime *int64   `protobuf:"varint,1,opt,name=placementPolicyCheckCTime" json:"placementPolicyCheckCTime,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *PlacementPolicyCheckFormat) Reset()         { *m = PlacementPolicyCheckFormat{} }
func (m *PlacementPolicyCheckFormat) String() string { return proto.CompactTextString(m) }
func (*PlacementPolicyCheckFormat) ProtoMessage()    {}
func (*PlacementPolicyCheckFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_38ce640af007ee02, []int{7}
}

func (m *PlacementPolicyCheckFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlacementPolicyCheckFormat.Unmarshal(m, b)
}
func (m *PlacementPolicyCheckFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlacementPolicyCheckFormat.Marshal(b, m, deterministic)
}
func (m *PlacementPolicyCheckFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlacementPolicyCheckFormat.Merge(m, src)
}
func (m *PlacementPolicyCheckFormat) XXX_Size() int {
	return xxx_messageInfo_PlacementPolicyCheckFormat.Size(m)
}
func (m *PlacementPolicyCheckFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_PlacementPolicyCheckFormat.DiscardUnknown(m)
}

var xxx_messageInfo_PlacementPolicyCheckFormat proto.InternalMessageInfo

func (m *PlacementPolicyCheckFormat) GetPlacementPolicyCheckCTime() int64 {
	if m != nil && m.PlacementPolicyCheckCTime != nil {
		return *m.PlacementPolicyCheckCTime
	}
	return 0
}

//*
// information of ReplicasCheck execution
type ReplicasCheckFormat struct {
	ReplicasCheckCTime   *int64   `protobuf:"varint,1,opt,name=replicasCheckCTime" json:"replicasCheckCTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplicasCheckFormat) Reset()         { *m = ReplicasCheckFormat{} }
func (m *ReplicasCheckFormat) String() string { return proto.CompactTextString(m) }
func (*ReplicasCheckFormat) ProtoMessage()    {}
func (*ReplicasCheckFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_38ce640af007ee02, []int{8}
}

func (m *ReplicasCheckFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplicasCheckFormat.Unmarshal(m, b)
}
func (m *ReplicasCheckFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplicasCheckFormat.Marshal(b, m, deterministic)
}
func (m *ReplicasCheckFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplicasCheckFormat.Merge(m, src)
}
func (m *ReplicasCheckFormat) XXX_Size() int {
	return xxx_messageInfo_ReplicasCheckFormat.Size(m)
}
func (m *ReplicasCheckFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplicasCheckFormat.DiscardUnknown(m)
}

var xxx_messageInfo_ReplicasCheckFormat proto.InternalMessageInfo

func (m *ReplicasCheckFormat) GetReplicasCheckCTime() int64 {
	if m != nil && m.ReplicasCheckCTime != nil {
		return *m.ReplicasCheckCTime
	}
	return 0
}

//*
// information about services exposed by a Bookie.
type BookieServiceInfoFormat struct {
	Endpoints            []*BookieServiceInfoFormat_Endpoint `protobuf:"bytes,6,rep,name=endpoints" json:"endpoints,omitempty"`
	Properties           map[string]string                   `protobuf:"bytes,7,rep,name=properties" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *BookieServiceInfoFormat) Reset()         { *m = BookieServiceInfoFormat{} }
func (m *BookieServiceInfoFormat) String() string { return proto.CompactTextString(m) }
func (*BookieServiceInfoFormat) ProtoMessage()    {}
func (*BookieServiceInfoFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_38ce640af007ee02, []int{9}
}

func (m *BookieServiceInfoFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookieServiceInfoFormat.Unmarshal(m, b)
}
func (m *BookieServiceInfoFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookieServiceInfoFormat.Marshal(b, m, deterministic)
}
func (m *BookieServiceInfoFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookieServiceInfoFormat.Merge(m, src)
}
func (m *BookieServiceInfoFormat) XXX_Size() int {
	return xxx_messageInfo_BookieServiceInfoFormat.Size(m)
}
func (m *BookieServiceInfoFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_BookieServiceInfoFormat.DiscardUnknown(m)
}

var xxx_messageInfo_BookieServiceInfoFormat proto.InternalMessageInfo

func (m *BookieServiceInfoFormat) GetEndpoints() []*BookieServiceInfoFormat_Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *BookieServiceInfoFormat) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

//*
// Information about an endpoint.
type BookieServiceInfoFormat_Endpoint struct {
	Id                   *string  `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Port                 *int32   `protobuf:"varint,2,req,name=port" json:"port,omitempty"`
	Host                 *string  `protobuf:"bytes,3,req,name=host" json:"host,omitempty"`
	Protocol             *string  `protobuf:"bytes,4,req,name=protocol" json:"protocol,omitempty"`
	Auth                 []string `protobuf:"bytes,5,rep,name=auth" json:"auth,omitempty"`
	Extensions           []string `protobuf:"bytes,6,rep,name=extensions" json:"extensions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookieServiceInfoFormat_Endpoint) Reset()         { *m = BookieServiceInfoFormat_Endpoint{} }
func (m *BookieServiceInfoFormat_Endpoint) String() string { return proto.CompactTextString(m) }
func (*BookieServiceInfoFormat_Endpoint) ProtoMessage()    {}
func (*BookieServiceInfoFormat_Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_38ce640af007ee02, []int{9, 0}
}

func (m *BookieServiceInfoFormat_Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookieServiceInfoFormat_Endpoint.Unmarshal(m, b)
}
func (m *BookieServiceInfoFormat_Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookieServiceInfoFormat_Endpoint.Marshal(b, m, deterministic)
}
func (m *BookieServiceInfoFormat_Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookieServiceInfoFormat_Endpoint.Merge(m, src)
}
func (m *BookieServiceInfoFormat_Endpoint) XXX_Size() int {
	return xxx_messageInfo_BookieServiceInfoFormat_Endpoint.Size(m)
}
func (m *BookieServiceInfoFormat_Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_BookieServiceInfoFormat_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_BookieServiceInfoFormat_Endpoint proto.InternalMessageInfo

func (m *BookieServiceInfoFormat_Endpoint) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *BookieServiceInfoFormat_Endpoint) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *BookieServiceInfoFormat_Endpoint) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *BookieServiceInfoFormat_Endpoint) GetProtocol() string {
	if m != nil && m.Protocol != nil {
		return *m.Protocol
	}
	return ""
}

func (m *BookieServiceInfoFormat_Endpoint) GetAuth() []string {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *BookieServiceInfoFormat_Endpoint) GetExtensions() []string {
	if m != nil {
		return m.Extensions
	}
	return nil
}

func init() {
	proto.RegisterEnum("bk.proto.LedgerMetadataFormat_State", LedgerMetadataFormat_State_name, LedgerMetadataFormat_State_value)
	proto.RegisterEnum("bk.proto.LedgerMetadataFormat_DigestType", LedgerMetadataFormat_DigestType_name, LedgerMetadataFormat_DigestType_value)
	proto.RegisterType((*LedgerMetadataFormat)(nil), "bk.proto.LedgerMetadataFormat")
	proto.RegisterType((*LedgerMetadataFormat_Segment)(nil), "bk.proto.LedgerMetadataFormat.Segment")
	proto.RegisterType((*LedgerMetadataFormatCMetadataMapEntry)(nil), "bk.proto.LedgerMetadataFormat.cMetadataMapEntry")
	proto.RegisterType((*LedgerRereplicationLayoutFormat)(nil), "bk.proto.LedgerRereplicationLayoutFormat")
	proto.RegisterType((*UnderreplicatedLedgerFormat)(nil), "bk.proto.UnderreplicatedLedgerFormat")
	proto.RegisterType((*CookieFormat)(nil), "bk.proto.CookieFormat")
	proto.RegisterType((*LockDataFormat)(nil), "bk.proto.LockDataFormat")
	proto.RegisterType((*AuditorVoteFormat)(nil), "bk.proto.AuditorVoteFormat")
	proto.RegisterType((*CheckAllLedgersFormat)(nil), "bk.proto.CheckAllLedgersFormat")
	proto.RegisterType((*PlacementPolicyCheckFormat)(nil), "bk.proto.PlacementPolicyCheckFormat")
	proto.RegisterType((*ReplicasCheckFormat)(nil), "bk.proto.ReplicasCheckFormat")
	proto.RegisterType((*BookieServiceInfoFormat)(nil), "bk.proto.BookieServiceInfoFormat")
	proto.RegisterMapType((map[string]string)(nil), "bk.proto.BookieServiceInfoFormat.PropertiesEntry")
	proto.RegisterType((*BookieServiceInfoFormat_Endpoint)(nil), "bk.proto.BookieServiceInfoFormat.Endpoint")
}

func init() { proto.RegisterFile("DataFormats.proto", fileDescriptor_38ce640af007ee02) }

var fileDescriptor_38ce640af007ee02 = []byte{
	// 882 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x51, 0x6f, 0xe3, 0x44,
	0x10, 0x96, 0x9d, 0xa4, 0x89, 0xa7, 0x25, 0x97, 0x2e, 0x05, 0x4c, 0x4e, 0xe2, 0x2c, 0xeb, 0x74,
	0x32, 0x08, 0x05, 0x08, 0x2f, 0xa7, 0x03, 0xa4, 0xeb, 0x25, 0x41, 0x8d, 0x68, 0xae, 0xbd, 0x6d,
	0x7b, 0xe2, 0x78, 0x41, 0x5b, 0x7b, 0x2e, 0x59, 0xe2, 0x78, 0xcd, 0x7a, 0x53, 0x08, 0x3f, 0x81,
	0x07, 0x1e, 0xf8, 0x1b, 0xfc, 0x49, 0xb4, 0xeb, 0x4d, 0xe2, 0x96, 0x9e, 0xc2, 0xdb, 0xce, 0x37,
	0x33, 0x9f, 0xc7, 0xdf, 0x7e, 0xb3, 0x70, 0x38, 0x64, 0x8a, 0x7d, 0x2f, 0xe4, 0x82, 0xa9, 0xa2,
	0x97, 0x4b, 0xa1, 0x04, 0x69, 0x5d, 0xcf, 0xcb, 0x53, 0xf8, 0xcf, 0x1e, 0x1c, 0x9d, 0x62, 0x32,
	0x45, 0x39, 0x41, 0xc5, 0x92, 0x4d, 0x25, 0xf9, 0x04, 0xe0, 0xd7, 0xa5, 0x90, 0xcb, 0xc5, 0x05,
	0xff, 0x03, 0x7d, 0x27, 0x70, 0xa3, 0x06, 0xad, 0x20, 0x24, 0x84, 0x03, 0xcc, 0x0a, 0x5c, 0x5c,
	0xa7, 0x68, 0x2a, 0x5c, 0x53, 0x71, 0x0b, 0x23, 0x1f, 0xc2, 0x5e, 0x8a, 0xd9, 0x54, 0xcd, 0xfc,
	0x5a, 0xe0, 0x46, 0x35, 0x6a, 0x23, 0x12, 0xc0, 0x7e, 0xca, 0x0a, 0x35, 0xca, 0x94, 0x5c, 0x8d,
	0x13, 0xbf, 0x1e, 0x38, 0x51, 0x8d, 0x56, 0x21, 0xf2, 0x1c, 0x1a, 0x85, 0x62, 0x0a, 0xfd, 0x46,
	0xe0, 0x46, 0xed, 0xfe, 0xe3, 0xde, 0x7a, 0xe0, 0xde, 0x7d, 0xc3, 0xf6, 0x2e, 0x74, 0xed, 0xb3,
	0xfa, 0xd9, 0xf9, 0xe8, 0x25, 0x2d, 0x1b, 0xc9, 0x73, 0x68, 0x16, 0x38, 0x5d, 0x60, 0xa6, 0xfc,
	0xbd, 0xa0, 0x16, 0xed, 0xf7, 0x9f, 0xec, 0xe2, 0x28, 0xab, 0xe9, 0xba, 0x8d, 0x8c, 0x01, 0x12,
	0x3e, 0xc5, 0x42, 0x5d, 0xae, 0x72, 0xf4, 0x9b, 0x81, 0x13, 0xb5, 0xfb, 0x9f, 0xee, 0x20, 0x19,
	0x6e, 0x1a, 0x68, 0xa5, 0x99, 0x74, 0xa1, 0x95, 0xb3, 0xa2, 0xf8, 0x4d, 0xc8, 0xc4, 0x6f, 0x05,
	0x4e, 0x74, 0x40, 0x37, 0x31, 0x79, 0x0c, 0xef, 0xb1, 0x78, 0xfe, 0x6a, 0xab, 0xb5, 0x17, 0x38,
	0x51, 0x83, 0xde, 0x06, 0xc9, 0x11, 0x34, 0x62, 0xc5, 0x17, 0xe8, 0x83, 0x11, 0xab, 0x0c, 0xc8,
	0x8f, 0xd0, 0x8e, 0x97, 0x85, 0x12, 0x8b, 0xf5, 0x18, 0xfe, 0xbe, 0xf9, 0xd7, 0x2f, 0x77, 0x8c,
	0x19, 0xaf, 0xe3, 0x09, 0xcb, 0x8d, 0xe4, 0xf4, 0x0e, 0x8f, 0xbe, 0xba, 0xf8, 0x52, 0xcc, 0x31,
	0xf3, 0x0f, 0xcc, 0x07, 0x6d, 0xd4, 0xbd, 0x82, 0xa6, 0x15, 0x8a, 0x3c, 0x81, 0xf6, 0xfa, 0xb6,
	0x27, 0xb8, 0xb8, 0x46, 0xe9, 0x3b, 0x41, 0x2d, 0xf2, 0xe8, 0x1d, 0x54, 0x3b, 0xe5, 0x2d, 0x97,
	0xdb, 0xeb, 0x76, 0x8d, 0x17, 0x6e, 0x61, 0xdd, 0x6f, 0xe0, 0xf0, 0x3f, 0x33, 0x91, 0x0e, 0xd4,
	0xe6, 0xb8, 0xf2, 0x9d, 0xc0, 0x89, 0x3c, 0xaa, 0x8f, 0x5a, 0x85, 0x1b, 0x96, 0x2e, 0xb5, 0xdb,
	0xb4, 0x88, 0x65, 0x10, 0xf6, 0xa0, 0x61, 0x0c, 0x40, 0x5a, 0x60, 0x2c, 0xd0, 0x71, 0xc8, 0x03,
	0xd8, 0x1f, 0xbf, 0xfc, 0x99, 0x8e, 0x06, 0x67, 0xaf, 0x47, 0xf4, 0x4d, 0xc7, 0x25, 0x00, 0x7b,
	0x83, 0xd3, 0xb3, 0x8b, 0xd1, 0xb0, 0x53, 0x0b, 0x9f, 0x02, 0x6c, 0xef, 0x89, 0x78, 0xd0, 0x18,
	0xd0, 0xc1, 0xd7, 0xfd, 0x8e, 0xa3, 0xfb, 0x4f, 0x26, 0xc7, 0x03, 0x5b, 0xae, 0xc1, 0x41, 0xa7,
	0xa6, 0x0b, 0x86, 0x57, 0x93, 0xc9, 0x9b, 0x4e, 0x3d, 0x3c, 0x83, 0x47, 0xa5, 0x9e, 0x14, 0x25,
	0xe6, 0x29, 0x8f, 0x99, 0xe2, 0x22, 0x3b, 0x65, 0x2b, 0xb1, 0x54, 0x76, 0x6f, 0x08, 0xd4, 0x95,
	0xf6, 0x8b, 0xde, 0x18, 0x8f, 0x9a, 0x33, 0xf1, 0xa1, 0x79, 0x83, 0xb2, 0xe0, 0x22, 0xb3, 0x6b,
	0xb2, 0x0e, 0xc3, 0x09, 0x3c, 0xbc, 0xca, 0x12, 0x94, 0x6b, 0x36, 0x4c, 0x4a, 0x7e, 0x4b, 0xe6,
	0x43, 0xd3, 0x66, 0xac, 0xb6, 0xeb, 0x70, 0xeb, 0x07, 0xb7, 0xe2, 0x87, 0xf0, 0x2f, 0x07, 0x0e,
	0x06, 0x42, 0xcc, 0x39, 0x6e, 0xb7, 0xf8, 0xda, 0xc4, 0x27, 0xa2, 0x50, 0x76, 0xa6, 0x0a, 0xa2,
	0xf3, 0xbf, 0x88, 0xa5, 0xcc, 0x58, 0x3a, 0xe4, 0xd2, 0x0c, 0xe7, 0xd1, 0x0a, 0xa2, 0xf3, 0xa9,
	0x19, 0x68, 0xc8, 0x65, 0x61, 0xb6, 0xd8, 0xa3, 0x15, 0x44, 0xe7, 0x79, 0x56, 0x28, 0x96, 0xc5,
	0x68, 0x17, 0xd9, 0xa3, 0x15, 0x24, 0xfc, 0x1c, 0xda, 0xa7, 0x22, 0x9e, 0x6f, 0x5f, 0x20, 0xbd,
	0x0a, 0xe5, 0xf7, 0xc7, 0x89, 0xbd, 0xd9, 0x4d, 0x1c, 0x7e, 0x01, 0x87, 0xc7, 0xcb, 0x84, 0x2b,
	0x21, 0x5f, 0x0b, 0x85, 0xff, 0xa3, 0xe1, 0x07, 0xf8, 0x60, 0x30, 0xc3, 0x78, 0x7e, 0x9c, 0xa6,
	0xa5, 0x6e, 0x85, 0x6d, 0xea, 0xc3, 0x51, 0x7c, 0x3b, 0x31, 0xb8, 0xd4, 0x6a, 0x39, 0x46, 0xad,
	0x7b, 0x73, 0xe1, 0x4f, 0xd0, 0x3d, 0x4f, 0x59, 0x8c, 0xda, 0xdc, 0xe7, 0x22, 0xe5, 0xf1, 0xca,
	0x70, 0x5b, 0xc6, 0x6f, 0xe1, 0xe3, 0xfc, 0x9e, 0x6c, 0x95, 0xf6, 0xdd, 0x05, 0xe1, 0x08, 0xde,
	0xa7, 0xe5, 0xcd, 0x15, 0x55, 0xd2, 0x1e, 0x10, 0x59, 0x85, 0xab, 0x6c, 0xf7, 0x64, 0xc2, 0x3f,
	0x6b, 0xf0, 0xd1, 0x0b, 0xf3, 0xf3, 0x17, 0x28, 0x6f, 0x78, 0x8c, 0xe3, 0xec, 0xad, 0xb0, 0x5c,
	0x27, 0xe0, 0x61, 0x96, 0xe4, 0x82, 0x67, 0xaa, 0xb0, 0x4f, 0xde, 0x67, 0xdb, 0x67, 0xe0, 0x1d,
	0x5d, 0xbd, 0x91, 0x6d, 0xa1, 0xdb, 0x66, 0xf2, 0x0a, 0x20, 0x97, 0x22, 0x47, 0xa9, 0x38, 0x16,
	0x7e, 0xd3, 0x50, 0x7d, 0xb5, 0x9b, 0xea, 0x7c, 0xd3, 0x53, 0x3e, 0x29, 0x15, 0x92, 0xee, 0xdf,
	0x0e, 0xb4, 0xd6, 0x9f, 0x22, 0x6d, 0x70, 0x79, 0x62, 0xcd, 0xe8, 0xf2, 0x44, 0xaf, 0x4c, 0x2e,
	0xa4, 0xb2, 0xbb, 0x61, 0xce, 0x1a, 0x9b, 0x69, 0xcb, 0x96, 0x96, 0x33, 0x67, 0xf3, 0x8a, 0xea,
	0x09, 0x62, 0x91, 0xfa, 0x75, 0x83, 0x6f, 0x62, 0x5d, 0xcf, 0x96, 0x6a, 0xe6, 0x37, 0xcc, 0x9a,
	0x98, 0xb3, 0x36, 0x27, 0xfe, 0xae, 0x30, 0xd3, 0x9b, 0x56, 0x4a, 0xe2, 0xd1, 0x0a, 0xd2, 0xfd,
	0x0e, 0x1e, 0xdc, 0x99, 0x79, 0xd7, 0x93, 0xe3, 0xd9, 0x27, 0xe7, 0x99, 0xfb, 0xd4, 0x79, 0xf1,
	0x08, 0x1e, 0x0a, 0x39, 0xed, 0xb1, 0x9c, 0xc5, 0x33, 0xec, 0x69, 0x4f, 0xce, 0x11, 0x73, 0x94,
	0xa5, 0x4c, 0x27, 0xce, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x06, 0xce, 0x51, 0xae, 0x79, 0x07,
	0x00, 0x00,
}