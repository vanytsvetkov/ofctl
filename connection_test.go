package ofctl_test

import (
	"ofctl"
	"testing"
)

func newConnection(t *testing.T) {
	var err error

	connection := ofctl.NewConnection()

	err = connection.Open(Bridge)
	if err != nil {
		t.Fatalf("Error connecting to OVS: %v", err)
	}

	t.Log("Connected to OVS")
}

func TestOpenConnection(t *testing.T) {
	defer ovs.Close()
	t.Run("NewConnection", newConnection)
}
