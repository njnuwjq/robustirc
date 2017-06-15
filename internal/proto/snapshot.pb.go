// Code generated by protoc-gen-go.
// source: snapshot.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Bool represents a boolean which can either be unset, true or
// false. This quirk is necessary because proto3 does not distinguish
// the default value (false) from a field not being set, which would
// not allow us to do upgrades from older versions of RobustIRC to
// newer versions of RobustIRC which introduce a new boolean field.
type Bool int32

const (
	Bool_UNSET Bool = 0
	Bool_TRUE  Bool = 1
	Bool_FALSE Bool = 2
)

var Bool_name = map[int32]string{
	0: "UNSET",
	1: "TRUE",
	2: "FALSE",
}
var Bool_value = map[string]int32{
	"UNSET": 0,
	"TRUE":  1,
	"FALSE": 2,
}

func (x Bool) String() string {
	return proto1.EnumName(Bool_name, int32(x))
}
func (Bool) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Timestamp serializes a Go time.Time value with correct IsZero()
// semantics. Merely serializing the UnixNano() value is not
// sufficient, see https://play.golang.org/p/n3ZWGwZCKR
type Timestamp struct {
	UnixNano int64 `protobuf:"varint,1,opt,name=unix_nano,json=unixNano" json:"unix_nano,omitempty"`
	IsZero   bool  `protobuf:"varint,2,opt,name=is_zero,json=isZero" json:"is_zero,omitempty"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto1.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Timestamp) GetUnixNano() int64 {
	if m != nil {
		return m.UnixNano
	}
	return 0
}

func (m *Timestamp) GetIsZero() bool {
	if m != nil {
		return m.IsZero
	}
	return false
}

// Snapshot contains the entire state of an IRCServer object, so that
// a new IRCServer object can be created with exactly the same state.
type Snapshot struct {
	Sessions      []*Snapshot_Session          `protobuf:"bytes,1,rep,name=sessions" json:"sessions,omitempty"`
	Channels      []*Snapshot_Channel          `protobuf:"bytes,2,rep,name=channels" json:"channels,omitempty"`
	Svsholds      map[string]*Snapshot_SVSHold `protobuf:"bytes,3,rep,name=svsholds" json:"svsholds,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	LastProcessed *RobustId                    `protobuf:"bytes,4,opt,name=last_processed,json=lastProcessed" json:"last_processed,omitempty"`
	Config        *Snapshot_Config             `protobuf:"bytes,5,opt,name=config" json:"config,omitempty"`
	// last_included_index is the last ircstore message index which was
	// included when taking the snapshot. This is relevant to store this
	// snapshot in fsm.lastSnapshotState when restoring after ircstore
	// was deleted.
	LastIncludedIndex uint64 `protobuf:"varint,6,opt,name=last_included_index,json=lastIncludedIndex" json:"last_included_index,omitempty"`
}

func (m *Snapshot) Reset()                    { *m = Snapshot{} }
func (m *Snapshot) String() string            { return proto1.CompactTextString(m) }
func (*Snapshot) ProtoMessage()               {}
func (*Snapshot) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Snapshot) GetSessions() []*Snapshot_Session {
	if m != nil {
		return m.Sessions
	}
	return nil
}

func (m *Snapshot) GetChannels() []*Snapshot_Channel {
	if m != nil {
		return m.Channels
	}
	return nil
}

func (m *Snapshot) GetSvsholds() map[string]*Snapshot_SVSHold {
	if m != nil {
		return m.Svsholds
	}
	return nil
}

func (m *Snapshot) GetLastProcessed() *RobustId {
	if m != nil {
		return m.LastProcessed
	}
	return nil
}

func (m *Snapshot) GetConfig() *Snapshot_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Snapshot) GetLastIncludedIndex() uint64 {
	if m != nil {
		return m.LastIncludedIndex
	}
	return 0
}

