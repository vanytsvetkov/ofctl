package ofctl

/*
#include "include/ofctl.h"
#include "include/ofp-port.h"
#include "include/ofp-table.h"

#include <ofp-actions.c>

// region OFPACT_OUTPUT_REG
struct mf_subfield
get_ofpact_output_reg_src(struct ofpact_output_reg *output_reg) {
    return output_reg->src;
}
// endregion

// region OFPACT_SET_FIELD

const struct mf_field *
get_ofpact_set_field_field(struct ofpact_set_field *set_field) {
	return set_field->field;
}

union mf_value *
get_ofpact_set_field_value(struct ofpact_set_field *set_field) {
	return set_field->value;
}

// endregion

// region OFPACT_REG_MOVE

struct mf_subfield
get_ofpact_reg_move_src(struct ofpact_reg_move *reg_move) {
    return reg_move->src;
}

struct mf_subfield
get_ofpact_reg_move_dst(struct ofpact_reg_move *reg_move) {
    return reg_move->dst;
}

// endregion

// region OFPACT_CONJUNCTION
uint32_t
get_ofpact_conjunction_id(struct ofpact_conjunction *conjunction) {
	return conjunction->id;
}

uint8_t
get_ofpact_conjunction_clause(struct ofpact_conjunction *conjunction) {
	return conjunction->clause;
}

uint8_t
get_ofpact_conjunction_n_clauses(struct ofpact_conjunction *conjunction) {
	return conjunction->n_clauses;
}
// endregion

// region OFPACT_LEARN
uint16_t
get_ofpact_learn_idle_timeout(struct ofpact_learn *learn) {
	return learn->idle_timeout;
}

uint16_t
get_ofpact_learn_hard_timeout(struct ofpact_learn *learn) {
	return learn->hard_timeout;
}

uint16_t
get_ofpact_learn_priority(struct ofpact_learn *learn) {
	return learn->priority;
}

uint8_t
get_ofpact_learn_table_id(struct ofpact_learn *learn) {
	return learn->table_id;
}

enum nx_learn_flags
get_ofpact_learn_flags(struct ofpact_learn *learn) {
	return learn->flags;
}

ovs_be64
get_ofpact_learn_cookie(struct ofpact_learn *learn) {
	return learn->cookie;
}

uint16_t
get_ofpact_learn_fin_idle_timeout(struct ofpact_learn *learn) {
	return learn->fin_idle_timeout;
}

uint16_t
get_ofpact_learn_fin_hard_timeout(struct ofpact_learn *learn) {
	return learn->fin_hard_timeout;
}

uint32_t
get_ofpact_learn_limit(struct ofpact_learn *learn) {
	return learn->limit;
}

struct mf_subfield
get_ofpact_learn_result_dst(struct ofpact_learn *learn) {
	return learn->result_dst;
}

struct ofpact_learn_spec *
get_ofpact_learn_specs(struct ofpact_learn *learn) {
	return learn->specs;
}

struct mf_subfield
get_ofpact_learn_spec_src(struct ofpact_learn_spec *spec) {
	return spec->src;
}

struct mf_subfield
get_ofpact_learn_spec_dst(struct ofpact_learn_spec *spec) {
	return spec->dst;
}

uint16_t
get_ofpact_learn_spec_src_type(struct ofpact_learn_spec *spec) {
	return spec->src_type;
}

uint16_t
get_ofpact_learn_spec_dst_type(struct ofpact_learn_spec *spec) {
	return spec->dst_type;
}

uint32_t
get_ofpact_learn_spec_n_bits(struct ofpact_learn_spec *spec) {
	return spec->n_bits;
}
// endregion

// region OFPACT_CT
uint16_t
get_ofpact_conntrack_flags(struct ofpact_conntrack *conntrack) {
	return conntrack->flags;
}

uint16_t
get_ofpact_conntrack_zone_imm(struct ofpact_conntrack *conntrack) {
	return conntrack->zone_imm;
}

struct mf_subfield
get_ofpact_conntrack_zone_src(struct ofpact_conntrack *conntrack) {
	return conntrack->zone_src;
}

uint16_t
get_ofpact_conntrack_alg(struct ofpact_conntrack *conntrack) {
	return conntrack->alg;
}

uint8_t
get_ofpact_conntrack_recirc_table(struct ofpact_conntrack *conntrack) {
	return conntrack->recirc_table;
}

struct ofpact *
get_ofpact_conntrack_actions(struct ofpact_conntrack *conntrack) {
	return conntrack->actions;
}

// endregion
*/
import "C"
import (
	"encoding/binary"
	"fmt"
	"strings"
	"unsafe"
)

type FlowActionType byte

