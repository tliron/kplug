// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: kplug.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PluginInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version    string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Api        string `protobuf:"bytes,4,opt,name=api,proto3" json:"api,omitempty"`
	ApiVersion string `protobuf:"bytes,5,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
}

func (x *PluginInformation) Reset() {
	*x = PluginInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kplug_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginInformation) ProtoMessage() {}

func (x *PluginInformation) ProtoReflect() protoreflect.Message {
	mi := &file_kplug_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginInformation.ProtoReflect.Descriptor instead.
func (*PluginInformation) Descriptor() ([]byte, []int) {
	return file_kplug_proto_rawDescGZIP(), []int{0}
}

func (x *PluginInformation) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PluginInformation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginInformation) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PluginInformation) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *PluginInformation) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

type PluginHeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accepted          bool   `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	NotAcceptedReason string `protobuf:"bytes,2,opt,name=notAcceptedReason,proto3" json:"notAcceptedReason,omitempty"`
}

func (x *PluginHeartbeatResponse) Reset() {
	*x = PluginHeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kplug_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginHeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginHeartbeatResponse) ProtoMessage() {}

func (x *PluginHeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kplug_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginHeartbeatResponse.ProtoReflect.Descriptor instead.
func (*PluginHeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_kplug_proto_rawDescGZIP(), []int{1}
}

func (x *PluginHeartbeatResponse) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

func (x *PluginHeartbeatResponse) GetNotAcceptedReason() string {
	if x != nil {
		return x.NotAcceptedReason
	}
	return ""
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Yaml string `protobuf:"bytes,2,opt,name=yaml,proto3" json:"yaml,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kplug_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_kplug_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_kplug_proto_rawDescGZIP(), []int{2}
}

func (x *Resource) GetYaml() string {
	if x != nil {
		return x.Yaml
	}
	return ""
}

type Resources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base       *Resource            `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Extensions map[string]*Resource `protobuf:"bytes,2,rep,name=extensions,proto3" json:"extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Resources) Reset() {
	*x = Resources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kplug_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resources) ProtoMessage() {}

func (x *Resources) ProtoReflect() protoreflect.Message {
	mi := &file_kplug_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resources.ProtoReflect.Descriptor instead.
func (*Resources) Descriptor() ([]byte, []int) {
	return file_kplug_proto_rawDescGZIP(), []int{3}
}

func (x *Resources) GetBase() *Resource {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *Resources) GetExtensions() map[string]*Resource {
	if x != nil {
		return x.Extensions
	}
	return nil
}

var File_kplug_proto protoreflect.FileDescriptor

var file_kplug_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6b, 0x70, 0x6c, 0x75, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6b,
	0x70, 0x6c, 0x75, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x11, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70,
	0x69, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x63, 0x0a, 0x17, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x6e, 0x6f, 0x74, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x6e, 0x6f, 0x74, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x22, 0xc2, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6b, 0x70, 0x6c, 0x75, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x6b, 0x70, 0x6c, 0x75, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x4e, 0x0a, 0x0f, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x6b, 0x70, 0x6c, 0x75, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x59, 0x0a, 0x0a, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x0f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x18, 0x2e, 0x6b,
	0x70, 0x6c, 0x75, 0x67, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1e, 0x2e, 0x6b, 0x70, 0x6c, 0x75, 0x67, 0x2e, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x98, 0x01, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x6b, 0x70,
	0x6c, 0x75, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x10, 0x2e,
	0x6b, 0x70, 0x6c, 0x75, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12,
	0x2c, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x6b, 0x70, 0x6c, 0x75,
	0x67, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x10, 0x2e, 0x6b, 0x70,
	0x6c, 0x75, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x32, 0x0a,
	0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x6b, 0x70, 0x6c, 0x75, 0x67, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x6c, 0x69, 0x72, 0x6f, 0x6e, 0x2f, 0x6b, 0x70, 0x6c, 0x75, 0x67, 0x2f, 0x6b, 0x70, 0x6c,
	0x75, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kplug_proto_rawDescOnce sync.Once
	file_kplug_proto_rawDescData = file_kplug_proto_rawDesc
)

func file_kplug_proto_rawDescGZIP() []byte {
	file_kplug_proto_rawDescOnce.Do(func() {
		file_kplug_proto_rawDescData = protoimpl.X.CompressGZIP(file_kplug_proto_rawDescData)
	})
	return file_kplug_proto_rawDescData
}

var file_kplug_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_kplug_proto_goTypes = []interface{}{
	(*PluginInformation)(nil),       // 0: kplug.PluginInformation
	(*PluginHeartbeatResponse)(nil), // 1: kplug.PluginHeartbeatResponse
	(*Resource)(nil),                // 2: kplug.Resource
	(*Resources)(nil),               // 3: kplug.Resources
	nil,                             // 4: kplug.Resources.ExtensionsEntry
	(*emptypb.Empty)(nil),           // 5: google.protobuf.Empty
}
var file_kplug_proto_depIdxs = []int32{
	2, // 0: kplug.Resources.base:type_name -> kplug.Resource
	4, // 1: kplug.Resources.extensions:type_name -> kplug.Resources.ExtensionsEntry
	2, // 2: kplug.Resources.ExtensionsEntry.value:type_name -> kplug.Resource
	0, // 3: kplug.Controller.pluginHeartbeat:input_type -> kplug.PluginInformation
	3, // 4: kplug.Plugin.create:input_type -> kplug.Resources
	3, // 5: kplug.Plugin.update:input_type -> kplug.Resources
	3, // 6: kplug.Plugin.delete:input_type -> kplug.Resources
	1, // 7: kplug.Controller.pluginHeartbeat:output_type -> kplug.PluginHeartbeatResponse
	3, // 8: kplug.Plugin.create:output_type -> kplug.Resources
	3, // 9: kplug.Plugin.update:output_type -> kplug.Resources
	5, // 10: kplug.Plugin.delete:output_type -> google.protobuf.Empty
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_kplug_proto_init() }
func file_kplug_proto_init() {
	if File_kplug_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kplug_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginInformation); i {
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
		file_kplug_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginHeartbeatResponse); i {
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
		file_kplug_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_kplug_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resources); i {
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
			RawDescriptor: file_kplug_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_kplug_proto_goTypes,
		DependencyIndexes: file_kplug_proto_depIdxs,
		MessageInfos:      file_kplug_proto_msgTypes,
	}.Build()
	File_kplug_proto = out.File
	file_kplug_proto_rawDesc = nil
	file_kplug_proto_goTypes = nil
	file_kplug_proto_depIdxs = nil
}