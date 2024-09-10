package ofctl

//#include "include/ofctl.h"
import "C"
import (
	"antrea.io/libOpenflow/protocol"
	"antrea.io/libOpenflow/util"
	"encoding/binary"
	"fmt"
	"net"
	"sync/atomic"
	"unsafe"

	"antrea.io/libOpenflow/openflow13"
)

type Message interface {
	MarshalBinary() (data []byte, err error)
	//UnmarshalBinary(data []byte) error

	Len() uint16
}

var messageXid uint32 = 1

func headerGenerator() func() Header {
	return func() Header {
		xid := atomic.AddUint32(&messageXid, 1)
		p := Header{version, 0, 8, xid}
		return p
	}
}

var newOfpHeader func() Header = headerGenerator()

// The version specifies the OpenFlow protocol version being
// used. During the current draft phase of the OpenFlow
// Protocol, the most significant bit will be set to indicate an
// experimental version and the lower bits will indicate a
// revision number. The current version is 0x01. The final
// version for a Type 0 switch will be 0x00. The length field
// indicates the total length of the message, so no additional
// framing is used to distinguish one frame from the next.

type Header struct {
	Version OpenFlowVersion
	Type    uint8
	Length  uint16
	Xid     uint32
}

func (h *Header) Header() *Header {
	return h
}

func (h *Header) Len() (n uint16) {
	return 8
}

func (h *Header) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 8)
	data[0] = uint8(h.Version)
	data[1] = h.Type
	binary.BigEndian.PutUint16(data[2:4], h.Length)
	binary.BigEndian.PutUint32(data[4:8], h.Xid)
	return
}

func (h *Header) UnmarshalBinary(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("the []byte is too short to unmarshel a full Header")
	}
	h.Version = OpenFlowVersion(data[0])
	h.Type = data[1]
	h.Length = binary.BigEndian.Uint16(data[2:4])
	h.Xid = binary.BigEndian.Uint32(data[4:8])
	return nil
}

type InstructionType uint16

const (
	OVSINST_OFPIT11_APPLY_ACTIONS InstructionType = C.OVSINST_OFPIT11_APPLY_ACTIONS
	OVSINST_OFPIT11_WRITE_ACTIONS InstructionType = C.OVSINST_OFPIT11_WRITE_ACTIONS
)

type Instruction interface {
	Message
}

type InstructionHeader struct {
	Type   InstructionType
	Length uint16
}

func (h *InstructionHeader) Header() *InstructionHeader {
	return h
}

func (h *InstructionHeader) Len() (n uint16) {
	return 4
}

func (h *InstructionHeader) MarshalBinary() (data []byte, err error) {
	data = make([]byte, h.Len())
	binary.BigEndian.PutUint16(data[:2], uint16(h.Type))
	binary.BigEndian.PutUint16(data[2:4], h.Length)
	return
}

type InstructionActions struct {
	InstructionHeader
	padding [4]byte
	Actions FlowActions /* 0 or more actions associated with OFPIT_WRITE_ACTIONS and OFPIT_APPLY_ACTIONS */
}

func (instruction *InstructionActions) Len() (n uint16) {
	n = instruction.Header().Len()
	n += uint16(len(instruction.padding))
	n += instruction.Actions.Len()
	return n
}

func (instruction *InstructionActions) MarshalBinary() (data []byte, err error) {
	data = make([]byte, instruction.Len())

	var bytes []byte
	{
		bytes, err = instruction.Header().MarshalBinary()
		if err != nil {
			return nil, err
		}
	}
	n := copy(data, bytes)
	n += copy(data[n:], instruction.padding[:]) // padding

	for _, action := range instruction.Actions {
		bytes, err = action.MarshalBinary()
		if err != nil {
			return nil, err
		}
		n += copy(data[n:], bytes)
	}

	return
}

func NewInstructionWriteActions() *InstructionActions {
	instr := new(InstructionActions)
	instr.Type = OVSINST_OFPIT11_WRITE_ACTIONS
	instr.Actions = make(FlowActions, 0)
	instr.Length = instr.Header().Len()

	return instr
}

func NewInstructionApplyActions() *InstructionActions {
	instr := new(InstructionActions)
	instr.Type = OVSINST_OFPIT11_APPLY_ACTIONS
	instr.Actions = make(FlowActions, 0)
	instr.Length = instr.Header().Len()

	return instr
}

type Instructions []Instruction

func (instructions *Instructions) Len() (n uint16) {
	for _, i := range *instructions {
		n += i.Len()
	}
	return n
}