const (
	/* Output. */

	OFPACT_OUTPUT     FlowActionType = C.OFPACT_OUTPUT
	OFPACT_GROUP      FlowActionType = C.OFPACT_GROUP
	OFPACT_CONTROLLER FlowActionType = C.OFPACT_CONTROLLER
	OFPACT_ENQUEUE    FlowActionType = C.OFPACT_ENQUEUE
	OFPACT_OUTPUT_REG FlowActionType = C.OFPACT_OUTPUT_REG
	OFPACT_BUNDLE     FlowActionType = C.OFPACT_BUNDLE

	/* Header changes. */

	OFPACT_SET_FIELD       FlowActionType = C.OFPACT_SET_FIELD
	OFPACT_SET_VLAN_VID    FlowActionType = C.OFPACT_SET_VLAN_VID
	OFPACT_SET_VLAN_PCP    FlowActionType = C.OFPACT_SET_VLAN_PCP
	OFPACT_STRIP_VLAN      FlowActionType = C.OFPACT_STRIP_VLAN
	OFPACT_PUSH_VLAN       FlowActionType = C.OFPACT_PUSH_VLAN
	OFPACT_SET_ETH_SRC     FlowActionType = C.OFPACT_SET_ETH_SRC
	OFPACT_SET_ETH_DST     FlowActionType = C.OFPACT_SET_ETH_DST
	OFPACT_SET_IPV4_SRC    FlowActionType = C.OFPACT_SET_IPV4_SRC
	OFPACT_SET_IPV4_DST    FlowActionType = C.OFPACT_SET_IPV4_DST
	OFPACT_SET_IP_DSCP     FlowActionType = C.OFPACT_SET_IP_DSCP
	OFPACT_SET_IP_ECN      FlowActionType = C.OFPACT_SET_IP_ECN
	OFPACT_SET_IP_TTL      FlowActionType = C.OFPACT_SET_IP_TTL
	OFPACT_SET_L4_SRC_PORT FlowActionType = C.OFPACT_SET_L4_SRC_PORT
	OFPACT_SET_L4_DST_PORT FlowActionType = C.OFPACT_SET_L4_DST_PORT
	OFPACT_REG_MOVE        FlowActionType = C.OFPACT_REG_MOVE
	OFPACT_STACK_PUSH      FlowActionType = C.OFPACT_STACK_PUSH
	OFPACT_STACK_POP       FlowActionType = C.OFPACT_STACK_POP
	OFPACT_DEC_TTL         FlowActionType = C.OFPACT_DEC_TTL
	OFPACT_SET_MPLS_LABEL  FlowActionType = C.OFPACT_SET_MPLS_LABEL
	OFPACT_SET_MPLS_TC     FlowActionType = C.OFPACT_SET_MPLS_TC
	OFPACT_SET_MPLS_TTL    FlowActionType = C.OFPACT_SET_MPLS_TTL
	OFPACT_DEC_MPLS_TTL    FlowActionType = C.OFPACT_DEC_MPLS_TTL
	OFPACT_PUSH_MPLS       FlowActionType = C.OFPACT_PUSH_MPLS
	OFPACT_POP_MPLS        FlowActionType = C.OFPACT_POP_MPLS
	OFPACT_DEC_NSH_TTL     FlowActionType = C.OFPACT_DEC_NSH_TTL
	OFPACT_DELETE_FIELD    FlowActionType = C.OFPACT_DELETE_FIELD

	/* Generic encap & decap */

	OFPACT_ENCAP FlowActionType = C.OFPACT_ENCAP
	OFPACT_DECAP FlowActionType = C.OFPACT_DECAP

	/* Metadata. */

	OFPACT_SET_TUNNEL  FlowActionType = C.OFPACT_SET_TUNNEL
	OFPACT_SET_QUEUE   FlowActionType = C.OFPACT_SET_QUEUE
	OFPACT_POP_QUEUE   FlowActionType = C.OFPACT_POP_QUEUE
	OFPACT_FIN_TIMEOUT FlowActionType = C.OFPACT_FIN_TIMEOUT

	/* Match table interaction. */

	OFPACT_RESUBMIT    FlowActionType = C.OFPACT_RESUBMIT
	OFPACT_LEARN       FlowActionType = C.OFPACT_LEARN
	OFPACT_CONJUNCTION FlowActionType = C.OFPACT_CONJUNCTION

	/* Arithmetic. */

	OFPACT_MULTIPATH FlowActionType = C.OFPACT_MULTIPATH

	/* Other. */

	OFPACT_NOTE             FlowActionType = C.OFPACT_NOTE
	OFPACT_EXIT             FlowActionType = C.OFPACT_EXIT
	OFPACT_SAMPLE           FlowActionType = C.OFPACT_SAMPLE
	OFPACT_UNROLL_XLATE     FlowActionType = C.OFPACT_UNROLL_XLATE
	OFPACT_CT               FlowActionType = C.OFPACT_CT
	OFPACT_CT_CLEAR         FlowActionType = C.OFPACT_CT_CLEAR
	OFPACT_NAT              FlowActionType = C.OFPACT_NAT
	OFPACT_OUTPUT_TRUNC     FlowActionType = C.OFPACT_OUTPUT_TRUNC
	OFPACT_CLONE            FlowActionType = C.OFPACT_CLONE
	OFPACT_CHECK_PKT_LARGER FlowActionType = C.OFPACT_CHECK_PKT_LARGER
)

