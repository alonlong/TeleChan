// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/api/api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e50ccc58c7b575d, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SignUpRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpRequest) Reset()         { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()    {}
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e50ccc58c7b575d, []int{1}
}

func (m *SignUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpRequest.Unmarshal(m, b)
}
func (m *SignUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpRequest.Marshal(b, m, deterministic)
}
func (m *SignUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpRequest.Merge(m, src)
}
func (m *SignUpRequest) XXX_Size() int {
	return xxx_messageInfo_SignUpRequest.Size(m)
}
func (m *SignUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpRequest proto.InternalMessageInfo

func (m *SignUpRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type SuccessReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuccessReply) Reset()         { *m = SuccessReply{} }
func (m *SuccessReply) String() string { return proto.CompactTextString(m) }
func (*SuccessReply) ProtoMessage()    {}
func (*SuccessReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e50ccc58c7b575d, []int{2}
}

func (m *SuccessReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuccessReply.Unmarshal(m, b)
}
func (m *SuccessReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuccessReply.Marshal(b, m, deterministic)
}
func (m *SuccessReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuccessReply.Merge(m, src)
}
func (m *SuccessReply) XXX_Size() int {
	return xxx_messageInfo_SuccessReply.Size(m)
}
func (m *SuccessReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SuccessReply.DiscardUnknown(m)
}

var xxx_messageInfo_SuccessReply proto.InternalMessageInfo

type SignInRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInRequest) Reset()         { *m = SignInRequest{} }
func (m *SignInRequest) String() string { return proto.CompactTextString(m) }
func (*SignInRequest) ProtoMessage()    {}
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e50ccc58c7b575d, []int{3}
}