// FlowMod (ofp_flow_mod v1.3)
type FlowMod struct {
	Header

	Cookie     uint64
	CookieMask uint64

	TableID uint8 /* ID of the table to put the flow in */
	Command uint8 /* flowmod command */

	IdleTimeout uint16 /* Idle time before discarding (seconds). */
	HardTimeout uint16 /* Max time before discarding (seconds). */

	Priority uint16 /* Priority level of flow entry. */
	BufferID uint32 /* Buffered packet to apply to */

	OutPort  FlowPort
	OutGroup uint32
	Flags    uint16

	padding [2]byte

	Match        *Match       // Fields to match
	Instructions Instructions //  Instructions set - zero or more.
}

func (msg *FlowMod) Len() (n uint16) {
	n = msg.Header.Len()
	n += 38
	n += uint16(len(msg.padding))
	if msg.Match != nil {
		//n += msg.Match.Len()
	}
	n += msg.Instructions.Len()

	return roundUp(n, 8)
}

// newFlowMod converts a Match object into an OpenFlow message.
func newFlowMod() *FlowMod {
	msg := new(FlowMod)

	msg.Header = newOfpHeader()
	msg.Header.Type = 14 // C.OFPTYPE_FLOW_MOD (for controller) // todo: откуда 14?

	msg.Cookie = 0
	msg.CookieMask = 0

	msg.TableID = OFP_DEFAULT_TABLE
	msg.Command = C.OFPFC_ADD

	msg.IdleTimeout = OFP_FLOW_PERMANENT
	msg.HardTimeout = OFP_FLOW_PERMANENT

	msg.Priority = OFP_DEFAULT_PRIORITY
	msg.BufferID = 0xffffffff
	msg.OutPort = OFPP_NONE
	msg.OutGroup = 0xffffffff
	msg.Flags = 0

	msg.Match = nil
	msg.Instructions = []Instruction{NewInstructionApplyActions()}

	return msg
}

func (msg *FlowMod) SetActions(actions ...FlowAction) {
	// Именно так, поскольку на текущий момент у нас нет других типов действий, кроме apply.
	apply := NewInstructionApplyActions()
	apply.Actions = actions

	msg.Instructions = []Instruction{apply}
}

func (msg *FlowMod) MarshalBinary() (data []byte, err error) {
	msg.Header.Length = msg.Len()
	data = make([]byte, msg.Header.Length)

	var bytes []byte
	{
		bytes, err = msg.Header.MarshalBinary()
		if err != nil {
			return nil, err
		}
	}
	n := copy(data, bytes)

	binary.BigEndian.PutUint64(data[n:], msg.Cookie)
	n += 8

	binary.BigEndian.PutUint64(data[n:], msg.CookieMask)
	n += 8

	data[n] = msg.TableID
	n += 1

	data[n] = msg.Command
	n += 1

	binary.BigEndian.PutUint16(data[n:], msg.IdleTimeout)
	n += 2

	binary.BigEndian.PutUint16(data[n:], msg.HardTimeout)
	n += 2

	binary.BigEndian.PutUint16(data[n:], msg.Priority)
	n += 2

	binary.BigEndian.PutUint32(data[n:], msg.BufferID)
	n += 4

	binary.BigEndian.PutUint32(data[n:], msg.OutPort.ofp11())
	n += 4

	binary.BigEndian.PutUint32(data[n:], msg.OutGroup)
	n += 4

	binary.BigEndian.PutUint16(data[n:], msg.Flags)
	n += 2

	n += copy(data[n:], msg.padding[:])

	if msg.Match != nil {
		// todo: do it!
		//bytes, err = msg.Match.MarshalBinary()
		//data = append(data, bytes...)
	}

	for _, instruction := range msg.Instructions {
		bytes, err = instruction.MarshalBinary()
		if err != nil {
			return nil, err
		}
		n += copy(data[n:], bytes)
	}

	return
}

func ParseFlow(bridge string, flow string) (err error) {
	connection := NewConnection()
	if err = connection.Open(bridge); err != nil {
		return err
	}
	defer connection.Close()

	//flow := fmt.Sprintf("table=%d, priority=%d actions=%s", table, priority, actions)
	cFlow := C.CString(flow)
	defer C.free(unsafe.Pointer(cFlow))

	// Prepare flow mod request

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

	var cRequest C.struct_ofputil_flow_mod
	C.memset(unsafe.Pointer(&cRequest), 0, C.sizeof_struct_ofputil_flow_mod)
	{
		var cUsableProtocols C.enum_ofputil_protocol
		errMessage := C.parse_ofp_flow_mod_str(
			&cRequest, cFlow, cPortMap, cTableMap, C.OFPFC_ADD, &cUsableProtocols,
		)
		if errMessage != nil {
			return fmt.Errorf("failed to parse flow mod request: %s", C.GoString(errMessage))
		}

		// Set protocol for mod flows
		usableProtocols := OpenFlowProtocol(cUsableProtocols)
		allowedProtocols := OFP_ANY
		for _, ofp := range []OpenFlowProtocol{
			OFP15_OXM,
			OFP14_OXM,
			OFP13_OXM,
			OFP12_OXM,
			OFP11_STD,
			OFP10_NXM_TID,
			OFP10_NXM,
			OFP10_STD_TID,
			OFP10_STD,
		} {
			if ofp&usableProtocols&allowedProtocols != 0 {
				err = connection.SetProtocol(ofp)
				if err == nil {
					break
				}
			}
		}
	}

	var cOfpbuf *C.struct_ofpbuf = C.ofputil_encode_flow_mod(&cRequest, connection.protocol)

	var ds []byte = C.GoBytes(unsafe.Pointer(cOfpbuf.data), C.int(cOfpbuf.size))

	fmt.Printf("%d\n", len(ds))
	fmt.Printf("%v\n", ds)
	fmt.Printf("0x%x\n", ds)

	//err = connection.transact(cOfpbuf)
	//err = connection.Send(cOfpbuf)
	if err != nil {
		return err
	}

	return nil
}