type FlowRawActionType FlowActionType

const (
	NXAST_RAW_REG_LOAD FlowRawActionType = C.NXAST_RAW_REG_LOAD
)

type ActionHeader struct {
	Type   FlowActionType
	Length uint16
}

func (header *ActionHeader) Header() *ActionHeader {
	return header
}

func (header *ActionHeader) Len() (n uint16) {
	return 4
}

func (header *ActionHeader) MarshalBinary() (data []byte, err error) {
	data = make([]byte, header.Len())
	binary.BigEndian.PutUint16(data[:2], uint16(header.Type))
	binary.BigEndian.PutUint16(data[2:4], header.Length)
	return
}

func (header *ActionHeader) UnmarshalBinary(data []byte) error {
	if len(data) != 4 {
		return fmt.Errorf("the []byte the wrong size to unmarshal an ActionHeader message")
	}
	header.Type = FlowActionType(binary.BigEndian.Uint16(data[:2]))
	header.Length = binary.BigEndian.Uint16(data[2:4])
	return nil
}

// region OFPACT_OUTPUT

// OutputAction represents an OFPACT_OUTPUT action.
type OutputAction struct {
	ActionHeader
	PortID   FlowPort
	PortName string
}

func (action *OutputAction) String() string {
	var prefix string
	if action.PortID < OFPP_MAX {
		prefix = "output:"
	}

	if action.PortName != "" {
		return fmt.Sprintf("%s%s", prefix, action.PortName)
	}
	return fmt.Sprintf("%s%d", prefix, action.PortID)
}

func (action *OutputAction) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("action not implemented")
}

func newOutputAction(ptr *C.struct_ofpact_output, cPortMap *C.struct_ofputil_port_map) *OutputAction {
	var cPortID C.ofp_port_t = C.get_ofpact_output_port(ptr)
	var cPortName *C.char = C.CString("")
	{
		C.ofputil_port_to_string(cPortID, cPortMap, cPortName, (C.size_t)(OFP_MAX_PORT_NAME_LEN))
		defer C.free(unsafe.Pointer(cPortName))
	}

	return &OutputAction{
		PortID:   FlowPort(cPortID),
		PortName: C.GoString(cPortName),
	}
}

// endregion

// region OFPACT_OUTPUT_REG

// OutputRegAction represents an OFPACT_OUTPUT_REG action.
type OutputRegAction struct {
	ActionHeader
	Src *MetaFlowSubfield
}

func (action *OutputRegAction) String() string {
	return fmt.Sprintf("output:%s", action.Src)
}

func (action *OutputRegAction) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("action not implemented")
}

func newOutputRegAction(ptr *C.struct_ofpact_output_reg) *OutputRegAction {
	var cSrc C.struct_mf_subfield = C.get_ofpact_output_reg_src(ptr)

	subfield := newMetaFlowSubfield(&cSrc)

	return &OutputRegAction{
		Src: subfield,
	}
}

// endregion

// region OFPACT_SET_FIELD

// region OFPACT_LOAD

// LoadAction represents an NXAST_RAW_REG_LOAD action.

// bitwiseScan scans the mask to find the next set of bits (1s or 0s) in the given range.
func bitwiseScan(mask []byte, bitValue uint, start uint, end uint) uint {
	for i := start; i < end; i++ {
		byteIndex := i / 8
		bitIndex := i % 8
		bit := (mask[byteIndex] >> (7 - bitIndex)) & 1
		if bit == byte(bitValue) {
			return i
		}
	}
	return end
}

// bitwiseGet extracts the value of a field from the specified range of bits.
func bitwiseGet(value []byte, start uint, nBits uint) uint64 {
	var result uint64
	for i := uint(0); i < nBits; i++ {
		bitIndex := start + i
		byteIndex := bitIndex / 8
		bitPosition := bitIndex % 8
		bit := (value[byteIndex] >> (7 - bitPosition)) & 1
		result |= uint64(bit) << (nBits - 1 - i)
	}
	return result
}

// min is a helper function to find the minimum of two integers.
//func min[T int | uint](a, b T) T {
//	if a < b {
//		return a
//	}
//	return b
//}

type LoadAction struct {
	ActionHeader
	Values []*MetaFlowValue
	Fields []*MetaFlowSubfield
}

func (action *LoadAction) String() string {
	var elems []string
	for i := range min(len(action.Values), len(action.Fields)) {
		elems = append(elems, fmt.Sprintf("%s->%s", action.Values[i], action.Fields[i]))
	}

	return fmt.Sprintf("load:%s", strings.Join(elems, ","))
}

func (action *LoadAction) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("action not implemented")
}