type Snapshot_IRCPrefix struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	User string `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Host string `protobuf:"bytes,3,opt,name=host" json:"host,omitempty"`
}

func (m *Snapshot_IRCPrefix) Reset()                    { *m = Snapshot_IRCPrefix{} }
func (m *Snapshot_IRCPrefix) String() string            { return proto1.CompactTextString(m) }
func (*Snapshot_IRCPrefix) ProtoMessage()               {}
func (*Snapshot_IRCPrefix) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

func (m *Snapshot_IRCPrefix) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Snapshot_IRCPrefix) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Snapshot_IRCPrefix) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type Snapshot_Session struct {
	Id                  *RobustId           `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Auth                string              `protobuf:"bytes,2,opt,name=auth" json:"auth,omitempty"`
	Nick                string              `protobuf:"bytes,3,opt,name=nick" json:"nick,omitempty"`
	Username            string              `protobuf:"bytes,4,opt,name=username" json:"username,omitempty"`
	Realname            string              `protobuf:"bytes,5,opt,name=realname" json:"realname,omitempty"`
	Channels            []string            `protobuf:"bytes,6,rep,name=channels" json:"channels,omitempty"`
	LastActivity        *Timestamp          `protobuf:"bytes,7,opt,name=last_activity,json=lastActivity" json:"last_activity,omitempty"`
	Operator            bool                `protobuf:"varint,8,opt,name=operator" json:"operator,omitempty"`
	AwayMsg             string              `protobuf:"bytes,9,opt,name=away_msg,json=awayMsg" json:"away_msg,omitempty"`
	Created             int64               `protobuf:"varint,22,opt,name=created" json:"created,omitempty"`
	ThrottlingExponent  int64               `protobuf:"varint,10,opt,name=throttling_exponent,json=throttlingExponent" json:"throttling_exponent,omitempty"`
	InvitedTo           []string            `protobuf:"bytes,11,rep,name=invited_to,json=invitedTo" json:"invited_to,omitempty"`
	Modes               []string            `protobuf:"bytes,12,rep,name=modes" json:"modes,omitempty"`
	Svid                string              `protobuf:"bytes,13,opt,name=svid" json:"svid,omitempty"`
	Pass                string              `protobuf:"bytes,14,opt,name=pass" json:"pass,omitempty"`
	Server              bool                `protobuf:"varint,15,opt,name=server" json:"server,omitempty"`
	LastClientMessageId uint64              `protobuf:"varint,17,opt,name=last_client_message_id,json=lastClientMessageId" json:"last_client_message_id,omitempty"`
	IrcPrefix           *Snapshot_IRCPrefix `protobuf:"bytes,18,opt,name=irc_prefix,json=ircPrefix" json:"irc_prefix,omitempty"`
	LastNonPing         *Timestamp          `protobuf:"bytes,19,opt,name=last_non_ping,json=lastNonPing" json:"last_non_ping,omitempty"`
	LastSolvedCaptcha   *Timestamp          `protobuf:"bytes,20,opt,name=last_solved_captcha,json=lastSolvedCaptcha" json:"last_solved_captcha,omitempty"`
	LoggedIn            Bool                `protobuf:"varint,21,opt,name=logged_in,json=loggedIn,enum=proto.Bool" json:"logged_in,omitempty"`
	RemoteAddr          string              `protobuf:"bytes,23,opt,name=remote_addr,json=remoteAddr" json:"remote_addr,omitempty"`
}

func (m *Snapshot_Session) Reset()                    { *m = Snapshot_Session{} }
func (m *Snapshot_Session) String() string            { return proto1.CompactTextString(m) }
func (*Snapshot_Session) ProtoMessage()               {}
func (*Snapshot_Session) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 1} }

func (m *Snapshot_Session) GetId() *RobustId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Snapshot_Session) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

func (m *Snapshot_Session) GetNick() string {
	if m != nil {
		return m.Nick
	}
	return ""
}

func (m *Snapshot_Session) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Snapshot_Session) GetRealname() string {
	if m != nil {
		return m.Realname
	}
	return ""
}

func (m *Snapshot_Session) GetChannels() []string {
	if m != nil {
		return m.Channels
	}
	return nil
}

func (m *Snapshot_Session) GetLastActivity() *Timestamp {
	if m != nil {
		return m.LastActivity
	}
	return nil
}

func (m *Snapshot_Session) GetOperator() bool {
	if m != nil {
		return m.Operator
	}
	return false
}

func (m *Snapshot_Session) GetAwayMsg() string {
	if m != nil {
		return m.AwayMsg
	}
	return ""
}

func (m *Snapshot_Session) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Snapshot_Session) GetThrottlingExponent() int64 {
	if m != nil {
		return m.ThrottlingExponent
	}
	return 0
}

func (m *Snapshot_Session) GetInvitedTo() []string {
	if m != nil {
		return m.InvitedTo
	}
	return nil
}

func (m *Snapshot_Session) GetModes() []string {
	if m != nil {
		return m.Modes
	}
	return nil
}

func (m *Snapshot_Session) GetSvid() string {
	if m != nil {
		return m.Svid
	}
	return ""
}

func (m *Snapshot_Session) GetPass() string {
	if m != nil {
		return m.Pass
	}
	return ""
}

func (m *Snapshot_Session) GetServer() bool {
	if m != nil {
		return m.Server
	}
	return false
}

func (m *Snapshot_Session) GetLastClientMessageId() uint64 {
	if m != nil {
		return m.LastClientMessageId
	}
	return 0
}

func (m *Snapshot_Session) GetIrcPrefix() *Snapshot_IRCPrefix {
	if m != nil {
		return m.IrcPrefix
	}
	return nil
}

func (m *Snapshot_Session) GetLastNonPing() *Timestamp {
	if m != nil {
		return m.LastNonPing
	}
	return nil
}

func (m *Snapshot_Session) GetLastSolvedCaptcha() *Timestamp {
	if m != nil {
		return m.LastSolvedCaptcha
	}
	return nil
}

