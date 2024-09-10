package ofctl

/*
#include "include/ofctl.h"
#include "include/connection.h"
#include "include/ofp-table.h"

struct table_iterator {
   struct vconn *vconn;

   enum { TI_STATS, TI_FEATURES } variant;
   struct ofpbuf *reply;
   ovs_be32 send_xid;
   bool more;

   struct ofputil_table_features features;
   struct ofpbuf raw_properties;
};

static void
table_iterator_init(struct table_iterator *ti, struct vconn *vconn)
{
   memset(ti, 0, sizeof *ti);
   ti->vconn = vconn;
   ti->variant = (vconn_get_version(vconn) < OFP13_VERSION
                  ? TI_STATS : TI_FEATURES);
   ti->more = true;

   enum ofpraw ofpraw = (ti->variant == TI_STATS
                         ? OFPRAW_OFPST_TABLE_REQUEST
                         : OFPRAW_OFPST13_TABLE_FEATURES_REQUEST);
   struct ofpbuf *rq = ofpraw_alloc(ofpraw, vconn_get_version(vconn), 0);
   ti->send_xid = ((struct ofp_header *) rq->data)->xid;
   send_openflow_buffer(ti->vconn, rq);
}

static const struct ofputil_table_features *
table_iterator_next(struct table_iterator *ti)
{
   for (;;) {
       if (ti->reply) {
           int retval;
           if (ti->variant == TI_STATS) {
               struct ofputil_table_stats ts;
               retval = ofputil_decode_table_stats_reply(ti->reply,
                                                         &ts, &ti->features);
           } else {
               ovs_assert(ti->variant == TI_FEATURES);
               retval = ofputil_decode_table_features(ti->reply,
                                                      &ti->features,
                                                      &ti->raw_properties);
           }
           if (!retval) {
               return &ti->features;
           } else if (retval != EOF) {
               ovs_fatal(0, "received bad reply: %s",
                         ofp_to_string(ti->reply->data, ti->reply->size,
                                       NULL, NULL, verbosity + 1));
           }
       }

       if (!ti->more) {
           return NULL;
       }

       ovs_be32 recv_xid;
       do {
           ofpbuf_delete(ti->reply);
           run(vconn_recv_block(ti->vconn, &ti->reply),
               "OpenFlow receive failed");
           recv_xid = ((struct ofp_header *) ti->reply->data)->xid;
       } while (ti->send_xid != recv_xid);

       struct ofp_header *oh = ti->reply->data;
       enum ofptype type;
       if (ofptype_pull(&type, ti->reply)
           || type != (ti->variant == TI_STATS
                       ? OFPTYPE_TABLE_STATS_REPLY
                       : OFPTYPE_TABLE_FEATURES_STATS_REPLY)) {
           ovs_fatal(0, "received bad reply: %s",
                     ofp_to_string(ti->reply->data, ti->reply->size, NULL,
                                   NULL, verbosity + 1));
       }

       ti->more = (ofpmp_flags(oh) & OFPSF_REPLY_MORE) != 0;
   }
}

static void
table_iterator_destroy(struct table_iterator *ti)
{
   if (ti) {
       while (ti->more) {
           // Drain vconn's queue of any other replies for this request.
           table_iterator_next(ti);
       }

       ofpbuf_delete(ti->reply);
   }
}

static const struct ofputil_table_map *
get_table_map(const char *vconn_name)
{
   static struct shash table_maps = SHASH_INITIALIZER(&table_maps);
   struct ofputil_table_map *map = shash_find_data(&table_maps, vconn_name);
   if (!map) {
       map = xmalloc(sizeof *map);
       ofputil_table_map_init(map);
       shash_add(&table_maps, vconn_name, map);

       if (!strchr(vconn_name, ':') || !vconn_verify_name(vconn_name)) {
           // For an active vconn (which includes a vconn constructed from a
           // bridge name), connect to it and pull down the port name-number
           // mapping.
           struct vconn *vconn;
           open_vconn(vconn_name, &vconn);

           struct table_iterator ti;
           table_iterator_init(&ti, vconn);
           for (;;) {
               const struct ofputil_table_features *tf
                   = table_iterator_next(&ti);
               if (!tf) {
                   break;
               }
               if (tf->name[0]) {
                   ofputil_table_map_put(map, tf->table_id, tf->name);
               }
           }
           table_iterator_destroy(&ti);

           vconn_close(vconn);
       } else {
           // Don't bother with passive vconns, since it could take a long
           // time for the remote to try to connect to us.  Don't bother with
           // invalid vconn names either.
       }
   }
   return map;
}

uint8_t
get_ofpact_resubmit_table_id(struct ofpact_resubmit *resubmit) {
    return resubmit->table_id;
}
*/
import "C"

const (
	OFP_DEFAULT_TABLE uint8 = 0

	OFP_MAX_TABLE_NAME_LEN uint8 = C.OFP_MAX_TABLE_NAME_LEN // 32 // https://www.openvswitch.org//support/dist-docs/ovs-ofctl.8.txt
)

func getTableMap(connection *Connection) *C.struct_ofputil_table_map {
	return C.get_table_map(connection.name)
}
