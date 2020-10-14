package utils

import (
	"fmt"
	"strings"
)

func ImagePullAndTag (imageSlice []string) ([]string, error) {
	var failedSlice []string
	for _, value := range imageSlice {
		pullAndTagCmd := fmt.Sprintf("docker pull %s && docker tag %s k8s.gcr.io/%s", value, value, strings.Split(value, "/")[len(strings.Split(value, "/"))-1])
		if _, err := RunCommand(pullAndTagCmd); err != nil {
			failedSlice = append(failedSlice, value)
		}
	}

	if len(failedSlice) != 0 {
		return failedSlice, fmt.Errorf("failed")
	}

	return nil, nil
}