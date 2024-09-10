package ofctl

/*
#include "include/ofctl.h"
*/
import "C"
import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"unsafe"
)

// FlowStat represents flow statistic.
type FlowStat struct {
	Cookie      uint64
	Duration    float64
	Table       uint8
	TableName   string
	Priority    uint16
	IdleTimeout uint16
	HardTimeout uint16
	PacketCount uint64
	ByteCount   uint64
	Match       *FlowMatch
	Actions     FlowActions
}

func (fs *FlowStat) weight() uint {
	return uint(fs.Priority) + 1<<fs.Table
}

// newFlowStats initializes a FlowStats structure from the C struct.
func newFlowStat(cStats *C.struct_ofputil_flow_stats, cPortMap *C.struct_ofputil_port_map, cTableMap *C.struct_ofputil_table_map) *FlowStat {
	var cTableID C.uint8_t = cStats.table_id
	cTableName := C.CString("")
	{
		C.ofputil_table_to_string(cTableID, cTableMap, cTableName, (C.size_t)(OFP_MAX_TABLE_NAME_LEN))
		defer C.free(unsafe.Pointer(cTableName))
	}

	return &FlowStat{
		Cookie:      uint64(cStats.cookie),
		Duration:    float64(cStats.duration_sec) + float64(cStats.duration_nsec)*1e-9,
		Table:       uint8(cTableID),
		TableName:   C.GoString(cTableName),
		Priority:    uint16(cStats.priority),
		IdleTimeout: uint16(cStats.idle_timeout),
		HardTimeout: uint16(cStats.hard_timeout),
		PacketCount: uint64(cStats.packet_count),
		ByteCount:   uint64(cStats.byte_count),
		Match:       newFlowMatch(&cStats.match),
		Actions:     newFlowActions(cStats.ofpacts, cStats.ofpacts_len, cPortMap, cTableMap),
	}
}

// String returns a string representation of the flow statistics.
func (fs *FlowStat) String() string {
	var stats string
	{
		var elems []string

		elems = append(elems, fmt.Sprintf("cookie=0x%x", fs.Cookie))
		elems = append(elems, fmt.Sprintf("duration=%.3fs", fs.Duration))

		if fs.TableName != "" {
			elems = append(elems, fmt.Sprintf("table=%s", fs.TableName))
		}

		elems = append(elems, fmt.Sprintf("n_packets=%d", fs.PacketCount))
		elems = append(elems, fmt.Sprintf("n_bytes=%d", fs.ByteCount))

		if fs.IdleTimeout != OFP_FLOW_PERMANENT {
			elems = append(elems, fmt.Sprintf("idle_timeout=%d", fs.IdleTimeout))
		}

		if fs.HardTimeout != OFP_FLOW_PERMANENT {
			elems = append(elems, fmt.Sprintf("hard_timeout=%d", fs.HardTimeout))
		}

		// todo: do we need any flags?

		// todo: do we need importance=?

		// todo: do we need idle_age=?
		// todo: do we need hard_age=?

		elems = append(elems, "")

		stats = strings.Join(elems, ", ")
	}

	var match string
	{
		var elems []string

		if fs.Priority != OFP_DEFAULT_PRIORITY {
			elems = append(elems, fmt.Sprintf("priority=%d", fs.Priority))
		}

		if fs.Match != nil {
			if str := fs.Match.String(); str != "" {
				elems = append(elems, str)
			}
		}

		match = strings.Join(elems, ",")
	}

	return fmt.Sprintf("%s%s actions=%s", stats, match, fs.Actions)
}

// FlowStats represents flow statistics.
type FlowStats []*FlowStat

func (fss *FlowStats) Sort(cmp func(a, b *FlowStat) int) {
	slices.SortFunc(*fss, cmp)
}

func SortFlowStatsByWeight(a, b *FlowStat) int {
	return cmp.Compare(a.weight(), b.weight())
}