func (m *SignInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInRequest.Unmarshal(m, b)
}
func (m *SignInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInRequest.Marshal(b, m, deterministic)
}
func (m *SignInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInRequest.Merge(m, src)
}
func (m *SignInRequest) XXX_Size() int {
	return xxx_messageInfo_SignInRequest.Size(m)
}
func (m *SignInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignInRequest proto.InternalMessageInfo

func (m *SignInRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type SignInReply struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInReply) Reset()         { *m = SignInReply{} }
func (m *SignInReply) String() string { return proto.CompactTextString(m) }
func (*SignInReply) ProtoMessage()    {}
func (*SignInReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e50ccc58c7b575d, []int{4}
}

func (m *SignInReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInReply.Unmarshal(m, b)
}
func (m *SignInReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInReply.Marshal(b, m, deterministic)
}
func (m *SignInReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInReply.Merge(m, src)
}
func (m *SignInReply) XXX_Size() int {
	return xxx_messageInfo_SignInReply.Size(m)
}
func (m *SignInReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInReply.DiscardUnknown(m)
}

var xxx_messageInfo_SignInReply proto.InternalMessageInfo

func (m *SignInReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RefreshRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshRequest) Reset()         { *m = RefreshRequest{} }
func (m *RefreshRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshRequest) ProtoMessage()    {}
func (*RefreshRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e50ccc58c7b575d, []int{5}
}

func (m *RefreshRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshRequest.Unmarshal(m, b)
}
func (m *RefreshRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshRequest.Marshal(b, m, deterministic)
}
func (m *RefreshRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshRequest.Merge(m, src)
}
func (m *RefreshRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshRequest.Size(m)
}
func (m *RefreshRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshRequest proto.InternalMessageInfo

type RefreshReply struct {
	RefreshToken         string   `protobuf:"bytes,1,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshReply) Reset()         { *m = RefreshReply{} }
func (m *RefreshReply) String() string { return proto.CompactTextString(m) }
func (*RefreshReply) ProtoMessage()    {}
func (*RefreshReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e50ccc58c7b575d, []int{6}
}

func (m *RefreshReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshReply.Unmarshal(m, b)
}
func (m *RefreshReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshReply.Marshal(b, m, deterministic)
}
func (m *RefreshReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshReply.Merge(m, src)
}
func (m *RefreshReply) XXX_Size() int {
	return xxx_messageInfo_RefreshReply.Size(m)
}
func (m *RefreshReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshReply.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshReply proto.InternalMessageInfo

func (m *RefreshReply) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type NewChanRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewChanRequest) Reset()         { *m = NewChanRequest{} }
func (m *NewChanRequest) String() string { return proto.CompactTextString(m) }
func (*NewChanRequest) ProtoMessage()    {}
func (*NewChanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e50ccc58c7b575d, []int{7}
}

func (m *NewChanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewChanRequest.Unmarshal(m, b)
}
func (m *NewChanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewChanRequest.Marshal(b, m, deterministic)
}
func (m *NewChanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewChanRequest.Merge(m, src)
}
func (m *NewChanRequest) XXX_Size() int {
	return xxx_messageInfo_NewChanRequest.Size(m)
}
func (m *NewChanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewChanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewChanRequest proto.InternalMessageInfo

func (m *NewChanRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type JoinChanRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinChanRequest) Reset()         { *m = JoinChanRequest{} }
func (m *JoinChanRequest) String() string { return proto.CompactTextString(m) }
func (*JoinChanRequest) ProtoMessage()    {}
func (*JoinChanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e50ccc58c7b575d, []int{8}
}

func (m *JoinChanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinChanRequest.Unmarshal(m, b)
}
func (m *JoinChanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinChanRequest.Marshal(b, m, deterministic)
}
func (m *JoinChanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinChanRequest.Merge(m, src)
}
func (m *JoinChanRequest) XXX_Size() int {
	return xxx_messageInfo_JoinChanRequest.Size(m)
}
func (m *JoinChanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinChanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinChanRequest proto.InternalMessageInfo

func (m *JoinChanRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Channel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owned                bool     `protobuf:"varint,3,opt,name=owned,proto3" json:"owned,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Channel) Reset()         { *m = Channel{} }
func (m *Channel) String() string { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()    {}
func (*Channel) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e50ccc58c7b575d, []int{9}
}

func (m *Channel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Channel.Unmarshal(m, b)
}
func (m *Channel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Channel.Marshal(b, m, deterministic)
}
func (m *Channel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Channel.Merge(m, src)
}
func (m *Channel) XXX_Size() int {
	return xxx_messageInfo_Channel.Size(m)
}
func (m *Channel) XXX_DiscardUnknown() {
	xxx_messageInfo_Channel.DiscardUnknown(m)
}

var xxx_messageInfo_Channel proto.InternalMessageInfo

func (m *Channel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Channel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Channel) GetOwned() bool {
	if m != nil {
		return m.Owned
	}
	return false
}

type GetChansRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetChansRequest) Reset()         { *m = GetChansRequest{} }
func (m *GetChansRequest) String() string { return proto.CompactTextString(m) }
func (*GetChansRequest) ProtoMessage()    {}
func (*GetChansRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e50ccc58c7b575d, []int{10}
}

func (m *GetChansRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetChansRequest.Unmarshal(m, b)
}
func (m *GetChansRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetChansRequest.Marshal(b, m, deterministic)
}
func (m *GetChansRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChansRequest.Merge(m, src)
}
func (m *GetChansRequest) XXX_Size() int {
	return xxx_messageInfo_GetChansRequest.Size(m)
}
func (m *GetChansRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChansRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetChansRequest proto.InternalMessageInfo

type SearchChansRequest struct {
	Input                string   `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchChansRequest) Reset()         { *m = SearchChansRequest{} }
func (m *SearchChansRequest) String() string { return proto.CompactTextString(m) }
func (*SearchChansRequest) ProtoMessage()    {}
func (*SearchChansRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e50ccc58c7b575d, []int{11}
}

func (m *SearchChansRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchChansRequest.Unmarshal(m, b)
}
func (m *SearchChansRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchChansRequest.Marshal(b, m, deterministic)
}
func (m *SearchChansRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchChansRequest.Merge(m, src)
}
func (m *SearchChansRequest) XXX_Size() int {
	return xxx_messageInfo_SearchChansRequest.Size(m)
}
func (m *SearchChansRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchChansRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchChansRequest proto.InternalMessageInfo

func (m *SearchChansRequest) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

type ChansReply struct {
	Channels             []*Channel `protobuf:"bytes,1,rep,name=channels,proto3" json:"channels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ChansReply) Reset()         { *m = ChansReply{} }
func (m *ChansReply) String() string { return proto.CompactTextString(m) }
func (*ChansReply) ProtoMessage()    {}
func (*ChansReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e50ccc58c7b575d, []int{12}
}

func (m *ChansReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChansReply.Unmarshal(m, b)
}
func (m *ChansReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChansReply.Marshal(b, m, deterministic)
}
func (m *ChansReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChansReply.Merge(m, src)
}
func (m *ChansReply) XXX_Size() int {
	return xxx_messageInfo_ChansReply.Size(m)
}
func (m *ChansReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ChansReply.DiscardUnknown(m)
}

var xxx_messageInfo_ChansReply proto.InternalMessageInfo

func (m *ChansReply) GetChannels() []*Channel {
	if m != nil {
		return m.Channels
	}
	return nil
}

type LeaveChanRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveChanRequest) Reset()         { *m = LeaveChanRequest{} }
func (m *LeaveChanRequest) String() string { return proto.CompactTextString(m) }
func (*LeaveChanRequest) ProtoMessage()    {}
func (*LeaveChanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e50ccc58c7b575d, []int{13}
}

func (m *LeaveChanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveChanRequest.Unmarshal(m, b)
}
func (m *LeaveChanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveChanRequest.Marshal(b, m, deterministic)
}
func (m *LeaveChanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveChanRequest.Merge(m, src)
}
func (m *LeaveChanRequest) XXX_Size() int {
	return xxx_messageInfo_LeaveChanRequest.Size(m)
}
func (m *LeaveChanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveChanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveChanRequest proto.InternalMessageInfo

func (m *LeaveChanRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ClientUpdate struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientUpdate) Reset()         { *m = ClientUpdate{} }
func (m *ClientUpdate) String() string { return proto.CompactTextString(m) }
func (*ClientUpdate) ProtoMessage()    {}
func (*ClientUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e50ccc58c7b575d, []int{14}
}

func (m *ClientUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientUpdate.Unmarshal(m, b)
}
func (m *ClientUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientUpdate.Marshal(b, m, deterministic)
}
func (m *ClientUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientUpdate.Merge(m, src)
}
func (m *ClientUpdate) XXX_Size() int {
	return xxx_messageInfo_ClientUpdate.Size(m)
}
func (m *ClientUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ClientUpdate proto.InternalMessageInfo

type ServerUpdate struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerUpdate) Reset()         { *m = ServerUpdate{} }
func (m *ServerUpdate) String() string { return proto.CompactTextString(m) }
func (*ServerUpdate) ProtoMessage()    {}
func (*ServerUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e50ccc58c7b575d, []int{15}
}

func (m *ServerUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerUpdate.Unmarshal(m, b)
}
func (m *ServerUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerUpdate.Marshal(b, m, deterministic)
}
func (m *ServerUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerUpdate.Merge(m, src)
}
func (m *ServerUpdate) XXX_Size() int {
	return xxx_messageInfo_ServerUpdate.Size(m)
}
func (m *ServerUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ServerUpdate proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "api.User")
	proto.RegisterType((*SignUpRequest)(nil), "api.SignUpRequest")
	proto.RegisterType((*SuccessReply)(nil), "api.SuccessReply")
	proto.RegisterType((*SignInRequest)(nil), "api.SignInRequest")
	proto.RegisterType((*SignInReply)(nil), "api.SignInReply")
	proto.RegisterType((*RefreshRequest)(nil), "api.RefreshRequest")
	proto.RegisterType((*RefreshReply)(nil), "api.RefreshReply")
	proto.RegisterType((*NewChanRequest)(nil), "api.NewChanRequest")
	proto.RegisterType((*JoinChanRequest)(nil), "api.JoinChanRequest")
	proto.RegisterType((*Channel)(nil), "api.Channel")
	proto.RegisterType((*GetChansRequest)(nil), "api.GetChansRequest")
	proto.RegisterType((*SearchChansRequest)(nil), "api.SearchChansRequest")
	proto.RegisterType((*ChansReply)(nil), "api.ChansReply")
	proto.RegisterType((*LeaveChanRequest)(nil), "api.LeaveChanRequest")
	proto.RegisterType((*ClientUpdate)(nil), "api.ClientUpdate")
	proto.RegisterType((*ServerUpdate)(nil), "api.ServerUpdate")
}

func init() { proto.RegisterFile("pkg/api/api.proto", fileDescriptor_7e50ccc58c7b575d) }

var fileDescriptor_7e50ccc58c7b575d = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdb, 0x6e, 0xda, 0x40,
	0x10, 0x8d, 0x81, 0x26, 0x66, 0xa0, 0x5c, 0xb6, 0xa9, 0x8a, 0x90, 0x2a, 0xa1, 0xed, 0x45, 0x56,
	0x1f, 0x68, 0x0a, 0xea, 0x25, 0xcf, 0x3c, 0x54, 0xa9, 0xaa, 0x3e, 0x98, 0xf0, 0x01, 0x5b, 0x7b,
	0x1a, 0x56, 0x31, 0xeb, 0xad, 0xd7, 0x0e, 0xca, 0x7f, 0xf7, 0x03, 0xaa, 0xbd, 0xd8, 0xb5, 0x29,
	0x52, 0xf2, 0x80, 0xc4, 0x99, 0x39, 0x67, 0x76, 0x76, 0xcf, 0x91, 0x61, 0x2c, 0x6f, 0x6f, 0xde,
	0x33, 0xc9, 0xf5, 0x6f, 0x2e, 0xb3, 0x34, 0x4f, 0x49, 0x9b, 0x49, 0x4e, 0xbf, 0x40, 0x67, 0xa3,
	0x30, 0x23, 0xe7, 0xf0, 0x04, 0x77, 0x8c, 0x27, 0x13, 0x6f, 0xe6, 0x05, 0xdd, 0xd0, 0x02, 0x32,
	0x05, 0x5f, 0x32, 0xa5, 0xf6, 0x69, 0x16, 0x4f, 0x5a, 0xa6, 0x51, 0x61, 0x3a, 0x87, 0xa7, 0x6b,
	0x7e, 0x23, 0x36, 0x32, 0xc4, 0xdf, 0x05, 0xaa, 0x9c, 0xbc, 0x84, 0x4e, 0xa1, 0x30, 0x33, 0x13,
	0x7a, 0x8b, 0xee, 0x5c, 0x9f, 0xa4, 0x67, 0x87, 0xa6, 0x4c, 0x07, 0xd0, 0x5f, 0x17, 0x51, 0x84,
	0x4a, 0x85, 0x28, 0x93, 0xfb, 0x52, 0x7f, 0x25, 0x1e, 0xa9, 0x7f, 0x05, 0xbd, 0x92, 0x2f, 0x93,
	0x7b, 0xbd, 0x70, 0x9e, 0xde, 0xa2, 0x28, 0x17, 0x36, 0x80, 0x8e, 0x60, 0x10, 0xe2, 0xaf, 0x0c,
	0xd5, 0xd6, 0x4d, 0xa5, 0x0b, 0xe8, 0x57, 0x15, 0xad, 0xa3, 0xd0, 0xcf, 0x2c, 0xbe, 0xae, 0xc9,
	0x1b, 0x35, 0xfa, 0x1a, 0x06, 0x3f, 0x70, 0xbf, 0xda, 0xb2, 0x6a, 0x37, 0x02, 0x1d, 0xc1, 0x76,
	0xe8, 0xd8, 0xe6, 0x3f, 0x7d, 0x03, 0xc3, 0x6f, 0x29, 0x17, 0x0f, 0xd1, 0x56, 0x70, 0xa6, 0x29,
	0x02, 0x13, 0x32, 0x80, 0x16, 0x8f, 0x5d, 0xb3, 0xc5, 0xe3, 0x8a, 0xde, 0xfa, 0x47, 0xd7, 0xf7,
	0x4a, 0xf7, 0x02, 0xe3, 0x49, 0x7b, 0xe6, 0x05, 0x7e, 0x68, 0x01, 0x1d, 0xc3, 0xf0, 0x2b, 0xe6,
	0x7a, 0x8e, 0x2a, 0x2f, 0xf6, 0x0e, 0xc8, 0x1a, 0x59, 0x16, 0x6d, 0xeb, 0x55, 0x2d, 0xe7, 0x42,
	0x16, 0x79, 0xf9, 0x2c, 0x06, 0xd0, 0x4f, 0x00, 0x8e, 0xa5, 0x9f, 0x20, 0x00, 0x3f, 0xb2, 0x1b,
	0xa9, 0x89, 0x37, 0x6b, 0x07, 0xbd, 0x45, 0xdf, 0x3c, 0xb6, 0x5b, 0x33, 0xac, 0xba, 0xf4, 0x2d,
	0x8c, 0xbe, 0x23, 0xbb, 0xc3, 0x87, 0xee, 0x38, 0x80, 0xfe, 0x2a, 0xe1, 0x28, 0xf2, 0x8d, 0x8c,
	0x59, 0x6e, 0xf0, 0x1a, 0xb3, 0x3b, 0xcc, 0x2c, 0x5e, 0xfc, 0x69, 0x83, 0x7f, 0x8d, 0x89, 0x99,
	0x43, 0x3e, 0xc0, 0xa9, 0x0d, 0x0e, 0x21, 0xe6, 0xd8, 0x46, 0x8a, 0xa6, 0x63, 0x5b, 0xab, 0x27,
	0xe5, 0x84, 0x5c, 0x58, 0xc9, 0x95, 0xa8, 0x49, 0xaa, 0xe0, 0x4c, 0x47, 0x8d, 0x9a, 0x55, 0x2c,
	0xe1, 0xcc, 0xd9, 0x4e, 0x9e, 0x99, 0x76, 0x33, 0x16, 0xee, 0x98, 0x7a, 0x32, 0xac, 0xc8, 0xf9,
	0xee, 0x44, 0xcd, 0x14, 0x1c, 0xdf, 0xed, 0x23, 0xf8, 0x65, 0x0c, 0xc8, 0xb9, 0x21, 0x1c, 0xa4,
	0xe2, 0xb8, 0x6c, 0x09, 0x7e, 0xe9, 0xa8, 0x93, 0x1d, 0x18, 0x3c, 0x1d, 0x56, 0xa6, 0x54, 0xa2,
	0x4b, 0xe8, 0xd5, 0x3c, 0x27, 0x2f, 0xec, 0xe0, 0xff, 0x52, 0x70, 0x4c, 0xfa, 0x19, 0xba, 0x95,
	0x95, 0xe4, 0xb9, 0xe9, 0x1f, 0x5a, 0x7b, 0x7c, 0xd1, 0x4b, 0xe8, 0xad, 0xd2, 0xdd, 0xae, 0x10,
	0x3c, 0x62, 0x39, 0x12, 0xcb, 0xa9, 0xbb, 0x5d, 0xca, 0x6a, 0x86, 0xd3, 0x93, 0xc0, 0xbb, 0xf0,
	0x7e, 0x9e, 0x9a, 0x0f, 0xcd, 0xf2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0x29, 0x8e, 0x80,
	0x7d, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TeleChanClient is the client API for TeleChan service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TeleChanClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SuccessReply, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInReply, error)
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshReply, error)
	NewChan(ctx context.Context, in *NewChanRequest, opts ...grpc.CallOption) (*SuccessReply, error)
	JoinChan(ctx context.Context, in *JoinChanRequest, opts ...grpc.CallOption) (*SuccessReply, error)
	GetChans(ctx context.Context, in *GetChansRequest, opts ...grpc.CallOption) (*ChansReply, error)
	SearchChans(ctx context.Context, in *SearchChansRequest, opts ...grpc.CallOption) (*ChansReply, error)
	LeaveChan(ctx context.Context, in *LeaveChanRequest, opts ...grpc.CallOption) (*SuccessReply, error)
	Communicate(ctx context.Context, opts ...grpc.CallOption) (TeleChan_CommunicateClient, error)
}

