package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"os/exec"
)

// RunCommand runs a given shell command
func RunCommand(cmd string) (string, error) {
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to run cmd: " + cmd + ", with out: " + string(out) + ", with error: " + err.Error())
	}
	return string(out), nil
}

// CreateDir create the target directory with error handling
func CreateDir(target string, mode int) error {
	fi, err := os.Lstat(target)

	if os.IsNotExist(err) {
		if err := os.MkdirAll(target, os.FileMode(mode)); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if fi != nil && !fi.IsDir() {
		return fmt.Errorf("%s already exist but it's not a directory", target)
	}
	return nil
}

// FileExisted checks if a file  or directory exists
func FileExisted(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func MarshalJsonToIOReader(v interface{}) (io.Reader, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	logrus.Infof("data is: %s", data)
	return bytes.NewBuffer(data), nil
}