func newLoadAction(ptr *C.struct_ofpact_set_field) *LoadAction {
	action := &LoadAction{
		Values: make([]*MetaFlowValue, 0),
		Fields: make([]*MetaFlowSubfield, 0),
	}

	var cValue C.uint64_t
	CSubfield := C.struct_mf_subfield{
		ofs:    0,
		n_bits: 0,
	}

	for C.next_load_segment(ptr, &CSubfield, &cValue) {
		action.Values = append(action.Values, newMetaFlowValue(cValue))
		action.Fields = append(action.Fields, newMetaFlowSubfield(&CSubfield))
	}

	return action
}

// endregion

// region OFPACT_SET_FIELD

// SetFieldAction represents an OFPACT_SET_FIELD action.
type SetFieldAction struct {
	FlowHasVlan bool
	Field       *MetaFlowField
	Value       *MetaFlowValue
}

func (a *SetFieldAction) String() string {
	return fmt.Sprintf("set_field:%s->%s", a.Value, a.Field)
}

func newSetFieldAction(ptr *C.struct_ofpact_set_field) *SetFieldAction {

	//var cFlowHasVlan C.bool_t = C.bool_t(false)
	//var cField *C.struct_mf_field = nil

	return &SetFieldAction{
		//FlowHasVlan: bool(cFlowHasVlan),
		//Field:       newMetaFlowField(cField),
		//Value:       "0x1",
	}
}

// endregion

// endregion

// region OFPACT_REG_MOVE

type RegMoveAction struct {
	ActionHeader
	Src *MetaFlowSubfield
	Dst *MetaFlowSubfield
}

func (action *RegMoveAction) String() string {
	return fmt.Sprintf("move(%s->%s)", action.Src, action.Dst)
}

func (action *RegMoveAction) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("action not implemented")
}

func newRegMoveAction(ptr *C.struct_ofpact_reg_move) *RegMoveAction {
	var cSrc C.struct_mf_subfield = C.get_ofpact_reg_move_src(ptr)
	var cDst C.struct_mf_subfield = C.get_ofpact_reg_move_dst(ptr)

	return &RegMoveAction{
		Src: newMetaFlowSubfield(&cSrc),
		Dst: newMetaFlowSubfield(&cDst),
	}
}

// endregion

// region OFPACT_RESUBMIT

// ActionResubmit represents a OFPACT_RESUBMIT action.
type ActionResubmit struct {
	ActionHeader

	vendorType uint16
	length     uint16
	vendor     uint32
	subtype    FlowActionType

	InPortID   FlowPort
	InPortName string

	TableID   uint8
	TableName string

	WithConntrackOrig bool

	padding [3]byte
}

func (action *ActionResubmit) Len() (n uint16) {
	n = action.Header().Len()
	n += 3
	n += uint16(len(action.padding))
	return n
}

func (action *ActionResubmit) String() string {
	if action.TableName != "" {
		return fmt.Sprintf("resubmit(,%s)", action.TableName)
	}
	return fmt.Sprintf("resubmit(,%d)", action.TableID)
}

func (action *ActionResubmit) MarshalBinary() (data []byte, err error) {
	data = make([]byte, action.Len())

	var bytes []byte
	bytes, err = action.ActionHeader.MarshalBinary()
	if err != nil {
		return nil, err
	}
	n := copy(data, bytes)

	data[n] = action.TableID
	n += 1

	binary.BigEndian.PutUint16(data[n:], uint16(action.InPortID))
	n += 2

	n += copy(data[n:], action.padding[:])

	return data, nil
}

func NewActionResubmit() *ActionResubmit {
	action := new(ActionResubmit)
	action.Type = 24
	action.Length = 0

	action.vendorType = C.OFPAT_VENDOR // 0xffff
	action.length = 16
	action.vendor = C.NX_VENDOR_ID //0xffff
	action.subtype = C.NXAST_RAW_RESUBMIT

	action.InPortID = 0 // OFPP_IN_PORT
	action.InPortName = ""

	action.TableID = OFP_DEFAULT_TABLE
	action.TableName = ""

	action.WithConntrackOrig = false

	return action
}

func newResubmitAction(ptr *C.struct_ofpact_resubmit, cTableMap *C.struct_ofputil_table_map) *ActionResubmit {
	cTableID := C.get_ofpact_resubmit_table_id(ptr)
	cTableName := C.CString("")
	{
		C.ofputil_table_to_string(cTableID, cTableMap, cTableName, (C.size_t)(OFP_MAX_TABLE_NAME_LEN))
		defer C.free(unsafe.Pointer(cTableName))
	}

	return &ActionResubmit{
		TableID:   uint8(cTableID),
		TableName: C.GoString(cTableName),
	}
}

// endregion

// region OFPACT_LEARN

type ActionLearnFlag byte

const (
	NX_LEARN_F_SEND_FLOW_REM  ActionLearnFlag = C.NX_LEARN_F_SEND_FLOW_REM
	NX_LEARN_F_DELETE_LEARNED ActionLearnFlag = C.NX_LEARN_F_DELETE_LEARNED
	NX_LEARN_F_WRITE_RESULT   ActionLearnFlag = C.NX_LEARN_F_WRITE_RESULT
)

