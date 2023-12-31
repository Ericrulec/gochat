// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: gochat.proto

package proto

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

type User_UserStatusCode int32

const (
	User_UNKNOWN_STATUS User_UserStatusCode = 0
	User_OFFLINE        User_UserStatusCode = 1
	User_ONLINE         User_UserStatusCode = 2
)

// Enum value maps for User_UserStatusCode.
var (
	User_UserStatusCode_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "OFFLINE",
		2: "ONLINE",
	}
	User_UserStatusCode_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"OFFLINE":        1,
		"ONLINE":         2,
	}
)

func (x User_UserStatusCode) Enum() *User_UserStatusCode {
	p := new(User_UserStatusCode)
	*p = x
	return p
}

func (x User_UserStatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (User_UserStatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_gochat_proto_enumTypes[0].Descriptor()
}

func (User_UserStatusCode) Type() protoreflect.EnumType {
	return &file_gochat_proto_enumTypes[0]
}

func (x User_UserStatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User_UserStatusCode.Descriptor instead.
func (User_UserStatusCode) EnumDescriptor() ([]byte, []int) {
	return file_gochat_proto_rawDescGZIP(), []int{1, 0}
}

type ChatMessageStatus_MessageStatus int32

const (
	ChatMessageStatus_UNKNOWN_STATUS ChatMessageStatus_MessageStatus = 0
	ChatMessageStatus_NOT_DELIVERED  ChatMessageStatus_MessageStatus = 1
	ChatMessageStatus_DELIVERED      ChatMessageStatus_MessageStatus = 2
)

// Enum value maps for ChatMessageStatus_MessageStatus.
var (
	ChatMessageStatus_MessageStatus_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "NOT_DELIVERED",
		2: "DELIVERED",
	}
	ChatMessageStatus_MessageStatus_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"NOT_DELIVERED":  1,
		"DELIVERED":      2,
	}
)

func (x ChatMessageStatus_MessageStatus) Enum() *ChatMessageStatus_MessageStatus {
	p := new(ChatMessageStatus_MessageStatus)
	*p = x
	return p
}

func (x ChatMessageStatus_MessageStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatMessageStatus_MessageStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_gochat_proto_enumTypes[1].Descriptor()
}

func (ChatMessageStatus_MessageStatus) Type() protoreflect.EnumType {
	return &file_gochat_proto_enumTypes[1]
}

func (x ChatMessageStatus_MessageStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatMessageStatus_MessageStatus.Descriptor instead.
func (ChatMessageStatus_MessageStatus) EnumDescriptor() ([]byte, []int) {
	return file_gochat_proto_rawDescGZIP(), []int{5, 0}
}

type ChatMessageResponse_MessageStatus int32

const (
	ChatMessageResponse_UNKNOWN_STATUS ChatMessageResponse_MessageStatus = 0
	ChatMessageResponse_NOT_DELIVERED  ChatMessageResponse_MessageStatus = 1
	ChatMessageResponse_DELIVERED      ChatMessageResponse_MessageStatus = 2
)

// Enum value maps for ChatMessageResponse_MessageStatus.
var (
	ChatMessageResponse_MessageStatus_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "NOT_DELIVERED",
		2: "DELIVERED",
	}
	ChatMessageResponse_MessageStatus_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"NOT_DELIVERED":  1,
		"DELIVERED":      2,
	}
)

func (x ChatMessageResponse_MessageStatus) Enum() *ChatMessageResponse_MessageStatus {
	p := new(ChatMessageResponse_MessageStatus)
	*p = x
	return p
}

func (x ChatMessageResponse_MessageStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatMessageResponse_MessageStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_gochat_proto_enumTypes[2].Descriptor()
}

func (ChatMessageResponse_MessageStatus) Type() protoreflect.EnumType {
	return &file_gochat_proto_enumTypes[2]
}

func (x ChatMessageResponse_MessageStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatMessageResponse_MessageStatus.Descriptor instead.
func (ChatMessageResponse_MessageStatus) EnumDescriptor() ([]byte, []int) {
	return file_gochat_proto_rawDescGZIP(), []int{6, 0}
}

type QueryUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *QueryUser) Reset() {
	*x = QueryUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gochat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUser) ProtoMessage() {}

