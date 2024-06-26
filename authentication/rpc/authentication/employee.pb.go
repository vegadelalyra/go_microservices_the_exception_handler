// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: employee.proto

package authentication

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

type EmployeeAddRequest_Department int32

const (
	EmployeeAddRequest_SALES     EmployeeAddRequest_Department = 0
	EmployeeAddRequest_MARKETING EmployeeAddRequest_Department = 1
	EmployeeAddRequest_ADMIN     EmployeeAddRequest_Department = 2
	EmployeeAddRequest_STAFF     EmployeeAddRequest_Department = 3
)

// Enum value maps for EmployeeAddRequest_Department.
var (
	EmployeeAddRequest_Department_name = map[int32]string{
		0: "SALES",
		1: "MARKETING",
		2: "ADMIN",
		3: "STAFF",
	}
	EmployeeAddRequest_Department_value = map[string]int32{
		"SALES":     0,
		"MARKETING": 1,
		"ADMIN":     2,
		"STAFF":     3,
	}
)

func (x EmployeeAddRequest_Department) Enum() *EmployeeAddRequest_Department {
	p := new(EmployeeAddRequest_Department)
	*p = x
	return p
}

func (x EmployeeAddRequest_Department) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmployeeAddRequest_Department) Descriptor() protoreflect.EnumDescriptor {
	return file_employee_proto_enumTypes[0].Descriptor()
}

func (EmployeeAddRequest_Department) Type() protoreflect.EnumType {
	return &file_employee_proto_enumTypes[0]
}

func (x EmployeeAddRequest_Department) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmployeeAddRequest_Department.Descriptor instead.
func (EmployeeAddRequest_Department) EnumDescriptor() ([]byte, []int) {
	return file_employee_proto_rawDescGZIP(), []int{0, 0}
}

type EmployeeAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string                         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address *int32                         `protobuf:"varint,2,opt,name=address,proto3,oneof" json:"address,omitempty"`
	Age     *int32                         `protobuf:"varint,3,opt,name=age,proto3,oneof" json:"age,omitempty"`
	Dept    *EmployeeAddRequest_Department `protobuf:"varint,4,opt,name=dept,proto3,enum=authentication.EmployeeAddRequest_Department,oneof" json:"dept,omitempty"`
	Emails  []string                       `protobuf:"bytes,5,rep,name=emails,proto3" json:"emails,omitempty"`
}

func (x *EmployeeAddRequest) Reset() {
	*x = EmployeeAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeAddRequest) ProtoMessage() {}

func (x *EmployeeAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeAddRequest.ProtoReflect.Descriptor instead.
func (*EmployeeAddRequest) Descriptor() ([]byte, []int) {
	return file_employee_proto_rawDescGZIP(), []int{0}
}

func (x *EmployeeAddRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EmployeeAddRequest) GetAddress() int32 {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return 0
}

func (x *EmployeeAddRequest) GetAge() int32 {
	if x != nil && x.Age != nil {
		return *x.Age
	}
	return 0
}

func (x *EmployeeAddRequest) GetDept() EmployeeAddRequest_Department {
	if x != nil && x.Dept != nil {
		return *x.Dept
	}
	return EmployeeAddRequest_SALES
}

func (x *EmployeeAddRequest) GetEmails() []string {
	if x != nil {
		return x.Emails
	}
	return nil
}

type EmployeeAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EmployeeAddResponse) Reset() {
	*x = EmployeeAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeAddResponse) ProtoMessage() {}

func (x *EmployeeAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeAddResponse.ProtoReflect.Descriptor instead.
func (*EmployeeAddResponse) Descriptor() ([]byte, []int) {
	return file_employee_proto_rawDescGZIP(), []int{1}
}

func (x *EmployeeAddResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EmployeeAddResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_employee_proto protoreflect.FileDescriptor

var file_employee_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x99, 0x02, 0x0a, 0x12, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x03, 0x61, 0x67, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x46, 0x0a, 0x04, 0x64, 0x65, 0x70, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x02,
	0x52, 0x04, 0x64, 0x65, 0x70, 0x74, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x73, 0x22, 0x3c, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x09, 0x0a, 0x05, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x41,
	0x52, 0x4b, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d,
	0x49, 0x4e, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x46, 0x46, 0x10, 0x03, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x5f,
	0x61, 0x67, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x22, 0x39, 0x0a, 0x13,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0x69, 0x0a, 0x0f, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x41, 0x64,
	0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_employee_proto_rawDescOnce sync.Once
	file_employee_proto_rawDescData = file_employee_proto_rawDesc
)

func file_employee_proto_rawDescGZIP() []byte {
	file_employee_proto_rawDescOnce.Do(func() {
		file_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_employee_proto_rawDescData)
	})
	return file_employee_proto_rawDescData
}

var file_employee_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_employee_proto_goTypes = []any{
	(EmployeeAddRequest_Department)(0), // 0: authentication.EmployeeAddRequest.Department
	(*EmployeeAddRequest)(nil),         // 1: authentication.EmployeeAddRequest
	(*EmployeeAddResponse)(nil),        // 2: authentication.EmployeeAddResponse
}
var file_employee_proto_depIdxs = []int32{
	0, // 0: authentication.EmployeeAddRequest.dept:type_name -> authentication.EmployeeAddRequest.Department
	1, // 1: authentication.EmployeeService.AddEmployee:input_type -> authentication.EmployeeAddRequest
	2, // 2: authentication.EmployeeService.AddEmployee:output_type -> authentication.EmployeeAddResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_employee_proto_init() }
func file_employee_proto_init() {
	if File_employee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_employee_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EmployeeAddRequest); i {
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
		file_employee_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EmployeeAddResponse); i {
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
	file_employee_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_employee_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_employee_proto_goTypes,
		DependencyIndexes: file_employee_proto_depIdxs,
		EnumInfos:         file_employee_proto_enumTypes,
		MessageInfos:      file_employee_proto_msgTypes,
	}.Build()
	File_employee_proto = out.File
	file_employee_proto_rawDesc = nil
	file_employee_proto_goTypes = nil
	file_employee_proto_depIdxs = nil
}
