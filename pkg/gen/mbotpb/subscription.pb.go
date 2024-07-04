// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: subscription.proto

package mbotpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionId string                 `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	CustomerId     string                 `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	StartDate      *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	ExpirationDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{0}
}

func (x *Subscription) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *Subscription) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Subscription) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *Subscription) GetExpirationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpirationDate
	}
	return nil
}

type CreateSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId     string                 `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	StartDate      *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	ExpirationDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
}

func (x *CreateSubscriptionRequest) Reset() {
	*x = CreateSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubscriptionRequest) ProtoMessage() {}

func (x *CreateSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*CreateSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSubscriptionRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CreateSubscriptionRequest) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *CreateSubscriptionRequest) GetExpirationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpirationDate
	}
	return nil
}

type CreateSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message      string        `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Subscription *Subscription `protobuf:"bytes,3,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *CreateSubscriptionResponse) Reset() {
	*x = CreateSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubscriptionResponse) ProtoMessage() {}

func (x *CreateSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*CreateSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSubscriptionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateSubscriptionResponse) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

type GetSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionId string `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
}

func (x *GetSubscriptionRequest) Reset() {
	*x = GetSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionRequest) ProtoMessage() {}

func (x *GetSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*GetSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{3}
}

func (x *GetSubscriptionRequest) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

type GetSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscription *Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *GetSubscriptionResponse) Reset() {
	*x = GetSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionResponse) ProtoMessage() {}

func (x *GetSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*GetSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{4}
}

func (x *GetSubscriptionResponse) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

type UpdateSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StartDate      *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	ExpirationDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
}

func (x *UpdateSubscriptionRequest) Reset() {
	*x = UpdateSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSubscriptionRequest) ProtoMessage() {}

func (x *UpdateSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*UpdateSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateSubscriptionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSubscriptionRequest) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *UpdateSubscriptionRequest) GetExpirationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpirationDate
	}
	return nil
}

type UpdateSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StartDate             *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	ExpirationDate        *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
	UpdatedStartDate      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_start_date,json=updatedStartDate,proto3" json:"updated_start_date,omitempty"`
	UpdatedExpirationDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_expiration_date,json=updatedExpirationDate,proto3" json:"updated_expiration_date,omitempty"`
}

func (x *UpdateSubscriptionResponse) Reset() {
	*x = UpdateSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSubscriptionResponse) ProtoMessage() {}

func (x *UpdateSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*UpdateSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateSubscriptionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSubscriptionResponse) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *UpdateSubscriptionResponse) GetExpirationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpirationDate
	}
	return nil
}

func (x *UpdateSubscriptionResponse) GetUpdatedStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedStartDate
	}
	return nil
}

func (x *UpdateSubscriptionResponse) GetUpdatedExpirationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedExpirationDate
	}
	return nil
}

type DeleteSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionId string `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
}

func (x *DeleteSubscriptionRequest) Reset() {
	*x = DeleteSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSubscriptionRequest) ProtoMessage() {}

func (x *DeleteSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*DeleteSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteSubscriptionRequest) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

type DeleteSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deleted bool `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeleteSubscriptionResponse) Reset() {
	*x = DeleteSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSubscriptionResponse) ProtoMessage() {}

func (x *DeleteSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*DeleteSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteSubscriptionResponse) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type GetSubscriptionsAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSubscriptionsAllRequest) Reset() {
	*x = GetSubscriptionsAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionsAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionsAllRequest) ProtoMessage() {}

func (x *GetSubscriptionsAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionsAllRequest.ProtoReflect.Descriptor instead.
func (*GetSubscriptionsAllRequest) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{9}
}

type GetSubscriptionsAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriptions []*Subscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
}

func (x *GetSubscriptionsAllResponse) Reset() {
	*x = GetSubscriptionsAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionsAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionsAllResponse) ProtoMessage() {}

func (x *GetSubscriptionsAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionsAllResponse.ProtoReflect.Descriptor instead.
func (*GetSubscriptionsAllResponse) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{10}
}

func (x *GetSubscriptionsAllResponse) GetSubscriptions() []*Subscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

type GetSubscriptionByCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *GetSubscriptionByCustomerRequest) Reset() {
	*x = GetSubscriptionByCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionByCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionByCustomerRequest) ProtoMessage() {}

func (x *GetSubscriptionByCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionByCustomerRequest.ProtoReflect.Descriptor instead.
func (*GetSubscriptionByCustomerRequest) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{11}
}

func (x *GetSubscriptionByCustomerRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type GetSubscriptionByCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriptions []*Subscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
}

func (x *GetSubscriptionByCustomerResponse) Reset() {
	*x = GetSubscriptionByCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionByCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionByCustomerResponse) ProtoMessage() {}

func (x *GetSubscriptionByCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionByCustomerResponse.ProtoReflect.Descriptor instead.
func (*GetSubscriptionByCustomerResponse) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{12}
}

func (x *GetSubscriptionByCustomerResponse) GetSubscriptions() []*Subscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

var File_subscription_proto protoreflect.FileDescriptor