type ActionLearnType uint16

const (
	// NX_LEARN_SRC_FIELD performs copy from field.
	NX_LEARN_SRC_FIELD ActionLearnType = C.NX_LEARN_SRC_FIELD

	// NX_LEARN_SRC_IMMEDIATE performs copy from immediate value.
	NX_LEARN_SRC_IMMEDIATE ActionLearnType = C.NX_LEARN_SRC_IMMEDIATE

	// NX_LEARN_DST_MATCH performs add match criterion.
	NX_LEARN_DST_MATCH ActionLearnType = C.NX_LEARN_DST_MATCH

	// NX_LEARN_DST_LOAD performs add NXAST_REG_LOAD action.
	NX_LEARN_DST_LOAD ActionLearnType = C.NX_LEARN_DST_LOAD

	// NX_LEARN_DST_OUTPUT performs add OFPACT_OUTPUT action.
	NX_LEARN_DST_OUTPUT ActionLearnType = C.NX_LEARN_DST_OUTPUT
)

// LearnSpecAction is a part of struct LearnAction, below.
type LearnSpecAction struct {
	// Src is NX_LEARN_SRC_FIELD only.
	Src *MetaFlowSubfield
	// Dst is NX_LEARN_DST_MATCH or NX_LEARN_DST_LOAD only.
	Dst *MetaFlowSubfield

	// SrcType is one of NX_LEARN_SRC_*.
	SrcType ActionLearnType
	// DstType is NX_LEARN_DST_*.
	DstType ActionLearnType
	// NumBits is a number of bits in a source and dest.
	NumBits uint32
}

func (spec *LearnSpecAction) String() string {
	switch {
	case spec.SrcType == NX_LEARN_SRC_IMMEDIATE && spec.DstType == NX_LEARN_DST_MATCH:
		// Immediate value loaded into a match field
		return fmt.Sprintf("%s=%s", spec.Dst, spec.Src)

	case spec.SrcType == NX_LEARN_SRC_FIELD && spec.DstType == NX_LEARN_DST_MATCH:
		// Field copied to match field
		if spec.Src.String() == spec.Dst.String() {
			// Same source and destination fields
			return fmt.Sprintf("%s", spec.Dst)
		}
		// Different source and destination fields
		return fmt.Sprintf("%s=%s", spec.Dst, spec.Src)

	case spec.SrcType == NX_LEARN_SRC_IMMEDIATE && spec.DstType == NX_LEARN_DST_LOAD:
		// Immediate value loaded into a field
		return fmt.Sprintf("load:%s->%s", spec.Src, spec.Dst)

	case spec.SrcType == NX_LEARN_SRC_FIELD && spec.DstType == NX_LEARN_DST_LOAD:
		// Field loaded into another field
		return fmt.Sprintf("load:%s->%s", spec.Src, spec.Dst)

	case spec.SrcType == NX_LEARN_SRC_FIELD && spec.DstType == NX_LEARN_DST_OUTPUT:
		// Field output to a port
		return fmt.Sprintf("output:%s", spec.Src)

	default:
		// Fallback case for unhandled combinations
		return fmt.Sprintf("%s,%s", spec.Src, spec.Dst)
	}
}

type LearnSpecsAction []*LearnSpecAction

func (lsa LearnSpecsAction) String() string {
	var specs []string
	for _, spec := range lsa {
		specs = append(specs, fmt.Sprintf("%s", spec))
	}
	return strings.Join(specs, ",")
}

func newLearnSpecsAction(cSpecs *C.struct_ofpact_learn_spec, cSpecsLen C.size_t) (specs LearnSpecsAction) {
	specs = make([]*LearnSpecAction, 0)

	for offset := C.size_t(0); offset < cSpecsLen; {
		cSpec := (*C.struct_ofpact_learn_spec)(unsafe.Pointer(uintptr(unsafe.Pointer(cSpecs)) + uintptr(offset)))

		var cSrc C.struct_mf_subfield = C.get_ofpact_learn_spec_src(cSpec)
		var cDst C.struct_mf_subfield = C.get_ofpact_learn_spec_dst(cSpec)

		var cSrcType C.uint16_t = C.get_ofpact_learn_spec_src_type(cSpec)
		var cDstType C.uint16_t = C.get_ofpact_learn_spec_dst_type(cSpec)
		var cNumBits C.uint32_t = C.get_ofpact_learn_spec_n_bits(cSpec)

		spec := &LearnSpecAction{
			Src:     newMetaFlowSubfield(&cSrc),
			Dst:     newMetaFlowSubfield(&cDst),
			SrcType: ActionLearnType(cSrcType),
			DstType: ActionLearnType(cDstType),
			NumBits: uint32(cNumBits),
		}

		specs = append(specs, spec)
		offset += C.sizeof_struct_ofpact_learn_spec
	}

	return specs
}

