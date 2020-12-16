package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
	"strings"
)

func GatewayConfig(gateway, privateIp string) error {
	log.Infof("GatewayConfig: starting")

	// check if public ip exist or not
	publicCmd := fmt.Sprintf("curl members.3322.org/dyndns/getip")
	if out, err := utils.RunCommand(publicCmd); err == nil {
		log.Infof("GatewayConfig: public ip is: %s, do not configure private gateway", out)
		return nil
	}

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
	netDeviceCfgFile := fmt.Sprintf("/etc/sysconfig/network-scripts/ifcfg-%s", strings.Replace(netDevice, "\n", "", -1))
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

	// confirm
	confirmCmd := fmt.Sprintf("ip route | grep %s", gateway)
	if out, err := utils.RunCommand(confirmCmd); err != nil {
		return err
	} else if out == "" {
		return fmt.Errorf("GatewayConfig: confirm gateway configuration failed")
	}

	log.Infof("GatewayConfig: succeed!")
	return nil
}
