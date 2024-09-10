package ofctl

/*
#include "include/ofctl.h"

uint8_t *
get_eth_addr_ea(struct eth_addr *addr) {
	return addr->ea;
}
*/
import "C"
import (
	"encoding/binary"
	"net/netip"
)

// Flow represents a network flow.
type Flow struct {
	Version OpenFlowVersion

	TableID  uint8
	Priority uint16
	Match    *Match
	Actions  FlowActions
}

func NewFlow() *Flow {
	flow := new(Flow)

	flow.TableID = OFP_DEFAULT_TABLE
	flow.Priority = OFP_DEFAULT_PRIORITY
	flow.Match = nil // todo
	flow.Actions = nil

	return flow
}

const (
	FLOW_N_REGS           = C.FLOW_N_REGS
	FLOW_MAX_VLAN_HEADERS = C.FLOW_MAX_VLAN_HEADERS
	FLOW_MAX_MPLS_LABELS  = 4 // C.FLOW_MAX_MPLS_LABELS
)

// Tunnel represents the encapsulating tunnel parameters.
type Tunnel struct {
	// Tunnel-related fields would go here, such as tunnel ID, IPs, etc.
}

// VLANHeader represents VLAN information.
type VLANHeader struct {
	TCI uint16
}

// Match represents a network flow.
type Match struct {
	/* Metadata */

	//    struct flow_tnl tunnel;     /* Encapsulating tunnel parameters. */
	Tunnel Tunnel

	//    ovs_be64 metadata;          /* OpenFlow Metadata. */
	Metadata uint64

	//    uint32_t regs[FLOW_N_REGS]; /* Registers. */
	Registers [FLOW_N_REGS]uint32

	//    uint32_t skb_priority;      /* Packet priority for QoS. */
	SKBPriority uint32

	//    uint32_t pkt_mark;          /* Packet mark. */
	PktMark uint32

	//    uint32_t dp_hash;           /* Datapath computed hash value. The exact
	//                                 * computation is opaque to the user space. */
	DPHash uint32

	//    union flow_in_port in_port; /* Input port.*/
	InPort uint32

	//    uint32_t recirc_id;         /* Must be exact match. */
	RecircID uint32

	//    uint8_t ct_state;           /* Connection tracking state. */
	CtState uint8

	//    uint8_t ct_nw_proto;        /* CT orig tuple IP protocol. */
	CtNWProto uint8

	//    uint16_t ct_zone;           /* Connection tracking zone. */
	CtZone uint16

	//    uint32_t ct_mark;           /* Connection mark.*/
	CtMark uint32

	//    ovs_be32 packet_type;       /* OpenFlow packet type. */
	PacketType uint32

	//    ovs_u128 ct_label;          /* Connection label. */
	CtLabel [2]uint64

	//    uint32_t conj_id;           /* Conjunction ID. */
	ConjID uint32

	//    ofp_port_t actset_output;   /* Output port in action set. */
	ActsetOutput uint32

	/* L2, Order the same as in the Ethernet header! (64-bit aligned) */

	//    struct eth_addr dl_dst;     /* Ethernet destination address. */
	DstMAC [6]byte
	//    struct eth_addr dl_src;     /* Ethernet source address. */
	SrcMAC [6]byte
	//    ovs_be16 dl_type;           /* Ethernet frame type.
	//                                   Note: This also holds the Ethertype for L3
	//                                   packets of type PACKET_TYPE(1, Ethertype) */
	DlType PacketType
	//    uint8_t pad1[2];            /* Pad to 64 bits. */

	//    union flow_vlan_hdr vlans[FLOW_MAX_VLAN_HEADERS]; /* VLANs */
	//    ovs_be32 mpls_lse[ROUND_UP(FLOW_MAX_MPLS_LABELS, 2)]; /* MPLS label stack
	//                                                             (with padding). */
	VLANs [FLOW_MAX_VLAN_HEADERS]VLANHeader
	MPLS  [FLOW_MAX_MPLS_LABELS]uint32

	/* L3 (64-bit aligned) */

	//    ovs_be32 nw_src;            /* IPv4 source address or ARP SPA. */
	//    ovs_be32 nw_dst;            /* IPv4 destination address or ARP TPA. */
	//    ovs_be32 ct_nw_src;         /* CT orig tuple IPv4 source address. */
	//    ovs_be32 ct_nw_dst;         /* CT orig tuple IPv4 destination address. */
	//    struct in6_addr ipv6_src;   /* IPv6 source address. */
	//    struct in6_addr ipv6_dst;   /* IPv6 destination address. */
	//    struct in6_addr ct_ipv6_src; /* CT orig tuple IPv6 source address. */
	//    struct in6_addr ct_ipv6_dst; /* CT orig tuple IPv6 destination address. */
	//    ovs_be32 ipv6_label;        /* IPv6 flow label. */
	//    uint8_t nw_frag;            /* FLOW_FRAG_* flags. */
	//    uint8_t nw_tos;             /* IP ToS (including DSCP and ECN). */
	//    uint8_t nw_ttl;             /* IP TTL/Hop Limit. */
	//    uint8_t nw_proto;           /* IP protocol or low 8 bits of ARP opcode. */
	NwSrc     netip.Addr
	NwDst     netip.Addr
	CtNWSrc   uint32
	CtNWDst   uint32
	IPv6Src   [16]byte
	IPv6Dst   [16]byte
	CtIPv6Src [16]byte
	CtIPv6Dst [16]byte
	IPv6Label uint32
	NWFrag    uint8
	NWTOS     uint8
	NWTTL     uint8
	NwProto   ProtocolType

	/* L4 (64-bit aligned) */

	//    struct in6_addr nd_target;  /* IPv6 neighbor discovery (ND) target. */
	//    struct eth_addr arp_sha;    /* ARP/ND source hardware address. */
	//    struct eth_addr arp_tha;    /* ARP/ND target hardware address. */
	//    ovs_be16 tcp_flags;         /* TCP flags/ICMPv6 ND options type. */
	//    ovs_be16 pad2;              /* Pad to 64 bits. */
	//    struct ovs_key_nsh nsh;     /* Network Service Header keys */

	//    ovs_be16 tp_src;            /* TCP/UDP/SCTP source port/ICMP type. */
	//    ovs_be16 tp_dst;            /* TCP/UDP/SCTP destination port/ICMP code. */
	//    ovs_be16 ct_tp_src;         /* CT original tuple source port/ICMP type. */
	//    ovs_be16 ct_tp_dst;         /* CT original tuple dst port/ICMP code. */
	//    ovs_be32 igmp_group_ip4;    /* IGMP group IPv4 address/ICMPv6 ND reserved
	//                                 * field.
	//                                 * Keep last for BUILD_ASSERT_DECL below. */
	//    ovs_be32 pad3;              /* Pad to 64 bits. */
	TpSrc     uint16
	TpDst     uint16
	CtTPSrc   uint16
	CtTPDst   uint16
	IGMPGroup uint32
}

