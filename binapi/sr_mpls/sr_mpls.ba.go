// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/sr_mpls.api.json

/*
 Package sr_mpls is a generated from VPP binary API module 'sr_mpls'.

 It contains following objects:
	 10 messages
	  5 services

*/
package sr_mpls

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// Services represents VPP binary API services:
//
//	"services": {
//	    "sr_mpls_steering_add_del": {
//	        "reply": "sr_mpls_steering_add_del_reply"
//	    },
//	    "sr_mpls_policy_assign_endpoint_color": {
//	        "reply": "sr_mpls_policy_assign_endpoint_color_reply"
//	    },
//	    "sr_mpls_policy_mod": {
//	        "reply": "sr_mpls_policy_mod_reply"
//	    },
//	    "sr_mpls_policy_add": {
//	        "reply": "sr_mpls_policy_add_reply"
//	    },
//	    "sr_mpls_policy_del": {
//	        "reply": "sr_mpls_policy_del_reply"
//	    }
//	},
//
type Services interface {
	SrMplsPolicyAdd(*SrMplsPolicyAdd) (*SrMplsPolicyAddReply, error)
	SrMplsPolicyAssignEndpointColor(*SrMplsPolicyAssignEndpointColor) (*SrMplsPolicyAssignEndpointColorReply, error)
	SrMplsPolicyDel(*SrMplsPolicyDel) (*SrMplsPolicyDelReply, error)
	SrMplsPolicyMod(*SrMplsPolicyMod) (*SrMplsPolicyModReply, error)
	SrMplsSteeringAddDel(*SrMplsSteeringAddDel) (*SrMplsSteeringAddDelReply, error)
}

/* Messages */

// SrMplsPolicyAdd represents VPP binary API message 'sr_mpls_policy_add':
//
//	"sr_mpls_policy_add",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u32",
//	    "bsid"
//	],
//	[
//	    "u32",
//	    "weight"
//	],
//	[
//	    "u8",
//	    "type"
//	],
//	[
//	    "u8",
//	    "n_segments"
//	],
//	[
//	    "u32",
//	    "segments",
//	    0,
//	    "n_segments"
//	],
//	{
//	    "crc": "0x6f5b21cc"
//	}
//
type SrMplsPolicyAdd struct {
	Bsid      uint32
	Weight    uint32
	Type      uint8
	NSegments uint8 `struc:"sizeof=Segments"`
	Segments  []uint32
}