func (m *Snapshot_Session) GetLoggedIn() Bool {
	if m != nil {
		return m.LoggedIn
	}
	return Bool_UNSET
}

func (m *Snapshot_Session) GetRemoteAddr() string {
	if m != nil {
		return m.RemoteAddr
	}
	return ""
}

type Snapshot_Channel struct {
	Name      string                             `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	TopicNick string                             `protobuf:"bytes,2,opt,name=topic_nick,json=topicNick" json:"topic_nick,omitempty"`
	TopicTime *Timestamp                         `protobuf:"bytes,3,opt,name=topic_time,json=topicTime" json:"topic_time,omitempty"`
	Topic     string                             `protobuf:"bytes,4,opt,name=topic" json:"topic,omitempty"`
	Nicks     map[string]*Snapshot_Channel_Modes `protobuf:"bytes,5,rep,name=nicks" json:"nicks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Modes     []string                           `protobuf:"bytes,6,rep,name=modes" json:"modes,omitempty"`
	Bans      []*Snapshot_Channel_BanPattern     `protobuf:"bytes,7,rep,name=bans" json:"bans,omitempty"`
}

func (m *Snapshot_Channel) Reset()                    { *m = Snapshot_Channel{} }
func (m *Snapshot_Channel) String() string            { return proto1.CompactTextString(m) }
func (*Snapshot_Channel) ProtoMessage()               {}
func (*Snapshot_Channel) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 2} }

func (m *Snapshot_Channel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Snapshot_Channel) GetTopicNick() string {
	if m != nil {
		return m.TopicNick
	}
	return ""
}

func (m *Snapshot_Channel) GetTopicTime() *Timestamp {
	if m != nil {
		return m.TopicTime
	}
	return nil
}

func (m *Snapshot_Channel) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Snapshot_Channel) GetNicks() map[string]*Snapshot_Channel_Modes {
	if m != nil {
		return m.Nicks
	}
	return nil
}

func (m *Snapshot_Channel) GetModes() []string {
	if m != nil {
		return m.Modes
	}
	return nil
}

func (m *Snapshot_Channel) GetBans() []*Snapshot_Channel_BanPattern {
	if m != nil {
		return m.Bans
	}
	return nil
}

// Modes is a workaround because proto3 does not support
// map<string, repeated string>.
type Snapshot_Channel_Modes struct {
	Mode []string `protobuf:"bytes,1,rep,name=mode" json:"mode,omitempty"`
}

func (m *Snapshot_Channel_Modes) Reset()                    { *m = Snapshot_Channel_Modes{} }
func (m *Snapshot_Channel_Modes) String() string            { return proto1.CompactTextString(m) }
func (*Snapshot_Channel_Modes) ProtoMessage()               {}
func (*Snapshot_Channel_Modes) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 2, 0} }

func (m *Snapshot_Channel_Modes) GetMode() []string {
	if m != nil {
		return m.Mode
	}
	return nil
}

type Snapshot_Channel_BanPattern struct {
	Pattern string `protobuf:"bytes,1,opt,name=pattern" json:"pattern,omitempty"`
	Regexp  string `protobuf:"bytes,2,opt,name=regexp" json:"regexp,omitempty"`
}

func (m *Snapshot_Channel_BanPattern) Reset()         { *m = Snapshot_Channel_BanPattern{} }
func (m *Snapshot_Channel_BanPattern) String() string { return proto1.CompactTextString(m) }
func (*Snapshot_Channel_BanPattern) ProtoMessage()    {}
func (*Snapshot_Channel_BanPattern) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1, 2, 2}
}

func (m *Snapshot_Channel_BanPattern) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

func (m *Snapshot_Channel_BanPattern) GetRegexp() string {
	if m != nil {
		return m.Regexp
	}
	return ""
}

type Snapshot_SVSHold struct {
	Added    *Timestamp `protobuf:"bytes,1,opt,name=added" json:"added,omitempty"`
	Duration string     `protobuf:"bytes,2,opt,name=duration" json:"duration,omitempty"`
	Reason   string     `protobuf:"bytes,3,opt,name=reason" json:"reason,omitempty"`
}

func (m *Snapshot_SVSHold) Reset()                    { *m = Snapshot_SVSHold{} }
func (m *Snapshot_SVSHold) String() string            { return proto1.CompactTextString(m) }
func (*Snapshot_SVSHold) ProtoMessage()               {}
func (*Snapshot_SVSHold) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 3} }

func (m *Snapshot_SVSHold) GetAdded() *Timestamp {
	if m != nil {
		return m.Added
	}
	return nil
}

