package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
	"strings"
)

func GatewayConfig(gateway, privateIp string) error {
	log.Infof("GatewayConfig: starting")

	// find the net device by ip
	var netDevice string
	getDeviceCmd := fmt.Sprintf("ip a | grep %s", privateIp)
	if out, err := utils.RunCommand(getDeviceCmd); err != nil {
		return err
	} else if out != "" {
		netDevice = strings.Split(out, " ")[len(strings.Split(out, " "))-1]
	} else {
		return fmt.Errorf("GatewayConfig: not found net device by privateIp %s", privateIp)
	}

	// config gateway to privateIP
	netDeviceCfgFile := fmt.Sprintf("/etc/sysconfig/network-scripts/ifcfg-%s", netDevice)
	if !utils.FileExisted(netDeviceCfgFile) {
		return fmt.Errorf("GatewayConfig: %s not exist", netDeviceCfgFile)
	}

	gatewayCmd := fmt.Sprintf("echo GATEWAY=%s >> %s", gateway, netDeviceCfgFile)
	if _, err := utils.RunCommand(gatewayCmd); err != nil {
		return err
	}

	restartNetCmd := fmt.Sprintf("systemctl restart network")
	if _, err := utils.RunCommand(restartNetCmd); err != nil {
		return err
	}

	log.Infof("GatewayConfig: succeed!")
	return nil
}
