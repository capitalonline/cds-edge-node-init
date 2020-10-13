package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
)

func SystemConfig () error {
	log.Infof("SystemConfig: Starting")
	// firewall
	firewallCmd := fmt.Sprintf("setenforce 0 && sed -i 's/SELINUX=enforcing/SELINUX=disabled/g' /etc/sysconfig/selinux && systemctl stop firewalld && systemctl disable firewalld")
	// log.Infof("firewallCmd is: %s", firewallCmd)

	_, err := utils.RunCommand(firewallCmd)
	if err != nil {
		log.Errorf("SystemConfig: firewallCmd error, err is: %s", err.Error())
		return  err
	}

	// rewrite /etc/sysctl.conf
	wgetCmd := fmt.Sprintf("wget -P /tmp http://%s/sysctl.conf", utils.CdsOssAddress)
	if _, err := utils.RunCommand(wgetCmd); err != nil {
		return err
	}

	// back /etc/sysctl.conf and rewrite it
	modifyCmd := fmt.Sprintf("mv /etc/sysctl.conf /etc/bak-sysctl.conf && cp /tmp/sysctl.conf /etc/sysctl.conf")
	if _, err := utils.RunCommand(modifyCmd); err != nil {
		return  err
	}

	log.Infof("SystemConfig: Succeed!")
	return nil
}
