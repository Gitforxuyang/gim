// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gim.proto

package gim

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

//消息类型
type MessageAction int32

const (
	MessageAction_Unknown  MessageAction = 0
	MessageAction_Text     MessageAction = 1
	MessageAction_Image    MessageAction = 2
	MessageAction_Sound    MessageAction = 3
	MessageAction_Video    MessageAction = 4
	MessageAction_Face     MessageAction = 5
	MessageAction_Location MessageAction = 6
	MessageAction_File     MessageAction = 7
	MessageAction_Gift     MessageAction = 8
	//    Notify = 9; //通知
	MessageAction_Push MessageAction = 10
)

var MessageAction_name = map[int32]string{
	0:  "Unknown",
	1:  "Text",
	2:  "Image",
	3:  "Sound",
	4:  "Video",
	5:  "Face",
	6:  "Location",
	7:  "File",
	8:  "Gift",
	10: "Push",
}

var MessageAction_value = map[string]int32{
	"Unknown":  0,
	"Text":     1,
	"Image":    2,
	"Sound":    3,
	"Video":    4,
	"Face":     5,
	"Location": 6,
	"File":     7,
	"Gift":     8,
	"Push":     10,
}

func (x MessageAction) String() string {
	return proto.EnumName(MessageAction_name, int32(x))
}

func (MessageAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{0}
}

type Ping struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{0}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

type Pong struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{1}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

type AuthReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Uid                  int64    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	SdkVersion           string   `protobuf:"bytes,4,opt,name=sdkVersion,proto3" json:"sdkVersion,omitempty"`
	DeviceId             string   `protobuf:"bytes,5,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	Platform             string   `protobuf:"bytes,6,opt,name=platform,proto3" json:"platform,omitempty"`
	Model                string   `protobuf:"bytes,7,opt,name=model,proto3" json:"model,omitempty"`
	System               string   `protobuf:"bytes,8,opt,name=system,proto3" json:"system,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthReq) Reset()         { *m = AuthReq{} }
func (m *AuthReq) String() string { return proto.CompactTextString(m) }
func (*AuthReq) ProtoMessage()    {}
func (*AuthReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{2}
}

func (m *AuthReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthReq.Unmarshal(m, b)
}
func (m *AuthReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthReq.Marshal(b, m, deterministic)
}
func (m *AuthReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthReq.Merge(m, src)
}
func (m *AuthReq) XXX_Size() int {
	return xxx_messageInfo_AuthReq.Size(m)
}
func (m *AuthReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthReq.DiscardUnknown(m)
}

var xxx_messageInfo_AuthReq proto.InternalMessageInfo

func (m *AuthReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *AuthReq) GetSdkVersion() string {
	if m != nil {
		return m.SdkVersion
	}
	return ""
}

func (m *AuthReq) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *AuthReq) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *AuthReq) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *AuthReq) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

type AuthRes struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRes) Reset()         { *m = AuthRes{} }
func (m *AuthRes) String() string { return proto.CompactTextString(m) }
func (*AuthRes) ProtoMessage()    {}
func (*AuthRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{3}
}

func (m *AuthRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRes.Unmarshal(m, b)
}
func (m *AuthRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRes.Marshal(b, m, deterministic)
}
func (m *AuthRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRes.Merge(m, src)
}
func (m *AuthRes) XXX_Size() int {
	return xxx_messageInfo_AuthRes.Size(m)
}
func (m *AuthRes) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRes.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRes proto.InternalMessageInfo

func (m *AuthRes) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AuthRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

