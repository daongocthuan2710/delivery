// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: internal/model/proto/delivery.proto

package proto

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

type EstimateFeeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderValue int64         `protobuf:"varint,1,opt,name=order_value,json=orderValue,proto3" json:"order_value,omitempty"`
	Cod        int64         `protobuf:"varint,2,opt,name=cod,proto3" json:"cod,omitempty"`
	Weight     int64         `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	From       *DeliveryInfo `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
	To         *DeliveryInfo `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *EstimateFeeReq) Reset() {
	*x = EstimateFeeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_model_proto_delivery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EstimateFeeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstimateFeeReq) ProtoMessage() {}

func (x *EstimateFeeReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_model_proto_delivery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstimateFeeReq.ProtoReflect.Descriptor instead.
func (*EstimateFeeReq) Descriptor() ([]byte, []int) {
	return file_internal_model_proto_delivery_proto_rawDescGZIP(), []int{0}
}

func (x *EstimateFeeReq) GetOrderValue() int64 {
	if x != nil {
		return x.OrderValue
	}
	return 0
}

func (x *EstimateFeeReq) GetCod() int64 {
	if x != nil {
		return x.Cod
	}
	return 0
}

func (x *EstimateFeeReq) GetWeight() int64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *EstimateFeeReq) GetFrom() *DeliveryInfo {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *EstimateFeeReq) GetTo() *DeliveryInfo {
	if x != nil {
		return x.To
	}
	return nil
}

type EstimateFeeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*DeliveryService `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *EstimateFeeRes) Reset() {
	*x = EstimateFeeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_model_proto_delivery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EstimateFeeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstimateFeeRes) ProtoMessage() {}

func (x *EstimateFeeRes) ProtoReflect() protoreflect.Message {
	mi := &file_internal_model_proto_delivery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstimateFeeRes.ProtoReflect.Descriptor instead.
func (*EstimateFeeRes) Descriptor() ([]byte, []int) {
	return file_internal_model_proto_delivery_proto_rawDescGZIP(), []int{1}
}

func (x *EstimateFeeRes) GetServices() []*DeliveryService {
	if x != nil {
		return x.Services
	}
	return nil
}

type DeliveryService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code         string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	TplCode      string `protobuf:"bytes,3,opt,name=tpl_code,json=tplCode,proto3" json:"tpl_code,omitempty"`
	TotalFee     int64  `protobuf:"varint,4,opt,name=total_fee,json=totalFee,proto3" json:"total_fee,omitempty"`
	EstimateTime string `protobuf:"bytes,5,opt,name=estimate_time,json=estimateTime,proto3" json:"estimate_time,omitempty"`
}

func (x *DeliveryService) Reset() {
	*x = DeliveryService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_model_proto_delivery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryService) ProtoMessage() {}

func (x *DeliveryService) ProtoReflect() protoreflect.Message {
	mi := &file_internal_model_proto_delivery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryService.ProtoReflect.Descriptor instead.
func (*DeliveryService) Descriptor() ([]byte, []int) {
	return file_internal_model_proto_delivery_proto_rawDescGZIP(), []int{2}
}

func (x *DeliveryService) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeliveryService) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DeliveryService) GetTplCode() string {
	if x != nil {
		return x.TplCode
	}
	return ""
}

func (x *DeliveryService) GetTotalFee() int64 {
	if x != nil {
		return x.TotalFee
	}
	return 0
}

func (x *DeliveryService) GetEstimateTime() string {
	if x != nil {
		return x.EstimateTime
	}
	return ""
}

type CreateDeliveryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderCode   string        `protobuf:"bytes,1,opt,name=order_code,json=orderCode,proto3" json:"order_code,omitempty"`
	Note        string        `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
	OrderValue  int64         `protobuf:"varint,3,opt,name=order_value,json=orderValue,proto3" json:"order_value,omitempty"`
	Cod         int64         `protobuf:"varint,4,opt,name=cod,proto3" json:"cod,omitempty"`
	Weight      int64         `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	ServiceCode string        `protobuf:"bytes,6,opt,name=service_code,json=serviceCode,proto3" json:"service_code,omitempty"`
	From        *DeliveryInfo `protobuf:"bytes,7,opt,name=from,proto3" json:"from,omitempty"`
	To          *DeliveryInfo `protobuf:"bytes,8,opt,name=to,proto3" json:"to,omitempty"`
	Items       []*Item       `protobuf:"bytes,9,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CreateDeliveryReq) Reset() {
	*x = CreateDeliveryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_model_proto_delivery_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeliveryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeliveryReq) ProtoMessage() {}

func (x *CreateDeliveryReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_model_proto_delivery_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeliveryReq.ProtoReflect.Descriptor instead.
func (*CreateDeliveryReq) Descriptor() ([]byte, []int) {
	return file_internal_model_proto_delivery_proto_rawDescGZIP(), []int{3}
}

func (x *CreateDeliveryReq) GetOrderCode() string {
	if x != nil {
		return x.OrderCode
	}
	return ""
}

func (x *CreateDeliveryReq) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *CreateDeliveryReq) GetOrderValue() int64 {
	if x != nil {
		return x.OrderValue
	}
	return 0
}

func (x *CreateDeliveryReq) GetCod() int64 {
	if x != nil {
		return x.Cod
	}
	return 0
}

func (x *CreateDeliveryReq) GetWeight() int64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *CreateDeliveryReq) GetServiceCode() string {
	if x != nil {
		return x.ServiceCode
	}
	return ""
}

func (x *CreateDeliveryReq) GetFrom() *DeliveryInfo {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *CreateDeliveryReq) GetTo() *DeliveryInfo {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *CreateDeliveryReq) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateDeliveryRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderCode    string `protobuf:"bytes,1,opt,name=order_code,json=orderCode,proto3" json:"order_code,omitempty"`
	TrackingCode string `protobuf:"bytes,2,opt,name=tracking_code,json=trackingCode,proto3" json:"tracking_code,omitempty"`
	Status       string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	TotalFee     int64  `protobuf:"varint,4,opt,name=total_fee,json=totalFee,proto3" json:"total_fee,omitempty"`
}

func (x *CreateDeliveryRes) Reset() {
	*x = CreateDeliveryRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_model_proto_delivery_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeliveryRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeliveryRes) ProtoMessage() {}

func (x *CreateDeliveryRes) ProtoReflect() protoreflect.Message {
	mi := &file_internal_model_proto_delivery_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeliveryRes.ProtoReflect.Descriptor instead.
func (*CreateDeliveryRes) Descriptor() ([]byte, []int) {
	return file_internal_model_proto_delivery_proto_rawDescGZIP(), []int{4}
}

func (x *CreateDeliveryRes) GetOrderCode() string {
	if x != nil {
		return x.OrderCode
	}
	return ""
}

func (x *CreateDeliveryRes) GetTrackingCode() string {
	if x != nil {
		return x.TrackingCode
	}
	return ""
}

func (x *CreateDeliveryRes) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateDeliveryRes) GetTotalFee() int64 {
	if x != nil {
		return x.TotalFee
	}
	return 0
}

type DeliveryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone      string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Address    string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	ProvinceId int32  `protobuf:"varint,4,opt,name=province_id,json=provinceId,proto3" json:"province_id,omitempty"`
	DistrictId int32  `protobuf:"varint,5,opt,name=district_id,json=districtId,proto3" json:"district_id,omitempty"`
	WardId     int32  `protobuf:"varint,6,opt,name=ward_id,json=wardId,proto3" json:"ward_id,omitempty"`
}

func (x *DeliveryInfo) Reset() {
	*x = DeliveryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_model_proto_delivery_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryInfo) ProtoMessage() {}

func (x *DeliveryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_internal_model_proto_delivery_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryInfo.ProtoReflect.Descriptor instead.
func (*DeliveryInfo) Descriptor() ([]byte, []int) {
	return file_internal_model_proto_delivery_proto_rawDescGZIP(), []int{5}
}

func (x *DeliveryInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeliveryInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *DeliveryInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DeliveryInfo) GetProvinceId() int32 {
	if x != nil {
		return x.ProvinceId
	}
	return 0
}

func (x *DeliveryInfo) GetDistrictId() int32 {
	if x != nil {
		return x.DistrictId
	}
	return 0
}

func (x *DeliveryInfo) GetWardId() int32 {
	if x != nil {
		return x.WardId
	}
	return 0
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Quantity int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Weight   int64  `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_model_proto_delivery_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_internal_model_proto_delivery_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_internal_model_proto_delivery_proto_rawDescGZIP(), []int{6}
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Item) GetWeight() int64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

