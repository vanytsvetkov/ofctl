enum open_target { MGMT, SNOOP };

extern int
open_vconn_socket(const char *name, struct vconn **vconnp);

extern enum ofputil_protocol
open_vconn__(const char *name, enum open_target target,
            struct vconn **vconnp);

extern enum ofputil_protocol
open_vconn(const char *name, struct vconn **vconnp);