// LearnAction represents a OFPACT_LEARN action.
type LearnAction struct {
	ActionHeader

	// IdleTimeout is an idle time before discarding (seconds).
	IdleTimeout uint16
	// HardTimeout is a max time before discarding (seconds).
	HardTimeout uint16
	// Priority level of flow entry.
	Priority uint16
	// TableID & TableName to insert flow entry.
	TableID   uint8
	TableName string
	// Flags is ActionLearnFlag (NX_LEARN_F_*).
	Flags ActionLearnFlag
	// Cookie for new flow.
	Cookie uint64
	// FinIdleTimeout is an idle timeout after FIN, if nonzero.
	FinIdleTimeout uint16
	// FinHardTimeout is a hard timeout after FIN, if nonzero.
	FinHardTimeout uint16
	// Limit, if the number of flows on 'TableID' with 'Cookie' exceeds this,
	// the action will not add a new flow. 0 indicates unlimited.
	Limit uint32
	// ResultDst used only if 'flags' has NX_LEARN_F_WRITE_RESULT.
	// If the execution failed to install a new flow because 'limit'
	// is exceeded, ResultDst will be set to 0, otherwise to 1.
	ResultDst *MetaFlowSubfield

	Specs LearnSpecsAction
}

func (action *LearnAction) String() string {
	var elems []string

	if action.TableName != "" {
		elems = append(elems, fmt.Sprintf("table=%s", action.TableName))
	}

	if action.IdleTimeout != OFP_FLOW_PERMANENT {
		elems = append(elems, fmt.Sprintf("idle_timeout=%d", action.IdleTimeout))
	}

	if action.HardTimeout != OFP_FLOW_PERMANENT {
		elems = append(elems, fmt.Sprintf("hard_timeout=%d", action.HardTimeout))
	}

	if action.FinIdleTimeout != 0 {
		elems = append(elems, fmt.Sprintf("fin_idle_timeout=%d", action.FinIdleTimeout))
	}

	if action.FinHardTimeout != 0 {
		elems = append(elems, fmt.Sprintf("fin_hard_timeout=%d", action.FinHardTimeout))
	}

	if action.Priority != OFP_DEFAULT_PRIORITY {
		elems = append(elems, fmt.Sprintf("priority=%d", action.Priority))
	}

	if action.Flags&NX_LEARN_F_SEND_FLOW_REM != 0 {
		elems = append(elems, "send_flow_rem")
	}

	if action.Flags&NX_LEARN_F_DELETE_LEARNED != 0 {
		elems = append(elems, "delete_learned")
	}

	if action.Cookie != 0 {
		elems = append(elems, fmt.Sprintf("cookie=0x%x", action.Cookie))
	}

	if action.Limit != 0 {
		elems = append(elems, fmt.Sprintf("limit=%d", action.Limit))
	}

	if action.Flags&NX_LEARN_F_WRITE_RESULT != 0 {
		elems = append(elems, fmt.Sprintf("result_dst=%s", action.ResultDst))
	}

	if action.Specs != nil {
		elems = append(elems, fmt.Sprintf("%s", action.Specs))
	}

	return fmt.Sprintf("learn(%s)", strings.Join(elems, ","))
}

func (action *LearnAction) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("action not implemented")
}

func newLearnAction(ptr *C.struct_ofpact_learn, cOfpactLen C.size_t, cTableMap *C.struct_ofputil_table_map) *LearnAction {
	cIdleTimeout := C.get_ofpact_learn_idle_timeout(ptr)
	cHardTimeout := C.get_ofpact_learn_hard_timeout(ptr)
	cPriority := C.get_ofpact_learn_priority(ptr)
	cTableID := C.get_ofpact_learn_table_id(ptr)
	cTableName := C.CString("")
	{
		C.ofputil_table_to_string(cTableID, cTableMap, cTableName, (C.size_t)(OFP_MAX_TABLE_NAME_LEN))
		defer C.free(unsafe.Pointer(cTableName))
	}
	cFlags := C.get_ofpact_learn_flags(ptr)
	cCookie := C.get_ofpact_learn_cookie(ptr)
	cFinIdleTimeout := C.get_ofpact_learn_fin_idle_timeout(ptr)
	cFinHardTimeout := C.get_ofpact_learn_fin_hard_timeout(ptr)
	cLimit := C.get_ofpact_learn_limit(ptr)
	var cResultDst C.struct_mf_subfield = C.get_ofpact_learn_result_dst(ptr)
	var cSpecs *C.struct_ofpact_learn_spec = C.get_ofpact_learn_specs(ptr)
	var cSpecsLen C.size_t = cOfpactLen - C.sizeof_struct_ofpact_learn

	return &LearnAction{
		IdleTimeout:    uint16(cIdleTimeout),
		HardTimeout:    uint16(cHardTimeout),
		Priority:       uint16(cPriority),
		TableID:        uint8(cTableID),
		TableName:      C.GoString(cTableName),
		Flags:          ActionLearnFlag(cFlags),
		Cookie:         uint64(cCookie),
		FinIdleTimeout: uint16(cFinIdleTimeout),
		FinHardTimeout: uint16(cFinHardTimeout),
		Limit:          uint32(cLimit),
		ResultDst:      newMetaFlowSubfield(&cResultDst),
		Specs:          newLearnSpecsAction(cSpecs, cSpecsLen),
	}
}

