package pkg

import "testing"

func TestPythonInstall(t *testing.T) {
	if err:= PythonInstall(); err != nil {
		t.Errorf("PythonInstall testing error, err is: %s", err)
	}
}
