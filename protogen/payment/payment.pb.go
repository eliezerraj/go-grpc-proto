// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: proto/payment/payment.proto

package payment

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Payment struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountId     string                 `protobuf:"bytes,1,opt,name=account_id,proto3" json:"account_id,omitempty"`
	CardNumber    string                 `protobuf:"bytes,2,opt,name=card_number,proto3" json:"card_number,omitempty"`
	CardType      string                 `protobuf:"bytes,3,opt,name=card_type,proto3" json:"card_type,omitempty"`
	Currency      string                 `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount        float64                `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Status        string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	PaymentAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=payment_at,proto3" json:"payment_at,omitempty"`
	Mcc           string                 `protobuf:"bytes,8,opt,name=mcc,proto3" json:"mcc,omitempty"`
	CardModel     string                 `protobuf:"bytes,9,opt,name=card_model,proto3" json:"card_model,omitempty"`
	Terminal      string                 `protobuf:"bytes,10,opt,name=terminal,proto3" json:"terminal,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Payment) Reset() {
	*x = Payment{}
	mi := &file_proto_payment_payment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_payment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_proto_payment_payment_proto_rawDescGZIP(), []int{0}
}

func (x *Payment) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Payment) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *Payment) GetCardType() string {
	if x != nil {
		return x.CardType
	}
	return ""
}

func (x *Payment) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Payment) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Payment) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Payment) GetPaymentAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PaymentAt
	}
	return nil
}

func (x *Payment) GetMcc() string {
	if x != nil {
		return x.Mcc
	}
	return ""
}

func (x *Payment) GetCardModel() string {
	if x != nil {
		return x.CardModel
	}
	return ""
}

func (x *Payment) GetTerminal() string {
	if x != nil {
		return x.Terminal
	}
	return ""
}

type PaymentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Payment       *Payment               `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaymentRequest) Reset() {
	*x = PaymentRequest{}
	mi := &file_proto_payment_payment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentRequest) ProtoMessage() {}

func (x *PaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_payment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentRequest.ProtoReflect.Descriptor instead.
func (*PaymentRequest) Descriptor() ([]byte, []int) {
	return file_proto_payment_payment_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentRequest) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

type PaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Payment       *Payment               `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaymentResponse) Reset() {
	*x = PaymentResponse{}
	mi := &file_proto_payment_payment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentResponse) ProtoMessage() {}

func (x *PaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_payment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentResponse.ProtoReflect.Descriptor instead.
func (*PaymentResponse) Descriptor() ([]byte, []int) {
	return file_proto_payment_payment_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentResponse) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

var File_proto_payment_payment_proto protoreflect.FileDescriptor

const file_proto_payment_payment_proto_rawDesc = "" +
	"\n" +
	"\x1bproto/payment/payment.proto\x12\apayment\x1a\x1fgoogle/protobuf/timestamp.proto\"\xbf\x02\n" +
	"\aPayment\x12\x1e\n" +
	"\n" +
	"account_id\x18\x01 \x01(\tR\n" +
	"account_id\x12 \n" +
	"\vcard_number\x18\x02 \x01(\tR\vcard_number\x12\x1c\n" +
	"\tcard_type\x18\x03 \x01(\tR\tcard_type\x12\x1a\n" +
	"\bcurrency\x18\x04 \x01(\tR\bcurrency\x12\x16\n" +
	"\x06amount\x18\x05 \x01(\x01R\x06amount\x12\x16\n" +
	"\x06status\x18\x06 \x01(\tR\x06status\x12:\n" +
	"\n" +
	"payment_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"payment_at\x12\x10\n" +
	"\x03mcc\x18\b \x01(\tR\x03mcc\x12\x1e\n" +
	"\n" +
	"card_model\x18\t \x01(\tR\n" +
	"card_model\x12\x1a\n" +
	"\bterminal\x18\n" +
	" \x01(\tR\bterminal\"<\n" +
	"\x0ePaymentRequest\x12*\n" +
	"\apayment\x18\x01 \x01(\v2\x10.payment.PaymentR\apayment\"=\n" +
	"\x0fPaymentResponse\x12*\n" +
	"\apayment\x18\x01 \x01(\v2\x10.payment.PaymentR\apaymentB6Z4github.com/eliezerraj/go-grpc-proto/protogen/paymentb\x06proto3"

var (
	file_proto_payment_payment_proto_rawDescOnce sync.Once
	file_proto_payment_payment_proto_rawDescData []byte
)

func file_proto_payment_payment_proto_rawDescGZIP() []byte {
	file_proto_payment_payment_proto_rawDescOnce.Do(func() {
		file_proto_payment_payment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_payment_payment_proto_rawDesc), len(file_proto_payment_payment_proto_rawDesc)))
	})
	return file_proto_payment_payment_proto_rawDescData
}

var file_proto_payment_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_payment_payment_proto_goTypes = []any{
	(*Payment)(nil),               // 0: payment.Payment
	(*PaymentRequest)(nil),        // 1: payment.PaymentRequest
	(*PaymentResponse)(nil),       // 2: payment.PaymentResponse
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_proto_payment_payment_proto_depIdxs = []int32{
	3, // 0: payment.Payment.payment_at:type_name -> google.protobuf.Timestamp
	0, // 1: payment.PaymentRequest.payment:type_name -> payment.Payment
	0, // 2: payment.PaymentResponse.payment:type_name -> payment.Payment
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_payment_payment_proto_init() }
func file_proto_payment_payment_proto_init() {
	if File_proto_payment_payment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_payment_payment_proto_rawDesc), len(file_proto_payment_payment_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_payment_payment_proto_goTypes,
		DependencyIndexes: file_proto_payment_payment_proto_depIdxs,
		MessageInfos:      file_proto_payment_payment_proto_msgTypes,
	}.Build()
	File_proto_payment_payment_proto = out.File
	file_proto_payment_payment_proto_goTypes = nil
	file_proto_payment_payment_proto_depIdxs = nil
}
