package ofctl

/*
#include "include/ofctl.h"
#include "include/connection.h"
#include "include/ofp-port.h"

static void
port_iterator_fetch_port_desc(struct port_iterator *pi)
{
   pi->variant = PI_PORT_DESC;
   pi->more = true;

   struct ofpbuf *rq = ofputil_encode_port_desc_stats_request(
       vconn_get_version(pi->vconn), OFPP_ANY);
   pi->send_xid = ((struct ofp_header *) rq->data)->xid;
   send_openflow_buffer(pi->vconn, rq);
}

static void
port_iterator_fetch_features(struct port_iterator *pi)
{
   pi->variant = PI_FEATURES;

   // Fetch the switch's ofp_switch_features.
   enum ofp_version version = vconn_get_version(pi->vconn);
   struct ofpbuf *rq = ofpraw_alloc(OFPRAW_OFPT_FEATURES_REQUEST, version, 0);
   run(vconn_transact(pi->vconn, rq, &pi->reply),
       "talking to %s", vconn_get_name(pi->vconn));

   enum ofptype type;
   if (ofptype_decode(&type, pi->reply->data)
       || type != OFPTYPE_FEATURES_REPLY) {
       ovs_fatal(0, "%s: received bad features reply",
                 vconn_get_name(pi->vconn));
   }
   if (!ofputil_switch_features_has_ports(pi->reply)) {
       // The switch features reply does not contain a complete list of ports.
       // Probably, there are more ports than will fit into a single 64 kB
       // OpenFlow message.  Use OFPST_PORT_DESC to get a complete list of
       // ports.
       ofpbuf_delete(pi->reply);
       pi->reply = NULL;
       port_iterator_fetch_port_desc(pi);
       return;
   }

   struct ofputil_switch_features features;
   enum ofperr error = ofputil_pull_switch_features(pi->reply, &features);
   if (error) {
       ovs_fatal(0, "%s: failed to decode features reply (%s)",
                 vconn_get_name(pi->vconn), ofperr_to_string(error));
   }
}

static void
port_iterator_init(struct port_iterator *pi, struct vconn *vconn)
{
   memset(pi, 0, sizeof *pi);
   pi->vconn = vconn;
   if (vconn_get_version(vconn) < OFP13_VERSION) {
       port_iterator_fetch_features(pi);
   } else {
       port_iterator_fetch_port_desc(pi);
   }
}

static bool
port_iterator_next(struct port_iterator *pi, struct ofputil_phy_port *pp)
{
   for (;;) {
       if (pi->reply) {
           int retval = ofputil_pull_phy_port(vconn_get_version(pi->vconn),
                                              pi->reply, pp);
           if (!retval) {
               return true;
           } else if (retval != EOF) {
               ovs_fatal(0, "received bad reply: %s",
                         ofp_to_string(pi->reply->data, pi->reply->size,
                                       NULL, NULL, verbosity + 1));
           }
       }

       if (pi->variant == PI_FEATURES || !pi->more) {
           return false;
       }

       ovs_be32 recv_xid;
       do {
           ofpbuf_delete(pi->reply);
           run(vconn_recv_block(pi->vconn, &pi->reply),
               "OpenFlow receive failed");
           recv_xid = ((struct ofp_header *) pi->reply->data)->xid;
       } while (pi->send_xid != recv_xid);

       struct ofp_header *oh = pi->reply->data;
       enum ofptype type;
       if (ofptype_pull(&type, pi->reply)
           || type != OFPTYPE_PORT_DESC_STATS_REPLY) {
           ovs_fatal(0, "received bad reply: %s",
                     ofp_to_string(pi->reply->data, pi->reply->size, NULL,
                                   NULL, verbosity + 1));
       }

       pi->more = (ofpmp_flags(oh) & OFPSF_REPLY_MORE) != 0;
   }
}

static void
port_iterator_destroy(struct port_iterator *pi)
{
   if (pi) {
       while (pi->variant == PI_PORT_DESC && pi->more) {
           // Drain vconn's queue of any other replies for this request.
           struct ofputil_phy_port pp;
           port_iterator_next(pi, &pp);
       }

       ofpbuf_delete(pi->reply);
   }
}

static const struct ofputil_port_map *
get_port_map(const char *vconn_name)
{
   static struct shash port_maps = SHASH_INITIALIZER(&port_maps);
   struct ofputil_port_map *map = shash_find_data(&port_maps, vconn_name);
   if (!map) {
       map = xmalloc(sizeof *map);
       ofputil_port_map_init(map);
       shash_add(&port_maps, vconn_name, map);

       if (!strchr(vconn_name, ':') || !vconn_verify_name(vconn_name)) {
			// For an active vconn (which includes a vconn constructed from a
           // bridge name), connect to it and pull down the port name-number
           // mapping.
			struct vconn *vconn;
           open_vconn(vconn_name, &vconn);

           struct port_iterator pi;
           struct ofputil_phy_port pp;
           for (port_iterator_init(&pi, vconn);
                port_iterator_next(&pi, &pp); ) {
               ofputil_port_map_put(map, pp.port_no, pp.name);
           }
           port_iterator_destroy(&pi);

           vconn_close(vconn);
       } else {
			// Don't bother with passive vconns, since it could take a long
           // time for the remote to try to connect to us.  Don't bother with
           // invalid vconn names either.
		}
   }
   return map;
}

ofp_port_t get_ofpact_output_port(struct ofpact_output *output) {
    return output->port;
}
*/
import "C"