type teleChanClient struct {
	cc *grpc.ClientConn
}

func NewTeleChanClient(cc *grpc.ClientConn) TeleChanClient {
	return &teleChanClient{cc}
}

func (c *teleChanClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SuccessReply, error) {
	out := new(SuccessReply)
	err := c.cc.Invoke(ctx, "/api.TeleChan/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teleChanClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInReply, error) {
	out := new(SignInReply)
	err := c.cc.Invoke(ctx, "/api.TeleChan/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teleChanClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshReply, error) {
	out := new(RefreshReply)
	err := c.cc.Invoke(ctx, "/api.TeleChan/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teleChanClient) NewChan(ctx context.Context, in *NewChanRequest, opts ...grpc.CallOption) (*SuccessReply, error) {
	out := new(SuccessReply)
	err := c.cc.Invoke(ctx, "/api.TeleChan/NewChan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teleChanClient) JoinChan(ctx context.Context, in *JoinChanRequest, opts ...grpc.CallOption) (*SuccessReply, error) {
	out := new(SuccessReply)
	err := c.cc.Invoke(ctx, "/api.TeleChan/JoinChan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teleChanClient) GetChans(ctx context.Context, in *GetChansRequest, opts ...grpc.CallOption) (*ChansReply, error) {
	out := new(ChansReply)
	err := c.cc.Invoke(ctx, "/api.TeleChan/GetChans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teleChanClient) SearchChans(ctx context.Context, in *SearchChansRequest, opts ...grpc.CallOption) (*ChansReply, error) {
	out := new(ChansReply)
	err := c.cc.Invoke(ctx, "/api.TeleChan/SearchChans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teleChanClient) LeaveChan(ctx context.Context, in *LeaveChanRequest, opts ...grpc.CallOption) (*SuccessReply, error) {
	out := new(SuccessReply)
	err := c.cc.Invoke(ctx, "/api.TeleChan/LeaveChan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teleChanClient) Communicate(ctx context.Context, opts ...grpc.CallOption) (TeleChan_CommunicateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TeleChan_serviceDesc.Streams[0], "/api.TeleChan/Communicate", opts...)
	if err != nil {
		return nil, err
	}
	x := &teleChanCommunicateClient{stream}
	return x, nil
}

type TeleChan_CommunicateClient interface {
	Send(*ClientUpdate) error
	Recv() (*ServerUpdate, error)
	grpc.ClientStream
}

type teleChanCommunicateClient struct {
	grpc.ClientStream
}

func (x *teleChanCommunicateClient) Send(m *ClientUpdate) error {
	return x.ClientStream.SendMsg(m)
}

func (x *teleChanCommunicateClient) Recv() (*ServerUpdate, error) {
	m := new(ServerUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TeleChanServer is the server API for TeleChan service.
type TeleChanServer interface {
	SignUp(context.Context, *SignUpRequest) (*SuccessReply, error)
	SignIn(context.Context, *SignInRequest) (*SignInReply, error)
	Refresh(context.Context, *RefreshRequest) (*RefreshReply, error)
	NewChan(context.Context, *NewChanRequest) (*SuccessReply, error)
	JoinChan(context.Context, *JoinChanRequest) (*SuccessReply, error)
	GetChans(context.Context, *GetChansRequest) (*ChansReply, error)
	SearchChans(context.Context, *SearchChansRequest) (*ChansReply, error)
	LeaveChan(context.Context, *LeaveChanRequest) (*SuccessReply, error)
	Communicate(TeleChan_CommunicateServer) error
}

// UnimplementedTeleChanServer can be embedded to have forward compatible implementations.
type UnimplementedTeleChanServer struct {
}

func (*UnimplementedTeleChanServer) SignUp(ctx context.Context, req *SignUpRequest) (*SuccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (*UnimplementedTeleChanServer) SignIn(ctx context.Context, req *SignInRequest) (*SignInReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (*UnimplementedTeleChanServer) Refresh(ctx context.Context, req *RefreshRequest) (*RefreshReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (*UnimplementedTeleChanServer) NewChan(ctx context.Context, req *NewChanRequest) (*SuccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewChan not implemented")
}
func (*UnimplementedTeleChanServer) JoinChan(ctx context.Context, req *JoinChanRequest) (*SuccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinChan not implemented")
}
func (*UnimplementedTeleChanServer) GetChans(ctx context.Context, req *GetChansRequest) (*ChansReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChans not implemented")
}
func (*UnimplementedTeleChanServer) SearchChans(ctx context.Context, req *SearchChansRequest) (*ChansReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchChans not implemented")
}
func (*UnimplementedTeleChanServer) LeaveChan(ctx context.Context, req *LeaveChanRequest) (*SuccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveChan not implemented")
}
func (*UnimplementedTeleChanServer) Communicate(srv TeleChan_CommunicateServer) error {
	return status.Errorf(codes.Unimplemented, "method Communicate not implemented")
}

func RegisterTeleChanServer(s *grpc.Server, srv TeleChanServer) {
	s.RegisterService(&_TeleChan_serviceDesc, srv)
}

func _TeleChan_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeleChanServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TeleChan/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeleChanServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeleChan_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeleChanServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TeleChan/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeleChanServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeleChan_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeleChanServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TeleChan/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeleChanServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeleChan_NewChan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewChanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeleChanServer).NewChan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TeleChan/NewChan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeleChanServer).NewChan(ctx, req.(*NewChanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeleChan_JoinChan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinChanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeleChanServer).JoinChan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TeleChan/JoinChan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeleChanServer).JoinChan(ctx, req.(*JoinChanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeleChan_GetChans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeleChanServer).GetChans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TeleChan/GetChans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeleChanServer).GetChans(ctx, req.(*GetChansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeleChan_SearchChans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchChansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeleChanServer).SearchChans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TeleChan/SearchChans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeleChanServer).SearchChans(ctx, req.(*SearchChansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeleChan_LeaveChan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveChanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeleChanServer).LeaveChan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TeleChan/LeaveChan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeleChanServer).LeaveChan(ctx, req.(*LeaveChanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeleChan_Communicate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TeleChanServer).Communicate(&teleChanCommunicateServer{stream})
}

type TeleChan_CommunicateServer interface {
	Send(*ServerUpdate) error
	Recv() (*ClientUpdate, error)
	grpc.ServerStream
}

type teleChanCommunicateServer struct {
	grpc.ServerStream
}

func (x *teleChanCommunicateServer) Send(m *ServerUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func (x *teleChanCommunicateServer) Recv() (*ClientUpdate, error) {
	m := new(ClientUpdate)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TeleChan_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.TeleChan",
	HandlerType: (*TeleChanServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _TeleChan_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _TeleChan_SignIn_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _TeleChan_Refresh_Handler,
		},
		{
			MethodName: "NewChan",
			Handler:    _TeleChan_NewChan_Handler,
		},
		{
			MethodName: "JoinChan",
			Handler:    _TeleChan_JoinChan_Handler,
		},
		{
			MethodName: "GetChans",
			Handler:    _TeleChan_GetChans_Handler,
		},
		{
			MethodName: "SearchChans",
			Handler:    _TeleChan_SearchChans_Handler,
		},
		{
			MethodName: "LeaveChan",
			Handler:    _TeleChan_LeaveChan_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Communicate",
			Handler:       _TeleChan_Communicate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/api/api.proto",
}
