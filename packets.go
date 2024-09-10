package ofctl

/*
#include "include/ofctl.h"
*/
import "C"
import "fmt"

type PacketType uint16

func (eth *PacketType) String() string {
	switch *eth {
	default:
		return ""
	case ETH_TYPE_IP:
		return "ip"
	case ETH_TYPE_IPV6:
		return "ipv6"
	case ETH_TYPE_ARP:
		return "arp"
	case ETH_TYPE_RARP:
		return "rarp"
	case ETH_TYPE_MPLS:
		return "mpls"
	case ETH_TYPE_MPLS_MCAST:
		return "mplsm"
	}
}

const (
	ETH_TYPE_IP          PacketType = C.ETH_TYPE_IP
	ETH_TYPE_ARP         PacketType = C.ETH_TYPE_ARP
	ETH_TYPE_TEB         PacketType = C.ETH_TYPE_TEB
	ETH_TYPE_VLAN_8021Q  PacketType = C.ETH_TYPE_VLAN_8021Q
	ETH_TYPE_VLAN        PacketType = C.ETH_TYPE_VLAN
	ETH_TYPE_VLAN_8021AD PacketType = C.ETH_TYPE_VLAN_8021AD
	ETH_TYPE_IPV6        PacketType = C.ETH_TYPE_IPV6
	ETH_TYPE_LACP        PacketType = C.ETH_TYPE_LACP
	ETH_TYPE_RARP        PacketType = C.ETH_TYPE_RARP
	ETH_TYPE_MPLS        PacketType = C.ETH_TYPE_MPLS
	ETH_TYPE_MPLS_MCAST  PacketType = C.ETH_TYPE_MPLS_MCAST
	ETH_TYPE_NSH         PacketType = C.ETH_TYPE_NSH
	ETH_TYPE_ERSPAN1     PacketType = C.ETH_TYPE_ERSPAN1 /* version 1 type II */
	ETH_TYPE_ERSPAN2     PacketType = C.ETH_TYPE_ERSPAN2 /* version 2 type III */
)

/* Standard well-defined IP protocols.  */

type ProtocolType uint8

func (proto *ProtocolType) WithPacket(eth PacketType) string {
	switch eth {
	case ETH_TYPE_IPV6:
		return fmt.Sprintf("%s6", proto.String())
	default:
		return proto.String()
	}
}

func (proto *ProtocolType) String() string {
	switch *proto {
	default:
		return "ip"
	case IPPROTO_ICMP:
		return "icmp"
	case IPPROTO_TCP:
		return "tcp"
	case IPPROTO_UDP:
		return "udp"
	case IPPROTO_SCTP:
		return "sctp"
	}
}

const (
	IPPROTO_ICMP   ProtocolType = C.IPPROTO_ICMP
	IPPROTO_ICMPV6 ProtocolType = C.IPPROTO_ICMPV6
	IPPROTO_TCP    ProtocolType = C.IPPROTO_TCP
	IPPROTO_UDP    ProtocolType = C.IPPROTO_UDP
	IPPROTO_SCTP   ProtocolType = C.IPPROTO_SCTP
)
