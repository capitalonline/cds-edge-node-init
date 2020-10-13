package pkg

import "testing"

func TestNetworkConfig(t *testing.T) {
	if err:= NetworkConfig(); err != nil {
		t.Errorf("NetworkConfig testing error, err is: %s", err)
	}
}
