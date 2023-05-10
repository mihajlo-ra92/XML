// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.0
// source: auth_service.proto

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
	return file_auth_service_proto_enumTypes[0].Descriptor()
}

func (User_UserType) Type() protoreflect.EnumType {
	return &file_auth_service_proto_enumTypes[0]
}

func (x User_UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User_UserType.Descriptor instead.
func (User_UserType) EnumDescriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{8, 0}
}

type AuthCreateAccommodationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt       string   `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	Name      string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location  string   `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Benefits  []string `protobuf:"bytes,4,rep,name=benefits,proto3" json:"benefits,omitempty"`
	Pictures  []string `protobuf:"bytes,5,rep,name=pictures,proto3" json:"pictures,omitempty"`
	MinGuests uint32   `protobuf:"varint,6,opt,name=minGuests,proto3" json:"minGuests,omitempty"`
	MaxGuests uint32   `protobuf:"varint,7,opt,name=maxGuests,proto3" json:"maxGuests,omitempty"`
}

func (x *AuthCreateAccommodationRequest) Reset() {
	*x = AuthCreateAccommodationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthCreateAccommodationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthCreateAccommodationRequest) ProtoMessage() {}

func (x *AuthCreateAccommodationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthCreateAccommodationRequest.ProtoReflect.Descriptor instead.
func (*AuthCreateAccommodationRequest) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{0}
}

func (x *AuthCreateAccommodationRequest) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *AuthCreateAccommodationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthCreateAccommodationRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *AuthCreateAccommodationRequest) GetBenefits() []string {
	if x != nil {
		return x.Benefits
	}
	return nil
}

func (x *AuthCreateAccommodationRequest) GetPictures() []string {
	if x != nil {
		return x.Pictures
	}
	return nil
}

func (x *AuthCreateAccommodationRequest) GetMinGuests() uint32 {
	if x != nil {
		return x.MinGuests
	}
	return 0
}

func (x *AuthCreateAccommodationRequest) GetMaxGuests() uint32 {
	if x != nil {
		return x.MaxGuests
	}
	return 0
}

type AuthCreateAccommodationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accomodation *Accommodation `protobuf:"bytes,1,opt,name=accomodation,proto3" json:"accomodation,omitempty"`
}

func (x *AuthCreateAccommodationResponse) Reset() {
	*x = AuthCreateAccommodationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthCreateAccommodationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthCreateAccommodationResponse) ProtoMessage() {}

func (x *AuthCreateAccommodationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthCreateAccommodationResponse.ProtoReflect.Descriptor instead.
func (*AuthCreateAccommodationResponse) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{1}
}

func (x *AuthCreateAccommodationResponse) GetAccomodation() *Accommodation {
	if x != nil {
		return x.Accomodation
	}
	return nil
}

type Accommodation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	HostId    string   `protobuf:"bytes,2,opt,name=hostId,proto3" json:"hostId,omitempty"`
	Name      string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Location  string   `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Benefits  []string `protobuf:"bytes,5,rep,name=benefits,proto3" json:"benefits,omitempty"`
	Pictures  []string `protobuf:"bytes,6,rep,name=pictures,proto3" json:"pictures,omitempty"`
	MinGuests uint32   `protobuf:"varint,7,opt,name=minGuests,proto3" json:"minGuests,omitempty"`
	MaxGuests uint32   `protobuf:"varint,8,opt,name=maxGuests,proto3" json:"maxGuests,omitempty"`
}

func (x *Accommodation) Reset() {
	*x = Accommodation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Accommodation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Accommodation) ProtoMessage() {}

func (x *Accommodation) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Accommodation.ProtoReflect.Descriptor instead.
func (*Accommodation) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{2}
}

func (x *Accommodation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Accommodation) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

func (x *Accommodation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Accommodation) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Accommodation) GetBenefits() []string {
	if x != nil {
		return x.Benefits
	}
	return nil
}

func (x *Accommodation) GetPictures() []string {
	if x != nil {
		return x.Pictures
	}
	return nil
}

func (x *Accommodation) GetMinGuests() uint32 {
	if x != nil {
		return x.MinGuests
	}
	return 0
}

func (x *Accommodation) GetMaxGuests() uint32 {
	if x != nil {
		return x.MaxGuests
	}
	return 0
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login *AuthLogin `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{3}
}

func (x *LoginRequest) GetLogin() *AuthLogin {
	if x != nil {
		return x.Login
	}
	return nil
}

type AuthLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthLogin) Reset() {
	*x = AuthLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthLogin) ProtoMessage() {}

func (x *AuthLogin) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthLogin.ProtoReflect.Descriptor instead.
func (*AuthLogin) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{4}
}

func (x *AuthLogin) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthLogin) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{5}
}

func (x *LoginResponse) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

type AuthUpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt  string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	User *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *AuthUpdateUserRequest) Reset() {
	*x = AuthUpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUpdateUserRequest) ProtoMessage() {}

