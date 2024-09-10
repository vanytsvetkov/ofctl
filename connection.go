package ofctl

/*
#include "include/ofctl.h"
#include "include/connection.h"

VLOG_DEFINE_THIS_MODULE(ofctl_connection);

extern int
open_vconn_socket(const char *name, struct vconn **vconnp)
{
   char *vconn_name = xasprintf("unix:%s", name);
   int error;

   error = vconn_open(vconn_name, get_allowed_ofp_versions(), DSCP_DEFAULT,
                      vconnp);
   if (error && error != ENOENT) {
       ovs_fatal(0, "%s: failed to open socket (%s)", name,
                 ovs_strerror(error));
   }
   free(vconn_name);

   return error;
}

extern enum ofputil_protocol
open_vconn__(const char *name, enum open_target target,
            struct vconn **vconnp)
{
   const char *suffix = target == MGMT ? "mgmt" : "snoop";
   char *datapath_name, *datapath_type, *socket_name;
   enum ofputil_protocol protocol;
   char *bridge_path;
   int ofp_version;
   int error;

   bridge_path = xasprintf("%s/%s.%s", ovs_rundir(), name, suffix);

   ofproto_parse_name(name, &datapath_name, &datapath_type);
   socket_name = xasprintf("%s/%s.%s", ovs_rundir(), datapath_name, suffix);
   free(datapath_name);
   free(datapath_type);

   if (strchr(name, ':')) {
       run(vconn_open(name, get_allowed_ofp_versions(), DSCP_DEFAULT, vconnp),
           "connecting to %s", name);
   } else if (!open_vconn_socket(name, vconnp)) {
       // Fall Through.
   } else if (!open_vconn_socket(bridge_path, vconnp)) {
       // Fall Through.
   } else if (!open_vconn_socket(socket_name, vconnp)) {
       // Fall Through.
   } else {
       free(bridge_path);
       free(socket_name);
       ovs_fatal(0, "%s is not a bridge or a socket", name);
   }

   if (target == SNOOP) {
       vconn_set_recv_any_version(*vconnp);
   }

   free(bridge_path);
   free(socket_name);

   VLOG_DBG("connecting to %s", vconn_get_name(*vconnp));
   error = vconn_connect_block(*vconnp, -1);
   if (error) {
       ovs_fatal(0, "%s: failed to connect to socket (%s)", name,
                 ovs_strerror(error));
   }

   ofp_version = vconn_get_version(*vconnp);
   protocol = ofputil_protocol_from_ofp_version(ofp_version);
   if (!protocol) {
       ovs_fatal(0, "%s: unsupported OpenFlow version 0x%02x",
                 name, ofp_version);
   }
   return protocol;
}

extern enum ofputil_protocol
open_vconn(const char *name, struct vconn **vconnp)
{
   return open_vconn__(name, MGMT, vconnp);
}
*/
import "C"
import (
	"fmt"
	"strings"
	"unsafe"
)

type openTarget string

const (
	MGMT  openTarget = "mgmt"
	SNOOP openTarget = "snoop"
)

// Connection представляет соединение с Open vSwitch
type Connection struct {
	ptr *C.struct_vconn

	name     *C.char
	version  C.enum_ofp_version
	protocol C.enum_ofputil_protocol

	target openTarget
}

func (conn *Connection) Name() string {
	return C.GoString(conn.name)
}

func (conn *Connection) Version() OpenFlowVersion {
	return OpenFlowVersion(conn.version)
}

func (conn *Connection) Protocol() OpenFlowProtocol {
	return OpenFlowProtocol(conn.protocol)
}

// NewConnection открывает соединение с указанным коммутатором Open vSwitch
//func NewConnection(bridge string) (*Connection, error) {
//	var vconn *C.struct_vconn
//	socketName := fmt.Sprintf(ovsSocketPathTemplate, bridge)
//	cSocketName := C.CString(socketName)
//	defer C.free(unsafe.Pointer(cSocketName))
//
//	ret := C.vconn_open(cSocketName, ofpVersion, 0, &vconn)
//	if ret != 0 {
//		return nil, fmt.Errorf("failed to open vconn: %v", ret)
//	}
//
//	return &Connection{ptr: vconn}, nil
//}

// NewConnection открывает новое соединение с указанным bridge
func NewConnection() *Connection {
	return &Connection{
		target: MGMT,
	}
}