func (x *QueryUser) ProtoReflect() protoreflect.Message {
	mi := &file_gochat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUser.ProtoReflect.Descriptor instead.
func (*QueryUser) Descriptor() ([]byte, []int) {
	return file_gochat_proto_rawDescGZIP(), []int{0}
}

func (x *QueryUser) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string              `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name   string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status User_UserStatusCode `protobuf:"varint,3,opt,name=status,proto3,enum=proto.gochat.User_UserStatusCode" json:"status,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gochat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_gochat_proto_msgTypes[1]
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
	return file_gochat_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetStatus() User_UserStatusCode {
	if x != nil {
		return x.Status
	}
	return User_UNKNOWN_STATUS
}

type ConnectUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Device string `protobuf:"bytes,3,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *ConnectUserRequest) Reset() {
	*x = ConnectUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gochat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectUserRequest) ProtoMessage() {}

func (x *ConnectUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gochat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectUserRequest.ProtoReflect.Descriptor instead.
func (*ConnectUserRequest) Descriptor() ([]byte, []int) {
	return file_gochat_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ConnectUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConnectUserRequest) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

type DisconnectUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Device string `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *DisconnectUserRequest) Reset() {
	*x = DisconnectUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gochat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectUserRequest) ProtoMessage() {}

func (x *DisconnectUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gochat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectUserRequest.ProtoReflect.Descriptor instead.
func (*DisconnectUserRequest) Descriptor() ([]byte, []int) {
	return file_gochat_proto_rawDescGZIP(), []int{3}
}

func (x *DisconnectUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *DisconnectUserRequest) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId    string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	FromUser  *User  `protobuf:"bytes,2,opt,name=from_user,json=fromUser,proto3" json:"from_user,omitempty"`
	ToUser    *User  `protobuf:"bytes,3,opt,name=to_user,json=toUser,proto3" json:"to_user,omitempty"`
	Body      string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	TimeStamp string `protobuf:"bytes,5,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gochat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_gochat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_gochat_proto_rawDescGZIP(), []int{4}
}

func (x *ChatMessage) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *ChatMessage) GetFromUser() *User {
	if x != nil {
		return x.FromUser
	}
	return nil
}

func (x *ChatMessage) GetToUser() *User {
	if x != nil {
		return x.ToUser
	}
	return nil
}

func (x *ChatMessage) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *ChatMessage) GetTimeStamp() string {
	if x != nil {
		return x.TimeStamp
	}
	return ""
}

type ChatMessageStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatMessage *ChatMessage                    `protobuf:"bytes,1,opt,name=chat_message,json=chatMessage,proto3" json:"chat_message,omitempty"`
	MsgStatus   ChatMessageStatus_MessageStatus `protobuf:"varint,2,opt,name=msg_status,json=msgStatus,proto3,enum=proto.gochat.ChatMessageStatus_MessageStatus" json:"msg_status,omitempty"`
}

func (x *ChatMessageStatus) Reset() {
	*x = ChatMessageStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gochat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessageStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessageStatus) ProtoMessage() {}

func (x *ChatMessageStatus) ProtoReflect() protoreflect.Message {
	mi := &file_gochat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessageStatus.ProtoReflect.Descriptor instead.
func (*ChatMessageStatus) Descriptor() ([]byte, []int) {
	return file_gochat_proto_rawDescGZIP(), []int{5}
}

func (x *ChatMessageStatus) GetChatMessage() *ChatMessage {
	if x != nil {
		return x.ChatMessage
	}
	return nil
}

func (x *ChatMessageStatus) GetMsgStatus() ChatMessageStatus_MessageStatus {
	if x != nil {
		return x.MsgStatus
	}
	return ChatMessageStatus_UNKNOWN_STATUS
}

type ChatMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatMessage *ChatMessage                      `protobuf:"bytes,1,opt,name=chat_message,json=chatMessage,proto3" json:"chat_message,omitempty"`
	MsgStatus   ChatMessageResponse_MessageStatus `protobuf:"varint,2,opt,name=msg_status,json=msgStatus,proto3,enum=proto.gochat.ChatMessageResponse_MessageStatus" json:"msg_status,omitempty"`
}

func (x *ChatMessageResponse) Reset() {
	*x = ChatMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gochat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessageResponse) ProtoMessage() {}

func (x *ChatMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gochat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessageResponse.ProtoReflect.Descriptor instead.
func (*ChatMessageResponse) Descriptor() ([]byte, []int) {
	return file_gochat_proto_rawDescGZIP(), []int{6}
}

func (x *ChatMessageResponse) GetChatMessage() *ChatMessage {
	if x != nil {
		return x.ChatMessage
	}
	return nil
}

func (x *ChatMessageResponse) GetMsgStatus() ChatMessageResponse_MessageStatus {
	if x != nil {
		return x.MsgStatus
	}
	return ChatMessageResponse_UNKNOWN_STATUS
}

type QueryMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User1          *User  `protobuf:"bytes,1,opt,name=user1,proto3" json:"user1,omitempty"`
	User2          *User  `protobuf:"bytes,2,opt,name=user2,proto3" json:"user2,omitempty"`
	TimeConstraint string `protobuf:"bytes,3,opt,name=time_constraint,json=timeConstraint,proto3" json:"time_constraint,omitempty"`
	Limit          int32  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *QueryMessagesRequest) Reset() {
	*x = QueryMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gochat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMessagesRequest) ProtoMessage() {}

func (x *QueryMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gochat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMessagesRequest.ProtoReflect.Descriptor instead.
func (*QueryMessagesRequest) Descriptor() ([]byte, []int) {
	return file_gochat_proto_rawDescGZIP(), []int{7}
}

func (x *QueryMessagesRequest) GetUser1() *User {
	if x != nil {
		return x.User1
	}
	return nil
}

func (x *QueryMessagesRequest) GetUser2() *User {
	if x != nil {
		return x.User2
	}
	return nil
}

func (x *QueryMessagesRequest) GetTimeConstraint() string {
	if x != nil {
		return x.TimeConstraint
	}
	return ""
}

func (x *QueryMessagesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type QueryMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User1    *User          `protobuf:"bytes,1,opt,name=user1,proto3" json:"user1,omitempty"`
	User2    *User          `protobuf:"bytes,2,opt,name=user2,proto3" json:"user2,omitempty"`
	Messages []*ChatMessage `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *QueryMessagesResponse) Reset() {
	*x = QueryMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gochat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMessagesResponse) ProtoMessage() {}

