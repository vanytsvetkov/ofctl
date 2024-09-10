package ofctl_test

import (
	"ofctl"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestAddFlows(t *testing.T) {
	var err error

	t0 := time.Now()

	//str := "table=22, priority=101 actions=drop"
	str := "table=22, priority=12307,tcp,nw_dst=10.100.2.238,tp_dst=39537 actions=drop" // resubmit(,23)" //
	err = ofctl.ParseFlow(Bridge, str)
	if err != nil {
		t.Fatalf(err.Error())
	}

	flow := ofctl.NewFlow()
	{
		flow.TableID = 22
		flow.Priority = 12307
		flow.Match = nil

		resubmit := ofctl.NewActionResubmit()
		{
			resubmit.TableID = 23
		}

		flow.Actions = []ofctl.FlowAction{
			//resubmit,
			//&ofctl.DropAction{},
		}
	}
	err = ofctl.AddFlow(Bridge, flow)
	if err != nil {
		t.Fatalf("Error adding flows: %v", err)
	}

	dt := time.Since(t0)

	_ = dt

	return
}
