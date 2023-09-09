// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: ecommerce_proto/order_proto/update_order.proto

package ecommerce_proto

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

type UpdateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer_ID string  `protobuf:"bytes,1,opt,name=Customer_ID,json=CustomerID,proto3" json:"Customer_ID,omitempty"`
	Sku         string  `protobuf:"bytes,2,opt,name=Sku,proto3" json:"Sku,omitempty"`
	Quantity    float32 `protobuf:"fixed32,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"` // repeated Shipping Shipping=3;
}

func (x *UpdateOrderRequest) Reset() {
	*x = UpdateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecommerce_proto_order_proto_update_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderRequest) ProtoMessage() {}

func (x *UpdateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_proto_order_proto_update_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrderRequest) Descriptor() ([]byte, []int) {
	return file_ecommerce_proto_order_proto_update_order_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateOrderRequest) GetCustomer_ID() string {
	if x != nil {
		return x.Customer_ID
	}
	return ""
}

func (x *UpdateOrderRequest) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *UpdateOrderRequest) GetQuantity() float32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type UpdateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *UpdateOrderResponse) Reset() {
	*x = UpdateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecommerce_proto_order_proto_update_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderResponse) ProtoMessage() {}

func (x *UpdateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_proto_order_proto_update_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderResponse.ProtoReflect.Descriptor instead.
func (*UpdateOrderResponse) Descriptor() ([]byte, []int) {
	return file_ecommerce_proto_order_proto_update_order_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateOrderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_ecommerce_proto_order_proto_update_order_proto protoreflect.FileDescriptor

var file_ecommerce_proto_order_proto_update_order_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x6b, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x53, 0x6b, 0x75, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x22, 0x2d, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x53, 0x75, 0x72, 0x65, 0x6e, 0x64, 0x68, 0x61, 0x72, 0x48, 0x4b, 0x2f, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ecommerce_proto_order_proto_update_order_proto_rawDescOnce sync.Once
	file_ecommerce_proto_order_proto_update_order_proto_rawDescData = file_ecommerce_proto_order_proto_update_order_proto_rawDesc
)

func file_ecommerce_proto_order_proto_update_order_proto_rawDescGZIP() []byte {
	file_ecommerce_proto_order_proto_update_order_proto_rawDescOnce.Do(func() {
		file_ecommerce_proto_order_proto_update_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_ecommerce_proto_order_proto_update_order_proto_rawDescData)
	})
	return file_ecommerce_proto_order_proto_update_order_proto_rawDescData
}

var file_ecommerce_proto_order_proto_update_order_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ecommerce_proto_order_proto_update_order_proto_goTypes = []interface{}{
	(*UpdateOrderRequest)(nil),  // 0: order_proto.UpdateOrderRequest
	(*UpdateOrderResponse)(nil), // 1: order_proto.UpdateOrderResponse
}
var file_ecommerce_proto_order_proto_update_order_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ecommerce_proto_order_proto_update_order_proto_init() }
func file_ecommerce_proto_order_proto_update_order_proto_init() {
	if File_ecommerce_proto_order_proto_update_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ecommerce_proto_order_proto_update_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderRequest); i {
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
		file_ecommerce_proto_order_proto_update_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderResponse); i {
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
			RawDescriptor: file_ecommerce_proto_order_proto_update_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ecommerce_proto_order_proto_update_order_proto_goTypes,
		DependencyIndexes: file_ecommerce_proto_order_proto_update_order_proto_depIdxs,
		MessageInfos:      file_ecommerce_proto_order_proto_update_order_proto_msgTypes,
	}.Build()
	File_ecommerce_proto_order_proto_update_order_proto = out.File
	file_ecommerce_proto_order_proto_update_order_proto_rawDesc = nil
	file_ecommerce_proto_order_proto_update_order_proto_goTypes = nil
	file_ecommerce_proto_order_proto_update_order_proto_depIdxs = nil
}
