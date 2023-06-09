// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: proto/payments.proto

package payments

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GatewayId int64   `protobuf:"varint,3,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	SourceId  int64   `protobuf:"varint,4,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	AccountId int64   `protobuf:"varint,5,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	ExtId     string  `protobuf:"bytes,6,opt,name=ext_id,json=extId,proto3" json:"ext_id,omitempty"`
	Cards     []*Card `protobuf:"bytes,7,rep,name=cards,proto3" json:"cards,omitempty"`
	Flags     int64   `protobuf:"varint,8,opt,name=flags,proto3" json:"flags,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_proto_payments_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetGatewayId() int64 {
	if x != nil {
		return x.GatewayId
	}
	return 0
}

func (x *Customer) GetSourceId() int64 {
	if x != nil {
		return x.SourceId
	}
	return 0
}

func (x *Customer) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *Customer) GetExtId() string {
	if x != nil {
		return x.ExtId
	}
	return ""
}

func (x *Customer) GetCards() []*Card {
	if x != nil {
		return x.Cards
	}
	return nil
}

func (x *Customer) GetFlags() int64 {
	if x != nil {
		return x.Flags
	}
	return 0
}

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ExtId    string `protobuf:"bytes,2,opt,name=ext_id,json=extId,proto3" json:"ext_id,omitempty"`
	Brand    string `protobuf:"bytes,3,opt,name=brand,proto3" json:"brand,omitempty"`
	Last4    string `protobuf:"bytes,4,opt,name=last4,proto3" json:"last4,omitempty"`
	ExpMonth uint32 `protobuf:"varint,5,opt,name=exp_month,json=expMonth,proto3" json:"exp_month,omitempty"`
	ExpYear  uint32 `protobuf:"varint,6,opt,name=exp_year,json=expYear,proto3" json:"exp_year,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_proto_payments_proto_rawDescGZIP(), []int{1}
}

func (x *Card) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Card) GetExtId() string {
	if x != nil {
		return x.ExtId
	}
	return ""
}

func (x *Card) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Card) GetLast4() string {
	if x != nil {
		return x.Last4
	}
	return ""
}

func (x *Card) GetExpMonth() uint32 {
	if x != nil {
		return x.ExpMonth
	}
	return 0
}

func (x *Card) GetExpYear() uint32 {
	if x != nil {
		return x.ExpYear
	}
	return 0
}

type CreateCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId  int64  `protobuf:"varint,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	OrgId     int64  `protobuf:"varint,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	AccountId int64  `protobuf:"varint,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateCustomerRequest) Reset() {
	*x = CreateCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerRequest) ProtoMessage() {}

func (x *CreateCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerRequest.ProtoReflect.Descriptor instead.
func (*CreateCustomerRequest) Descriptor() ([]byte, []int) {
	return file_proto_payments_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCustomerRequest) GetSourceId() int64 {
	if x != nil {
		return x.SourceId
	}
	return 0
}

func (x *CreateCustomerRequest) GetOrgId() int64 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

func (x *CreateCustomerRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *CreateCustomerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *CreateCustomerResponse) Reset() {
	*x = CreateCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerResponse) ProtoMessage() {}

func (x *CreateCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payments_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerResponse.ProtoReflect.Descriptor instead.
func (*CreateCustomerResponse) Descriptor() ([]byte, []int) {
	return file_proto_payments_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCustomerResponse) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

type GetCustomerByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId  int64 `protobuf:"varint,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	AccountId int64 `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *GetCustomerByIdRequest) Reset() {
	*x = GetCustomerByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payments_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerByIdRequest) ProtoMessage() {}

func (x *GetCustomerByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payments_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerByIdRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerByIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_payments_proto_rawDescGZIP(), []int{4}
}

func (x *GetCustomerByIdRequest) GetSourceId() int64 {
	if x != nil {
		return x.SourceId
	}
	return 0
}

func (x *GetCustomerByIdRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

type GetCustomerByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *GetCustomerByIdResponse) Reset() {
	*x = GetCustomerByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payments_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerByIdResponse) ProtoMessage() {}

func (x *GetCustomerByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payments_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerByIdResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerByIdResponse) Descriptor() ([]byte, []int) {
	return file_proto_payments_proto_rawDescGZIP(), []int{5}
}

func (x *GetCustomerByIdResponse) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

type AddCustomerPaymentMethodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId  int64 `protobuf:"varint,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	AccountId int64 `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Card      *Card `protobuf:"bytes,3,opt,name=card,proto3" json:"card,omitempty"`
}

