// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: vagrant-ruby/builtin/myplugin/plugin.proto

package myplugin

import (
	proto "github.com/golang/protobuf/proto"
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

type UpResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpResult) Reset() {
	*x = UpResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vagrant_ruby_builtin_myplugin_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpResult) ProtoMessage() {}

func (x *UpResult) ProtoReflect() protoreflect.Message {
	mi := &file_vagrant_ruby_builtin_myplugin_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpResult.ProtoReflect.Descriptor instead.
func (*UpResult) Descriptor() ([]byte, []int) {
	return file_vagrant_ruby_builtin_myplugin_plugin_proto_rawDescGZIP(), []int{0}
}

var File_vagrant_ruby_builtin_myplugin_plugin_proto protoreflect.FileDescriptor

var file_vagrant_ruby_builtin_myplugin_plugin_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x76, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2d, 0x72, 0x75, 0x62, 0x79, 0x2f, 0x62,
	0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x2f, 0x6d, 0x79, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x79,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x22, 0x0a, 0x0a, 0x08, 0x55, 0x70, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x42, 0x1f, 0x5a, 0x1d, 0x76, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2d, 0x72, 0x75,
	0x62, 0x79, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x2f, 0x6d, 0x79, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vagrant_ruby_builtin_myplugin_plugin_proto_rawDescOnce sync.Once
	file_vagrant_ruby_builtin_myplugin_plugin_proto_rawDescData = file_vagrant_ruby_builtin_myplugin_plugin_proto_rawDesc
)

func file_vagrant_ruby_builtin_myplugin_plugin_proto_rawDescGZIP() []byte {
	file_vagrant_ruby_builtin_myplugin_plugin_proto_rawDescOnce.Do(func() {
		file_vagrant_ruby_builtin_myplugin_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_vagrant_ruby_builtin_myplugin_plugin_proto_rawDescData)
	})
	return file_vagrant_ruby_builtin_myplugin_plugin_proto_rawDescData
}

var file_vagrant_ruby_builtin_myplugin_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_vagrant_ruby_builtin_myplugin_plugin_proto_goTypes = []interface{}{
	(*UpResult)(nil), // 0: myplugin.UpResult
}
var file_vagrant_ruby_builtin_myplugin_plugin_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vagrant_ruby_builtin_myplugin_plugin_proto_init() }
func file_vagrant_ruby_builtin_myplugin_plugin_proto_init() {
	if File_vagrant_ruby_builtin_myplugin_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vagrant_ruby_builtin_myplugin_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpResult); i {
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
			RawDescriptor: file_vagrant_ruby_builtin_myplugin_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vagrant_ruby_builtin_myplugin_plugin_proto_goTypes,
		DependencyIndexes: file_vagrant_ruby_builtin_myplugin_plugin_proto_depIdxs,
		MessageInfos:      file_vagrant_ruby_builtin_myplugin_plugin_proto_msgTypes,
	}.Build()
	File_vagrant_ruby_builtin_myplugin_plugin_proto = out.File
	file_vagrant_ruby_builtin_myplugin_plugin_proto_rawDesc = nil
	file_vagrant_ruby_builtin_myplugin_plugin_proto_goTypes = nil
	file_vagrant_ruby_builtin_myplugin_plugin_proto_depIdxs = nil
}