func addFlow(bridge string, flow string) (err error) {
	connection := NewConnection()
	if err = connection.Open(bridge); err != nil {
		return err
	}
	defer connection.Close()

	//flow := fmt.Sprintf("table=%d, priority=%d actions=%s", table, priority, actions)

	cFlow := C.CString(flow)
	defer C.free(unsafe.Pointer(cFlow))

	// Prepare flow mod request

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

	var cRequest C.struct_ofputil_flow_mod
	C.memset(unsafe.Pointer(&cRequest), 0, C.sizeof_struct_ofputil_flow_mod)
	{
		var cUsableProtocols C.enum_ofputil_protocol
		errMessage := C.parse_ofp_flow_mod_str(
			&cRequest, cFlow, cPortMap, cTableMap, C.OFPFC_ADD, &cUsableProtocols,
		)
		if errMessage != nil {
			return fmt.Errorf("failed to parse flow mod request: %s", C.GoString(errMessage))
		}

		// Set protocol for mod flows
		usableProtocols := OpenFlowProtocol(cUsableProtocols)
		allowedProtocols := OFP_ANY
		for _, ofp := range []OpenFlowProtocol{
			OFP15_OXM,
			OFP14_OXM,
			OFP13_OXM,
			OFP12_OXM,
			OFP11_STD,
			OFP10_NXM_TID,
			OFP10_NXM,
			OFP10_STD_TID,
			OFP10_STD,
		} {
			if ofp&usableProtocols&allowedProtocols != 0 {
				err = connection.SetProtocol(ofp)
				if err == nil {
					break
				}
			}
		}
	}

	var cOfpbuf *C.struct_ofpbuf = C.ofputil_encode_flow_mod(&cRequest, connection.protocol)

	var ds []byte = C.GoBytes(unsafe.Pointer(cOfpbuf.data), C.int(cOfpbuf.size))

	fmt.Printf("%v\n", ds)
	fmt.Printf("0x%x\n", ds)

	err = connection.transact(cOfpbuf)
	//err = connection.Send(cOfpbuf)
	if err != nil {
		return err
	}

	return nil
}

func AddFlow(bridge string, flow *Flow) (err error) {
	//msg := newFlowMod()
	//{
	//	msg.Command = C.OFPFC_ADD
	//	msg.TableID = flow.TableID
	//	msg.Priority = flow.Priority
	//	msg.SetActions(flow.Actions...)
	//}

	msg := openflow13.NewFlowMod()
	{
		msg.TableId = flow.TableID
		msg.Priority = flow.Priority

		//str := "table=22, priority=12307,tcp,nw_dst=10.100.2.238,tp_dst=39537 actions=drop"

		msg.Match.AddField(*openflow13.NewEthTypeField(protocol.IPv4_MSG))
		msg.Match.AddField(*openflow13.NewIpProtoField(protocol.Type_TCP))
		msg.Match.AddField(*openflow13.NewIpv4DstField(net.IPv4(10, 100, 2, 238), nil))
		msg.Match.AddField(*openflow13.NewTcpDstField(39537))

		instruction := openflow13.NewInstrApplyActions()

		for _, flowAction := range flow.Actions {
			switch fa := flowAction.(type) {
			case *ActionResubmit:
				action := openflow13.NewNXActionResubmitTableAction(uint16(OFPP_NONE), fa.TableID)
				err = instruction.AddAction(action, false)
				if err != nil {
					return err
				}
			}
		}

		if len(instruction.Actions) > 0 {
			msg.AddInstruction(instruction)
		}
	}

	err = ModFlow(bridge, msg)
	if err != nil {
		return err
	}

	return nil
}

func ModFlow(bridge string, msg util.Message) (err error) {
	connection := NewConnection()
	if err = connection.Open(bridge); err != nil {
		return err
	}
	defer connection.Close()

	err = connection.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