func (x *AddCustomerPaymentMethodRequest) Reset() {
	*x = AddCustomerPaymentMethodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payments_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCustomerPaymentMethodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCustomerPaymentMethodRequest) ProtoMessage() {}

func (x *AddCustomerPaymentMethodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payments_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCustomerPaymentMethodRequest.ProtoReflect.Descriptor instead.
func (*AddCustomerPaymentMethodRequest) Descriptor() ([]byte, []int) {
	return file_proto_payments_proto_rawDescGZIP(), []int{6}
}

func (x *AddCustomerPaymentMethodRequest) GetSourceId() int64 {
	if x != nil {
		return x.SourceId
	}
	return 0
}

func (x *AddCustomerPaymentMethodRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *AddCustomerPaymentMethodRequest) GetCard() *Card {
	if x != nil {
		return x.Card
	}
	return nil
}

type AddCustomerPaymentMethodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddCustomerPaymentMethodResponse) Reset() {
	*x = AddCustomerPaymentMethodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payments_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCustomerPaymentMethodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCustomerPaymentMethodResponse) ProtoMessage() {}

func (x *AddCustomerPaymentMethodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payments_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCustomerPaymentMethodResponse.ProtoReflect.Descriptor instead.
func (*AddCustomerPaymentMethodResponse) Descriptor() ([]byte, []int) {
	return file_proto_payments_proto_rawDescGZIP(), []int{7}
}

func (x *AddCustomerPaymentMethodResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RemoveCustomerPaymentMethodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId  int64 `protobuf:"varint,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	AccountId int64 `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	CardId    int64 `protobuf:"varint,3,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`
}

func (x *RemoveCustomerPaymentMethodRequest) Reset() {
	*x = RemoveCustomerPaymentMethodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payments_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveCustomerPaymentMethodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCustomerPaymentMethodRequest) ProtoMessage() {}

func (x *RemoveCustomerPaymentMethodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payments_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCustomerPaymentMethodRequest.ProtoReflect.Descriptor instead.
func (*RemoveCustomerPaymentMethodRequest) Descriptor() ([]byte, []int) {
	return file_proto_payments_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveCustomerPaymentMethodRequest) GetSourceId() int64 {
	if x != nil {
		return x.SourceId
	}
	return 0
}

func (x *RemoveCustomerPaymentMethodRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *RemoveCustomerPaymentMethodRequest) GetCardId() int64 {
	if x != nil {
		return x.CardId
	}
	return 0
}

type RemoveCustomerPaymentMethodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RemoveCustomerPaymentMethodResponse) Reset() {
	*x = RemoveCustomerPaymentMethodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payments_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveCustomerPaymentMethodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCustomerPaymentMethodResponse) ProtoMessage() {}

func (x *RemoveCustomerPaymentMethodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payments_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCustomerPaymentMethodResponse.ProtoReflect.Descriptor instead.
func (*RemoveCustomerPaymentMethodResponse) Descriptor() ([]byte, []int) {
	return file_proto_payments_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveCustomerPaymentMethodResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_payments_proto protoreflect.FileDescriptor

var file_proto_payments_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0xdc, 0x01, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x65, 0x78, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x78,
	0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x61,
	0x72, 0x64, 0x52, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x22,
	0x91, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x78, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x78, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x73, 0x74, 0x34, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x73, 0x74, 0x34, 0x12, 0x1b, 0x0a, 0x09, 0x65,
	0x78, 0x70, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x65, 0x78, 0x70, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x5f,
	0x79, 0x65, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x78, 0x70, 0x59,
	0x65, 0x61, 0x72, 0x22, 0x7e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x54, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x81,
	0x01, 0x0a, 0x1f, 0x41, 0x64, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x63, 0x61,
	0x72, 0x64, 0x22, 0x3c, 0x0a, 0x20, 0x41, 0x64, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x79, 0x0a, 0x22, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x23, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xb4, 0x03, 0x0a,
	0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x55, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x12, 0x1f, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x20, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x73, 0x0a, 0x18, 0x41, 0x64, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x29, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7c, 0x0a, 0x1b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x2c, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x74, 0x6b, 0x6f, 0x68, 0x75, 0x74, 0x2f, 0x67, 0x6f,
	0x2d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_payments_proto_rawDescOnce sync.Once
	file_proto_payments_proto_rawDescData = file_proto_payments_proto_rawDesc
)

func file_proto_payments_proto_rawDescGZIP() []byte {
	file_proto_payments_proto_rawDescOnce.Do(func() {
		file_proto_payments_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_payments_proto_rawDescData)
	})
	return file_proto_payments_proto_rawDescData
}

var file_proto_payments_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_payments_proto_goTypes = []interface{}{
	(*Customer)(nil),                            // 0: payments.Customer
	(*Card)(nil),                                // 1: payments.Card
	(*CreateCustomerRequest)(nil),               // 2: payments.CreateCustomerRequest
	(*CreateCustomerResponse)(nil),              // 3: payments.CreateCustomerResponse
	(*GetCustomerByIdRequest)(nil),              // 4: payments.GetCustomerByIdRequest
	(*GetCustomerByIdResponse)(nil),             // 5: payments.GetCustomerByIdResponse
	(*AddCustomerPaymentMethodRequest)(nil),     // 6: payments.AddCustomerPaymentMethodRequest
	(*AddCustomerPaymentMethodResponse)(nil),    // 7: payments.AddCustomerPaymentMethodResponse
	(*RemoveCustomerPaymentMethodRequest)(nil),  // 8: payments.RemoveCustomerPaymentMethodRequest
	(*RemoveCustomerPaymentMethodResponse)(nil), // 9: payments.RemoveCustomerPaymentMethodResponse
}
var file_proto_payments_proto_depIdxs = []int32{
	1, // 0: payments.Customer.cards:type_name -> payments.Card
	0, // 1: payments.CreateCustomerResponse.customer:type_name -> payments.Customer
	0, // 2: payments.GetCustomerByIdResponse.customer:type_name -> payments.Customer
	1, // 3: payments.AddCustomerPaymentMethodRequest.card:type_name -> payments.Card
	2, // 4: payments.PaymentService.CreateCustomer:input_type -> payments.CreateCustomerRequest
	4, // 5: payments.PaymentService.GetCustomerById:input_type -> payments.GetCustomerByIdRequest
	6, // 6: payments.PaymentService.AddCustomerPaymentMethod:input_type -> payments.AddCustomerPaymentMethodRequest
	8, // 7: payments.PaymentService.RemoveCustomerPaymentMethod:input_type -> payments.RemoveCustomerPaymentMethodRequest
	3, // 8: payments.PaymentService.CreateCustomer:output_type -> payments.CreateCustomerResponse
	5, // 9: payments.PaymentService.GetCustomerById:output_type -> payments.GetCustomerByIdResponse
	7, // 10: payments.PaymentService.AddCustomerPaymentMethod:output_type -> payments.AddCustomerPaymentMethodResponse
	9, // 11: payments.PaymentService.RemoveCustomerPaymentMethod:output_type -> payments.RemoveCustomerPaymentMethodResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_payments_proto_init() }
func file_proto_payments_proto_init() {
	if File_proto_payments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_payments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_payments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_payments_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCustomerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_payments_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCustomerResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_payments_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerByIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_payments_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerByIdResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_payments_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCustomerPaymentMethodRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_payments_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCustomerPaymentMethodResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_payments_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveCustomerPaymentMethodRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_payments_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveCustomerPaymentMethodResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_payments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_payments_proto_goTypes,
		DependencyIndexes: file_proto_payments_proto_depIdxs,
		MessageInfos:      file_proto_payments_proto_msgTypes,
	}.Build()
	File_proto_payments_proto = out.File
	file_proto_payments_proto_rawDesc = nil
	file_proto_payments_proto_goTypes = nil
	file_proto_payments_proto_depIdxs = nil
}