//客户端发消息
type SendMessageReq struct {
	Seq int64 `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	//消息类型  1-单聊 2-群聊 3-聊天室
	//群聊跟聊天室的区别  群聊可以加入多个，且不在线情况下仍可收到消息  聊天室只能加入一个，且如果不在线则收不到消息
	Type                 int32         `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Action               MessageAction `protobuf:"varint,3,opt,name=action,proto3,enum=gim.MessageAction" json:"action,omitempty"`
	From                 int64         `protobuf:"varint,4,opt,name=from,proto3" json:"from,omitempty"`
	To                   int64         `protobuf:"varint,5,opt,name=to,proto3" json:"to,omitempty"`
	Content              string        `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SendMessageReq) Reset()         { *m = SendMessageReq{} }
func (m *SendMessageReq) String() string { return proto.CompactTextString(m) }
func (*SendMessageReq) ProtoMessage()    {}
func (*SendMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{4}
}

func (m *SendMessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageReq.Unmarshal(m, b)
}
func (m *SendMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageReq.Marshal(b, m, deterministic)
}
func (m *SendMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageReq.Merge(m, src)
}
func (m *SendMessageReq) XXX_Size() int {
	return xxx_messageInfo_SendMessageReq.Size(m)
}
func (m *SendMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageReq proto.InternalMessageInfo

func (m *SendMessageReq) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *SendMessageReq) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *SendMessageReq) GetAction() MessageAction {
	if m != nil {
		return m.Action
	}
	return MessageAction_Unknown
}

func (m *SendMessageReq) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *SendMessageReq) GetTo() int64 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *SendMessageReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type SendMessageResp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Seq                  int64    `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`
	MsgId                int64    `protobuf:"varint,4,opt,name=msgId,proto3" json:"msgId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageResp) Reset()         { *m = SendMessageResp{} }
func (m *SendMessageResp) String() string { return proto.CompactTextString(m) }
func (*SendMessageResp) ProtoMessage()    {}
func (*SendMessageResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{5}
}

func (m *SendMessageResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageResp.Unmarshal(m, b)
}
func (m *SendMessageResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageResp.Marshal(b, m, deterministic)
}
func (m *SendMessageResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageResp.Merge(m, src)
}
func (m *SendMessageResp) XXX_Size() int {
	return xxx_messageInfo_SendMessageResp.Size(m)
}
func (m *SendMessageResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageResp.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageResp proto.InternalMessageInfo

func (m *SendMessageResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SendMessageResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *SendMessageResp) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *SendMessageResp) GetMsgId() int64 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

//通知客户端有新消息
type Notify struct {
	MsgId                int64    `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notify) Reset()         { *m = Notify{} }
func (m *Notify) String() string { return proto.CompactTextString(m) }
func (*Notify) ProtoMessage()    {}
func (*Notify) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{6}
}

func (m *Notify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notify.Unmarshal(m, b)
}
func (m *Notify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notify.Marshal(b, m, deterministic)
}
func (m *Notify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notify.Merge(m, src)
}
func (m *Notify) XXX_Size() int {
	return xxx_messageInfo_Notify.Size(m)
}
func (m *Notify) XXX_DiscardUnknown() {
	xxx_messageInfo_Notify.DiscardUnknown(m)
}

var xxx_messageInfo_Notify proto.InternalMessageInfo

func (m *Notify) GetMsgId() int64 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

type NotifyAck struct {
	MsgId                int64    `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyAck) Reset()         { *m = NotifyAck{} }
func (m *NotifyAck) String() string { return proto.CompactTextString(m) }
func (*NotifyAck) ProtoMessage()    {}
func (*NotifyAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{7}
}

func (m *NotifyAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyAck.Unmarshal(m, b)
}
func (m *NotifyAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyAck.Marshal(b, m, deterministic)
}
func (m *NotifyAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyAck.Merge(m, src)
}
func (m *NotifyAck) XXX_Size() int {
	return xxx_messageInfo_NotifyAck.Size(m)
}
func (m *NotifyAck) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyAck.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyAck proto.InternalMessageInfo

func (m *NotifyAck) GetMsgId() int64 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

type SyncMessageReq struct {
	LastId               int64    `protobuf:"varint,1,opt,name=lastId,proto3" json:"lastId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncMessageReq) Reset()         { *m = SyncMessageReq{} }
func (m *SyncMessageReq) String() string { return proto.CompactTextString(m) }
func (*SyncMessageReq) ProtoMessage()    {}
func (*SyncMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{8}
}

func (m *SyncMessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncMessageReq.Unmarshal(m, b)
}
func (m *SyncMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncMessageReq.Marshal(b, m, deterministic)
}
func (m *SyncMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncMessageReq.Merge(m, src)
}
func (m *SyncMessageReq) XXX_Size() int {
	return xxx_messageInfo_SyncMessageReq.Size(m)
}
func (m *SyncMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_SyncMessageReq proto.InternalMessageInfo

func (m *SyncMessageReq) GetLastId() int64 {
	if m != nil {
		return m.LastId
	}
	return 0
}

type MessageBody struct {
	Type                 int32         `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Action               MessageAction `protobuf:"varint,2,opt,name=action,proto3,enum=gim.MessageAction" json:"action,omitempty"`
	From                 int64         `protobuf:"varint,3,opt,name=from,proto3" json:"from,omitempty"`
	To                   int64         `protobuf:"varint,4,opt,name=to,proto3" json:"to,omitempty"`
	Content              string        `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MessageBody) Reset()         { *m = MessageBody{} }
func (m *MessageBody) String() string { return proto.CompactTextString(m) }
func (*MessageBody) ProtoMessage()    {}
func (*MessageBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{9}
}

func (m *MessageBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageBody.Unmarshal(m, b)
}
func (m *MessageBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageBody.Marshal(b, m, deterministic)
}
func (m *MessageBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageBody.Merge(m, src)
}
func (m *MessageBody) XXX_Size() int {
	return xxx_messageInfo_MessageBody.Size(m)
}
func (m *MessageBody) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageBody.DiscardUnknown(m)
}

var xxx_messageInfo_MessageBody proto.InternalMessageInfo

func (m *MessageBody) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MessageBody) GetAction() MessageAction {
	if m != nil {
		return m.Action
	}
	return MessageAction_Unknown
}

func (m *MessageBody) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *MessageBody) GetTo() int64 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *MessageBody) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type SyncMessageResp struct {
	Code                 int32          `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string         `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Msgs                 []*MessageBody `protobuf:"bytes,3,rep,name=msgs,proto3" json:"msgs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SyncMessageResp) Reset()         { *m = SyncMessageResp{} }
func (m *SyncMessageResp) String() string { return proto.CompactTextString(m) }
func (*SyncMessageResp) ProtoMessage()    {}
func (*SyncMessageResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{10}
}

func (m *SyncMessageResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncMessageResp.Unmarshal(m, b)
}
func (m *SyncMessageResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncMessageResp.Marshal(b, m, deterministic)
}
func (m *SyncMessageResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncMessageResp.Merge(m, src)
}
func (m *SyncMessageResp) XXX_Size() int {
	return xxx_messageInfo_SyncMessageResp.Size(m)
}
func (m *SyncMessageResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncMessageResp.DiscardUnknown(m)
}

var xxx_messageInfo_SyncMessageResp proto.InternalMessageInfo

func (m *SyncMessageResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SyncMessageResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *SyncMessageResp) GetMsgs() []*MessageBody {
	if m != nil {
		return m.Msgs
	}
	return nil
}

type FetchMessageReq struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	MsgId                int64    `protobuf:"varint,2,opt,name=msgId,proto3" json:"msgId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchMessageReq) Reset()         { *m = FetchMessageReq{} }
func (m *FetchMessageReq) String() string { return proto.CompactTextString(m) }
func (*FetchMessageReq) ProtoMessage()    {}
func (*FetchMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{11}
}

func (m *FetchMessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchMessageReq.Unmarshal(m, b)
}
func (m *FetchMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchMessageReq.Marshal(b, m, deterministic)
}
func (m *FetchMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchMessageReq.Merge(m, src)
}
func (m *FetchMessageReq) XXX_Size() int {
	return xxx_messageInfo_FetchMessageReq.Size(m)
}
func (m *FetchMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_FetchMessageReq proto.InternalMessageInfo

func (m *FetchMessageReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *FetchMessageReq) GetMsgId() int64 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

type FetchMessageResp struct {
	Code                 int32          `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string         `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Msgs                 []*MessageBody `protobuf:"bytes,3,rep,name=msgs,proto3" json:"msgs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FetchMessageResp) Reset()         { *m = FetchMessageResp{} }
func (m *FetchMessageResp) String() string { return proto.CompactTextString(m) }
func (*FetchMessageResp) ProtoMessage()    {}
func (*FetchMessageResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{12}
}

func (m *FetchMessageResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchMessageResp.Unmarshal(m, b)
}
func (m *FetchMessageResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchMessageResp.Marshal(b, m, deterministic)
}
func (m *FetchMessageResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchMessageResp.Merge(m, src)
}
func (m *FetchMessageResp) XXX_Size() int {
	return xxx_messageInfo_FetchMessageResp.Size(m)
}
func (m *FetchMessageResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchMessageResp.DiscardUnknown(m)
}

var xxx_messageInfo_FetchMessageResp proto.InternalMessageInfo

func (m *FetchMessageResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *FetchMessageResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *FetchMessageResp) GetMsgs() []*MessageBody {
	if m != nil {
		return m.Msgs
	}
	return nil
}

type SyncLastIdReq struct {
	LastId               int64    `protobuf:"varint,1,opt,name=lastId,proto3" json:"lastId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncLastIdReq) Reset()         { *m = SyncLastIdReq{} }
func (m *SyncLastIdReq) String() string { return proto.CompactTextString(m) }
func (*SyncLastIdReq) ProtoMessage()    {}
func (*SyncLastIdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{13}
}

func (m *SyncLastIdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncLastIdReq.Unmarshal(m, b)
}
func (m *SyncLastIdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncLastIdReq.Marshal(b, m, deterministic)
}
func (m *SyncLastIdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncLastIdReq.Merge(m, src)
}
func (m *SyncLastIdReq) XXX_Size() int {
	return xxx_messageInfo_SyncLastIdReq.Size(m)
}
func (m *SyncLastIdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncLastIdReq.DiscardUnknown(m)
}

var xxx_messageInfo_SyncLastIdReq proto.InternalMessageInfo

func (m *SyncLastIdReq) GetLastId() int64 {
	if m != nil {
		return m.LastId
	}
	return 0
}

type SyncLastIdResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncLastIdResp) Reset()         { *m = SyncLastIdResp{} }
func (m *SyncLastIdResp) String() string { return proto.CompactTextString(m) }
func (*SyncLastIdResp) ProtoMessage()    {}
func (*SyncLastIdResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{14}
}

func (m *SyncLastIdResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncLastIdResp.Unmarshal(m, b)
}
func (m *SyncLastIdResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncLastIdResp.Marshal(b, m, deterministic)
}
func (m *SyncLastIdResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncLastIdResp.Merge(m, src)
}
func (m *SyncLastIdResp) XXX_Size() int {
	return xxx_messageInfo_SyncLastIdResp.Size(m)
}
func (m *SyncLastIdResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncLastIdResp.DiscardUnknown(m)
}

var xxx_messageInfo_SyncLastIdResp proto.InternalMessageInfo

type KickOutNotify struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KickOutNotify) Reset()         { *m = KickOutNotify{} }
func (m *KickOutNotify) String() string { return proto.CompactTextString(m) }
func (*KickOutNotify) ProtoMessage()    {}
func (*KickOutNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{15}
}

func (m *KickOutNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KickOutNotify.Unmarshal(m, b)
}
func (m *KickOutNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KickOutNotify.Marshal(b, m, deterministic)
}
func (m *KickOutNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KickOutNotify.Merge(m, src)
}
func (m *KickOutNotify) XXX_Size() int {
	return xxx_messageInfo_KickOutNotify.Size(m)
}
func (m *KickOutNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_KickOutNotify.DiscardUnknown(m)
}

var xxx_messageInfo_KickOutNotify proto.InternalMessageInfo

func (m *KickOutNotify) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *KickOutNotify) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type LogoutReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutReq) Reset()         { *m = LogoutReq{} }
func (m *LogoutReq) String() string { return proto.CompactTextString(m) }
func (*LogoutReq) ProtoMessage()    {}
func (*LogoutReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{16}
}

func (m *LogoutReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutReq.Unmarshal(m, b)
}
func (m *LogoutReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutReq.Marshal(b, m, deterministic)
}
func (m *LogoutReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutReq.Merge(m, src)
}
func (m *LogoutReq) XXX_Size() int {
	return xxx_messageInfo_LogoutReq.Size(m)
}
func (m *LogoutReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutReq.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutReq proto.InternalMessageInfo

type LogoutResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResp) Reset()         { *m = LogoutResp{} }
func (m *LogoutResp) String() string { return proto.CompactTextString(m) }
func (*LogoutResp) ProtoMessage()    {}
func (*LogoutResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{17}
}

func (m *LogoutResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResp.Unmarshal(m, b)
}
func (m *LogoutResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResp.Marshal(b, m, deterministic)
}
func (m *LogoutResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResp.Merge(m, src)
}
func (m *LogoutResp) XXX_Size() int {
	return xxx_messageInfo_LogoutResp.Size(m)
}
func (m *LogoutResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResp.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResp proto.InternalMessageInfo

type BaseResp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseResp) Reset()         { *m = BaseResp{} }
func (m *BaseResp) String() string { return proto.CompactTextString(m) }
func (*BaseResp) ProtoMessage()    {}
func (*BaseResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5533614fb47f913, []int{18}
}

func (m *BaseResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseResp.Unmarshal(m, b)
}
func (m *BaseResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseResp.Marshal(b, m, deterministic)
}
func (m *BaseResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseResp.Merge(m, src)
}
func (m *BaseResp) XXX_Size() int {
	return xxx_messageInfo_BaseResp.Size(m)
}
func (m *BaseResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseResp.DiscardUnknown(m)
}

var xxx_messageInfo_BaseResp proto.InternalMessageInfo

func (m *BaseResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BaseResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterEnum("gim.MessageAction", MessageAction_name, MessageAction_value)
	proto.RegisterType((*Ping)(nil), "gim.Ping")
	proto.RegisterType((*Pong)(nil), "gim.Pong")
	proto.RegisterType((*AuthReq)(nil), "gim.AuthReq")
	proto.RegisterType((*AuthRes)(nil), "gim.AuthRes")
	proto.RegisterType((*SendMessageReq)(nil), "gim.SendMessageReq")
	proto.RegisterType((*SendMessageResp)(nil), "gim.SendMessageResp")
	proto.RegisterType((*Notify)(nil), "gim.Notify")
	proto.RegisterType((*NotifyAck)(nil), "gim.NotifyAck")
	proto.RegisterType((*SyncMessageReq)(nil), "gim.SyncMessageReq")
	proto.RegisterType((*MessageBody)(nil), "gim.MessageBody")
	proto.RegisterType((*SyncMessageResp)(nil), "gim.SyncMessageResp")
	proto.RegisterType((*FetchMessageReq)(nil), "gim.FetchMessageReq")
	proto.RegisterType((*FetchMessageResp)(nil), "gim.FetchMessageResp")
	proto.RegisterType((*SyncLastIdReq)(nil), "gim.SyncLastIdReq")
	proto.RegisterType((*SyncLastIdResp)(nil), "gim.SyncLastIdResp")
	proto.RegisterType((*KickOutNotify)(nil), "gim.KickOutNotify")
	proto.RegisterType((*LogoutReq)(nil), "gim.LogoutReq")
	proto.RegisterType((*LogoutResp)(nil), "gim.LogoutResp")
	proto.RegisterType((*BaseResp)(nil), "gim.BaseResp")
}

func init() { proto.RegisterFile("gim.proto", fileDescriptor_e5533614fb47f913) }

var fileDescriptor_e5533614fb47f913 = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xd1, 0x6e, 0xd3, 0x3c,
	0x18, 0xfd, 0x13, 0x27, 0x69, 0xf2, 0x75, 0xeb, 0x2c, 0xeb, 0xd7, 0x14, 0x71, 0x31, 0x95, 0x08,
	0x89, 0x6a, 0x17, 0x03, 0x0d, 0x71, 0xc1, 0xe5, 0x76, 0x31, 0x34, 0x51, 0x60, 0x72, 0x61, 0x77,
	0x80, 0x42, 0xe2, 0xa6, 0x51, 0x1b, 0x3b, 0xad, 0x1d, 0xa0, 0x17, 0x3c, 0x01, 0x2f, 0xc1, 0x5b,
	0xf0, 0x7a, 0xc8, 0x4e, 0xd2, 0x36, 0x62, 0xa0, 0x22, 0x71, 0xd5, 0x73, 0xfc, 0x7d, 0x9f, 0x73,
	0xce, 0xb1, 0x5d, 0x08, 0xb2, 0xbc, 0x38, 0x2b, 0x57, 0x42, 0x09, 0x82, 0xb2, 0xbc, 0x88, 0x3c,
	0x70, 0x6e, 0x72, 0x9e, 0x99, 0x5f, 0xc1, 0xb3, 0xe8, 0x87, 0x05, 0xbd, 0x8b, 0x4a, 0xcd, 0x28,
	0x5b, 0x92, 0xff, 0xc1, 0x55, 0x62, 0xce, 0x78, 0x68, 0x0d, 0xad, 0x51, 0x40, 0x6b, 0x42, 0x30,
	0xa0, 0x2a, 0x4f, 0x43, 0x7b, 0x68, 0x8d, 0x10, 0xd5, 0x90, 0x9c, 0x00, 0xc8, 0x74, 0x7e, 0xcb,
	0x56, 0x32, 0x17, 0x3c, 0x74, 0x4c, 0xf3, 0xce, 0x0a, 0xb9, 0x07, 0x7e, 0xca, 0x3e, 0xe5, 0x09,
	0xbb, 0x4e, 0x43, 0xd7, 0x54, 0x37, 0x5c, 0xd7, 0xca, 0x45, 0xac, 0xa6, 0x62, 0x55, 0x84, 0x5e,
	0x5d, 0x6b, 0xb9, 0xfe, 0x7e, 0x21, 0x52, 0xb6, 0x08, 0x7b, 0xf5, 0xf7, 0x0d, 0x21, 0xc7, 0xe0,
	0xc9, 0xb5, 0x54, 0xac, 0x08, 0x7d, 0xb3, 0xdc, 0xb0, 0xe8, 0x51, 0x2b, 0x5c, 0x12, 0x02, 0x4e,
	0x22, 0x52, 0x66, 0x74, 0xbb, 0xd4, 0x60, 0x2d, 0xbb, 0x90, 0x99, 0x91, 0x1d, 0x50, 0x0d, 0xa3,
	0xef, 0x16, 0x0c, 0x26, 0x8c, 0xa7, 0x2f, 0x99, 0x94, 0x71, 0xc6, 0xb4, 0x63, 0x0c, 0x48, 0xb2,
	0xa5, 0x99, 0x43, 0x54, 0x43, 0xbd, 0x95, 0x5a, 0x97, 0xcc, 0xcc, 0xb9, 0xd4, 0x60, 0x72, 0x0a,
	0x5e, 0x9c, 0x28, 0xed, 0x15, 0x0d, 0xad, 0xd1, 0xe0, 0x9c, 0x9c, 0xe9, 0x50, 0x9b, 0x6d, 0x2e,
	0x4c, 0x85, 0x36, 0x1d, 0x7a, 0x7e, 0xba, 0x12, 0x85, 0x49, 0x05, 0x51, 0x83, 0xc9, 0x00, 0x6c,
	0x25, 0x4c, 0x12, 0x88, 0xda, 0x4a, 0x90, 0x10, 0x7a, 0x89, 0xe0, 0x8a, 0x71, 0xd5, 0x44, 0xd0,
	0xd2, 0xe8, 0x03, 0x1c, 0x75, 0x14, 0xca, 0x72, 0x3f, 0x6f, 0xad, 0x11, 0xb4, 0x35, 0xa2, 0xc3,
	0x94, 0xd9, 0x75, 0xda, 0x28, 0xa9, 0x49, 0x74, 0x02, 0xde, 0x2b, 0xa1, 0xf2, 0xe9, 0x7a, 0x5b,
	0xb7, 0x76, 0xeb, 0xf7, 0x21, 0xa8, 0xeb, 0x17, 0xc9, 0xfc, 0x37, 0x2d, 0x23, 0x18, 0x4c, 0xd6,
	0x3c, 0xd9, 0x49, 0xf1, 0x18, 0xbc, 0x45, 0x2c, 0xd5, 0xa6, 0xb1, 0x61, 0xd1, 0x37, 0x0b, 0xfa,
	0x4d, 0xdb, 0xa5, 0x48, 0xd7, 0x9b, 0x6c, 0xad, 0x3b, 0xb3, 0xb5, 0xf7, 0xce, 0x16, 0xfd, 0x92,
	0xad, 0x73, 0x57, 0xb6, 0x6e, 0x37, 0xdb, 0x77, 0x70, 0xd4, 0xd1, 0xbd, 0x77, 0xb6, 0x0f, 0xc0,
	0x29, 0x64, 0x26, 0x43, 0x34, 0x44, 0xa3, 0xfe, 0x39, 0xde, 0x15, 0xa8, 0x6d, 0x51, 0x53, 0x8d,
	0x9e, 0xc1, 0xd1, 0x15, 0x53, 0xc9, 0xac, 0x7b, 0xbb, 0xf4, 0xcb, 0xb1, 0xb6, 0x2f, 0x67, 0x93,
	0xa8, 0xbd, 0x9b, 0xe8, 0x7b, 0xc0, 0xdd, 0xd1, 0x7f, 0x2c, 0xed, 0x21, 0x1c, 0x6a, 0xe7, 0x63,
	0x73, 0x2a, 0x7f, 0x3a, 0x30, 0x5c, 0x1f, 0x6d, 0xdb, 0x28, 0xcb, 0xe8, 0x29, 0x1c, 0xbe, 0xc8,
	0x93, 0xf9, 0xeb, 0x4a, 0x35, 0xd7, 0x66, 0xbf, 0xa7, 0xd6, 0x87, 0x60, 0x2c, 0x32, 0x51, 0x29,
	0xca, 0x96, 0xd1, 0x01, 0x40, 0x4b, 0x64, 0x19, 0x3d, 0x06, 0xff, 0x32, 0x96, 0x7f, 0x61, 0xf2,
	0xf4, 0x2b, 0x1c, 0x76, 0xee, 0x03, 0xe9, 0x43, 0xef, 0x2d, 0x9f, 0x73, 0xf1, 0x99, 0xe3, 0xff,
	0x88, 0x0f, 0xce, 0x1b, 0xf6, 0x45, 0x61, 0x8b, 0x04, 0xe0, 0x5e, 0x17, 0x71, 0xc6, 0xb0, 0xad,
	0xe1, 0x44, 0x54, 0x3c, 0xc5, 0x48, 0xc3, 0xdb, 0x3c, 0x65, 0x02, 0x3b, 0xba, 0xf5, 0x2a, 0x4e,
	0x18, 0x76, 0xc9, 0x01, 0xf8, 0x63, 0x91, 0xc4, 0x7a, 0x37, 0xec, 0x99, 0xf5, 0x7c, 0xc1, 0x70,
	0x4f, 0xa3, 0xe7, 0xf9, 0x54, 0x61, 0x5f, 0xa3, 0x9b, 0x4a, 0xce, 0x30, 0x7c, 0xf4, 0xcc, 0xbf,
	0xe7, 0x93, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x44, 0x34, 0xf6, 0x94, 0x4a, 0x05, 0x00, 0x00,
}
