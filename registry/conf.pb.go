// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: conf.proto

package registry

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

type RegistryConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Nacos         *RegistryConfig_Nacos  `protobuf:"bytes,1,opt,name=nacos,proto3" json:"nacos,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegistryConfig) Reset() {
	*x = RegistryConfig{}
	mi := &file_conf_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegistryConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryConfig) ProtoMessage() {}

func (x *RegistryConfig) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryConfig.ProtoReflect.Descriptor instead.
func (*RegistryConfig) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0}
}

func (x *RegistryConfig) GetNacos() *RegistryConfig_Nacos {
	if x != nil {
		return x.Nacos
	}
	return nil
}

type RegistryConfig_Nacos struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	Enable        bool                           `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	Config        *RegistryConfig_Nacos_Config   `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	Registry      *RegistryConfig_Nacos_Registry `protobuf:"bytes,3,opt,name=registry,proto3" json:"registry,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegistryConfig_Nacos) Reset() {
	*x = RegistryConfig_Nacos{}
	mi := &file_conf_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegistryConfig_Nacos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryConfig_Nacos) ProtoMessage() {}

func (x *RegistryConfig_Nacos) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryConfig_Nacos.ProtoReflect.Descriptor instead.
func (*RegistryConfig_Nacos) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RegistryConfig_Nacos) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *RegistryConfig_Nacos) GetConfig() *RegistryConfig_Nacos_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *RegistryConfig_Nacos) GetRegistry() *RegistryConfig_Nacos_Registry {
	if x != nil {
		return x.Registry
	}
	return nil
}

type RegistryConfig_Nacos_Config struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port          uint64                 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Namespace     string                 `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Group         string                 `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	DataId        string                 `protobuf:"bytes,5,opt,name=data_id,json=dataId,proto3" json:"data_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegistryConfig_Nacos_Config) Reset() {
	*x = RegistryConfig_Nacos_Config{}
	mi := &file_conf_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegistryConfig_Nacos_Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryConfig_Nacos_Config) ProtoMessage() {}

func (x *RegistryConfig_Nacos_Config) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryConfig_Nacos_Config.ProtoReflect.Descriptor instead.
func (*RegistryConfig_Nacos_Config) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *RegistryConfig_Nacos_Config) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RegistryConfig_Nacos_Config) GetPort() uint64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RegistryConfig_Nacos_Config) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *RegistryConfig_Nacos_Config) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *RegistryConfig_Nacos_Config) GetDataId() string {
	if x != nil {
		return x.DataId
	}
	return ""
}

type RegistryConfig_Nacos_Registry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port          uint64                 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Namespace     string                 `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegistryConfig_Nacos_Registry) Reset() {
	*x = RegistryConfig_Nacos_Registry{}
	mi := &file_conf_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegistryConfig_Nacos_Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryConfig_Nacos_Registry) ProtoMessage() {}

func (x *RegistryConfig_Nacos_Registry) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryConfig_Nacos_Registry.ProtoReflect.Descriptor instead.
func (*RegistryConfig_Nacos_Registry) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *RegistryConfig_Nacos_Registry) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RegistryConfig_Nacos_Registry) GetPort() uint64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RegistryConfig_Nacos_Registry) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_conf_proto protoreflect.FileDescriptor

const file_conf_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"conf.proto\x12\bregistry\"\xca\x03\n" +
	"\x0eRegistryConfig\x124\n" +
	"\x05nacos\x18\x01 \x01(\v2\x1e.registry.RegistryConfig.NacosR\x05nacos\x1a\x81\x03\n" +
	"\x05Nacos\x12\x16\n" +
	"\x06enable\x18\x01 \x01(\bR\x06enable\x12=\n" +
	"\x06config\x18\x02 \x01(\v2%.registry.RegistryConfig.Nacos.ConfigR\x06config\x12C\n" +
	"\bregistry\x18\x03 \x01(\v2'.registry.RegistryConfig.Nacos.RegistryR\bregistry\x1a\x83\x01\n" +
	"\x06Config\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\tR\aaddress\x12\x12\n" +
	"\x04port\x18\x02 \x01(\x04R\x04port\x12\x1c\n" +
	"\tnamespace\x18\x03 \x01(\tR\tnamespace\x12\x14\n" +
	"\x05group\x18\x04 \x01(\tR\x05group\x12\x17\n" +
	"\adata_id\x18\x05 \x01(\tR\x06dataId\x1aV\n" +
	"\bRegistry\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\tR\aaddress\x12\x12\n" +
	"\x04port\x18\x02 \x01(\x04R\x04port\x12\x1c\n" +
	"\tnamespace\x18\x03 \x01(\tR\tnamespaceB,Z*github.com/adgers/common/registry;registryb\x06proto3"

var (
	file_conf_proto_rawDescOnce sync.Once
	file_conf_proto_rawDescData []byte
)

func file_conf_proto_rawDescGZIP() []byte {
	file_conf_proto_rawDescOnce.Do(func() {
		file_conf_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_conf_proto_rawDesc), len(file_conf_proto_rawDesc)))
	})
	return file_conf_proto_rawDescData
}

var file_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_conf_proto_goTypes = []any{
	(*RegistryConfig)(nil),                // 0: registry.RegistryConfig
	(*RegistryConfig_Nacos)(nil),          // 1: registry.RegistryConfig.Nacos
	(*RegistryConfig_Nacos_Config)(nil),   // 2: registry.RegistryConfig.Nacos.Config
	(*RegistryConfig_Nacos_Registry)(nil), // 3: registry.RegistryConfig.Nacos.Registry
}
var file_conf_proto_depIdxs = []int32{
	1, // 0: registry.RegistryConfig.nacos:type_name -> registry.RegistryConfig.Nacos
	2, // 1: registry.RegistryConfig.Nacos.config:type_name -> registry.RegistryConfig.Nacos.Config
	3, // 2: registry.RegistryConfig.Nacos.registry:type_name -> registry.RegistryConfig.Nacos.Registry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_conf_proto_init() }
func file_conf_proto_init() {
	if File_conf_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_conf_proto_rawDesc), len(file_conf_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_proto_goTypes,
		DependencyIndexes: file_conf_proto_depIdxs,
		MessageInfos:      file_conf_proto_msgTypes,
	}.Build()
	File_conf_proto = out.File
	file_conf_proto_goTypes = nil
	file_conf_proto_depIdxs = nil
}
