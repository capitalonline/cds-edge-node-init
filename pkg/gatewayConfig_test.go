package pkg

import "testing"

func TestGatewayConfig(t *testing.T) {
	// gateway, privateIp string
	gateway := ""
	privateIp := ""

	if err:= GatewayConfig(gateway, privateIp); err != nil {
		t.Errorf("GatewayConfig testing error, err is: %s", err)
	}

}