func (x *AuthUpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUpdateUserRequest.ProtoReflect.Descriptor instead.
func (*AuthUpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{6}
}

func (x *AuthUpdateUserRequest) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *AuthUpdateUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type AuthUpdateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *AuthUpdateUserResponse) Reset() {
	*x = AuthUpdateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUpdateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUpdateUserResponse) ProtoMessage() {}

func (x *AuthUpdateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUpdateUserResponse.ProtoReflect.Descriptor instead.
func (*AuthUpdateUserResponse) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{7}
}

func (x *AuthUpdateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
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
	Address   string        `protobuf:"bytes,8,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[8]
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
	return file_auth_service_proto_rawDescGZIP(), []int{8}
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

func (x *User) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_auth_service_proto protoreflect.FileDescriptor

var file_auth_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x01, 0x0a, 0x1e, 0x41, 0x75, 0x74,
	0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6a,
	0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x47, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x47, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x47, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x22, 0x5a, 0x0a, 0x1f, 0x41, 0x75, 0x74, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xdb, 0x01,
	0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x65, 0x6e, 0x65, 0x66,
	0x69, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x62, 0x65, 0x6e, 0x65, 0x66,
	0x69, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x6d, 0x61, 0x78, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x6d, 0x61, 0x78, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x35, 0x0a, 0x0c, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x22, 0x43, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x21, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x22, 0x49, 0x0a, 0x15, 0x41, 0x75,
	0x74, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6a, 0x77, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x38, 0x0a, 0x16, 0x41, 0x75, 0x74, 0x68, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x95, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x2a, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x75, 0x65, 0x73, 0x74,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x10, 0x02, 0x32, 0xbf, 0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x3a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x06, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x80, 0x01, 0x0a, 0x17, 0x41, 0x75, 0x74, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x64, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x68, 0x61, 0x6a, 0x6c, 0x6f, 0x2d,
	0x72, 0x61, 0x39, 0x32, 0x2f, 0x58, 0x4d, 0x4c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_service_proto_rawDescOnce sync.Once
	file_auth_service_proto_rawDescData = file_auth_service_proto_rawDesc
)

func file_auth_service_proto_rawDescGZIP() []byte {
	file_auth_service_proto_rawDescOnce.Do(func() {
		file_auth_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_service_proto_rawDescData)
	})
	return file_auth_service_proto_rawDescData
}

var file_auth_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_auth_service_proto_goTypes = []interface{}{
	(User_UserType)(0),                      // 0: user.User.UserType
	(*AuthCreateAccommodationRequest)(nil),  // 1: user.AuthCreateAccommodationRequest
	(*AuthCreateAccommodationResponse)(nil), // 2: user.AuthCreateAccommodationResponse
	(*Accommodation)(nil),                   // 3: user.Accommodation
	(*LoginRequest)(nil),                    // 4: user.LoginRequest
	(*AuthLogin)(nil),                       // 5: user.AuthLogin
	(*LoginResponse)(nil),                   // 6: user.LoginResponse
	(*AuthUpdateUserRequest)(nil),           // 7: user.AuthUpdateUserRequest
	(*AuthUpdateUserResponse)(nil),          // 8: user.AuthUpdateUserResponse
	(*User)(nil),                            // 9: user.User
}
var file_auth_service_proto_depIdxs = []int32{
	3, // 0: user.AuthCreateAccommodationResponse.accomodation:type_name -> user.Accommodation
	5, // 1: user.LoginRequest.login:type_name -> user.AuthLogin
	9, // 2: user.AuthUpdateUserRequest.user:type_name -> user.User
	9, // 3: user.AuthUpdateUserResponse.user:type_name -> user.User
	0, // 4: user.User.userType:type_name -> user.User.UserType
	4, // 5: user.AuthService.Login:input_type -> user.LoginRequest
	1, // 6: user.AuthService.AuthCreateAccommodation:input_type -> user.AuthCreateAccommodationRequest
	7, // 7: user.AuthService.AuthUpdateUser:input_type -> user.AuthUpdateUserRequest
	6, // 8: user.AuthService.Login:output_type -> user.LoginResponse
	2, // 9: user.AuthService.AuthCreateAccommodation:output_type -> user.AuthCreateAccommodationResponse
	8, // 10: user.AuthService.AuthUpdateUser:output_type -> user.AuthUpdateUserResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_auth_service_proto_init() }
func file_auth_service_proto_init() {
	if File_auth_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthCreateAccommodationRequest); i {
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
		file_auth_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthCreateAccommodationResponse); i {
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
		file_auth_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Accommodation); i {
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
		file_auth_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_auth_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthLogin); i {
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
		file_auth_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_auth_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthUpdateUserRequest); i {
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
		file_auth_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthUpdateUserResponse); i {
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
		file_auth_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_auth_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_service_proto_goTypes,
		DependencyIndexes: file_auth_service_proto_depIdxs,
		EnumInfos:         file_auth_service_proto_enumTypes,
		MessageInfos:      file_auth_service_proto_msgTypes,
	}.Build()
	File_auth_service_proto = out.File
	file_auth_service_proto_rawDesc = nil
	file_auth_service_proto_goTypes = nil
	file_auth_service_proto_depIdxs = nil
}