// endregion

// region OFPACT_CONJUNCTION

// ConjunctionAction represents a OFPACT_CONJUNCTION action.
type ConjunctionAction struct {
	ActionHeader

	ID         uint32
	Clause     uint8
	NumClauses uint8
}

func (action *ConjunctionAction) String() string {
	return fmt.Sprintf("conjunction(%d,%d/%d)", action.ID, action.Clause+1, action.NumClauses)
}

func (action *ConjunctionAction) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("action not implemented")
}

func newConjunctionAction(ptr *C.struct_ofpact_conjunction) *ConjunctionAction {
	conjunction := (*C.struct_ofpact_conjunction)(unsafe.Pointer(ptr))

	cID := C.get_ofpact_conjunction_id(conjunction)
	cClause := C.get_ofpact_conjunction_clause(conjunction)
	cNumClauses := C.get_ofpact_conjunction_n_clauses(conjunction)

	return &ConjunctionAction{
		ID:         uint32(cID),
		Clause:     uint8(cClause),
		NumClauses: uint8(cNumClauses),
	}
}

// endregion

// region OFPACT_CT

type ConntrackFlag uint16

const (
	NX_CT_F_COMMIT ConntrackFlag = C.NX_CT_F_COMMIT
	NX_CT_F_FORCE  ConntrackFlag = C.NX_CT_F_FORCE
)

const NX_CT_RECIRC_NONE uint8 = C.NX_CT_RECIRC_NONE

// ConntrackAction represents an OFPACT_CT action.
type ConntrackAction struct {
	ActionHeader

	Flags           ConntrackFlag
	ZoneImm         uint16
	ZoneSrc         *MetaFlowSubfield
	Alg             uint16
	RecircTable     uint8
	RecircTableName string
	Actions         FlowActions
}

func (action *ConntrackAction) String() string {
	var elems []string

	if action.Flags&NX_CT_F_COMMIT != 0 {
		elems = append(elems, "commit")
	}

	if action.Flags&NX_CT_F_FORCE != 0 {
		elems = append(elems, "force")
	}

	if action.RecircTable != NX_CT_RECIRC_NONE {
		elems = append(elems, fmt.Sprintf("table=%s", action.RecircTableName))
	}

	if action.ZoneSrc.Field != nil {
		elems = append(elems, fmt.Sprintf("zone=%s", action.ZoneSrc))
	} else if action.ZoneImm != 0 {
		elems = append(elems, fmt.Sprintf("zone=%d", action.ZoneImm))
	}

	// /* If the first action is a NAT action, format it outside of the 'exec'
	//     * envelope. */
	//    const struct ofpact *action = a->actions;
	//    size_t actions_len = ofpact_ct_get_action_len(a);
	//    if (actions_len && action->type == OFPACT_NAT) {
	//        format_NAT(ofpact_get_NAT(action), fp);
	//        ds_put_char(fp->s, ',');
	//        actions_len -= OFPACT_ALIGN(action->len);
	//        action = ofpact_next(action);
	//    }

	//    if (actions_len) {
	//        ds_put_format(fp->s, "%sexec(%s", colors.paren, colors.end);
	//        ofpacts_format(action, actions_len, fp);
	//        ds_put_format(fp->s, "%s),%s", colors.paren, colors.end);
	//    }
	if action.Actions != nil {
		//var subelems []string
		//for _, subaction := range action.Actions {
		//	subelems = append(subelems, subaction.String())
		//}

		//elems = append(elems, fmt.Sprintf("exec(%s)", strings.Join(subelems, ",")))
		elems = append(elems, fmt.Sprintf("exec(%s)", action.Actions))
	}

	switch action.Alg {
	case 0:
	case C.IPPORT_FTP:
		elems = append(elems, "alg=ftp")
	case C.IPPORT_TFTP:
		elems = append(elems, "alg=tftp")
	default:
		elems = append(elems, fmt.Sprintf("alg=%d", action.Alg))
	}

	return fmt.Sprintf("ct(%s)", strings.Join(elems, ","))
}

func (action *ConntrackAction) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("action not implemented")
}