func (m *Snapshot_SVSHold) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *Snapshot_SVSHold) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type Snapshot_Config struct {
	Revision                uint64               `protobuf:"varint,1,opt,name=revision" json:"revision,omitempty"`
	Irc                     *Snapshot_Config_IRC `protobuf:"bytes,2,opt,name=irc" json:"irc,omitempty"`
	SessionExpiration       string               `protobuf:"bytes,3,opt,name=session_expiration,json=sessionExpiration" json:"session_expiration,omitempty"`
	PostMessageCooloff      string               `protobuf:"bytes,4,opt,name=post_message_cooloff,json=postMessageCooloff" json:"post_message_cooloff,omitempty"`
	TrustedBridges          map[string]string    `protobuf:"bytes,5,rep,name=trusted_bridges,json=trustedBridges" json:"trusted_bridges,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	CaptchaUrl              string               `protobuf:"bytes,6,opt,name=captcha_url,json=captchaUrl" json:"captcha_url,omitempty"`
	CaptchaHmacSecret       string               `protobuf:"bytes,7,opt,name=captcha_hmac_secret,json=captchaHmacSecret" json:"captcha_hmac_secret,omitempty"`
	CaptchaRequiredForLogin bool                 `protobuf:"varint,8,opt,name=captcha_required_for_login,json=captchaRequiredForLogin" json:"captcha_required_for_login,omitempty"`
	MaxSessions             uint64               `protobuf:"varint,9,opt,name=max_sessions,json=maxSessions" json:"max_sessions,omitempty"`
	MaxChannels             uint64               `protobuf:"varint,10,opt,name=max_channels,json=maxChannels" json:"max_channels,omitempty"`
	Banned                  map[string]string    `protobuf:"bytes,11,rep,name=banned" json:"banned,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Snapshot_Config) Reset()                    { *m = Snapshot_Config{} }
func (m *Snapshot_Config) String() string            { return proto1.CompactTextString(m) }
func (*Snapshot_Config) ProtoMessage()               {}
func (*Snapshot_Config) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 5} }

func (m *Snapshot_Config) GetRevision() uint64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *Snapshot_Config) GetIrc() *Snapshot_Config_IRC {
	if m != nil {
		return m.Irc
	}
	return nil
}

func (m *Snapshot_Config) GetSessionExpiration() string {
	if m != nil {
		return m.SessionExpiration
	}
	return ""
}

func (m *Snapshot_Config) GetPostMessageCooloff() string {
	if m != nil {
		return m.PostMessageCooloff
	}
	return ""
}

func (m *Snapshot_Config) GetTrustedBridges() map[string]string {
	if m != nil {
		return m.TrustedBridges
	}
	return nil
}

func (m *Snapshot_Config) GetCaptchaUrl() string {
	if m != nil {
		return m.CaptchaUrl
	}
	return ""
}

func (m *Snapshot_Config) GetCaptchaHmacSecret() string {
	if m != nil {
		return m.CaptchaHmacSecret
	}
	return ""
}

func (m *Snapshot_Config) GetCaptchaRequiredForLogin() bool {
	if m != nil {
		return m.CaptchaRequiredForLogin
	}
	return false
}

func (m *Snapshot_Config) GetMaxSessions() uint64 {
	if m != nil {
		return m.MaxSessions
	}
	return 0
}

func (m *Snapshot_Config) GetMaxChannels() uint64 {
	if m != nil {
		return m.MaxChannels
	}
	return 0
}

func (m *Snapshot_Config) GetBanned() map[string]string {
	if m != nil {
		return m.Banned
	}
	return nil
}

type Snapshot_Config_IRC struct {
	Operators []*Snapshot_Config_IRC_Operator `protobuf:"bytes,1,rep,name=operators" json:"operators,omitempty"`
	Services  []*Snapshot_Config_IRC_Service  `protobuf:"bytes,2,rep,name=services" json:"services,omitempty"`
}

func (m *Snapshot_Config_IRC) Reset()                    { *m = Snapshot_Config_IRC{} }
func (m *Snapshot_Config_IRC) String() string            { return proto1.CompactTextString(m) }
func (*Snapshot_Config_IRC) ProtoMessage()               {}
func (*Snapshot_Config_IRC) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 5, 0} }

func (m *Snapshot_Config_IRC) GetOperators() []*Snapshot_Config_IRC_Operator {
	if m != nil {
		return m.Operators
	}
	return nil
}

func (m *Snapshot_Config_IRC) GetServices() []*Snapshot_Config_IRC_Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type Snapshot_Config_IRC_Operator struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *Snapshot_Config_IRC_Operator) Reset()         { *m = Snapshot_Config_IRC_Operator{} }
func (m *Snapshot_Config_IRC_Operator) String() string { return proto1.CompactTextString(m) }
func (*Snapshot_Config_IRC_Operator) ProtoMessage()    {}
func (*Snapshot_Config_IRC_Operator) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1, 5, 0, 0}
}

