package ofctl

/*
#include "include/ofctl.h"
*/
import "C"
import (
	"fmt"
	"net/netip"
	"strings"
)

// FlowMatch represents a flow match condition.
type FlowMatch struct {
	Flow *Match
	WC   *FlowWildCards
}

func (fm *FlowMatch) DlType() *PacketType {
	if fm.WC.DlType() {
		return &fm.Flow.DlType
	}
	return nil
}

func (fm *FlowMatch) NwSrc() *netip.Addr {
	if fm.WC.NwSrc() {
		return &fm.Flow.NwSrc
	}
	return nil
}

func (fm *FlowMatch) NwDst() *netip.Addr {
	if fm.WC.NwDst() {
		return &fm.Flow.NwDst
	}
	return nil
}

func (fm *FlowMatch) NwProto() *ProtocolType {
	if fm.WC.NwProto() {
		return &fm.Flow.NwProto
	}
	return nil
}

func (fm *FlowMatch) TpSrc() *uint16 {
	if fm.WC.TpSrc() {
		return &fm.Flow.TpSrc
	}

	return nil
}

func (fm *FlowMatch) TpDst() *uint16 {
	if fm.WC.TpDst() {
		return &fm.Flow.TpDst
	}

	return nil
}

func (fm *FlowMatch) String() string {
	//flow := fm.Match
	//wc := fm.WC

	var elems []string

	//if wc.Masks.PktMark != 0 {
	//	elems = append(elems, fmt.Sprintf("pkt_mark=%d/%d", flow.PktMark, wc.Masks.PktMark))
	//}
	//
	//if wc.Masks.RecircID != 0 {
	//	elems = append(elems, fmt.Sprintf("recirc_id=%d/%d", flow.RecircID, wc.Masks.RecircID))
	//}

	// if (wc->masks.dp_hash) {
	//        format_uint32_masked(s, "dp_hash", f->dp_hash,
	//                             wc->masks.dp_hash);
	//    }
	//
	//    if (wc->masks.conj_id) {
	//        ds_put_format(s, "%sconj_id%s=%"PRIu32",",
	//                      colors.param, colors.end, f->conj_id);
	//    }
	//
	//    if (wc->masks.skb_priority) {
	//        ds_put_format(s, "%sskb_priority=%s%#"PRIx32",",
	//                      colors.param, colors.end, f->skb_priority);
	//    }
	//
	//    if (wc->masks.actset_output) {
	//        ds_put_format(s, "%sactset_output=%s", colors.param, colors.end);
	//        ofputil_format_port(f->actset_output, port_map, s);
	//        ds_put_char(s, ',');
	//    }

	//    if (wc->masks.ct_state) {
	//        if (wc->masks.ct_state == UINT8_MAX) {
	//            ds_put_format(s, "%sct_state=%s", colors.param, colors.end);
	//            if (f->ct_state) {
	//                format_flags(s, ct_state_to_string, f->ct_state, '|');
	//            } else {
	//                ds_put_cstr(s, "0"); /* No state. */
	//            }
	//        } else {
	//            format_flags_masked(s, "ct_state", ct_state_to_string,
	//                                f->ct_state, wc->masks.ct_state, UINT8_MAX);
	//        }
	//        ds_put_char(s, ',');
	//    }

	//if wc.CtState() {
	//	var flags []string
	//	if wc.Masks.CtState == math.MaxUint8 {
	//		if flow.CtState != 0 {
	//			flags = []string{}
	//		} else {
	//			flags = append(flags, fmt.Sprintf("%d", flow.CtState))
	//		}
	//	} else {
	//		flags = []string{}
	//	}
	//	elems = append(elems, fmt.Sprintf("ct_state=%s", strings.Join(flags, ",")))
	//}

	// if (wc->masks.ct_zone) {
	//        format_uint16_masked(s, "ct_zone", f->ct_zone, wc->masks.ct_zone);
	//    }
	//
	//    if (wc->masks.ct_mark) {
	//        format_uint32_masked(s, "ct_mark", f->ct_mark, wc->masks.ct_mark);
	//    }
	//
	//    if (!ovs_u128_is_zero(wc->masks.ct_label)) {
	//        format_ct_label_masked(s, &f->ct_label, &wc->masks.ct_label);
	//    }
	//
	//    format_ip_netmask(s, "ct_nw_src", f->ct_nw_src,
	//                      wc->masks.ct_nw_src);
	//    format_ipv6_netmask(s, "ct_ipv6_src", &f->ct_ipv6_src,
	//                        &wc->masks.ct_ipv6_src);
	//    format_ip_netmask(s, "ct_nw_dst", f->ct_nw_dst,
	//                      wc->masks.ct_nw_dst);
	//    format_ipv6_netmask(s, "ct_ipv6_dst", &f->ct_ipv6_dst,
	//                        &wc->masks.ct_ipv6_dst);
	//    if (wc->masks.ct_nw_proto) {
	//        ds_put_format(s, "%sct_nw_proto=%s%"PRIu8",",
	//                      colors.param, colors.end, f->ct_nw_proto);
	//        format_be16_masked(s, "ct_tp_src", f->ct_tp_src, wc->masks.ct_tp_src);
	//        format_be16_masked(s, "ct_tp_dst", f->ct_tp_dst, wc->masks.ct_tp_dst);
	//    }
	//
	//    if (wc->masks.packet_type &&
	//        (!match_has_default_packet_type(match) || is_megaflow)) {
	//        format_packet_type_masked(s, f->packet_type, wc->masks.packet_type);
	//        ds_put_char(s, ',');
	//        if (pt_ns(f->packet_type) == OFPHTN_ETHERTYPE) {
	//            dl_type = pt_ns_type_be(f->packet_type);
	//        }
	//    }

	//if fm.DlType() {
	//	elems = append(elems, flow.NwProto.WithPacket(flow.DlType))
	//}

	dlType := fm.DlType()
	nwProto := fm.NwProto()

	if dlType != nil {
		if nwProto != nil {
			elems = append(elems, nwProto.WithPacket(*dlType))
		}
	}

	nwSrc := fm.NwSrc()
	nwDst := fm.NwDst()

	if nwSrc != nil {
		elems = append(elems, fmt.Sprintf("nw_src=%s", *nwSrc))
	}

	if nwDst != nil {
		elems = append(elems, fmt.Sprintf("nw_dst=%s", *nwDst))
	}

	tpSrc := fm.TpSrc()
	tpDst := fm.TpDst()

	if dlType != nil {
		switch {
		case *dlType == ETH_TYPE_IP && *nwProto == IPPROTO_ICMP:
			if tpSrc != nil {
				elems = append(elems, fmt.Sprintf("icmp_type=%d", *tpSrc))
			}
			if tpDst != nil {
				elems = append(elems, fmt.Sprintf("icmp_code=%d", *tpDst))
			}
		case *dlType == ETH_TYPE_IPV6 && *nwProto == IPPROTO_ICMPV6:
			if tpSrc != nil {
				elems = append(elems, fmt.Sprintf("icmp_type=%d", *tpSrc))
			}
			if tpDst != nil {
				elems = append(elems, fmt.Sprintf("icmp_code=%d", *tpDst))
			}
			// todo: complete others
		default:
			if tpSrc != nil {
				elems = append(elems, fmt.Sprintf("tp_src=%d", *tpSrc))
			}
			if tpDst != nil {
				elems = append(elems, fmt.Sprintf("tp_dst=%d", *tpDst))
			}
		}
	}

	// ...

	// if (dl_type == htons(ETH_TYPE_IP) &&
	//        f->nw_proto == IPPROTO_ICMP) {
	//        format_be16_masked(s, "icmp_type", f->tp_src, wc->masks.tp_src);
	//        format_be16_masked(s, "icmp_code", f->tp_dst, wc->masks.tp_dst);
	//    } else if (dl_type == htons(ETH_TYPE_IPV6) &&
	//               f->nw_proto == IPPROTO_ICMPV6) {
	//        format_be16_masked(s, "icmp_type", f->tp_src, wc->masks.tp_src);
	//        format_be16_masked(s, "icmp_code", f->tp_dst, wc->masks.tp_dst);
	//        format_ipv6_netmask(s, "nd_target", &f->nd_target,
	//                            &wc->masks.nd_target);
	//        format_eth_masked(s, "nd_sll", f->arp_sha, wc->masks.arp_sha);
	//        format_eth_masked(s, "nd_tll", f->arp_tha, wc->masks.arp_tha);
	//        if (wc->masks.igmp_group_ip4) {
	//            format_be32_masked(s,"nd_reserved", f->igmp_group_ip4,
	//                               wc->masks.igmp_group_ip4);
	//        }
	//        if (wc->masks.tcp_flags) {
	//            format_be16_masked(s,"nd_options_type", f->tcp_flags,
	//                               wc->masks.tcp_flags);
	//        }
	//    } else {
	//        format_be16_masked(s, "tp_src", f->tp_src, wc->masks.tp_src);
	//        format_be16_masked(s, "tp_dst", f->tp_dst, wc->masks.tp_dst);
	//    }

	return strings.Join(elems, ",")
}

// newFlowMatch function to create Go FlowMatch from C.struct_match
func newFlowMatch(ptr *C.struct_match) *FlowMatch {
	flowMatch := &FlowMatch{
		Flow: newFlow(&ptr.flow),
		WC:   newFlowWildCards(&ptr.wc),
	}

	return flowMatch
}