var File_internal_model_proto_delivery_proto protoreflect.FileDescriptor

var file_internal_model_proto_delivery_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x0e, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61,
	0x74, 0x65, 0x46, 0x65, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6f, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1d, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x0e, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x65, 0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x96, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x70, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x70, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x65, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x73, 0x74,
	0x69, 0x6d, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x93,
	0x02, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6f, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1d, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x66, 0x65, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x46, 0x65, 0x65, 0x22, 0xad, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x61,
	0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x61, 0x72,
	0x64, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x32, 0x79, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12,
	0x3a, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0b, 0x45,
	0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x12, 0x0f, 0x2e, 0x45, 0x73, 0x74,
	0x69, 0x6d, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x45, 0x73,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x1f,
	0x5a, 0x1d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_model_proto_delivery_proto_rawDescOnce sync.Once
	file_internal_model_proto_delivery_proto_rawDescData = file_internal_model_proto_delivery_proto_rawDesc
)

func file_internal_model_proto_delivery_proto_rawDescGZIP() []byte {
	file_internal_model_proto_delivery_proto_rawDescOnce.Do(func() {
		file_internal_model_proto_delivery_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_model_proto_delivery_proto_rawDescData)
	})
	return file_internal_model_proto_delivery_proto_rawDescData
}

var file_internal_model_proto_delivery_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internal_model_proto_delivery_proto_goTypes = []interface{}{
	(*EstimateFeeReq)(nil),    // 0: EstimateFeeReq
	(*EstimateFeeRes)(nil),    // 1: EstimateFeeRes
	(*DeliveryService)(nil),   // 2: DeliveryService
	(*CreateDeliveryReq)(nil), // 3: CreateDeliveryReq
	(*CreateDeliveryRes)(nil), // 4: CreateDeliveryRes
	(*DeliveryInfo)(nil),      // 5: DeliveryInfo
	(*Item)(nil),              // 6: Item
}
var file_internal_model_proto_delivery_proto_depIdxs = []int32{
	5, // 0: EstimateFeeReq.from:type_name -> DeliveryInfo
	5, // 1: EstimateFeeReq.to:type_name -> DeliveryInfo
	2, // 2: EstimateFeeRes.services:type_name -> DeliveryService
	5, // 3: CreateDeliveryReq.from:type_name -> DeliveryInfo
	5, // 4: CreateDeliveryReq.to:type_name -> DeliveryInfo
	6, // 5: CreateDeliveryReq.items:type_name -> Item
	3, // 6: Delivery.CreateDelivery:input_type -> CreateDeliveryReq
	0, // 7: Delivery.EstimateFee:input_type -> EstimateFeeReq
	4, // 8: Delivery.CreateDelivery:output_type -> CreateDeliveryRes
	1, // 9: Delivery.EstimateFee:output_type -> EstimateFeeRes
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_internal_model_proto_delivery_proto_init() }
func file_internal_model_proto_delivery_proto_init() {
	if File_internal_model_proto_delivery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_model_proto_delivery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EstimateFeeReq); i {
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
		file_internal_model_proto_delivery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EstimateFeeRes); i {
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
		file_internal_model_proto_delivery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryService); i {
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
		file_internal_model_proto_delivery_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeliveryReq); i {
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
		file_internal_model_proto_delivery_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeliveryRes); i {
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
		file_internal_model_proto_delivery_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryInfo); i {
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
		file_internal_model_proto_delivery_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
			RawDescriptor: file_internal_model_proto_delivery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_model_proto_delivery_proto_goTypes,
		DependencyIndexes: file_internal_model_proto_delivery_proto_depIdxs,
		MessageInfos:      file_internal_model_proto_delivery_proto_msgTypes,
	}.Build()
	File_internal_model_proto_delivery_proto = out.File
	file_internal_model_proto_delivery_proto_rawDesc = nil
	file_internal_model_proto_delivery_proto_goTypes = nil
	file_internal_model_proto_delivery_proto_depIdxs = nil
}
