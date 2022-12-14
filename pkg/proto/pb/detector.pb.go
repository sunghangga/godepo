// To get an in-depth walk-through of this file and the related examples, see:
// https://developers.google.com/protocol-buffers/docs/tutorials

// [START declaration]

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: detector.proto

package pb

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

// [START messages]
type Counter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RollingPeriodStartUnix int64   `protobuf:"varint,1,opt,name=rolling_period_start_unix,json=rollingPeriodStartUnix,proto3" json:"rolling_period_start_unix,omitempty"`
	Received               float64 `protobuf:"fixed64,2,opt,name=received,proto3" json:"received,omitempty"`
}

func (x *Counter) Reset() {
	*x = Counter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Counter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Counter) ProtoMessage() {}

func (x *Counter) ProtoReflect() protoreflect.Message {
	mi := &file_detector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Counter.ProtoReflect.Descriptor instead.
func (*Counter) Descriptor() ([]byte, []int) {
	return file_detector_proto_rawDescGZIP(), []int{0}
}

func (x *Counter) GetRollingPeriodStartUnix() int64 {
	if x != nil {
		return x.RollingPeriodStartUnix
	}
	return 0
}

func (x *Counter) GetReceived() float64 {
	if x != nil {
		return x.Received
	}
	return 0
}

var File_detector_proto protoreflect.FileDescriptor

var file_detector_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x60, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x19, 0x72,
	0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x75, 0x6e, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16,
	0x72, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x55, 0x6e, 0x69, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_detector_proto_rawDescOnce sync.Once
	file_detector_proto_rawDescData = file_detector_proto_rawDesc
)

func file_detector_proto_rawDescGZIP() []byte {
	file_detector_proto_rawDescOnce.Do(func() {
		file_detector_proto_rawDescData = protoimpl.X.CompressGZIP(file_detector_proto_rawDescData)
	})
	return file_detector_proto_rawDescData
}

var file_detector_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_detector_proto_goTypes = []interface{}{
	(*Counter)(nil), // 0: Counter
}
var file_detector_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_detector_proto_init() }
func file_detector_proto_init() {
	if File_detector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_detector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Counter); i {
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
			RawDescriptor: file_detector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_detector_proto_goTypes,
		DependencyIndexes: file_detector_proto_depIdxs,
		MessageInfos:      file_detector_proto_msgTypes,
	}.Build()
	File_detector_proto = out.File
	file_detector_proto_rawDesc = nil
	file_detector_proto_goTypes = nil
	file_detector_proto_depIdxs = nil
}
