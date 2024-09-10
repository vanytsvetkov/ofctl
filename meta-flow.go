package ofctl

/*
#include "include/ofctl.h"
*/
import "C"
import "fmt"

type MetaFlowFieldID uint8

const (
	/* ## -------- ## */
	/* ## Metadata ## */
	/* ## -------- ## */

	/* "ct_label".
	 *
	 * Connection tracking label.  The label is carried with the
	 * connection tracking state.  On Linux this is held in the
	 * conntrack label extension but the exact implementation is
	 * platform-dependent.
	 *
	 * Writable only from nested actions within the NXAST_CT action.
	 *
	 * Type: be128.
	 * Maskable: bitwise.
	 * Formatting: hexadecimal.
	 * Prerequisites: none.
	 * Access: read/write.
	 * NXM: NXM_NX_CT_LABEL(108) since v2.5.
	 * OXM: none.
	 */

	MFF_CT_LABEL MetaFlowFieldID = C.MFF_CT_LABEL

	/* "reg<N>".
	 *
	 * Nicira extension scratch pad register with initial value 0.
	 *
	 * Type: be32.
	 * Maskable: bitwise.
	 * Formatting: hexadecimal.
	 * Prerequisites: none.
	 * Access: read/write.
	 * NXM: NXM_NX_REG0(0) since v1.1.        <0>
	 * NXM: NXM_NX_REG1(1) since v1.1.        <1>
	 * NXM: NXM_NX_REG2(2) since v1.1.        <2>
	 * NXM: NXM_NX_REG3(3) since v1.1.        <3>
	 * NXM: NXM_NX_REG4(4) since v1.3.        <4>
	 * NXM: NXM_NX_REG5(5) since v1.7.        <5>
	 * NXM: NXM_NX_REG6(6) since v1.7.        <6>
	 * NXM: NXM_NX_REG7(7) since v1.7.        <7>
	 * NXM: NXM_NX_REG8(8) since v2.6.        <8>
	 * NXM: NXM_NX_REG9(9) since v2.6.        <9>
	 * NXM: NXM_NX_REG10(10) since v2.6.      <10>
	 * NXM: NXM_NX_REG11(11) since v2.6.      <11>
	 * NXM: NXM_NX_REG12(12) since v2.6.      <12>
	 * NXM: NXM_NX_REG13(13) since v2.6.      <13>
	 * NXM: NXM_NX_REG14(14) since v2.6.      <14>
	 * NXM: NXM_NX_REG15(15) since v2.6.      <15>
	 * OXM: none.
	 */

	MFF_REG0  MetaFlowFieldID = C.MFF_REG0
	MFF_REG1  MetaFlowFieldID = C.MFF_REG1
	MFF_REG2  MetaFlowFieldID = C.MFF_REG2
	MFF_REG3  MetaFlowFieldID = C.MFF_REG3
	MFF_REG4  MetaFlowFieldID = C.MFF_REG4
	MFF_REG5  MetaFlowFieldID = C.MFF_REG5
	MFF_REG6  MetaFlowFieldID = C.MFF_REG6
	MFF_REG7  MetaFlowFieldID = C.MFF_REG7
	MFF_REG8  MetaFlowFieldID = C.MFF_REG8
	MFF_REG9  MetaFlowFieldID = C.MFF_REG9
	MFF_REG10 MetaFlowFieldID = C.MFF_REG10
	MFF_REG11 MetaFlowFieldID = C.MFF_REG11
	MFF_REG12 MetaFlowFieldID = C.MFF_REG12
	MFF_REG13 MetaFlowFieldID = C.MFF_REG13
	MFF_REG14 MetaFlowFieldID = C.MFF_REG14
	MFF_REG15 MetaFlowFieldID = C.MFF_REG15

	/* "in_port".
	 *
	 * 16-bit (OpenFlow 1.0) view of the physical or virtual port on which the
	 * packet was received.
	 *
	 * Type: be16.
	 * Maskable: no.
	 * Formatting: OpenFlow 1.0 port.
	 * Prerequisites: none.
	 * Access: read/write.
	 * NXM: NXM_OF_IN_PORT(0) since v1.1.
	 * OXM: none.
	 * OF1.0: exact match.
	 * OF1.1: exact match.
	 */

	MFF_IN_PORT MetaFlowFieldID = C.MFF_IN_PORT

	/* ## -------- ## */
	/* ## Ethernet ## */
	/* ## -------- ## */

	/* "eth_src" (aka "dl_src").
	 *
	 * Source address in Ethernet header.
	 *
	 * Type: MAC.
	 * Maskable: bitwise.
	 * Formatting: Ethernet.
	 * Prerequisites: Ethernet.
	 * Access: read/write.
	 * NXM: NXM_OF_ETH_SRC(2) since v1.1.
	 * OXM: OXM_OF_ETH_SRC(4) since OF1.2 and v1.7.
	 * OF1.0: exact match.
	 * OF1.1: bitwise mask.
	 */

	MFF_ETH_SRC MetaFlowFieldID = C.MFF_ETH_SRC

	/* "eth_dst" (aka "dl_dst").
	 *
	 * Destination address in Ethernet header.
	 *
	 * Type: MAC.
	 * Maskable: bitwise.
	 * Formatting: Ethernet.
	 * Prerequisites: Ethernet.
	 * Access: read/write.
	 * NXM: NXM_OF_ETH_DST(1) since v1.1.
	 * OXM: OXM_OF_ETH_DST(3) since OF1.2 and v1.7.
	 * OF1.0: exact match.
	 * OF1.1: bitwise mask.
	 */

	MFF_ETH_DST MetaFlowFieldID = C.MFF_ETH_DST

	/* ## ---- ## */
	/* ## VLAN ## */
	/* ## ---- ## */

	/* It looks odd for vlan_tci, vlan_vid, and vlan_pcp to say that they are
	 * supported in OF1.0 and OF1.1, since the detailed semantics of these fields
	 * only apply to NXM or OXM.  They are marked as supported for exact matches in
	 * OF1.0 and OF1.1 because exact matches on those fields can be successfully
	 * translated into the OF1.0 and OF1.1 flow formats. */

	/* "vlan_tci".
	 *
	 * 802.1Q TCI.
	 *
	 * For a packet with an 802.1Q header, this is the Tag Control Information
	 * (TCI) field, with the CFI bit forced to 1.  For a packet with no 802.1Q
	 * header, this has value 0.
	 *
	 * Type: be16.
	 * Maskable: bitwise.
	 * Formatting: hexadecimal.
	 * Prerequisites: Ethernet.
	 * Access: read/write.
	 * NXM: NXM_OF_VLAN_TCI(4) since v1.1.
	 * OXM: none.
	 * OF1.0: exact match.
	 * OF1.1: exact match.
	 */

	MFF_VLAN_TCI MetaFlowFieldID = C.MFF_VLAN_TCI
)

