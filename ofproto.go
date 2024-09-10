package ofctl

/*
#include "include/ofctl.h"
*/
import "C"

type OpenFlowVersion uint8

const (
	OFP10_VERSION OpenFlowVersion = C.OFP10_VERSION
	OFP11_VERSION OpenFlowVersion = C.OFP11_VERSION
	OFP12_VERSION OpenFlowVersion = C.OFP12_VERSION
	OFP13_VERSION OpenFlowVersion = C.OFP13_VERSION
	OFP14_VERSION OpenFlowVersion = C.OFP14_VERSION
	OFP15_VERSION OpenFlowVersion = C.OFP15_VERSION
)

var version OpenFlowVersion

func allowedVersion() C.uint32_t {
	if version != 0 {
		return C.uint32_t(
			1 << version,
		)
	}

	return C.get_allowed_ofp_versions()
}

func SetVersion(v OpenFlowVersion) {
	version = v
}

type OpenFlowProtocol uint16

const (
	OFP_NONE OpenFlowProtocol = C.OFPUTIL_P_NONE

	/* OpenFlow 1.0 protocols.
	 *
	 * The "STD" protocols use the standard OpenFlow 1.0 flow format.
	 * The "NXM" protocols use the Nicira Extensible Match (NXM) flow format.
	 *
	 * The protocols with "TID" mean that the nx_flow_mod_table_id Nicira
	 * extension has been enabled.  The other protocols have it disabled.
	 */

	OFP10_STD     OpenFlowProtocol = C.OFPUTIL_P_OF10_STD
	OFP10_STD_TID OpenFlowProtocol = C.OFPUTIL_P_OF10_STD_TID
	OFP10_NXM     OpenFlowProtocol = C.OFPUTIL_P_OF10_NXM
	OFP10_NXM_TID OpenFlowProtocol = C.OFPUTIL_P_OF10_NXM_TID

	/* OpenFlow 1.1 protocol.
	 *
	 * We only support the standard OpenFlow 1.1 flow format.
	 *
	 * OpenFlow 1.1 always operates with an equivalent of the
	 * nx_flow_mod_table_id Nicira extension enabled, so there is no "TID"
	 * variant.
	 */

	OFP11_STD OpenFlowProtocol = C.OFPUTIL_P_OF11_STD

	/* OpenFlow 1.2+ protocols (only one variant each).
	 *
	 * These use the standard OpenFlow Extensible Match (OXM) flow format.
	 *
	 * OpenFlow 1.2+ always operates with an equivalent of the
	 * nx_flow_mod_table_id Nicira extension enabled, so there is no "TID"
	 * variant.
	 */

	OFP12_OXM OpenFlowProtocol = C.OFPUTIL_P_OF12_OXM
	OFP13_OXM OpenFlowProtocol = C.OFPUTIL_P_OF13_OXM
	OFP14_OXM OpenFlowProtocol = C.OFPUTIL_P_OF14_OXM
	OFP15_OXM OpenFlowProtocol = C.OFPUTIL_P_OF15_OXM

	/* All protocols. */

	OFP_ANY OpenFlowProtocol = C.OFPUTIL_P_ANY
)
