// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: proto/health/health.proto

package health

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN         HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING         HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING     HealthCheckResponse_ServingStatus = 2
	HealthCheckResponse_SERVICE_UNKNOWN HealthCheckResponse_ServingStatus = 3 // Used only by the Watch method.
)

// Enum value maps for HealthCheckResponse_ServingStatus.
var (
	HealthCheckResponse_ServingStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "SERVING",
		2: "NOT_SERVING",
		3: "SERVICE_UNKNOWN",
	}
	HealthCheckResponse_ServingStatus_value = map[string]int32{
		"UNKNOWN":         0,
		"SERVING":         1,
		"NOT_SERVING":     2,
		"SERVICE_UNKNOWN": 3,
	}
)

func (x HealthCheckResponse_ServingStatus) Enum() *HealthCheckResponse_ServingStatus {
	p := new(HealthCheckResponse_ServingStatus)
	*p = x
	return p
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthCheckResponse_ServingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_health_health_proto_enumTypes[0].Descriptor()
}

func (HealthCheckResponse_ServingStatus) Type() protoreflect.EnumType {
	return &file_proto_health_health_proto_enumTypes[0]
}

func (x HealthCheckResponse_ServingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthCheckResponse_ServingStatus.Descriptor instead.
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_health_health_proto_rawDescGZIP(), []int{1, 0}
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Service       string                 `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	mi := &file_proto_health_health_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_health_health_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_proto_health_health_proto_rawDescGZIP(), []int{0}
}

func (x *HealthCheckRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	Status        HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=health.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	mi := &file_proto_health_health_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_health_health_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_proto_health_health_proto_rawDescGZIP(), []int{1}
}

func (x *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if x != nil {
		return x.Status
	}
	return HealthCheckResponse_UNKNOWN
}

var File_proto_health_health_proto protoreflect.FileDescriptor

const file_proto_health_health_proto_rawDesc = "" +
	"\n" +
	"\x19proto/health/health.proto\x12\x06health\".\n" +
	"\x12HealthCheckRequest\x12\x18\n" +
	"\aservice\x18\x01 \x01(\tR\aservice\"\xa9\x01\n" +
	"\x13HealthCheckResponse\x12A\n" +
	"\x06status\x18\x01 \x01(\x0e2).health.HealthCheckResponse.ServingStatusR\x06status\"O\n" +
	"\rServingStatus\x12\v\n" +
	"\aUNKNOWN\x10\x00\x12\v\n" +
	"\aSERVING\x10\x01\x12\x0f\n" +
	"\vNOT_SERVING\x10\x02\x12\x13\n" +
	"\x0fSERVICE_UNKNOWN\x10\x03B5Z3github.com/eliezerraj/go-grpc-proto/protogen/healthb\x06proto3"

var (
	file_proto_health_health_proto_rawDescOnce sync.Once
	file_proto_health_health_proto_rawDescData []byte
)

func file_proto_health_health_proto_rawDescGZIP() []byte {
	file_proto_health_health_proto_rawDescOnce.Do(func() {
		file_proto_health_health_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_health_health_proto_rawDesc), len(file_proto_health_health_proto_rawDesc)))
	})
	return file_proto_health_health_proto_rawDescData
}

var file_proto_health_health_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_health_health_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_health_health_proto_goTypes = []any{
	(HealthCheckResponse_ServingStatus)(0), // 0: health.HealthCheckResponse.ServingStatus
	(*HealthCheckRequest)(nil),             // 1: health.HealthCheckRequest
	(*HealthCheckResponse)(nil),            // 2: health.HealthCheckResponse
}
var file_proto_health_health_proto_depIdxs = []int32{
	0, // 0: health.HealthCheckResponse.status:type_name -> health.HealthCheckResponse.ServingStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_health_health_proto_init() }
func file_proto_health_health_proto_init() {
	if File_proto_health_health_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_health_health_proto_rawDesc), len(file_proto_health_health_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_health_health_proto_goTypes,
		DependencyIndexes: file_proto_health_health_proto_depIdxs,
		EnumInfos:         file_proto_health_health_proto_enumTypes,
		MessageInfos:      file_proto_health_health_proto_msgTypes,
	}.Build()
	File_proto_health_health_proto = out.File
	file_proto_health_health_proto_goTypes = nil
	file_proto_health_health_proto_depIdxs = nil
}
