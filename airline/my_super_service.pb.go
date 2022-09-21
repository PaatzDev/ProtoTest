// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.3
// source: my_super_service.proto

package v1

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

type ListNotamsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListNotamsInput) Reset() {
	*x = ListNotamsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_super_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNotamsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotamsInput) ProtoMessage() {}

func (x *ListNotamsInput) ProtoReflect() protoreflect.Message {
	mi := &file_my_super_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotamsInput.ProtoReflect.Descriptor instead.
func (*ListNotamsInput) Descriptor() ([]byte, []int) {
	return file_my_super_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListNotamsInput) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListNotamsOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notams []*Notam `protobuf:"bytes,1,rep,name=notams,proto3" json:"notams,omitempty"`
}

func (x *ListNotamsOutput) Reset() {
	*x = ListNotamsOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_super_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNotamsOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotamsOutput) ProtoMessage() {}

func (x *ListNotamsOutput) ProtoReflect() protoreflect.Message {
	mi := &file_my_super_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotamsOutput.ProtoReflect.Descriptor instead.
func (*ListNotamsOutput) Descriptor() ([]byte, []int) {
	return file_my_super_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListNotamsOutput) GetNotams() []*Notam {
	if x != nil {
		return x.Notams
	}
	return nil
}

type AddNotamInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notam *Notam `protobuf:"bytes,1,opt,name=notam,proto3" json:"notam,omitempty"`
}

func (x *AddNotamInput) Reset() {
	*x = AddNotamInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_super_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNotamInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNotamInput) ProtoMessage() {}

func (x *AddNotamInput) ProtoReflect() protoreflect.Message {
	mi := &file_my_super_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNotamInput.ProtoReflect.Descriptor instead.
func (*AddNotamInput) Descriptor() ([]byte, []int) {
	return file_my_super_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddNotamInput) GetNotam() *Notam {
	if x != nil {
		return x.Notam
	}
	return nil
}

type AddNotamOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notam *Notam `protobuf:"bytes,1,opt,name=notam,proto3" json:"notam,omitempty"`
}

func (x *AddNotamOutput) Reset() {
	*x = AddNotamOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_super_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNotamOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNotamOutput) ProtoMessage() {}

func (x *AddNotamOutput) ProtoReflect() protoreflect.Message {
	mi := &file_my_super_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNotamOutput.ProtoReflect.Descriptor instead.
func (*AddNotamOutput) Descriptor() ([]byte, []int) {
	return file_my_super_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddNotamOutput) GetNotam() *Notam {
	if x != nil {
		return x.Notam
	}
	return nil
}

type Notam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Notam) Reset() {
	*x = Notam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_super_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notam) ProtoMessage() {}

func (x *Notam) ProtoReflect() protoreflect.Message {
	mi := &file_my_super_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notam.ProtoReflect.Descriptor instead.
func (*Notam) Descriptor() ([]byte, []int) {
	return file_my_super_service_proto_rawDescGZIP(), []int{4}
}

func (x *Notam) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notam) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_my_super_service_proto protoreflect.FileDescriptor

var file_my_super_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x79, 0x5f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x79, 0x6f, 0x75, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x27, 0x0a, 0x0f, 0x4c, 0x69, 0x73,
	0x74, 0x4e, 0x6f, 0x74, 0x61, 0x6d, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x42, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x61, 0x6d, 0x73,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x61, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x79, 0x6f, 0x75, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x61, 0x6d, 0x52, 0x06,
	0x6e, 0x6f, 0x74, 0x61, 0x6d, 0x73, 0x22, 0x3d, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74,
	0x61, 0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x61, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x79, 0x6f, 0x75, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x61, 0x6d, 0x52, 0x05,
	0x6e, 0x6f, 0x74, 0x61, 0x6d, 0x22, 0x3e, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x61,
	0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x61, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x79, 0x6f, 0x75, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x61, 0x6d, 0x52, 0x05,
	0x6e, 0x6f, 0x74, 0x61, 0x6d, 0x22, 0x31, 0x0a, 0x05, 0x4e, 0x6f, 0x74, 0x61, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xb0, 0x01, 0x0a, 0x0e, 0x4d, 0x79, 0x53,
	0x75, 0x70, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x4c,
	0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x61, 0x6d, 0x73, 0x12, 0x20, 0x2e, 0x79, 0x6f, 0x75, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4e, 0x6f, 0x74, 0x61, 0x6d, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x6f,
	0x75, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4e, 0x6f, 0x74, 0x61, 0x6d, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x4b,
	0x0a, 0x08, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x61, 0x6d, 0x12, 0x1e, 0x2e, 0x79, 0x6f, 0x75,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x4e, 0x6f, 0x74, 0x61, 0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1f, 0x2e, 0x79, 0x6f, 0x75,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x4e, 0x6f, 0x74, 0x61, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x25, 0x5a, 0x23, 0x61,
	0x65, 0x72, 0x6f, 0x2e, 0x73, 0x6b, 0x79, 0x6c, 0x69, 0x6e, 0x78, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x79, 0x6f, 0x75, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_my_super_service_proto_rawDescOnce sync.Once
	file_my_super_service_proto_rawDescData = file_my_super_service_proto_rawDesc
)

func file_my_super_service_proto_rawDescGZIP() []byte {
	file_my_super_service_proto_rawDescOnce.Do(func() {
		file_my_super_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_my_super_service_proto_rawDescData)
	})
	return file_my_super_service_proto_rawDescData
}

var file_my_super_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_my_super_service_proto_goTypes = []interface{}{
	(*ListNotamsInput)(nil),  // 0: your.service.v1.ListNotamsInput
	(*ListNotamsOutput)(nil), // 1: your.service.v1.ListNotamsOutput
	(*AddNotamInput)(nil),    // 2: your.service.v1.AddNotamInput
	(*AddNotamOutput)(nil),   // 3: your.service.v1.AddNotamOutput
	(*Notam)(nil),            // 4: your.service.v1.Notam
}
var file_my_super_service_proto_depIdxs = []int32{
	4, // 0: your.service.v1.ListNotamsOutput.notams:type_name -> your.service.v1.Notam
	4, // 1: your.service.v1.AddNotamInput.notam:type_name -> your.service.v1.Notam
	4, // 2: your.service.v1.AddNotamOutput.notam:type_name -> your.service.v1.Notam
	0, // 3: your.service.v1.MySuperService.ListNotams:input_type -> your.service.v1.ListNotamsInput
	2, // 4: your.service.v1.MySuperService.AddNotam:input_type -> your.service.v1.AddNotamInput
	1, // 5: your.service.v1.MySuperService.ListNotams:output_type -> your.service.v1.ListNotamsOutput
	3, // 6: your.service.v1.MySuperService.AddNotam:output_type -> your.service.v1.AddNotamOutput
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_my_super_service_proto_init() }
func file_my_super_service_proto_init() {
	if File_my_super_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_my_super_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNotamsInput); i {
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
		file_my_super_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNotamsOutput); i {
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
		file_my_super_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNotamInput); i {
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
		file_my_super_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNotamOutput); i {
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
		file_my_super_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notam); i {
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
			RawDescriptor: file_my_super_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_my_super_service_proto_goTypes,
		DependencyIndexes: file_my_super_service_proto_depIdxs,
		MessageInfos:      file_my_super_service_proto_msgTypes,
	}.Build()
	File_my_super_service_proto = out.File
	file_my_super_service_proto_rawDesc = nil
	file_my_super_service_proto_goTypes = nil
	file_my_super_service_proto_depIdxs = nil
}