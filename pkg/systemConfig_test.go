package pkg

import "testing"

func TestSystemConfig(t *testing.T) {
	if err:= SystemConfig(); err != nil {
		t.Errorf("SystemConfig testing error, err is: %s", err)
	}
}
