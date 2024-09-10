package ofctl_test

import (
	"ofctl"
)

var ovs = new(ofctl.Connection)

const Bridge = "br0"

// deprecated
//func TestRun(t *testing.T) {
//	stdout, err := ofctl.Run("--help")
//	if err != nil {
//		t.Fatalf("Error running dump-flows: %v", err)
//	}
//
//	t.Log(stdout)
//}