const OFP_MAX_PORT_NAME_LEN uint8 = C.OFP_MAX_PORT_NAME_LEN

type FlowPort uint16

/* Returns the OpenFlow 1.1+ port number equivalent to the OpenFlow 1.0 port
 * number 'ofp10_port', for encoding OpenFlow 1.1+ protocol messages.
 *
 * See the definition of OFP11_MAX for an explanation of the mapping. */
//ovs_be32
//ofputil_port_to_ofp11(ofp_port_t ofp10_port)
//{
//return htonl(ofp_to_u16(ofp10_port) < ofp_to_u16(OFPP_MAX)
//? ofp_to_u16(ofp10_port)
//: ofp_to_u16(ofp10_port) + OFPP11_OFFSET);
//}

func (fp *FlowPort) ofp11() uint32 {
	if fp == nil {
		return 0
	}

	if *fp < OFPP_MAX {
		return uint32(*fp)
	}

	return uint32(*fp) + OFPP11_OFFSET
}

const (
	/* Port number(s)   meaning
	 * ---------------  --------------------------------------
	 * 0x0000           not assigned a meaning by OpenFlow 1.0
	 * 0x0001...0xfeff  "physical" ports
	 * 0xff00...0xfff6  "reserved" but not assigned a meaning by OpenFlow 1.x
	 * 0xfff7...0xffff  "reserved" OFPP_* ports with assigned meanings
	 */

	/* Ranges. */

	OFPP_MAX        FlowPort = C.OFPP_MAX        /* Max # of switch ports. */
	OFPP_FIRST_RESV FlowPort = C.OFPP_FIRST_RESV /* First assigned reserved port. */
	OFPP_LAST_RESV  FlowPort = C.OFPP_LAST_RESV  /* Last assigned reserved port. */

	/* Reserved output "ports". */

	OFPP_UNSET   FlowPort = C.OFPP_UNSET   /* For OXM_OF_ACTSET_OUTPUT only. */
	OFPP_IN_PORT FlowPort = C.OFPP_IN_PORT /* Where the packet came in. */
	OFPP_TABLE   FlowPort = C.OFPP_TABLE   /* Perform actions in flow table. */
	OFPP_NORMAL  FlowPort = C.OFPP_NORMAL  /* Process with normal L2/L3. */
	OFPP_FLOOD   FlowPort = C.OFPP_FLOOD   /* All ports except input port and ports disabled by STP. */

	OFPP_ALL        FlowPort = C.OFPP_ALL        /* All ports except input port. */
	OFPP_CONTROLLER FlowPort = C.OFPP_CONTROLLER /* Send to controller. */
	OFPP_LOCAL      FlowPort = C.OFPP_LOCAL      /* Local openflow "port". */
	OFPP_NONE       FlowPort = C.OFPP_NONE       /* Not associated with any port. */
)

const (
	/* OpenFlow 1.1 uses 32-bit port numbers.  Open vSwitch, for now, uses OpenFlow
	 * 1.0 port numbers internally.  We map them to OpenFlow 1.0 as follows:
	 *
	 * OF1.1                    <=>  OF1.0
	 * -----------------------       ---------------
	 * 0x00000000...0x0000feff  <=>  0x0000...0xfeff  "physical" ports
	 * 0x0000ff00...0xfffffeff  <=>  not supported
	 * 0xffffff00...0xffffffff  <=>  0xff00...0xffff  "reserved" OFPP_* ports
	 *
	 * OFPP11_OFFSET is the value that must be added or subtracted to convert
	 * an OpenFlow 1.0 reserved port number to or from, respectively, the
	 * corresponding OpenFlow 1.1 reserved port number.
	 */

	OFPP11_MAX    uint32 = C.OFPP11_MAX
	OFPP11_OFFSET uint32 = C.OFPP11_OFFSET /* OFPP11_MAX - OFPP_MAX */
)

func getPortMap(connection *Connection) *C.struct_ofputil_port_map {
	return C.get_port_map(connection.name)
}
