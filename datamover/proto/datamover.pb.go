// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datamover.proto

package datamover

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

type KV struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KV) Reset()         { *m = KV{} }
func (m *KV) String() string { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()    {}
func (*KV) Descriptor() ([]byte, []int) {
	return fileDescriptor_302d3374d95ec8eb, []int{0}
}

func (m *KV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KV.Unmarshal(m, b)
}
func (m *KV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KV.Marshal(b, m, deterministic)
}
func (m *KV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KV.Merge(m, src)
}
func (m *KV) XXX_Size() int {
	return xxx_messageInfo_KV.Size(m)
}
func (m *KV) XXX_DiscardUnknown() {
	xxx_messageInfo_KV.DiscardUnknown(m)
}

var xxx_messageInfo_KV proto.InternalMessageInfo

func (m *KV) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KV) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Filter struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Tag                  []*KV    `protobuf:"bytes,2,rep,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_302d3374d95ec8eb, []int{1}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Filter) GetTag() []*KV {
	if m != nil {
		return m.Tag
	}
	return nil
}

type Connector struct {
	Type                 string   `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	BucketName           string   `protobuf:"bytes,2,opt,name=BucketName,proto3" json:"BucketName,omitempty"`
	ConnConfig           []*KV    `protobuf:"bytes,3,rep,name=ConnConfig,proto3" json:"ConnConfig,omitempty"`
	BackendName          string   `protobuf:"bytes,4,opt,name=BackendName,proto3" json:"BackendName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connector) Reset()         { *m = Connector{} }
func (m *Connector) String() string { return proto.CompactTextString(m) }
func (*Connector) ProtoMessage()    {}
func (*Connector) Descriptor() ([]byte, []int) {
	return fileDescriptor_302d3374d95ec8eb, []int{2}
}

func (m *Connector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connector.Unmarshal(m, b)
}
func (m *Connector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connector.Marshal(b, m, deterministic)
}
func (m *Connector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connector.Merge(m, src)
}
func (m *Connector) XXX_Size() int {
	return xxx_messageInfo_Connector.Size(m)
}
func (m *Connector) XXX_DiscardUnknown() {
	xxx_messageInfo_Connector.DiscardUnknown(m)
}

var xxx_messageInfo_Connector proto.InternalMessageInfo

func (m *Connector) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Connector) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

func (m *Connector) GetConnConfig() []*KV {
	if m != nil {
		return m.ConnConfig
	}
	return nil
}

func (m *Connector) GetBackendName() string {
	if m != nil {
		return m.BackendName
	}
	return ""
}

type RunJobRequest struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TenanId              string     `protobuf:"bytes,2,opt,name=tenanId,proto3" json:"tenanId,omitempty"`
	UserId               string     `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	SourceConn           *Connector `protobuf:"bytes,4,opt,name=sourceConn,proto3" json:"sourceConn,omitempty"`
	DestConn             *Connector `protobuf:"bytes,5,opt,name=destConn,proto3" json:"destConn,omitempty"`
	Filt                 *Filter    `protobuf:"bytes,6,opt,name=filt,proto3" json:"filt,omitempty"`
	RemainSource         bool       `protobuf:"varint,7,opt,name=remainSource,proto3" json:"remainSource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RunJobRequest) Reset()         { *m = RunJobRequest{} }
func (m *RunJobRequest) String() string { return proto.CompactTextString(m) }
func (*RunJobRequest) ProtoMessage()    {}
func (*RunJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_302d3374d95ec8eb, []int{3}
}

func (m *RunJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunJobRequest.Unmarshal(m, b)
}
func (m *RunJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunJobRequest.Marshal(b, m, deterministic)
}
func (m *RunJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunJobRequest.Merge(m, src)
}
func (m *RunJobRequest) XXX_Size() int {
	return xxx_messageInfo_RunJobRequest.Size(m)
}
func (m *RunJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunJobRequest proto.InternalMessageInfo

func (m *RunJobRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RunJobRequest) GetTenanId() string {
	if m != nil {
		return m.TenanId
	}
	return ""
}

func (m *RunJobRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *RunJobRequest) GetSourceConn() *Connector {
	if m != nil {
		return m.SourceConn
	}
	return nil
}

func (m *RunJobRequest) GetDestConn() *Connector {
	if m != nil {
		return m.DestConn
	}
	return nil
}

func (m *RunJobRequest) GetFilt() *Filter {
	if m != nil {
		return m.Filt
	}
	return nil
}

func (m *RunJobRequest) GetRemainSource() bool {
	if m != nil {
		return m.RemainSource
	}
	return false
}

type RunJobResponse struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunJobResponse) Reset()         { *m = RunJobResponse{} }
func (m *RunJobResponse) String() string { return proto.CompactTextString(m) }
func (*RunJobResponse) ProtoMessage()    {}
func (*RunJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_302d3374d95ec8eb, []int{4}
}

func (m *RunJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunJobResponse.Unmarshal(m, b)
}
func (m *RunJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunJobResponse.Marshal(b, m, deterministic)
}
func (m *RunJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunJobResponse.Merge(m, src)
}
func (m *RunJobResponse) XXX_Size() int {
	return xxx_messageInfo_RunJobResponse.Size(m)
}
func (m *RunJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RunJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RunJobResponse proto.InternalMessageInfo

func (m *RunJobResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type LifecycleActionRequest struct {
	ObjKey               string   `protobuf:"bytes,1,opt,name=objKey,proto3" json:"objKey,omitempty"`
	BucketName           string   `protobuf:"bytes,2,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	VersionId            string   `protobuf:"bytes,3,opt,name=versionId,proto3" json:"versionId,omitempty"`
	StorageMeta          string   `protobuf:"bytes,4,opt,name=storageMeta,proto3" json:"storageMeta,omitempty"`
	ObjectId             string   `protobuf:"bytes,5,opt,name=objectId,proto3" json:"objectId,omitempty"`
	Action               int32    `protobuf:"varint,6,opt,name=action,proto3" json:"action,omitempty"`
	SourceTier           int32    `protobuf:"varint,7,opt,name=sourceTier,proto3" json:"sourceTier,omitempty"`
	TargetTier           int32    `protobuf:"varint,8,opt,name=targetTier,proto3" json:"targetTier,omitempty"`
	SourceBackend        string   `protobuf:"bytes,9,opt,name=sourceBackend,proto3" json:"sourceBackend,omitempty"`
	TargetBackend        string   `protobuf:"bytes,10,opt,name=targetBackend,proto3" json:"targetBackend,omitempty"`
	ObjSize              int64    `protobuf:"varint,11,opt,name=objSize,proto3" json:"objSize,omitempty"`
	UploadId             string   `protobuf:"bytes,12,opt,name=uploadId,proto3" json:"uploadId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LifecycleActionRequest) Reset()         { *m = LifecycleActionRequest{} }
func (m *LifecycleActionRequest) String() string { return proto.CompactTextString(m) }
func (*LifecycleActionRequest) ProtoMessage()    {}
func (*LifecycleActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_302d3374d95ec8eb, []int{5}
}

func (m *LifecycleActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LifecycleActionRequest.Unmarshal(m, b)
}
func (m *LifecycleActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LifecycleActionRequest.Marshal(b, m, deterministic)
}
func (m *LifecycleActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LifecycleActionRequest.Merge(m, src)
}
func (m *LifecycleActionRequest) XXX_Size() int {
	return xxx_messageInfo_LifecycleActionRequest.Size(m)
}
func (m *LifecycleActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LifecycleActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LifecycleActionRequest proto.InternalMessageInfo

func (m *LifecycleActionRequest) GetObjKey() string {
	if m != nil {
		return m.ObjKey
	}
	return ""
}

func (m *LifecycleActionRequest) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

func (m *LifecycleActionRequest) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *LifecycleActionRequest) GetStorageMeta() string {
	if m != nil {
		return m.StorageMeta
	}
	return ""
}

func (m *LifecycleActionRequest) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *LifecycleActionRequest) GetAction() int32 {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *LifecycleActionRequest) GetSourceTier() int32 {
	if m != nil {
		return m.SourceTier
	}
	return 0
}

func (m *LifecycleActionRequest) GetTargetTier() int32 {
	if m != nil {
		return m.TargetTier
	}
	return 0
}

func (m *LifecycleActionRequest) GetSourceBackend() string {
	if m != nil {
		return m.SourceBackend
	}
	return ""
}

func (m *LifecycleActionRequest) GetTargetBackend() string {
	if m != nil {
		return m.TargetBackend
	}
	return ""
}

func (m *LifecycleActionRequest) GetObjSize() int64 {
	if m != nil {
		return m.ObjSize
	}
	return 0
}

func (m *LifecycleActionRequest) GetUploadId() string {
	if m != nil {
		return m.UploadId
	}
	return ""
}

type LifecycleActionResonse struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LifecycleActionResonse) Reset()         { *m = LifecycleActionResonse{} }
func (m *LifecycleActionResonse) String() string { return proto.CompactTextString(m) }
func (*LifecycleActionResonse) ProtoMessage()    {}
func (*LifecycleActionResonse) Descriptor() ([]byte, []int) {
	return fileDescriptor_302d3374d95ec8eb, []int{6}
}

func (m *LifecycleActionResonse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LifecycleActionResonse.Unmarshal(m, b)
}
func (m *LifecycleActionResonse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LifecycleActionResonse.Marshal(b, m, deterministic)
}
func (m *LifecycleActionResonse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LifecycleActionResonse.Merge(m, src)
}
func (m *LifecycleActionResonse) XXX_Size() int {
	return xxx_messageInfo_LifecycleActionResonse.Size(m)
}
func (m *LifecycleActionResonse) XXX_DiscardUnknown() {
	xxx_messageInfo_LifecycleActionResonse.DiscardUnknown(m)
}

var xxx_messageInfo_LifecycleActionResonse proto.InternalMessageInfo

func (m *LifecycleActionResonse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*KV)(nil), "KV")
	proto.RegisterType((*Filter)(nil), "Filter")
	proto.RegisterType((*Connector)(nil), "Connector")
	proto.RegisterType((*RunJobRequest)(nil), "RunJobRequest")
	proto.RegisterType((*RunJobResponse)(nil), "RunJobResponse")
	proto.RegisterType((*LifecycleActionRequest)(nil), "LifecycleActionRequest")
	proto.RegisterType((*LifecycleActionResonse)(nil), "LifecycleActionResonse")
}

func init() {
	proto.RegisterFile("datamover.proto", fileDescriptor_302d3374d95ec8eb)
}

var fileDescriptor_302d3374d95ec8eb = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0x4b, 0x6f, 0xd3, 0x4c,
	0x14, 0xfd, 0x1c, 0x37, 0x6e, 0x7c, 0xd3, 0xa6, 0x1f, 0x23, 0x08, 0x56, 0x40, 0x28, 0x32, 0x08,
	0x45, 0x05, 0x79, 0x11, 0x16, 0xac, 0x69, 0x11, 0x28, 0x04, 0x58, 0x4c, 0xab, 0xee, 0xc7, 0xf6,
	0x4d, 0x34, 0x79, 0xcc, 0x84, 0xf1, 0x38, 0x22, 0xec, 0xd8, 0xf4, 0xa7, 0xf2, 0x3b, 0xd0, 0x8c,
	0x1f, 0x71, 0x68, 0x76, 0x3e, 0xe7, 0xde, 0xb9, 0x8f, 0x73, 0x6e, 0x02, 0x17, 0x29, 0xd3, 0x6c,
	0x2d, 0xb7, 0xa8, 0xa2, 0x8d, 0x92, 0x5a, 0x86, 0x6f, 0xa1, 0x35, 0xbd, 0x23, 0xff, 0x83, 0xbb,
	0xc4, 0x5d, 0xe0, 0x0c, 0x9d, 0x91, 0x4f, 0xcd, 0x27, 0x79, 0x0c, 0xed, 0x2d, 0x5b, 0xe5, 0x18,
	0xb4, 0x2c, 0x57, 0x80, 0xf0, 0x3d, 0x78, 0x9f, 0xf8, 0x4a, 0xa3, 0x22, 0x7d, 0xf0, 0x36, 0x0a,
	0x67, 0xfc, 0x67, 0xf9, 0xa8, 0x44, 0xe4, 0x09, 0xb8, 0x9a, 0xcd, 0x83, 0xd6, 0xd0, 0x1d, 0x75,
	0xc7, 0x6e, 0x34, 0xbd, 0xa3, 0x06, 0x87, 0xf7, 0x0e, 0xf8, 0xd7, 0x52, 0x08, 0x4c, 0xb4, 0x54,
	0x84, 0xc0, 0xc9, 0xed, 0x6e, 0x83, 0xe5, 0x53, 0xfb, 0x4d, 0x5e, 0x00, 0x5c, 0xe5, 0xc9, 0x12,
	0xf5, 0x77, 0xb6, 0xae, 0xba, 0x36, 0x18, 0xf2, 0x12, 0xc0, 0x14, 0xb8, 0x96, 0x62, 0xc6, 0xe7,
	0x81, 0xbb, 0xaf, 0xdf, 0xa0, 0xc9, 0x10, 0xba, 0x57, 0x2c, 0x59, 0xa2, 0x48, 0x6d, 0x95, 0x13,
	0x5b, 0xa5, 0x49, 0x85, 0x7f, 0x1c, 0x38, 0xa7, 0xb9, 0xf8, 0x22, 0x63, 0x8a, 0x3f, 0x72, 0xcc,
	0x34, 0xe9, 0x41, 0x8b, 0xa7, 0xe5, 0x28, 0x2d, 0x9e, 0x92, 0x00, 0x4e, 0x35, 0x0a, 0x26, 0x26,
	0x69, 0x39, 0x45, 0x05, 0xcd, 0xce, 0x79, 0x86, 0x6a, 0x92, 0x06, 0x6e, 0xb1, 0x73, 0x81, 0xc8,
	0x25, 0x40, 0x26, 0x73, 0x95, 0xa0, 0x99, 0xc4, 0x36, 0xed, 0x8e, 0x21, 0xaa, 0xd7, 0xa5, 0x8d,
	0x28, 0x79, 0x0d, 0x9d, 0x14, 0x33, 0x6d, 0x33, 0xdb, 0x0f, 0x32, 0xeb, 0x18, 0x79, 0x06, 0x27,
	0x33, 0xbe, 0xd2, 0x81, 0x67, 0x73, 0x4e, 0xa3, 0x42, 0x76, 0x6a, 0x49, 0x12, 0xc2, 0x99, 0xc2,
	0x35, 0xe3, 0xe2, 0xc6, 0x16, 0x0e, 0x4e, 0x87, 0xce, 0xa8, 0x43, 0x0f, 0xb8, 0x30, 0x84, 0x5e,
	0xb5, 0x67, 0xb6, 0x91, 0x22, 0x43, 0x63, 0x32, 0x2a, 0x55, 0x99, 0x8c, 0x4a, 0x85, 0xf7, 0x2e,
	0xf4, 0xbf, 0xf2, 0x19, 0x26, 0xbb, 0x64, 0x85, 0x1f, 0x12, 0xcd, 0xa5, 0xa8, 0x54, 0xe9, 0x83,
	0x27, 0xe3, 0xc5, 0xb4, 0x3e, 0x8a, 0x12, 0x19, 0x9b, 0xe2, 0x07, 0x36, 0xed, 0x19, 0xf2, 0x1c,
	0xfc, 0x2d, 0xaa, 0x8c, 0x4b, 0x51, 0xcb, 0xb4, 0x27, 0x8c, 0x3f, 0x99, 0x96, 0x8a, 0xcd, 0xf1,
	0x1b, 0x6a, 0x56, 0xf9, 0xd3, 0xa0, 0xc8, 0x00, 0x3a, 0x32, 0x5e, 0x60, 0xa2, 0x27, 0xa9, 0xd5,
	0xc7, 0xa7, 0x35, 0x36, 0x33, 0x31, 0x3b, 0xa4, 0x55, 0xa5, 0x4d, 0x4b, 0x64, 0x66, 0x2a, 0x14,
	0xbe, 0xe5, 0xa8, 0xac, 0x18, 0x6d, 0xda, 0x60, 0x4c, 0x5c, 0x33, 0x35, 0x47, 0x6d, 0xe3, 0x9d,
	0x22, 0xbe, 0x67, 0xc8, 0x2b, 0x38, 0x2f, 0xb2, 0xcb, 0x43, 0x09, 0x7c, 0xdb, 0xf8, 0x90, 0x34,
	0x59, 0xc5, 0x9b, 0x2a, 0x0b, 0x8a, 0xac, 0x03, 0xd2, 0x5c, 0x8f, 0x8c, 0x17, 0x37, 0xfc, 0x17,
	0x06, 0xdd, 0xa1, 0x33, 0x72, 0x69, 0x05, 0xcd, 0x66, 0xf9, 0x66, 0x25, 0x59, 0x3a, 0x49, 0x83,
	0xb3, 0x62, 0xb3, 0x0a, 0x87, 0x97, 0x47, 0x7c, 0xc8, 0x8e, 0x9b, 0x36, 0xfe, 0xed, 0x80, 0x5f,
	0xff, 0x8a, 0xc9, 0x1b, 0xf0, 0x68, 0x2e, 0x16, 0x32, 0x26, 0xbd, 0xe8, 0xe0, 0xae, 0x07, 0x17,
	0xd1, 0xa1, 0xff, 0xe1, 0x7f, 0xe4, 0x33, 0x3c, 0xfa, 0x28, 0xff, 0x69, 0x44, 0x9e, 0x46, 0xc7,
	0x4f, 0x60, 0x70, 0x24, 0x90, 0x15, 0x85, 0x62, 0xcf, 0xfe, 0x79, 0xbc, 0xfb, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x52, 0x27, 0x34, 0x4b, 0x4f, 0x04, 0x00, 0x00,
}
