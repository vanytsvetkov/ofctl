package ofctl

/*
#cgo CFLAGS: -I/usr/local/include -I/usr/src/ovs -I/usr/src/ovs/include -I/usr/src/ovs/lib
#cgo LDFLAGS: -L/usr/local/lib -lopenvswitch -lofproto -lm

#include "include/ofctl.h"
#include "include/connection.h"

VLOG_DEFINE_THIS_MODULE(ofctl);

extern void
run(int retval, const char *message, ...)
{
   if (retval) {
       va_list args;

       va_start(args, message);
       ovs_fatal_valist(retval, message, args);
   }
}

extern void
send_openflow_buffer(struct vconn *vconn, struct ofpbuf *buffer)
{
    run(vconn_send_block(vconn, buffer), "failed to send packet to switch");
}

*/
import "C"

const (
	// OFP_FLOW_PERMANENT == 0
	// Value used in "idle_timeout" and "hard_timeout" to indicate that the entry is permanent.
	OFP_FLOW_PERMANENT uint16 = C.OFP_FLOW_PERMANENT

	// OFP_DEFAULT_PRIORITY == 0x8000
	// By default, choose a priority in the middle.
	OFP_DEFAULT_PRIORITY uint16 = C.OFP_DEFAULT_PRIORITY
)

func init() {
	SetVersion(OFP13_VERSION)
}

/*
// Run вызывает ovs-ofctl с заданными аргументами
func Run(args ...string) (string, error) {
	args = append([]string{"ovs-ofctl"}, args...)

	cArgs := make([]*C.char, len(args))
	for i, arg := range args {
		cArgs[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(cArgs[i]))
	}

	var res int
	stdout, stderr := catchStd(func() {
		//res = int(C.ofctl_run(C.int(len(args)), &cArgs[0]))
		res = 0
	})
	if res != 0 {
		return "", fmt.Errorf("failed to run ofctl")
	}

	if stderr != "" {
		return stdout, fmt.Errorf(stderr)
	}

	return stdout, nil
}

func run(args ...string) (string, error) {
	args = append([]string{"ovs-ofctl"}, args...)

	argc := C.int(len(args))
	argv := make([]*C.char, len(args))
	for i, arg := range args {
		argv[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(argv[i]))
	}

	C.parse_options(argc, &argv[0])

	var ctx *C.struct_ovs_cmdl_context
	{
		size := C.size_t(unsafe.Sizeof(C.struct_ovs_cmdl_context{}))
		ctx = (*C.struct_ovs_cmdl_context)(C.malloc(size))
		defer C.free(unsafe.Pointer(ctx))
	}

	ctx.argc = argc - C.optind
	ctx.argv = &argv[C.optind]

	stdout, stderr := catchStd(func() {
		C.ovs_cmdl_run_command(ctx, C.get_all_commands())
	})

	if stderr != "" {
		return stdout, fmt.Errorf(stderr)
	}

	return stdout, nil
}
*/
