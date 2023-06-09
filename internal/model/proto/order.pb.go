// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: internal/model/proto/order.proto

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

type UpdateDeliveryStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderCode    string `protobuf:"bytes,1,opt,name=order_code,json=orderCode,proto3" json:"order_code,omitempty"`
	TrackingCode string `protobuf:"bytes,2,opt,name=tracking_code,json=trackingCode,proto3" json:"tracking_code,omitempty"`
	Status       string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	UpdatedAt    int64  `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UpdateDeliveryStatusReq) Reset() {
	*x = UpdateDeliveryStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_model_proto_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeliveryStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeliveryStatusReq) ProtoMessage() {}

func (x *UpdateDeliveryStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_model_proto_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeliveryStatusReq.ProtoReflect.Descriptor instead.
func (*UpdateDeliveryStatusReq) Descriptor() ([]byte, []int) {
	return file_internal_model_proto_order_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateDeliveryStatusReq) GetOrderCode() string {
	if x != nil {
		return x.OrderCode
	}
	return ""
}

func (x *UpdateDeliveryStatusReq) GetTrackingCode() string {
	if x != nil {
		return x.TrackingCode
	}
	return ""
}

func (x *UpdateDeliveryStatusReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdateDeliveryStatusReq) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type UpdateDeliveryStatusRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UpdateDeliveryStatusRes) Reset() {
	*x = UpdateDeliveryStatusRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_model_proto_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeliveryStatusRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeliveryStatusRes) ProtoMessage() {}

func (x *UpdateDeliveryStatusRes) ProtoReflect() protoreflect.Message {
	mi := &file_internal_model_proto_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeliveryStatusRes.ProtoReflect.Descriptor instead.
func (*UpdateDeliveryStatusRes) Descriptor() ([]byte, []int) {
	return file_internal_model_proto_order_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateDeliveryStatusRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_internal_model_proto_order_proto protoreflect.FileDescriptor

var file_internal_model_proto_order_proto_rawDesc = []byte{
	0x0a, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1d,
	0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x33, 0x0a, 0x17, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x55,
	0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_model_proto_order_proto_rawDescOnce sync.Once
	file_internal_model_proto_order_proto_rawDescData = file_internal_model_proto_order_proto_rawDesc
)

func file_internal_model_proto_order_proto_rawDescGZIP() []byte {
	file_internal_model_proto_order_proto_rawDescOnce.Do(func() {
		file_internal_model_proto_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_model_proto_order_proto_rawDescData)
	})
	return file_internal_model_proto_order_proto_rawDescData
}

var file_internal_model_proto_order_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_model_proto_order_proto_goTypes = []interface{}{
	(*UpdateDeliveryStatusReq)(nil), // 0: UpdateDeliveryStatusReq
	(*UpdateDeliveryStatusRes)(nil), // 1: UpdateDeliveryStatusRes
}
var file_internal_model_proto_order_proto_depIdxs = []int32{
	0, // 0: Order.UpdateDeliveryStatus:input_type -> UpdateDeliveryStatusReq
	1, // 1: Order.UpdateDeliveryStatus:output_type -> UpdateDeliveryStatusRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_model_proto_order_proto_init() }
func file_internal_model_proto_order_proto_init() {
	if File_internal_model_proto_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_model_proto_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeliveryStatusReq); i {
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
		file_internal_model_proto_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeliveryStatusRes); i {
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
			RawDescriptor: file_internal_model_proto_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_model_proto_order_proto_goTypes,
		DependencyIndexes: file_internal_model_proto_order_proto_depIdxs,
		MessageInfos:      file_internal_model_proto_order_proto_msgTypes,
	}.Build()
	File_internal_model_proto_order_proto = out.File
	file_internal_model_proto_order_proto_rawDesc = nil
	file_internal_model_proto_order_proto_goTypes = nil
	file_internal_model_proto_order_proto_depIdxs = nil
}
