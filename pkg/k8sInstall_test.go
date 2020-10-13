package pkg

import "testing"

func TestK8sInstall(t *testing.T) {
	if err:= K8sInstall(); err != nil {
		t.Errorf("K8sInstall testing error, err is: %s", err)
	}
}
