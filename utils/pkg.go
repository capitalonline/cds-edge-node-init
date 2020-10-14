package utils

import (
	"fmt"
)

func InstallPkgs(pkgs []string, group bool) ([]string, error) {
	// Logger.Printf("installPkgs: %s starting", pkgs)
	yumCmd := "yum install -y"
	if group {
		yumCmd = "yum -y groupinstall"
	}

	var failedSlice []string
	for _, value := range pkgs {
		installCmd := fmt.Sprintf("%s %s", yumCmd, value)
		if _, err := RunCommand(installCmd); err != nil {
			// Logger.Printf("install: %s error, err is: %s", value, err.Error())
			failedSlice = append(failedSlice, value)
		}
	}

	if len(failedSlice) != 0 {
		// Logger.Printf("some pkgs installed failed, return in failedSlice")
		return failedSlice, fmt.Errorf("failed")
	}

	// Logger.Printf("installPkgs: Succeed!")
	return nil, nil
}