
struct port_iterator {
   struct vconn *vconn;

   enum { PI_FEATURES, PI_PORT_DESC } variant;
   struct ofpbuf *reply;
   ovs_be32 send_xid;
   bool more;
};

static void
port_iterator_fetch_port_desc(struct port_iterator *pi);

static void
port_iterator_fetch_features(struct port_iterator *pi);

static void
port_iterator_init(struct port_iterator *pi, struct vconn *vconn);

static bool
port_iterator_next(struct port_iterator *pi, struct ofputil_phy_port *pp);

static void
port_iterator_destroy(struct port_iterator *pi);

static const struct ofputil_port_map *
get_port_map(const char *vconn_name);

ofp_port_t
get_ofpact_output_port(struct ofpact_output *output);
