// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: cbws/service_management/service_usage/v1alpha1/service.proto

package v1alpha1

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ServiceState int32

const (
	ServiceState_SERVICE_STATE_UNSPECIFIED ServiceState = 0
	// Service is currently not enabled in the consumer project
	ServiceState_SERVICE_STATE_DISABLED ServiceState = 1
	// Service is enabled in the consumer project and can be used
	ServiceState_SERVICE_STATE_ENABLED ServiceState = 2
)

// Enum value maps for ServiceState.
var (
	ServiceState_name = map[int32]string{
		0: "SERVICE_STATE_UNSPECIFIED",
		1: "SERVICE_STATE_DISABLED",
		2: "SERVICE_STATE_ENABLED",
	}
	ServiceState_value = map[string]int32{
		"SERVICE_STATE_UNSPECIFIED": 0,
		"SERVICE_STATE_DISABLED":    1,
		"SERVICE_STATE_ENABLED":     2,
	}
)

func (x ServiceState) Enum() *ServiceState {
	p := new(ServiceState)
	*p = x
	return p
}

func (x ServiceState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceState) Descriptor() protoreflect.EnumDescriptor {
	return file_cbws_service_management_service_usage_v1alpha1_service_proto_enumTypes[0].Descriptor()
}

func (ServiceState) Type() protoreflect.EnumType {
	return &file_cbws_service_management_service_usage_v1alpha1_service_proto_enumTypes[0]
}

func (x ServiceState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceState.Descriptor instead.
func (ServiceState) EnumDescriptor() ([]byte, []int) {
	return file_cbws_service_management_service_usage_v1alpha1_service_proto_rawDescGZIP(), []int{0}
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Service relative resource name, in the format of projects/test-project/services/k8s.cbws.xyz
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Relative resource name of the parent resource, for example projects/test-project
	Parent string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	// Overall information about the service, including service name (k8s.cbws.xyz) and title
	Config *ServiceConfig `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	// Whether the service is enabled on this project or not
	State ServiceState `protobuf:"varint,4,opt,name=state,proto3,enum=cbws.service_management.service_usage.v1alpha1.ServiceState" json:"state,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cbws_service_management_service_usage_v1alpha1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_cbws_service_management_service_usage_v1alpha1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_cbws_service_management_service_usage_v1alpha1_service_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *Service) GetConfig() *ServiceConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Service) GetState() ServiceState {
	if x != nil {
		return x.State
	}
	return ServiceState_SERVICE_STATE_UNSPECIFIED
}

type ServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the service, for example k8s.cbws.xyz
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The title of the product of the service, for example Kubernetes
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *ServiceConfig) Reset() {
	*x = ServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cbws_service_management_service_usage_v1alpha1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceConfig) ProtoMessage() {}

func (x *ServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cbws_service_management_service_usage_v1alpha1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceConfig.ProtoReflect.Descriptor instead.
func (*ServiceConfig) Descriptor() ([]byte, []int) {
	return file_cbws_service_management_service_usage_v1alpha1_service_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceConfig) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_cbws_service_management_service_usage_v1alpha1_service_proto protoreflect.FileDescriptor

var file_cbws_service_management_service_usage_v1alpha1_service_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x63, 0x62, 0x77, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e,
	0x63, 0x62, 0x77, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x02, 0x0a, 0x07, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x55, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x62, 0x77, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x52, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x63, 0x62, 0x77, 0x73, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x51, 0xea, 0x41,
	0x4e, 0x0a, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x63, 0x62, 0x77, 0x73, 0x2e, 0x78, 0x79, 0x7a, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x7d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x7d, 0x22,
	0x39, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x2a, 0x64, 0x0a, 0x0c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42,
	0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02,
	0x42, 0x87, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x62, 0x77, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x62, 0x77, 0x73, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x63, 0x62, 0x77, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1f,
	0x43, 0x62, 0x77, 0x73, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x5c, 0x47, 0x52, 0x50, 0x43, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2,
	0x02, 0x28, 0x43, 0x62, 0x77, 0x73, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x5c, 0x47, 0x52, 0x50, 0x43, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cbws_service_management_service_usage_v1alpha1_service_proto_rawDescOnce sync.Once
	file_cbws_service_management_service_usage_v1alpha1_service_proto_rawDescData = file_cbws_service_management_service_usage_v1alpha1_service_proto_rawDesc
)

func file_cbws_service_management_service_usage_v1alpha1_service_proto_rawDescGZIP() []byte {
	file_cbws_service_management_service_usage_v1alpha1_service_proto_rawDescOnce.Do(func() {
		file_cbws_service_management_service_usage_v1alpha1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cbws_service_management_service_usage_v1alpha1_service_proto_rawDescData)
	})
	return file_cbws_service_management_service_usage_v1alpha1_service_proto_rawDescData
}

var file_cbws_service_management_service_usage_v1alpha1_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cbws_service_management_service_usage_v1alpha1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cbws_service_management_service_usage_v1alpha1_service_proto_goTypes = []interface{}{
	(ServiceState)(0),     // 0: cbws.service_management.service_usage.v1alpha1.ServiceState
	(*Service)(nil),       // 1: cbws.service_management.service_usage.v1alpha1.Service
	(*ServiceConfig)(nil), // 2: cbws.service_management.service_usage.v1alpha1.ServiceConfig
}
var file_cbws_service_management_service_usage_v1alpha1_service_proto_depIdxs = []int32{
	2, // 0: cbws.service_management.service_usage.v1alpha1.Service.config:type_name -> cbws.service_management.service_usage.v1alpha1.ServiceConfig
	0, // 1: cbws.service_management.service_usage.v1alpha1.Service.state:type_name -> cbws.service_management.service_usage.v1alpha1.ServiceState
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cbws_service_management_service_usage_v1alpha1_service_proto_init() }
func file_cbws_service_management_service_usage_v1alpha1_service_proto_init() {
	if File_cbws_service_management_service_usage_v1alpha1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cbws_service_management_service_usage_v1alpha1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_cbws_service_management_service_usage_v1alpha1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceConfig); i {
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
			RawDescriptor: file_cbws_service_management_service_usage_v1alpha1_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cbws_service_management_service_usage_v1alpha1_service_proto_goTypes,
		DependencyIndexes: file_cbws_service_management_service_usage_v1alpha1_service_proto_depIdxs,
		EnumInfos:         file_cbws_service_management_service_usage_v1alpha1_service_proto_enumTypes,
		MessageInfos:      file_cbws_service_management_service_usage_v1alpha1_service_proto_msgTypes,
	}.Build()
	File_cbws_service_management_service_usage_v1alpha1_service_proto = out.File
	file_cbws_service_management_service_usage_v1alpha1_service_proto_rawDesc = nil
	file_cbws_service_management_service_usage_v1alpha1_service_proto_goTypes = nil
	file_cbws_service_management_service_usage_v1alpha1_service_proto_depIdxs = nil
}
