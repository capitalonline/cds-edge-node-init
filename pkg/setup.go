package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
)

func SetUp (setUpSlice []string) error {
	for _, value := range setUpSlice {
		setUpCmd := fmt.Sprintf(" systemctl enable %s", value)
		if _, err := utils.RunCommand(setUpCmd); err != nil {
			return err
		}
	}

	return nil
}