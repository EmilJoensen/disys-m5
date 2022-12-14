// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: auction/auction.proto

package auction

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

type BidAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount int32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *BidAmount) Reset() {
	*x = BidAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_auction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidAmount) ProtoMessage() {}

func (x *BidAmount) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidAmount.ProtoReflect.Descriptor instead.
func (*BidAmount) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{0}
}

func (x *BidAmount) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BidAmount) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BidAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack string `protobuf:"bytes,1,opt,name=ack,proto3" json:"ack,omitempty"`
}

func (x *BidAck) Reset() {
	*x = BidAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_auction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidAck) ProtoMessage() {}

func (x *BidAck) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidAck.ProtoReflect.Descriptor instead.
func (*BidAck) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{1}
}

func (x *BidAck) GetAck() string {
	if x != nil {
		return x.Ack
	}
	return ""
}

type ResultOutcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Outcome   int32  `protobuf:"varint,2,opt,name=outcome,proto3" json:"outcome,omitempty"`
	Starttime int64  `protobuf:"varint,3,opt,name=starttime,proto3" json:"starttime,omitempty"`
}

func (x *ResultOutcome) Reset() {
	*x = ResultOutcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_auction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultOutcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultOutcome) ProtoMessage() {}

func (x *ResultOutcome) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultOutcome.ProtoReflect.Descriptor instead.
func (*ResultOutcome) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{2}
}

func (x *ResultOutcome) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ResultOutcome) GetOutcome() int32 {
	if x != nil {
		return x.Outcome
	}
	return 0
}

func (x *ResultOutcome) GetStarttime() int64 {
	if x != nil {
		return x.Starttime
	}
	return 0
}

type ResultVoid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResultVoid) Reset() {
	*x = ResultVoid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_auction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultVoid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultVoid) ProtoMessage() {}

func (x *ResultVoid) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultVoid.ProtoReflect.Descriptor instead.
func (*ResultVoid) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{3}
}

var File_auction_auction_proto protoreflect.FileDescriptor

var file_auction_auction_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x33, 0x0a, 0x09, 0x42, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1a, 0x0a, 0x06, 0x42, 0x69, 0x64, 0x41, 0x63, 0x6b, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x63,
	0x6b, 0x22, 0x5f, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4f, 0x75, 0x74, 0x63, 0x6f,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x75,
	0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x75, 0x74,
	0x63, 0x6f, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x74, 0x69,
	0x6d, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x56, 0x6f, 0x69, 0x64,
	0x32, 0x70, 0x0a, 0x07, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x03, 0x42,
	0x69, 0x64, 0x12, 0x12, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x69, 0x64,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x42, 0x69, 0x64, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x16, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65,
	0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x6d, 0x69, 0x6c, 0x4a, 0x6f, 0x65,
	0x6e, 0x73, 0x65, 0x6e, 0x2f, 0x64, 0x69, 0x73, 0x79, 0x73, 0x2d, 0x6d, 0x35, 0x3b, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auction_auction_proto_rawDescOnce sync.Once
	file_auction_auction_proto_rawDescData = file_auction_auction_proto_rawDesc
)

func file_auction_auction_proto_rawDescGZIP() []byte {
	file_auction_auction_proto_rawDescOnce.Do(func() {
		file_auction_auction_proto_rawDescData = protoimpl.X.CompressGZIP(file_auction_auction_proto_rawDescData)
	})
	return file_auction_auction_proto_rawDescData
}

var file_auction_auction_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_auction_auction_proto_goTypes = []interface{}{
	(*BidAmount)(nil),     // 0: auction.BidAmount
	(*BidAck)(nil),        // 1: auction.BidAck
	(*ResultOutcome)(nil), // 2: auction.ResultOutcome
	(*ResultVoid)(nil),    // 3: auction.ResultVoid
}
var file_auction_auction_proto_depIdxs = []int32{
	0, // 0: auction.Auction.Bid:input_type -> auction.BidAmount
	3, // 1: auction.Auction.Result:input_type -> auction.ResultVoid
	1, // 2: auction.Auction.Bid:output_type -> auction.BidAck
	2, // 3: auction.Auction.Result:output_type -> auction.ResultOutcome
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auction_auction_proto_init() }
func file_auction_auction_proto_init() {
	if File_auction_auction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auction_auction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidAmount); i {
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
		file_auction_auction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidAck); i {
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
		file_auction_auction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultOutcome); i {
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
		file_auction_auction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultVoid); i {
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
			RawDescriptor: file_auction_auction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auction_auction_proto_goTypes,
		DependencyIndexes: file_auction_auction_proto_depIdxs,
		MessageInfos:      file_auction_auction_proto_msgTypes,
	}.Build()
	File_auction_auction_proto = out.File
	file_auction_auction_proto_rawDesc = nil
	file_auction_auction_proto_goTypes = nil
	file_auction_auction_proto_depIdxs = nil
}