// Open открывает соединение с OpenvSwitch
func (conn *Connection) Open(name string) (err error) {
	if strings.Contains(name, "unix:") {
		cName := C.CString(name)

		ret := C.vconn_open(cName, allowedVersion(), C.DSCP_DEFAULT, &conn.ptr)
		if ret != 0 {
			return fmt.Errorf("%s: failed to open socket (%s)", name, C.GoString(C.ovs_strerror(ret)))
		}

		conn.name = cName

	} else if err = conn.Open(fmt.Sprintf("unix:%s", name)); err == nil {
		/* Fall Through. */
	} else if err = conn.Open(fmt.Sprintf("unix:%s/%s.%s", C.GoString(C.ovs_rundir()), name, conn.target)); err == nil {
		/* Fall Through. */
	} else {
		return fmt.Errorf("%s is not a bridge or a socket", name)
	}

	// Дополнительные настройки для SNOOP
	// (если требуется, нужно будет добавить соответствующий параметр)

	// Устанавливаем соединение
	ret := C.vconn_connect_block(conn.ptr, -1)
	if ret != 0 {
		return fmt.Errorf("%s: failed to connect to socket (%s)", name, C.GoString(C.ovs_strerror(ret)))
	}

	conn.version = C.enum_ofp_version(C.vconn_get_version(conn.ptr))
	conn.protocol = C.ofputil_protocol_from_ofp_version(conn.version)
	if conn.protocol == 0 {
		return fmt.Errorf("%s: unsupported OpenFlow version 0x%02x", name, conn.version)
	}

	return nil
}

func (conn *Connection) Send(msg Message) error {
	bytes, err := msg.MarshalBinary()
	if err != nil {
		return err
	}

	fmt.Printf("0x%x\n", bytes)

	var cOfpbuf *C.struct_ofpbuf
	{
		cOfpbuf = (*C.struct_ofpbuf)(C.malloc(C.sizeof_struct_ofpbuf))
		defer C.free(unsafe.Pointer(cOfpbuf))

		cOfpbuf = C.ofpbuf_new(0)
		C.ofpbuf_put_zeros(cOfpbuf, C.size_t(len(bytes)))

		cBytes := C.CBytes(bytes)
		defer C.free(unsafe.Pointer(cBytes))

		cOfpbuf.data = cBytes
	}

	cTest := C.ofp_to_string(
		cOfpbuf.data,
		C.size_t(cOfpbuf.size),
		nil,
		nil,
		4,
	)

	fmt.Printf("%s\n", C.GoString(cTest))

	err = conn.send(cOfpbuf)
	if err != nil {
		return err
	}

	return nil
}

func (conn *Connection) send(cOfpbuf *C.struct_ofpbuf) error {
	ret := C.vconn_send_block(conn.ptr, cOfpbuf)
	if ret != 0 {
		return fmt.Errorf("failed to send message") // todo: prettify error
	}

	return nil
}

func (conn *Connection) receive() (cReply *C.struct_ofpbuf, err error) {
	ret := C.vconn_recv_block(conn.ptr, &cReply)
	if ret != 0 {
		return cReply, fmt.Errorf("failed to receive message")
	}

	return cReply, nil
}

// Close закрывает соединение с OpenvSwitch
func (conn *Connection) Close() {
	if conn.ptr != nil {
		C.vconn_close(conn.ptr)
		conn.ptr = nil
	}

	if conn.name != nil {
		C.free(unsafe.Pointer(conn.name))
		conn.name = nil
	}
}

func (conn *Connection) SetProtocol(want OpenFlowProtocol) (err error) {
	cur := conn.protocol
	defer func() {
		if err != nil {
			conn.protocol = cur
		}
	}()

	var request, reply *C.struct_ofpbuf
	var next C.enum_ofputil_protocol

	request = C.ofputil_encode_set_protocol(cur, C.enum_ofputil_protocol(want), &next)
	if request == nil {
		return nil // i.e. conn.protocol == want
	}

	C.vconn_transact_noreply(conn.ptr, request, &reply)

	if reply != nil {
		errMessage := C.ofp_to_string(
			reply.data, C.size_t(reply.size),
			(*C.struct_ofputil_port_map)(C.NULL), (*C.struct_ofputil_table_map)(C.NULL), 2,
		)
		if errMessage != nil {
			return fmt.Errorf("%s: failed to set protocol, switch replied: %s", conn.Name(), C.GoString(errMessage))
		}

		return fmt.Errorf("%s: failed to set protocol, switch not replied", conn.Name())
	}

	conn.protocol = next

	return nil
}

// multiTransact performs multi transactions to socket without reply.
func (conn *Connection) multiTransact(cRequests *C.struct_ovs_list) error {
	var cReply *C.struct_ofpbuf
	defer C.ofpbuf_delete(cReply)

	C.vconn_transact_multiple_noreply(conn.ptr, cRequests, &cReply)

	if cReply != nil {
		errorMessage := C.ofp_to_string(
			cReply.data,
			C.size_t(cReply.size),
			nil,
			nil,
			0,
		)
		return fmt.Errorf("%s: failed to transact: %s", conn.Name(), C.GoString(errorMessage))
	}

	return nil
}

// transact performs a single transaction to socket without a reply.
func (conn *Connection) transact(cRequest *C.struct_ofpbuf) error {
	var cRequests *C.struct_ovs_list
	{
		cRequests = (*C.struct_ovs_list)(C.malloc(C.sizeof_struct_ovs_list))

		C.ovs_list_init(cRequests)
		C.ovs_list_push_back(cRequests, &cRequest.list_node)

		defer C.free(unsafe.Pointer(cRequests))
	}

	err := conn.multiTransact(cRequests)
	if err != nil {
		return err
	}

	return nil
}
