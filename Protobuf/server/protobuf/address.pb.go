// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: address.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Our address book file is just one of these.
type AddressBook struct {
	state             protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_People *[]*Person             `protobuf:"bytes,1,rep,name=people,proto3"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *AddressBook) Reset() {
	*x = AddressBook{}
	mi := &file_address_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressBook) ProtoMessage() {}

func (x *AddressBook) ProtoReflect() protoreflect.Message {
	mi := &file_address_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *AddressBook) GetPeople() []*Person {
	if x != nil {
		if x.xxx_hidden_People != nil {
			return *x.xxx_hidden_People
		}
	}
	return nil
}

func (x *AddressBook) SetPeople(v []*Person) {
	x.xxx_hidden_People = &v
}

type AddressBook_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	People []*Person
}

func (b0 AddressBook_builder) Build() *AddressBook {
	m0 := &AddressBook{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_People = &b.People
	return m0
}

var File_address_proto protoreflect.FileDescriptor

var file_address_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x6f,
	0x6f, 0x6b, 0x12, 0x24, 0x0a, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x52, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_address_proto_goTypes = []any{
	(*AddressBook)(nil), // 0: test.AddressBook
	(*Person)(nil),      // 1: test.Person
}
var file_address_proto_depIdxs = []int32{
	1, // 0: test.AddressBook.people:type_name -> test.Person
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_address_proto_init() }
func file_address_proto_init() {
	if File_address_proto != nil {
		return
	}
	file_person_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_address_proto_rawDesc), len(file_address_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_address_proto_goTypes,
		DependencyIndexes: file_address_proto_depIdxs,
		MessageInfos:      file_address_proto_msgTypes,
	}.Build()
	File_address_proto = out.File
	file_address_proto_goTypes = nil
	file_address_proto_depIdxs = nil
}