func (x *QueryMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gochat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMessagesResponse.ProtoReflect.Descriptor instead.
func (*QueryMessagesResponse) Descriptor() ([]byte, []int) {
	return file_gochat_proto_rawDescGZIP(), []int{8}
}

func (x *QueryMessagesResponse) GetUser1() *User {
	if x != nil {
		return x.User1
	}
	return nil
}

func (x *QueryMessagesResponse) GetUser2() *User {
	if x != nil {
		return x.User2
	}
	return nil
}

func (x *QueryMessagesResponse) GetMessages() []*ChatMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_gochat_proto protoreflect.FileDescriptor

var file_gochat_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x22, 0x24, 0x0a, 0x09,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0xad, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x3d, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x46, 0x46,
	0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45,
	0x10, 0x02, 0x22, 0x59, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x57, 0x0a,
	0x15, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0xb7, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12,
	0x2f, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x2b, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0xe6, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3c, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x6d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x45, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x44,
	0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45,
	0x4c, 0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x02, 0x22, 0xea, 0x01, 0x0a, 0x13, 0x43, 0x68,
	0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x4e, 0x0a, 0x0a, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x6d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x45, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x49,
	0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x4c, 0x49, 0x56,
	0x45, 0x52, 0x45, 0x44, 0x10, 0x02, 0x22, 0xa9, 0x01, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x28, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x31, 0x12, 0x28, 0x0a, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x32, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x69,
	0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0xa2, 0x01, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x31, 0x12, 0x28, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x32, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x32,
	0x12, 0x35, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x32, 0xb6, 0x04, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x30, 0x01, 0x12, 0x49, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x50, 0x0a,
	0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x21, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x57, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67,
	0x6f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x56, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45,
	0x72, 0x69, 0x63, 0x72, 0x75, 0x6c, 0x65, 0x63, 0x2f, 0x67, 0x6f, 0x63, 0x68, 0x61, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gochat_proto_rawDescOnce sync.Once
	file_gochat_proto_rawDescData = file_gochat_proto_rawDesc
)

func file_gochat_proto_rawDescGZIP() []byte {
	file_gochat_proto_rawDescOnce.Do(func() {
		file_gochat_proto_rawDescData = protoimpl.X.CompressGZIP(file_gochat_proto_rawDescData)
	})
	return file_gochat_proto_rawDescData
}