func newConntrackAction(ptr *C.struct_ofpact_conntrack, cPortMap *C.struct_ofputil_port_map, cTableMap *C.struct_ofputil_table_map) *ConntrackAction {
	var cFlags C.uint16_t = C.get_ofpact_conntrack_flags(ptr)
	var cZoneImm C.uint16_t = C.get_ofpact_conntrack_zone_imm(ptr)
	var cZoneSrc C.struct_mf_subfield = C.get_ofpact_conntrack_zone_src(ptr)
	var cAlg C.uint16_t = C.get_ofpact_conntrack_alg(ptr)
	var cRecircTable C.uint8_t = C.get_ofpact_conntrack_recirc_table(ptr)
	cRecircTableName := C.CString("")
	{
		C.ofputil_table_to_string(cRecircTable, cTableMap, cRecircTableName, (C.size_t)(OFP_MAX_TABLE_NAME_LEN))
		defer C.free(unsafe.Pointer(cRecircTableName))
	}
	var cConntrackActions *C.struct_ofpact = C.get_ofpact_conntrack_actions(ptr)
	var cConntrackActionsLen C.size_t = C.ofpact_ct_get_action_len(ptr)

	var actions FlowActions
	if cConntrackActionsLen != 0 {
		actions = newFlowActions(cConntrackActions, cConntrackActionsLen, cPortMap, cTableMap)
	}

	return &ConntrackAction{
		Flags:           ConntrackFlag(cFlags),
		ZoneImm:         uint16(cZoneImm),
		ZoneSrc:         newMetaFlowSubfield(&cZoneSrc),
		Alg:             uint16(cAlg),
		RecircTable:     uint8(cRecircTable),
		RecircTableName: C.GoString(cRecircTableName),
		Actions:         actions,
	}
}

// endregion

// region OFPACT_UNIMPL

// UnimplementedAction is used for actions that are not explicitly handled.
type UnimplementedAction struct {
	ActionHeader
}

func (action *UnimplementedAction) String() string {
	return fmt.Sprintf("<unimplemented(%d)>", action.Type)
}

func (action *UnimplementedAction) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("action not implemented")
}

// endregion

// region OFPACT_DROP

// DropAction represents a drop action, which is applied when no other actions are specified.
type DropAction struct {
	ActionHeader
}

func (action *DropAction) String() string {
	return "drop"
}

func (action *DropAction) MarshalBinary() (data []byte, err error) {
	return
}

// endregion

// FlowAction is an interface that all actions should implement.
type FlowAction interface {
	Message
	String() string
}

func newFlowAction(ptr *C.struct_ofpact, cPortMap *C.struct_ofputil_port_map, cTableMap *C.struct_ofputil_table_map) FlowAction {
	action := FlowActionType(ptr._type)

	switch action {
	case OFPACT_OUTPUT:
		return newOutputAction((*C.struct_ofpact_output)(unsafe.Pointer(ptr)), cPortMap)
	case OFPACT_OUTPUT_REG:
		return newOutputRegAction((*C.struct_ofpact_output_reg)(unsafe.Pointer(ptr)))
	case OFPACT_SET_FIELD:
		if FlowRawActionType(ptr.raw) == NXAST_RAW_REG_LOAD {
			return newLoadAction((*C.struct_ofpact_set_field)(unsafe.Pointer(ptr)))
		} else {
			//return newSetFieldAction((*C.struct_ofpact_set_field)(unsafe.Pointer(ptr)))
		}
	case OFPACT_REG_MOVE:
		return newRegMoveAction((*C.struct_ofpact_reg_move)(unsafe.Pointer(ptr)))
	case OFPACT_RESUBMIT:
		return newResubmitAction((*C.struct_ofpact_resubmit)(unsafe.Pointer(ptr)), cTableMap)
	case OFPACT_LEARN:
		return newLearnAction((*C.struct_ofpact_learn)(unsafe.Pointer(ptr)), C.size_t(ptr.len), cTableMap)
	case OFPACT_CONJUNCTION:
		return newConjunctionAction((*C.struct_ofpact_conjunction)(unsafe.Pointer(ptr)))
	case OFPACT_CT:
		return newConntrackAction((*C.struct_ofpact_conntrack)(unsafe.Pointer(ptr)), cPortMap, cTableMap)
	}

	return &UnimplementedAction{
		ActionHeader: ActionHeader{
			Type: action,
		},
	}
}

// FlowActions represents a list of actions to be applied to matching packets.
type FlowActions []FlowAction

func (fa FlowActions) Len() (n uint16) {
	for _, action := range fa {
		n += action.Len()
	}
	return n
}

func (fa FlowActions) String() string {
	var actions []string
	for _, action := range fa {
		actions = append(actions, fmt.Sprintf("%s", action))
	}
	return strings.Join(actions, ",")
}

func newFlowActions(
	cPtr *C.struct_ofpact, cLen C.size_t, cPortMap *C.struct_ofputil_port_map, cTableMap *C.struct_ofputil_table_map,
) (actions FlowActions) {
	actions = make([]FlowAction, 0)

	if cLen == 0 {
		actions = append(actions, &DropAction{})
		return actions
	}

	for offset := C.size_t(0); offset < cLen; {
		ofpact := (*C.struct_ofpact)(unsafe.Pointer(uintptr(unsafe.Pointer(cPtr)) + uintptr(offset)))

		actions = append(actions,
			newFlowAction(ofpact, cPortMap, cTableMap),
		)

		offset += C.size_t(ofpact.len)
	}

	return actions
}
