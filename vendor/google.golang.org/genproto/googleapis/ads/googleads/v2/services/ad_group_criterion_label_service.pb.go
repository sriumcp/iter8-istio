// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/ad_group_criterion_label_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for
// [AdGroupCriterionLabelService.GetAdGroupCriterionLabel][google.ads.googleads.v2.services.AdGroupCriterionLabelService.GetAdGroupCriterionLabel].
type GetAdGroupCriterionLabelRequest struct {
	// The resource name of the ad group criterion label to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupCriterionLabelRequest) Reset()         { *m = GetAdGroupCriterionLabelRequest{} }
func (m *GetAdGroupCriterionLabelRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupCriterionLabelRequest) ProtoMessage()    {}
func (*GetAdGroupCriterionLabelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3ff5931695c10a3, []int{0}
}

func (m *GetAdGroupCriterionLabelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupCriterionLabelRequest.Unmarshal(m, b)
}
func (m *GetAdGroupCriterionLabelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupCriterionLabelRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupCriterionLabelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupCriterionLabelRequest.Merge(m, src)
}
func (m *GetAdGroupCriterionLabelRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupCriterionLabelRequest.Size(m)
}
func (m *GetAdGroupCriterionLabelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupCriterionLabelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupCriterionLabelRequest proto.InternalMessageInfo

func (m *GetAdGroupCriterionLabelRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [AdGroupCriterionLabelService.MutateAdGroupCriterionLabels][google.ads.googleads.v2.services.AdGroupCriterionLabelService.MutateAdGroupCriterionLabels].
type MutateAdGroupCriterionLabelsRequest struct {
	// ID of the customer whose ad group criterion labels are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on ad group criterion labels.
	Operations []*AdGroupCriterionLabelOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupCriterionLabelsRequest) Reset()         { *m = MutateAdGroupCriterionLabelsRequest{} }
func (m *MutateAdGroupCriterionLabelsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupCriterionLabelsRequest) ProtoMessage()    {}
func (*MutateAdGroupCriterionLabelsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3ff5931695c10a3, []int{1}
}

func (m *MutateAdGroupCriterionLabelsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupCriterionLabelsRequest.Unmarshal(m, b)
}
func (m *MutateAdGroupCriterionLabelsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupCriterionLabelsRequest.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupCriterionLabelsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupCriterionLabelsRequest.Merge(m, src)
}
func (m *MutateAdGroupCriterionLabelsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupCriterionLabelsRequest.Size(m)
}
func (m *MutateAdGroupCriterionLabelsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupCriterionLabelsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupCriterionLabelsRequest proto.InternalMessageInfo

func (m *MutateAdGroupCriterionLabelsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdGroupCriterionLabelsRequest) GetOperations() []*AdGroupCriterionLabelOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateAdGroupCriterionLabelsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateAdGroupCriterionLabelsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on an ad group criterion label.
type AdGroupCriterionLabelOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdGroupCriterionLabelOperation_Create
	//	*AdGroupCriterionLabelOperation_Remove
	Operation            isAdGroupCriterionLabelOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *AdGroupCriterionLabelOperation) Reset()         { *m = AdGroupCriterionLabelOperation{} }
func (m *AdGroupCriterionLabelOperation) String() string { return proto.CompactTextString(m) }
func (*AdGroupCriterionLabelOperation) ProtoMessage()    {}
func (*AdGroupCriterionLabelOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3ff5931695c10a3, []int{2}
}

func (m *AdGroupCriterionLabelOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupCriterionLabelOperation.Unmarshal(m, b)
}
func (m *AdGroupCriterionLabelOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupCriterionLabelOperation.Marshal(b, m, deterministic)
}
func (m *AdGroupCriterionLabelOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupCriterionLabelOperation.Merge(m, src)
}
func (m *AdGroupCriterionLabelOperation) XXX_Size() int {
	return xxx_messageInfo_AdGroupCriterionLabelOperation.Size(m)
}
func (m *AdGroupCriterionLabelOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupCriterionLabelOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupCriterionLabelOperation proto.InternalMessageInfo

type isAdGroupCriterionLabelOperation_Operation interface {
	isAdGroupCriterionLabelOperation_Operation()
}

type AdGroupCriterionLabelOperation_Create struct {
	Create *resources.AdGroupCriterionLabel `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupCriterionLabelOperation_Remove struct {
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*AdGroupCriterionLabelOperation_Create) isAdGroupCriterionLabelOperation_Operation() {}

func (*AdGroupCriterionLabelOperation_Remove) isAdGroupCriterionLabelOperation_Operation() {}

func (m *AdGroupCriterionLabelOperation) GetOperation() isAdGroupCriterionLabelOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdGroupCriterionLabelOperation) GetCreate() *resources.AdGroupCriterionLabel {
	if x, ok := m.GetOperation().(*AdGroupCriterionLabelOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdGroupCriterionLabelOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdGroupCriterionLabelOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupCriterionLabelOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupCriterionLabelOperation_Create)(nil),
		(*AdGroupCriterionLabelOperation_Remove)(nil),
	}
}

// Response message for an ad group criterion labels mutate.
type MutateAdGroupCriterionLabelsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateAdGroupCriterionLabelResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *MutateAdGroupCriterionLabelsResponse) Reset()         { *m = MutateAdGroupCriterionLabelsResponse{} }
func (m *MutateAdGroupCriterionLabelsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupCriterionLabelsResponse) ProtoMessage()    {}
func (*MutateAdGroupCriterionLabelsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3ff5931695c10a3, []int{3}
}

func (m *MutateAdGroupCriterionLabelsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupCriterionLabelsResponse.Unmarshal(m, b)
}
func (m *MutateAdGroupCriterionLabelsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupCriterionLabelsResponse.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupCriterionLabelsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupCriterionLabelsResponse.Merge(m, src)
}
func (m *MutateAdGroupCriterionLabelsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupCriterionLabelsResponse.Size(m)
}
func (m *MutateAdGroupCriterionLabelsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupCriterionLabelsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupCriterionLabelsResponse proto.InternalMessageInfo

func (m *MutateAdGroupCriterionLabelsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateAdGroupCriterionLabelsResponse) GetResults() []*MutateAdGroupCriterionLabelResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for an ad group criterion label mutate.
type MutateAdGroupCriterionLabelResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupCriterionLabelResult) Reset()         { *m = MutateAdGroupCriterionLabelResult{} }
func (m *MutateAdGroupCriterionLabelResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupCriterionLabelResult) ProtoMessage()    {}
func (*MutateAdGroupCriterionLabelResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3ff5931695c10a3, []int{4}
}

func (m *MutateAdGroupCriterionLabelResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupCriterionLabelResult.Unmarshal(m, b)
}
func (m *MutateAdGroupCriterionLabelResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupCriterionLabelResult.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupCriterionLabelResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupCriterionLabelResult.Merge(m, src)
}
func (m *MutateAdGroupCriterionLabelResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupCriterionLabelResult.Size(m)
}
func (m *MutateAdGroupCriterionLabelResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupCriterionLabelResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupCriterionLabelResult proto.InternalMessageInfo

func (m *MutateAdGroupCriterionLabelResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupCriterionLabelRequest)(nil), "google.ads.googleads.v2.services.GetAdGroupCriterionLabelRequest")
	proto.RegisterType((*MutateAdGroupCriterionLabelsRequest)(nil), "google.ads.googleads.v2.services.MutateAdGroupCriterionLabelsRequest")
	proto.RegisterType((*AdGroupCriterionLabelOperation)(nil), "google.ads.googleads.v2.services.AdGroupCriterionLabelOperation")
	proto.RegisterType((*MutateAdGroupCriterionLabelsResponse)(nil), "google.ads.googleads.v2.services.MutateAdGroupCriterionLabelsResponse")
	proto.RegisterType((*MutateAdGroupCriterionLabelResult)(nil), "google.ads.googleads.v2.services.MutateAdGroupCriterionLabelResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/ad_group_criterion_label_service.proto", fileDescriptor_d3ff5931695c10a3)
}

var fileDescriptor_d3ff5931695c10a3 = []byte{
	// 682 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcf, 0x6b, 0xd4, 0x4e,
	0x14, 0xff, 0x26, 0x2d, 0xfd, 0xda, 0xd9, 0xaa, 0x30, 0x22, 0x86, 0xb5, 0xd8, 0x6d, 0x5a, 0xb0,
	0xec, 0x21, 0x81, 0x78, 0x29, 0x29, 0x85, 0x4d, 0x4b, 0x77, 0x2b, 0xa8, 0x2d, 0x29, 0xf4, 0x20,
	0x2b, 0x71, 0x9a, 0x8c, 0x21, 0x90, 0x64, 0xe2, 0xcc, 0x64, 0xa1, 0x94, 0x5e, 0x3c, 0x79, 0xf3,
	0xe0, 0xd1, 0x9b, 0x47, 0xff, 0x94, 0x82, 0x27, 0x4f, 0xde, 0x3d, 0x79, 0xd7, 0xb3, 0x4c, 0x26,
	0xb3, 0xdb, 0x96, 0xfd, 0x21, 0xf6, 0xf6, 0xf6, 0xcd, 0x27, 0x9f, 0xcf, 0xfb, 0xbc, 0xf7, 0x66,
	0x16, 0xf4, 0x62, 0x42, 0xe2, 0x14, 0xdb, 0x28, 0x62, 0xb6, 0x0c, 0x45, 0x34, 0x70, 0x6c, 0x86,
	0xe9, 0x20, 0x09, 0x31, 0xb3, 0x51, 0x14, 0xc4, 0x94, 0x94, 0x45, 0x10, 0xd2, 0x84, 0x63, 0x9a,
	0x90, 0x3c, 0x48, 0xd1, 0x09, 0x4e, 0x83, 0x1a, 0x61, 0x15, 0x94, 0x70, 0x02, 0x5b, 0xf2, 0x6b,
	0x0b, 0x45, 0xcc, 0x1a, 0x12, 0x59, 0x03, 0xc7, 0x52, 0x44, 0xcd, 0xce, 0x24, 0x29, 0x8a, 0x19,
	0x29, 0xe9, 0x34, 0x2d, 0xa9, 0xd1, 0x5c, 0x56, 0x0c, 0x45, 0x62, 0xa3, 0x3c, 0x27, 0x1c, 0xf1,
	0x84, 0xe4, 0xac, 0x3e, 0x7d, 0x50, 0x9f, 0xd2, 0x22, 0xb4, 0x19, 0x47, 0xbc, 0xbc, 0x7e, 0x20,
	0x3e, 0x0b, 0xd3, 0x04, 0xe7, 0x5c, 0x1e, 0x98, 0x5d, 0xb0, 0xd2, 0xc3, 0xdc, 0x8b, 0x7a, 0x42,
	0x73, 0x57, 0x49, 0x3e, 0x13, 0x8a, 0x3e, 0x7e, 0x5b, 0x62, 0xc6, 0xe1, 0x1a, 0xb8, 0xad, 0xca,
	0x0b, 0x72, 0x94, 0x61, 0x43, 0x6b, 0x69, 0x1b, 0x8b, 0xfe, 0x92, 0x4a, 0xbe, 0x40, 0x19, 0x36,
	0x7f, 0x6b, 0x60, 0xed, 0x79, 0xc9, 0x11, 0xc7, 0x63, 0xb9, 0x98, 0x22, 0x5b, 0x01, 0x8d, 0xb0,
	0x64, 0x9c, 0x64, 0x98, 0x06, 0x49, 0x54, 0x53, 0x01, 0x95, 0x7a, 0x1a, 0xc1, 0xd7, 0x00, 0x90,
	0x02, 0x53, 0x69, 0xcb, 0xd0, 0x5b, 0x73, 0x1b, 0x0d, 0xa7, 0x63, 0xcd, 0xea, 0xac, 0x35, 0x56,
	0xf5, 0x40, 0x11, 0xf9, 0x97, 0x38, 0xe1, 0x63, 0x70, 0xb7, 0x40, 0x94, 0x27, 0x28, 0x0d, 0xde,
	0xa0, 0x24, 0x2d, 0x29, 0x36, 0xe6, 0x5a, 0xda, 0xc6, 0x2d, 0xff, 0x4e, 0x9d, 0xee, 0xca, 0xac,
	0x30, 0x3e, 0x40, 0x69, 0x12, 0x21, 0x8e, 0x03, 0x92, 0xa7, 0xa7, 0xc6, 0x7c, 0x05, 0x5b, 0x52,
	0xc9, 0x83, 0x3c, 0x3d, 0x35, 0x3f, 0x69, 0xe0, 0xd1, 0x74, 0x71, 0xe8, 0x83, 0x85, 0x90, 0x62,
	0xc4, 0x65, 0xe7, 0x1a, 0xce, 0xe6, 0x44, 0x3b, 0xc3, 0x35, 0x18, 0xef, 0x67, 0xff, 0x3f, 0xbf,
	0x66, 0x82, 0x06, 0x58, 0xa0, 0x38, 0x23, 0x03, 0x6c, 0xe8, 0xa2, 0x85, 0xe2, 0x44, 0xfe, 0xde,
	0x69, 0x80, 0xc5, 0xa1, 0x59, 0xf3, 0xab, 0x06, 0xd6, 0xa7, 0x8f, 0x85, 0x15, 0x24, 0x67, 0x18,
	0x76, 0xc1, 0xfd, 0x6b, 0x4d, 0x09, 0x30, 0xa5, 0x84, 0x56, 0xad, 0x69, 0x38, 0x50, 0x95, 0x4c,
	0x8b, 0xd0, 0x3a, 0xaa, 0x36, 0xcb, 0xbf, 0x77, 0xb5, 0x5d, 0x7b, 0x02, 0x0e, 0x5f, 0x81, 0xff,
	0x29, 0x66, 0x65, 0xca, 0xd5, 0xec, 0x76, 0x67, 0xcf, 0x6e, 0x4a, 0x81, 0x7e, 0xc5, 0xe5, 0x2b,
	0x4e, 0x73, 0x1f, 0xac, 0xce, 0x44, 0xff, 0xd5, 0xc2, 0x3a, 0x1f, 0xe6, 0xc1, 0xf2, 0x58, 0x92,
	0x23, 0x59, 0x16, 0xfc, 0xae, 0x01, 0x63, 0xd2, 0xd5, 0x80, 0xde, 0x6c, 0x57, 0x33, 0xae, 0x55,
	0xf3, 0x9f, 0xb7, 0xc0, 0xec, 0xbc, 0xfb, 0xf6, 0xe3, 0xa3, 0xee, 0xc2, 0x4d, 0xf1, 0x72, 0x9c,
	0x5d, 0xb1, 0xba, 0xad, 0x6e, 0x12, 0xb3, 0xdb, 0x36, 0x1a, 0x3b, 0x72, 0xbb, 0x7d, 0x0e, 0x7f,
	0x69, 0x60, 0x79, 0xda, 0x5a, 0xc0, 0xbd, 0x1b, 0x4d, 0x4d, 0xdd, 0xf6, 0x66, 0xf7, 0xa6, 0x34,
	0x72, 0x3b, 0xcd, 0x6e, 0xe5, 0xb8, 0x63, 0x6e, 0x09, 0xc7, 0x23, 0x8b, 0x67, 0x97, 0x9e, 0x92,
	0xed, 0xf6, 0xf9, 0x04, 0xc3, 0x6e, 0x56, 0x49, 0xb8, 0x5a, 0xbb, 0xf9, 0xf0, 0xc2, 0x33, 0x46,
	0x65, 0xd4, 0x51, 0x91, 0x30, 0x2b, 0x24, 0xd9, 0xce, 0x7b, 0x1d, 0xac, 0x87, 0x24, 0x9b, 0x59,
	0xf2, 0xce, 0xea, 0xb4, 0xbd, 0x39, 0x14, 0xcf, 0xea, 0xa1, 0xf6, 0x72, 0xbf, 0xa6, 0x89, 0x49,
	0x8a, 0xf2, 0xd8, 0x22, 0x34, 0xb6, 0x63, 0x9c, 0x57, 0x8f, 0xae, 0x3d, 0x12, 0x9e, 0xfc, 0xa7,
	0xb3, 0xa5, 0x82, 0xcf, 0xfa, 0x5c, 0xcf, 0xf3, 0xbe, 0xe8, 0xad, 0x9e, 0x24, 0xf4, 0x22, 0x66,
	0xc9, 0x50, 0x44, 0xc7, 0x8e, 0x55, 0x0b, 0xb3, 0x0b, 0x05, 0xe9, 0x7b, 0x11, 0xeb, 0x0f, 0x21,
	0xfd, 0x63, 0xa7, 0xaf, 0x20, 0x3f, 0xf5, 0x75, 0x99, 0x77, 0x5d, 0x2f, 0x62, 0xae, 0x3b, 0x04,
	0xb9, 0xee, 0xb1, 0xe3, 0xba, 0x0a, 0x76, 0xb2, 0x50, 0xd5, 0xf9, 0xe4, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x28, 0x6d, 0xc2, 0x67, 0x1b, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdGroupCriterionLabelServiceClient is the client API for AdGroupCriterionLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupCriterionLabelServiceClient interface {
	// Returns the requested ad group criterion label in full detail.
	GetAdGroupCriterionLabel(ctx context.Context, in *GetAdGroupCriterionLabelRequest, opts ...grpc.CallOption) (*resources.AdGroupCriterionLabel, error)
	// Creates and removes ad group criterion labels.
	// Operation statuses are returned.
	MutateAdGroupCriterionLabels(ctx context.Context, in *MutateAdGroupCriterionLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupCriterionLabelsResponse, error)
}

type adGroupCriterionLabelServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdGroupCriterionLabelServiceClient(cc *grpc.ClientConn) AdGroupCriterionLabelServiceClient {
	return &adGroupCriterionLabelServiceClient{cc}
}

func (c *adGroupCriterionLabelServiceClient) GetAdGroupCriterionLabel(ctx context.Context, in *GetAdGroupCriterionLabelRequest, opts ...grpc.CallOption) (*resources.AdGroupCriterionLabel, error) {
	out := new(resources.AdGroupCriterionLabel)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AdGroupCriterionLabelService/GetAdGroupCriterionLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupCriterionLabelServiceClient) MutateAdGroupCriterionLabels(ctx context.Context, in *MutateAdGroupCriterionLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupCriterionLabelsResponse, error) {
	out := new(MutateAdGroupCriterionLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AdGroupCriterionLabelService/MutateAdGroupCriterionLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupCriterionLabelServiceServer is the server API for AdGroupCriterionLabelService service.
type AdGroupCriterionLabelServiceServer interface {
	// Returns the requested ad group criterion label in full detail.
	GetAdGroupCriterionLabel(context.Context, *GetAdGroupCriterionLabelRequest) (*resources.AdGroupCriterionLabel, error)
	// Creates and removes ad group criterion labels.
	// Operation statuses are returned.
	MutateAdGroupCriterionLabels(context.Context, *MutateAdGroupCriterionLabelsRequest) (*MutateAdGroupCriterionLabelsResponse, error)
}

// UnimplementedAdGroupCriterionLabelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupCriterionLabelServiceServer struct {
}

func (*UnimplementedAdGroupCriterionLabelServiceServer) GetAdGroupCriterionLabel(ctx context.Context, req *GetAdGroupCriterionLabelRequest) (*resources.AdGroupCriterionLabel, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetAdGroupCriterionLabel not implemented")
}
func (*UnimplementedAdGroupCriterionLabelServiceServer) MutateAdGroupCriterionLabels(ctx context.Context, req *MutateAdGroupCriterionLabelsRequest) (*MutateAdGroupCriterionLabelsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateAdGroupCriterionLabels not implemented")
}

func RegisterAdGroupCriterionLabelServiceServer(s *grpc.Server, srv AdGroupCriterionLabelServiceServer) {
	s.RegisterService(&_AdGroupCriterionLabelService_serviceDesc, srv)
}

func _AdGroupCriterionLabelService_GetAdGroupCriterionLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupCriterionLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupCriterionLabelServiceServer).GetAdGroupCriterionLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AdGroupCriterionLabelService/GetAdGroupCriterionLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupCriterionLabelServiceServer).GetAdGroupCriterionLabel(ctx, req.(*GetAdGroupCriterionLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupCriterionLabelService_MutateAdGroupCriterionLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupCriterionLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupCriterionLabelServiceServer).MutateAdGroupCriterionLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AdGroupCriterionLabelService/MutateAdGroupCriterionLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupCriterionLabelServiceServer).MutateAdGroupCriterionLabels(ctx, req.(*MutateAdGroupCriterionLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupCriterionLabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.AdGroupCriterionLabelService",
	HandlerType: (*AdGroupCriterionLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupCriterionLabel",
			Handler:    _AdGroupCriterionLabelService_GetAdGroupCriterionLabel_Handler,
		},
		{
			MethodName: "MutateAdGroupCriterionLabels",
			Handler:    _AdGroupCriterionLabelService_MutateAdGroupCriterionLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/ad_group_criterion_label_service.proto",
}
