// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: user_service.proto

package user

import (
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

type User_UserType int32

const (
	User_Guest User_UserType = 0
	User_Host  User_UserType = 1
	User_Admin User_UserType = 2
)

// Enum value maps for User_UserType.
var (
	User_UserType_name = map[int32]string{
		0: "Guest",
		1: "Host",
		2: "Admin",
	}
	User_UserType_value = map[string]int32{
		"Guest": 0,
		"Host":  1,
		"Admin": 2,
	}
)

func (x User_UserType) Enum() *User_UserType {
	p := new(User_UserType)
	*p = x
	return p
}

func (x User_UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (User_UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_user_service_proto_enumTypes[0].Descriptor()
}

func (User_UserType) Type() protoreflect.EnumType {
	return &file_user_service_proto_enumTypes[0]
}

func (x User_UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User_UserType.Descriptor instead.
func (User_UserType) EnumDescriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{4, 0}
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{2}
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserType  User_UserType `protobuf:"varint,2,opt,name=userType,proto3,enum=user.User_UserType" json:"userType,omitempty"`
	Username  string        `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password  string        `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Email     string        `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	FirstName string        `protobuf:"bytes,6,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName  string        `protobuf:"bytes,7,opt,name=LastName,proto3" json:"LastName,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetUserType() User_UserType {
	if x != nil {
		return x.UserType
	}
	return User_Guest
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

var File_user_service_proto protoreflect.FileDescriptor

var file_user_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0xfb, 0x01, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x75, 0x65, 0x73,
	0x74, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x10, 0x02, 0x32, 0x91, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x42, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x42, 0x3a, 0x5a, 0x38,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x6d, 0x61, 0x72,
	0x61, 0x72, 0x61, 0x6e, 0x6b, 0x6f, 0x76, 0x69, 0x63, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_service_proto_rawDescOnce sync.Once
	file_user_service_proto_rawDescData = file_user_service_proto_rawDesc
)

func file_user_service_proto_rawDescGZIP() []byte {
	file_user_service_proto_rawDescOnce.Do(func() {
		file_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_service_proto_rawDescData)
	})
	return file_user_service_proto_rawDescData
}

var file_user_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_service_proto_goTypes = []interface{}{
	(User_UserType)(0),     // 0: user.User.UserType
	(*GetRequest)(nil),     // 1: user.GetRequest
	(*GetResponse)(nil),    // 2: user.GetResponse
	(*GetAllRequest)(nil),  // 3: user.GetAllRequest
	(*GetAllResponse)(nil), // 4: user.GetAllResponse
	(*User)(nil),           // 5: user.User
}
var file_user_service_proto_depIdxs = []int32{
	5, // 0: user.GetResponse.user:type_name -> user.User
	5, // 1: user.GetAllResponse.users:type_name -> user.User
	0, // 2: user.User.userType:type_name -> user.User.UserType
	1, // 3: user.UserService.Get:input_type -> user.GetRequest
	3, // 4: user.UserService.GetAll:input_type -> user.GetAllRequest
	2, // 5: user.UserService.Get:output_type -> user.GetResponse
	4, // 6: user.UserService.GetAll:output_type -> user.GetAllResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_user_service_proto_init() }
func file_user_service_proto_init() {
	if File_user_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
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
		file_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
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
		file_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_user_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_proto_goTypes,
		DependencyIndexes: file_user_service_proto_depIdxs,
		EnumInfos:         file_user_service_proto_enumTypes,
		MessageInfos:      file_user_service_proto_msgTypes,
	}.Build()
	File_user_service_proto = out.File
	file_user_service_proto_rawDesc = nil
	file_user_service_proto_goTypes = nil
	file_user_service_proto_depIdxs = nil
}
