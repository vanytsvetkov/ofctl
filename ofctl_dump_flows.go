package ofctl

/*
#include "include/ofctl.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

/*
type dumpFlowsSettings struct {
	match FlowMatch
}

func DumpFlowsWithTableID(id uint8) func(setting *dumpFlowsSettings) {
	return func(setting *dumpFlowsSettings) {
		setting.match = fmt.Sprintf("table=%d", id)
	}
*/

// DumpFlows fetches and returns the flow statistics from the bridge.
func DumpFlows(bridge string) (flowStats FlowStats, err error) {
	/*
		opts ...func(setting *dumpFlowsSettings)

		setting := &dumpFlowsSettings{
			match: make([]string, 0),
			//version: OFP10_VERSION,
		}

		for _, opt := range opts {
			opt(setting)
		}
	*/

	// todo: handle options!

	connection := NewConnection()
	if err = connection.Open(bridge); err != nil {
		return nil, err
	}
	defer connection.Close()

	cMatch := C.CString("")
	defer C.free(unsafe.Pointer(cMatch))

	// Prepare flow stats request

	var cPortMap *C.struct_ofputil_port_map
	{
		cPortMap = getPortMap(connection)
		defer C.free(unsafe.Pointer(cPortMap))
	}

	var cTableMap *C.struct_ofputil_table_map
	{
		cTableMap = getTableMap(connection)
		defer C.free(unsafe.Pointer(cTableMap))
	}

	var cRequest C.struct_ofputil_flow_stats_request
	C.memset(unsafe.Pointer(&cRequest), 0, C.sizeof_struct_ofputil_flow_stats_request)
	{
		var cUsableProtocols C.enum_ofputil_protocol
		errMessage := C.parse_ofp_flow_stats_request_str(
			&cRequest, C.bool(false), cMatch, cPortMap, cTableMap, &cUsableProtocols,
		)
		if errMessage != nil {
			return nil, fmt.Errorf("failed to parse flow stats request: %s", C.GoString(errMessage))
		}

		// Set protocol for dump flows
		usableProtocols := OpenFlowProtocol(cUsableProtocols)
		allowedProtocols := OFP_ANY
		for _, ofp := range []OpenFlowProtocol{
			OFP15_OXM, OFP14_OXM, OFP13_OXM, OFP12_OXM, OFP11_STD, OFP10_NXM, OFP10_STD,
		} {
			if ofp&usableProtocols&allowedProtocols != 0 {
				err = connection.SetProtocol(ofp)
				if err == nil {
					break
				}
			}
		}
	}

	// Fetch flow stats
	var n_fses C.size_t
	var fses *C.struct_ofputil_flow_stats
	{
		fses = (*C.struct_ofputil_flow_stats)(C.malloc(C.sizeof_struct_ofputil_flow_stats))
		defer C.free(unsafe.Pointer(fses))
		defer func() {
			for i := C.size_t(0); i < n_fses; i++ {
				C.free(
					unsafe.Pointer(
						(*C.struct_ofputil_flow_stats)(
							unsafe.Pointer(uintptr(unsafe.Pointer(fses)) + uintptr(i)*unsafe.Sizeof(*fses)),
						).ofpacts,
					),
				)
			}
		}()
	}

	connection.dumpFlows(&cRequest, &fses, &n_fses)

	// Convert the C structs to Go structs
	flowStats = make([]*FlowStat, int(n_fses))
	for i := 0; i < int(n_fses); i++ {
		flowStats[i] = newFlowStat(
			(*C.struct_ofputil_flow_stats)(unsafe.Pointer(uintptr(unsafe.Pointer(fses))+uintptr(i)*unsafe.Sizeof(*fses))),
			cPortMap, cTableMap,
		)
	}

	return flowStats, nil
}

func (conn *Connection) dumpFlows(request *C.struct_ofputil_flow_stats_request, reply **C.struct_ofputil_flow_stats, length *C.size_t) {
	C.vconn_dump_flows(conn.ptr, request, conn.protocol, reply, length)
}
