package pkg

import "testing"

func TestDockerInstall(t *testing.T) {
	if err:= DockerInstall(); err != nil {
		t.Errorf("DockerInstall testing error, err is: %s", err)
	}
}
