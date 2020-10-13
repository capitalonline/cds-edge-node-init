package pkg

import "testing"

func TestSetUp(t *testing.T) {
	if err:= SetUp(); err != nil {
		t.Errorf("SetUp testing error, err is: %s", err)
	}
}
