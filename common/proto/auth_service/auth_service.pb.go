// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: auth_service/auth_service.proto

package auth_service

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

type ChangePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChangePasswordMessage *ChangePasswordMessage `protobuf:"bytes,1,opt,name=changePasswordMessage,proto3" json:"changePasswordMessage,omitempty"`
}

func (x *ChangePasswordRequest) Reset() {
	*x = ChangePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_auth_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordRequest) ProtoMessage() {}

func (x *ChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_auth_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_auth_service_auth_service_proto_rawDescGZIP(), []int{0}
}

func (x *ChangePasswordRequest) GetChangePasswordMessage() *ChangePasswordMessage {
	if x != nil {
		return x.ChangePasswordMessage
	}
	return nil
}

type ChangePasswordMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	OldPassword string `protobuf:"bytes,2,opt,name=oldPassword,proto3" json:"oldPassword,omitempty"`
	NewPassword string `protobuf:"bytes,3,opt,name=NewPassword,proto3" json:"NewPassword,omitempty"`
}

func (x *ChangePasswordMessage) Reset() {
	*x = ChangePasswordMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_auth_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordMessage) ProtoMessage() {}

func (x *ChangePasswordMessage) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_auth_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordMessage.ProtoReflect.Descriptor instead.
func (*ChangePasswordMessage) Descriptor() ([]byte, []int) {
	return file_auth_service_auth_service_proto_rawDescGZIP(), []int{1}
}

func (x *ChangePasswordMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ChangePasswordMessage) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *ChangePasswordMessage) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_auth_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_auth_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_auth_service_auth_service_proto_rawDescGZIP(), []int{2}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_auth_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_auth_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_auth_service_auth_service_proto_rawDescGZIP(), []int{3}
}

