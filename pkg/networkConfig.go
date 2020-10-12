package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
)

func NetworkConfig () error {
	// install necessary pkgs
	installPkgs := []string{"ipvsadm", "ipset"}
	if out, err := utils.InstallPkgs(installPkgs, false ); err != nil {
		if _, err := utils.InstallPkgs(out, false); err != nil {
			return nil
		}
	}

	// wget ipvs.modules
	wgetCmd := fmt.Sprintf("wget -P /etc/sysconfig/modules https://***/ipvs.modules")
	if _, err := utils.RunCommand(wgetCmd); err != nil {
		return err
	}

	// config nerwork
	configCmd := fmt.Sprintf("chmod 755 /etc/sysconfig/modules/ipvs.modules && bash /etc/sysconfig/modules/ipvs.modules && lsmod | grep -E 'ip_vs|nf_conntrack_ipv4'")
	if _, err := utils.RunCommand(configCmd); err != nil {
		return err
	}

	return nil
}
