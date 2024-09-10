#include <config.h>

#include <stdlib.h>
#include <unistd.h>

#include <dirs.h>

#include <dirs.h>
#include <ofp-version-opt.h>
#include <ofproto/ofproto.h>
#include <openvswitch/dynamic-string.h>
#include <openvswitch/ofp-actions.h>
#include <openvswitch/ofp-port.h>
#include <openvswitch/ofp-util.h>
#include <openvswitch/ofp-print.h>
#include <openvswitch/ofp-msgs.h>
#include <openvswitch/ofp-switch.h>
#include <openvswitch/ofp-table.h>
#include <openvswitch/ofp-flow.h>
#include <openvswitch/meta-flow.h>
#include <openvswitch/ofpbuf.h>
#include <openvswitch/vconn.h>
#include <openvswitch/vlog.h>
#include <openvswitch/shash.h>

#include <socket-util.h>
#include <util.h>

// todo: delete it!

static int verbosity;

void
run(int retval, const char *message, ...);

void
send_openflow_buffer(struct vconn *vconn, struct ofpbuf *buffer);
