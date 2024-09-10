package ofctl_test

import (
	"fmt"
	"ofctl"
	"testing"
	"time"
)

func TestDumpFlows(t *testing.T) {
	t0 := time.Now()

	flows, err := ofctl.DumpFlows(Bridge)
	if err != nil {
		t.Fatalf("Error dumping flows: %v", err)
	}

	dt := time.Since(t0)

	_ = dt

	flows.Sort(ofctl.SortFlowStatsByWeight)

	t.Log("Match rules:")
	for _, flow := range flows {
		fmt.Printf("%s\n", flow)

		if len(flow.Actions) > 1 {
			continue
		}
	}

	dt2 := time.Since(t0)

	_ = dt2

	return
}

type IFlow interface {
	Get()
	Add()
	Delete()
}
