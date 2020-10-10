package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
)

func SystemConfig () error {
	log.Infof("SystemConfig: Starting")

	firewallCmd := fmt.Sprintf("setenforce 0 && sed -i '%s' /etc/sysconfig/selinux && systemctl stop firewalld && systemctl disable firewalld", "s/SELINUX=enforcing/SELINUX=disabled/g")
	fmt.Printf("firewallCmd is: %s", firewallCmd)

	_, err := utils.RunCommand(firewallCmd)
	if err != nil {
		log.Errorf("SystemConfig: firewallCmd error, err is: %s", err.Error())
		return  err
	}

	log.Infof("SystemConfig: Succeed!")
	return nil
}