var file_subscription_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x62, 0x6f, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x01, 0x0a, 0x0c,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x43, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0xbc, 0x01, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x43, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0x6e, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a,
	0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x62, 0x6f, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x41, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x62, 0x6f, 0x74,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xab, 0x01, 0x0a, 0x19,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0xca, 0x02, 0x0a, 0x1a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x48, 0x0a, 0x12, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x10, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x52, 0x0a, 0x17, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x15, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0x44, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x1a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x22, 0x1c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x57, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x38, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x62, 0x6f, 0x74, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x43, 0x0a, 0x20, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x5d, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x79, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d,
	0x62, 0x6f, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x75, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x62, 0x6f, 0x74, 0x42, 0x11, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6b, 0x75,
	0x70, 0x72, 0x6f, 0x76, 0x2f, 0x6d, 0x62, 0x6f, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x6d, 0x62, 0x6f, 0x74, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02,
	0x04, 0x4d, 0x62, 0x6f, 0x74, 0xca, 0x02, 0x04, 0x4d, 0x62, 0x6f, 0x74, 0xe2, 0x02, 0x10, 0x4d,
	0x62, 0x6f, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x04, 0x4d, 0x62, 0x6f, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subscription_proto_rawDescOnce sync.Once
	file_subscription_proto_rawDescData = file_subscription_proto_rawDesc
)

func file_subscription_proto_rawDescGZIP() []byte {
	file_subscription_proto_rawDescOnce.Do(func() {
		file_subscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_subscription_proto_rawDescData)
	})
	return file_subscription_proto_rawDescData
}

var file_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_subscription_proto_goTypes = []any{
	(*Subscription)(nil),                      // 0: mbot.Subscription
	(*CreateSubscriptionRequest)(nil),         // 1: mbot.CreateSubscriptionRequest
	(*CreateSubscriptionResponse)(nil),        // 2: mbot.CreateSubscriptionResponse
	(*GetSubscriptionRequest)(nil),            // 3: mbot.GetSubscriptionRequest
	(*GetSubscriptionResponse)(nil),           // 4: mbot.GetSubscriptionResponse
	(*UpdateSubscriptionRequest)(nil),         // 5: mbot.UpdateSubscriptionRequest
	(*UpdateSubscriptionResponse)(nil),        // 6: mbot.UpdateSubscriptionResponse
	(*DeleteSubscriptionRequest)(nil),         // 7: mbot.DeleteSubscriptionRequest
	(*DeleteSubscriptionResponse)(nil),        // 8: mbot.DeleteSubscriptionResponse
	(*GetSubscriptionsAllRequest)(nil),        // 9: mbot.GetSubscriptionsAllRequest
	(*GetSubscriptionsAllResponse)(nil),       // 10: mbot.GetSubscriptionsAllResponse
	(*GetSubscriptionByCustomerRequest)(nil),  // 11: mbot.GetSubscriptionByCustomerRequest
	(*GetSubscriptionByCustomerResponse)(nil), // 12: mbot.GetSubscriptionByCustomerResponse
	(*timestamppb.Timestamp)(nil),             // 13: google.protobuf.Timestamp
}
var file_subscription_proto_depIdxs = []int32{
	13, // 0: mbot.Subscription.start_date:type_name -> google.protobuf.Timestamp
	13, // 1: mbot.Subscription.expiration_date:type_name -> google.protobuf.Timestamp
	13, // 2: mbot.CreateSubscriptionRequest.start_date:type_name -> google.protobuf.Timestamp
	13, // 3: mbot.CreateSubscriptionRequest.expiration_date:type_name -> google.protobuf.Timestamp
	0,  // 4: mbot.CreateSubscriptionResponse.subscription:type_name -> mbot.Subscription
	0,  // 5: mbot.GetSubscriptionResponse.subscription:type_name -> mbot.Subscription
	13, // 6: mbot.UpdateSubscriptionRequest.start_date:type_name -> google.protobuf.Timestamp
	13, // 7: mbot.UpdateSubscriptionRequest.expiration_date:type_name -> google.protobuf.Timestamp
	13, // 8: mbot.UpdateSubscriptionResponse.start_date:type_name -> google.protobuf.Timestamp
	13, // 9: mbot.UpdateSubscriptionResponse.expiration_date:type_name -> google.protobuf.Timestamp
	13, // 10: mbot.UpdateSubscriptionResponse.updated_start_date:type_name -> google.protobuf.Timestamp
	13, // 11: mbot.UpdateSubscriptionResponse.updated_expiration_date:type_name -> google.protobuf.Timestamp
	0,  // 12: mbot.GetSubscriptionsAllResponse.subscriptions:type_name -> mbot.Subscription
	0,  // 13: mbot.GetSubscriptionByCustomerResponse.subscriptions:type_name -> mbot.Subscription
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_subscription_proto_init() }
func file_subscription_proto_init() {
	if File_subscription_proto != nil {
		return
	}
	file_customer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_subscription_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Subscription); i {
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
		file_subscription_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateSubscriptionRequest); i {
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
		file_subscription_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateSubscriptionResponse); i {
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
		file_subscription_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetSubscriptionRequest); i {
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
		file_subscription_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetSubscriptionResponse); i {
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
		file_subscription_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateSubscriptionRequest); i {
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
		file_subscription_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateSubscriptionResponse); i {
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
		file_subscription_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteSubscriptionRequest); i {
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
		file_subscription_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteSubscriptionResponse); i {
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
		file_subscription_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetSubscriptionsAllRequest); i {
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
		file_subscription_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*GetSubscriptionsAllResponse); i {
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
		file_subscription_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*GetSubscriptionByCustomerRequest); i {
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
		file_subscription_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*GetSubscriptionByCustomerResponse); i {
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
			RawDescriptor: file_subscription_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_subscription_proto_goTypes,
		DependencyIndexes: file_subscription_proto_depIdxs,
		MessageInfos:      file_subscription_proto_msgTypes,
	}.Build()
	File_subscription_proto = out.File
	file_subscription_proto_rawDesc = nil
	file_subscription_proto_goTypes = nil
	file_subscription_proto_depIdxs = nil
}
