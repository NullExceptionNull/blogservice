// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: order.proto

package opb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type OrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source    string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	SourceId  string `protobuf:"bytes,2,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	Mobile    string `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	ShareNum  string `protobuf:"bytes,4,opt,name=share_num,json=shareNum,proto3" json:"share_num,omitempty"`
	SendMsn   string `protobuf:"bytes,5,opt,name=send_msn,json=sendMsn,proto3" json:"send_msn,omitempty"`
	OrderInfo string `protobuf:"bytes,6,opt,name=order_info,json=orderInfo,proto3" json:"order_info,omitempty"`
}

func (x *OrderReq) Reset() {
	*x = OrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderReq) ProtoMessage() {}

func (x *OrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderReq.ProtoReflect.Descriptor instead.
func (*OrderReq) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderReq) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *OrderReq) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *OrderReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *OrderReq) GetShareNum() string {
	if x != nil {
		return x.ShareNum
	}
	return ""
}

func (x *OrderReq) GetSendMsn() string {
	if x != nil {
		return x.SendMsn
	}
	return ""
}

func (x *OrderReq) GetOrderInfo() string {
	if x != nil {
		return x.OrderInfo
	}
	return ""
}

type OrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stat int32            `protobuf:"varint,1,opt,name=stat,proto3" json:"stat,omitempty"`
	Data *OrderResultInfo `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OrderResp) Reset() {
	*x = OrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResp) ProtoMessage() {}

func (x *OrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResp.ProtoReflect.Descriptor instead.
func (*OrderResp) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderResp) GetStat() int32 {
	if x != nil {
		return x.Stat
	}
	return 0
}

func (x *OrderResp) GetData() *OrderResultInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type OrderResultInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderNum     string `protobuf:"bytes,1,opt,name=order_num,json=orderNum,proto3" json:"order_num,omitempty"`
	UserId       string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IsExistOrder string `protobuf:"bytes,3,opt,name=is_exist_order,json=isExistOrder,proto3" json:"is_exist_order,omitempty"`
}

func (x *OrderResultInfo) Reset() {
	*x = OrderResultInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResultInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResultInfo) ProtoMessage() {}

func (x *OrderResultInfo) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResultInfo.ProtoReflect.Descriptor instead.
func (*OrderResultInfo) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderResultInfo) GetOrderNum() string {
	if x != nil {
		return x.OrderNum
	}
	return ""
}

func (x *OrderResultInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderResultInfo) GetIsExistOrder() string {
	if x != nil {
		return x.IsExistOrder
	}
	return ""
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x68, 0x61, 0x72, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x6e,
	0x64, 0x5f, 0x6d, 0x73, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x6e,
	0x64, 0x4d, 0x73, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x4b, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x73, 0x74, 0x61, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x6d, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x69, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x32,
	0x5c, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4c, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x6f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f,
	0x61, 0x75, 0x74, 0x6f, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x42, 0x10, 0x5a,
	0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x3b, 0x6f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_order_proto_goTypes = []interface{}{
	(*OrderReq)(nil),        // 0: order.OrderReq
	(*OrderResp)(nil),       // 1: order.OrderResp
	(*OrderResultInfo)(nil), // 2: order.OrderResultInfo
}
var file_order_proto_depIdxs = []int32{
	2, // 0: order.OrderResp.data:type_name -> order.OrderResultInfo
	0, // 1: order.OrderService.AutoOrder:input_type -> order.OrderReq
	1, // 2: order.OrderService.AutoOrder:output_type -> order.OrderResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderReq); i {
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
		file_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResp); i {
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
		file_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResultInfo); i {
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
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