func (x *Permission) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type UserPermissions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permissions []*Permission `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *UserPermissions) Reset() {
	*x = UserPermissions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_auth_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPermissions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPermissions) ProtoMessage() {}

func (x *UserPermissions) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_auth_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPermissions.ProtoReflect.Descriptor instead.
func (*UserPermissions) Descriptor() ([]byte, []int) {
	return file_auth_service_auth_service_proto_rawDescGZIP(), []int{4}
}

func (x *UserPermissions) GetPermissions() []*Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type ValidateToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ValidateToken) Reset() {
	*x = ValidateToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_auth_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateToken) ProtoMessage() {}

func (x *ValidateToken) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_auth_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateToken.ProtoReflect.Descriptor instead.
func (*ValidateToken) Descriptor() ([]byte, []int) {
	return file_auth_service_auth_service_proto_rawDescGZIP(), []int{5}
}

func (x *ValidateToken) GetToken() *Token {
	if x != nil {
		return x.Token
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
		mi := &file_auth_service_auth_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_auth_service_proto_msgTypes[6]
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
	return file_auth_service_auth_service_proto_rawDescGZIP(), []int{6}
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
		mi := &file_auth_service_auth_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_auth_service_proto_msgTypes[7]
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
	return file_auth_service_auth_service_proto_rawDescGZIP(), []int{7}
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

	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname   string `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	Password  string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Telephone string `protobuf:"bytes,6,opt,name=telephone,proto3" json:"telephone,omitempty"`
	Gender    string `protobuf:"bytes,7,opt,name=gender,proto3" json:"gender,omitempty"`
	BirthDate string `protobuf:"bytes,8,opt,name=birthDate,proto3" json:"birthDate,omitempty"`
	Biography string `protobuf:"bytes,9,opt,name=biography,proto3" json:"biography,omitempty"`
	IsPublic  bool   `protobuf:"varint,17,opt,name=isPublic,proto3" json:"isPublic,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_auth_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_auth_service_proto_msgTypes[8]
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
	return file_auth_service_auth_service_proto_rawDescGZIP(), []int{8}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetSurname() string {
	if x != nil {
		return x.Surname
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

func (x *User) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

func (x *User) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *User) GetBirthDate() string {
	if x != nil {
		return x.BirthDate
	}
	return ""
}

func (x *User) GetBiography() string {
	if x != nil {
		return x.Biography
	}
	return ""
}

func (x *User) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

type CreateNewUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateNewUser) Reset() {
	*x = CreateNewUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_auth_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewUser) ProtoMessage() {}

func (x *CreateNewUser) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_auth_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewUser.ProtoReflect.Descriptor instead.
func (*CreateNewUser) Descriptor() ([]byte, []int) {
	return file_auth_service_auth_service_proto_rawDescGZIP(), []int{9}
}

func (x *CreateNewUser) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type PutNewUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *PutNewUser) Reset() {
	*x = PutNewUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_auth_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutNewUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutNewUser) ProtoMessage() {}

func (x *PutNewUser) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_auth_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutNewUser.ProtoReflect.Descriptor instead.
func (*PutNewUser) Descriptor() ([]byte, []int) {
	return file_auth_service_auth_service_proto_rawDescGZIP(), []int{10}
}

func (x *PutNewUser) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateNewUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateNewUserResponse) Reset() {
	*x = CreateNewUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_auth_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewUserResponse) ProtoMessage() {}

func (x *CreateNewUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_auth_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewUserResponse.ProtoReflect.Descriptor instead.
func (*CreateNewUserResponse) Descriptor() ([]byte, []int) {
	return file_auth_service_auth_service_proto_rawDescGZIP(), []int{11}
}

func (x *CreateNewUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_auth_service_auth_service_proto protoreflect.FileDescriptor

var file_auth_service_auth_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a,
	0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x59, 0x0a, 0x15, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x15, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x77, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e,
	0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x1d, 0x0a, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x22, 0x0a, 0x0a, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4d, 0x0a,
	0x0f, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x3a, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3a, 0x0a, 0x0d,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x29, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3a, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x90, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x62, 0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x62, 0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x37, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x34, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x26, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x3f, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xbc, 0x05, 0x0a, 0x0b, 0x41, 0x75, 0x74,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x4e,
	0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x55,
	0x73, 0x65, 0x72, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x8f, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x22, 0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31,
	0x1a, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x15, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x5c, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1b, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x5e, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x1a, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x76, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x4a, 0x57, 0x54, 0x12,
	0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x1b, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x26, 0x1a, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x3a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x7e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x1d, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x26, 0x1a, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x3a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x57, 0x53, 0x2d, 0x42, 0x53, 0x2d, 0x45, 0x50, 0x2d,
	0x54, 0x49, 0x4d, 0x32, 0x2d, 0x32, 0x30, 0x32, 0x32, 0x2f, 0x78, 0x77, 0x73, 0x62, 0x73, 0x2d,
	0x65, 0x70, 0x74, 0x69, 0x6d, 0x36, 0x2d, 0x32, 0x30, 0x32, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_service_auth_service_proto_rawDescOnce sync.Once
	file_auth_service_auth_service_proto_rawDescData = file_auth_service_auth_service_proto_rawDesc
)

func file_auth_service_auth_service_proto_rawDescGZIP() []byte {
	file_auth_service_auth_service_proto_rawDescOnce.Do(func() {
		file_auth_service_auth_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_service_auth_service_proto_rawDescData)
	})
	return file_auth_service_auth_service_proto_rawDescData
}

var file_auth_service_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_auth_service_auth_service_proto_goTypes = []interface{}{
	(*ChangePasswordRequest)(nil), // 0: auth_service.ChangePasswordRequest
	(*ChangePasswordMessage)(nil), // 1: auth_service.ChangePasswordMessage
	(*Token)(nil),                 // 2: auth_service.Token
	(*Permission)(nil),            // 3: auth_service.Permission
	(*UserPermissions)(nil),       // 4: auth_service.UserPermissions
	(*ValidateToken)(nil),         // 5: auth_service.ValidateToken
	(*GetAllRequest)(nil),         // 6: auth_service.GetAllRequest
	(*GetAllResponse)(nil),        // 7: auth_service.GetAllResponse
	(*User)(nil),                  // 8: auth_service.User
	(*CreateNewUser)(nil),         // 9: auth_service.CreateNewUser
	(*PutNewUser)(nil),            // 10: auth_service.PutNewUser
	(*CreateNewUserResponse)(nil), // 11: auth_service.CreateNewUserResponse
}
var file_auth_service_auth_service_proto_depIdxs = []int32{
	1,  // 0: auth_service.ChangePasswordRequest.changePasswordMessage:type_name -> auth_service.ChangePasswordMessage
	3,  // 1: auth_service.UserPermissions.permissions:type_name -> auth_service.Permission
	2,  // 2: auth_service.ValidateToken.token:type_name -> auth_service.Token
	8,  // 3: auth_service.GetAllResponse.users:type_name -> auth_service.User
	8,  // 4: auth_service.CreateNewUser.user:type_name -> auth_service.User
	8,  // 5: auth_service.PutNewUser.user:type_name -> auth_service.User
	8,  // 6: auth_service.CreateNewUserResponse.user:type_name -> auth_service.User
	9,  // 7: auth_service.AuthService.AddNewUser:input_type -> auth_service.CreateNewUser
	0,  // 8: auth_service.AuthService.UpdateUserPassword:input_type -> auth_service.ChangePasswordRequest
	6,  // 9: auth_service.AuthService.GetAll:input_type -> auth_service.GetAllRequest
	9,  // 10: auth_service.AuthService.LoginUser:input_type -> auth_service.CreateNewUser
	5,  // 11: auth_service.AuthService.AuthorizeJWT:input_type -> auth_service.ValidateToken
	5,  // 12: auth_service.AuthService.GetUserPermissions:input_type -> auth_service.ValidateToken
	9,  // 13: auth_service.AuthService.AddNewUser:output_type -> auth_service.CreateNewUser
	9,  // 14: auth_service.AuthService.UpdateUserPassword:output_type -> auth_service.CreateNewUser
	7,  // 15: auth_service.AuthService.GetAll:output_type -> auth_service.GetAllResponse
	2,  // 16: auth_service.AuthService.LoginUser:output_type -> auth_service.Token
	9,  // 17: auth_service.AuthService.AuthorizeJWT:output_type -> auth_service.CreateNewUser
	4,  // 18: auth_service.AuthService.GetUserPermissions:output_type -> auth_service.UserPermissions
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_auth_service_auth_service_proto_init() }
func file_auth_service_auth_service_proto_init() {
	if File_auth_service_auth_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_service_auth_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordRequest); i {
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
		file_auth_service_auth_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordMessage); i {
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
		file_auth_service_auth_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_auth_service_auth_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
		file_auth_service_auth_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPermissions); i {
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
		file_auth_service_auth_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateToken); i {
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
		file_auth_service_auth_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_auth_service_auth_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_auth_service_auth_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_auth_service_auth_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewUser); i {
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
		file_auth_service_auth_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutNewUser); i {
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
		file_auth_service_auth_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewUserResponse); i {
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
			RawDescriptor: file_auth_service_auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_service_auth_service_proto_goTypes,
		DependencyIndexes: file_auth_service_auth_service_proto_depIdxs,
		MessageInfos:      file_auth_service_auth_service_proto_msgTypes,
	}.Build()
	File_auth_service_auth_service_proto = out.File
	file_auth_service_auth_service_proto_rawDesc = nil
	file_auth_service_auth_service_proto_goTypes = nil
	file_auth_service_auth_service_proto_depIdxs = nil
}