func (*SrMplsPolicyAdd) GetMessageName() string {
	return "sr_mpls_policy_add"
}
func (*SrMplsPolicyAdd) GetCrcString() string {
	return "6f5b21cc"
}
func (*SrMplsPolicyAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SrMplsPolicyAddReply represents VPP binary API message 'sr_mpls_policy_add_reply':
//
//	"sr_mpls_policy_add_reply",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	{
//	    "crc": "0xe8d4e804"
//	}
//
type SrMplsPolicyAddReply struct {
	Retval int32
}

func (*SrMplsPolicyAddReply) GetMessageName() string {
	return "sr_mpls_policy_add_reply"
}
func (*SrMplsPolicyAddReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SrMplsPolicyAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SrMplsPolicyMod represents VPP binary API message 'sr_mpls_policy_mod':
//
//	"sr_mpls_policy_mod",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u32",
//	    "bsid"
//	],
//	[
//	    "u8",
//	    "operation"
//	],
//	[
//	    "u32",
//	    "sl_index"
//	],
//	[
//	    "u32",
//	    "weight"
//	],
//	[
//	    "u8",
//	    "n_segments"
//	],
//	[
//	    "u32",
//	    "segments",
//	    0,
//	    "n_segments"
//	],
//	{
//	    "crc": "0x09d338ac"
//	}
//
type SrMplsPolicyMod struct {
	Bsid      uint32
	Operation uint8
	SlIndex   uint32
	Weight    uint32
	NSegments uint8 `struc:"sizeof=Segments"`
	Segments  []uint32
}

func (*SrMplsPolicyMod) GetMessageName() string {
	return "sr_mpls_policy_mod"
}
func (*SrMplsPolicyMod) GetCrcString() string {
	return "09d338ac"
}
func (*SrMplsPolicyMod) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SrMplsPolicyModReply represents VPP binary API message 'sr_mpls_policy_mod_reply':
//
//	"sr_mpls_policy_mod_reply",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	{
//	    "crc": "0xe8d4e804"
//	}
//
type SrMplsPolicyModReply struct {
	Retval int32
}

func (*SrMplsPolicyModReply) GetMessageName() string {
	return "sr_mpls_policy_mod_reply"
}
func (*SrMplsPolicyModReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SrMplsPolicyModReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SrMplsPolicyDel represents VPP binary API message 'sr_mpls_policy_del':
//
//	"sr_mpls_policy_del",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u32",
//	    "bsid"
//	],
//	{
//	    "crc": "0xe29d34fa"
//	}
//
type SrMplsPolicyDel struct {
	Bsid uint32
}

func (*SrMplsPolicyDel) GetMessageName() string {
	return "sr_mpls_policy_del"
}
func (*SrMplsPolicyDel) GetCrcString() string {
	return "e29d34fa"
}
func (*SrMplsPolicyDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SrMplsPolicyDelReply represents VPP binary API message 'sr_mpls_policy_del_reply':
//
//	"sr_mpls_policy_del_reply",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	{
//	    "crc": "0xe8d4e804"
//	}
//
type SrMplsPolicyDelReply struct {
	Retval int32
}

func (*SrMplsPolicyDelReply) GetMessageName() string {
	return "sr_mpls_policy_del_reply"
}
func (*SrMplsPolicyDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SrMplsPolicyDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SrMplsSteeringAddDel represents VPP binary API message 'sr_mpls_steering_add_del':
//
//	"sr_mpls_steering_add_del",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u8",
//	    "is_del"
//	],
//	[
//	    "u32",
//	    "bsid"
//	],
//	[
//	    "u32",
//	    "table_id"
//	],
//	[
//	    "u8",
//	    "prefix_addr",
//	    16
//	],
//	[
//	    "u32",
//	    "mask_width"
//	],
//	[
//	    "u8",
//	    "traffic_type"
//	],
//	[
//	    "u8",
//	    "next_hop",
//	    16
//	],
//	[
//	    "u8",
//	    "nh_type"
//	],
//	[
//	    "u32",
//	    "color"
//	],
//	[
//	    "u8",
//	    "co_bits"
//	],
//	[
//	    "u32",
//	    "vpn_label"
//	],
//	{
//	    "crc": "0x1591f94a"
//	}
//
type SrMplsSteeringAddDel struct {
	IsDel       uint8
	Bsid        uint32
	TableID     uint32
	PrefixAddr  []byte `struc:"[16]byte"`
	MaskWidth   uint32
	TrafficType uint8
	NextHop     []byte `struc:"[16]byte"`
	NhType      uint8
	Color       uint32
	CoBits      uint8
	VPNLabel    uint32
}

func (*SrMplsSteeringAddDel) GetMessageName() string {
	return "sr_mpls_steering_add_del"
}
func (*SrMplsSteeringAddDel) GetCrcString() string {
	return "1591f94a"
}
func (*SrMplsSteeringAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SrMplsSteeringAddDelReply represents VPP binary API message 'sr_mpls_steering_add_del_reply':
//
//	"sr_mpls_steering_add_del_reply",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	{
//	    "crc": "0xe8d4e804"
//	}
//
type SrMplsSteeringAddDelReply struct {
	Retval int32
}

func (*SrMplsSteeringAddDelReply) GetMessageName() string {
	return "sr_mpls_steering_add_del_reply"
}
func (*SrMplsSteeringAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SrMplsSteeringAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SrMplsPolicyAssignEndpointColor represents VPP binary API message 'sr_mpls_policy_assign_endpoint_color':
//
//	"sr_mpls_policy_assign_endpoint_color",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u32",
//	    "bsid"
//	],
//	[
//	    "u8",
//	    "endpoint",
//	    16
//	],
//	[
//	    "u8",
//	    "endpoint_type"
//	],
//	[
//	    "u32",
//	    "color"
//	],
//	{
//	    "crc": "0x6c82a6da"
//	}
//
type SrMplsPolicyAssignEndpointColor struct {
	Bsid         uint32
	Endpoint     []byte `struc:"[16]byte"`
	EndpointType uint8
	Color        uint32
}

func (*SrMplsPolicyAssignEndpointColor) GetMessageName() string {
	return "sr_mpls_policy_assign_endpoint_color"
}
func (*SrMplsPolicyAssignEndpointColor) GetCrcString() string {
	return "6c82a6da"
}
func (*SrMplsPolicyAssignEndpointColor) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SrMplsPolicyAssignEndpointColorReply represents VPP binary API message 'sr_mpls_policy_assign_endpoint_color_reply':
//
//	"sr_mpls_policy_assign_endpoint_color_reply",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	{
//	    "crc": "0xe8d4e804"
//	}
//
type SrMplsPolicyAssignEndpointColorReply struct {
	Retval int32
}

func (*SrMplsPolicyAssignEndpointColorReply) GetMessageName() string {
	return "sr_mpls_policy_assign_endpoint_color_reply"
}
func (*SrMplsPolicyAssignEndpointColorReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SrMplsPolicyAssignEndpointColorReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*SrMplsPolicyAdd)(nil), "sr_mpls.SrMplsPolicyAdd")
	api.RegisterMessage((*SrMplsPolicyAddReply)(nil), "sr_mpls.SrMplsPolicyAddReply")
	api.RegisterMessage((*SrMplsPolicyMod)(nil), "sr_mpls.SrMplsPolicyMod")
	api.RegisterMessage((*SrMplsPolicyModReply)(nil), "sr_mpls.SrMplsPolicyModReply")
	api.RegisterMessage((*SrMplsPolicyDel)(nil), "sr_mpls.SrMplsPolicyDel")
	api.RegisterMessage((*SrMplsPolicyDelReply)(nil), "sr_mpls.SrMplsPolicyDelReply")
	api.RegisterMessage((*SrMplsSteeringAddDel)(nil), "sr_mpls.SrMplsSteeringAddDel")
	api.RegisterMessage((*SrMplsSteeringAddDelReply)(nil), "sr_mpls.SrMplsSteeringAddDelReply")
	api.RegisterMessage((*SrMplsPolicyAssignEndpointColor)(nil), "sr_mpls.SrMplsPolicyAssignEndpointColor")
	api.RegisterMessage((*SrMplsPolicyAssignEndpointColorReply)(nil), "sr_mpls.SrMplsPolicyAssignEndpointColorReply")
}