var file_gochat_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_gochat_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_gochat_proto_goTypes = []interface{}{
	(User_UserStatusCode)(0),               // 0: proto.gochat.User.UserStatusCode
	(ChatMessageStatus_MessageStatus)(0),   // 1: proto.gochat.ChatMessageStatus.MessageStatus
	(ChatMessageResponse_MessageStatus)(0), // 2: proto.gochat.ChatMessageResponse.MessageStatus
	(*QueryUser)(nil),                      // 3: proto.gochat.QueryUser
	(*User)(nil),                           // 4: proto.gochat.User
	(*ConnectUserRequest)(nil),             // 5: proto.gochat.ConnectUserRequest
	(*DisconnectUserRequest)(nil),          // 6: proto.gochat.DisconnectUserRequest
	(*ChatMessage)(nil),                    // 7: proto.gochat.ChatMessage
	(*ChatMessageStatus)(nil),              // 8: proto.gochat.ChatMessageStatus
	(*ChatMessageResponse)(nil),            // 9: proto.gochat.ChatMessageResponse
	(*QueryMessagesRequest)(nil),           // 10: proto.gochat.QueryMessagesRequest
	(*QueryMessagesResponse)(nil),          // 11: proto.gochat.QueryMessagesResponse
}
var file_gochat_proto_depIdxs = []int32{
	0,  // 0: proto.gochat.User.status:type_name -> proto.gochat.User.UserStatusCode
	4,  // 1: proto.gochat.DisconnectUserRequest.user:type_name -> proto.gochat.User
	4,  // 2: proto.gochat.ChatMessage.from_user:type_name -> proto.gochat.User
	4,  // 3: proto.gochat.ChatMessage.to_user:type_name -> proto.gochat.User
	7,  // 4: proto.gochat.ChatMessageStatus.chat_message:type_name -> proto.gochat.ChatMessage
	1,  // 5: proto.gochat.ChatMessageStatus.msg_status:type_name -> proto.gochat.ChatMessageStatus.MessageStatus
	7,  // 6: proto.gochat.ChatMessageResponse.chat_message:type_name -> proto.gochat.ChatMessage
	2,  // 7: proto.gochat.ChatMessageResponse.msg_status:type_name -> proto.gochat.ChatMessageResponse.MessageStatus
	4,  // 8: proto.gochat.QueryMessagesRequest.user1:type_name -> proto.gochat.User
	4,  // 9: proto.gochat.QueryMessagesRequest.user2:type_name -> proto.gochat.User
	4,  // 10: proto.gochat.QueryMessagesResponse.user1:type_name -> proto.gochat.User
	4,  // 11: proto.gochat.QueryMessagesResponse.user2:type_name -> proto.gochat.User
	7,  // 12: proto.gochat.QueryMessagesResponse.messages:type_name -> proto.gochat.ChatMessage
	3,  // 13: proto.gochat.MessagingService.GetUser:input_type -> proto.gochat.QueryUser
	5,  // 14: proto.gochat.MessagingService.ConnectUser:input_type -> proto.gochat.ConnectUserRequest
	6,  // 15: proto.gochat.MessagingService.DisconnectUser:input_type -> proto.gochat.DisconnectUserRequest
	7,  // 16: proto.gochat.MessagingService.BroadcastMessage:input_type -> proto.gochat.ChatMessage
	8,  // 17: proto.gochat.MessagingService.UpdateMessageStatus:input_type -> proto.gochat.ChatMessageStatus
	7,  // 18: proto.gochat.MessagingService.GetMessageStatus:input_type -> proto.gochat.ChatMessage
	10, // 19: proto.gochat.MessagingService.GetMessages:input_type -> proto.gochat.QueryMessagesRequest
	4,  // 20: proto.gochat.MessagingService.GetUser:output_type -> proto.gochat.User
	7,  // 21: proto.gochat.MessagingService.ConnectUser:output_type -> proto.gochat.ChatMessage
	4,  // 22: proto.gochat.MessagingService.DisconnectUser:output_type -> proto.gochat.User
	9,  // 23: proto.gochat.MessagingService.BroadcastMessage:output_type -> proto.gochat.ChatMessageResponse
	8,  // 24: proto.gochat.MessagingService.UpdateMessageStatus:output_type -> proto.gochat.ChatMessageStatus
	8,  // 25: proto.gochat.MessagingService.GetMessageStatus:output_type -> proto.gochat.ChatMessageStatus
	11, // 26: proto.gochat.MessagingService.GetMessages:output_type -> proto.gochat.QueryMessagesResponse
	20, // [20:27] is the sub-list for method output_type
	13, // [13:20] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_gochat_proto_init() }
func file_gochat_proto_init() {
	if File_gochat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gochat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUser); i {
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
		file_gochat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_gochat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectUserRequest); i {
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
		file_gochat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectUserRequest); i {
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
		file_gochat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
		file_gochat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessageStatus); i {
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
		file_gochat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessageResponse); i {
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
		file_gochat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMessagesRequest); i {
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
		file_gochat_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMessagesResponse); i {
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
			RawDescriptor: file_gochat_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gochat_proto_goTypes,
		DependencyIndexes: file_gochat_proto_depIdxs,
		EnumInfos:         file_gochat_proto_enumTypes,
		MessageInfos:      file_gochat_proto_msgTypes,
	}.Build()
	File_gochat_proto = out.File
	file_gochat_proto_rawDesc = nil
	file_gochat_proto_goTypes = nil
	file_gochat_proto_depIdxs = nil
}
