package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/run"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
)

func NetworkConfig (k8sV17InitData *run.K8sV17Config) error {
	log.Infof("NetworkConfig: starting")

	// install necessary pkgs
	//installPkgs := []string{"ipvsadm", "ipset"}
	if out, err := utils.InstallPkgs(k8sV17InitData.NetworkConfig.Pkgs, false ); err != nil {
		if _, err := utils.InstallPkgs(out, false); err != nil {
			return nil
		}
	}

	// wget ipvs.modules
	wgetCmd := fmt.Sprintf("wget -P /etc/sysconfig/modules %s", k8sV17InitData.NetworkConfig.Ipvs)
	if _, err := utils.RunCommand(wgetCmd); err != nil {
		return err
	}

	// config network
	configCmd := fmt.Sprintf("chmod 755 /etc/sysconfig/modules/ipvs.modules && bash /etc/sysconfig/modules/ipvs.modules")
	if _, err := utils.RunCommand(configCmd); err != nil {
		return err
	}

	// confirm
	confirmCmd := fmt.Sprintf("lsmod | grep -E 'ip_vs|nf_conntrack_ipv4'")
	if _, err := utils.RunCommand(confirmCmd); err != nil {
		return err
	}

	log.Infof("NetworkConfig: Succeed!")
	return nil
}
