package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
)


func GatewayConfig (gateway, device string) error {
	log.Infof("GatewayConfig: starting")

	netCmd := fmt.Sprintf("route add default gw %s dev %s", gateway, device)
	if _,err := utils.RunCommand(netCmd); err != nil {
		return err
	}

	reloadCmd := fmt.Sprintf("service network reload")
	if _,err := utils.RunCommand(reloadCmd); err != nil {
		return err
	}

	return nil
}