type MetaFlowValue uint64

// String returns the value as a hexadecimal string.
func (value *MetaFlowValue) String() string {
	return fmt.Sprintf("0x%x", *value)
}

func newMetaFlowValue(value C.uint64_t) *MetaFlowValue {
	tmp := MetaFlowValue(value)
	return &tmp
}

type MetaFlowField struct {
	ID          MetaFlowFieldID
	Name        string
	ExtraName   string
	bytes       uint
	bits        uint
	variableLen bool
	maskable    uint8
	// ...
	usableProtocolsExact   uint32
	usableProtocolsCIRD    uint32
	usableProtocolsBitwise uint32
}

func (field *MetaFlowField) String() string {
	switch field.ID {
	case MFF_CT_LABEL:
		return "NXM_NX_CT_LABEL"
	case MFF_VLAN_TCI:
		return "NXM_OF_VLAN_TCI"
	case MFF_ETH_SRC:
		return "NXM_OF_ETH_SRC"
	case MFF_ETH_DST:
		return "NXM_OF_ETH_DST"
	case MFF_IN_PORT:
		return "NXM_OF_IN_PORT"
	case MFF_REG0, MFF_REG1, MFF_REG2, MFF_REG3,
		MFF_REG4, MFF_REG5, MFF_REG6, MFF_REG7,
		MFF_REG8, MFF_REG9, MFF_REG10, MFF_REG11,
		MFF_REG12, MFF_REG13, MFF_REG14, MFF_REG15:
		return fmt.Sprintf("NXM_NX_REG%d", field.ID-MFF_REG0)
	}

	return "<unknown>"
}

func newMetaFlowField(cField *C.struct_mf_field) *MetaFlowField {
	if cField == nil {
		return nil
	}

	return &MetaFlowField{
		ID:          MetaFlowFieldID(cField.id),
		Name:        C.GoString(cField.name),
		ExtraName:   C.GoString(cField.extra_name),
		bytes:       uint(cField.n_bytes),
		bits:        uint(cField.n_bits),
		variableLen: bool(cField.variable_len),
		maskable:    uint8(cField.maskable),
		// ...
		usableProtocolsExact:   uint32(cField.usable_protocols_exact),
		usableProtocolsCIRD:    uint32(cField.usable_protocols_cidr),
		usableProtocolsBitwise: uint32(cField.usable_protocols_bitwise),
	}
}

type MetaFlowSubfield struct {
	Field  *MetaFlowField
	offset uint
	bits   uint
}

func (subfield *MetaFlowSubfield) String() string {
	if subfield.Field == nil {
		return "<unknown>"
	}

	str := fmt.Sprintf("%s", subfield.Field)
	if subfield.offset == 0 && subfield.bits == subfield.Field.bits {
		str += "[]"
	} else if subfield.bits == 1 {
		str += fmt.Sprintf("[%d]", subfield.offset)
	} else {
		str += fmt.Sprintf("[%d..%d]", subfield.offset, subfield.offset+subfield.bits-1)
	}

	return str
}

func newMetaFlowSubfield(cSubfield *C.struct_mf_subfield) *MetaFlowSubfield {
	return &MetaFlowSubfield{
		Field:  newMetaFlowField(cSubfield.field),
		offset: uint(cSubfield.ofs),
		bits:   uint(cSubfield.n_bits),
	}
}
