package utilities

import (
	"testing"
)

func TestListenPort(t *testing.T) {
	listener, port, err := ListenPort()
	if err != nil {
		t.Fatalf("ListenPort() failed: %v", err)
	}
	defer listener.Close()

	// Verify that the port is within the valid range (1-65535)
	if port <= 0 || port > 65535 {
		t.Errorf("ListenPort() returned invalid port: %d", port)
	}
}