func (m *Snapshot_Config_IRC_Operator) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Snapshot_Config_IRC_Operator) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Snapshot_Config_IRC_Service struct {
	Password string `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
}

func (m *Snapshot_Config_IRC_Service) Reset()         { *m = Snapshot_Config_IRC_Service{} }
func (m *Snapshot_Config_IRC_Service) String() string { return proto1.CompactTextString(m) }
func (*Snapshot_Config_IRC_Service) ProtoMessage()    {}
func (*Snapshot_Config_IRC_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1, 5, 0, 1}
}

func (m *Snapshot_Config_IRC_Service) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto1.RegisterType((*Timestamp)(nil), "proto.Timestamp")
	proto1.RegisterType((*Snapshot)(nil), "proto.Snapshot")
	proto1.RegisterType((*Snapshot_IRCPrefix)(nil), "proto.Snapshot.IRCPrefix")
	proto1.RegisterType((*Snapshot_Session)(nil), "proto.Snapshot.Session")
	proto1.RegisterType((*Snapshot_Channel)(nil), "proto.Snapshot.Channel")
	proto1.RegisterType((*Snapshot_Channel_Modes)(nil), "proto.Snapshot.Channel.Modes")
	proto1.RegisterType((*Snapshot_Channel_BanPattern)(nil), "proto.Snapshot.Channel.BanPattern")
	proto1.RegisterType((*Snapshot_SVSHold)(nil), "proto.Snapshot.SVSHold")
	proto1.RegisterType((*Snapshot_Config)(nil), "proto.Snapshot.Config")
	proto1.RegisterType((*Snapshot_Config_IRC)(nil), "proto.Snapshot.Config.IRC")
	proto1.RegisterType((*Snapshot_Config_IRC_Operator)(nil), "proto.Snapshot.Config.IRC.Operator")
	proto1.RegisterType((*Snapshot_Config_IRC_Service)(nil), "proto.Snapshot.Config.IRC.Service")
	proto1.RegisterEnum("proto.Bool", Bool_name, Bool_value)
}

func init() { proto1.RegisterFile("snapshot.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xdd, 0x6e, 0x1b, 0xb7,
	0x12, 0x3e, 0xb2, 0x7e, 0x77, 0x94, 0x38, 0x36, 0xed, 0xd8, 0xcc, 0x06, 0xc1, 0xf1, 0xf1, 0x41,
	0x03, 0x21, 0x68, 0x94, 0x22, 0x6e, 0x83, 0x24, 0x05, 0x82, 0xda, 0x82, 0xd3, 0xa8, 0x48, 0x5c,
	0x63, 0xe5, 0xb4, 0x40, 0x6f, 0x16, 0xf4, 0x92, 0x96, 0x88, 0xac, 0xc8, 0x2d, 0x49, 0x29, 0x72,
	0xaf, 0x7b, 0xd9, 0xf7, 0xe8, 0x73, 0xf4, 0x79, 0xfa, 0x12, 0x05, 0x7f, 0x76, 0x6d, 0xc7, 0x72,
	0x80, 0x5e, 0x2d, 0x67, 0xbe, 0x8f, 0x33, 0xdc, 0xe1, 0xfc, 0x10, 0x56, 0xb5, 0x20, 0x85, 0x9e,
	0x48, 0xd3, 0x2f, 0x94, 0x34, 0x12, 0x35, 0xdd, 0x27, 0xee, 0x9a, 0xf3, 0x82, 0x69, 0xaf, 0xdb,
	0xdd, 0x87, 0xe8, 0x84, 0x4f, 0x99, 0x36, 0x64, 0x5a, 0xa0, 0xfb, 0x10, 0xcd, 0x04, 0x5f, 0xa4,
	0x82, 0x08, 0x89, 0x6b, 0x3b, 0xb5, 0x5e, 0x3d, 0xe9, 0x58, 0xc5, 0x11, 0x11, 0x12, 0x6d, 0x43,
	0x9b, 0xeb, 0xf4, 0x37, 0xa6, 0x24, 0x5e, 0xd9, 0xa9, 0xf5, 0x3a, 0x49, 0x8b, 0xeb, 0x5f, 0x98,
	0x92, 0xbb, 0x7f, 0xde, 0x85, 0xce, 0x28, 0x78, 0x42, 0x7b, 0xd0, 0xd1, 0x4c, 0x6b, 0x2e, 0x85,
	0xc6, 0xb5, 0x9d, 0x7a, 0xaf, 0xfb, 0x74, 0xdb, 0x7b, 0xea, 0x97, 0x94, 0xfe, 0xc8, 0xe3, 0x49,
	0x45, 0xb4, 0x9b, 0xb2, 0x09, 0x11, 0x82, 0xe5, 0x1a, 0xaf, 0x2c, 0xdf, 0x34, 0xf0, 0x78, 0x52,
	0x11, 0xd1, 0x0b, 0xe8, 0xe8, 0xb9, 0x9e, 0xc8, 0x9c, 0x6a, 0x5c, 0x77, 0x9b, 0x1e, 0x5c, 0xf3,
	0x14, 0xf0, 0x43, 0x61, 0xd4, 0x79, 0x52, 0xd1, 0xd1, 0x33, 0x58, 0xcd, 0x89, 0x36, 0x69, 0xa1,
	0x64, 0xc6, 0xb4, 0x66, 0x14, 0x37, 0x76, 0x6a, 0xbd, 0xee, 0xd3, 0x3b, 0xc1, 0x40, 0x22, 0x4f,
	0x67, 0xda, 0x0c, 0x69, 0x72, 0xdb, 0xd2, 0x8e, 0x4b, 0x16, 0xea, 0x43, 0x2b, 0x93, 0xe2, 0x8c,
	0x8f, 0x71, 0xd3, 0xf1, 0xb7, 0xae, 0x9d, 0xd2, 0xa1, 0x49, 0x60, 0xa1, 0x3e, 0x6c, 0x38, 0x3f,
	0x5c, 0x64, 0xf9, 0x8c, 0x32, 0x9a, 0x72, 0x41, 0xd9, 0x02, 0xb7, 0x76, 0x6a, 0xbd, 0x46, 0xb2,
	0x6e, 0xa1, 0x61, 0x40, 0x86, 0x16, 0x88, 0xbf, 0x87, 0x68, 0x98, 0x0c, 0x8e, 0x15, 0x3b, 0xe3,
	0x0b, 0x84, 0xa0, 0x21, 0xc8, 0x94, 0xb9, 0x7b, 0x88, 0x12, 0xb7, 0xb6, 0xba, 0x99, 0x66, 0xca,
	0x5d, 0x40, 0x94, 0xb8, 0xb5, 0xd5, 0x4d, 0xa4, 0x36, 0xb8, 0xee, 0x75, 0x76, 0x1d, 0xff, 0xd1,
	0x82, 0x76, 0x08, 0x33, 0xfa, 0x2f, 0xac, 0x70, 0xea, 0xac, 0x2c, 0xf9, 0xc1, 0x15, 0x4e, 0xad,
	0x01, 0x32, 0x33, 0x93, 0xd2, 0xa8, 0x5d, 0x3b, 0xe7, 0x3c, 0xfb, 0x50, 0x1a, 0xb5, 0x6b, 0x14,
	0x43, 0xc7, 0x3a, 0x74, 0x87, 0x6a, 0x38, 0x7d, 0x25, 0x5b, 0x4c, 0x31, 0x92, 0x3b, 0xac, 0xe9,
	0xb1, 0x52, 0xb6, 0x58, 0x75, 0xbb, 0xad, 0x9d, 0xba, 0xc5, 0xaa, 0x4b, 0xfc, 0x06, 0x5c, 0x88,
	0x53, 0x92, 0x19, 0x3e, 0xe7, 0xe6, 0x1c, 0xb7, 0xdd, 0x39, 0xd7, 0xc2, 0x39, 0xab, 0xd4, 0x4c,
	0x6e, 0x59, 0xda, 0x7e, 0x60, 0x59, 0x93, 0xb2, 0x60, 0x8a, 0x18, 0xa9, 0x70, 0xc7, 0x25, 0x63,
	0x25, 0xa3, 0x7b, 0xd0, 0x21, 0x1f, 0xc9, 0x79, 0x3a, 0xd5, 0x63, 0x1c, 0xb9, 0xa3, 0xb4, 0xad,
	0xfc, 0x4e, 0x8f, 0x11, 0x86, 0x76, 0xa6, 0x18, 0x31, 0x8c, 0xe2, 0x2d, 0x97, 0xdd, 0xa5, 0x88,
	0x9e, 0xc0, 0x86, 0x99, 0x28, 0x69, 0x4c, 0xce, 0xc5, 0x38, 0x65, 0x8b, 0x42, 0x0a, 0x26, 0x0c,
	0x06, 0xc7, 0x42, 0x17, 0xd0, 0x61, 0x40, 0xd0, 0x03, 0x00, 0x2e, 0xe6, 0xdc, 0x30, 0x9a, 0x1a,
	0x89, 0xbb, 0xee, 0xb7, 0xa2, 0xa0, 0x39, 0x91, 0x68, 0x13, 0x9a, 0x53, 0x49, 0x99, 0xc6, 0xb7,
	0x1c, 0xe2, 0x05, 0x1b, 0x55, 0x3d, 0xe7, 0x14, 0xdf, 0xf6, 0x51, 0xb5, 0x6b, 0xab, 0x2b, 0x88,
	0xd6, 0x78, 0xd5, 0xeb, 0xec, 0x1a, 0x6d, 0x41, 0x4b, 0x33, 0x35, 0x67, 0x0a, 0xdf, 0xf1, 0x95,
	0xe6, 0x25, 0xb4, 0x07, 0x5b, 0x2e, 0x5a, 0x59, 0xce, 0x99, 0x30, 0xe9, 0x94, 0x69, 0x4d, 0xc6,
	0x2c, 0xe5, 0x14, 0xaf, 0xbb, 0x94, 0x72, 0xd9, 0x36, 0x70, 0xe0, 0x3b, 0x8f, 0x0d, 0x29, 0x7a,
	0x0e, 0xc0, 0x55, 0x96, 0x16, 0x2e, 0xab, 0x30, 0x72, 0xf1, 0xbd, 0xf7, 0x69, 0xe2, 0x56, 0x69,
	0x97, 0x44, 0x5c, 0x65, 0x21, 0x03, 0xbf, 0x0e, 0x97, 0x23, 0xa4, 0x48, 0x0b, 0x2e, 0xc6, 0x78,
	0xe3, 0x86, 0xcb, 0xe9, 0x5a, 0xda, 0x91, 0x14, 0xc7, 0x5c, 0x8c, 0xd1, 0x77, 0x21, 0xe9, 0xb5,
	0xcc, 0xe7, 0x8c, 0xa6, 0x19, 0x29, 0x4c, 0x36, 0x21, 0x78, 0xf3, 0x86, 0xbd, 0xae, 0x0c, 0x46,
	0x8e, 0x3b, 0xf0, 0x54, 0xd4, 0x83, 0x28, 0x97, 0xe3, 0xb1, 0xab, 0x17, 0x7c, 0x77, 0xa7, 0xd6,
	0x5b, 0x7d, 0xda, 0x0d, 0xfb, 0x0e, 0xa4, 0xcc, 0x93, 0x8e, 0x47, 0x87, 0x36, 0xb7, 0xbb, 0x8a,
	0x4d, 0xa5, 0x61, 0x29, 0xa1, 0x54, 0xe1, 0x6d, 0x17, 0x43, 0xf0, 0xaa, 0x7d, 0x4a, 0xd5, 0x0f,
	0x8d, 0xce, 0xda, 0xda, 0x7a, 0xfc, 0x57, 0x1d, 0xda, 0xa1, 0x81, 0x2c, 0x2d, 0xab, 0x07, 0x00,
	0x46, 0x16, 0x3c, 0x4b, 0x5d, 0xce, 0xfb, 0x3a, 0x88, 0x9c, 0xe6, 0xc8, 0x26, 0xfe, 0x93, 0x12,
	0x36, 0x7c, 0xca, 0x5c, 0x49, 0x2c, 0xfb, 0x11, 0xbf, 0xc1, 0xca, 0xf6, 0xf6, 0x9d, 0x10, 0xca,
	0xc4, 0x0b, 0xe8, 0x39, 0x34, 0xad, 0x7d, 0x8d, 0x9b, 0xae, 0x5b, 0xed, 0xde, 0xd0, 0xe2, 0xfa,
	0xd6, 0x67, 0x68, 0x59, 0x7e, 0xc3, 0x45, 0x36, 0xb5, 0x2e, 0x67, 0xd3, 0x33, 0x68, 0x9c, 0x12,
	0xa1, 0x71, 0xfb, 0xf3, 0xe6, 0x0e, 0x88, 0x38, 0x26, 0xc6, 0x30, 0x25, 0x12, 0xc7, 0x8f, 0xef,
	0x43, 0xf3, 0x5d, 0x99, 0x8e, 0xd6, 0x92, 0xeb, 0xd3, 0x51, 0xe2, 0xd6, 0xf1, 0xcf, 0x00, 0x17,
	0xfe, 0xd1, 0x1a, 0xd4, 0x3f, 0xb0, 0xf3, 0x10, 0x2b, 0xbb, 0x44, 0x7b, 0xd0, 0x9c, 0x93, 0x7c,
	0xc6, 0x5c, 0x94, 0x96, 0xb4, 0xdc, 0xd2, 0xab, 0xf3, 0x90, 0x78, 0xee, 0xcb, 0x95, 0xe7, 0xb5,
	0xf8, 0x15, 0xc0, 0xc5, 0x49, 0x6c, 0x25, 0x16, 0x7e, 0x19, 0x8c, 0x97, 0xa2, 0xcd, 0x7d, 0xc5,
	0xc6, 0x6c, 0x51, 0x84, 0x7b, 0x08, 0x52, 0xcc, 0xa0, 0x3d, 0xfa, 0x69, 0xf4, 0x46, 0xe6, 0x14,
	0x3d, 0x84, 0x26, 0xa1, 0x94, 0x95, 0x4d, 0xed, 0xfa, 0x55, 0x78, 0xd8, 0x76, 0x09, 0x3a, 0x53,
	0xc4, 0x70, 0x29, 0x82, 0xb1, 0x4a, 0xf6, 0x6e, 0x88, 0x96, 0x22, 0xb4, 0xb8, 0x20, 0xc5, 0x27,
	0x70, 0xfb, 0xca, 0xd4, 0x58, 0x12, 0x82, 0xc7, 0x57, 0x43, 0x70, 0x7d, 0xbe, 0xf9, 0x63, 0x5e,
	0xfe, 0xf9, 0xdf, 0xdb, 0xd0, 0xf2, 0xb3, 0xc1, 0x77, 0xca, 0x39, 0xb7, 0xad, 0xd9, 0x19, 0x6d,
	0x24, 0x95, 0x8c, 0xbe, 0x84, 0x3a, 0x57, 0x59, 0xb0, 0x1b, 0x2f, 0x1f, 0x2e, 0xb6, 0x54, 0x13,
	0x4b, 0x43, 0x8f, 0x01, 0x85, 0x09, 0x6a, 0x1b, 0x16, 0x0f, 0x3f, 0xea, 0x7f, 0x67, 0x3d, 0x20,
	0x87, 0x15, 0x80, 0xbe, 0x82, 0xcd, 0x42, 0xea, 0x8b, 0xae, 0x91, 0x49, 0x99, 0xcb, 0xb3, 0xb3,
	0x90, 0xa3, 0xc8, 0x62, 0xa1, 0x69, 0x0c, 0x3c, 0x82, 0x46, 0x70, 0xc7, 0xa8, 0x99, 0xb6, 0x3d,
	0xee, 0x54, 0x71, 0x3a, 0x66, 0x65, 0xea, 0x3e, 0xba, 0xe1, 0x68, 0x27, 0x9e, 0x7d, 0xe0, 0xc9,
	0x3e, 0x85, 0x57, 0xcd, 0x15, 0xa5, 0x2d, 0xd9, 0xd0, 0x12, 0xd2, 0x99, 0xca, 0xdd, 0x2c, 0x8c,
	0x12, 0x08, 0xaa, 0xf7, 0x2a, 0xb7, 0x43, 0xb3, 0x24, 0x4c, 0xa6, 0x24, 0x4b, 0x35, 0xcb, 0x14,
	0x33, 0x6e, 0x30, 0x44, 0xc9, 0x7a, 0x80, 0xde, 0x4c, 0x49, 0x36, 0x72, 0x00, 0xfa, 0x16, 0xe2,
	0x92, 0xaf, 0xd8, 0xaf, 0x33, 0xae, 0x18, 0x4d, 0xcf, 0xa4, 0x4a, 0x73, 0x39, 0xe6, 0x22, 0x4c,
	0x87, 0xed, 0xc0, 0x48, 0x02, 0xe1, 0xb5, 0x54, 0x6f, 0x2d, 0x8c, 0xfe, 0x07, 0xb7, 0xa6, 0x64,
	0x91, 0x56, 0x4f, 0x96, 0xc8, 0xdd, 0x48, 0x77, 0x4a, 0x16, 0xa3, 0xf2, 0x71, 0x12, 0x28, 0xd5,
	0x08, 0x83, 0x8a, 0x32, 0x28, 0xa7, 0xd8, 0x4b, 0x68, 0x9d, 0xda, 0x25, 0x75, 0x83, 0x60, 0x59,
	0x2d, 0xfa, 0xf8, 0x1c, 0x38, 0x92, 0x8f, 0x4b, 0xd8, 0x11, 0xff, 0x5d, 0x83, 0xfa, 0x30, 0x19,
	0xa0, 0x7d, 0x88, 0xca, 0x11, 0x56, 0xbe, 0x9c, 0xfe, 0x7f, 0x73, 0x06, 0xf4, 0x7f, 0x0c, 0xdc,
	0xe4, 0x62, 0x17, 0x7a, 0x65, 0xdf, 0x5e, 0x6a, 0xce, 0x33, 0x56, 0x3e, 0xa3, 0x76, 0x3f, 0x63,
	0x61, 0xe4, 0xa9, 0x49, 0xb5, 0x27, 0x7e, 0x09, 0x9d, 0xd2, 0xec, 0xd2, 0x36, 0x19, 0x43, 0xc7,
	0x8e, 0xa7, 0x8f, 0x52, 0xd1, 0xb2, 0x9e, 0x4a, 0x39, 0xfe, 0xc2, 0x3e, 0x38, 0x9c, 0x9d, 0x2b,
	0xb4, 0xda, 0x27, 0xb4, 0x7d, 0xd8, 0x58, 0x92, 0x24, 0x4b, 0x8a, 0x6c, 0xf3, 0x72, 0x91, 0x45,
	0x97, 0x6b, 0xe9, 0x05, 0x74, 0x2f, 0xc5, 0xf1, 0xdf, 0x6c, 0x7d, 0xf4, 0x10, 0x1a, 0x76, 0x80,
	0xa0, 0x08, 0x9a, 0xef, 0x8f, 0x46, 0x87, 0x27, 0x6b, 0xff, 0x41, 0x1d, 0x68, 0x9c, 0x24, 0xef,
	0x0f, 0xd7, 0x6a, 0x56, 0xf9, 0x7a, 0xff, 0xed, 0xe8, 0x70, 0x6d, 0xe5, 0xb4, 0xe5, 0xa2, 0xb6,
	0xf7, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x43, 0xe8, 0x84, 0x41, 0x0b, 0x00, 0x00,
}