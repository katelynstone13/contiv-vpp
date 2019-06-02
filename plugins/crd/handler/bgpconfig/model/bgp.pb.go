// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bgp.proto

package model

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PeerConf_RemovePrivateAs int32

const (
	PeerConf_NONE    PeerConf_RemovePrivateAs = 0
	PeerConf_ALL     PeerConf_RemovePrivateAs = 1
	PeerConf_REPLACE PeerConf_RemovePrivateAs = 2
)

var PeerConf_RemovePrivateAs_name = map[int32]string{
	0: "NONE",
	1: "ALL",
	2: "REPLACE",
}

var PeerConf_RemovePrivateAs_value = map[string]int32{
	"NONE":    0,
	"ALL":     1,
	"REPLACE": 2,
}

func (x PeerConf_RemovePrivateAs) String() string {
	return proto.EnumName(PeerConf_RemovePrivateAs_name, int32(x))
}

func (PeerConf_RemovePrivateAs) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2e12fe45eec524a6, []int{2, 0}
}

//BGP configuration
type BgpConf struct {
	Global               *GlobalConf `protobuf:"bytes,1,opt,name=global,proto3" json:"global,omitempty"`
	Peers                []*PeerConf `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BgpConf) Reset()         { *m = BgpConf{} }
func (m *BgpConf) String() string { return proto.CompactTextString(m) }
func (*BgpConf) ProtoMessage()    {}
func (*BgpConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e12fe45eec524a6, []int{0}
}
func (m *BgpConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpConf.Unmarshal(m, b)
}
func (m *BgpConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpConf.Marshal(b, m, deterministic)
}
func (m *BgpConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpConf.Merge(m, src)
}
func (m *BgpConf) XXX_Size() int {
	return xxx_messageInfo_BgpConf.Size(m)
}
func (m *BgpConf) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpConf.DiscardUnknown(m)
}

var xxx_messageInfo_BgpConf proto.InternalMessageInfo

func (m *BgpConf) GetGlobal() *GlobalConf {
	if m != nil {
		return m.Global
	}
	return nil
}

func (m *BgpConf) GetPeers() []*PeerConf {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (*BgpConf) XXX_MessageName() string {
	return "model.BgpConf"
}

// global configuration
type GlobalConf struct {
	As                   uint32   `protobuf:"varint,1,opt,name=as,proto3" json:"as,omitempty"`
	RouterId             string   `protobuf:"bytes,2,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
	ListenPort           int32    `protobuf:"varint,3,opt,name=listen_port,json=listenPort,proto3" json:"listen_port,omitempty"`
	ListenAddresses      []string `protobuf:"bytes,4,rep,name=listen_addresses,json=listenAddresses,proto3" json:"listen_addresses,omitempty"`
	Families             []uint32 `protobuf:"varint,5,rep,packed,name=families,proto3" json:"families,omitempty"`
	UseMultiplePaths     bool     `protobuf:"varint,6,opt,name=use_multiple_paths,json=useMultiplePaths,proto3" json:"use_multiple_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GlobalConf) Reset()         { *m = GlobalConf{} }
func (m *GlobalConf) String() string { return proto.CompactTextString(m) }
func (*GlobalConf) ProtoMessage()    {}
func (*GlobalConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e12fe45eec524a6, []int{1}
}
func (m *GlobalConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalConf.Unmarshal(m, b)
}
func (m *GlobalConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalConf.Marshal(b, m, deterministic)
}
func (m *GlobalConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalConf.Merge(m, src)
}
func (m *GlobalConf) XXX_Size() int {
	return xxx_messageInfo_GlobalConf.Size(m)
}
func (m *GlobalConf) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalConf.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalConf proto.InternalMessageInfo

func (m *GlobalConf) GetAs() uint32 {
	if m != nil {
		return m.As
	}
	return 0
}

func (m *GlobalConf) GetRouterId() string {
	if m != nil {
		return m.RouterId
	}
	return ""
}

func (m *GlobalConf) GetListenPort() int32 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

func (m *GlobalConf) GetListenAddresses() []string {
	if m != nil {
		return m.ListenAddresses
	}
	return nil
}

func (m *GlobalConf) GetFamilies() []uint32 {
	if m != nil {
		return m.Families
	}
	return nil
}

func (m *GlobalConf) GetUseMultiplePaths() bool {
	if m != nil {
		return m.UseMultiplePaths
	}
	return false
}

func (*GlobalConf) XXX_MessageName() string {
	return "model.GlobalConf"
}

// neighbor configuration, one struct will be created per peer
type PeerConf struct {
	Name                 string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AuthPassword         string                   `protobuf:"bytes,2,opt,name=auth_password,json=authPassword,proto3" json:"auth_password,omitempty"`
	Description          string                   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	LocalAs              uint32                   `protobuf:"varint,4,opt,name=local_as,json=localAs,proto3" json:"local_as,omitempty"`
	NeighborAddress      string                   `protobuf:"bytes,5,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	PeerAs               uint32                   `protobuf:"varint,6,opt,name=peer_as,json=peerAs,proto3" json:"peer_as,omitempty"`
	PeerGroup            string                   `protobuf:"bytes,7,opt,name=peer_group,json=peerGroup,proto3" json:"peer_group,omitempty"`
	PeerType             uint32                   `protobuf:"varint,8,opt,name=peer_type,json=peerType,proto3" json:"peer_type,omitempty"`
	RemovePrivateAs      PeerConf_RemovePrivateAs `protobuf:"varint,9,opt,name=remove_private_as,json=removePrivateAs,proto3,enum=model.PeerConf_RemovePrivateAs" json:"remove_private_as,omitempty"`
	RouteFlapDamping     bool                     `protobuf:"varint,10,opt,name=route_flap_damping,json=routeFlapDamping,proto3" json:"route_flap_damping,omitempty"`
	SendCommunity        uint32                   `protobuf:"varint,11,opt,name=send_community,json=sendCommunity,proto3" json:"send_community,omitempty"`
	NeighborInterface    string                   `protobuf:"bytes,12,opt,name=neighbor_interface,json=neighborInterface,proto3" json:"neighbor_interface,omitempty"`
	Vrf                  string                   `protobuf:"bytes,13,opt,name=vrf,proto3" json:"vrf,omitempty"`
	AllowOwnAs           uint32                   `protobuf:"varint,14,opt,name=allow_own_as,json=allowOwnAs,proto3" json:"allow_own_as,omitempty"`
	ReplacePeerAs        bool                     `protobuf:"varint,15,opt,name=replace_peer_as,json=replacePeerAs,proto3" json:"replace_peer_as,omitempty"`
	AdminDown            bool                     `protobuf:"varint,16,opt,name=admin_down,json=adminDown,proto3" json:"admin_down,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PeerConf) Reset()         { *m = PeerConf{} }
func (m *PeerConf) String() string { return proto.CompactTextString(m) }
func (*PeerConf) ProtoMessage()    {}
func (*PeerConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e12fe45eec524a6, []int{2}
}
func (m *PeerConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerConf.Unmarshal(m, b)
}
func (m *PeerConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerConf.Marshal(b, m, deterministic)
}
func (m *PeerConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerConf.Merge(m, src)
}
func (m *PeerConf) XXX_Size() int {
	return xxx_messageInfo_PeerConf.Size(m)
}
func (m *PeerConf) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerConf.DiscardUnknown(m)
}

var xxx_messageInfo_PeerConf proto.InternalMessageInfo

func (m *PeerConf) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PeerConf) GetAuthPassword() string {
	if m != nil {
		return m.AuthPassword
	}
	return ""
}

func (m *PeerConf) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PeerConf) GetLocalAs() uint32 {
	if m != nil {
		return m.LocalAs
	}
	return 0
}

func (m *PeerConf) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *PeerConf) GetPeerAs() uint32 {
	if m != nil {
		return m.PeerAs
	}
	return 0
}

func (m *PeerConf) GetPeerGroup() string {
	if m != nil {
		return m.PeerGroup
	}
	return ""
}

func (m *PeerConf) GetPeerType() uint32 {
	if m != nil {
		return m.PeerType
	}
	return 0
}

func (m *PeerConf) GetRemovePrivateAs() PeerConf_RemovePrivateAs {
	if m != nil {
		return m.RemovePrivateAs
	}
	return PeerConf_NONE
}

func (m *PeerConf) GetRouteFlapDamping() bool {
	if m != nil {
		return m.RouteFlapDamping
	}
	return false
}

func (m *PeerConf) GetSendCommunity() uint32 {
	if m != nil {
		return m.SendCommunity
	}
	return 0
}

func (m *PeerConf) GetNeighborInterface() string {
	if m != nil {
		return m.NeighborInterface
	}
	return ""
}

func (m *PeerConf) GetVrf() string {
	if m != nil {
		return m.Vrf
	}
	return ""
}

func (m *PeerConf) GetAllowOwnAs() uint32 {
	if m != nil {
		return m.AllowOwnAs
	}
	return 0
}

func (m *PeerConf) GetReplacePeerAs() bool {
	if m != nil {
		return m.ReplacePeerAs
	}
	return false
}

func (m *PeerConf) GetAdminDown() bool {
	if m != nil {
		return m.AdminDown
	}
	return false
}

func (*PeerConf) XXX_MessageName() string {
	return "model.PeerConf"
}
func init() {
	proto.RegisterEnum("model.PeerConf_RemovePrivateAs", PeerConf_RemovePrivateAs_name, PeerConf_RemovePrivateAs_value)
	proto.RegisterType((*BgpConf)(nil), "model.BgpConf")
	proto.RegisterType((*GlobalConf)(nil), "model.GlobalConf")
	proto.RegisterType((*PeerConf)(nil), "model.PeerConf")
}

func init() { proto.RegisterFile("bgp.proto", fileDescriptor_2e12fe45eec524a6) }

var fileDescriptor_2e12fe45eec524a6 = []byte{
	// 627 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x49, 0x3f, 0x93, 0xd3, 0xf5, 0x63, 0xbe, 0xc1, 0x0c, 0xc1, 0xa2, 0xa2, 0xa1, 0x4c,
	0x62, 0x9d, 0x18, 0x4f, 0xd0, 0x7d, 0x30, 0x4d, 0x8c, 0xad, 0x8a, 0xb8, 0xe3, 0x22, 0x72, 0x1b,
	0x37, 0x8d, 0x94, 0xd8, 0x96, 0xed, 0xac, 0xda, 0x1b, 0x72, 0xc7, 0x43, 0xf0, 0x20, 0x20, 0x9f,
	0x34, 0x03, 0x76, 0xe7, 0xf3, 0xfb, 0xff, 0xcf, 0xa9, 0xff, 0xc7, 0x0d, 0x04, 0xcb, 0x4c, 0xcd,
	0x94, 0x96, 0x56, 0x92, 0x6e, 0x29, 0x53, 0x5e, 0x1c, 0x9c, 0x64, 0xb9, 0xdd, 0x54, 0xcb, 0xd9,
	0x4a, 0x96, 0xa7, 0x99, 0xcc, 0xe4, 0x29, 0xaa, 0xcb, 0x6a, 0x8d, 0x15, 0x16, 0x78, 0xaa, 0xbb,
	0xa6, 0xdf, 0xa1, 0x7f, 0x9e, 0xa9, 0x0b, 0x29, 0xd6, 0xe4, 0x18, 0x7a, 0x59, 0x21, 0x97, 0xac,
	0xa0, 0x5e, 0xe8, 0x45, 0x83, 0xb3, 0xfd, 0x19, 0x4e, 0x9c, 0x5d, 0x23, 0x74, 0x96, 0x78, 0x67,
	0x20, 0x47, 0xd0, 0x55, 0x9c, 0x6b, 0x43, 0x5b, 0x61, 0x3b, 0x1a, 0x9c, 0x8d, 0x77, 0xce, 0x05,
	0xe7, 0x1a, 0x7d, 0xb5, 0x3a, 0xfd, 0xe9, 0x01, 0xfc, 0xed, 0x26, 0x23, 0x68, 0x31, 0x83, 0xc3,
	0x87, 0x71, 0x8b, 0x19, 0xf2, 0x1a, 0x02, 0x2d, 0x2b, 0xcb, 0x75, 0x92, 0xa7, 0xb4, 0x15, 0x7a,
	0x51, 0x10, 0xfb, 0x35, 0xb8, 0x49, 0xc9, 0x21, 0x0c, 0x8a, 0xdc, 0x58, 0x2e, 0x12, 0x25, 0xb5,
	0xa5, 0xed, 0xd0, 0x8b, 0xba, 0x31, 0xd4, 0x68, 0x21, 0xb5, 0x25, 0xc7, 0x30, 0xd9, 0x19, 0x58,
	0x9a, 0x6a, 0x6e, 0x0c, 0x37, 0xb4, 0x13, 0xb6, 0xa3, 0x20, 0x1e, 0xd7, 0x7c, 0xde, 0x60, 0x72,
	0x00, 0xfe, 0x9a, 0x95, 0x79, 0x91, 0x73, 0x43, 0xbb, 0x61, 0x3b, 0x1a, 0xc6, 0x4f, 0x35, 0xf9,
	0x00, 0xa4, 0x32, 0x3c, 0x29, 0xab, 0xc2, 0xe6, 0xaa, 0xe0, 0x89, 0x62, 0x76, 0x63, 0x68, 0x2f,
	0xf4, 0x22, 0x3f, 0x9e, 0x54, 0x86, 0x7f, 0xdd, 0x09, 0x0b, 0xc7, 0xa7, 0xbf, 0x3b, 0xe0, 0x37,
	0x29, 0x09, 0x81, 0x8e, 0x60, 0x25, 0xc7, 0x44, 0x41, 0x8c, 0x67, 0xf2, 0x0e, 0x86, 0xac, 0xb2,
	0x9b, 0x44, 0x31, 0x63, 0xb6, 0x52, 0x37, 0xb9, 0xf6, 0x1c, 0x5c, 0xec, 0x18, 0x09, 0x61, 0x90,
	0x72, 0xb3, 0xd2, 0xb9, 0xb2, 0xb9, 0x14, 0x98, 0x2d, 0x88, 0xff, 0x45, 0xe4, 0x15, 0xf8, 0x85,
	0x5c, 0xb1, 0x22, 0x61, 0x2e, 0x94, 0x5b, 0x58, 0x1f, 0xeb, 0xb9, 0x71, 0xb9, 0x05, 0xcf, 0xb3,
	0xcd, 0x52, 0xea, 0x26, 0x39, 0xed, 0xe2, 0x84, 0x71, 0xc3, 0x77, 0xc9, 0xc9, 0x4b, 0xe8, 0xbb,
	0x87, 0x70, 0x43, 0x7a, 0x38, 0xa4, 0xe7, 0xca, 0xb9, 0x21, 0x6f, 0x00, 0x50, 0xc8, 0xb4, 0xac,
	0x14, 0xed, 0x63, 0x77, 0xe0, 0xc8, 0xb5, 0x03, 0xee, 0x61, 0x50, 0xb6, 0x8f, 0x8a, 0x53, 0x1f,
	0x3b, 0x7d, 0x07, 0xbe, 0x3d, 0x2a, 0x4e, 0xbe, 0xc0, 0xbe, 0xe6, 0xa5, 0x7c, 0xe0, 0x89, 0xd2,
	0xf9, 0x03, 0xb3, 0xdc, 0x8d, 0x0f, 0x42, 0x2f, 0x1a, 0x9d, 0x1d, 0x3e, 0xfb, 0x1f, 0xcc, 0x62,
	0x34, 0x2e, 0x6a, 0xdf, 0xdc, 0xc4, 0x63, 0xfd, 0x3f, 0x70, 0xdb, 0xc7, 0x17, 0x4f, 0xd6, 0x05,
	0x53, 0x49, 0xca, 0x4a, 0x95, 0x8b, 0x8c, 0x42, 0xbd, 0x7d, 0x54, 0x3e, 0x17, 0x4c, 0x5d, 0xd6,
	0x9c, 0x1c, 0xc1, 0xc8, 0x70, 0x91, 0x26, 0x2b, 0x59, 0x96, 0x95, 0xc8, 0xed, 0x23, 0x1d, 0xe0,
	0xe5, 0x86, 0x8e, 0x5e, 0x34, 0x90, 0x9c, 0x00, 0x79, 0xda, 0x50, 0x2e, 0x2c, 0xd7, 0x6b, 0xb6,
	0xe2, 0x74, 0x0f, 0x53, 0xee, 0x37, 0xca, 0x4d, 0x23, 0x90, 0x09, 0xb4, 0x1f, 0xf4, 0x9a, 0x0e,
	0x51, 0x77, 0x47, 0x12, 0xc2, 0x1e, 0x2b, 0x0a, 0xb9, 0x4d, 0xe4, 0x56, 0xb8, 0x74, 0x23, 0xfc,
	0x15, 0x40, 0x76, 0xbf, 0x15, 0x73, 0x43, 0xde, 0xc3, 0x58, 0x73, 0x55, 0xb0, 0x15, 0x4f, 0x9a,
	0x0d, 0x8f, 0xf1, 0xd2, 0xc3, 0x1d, 0x5e, 0x3c, 0x2d, 0x9a, 0xa5, 0x65, 0x2e, 0x92, 0x54, 0x6e,
	0x05, 0x9d, 0xa0, 0x25, 0x40, 0x72, 0x29, 0xb7, 0x62, 0xfa, 0x11, 0xc6, 0xcf, 0x56, 0x44, 0x7c,
	0xe8, 0xdc, 0xdd, 0xdf, 0x5d, 0x4d, 0x5e, 0x90, 0x3e, 0xb4, 0xe7, 0xb7, 0xb7, 0x13, 0x8f, 0x0c,
	0xa0, 0x1f, 0x5f, 0x2d, 0x6e, 0xe7, 0x17, 0x57, 0x93, 0xd6, 0x79, 0xe7, 0xc7, 0xaf, 0xb7, 0xde,
	0xb2, 0x87, 0x5f, 0xef, 0xa7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xc5, 0x18, 0x96, 0x00,
	0x04, 0x00, 0x00,
}