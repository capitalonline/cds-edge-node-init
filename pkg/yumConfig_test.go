package pkg

import "testing"

func TestYumConfig(t *testing.T) {
	if err:= YumConfig(); err != nil {
		t.Errorf("YumConfig testing error, err is: %s", err)
	}
}