//func newFlow(ptr *C.struct_flow) *Match {
//	return &Match{
//		PktMark:  uint32(ptr.pkt_mark),
//		RecircID: uint32(ptr.recirc_id),
//		CtState:  uint8(ptr.ct_state),
//	}
//}

func ipFromUint32(value uint32) (netip.Addr, bool) {
	a := make([]byte, 4)
	binary.LittleEndian.PutUint32(a[:4], value)

	return netip.AddrFromSlice(a)
}

// newFlow converts a C struct flow pointer to a Go Match struct.
func newFlow(cFlow *C.struct_flow) *Match {
	flow := &Match{}

	// Metadata
	//flow.Metadata = uint64(cFlow.metadata)
	//flow.SKBPriority = uint32(cFlow.skb_priority)
	//flow.PktMark = uint32(cFlow.pkt_mark)
	//flow.DPHash = uint32(cFlow.dp_hash)
	//flow.InPort = uint32(cFlow.in_port.ofp_port)
	//flow.RecircID = uint32(cFlow.recirc_id)
	//flow.CtState = uint8(cFlow.ct_state)
	//flow.CtNWProto = uint8(cFlow.ct_nw_proto)
	//flow.CtZone = uint16(cFlow.ct_zone)
	//flow.CtMark = uint32(cFlow.ct_mark)
	//flow.PacketType = uint32(cFlow.packet_type)

	// Registers
	//for i := 0; i < FLOW_N_REGS; i++ {
	//	flow.Registers[i] = uint32(cFlow.regs[i])
	//}

	// Layer 2 (Ethernet)
	//copy(flow.DstMAC[:], C.GoBytes(unsafe.Pointer(&cFlow.dl_dst), 6))
	//copy(flow.SrcMAC[:], C.GoBytes(unsafe.Pointer(&cFlow.dl_src), 6))
	flow.DlType = PacketType(cFlow.dl_type)

	// VLAN Headers
	//for i := 0; i < FLOW_MAX_VLAN_HEADERS; i++ {
	//	flow.VLANs[i].TCI = uint16(C.ntohs(cFlow.vlans[i].tci))
	//}

	// MPLS Labels
	//for i := 0; i < FLOW_MAX_MPLS_LABELS; i++ {
	//	flow.MPLS[i] = uint32(C.ntohl(cFlow.mpls_lse[i]))
	//}

	// Layer 3 (IP)
	//netip.AddrFrom4()
	//var ret [4]byte
	//byteorder.BePutUint32(ret[:], uint32(ip.addr.lo))
	//return ret[:]

	flow.NwSrc, _ = ipFromUint32(uint32(cFlow.nw_src))
	flow.NwDst, _ = ipFromUint32(uint32(cFlow.nw_dst))

	//flow.CtNWSrc = uint32(C.ntohl(cFlow.ct_nw_src))
	//flow.CtNWDst = uint32(C.ntohl(cFlow.ct_nw_dst))
	//copy(flow.IPv6Src[:], C.GoBytes(unsafe.Pointer(&cFlow.ipv6_src), 16))
	//copy(flow.IPv6Dst[:], C.GoBytes(unsafe.Pointer(&cFlow.ipv6_dst), 16))
	//copy(flow.CtIPv6Src[:], C.GoBytes(unsafe.Pointer(&cFlow.ct_ipv6_src), 16))
	//copy(flow.CtIPv6Dst[:], C.GoBytes(unsafe.Pointer(&cFlow.ct_ipv6_dst), 16))
	//flow.IPv6Label = uint32(C.ntohl(cFlow.ipv6_label))
	//flow.NWFrag = uint8(cFlow.nw_frag)
	//flow.NWTOS = uint8(cFlow.nw_tos)
	//flow.NWTTL = uint8(cFlow.nw_ttl)
	flow.NwProto = ProtocolType(cFlow.nw_proto)

	// Layer 4 (Transport)
	flow.TpSrc = uint16(cFlow.tp_src)
	flow.TpDst = uint16(cFlow.tp_dst)
	//flow.CtTPSrc = uint16(C.ntohs(cFlow.ct_tp_src))
	//flow.CtTPDst = uint16(C.ntohs(cFlow.ct_tp_dst))
	//flow.IGMPGroup = uint32(C.ntohl(cFlow.igmp_group_ip4))

	return flow
}

type FlowWildCards struct {
	Masks *Match
}

func (wc *FlowWildCards) DlType() bool {
	return wc.Masks.DlType != 0
}

func (wc *FlowWildCards) NwSrc() bool {
	return !wc.Masks.NwSrc.IsUnspecified()
}

func (wc *FlowWildCards) NwDst() bool {
	return !wc.Masks.NwDst.IsUnspecified()
}

func (wc *FlowWildCards) NwProto() bool {
	return wc.Masks.NwProto != 0
}

func (wc *FlowWildCards) TpSrc() bool {
	return wc.Masks.TpSrc != 0
}

func (wc *FlowWildCards) TpDst() bool {
	return wc.Masks.TpDst != 0
}

func newFlowWildCards(ptr *C.struct_flow_wildcards) *FlowWildCards {
	return &FlowWildCards{
		Masks: newFlow(&ptr.masks),
	}
}
