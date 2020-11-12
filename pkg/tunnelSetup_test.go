package pkg

import (
	"github.com/capitalonline/cds-edge-node-init/utils"
	"testing"
)

func TestTunnelSetup (t *testing.T) {
	// initData *utils.InitData
	var testingInitData utils.InitData
	testingInitData.PrivateIP = ""
	testingInitData.RootPassword = ""
	testingInitData.ClusterID = ""
	testingInitData.CustomerID = ""
	testingInitData.UserID = ""
	testingInitData.Gateway = ""
	testingInitData.Ak = ""
	testingInitData.Sk = ""
	testingInitData.K8sVersion = ""

	if err:= TunnelSetup(&testingInitData); err != nil {
		t.Errorf("TunnelSetup testing error, err is: %s", err)
	}

}
