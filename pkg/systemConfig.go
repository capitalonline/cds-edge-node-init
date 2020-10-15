package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
	"strings"
)

func SystemConfig () error {
	log.Infof("SystemConfig: Starting")

	// selinux config
	if err := selinuxConfig(); err != nil {
		return err
	}

	// firewalld and /etc/sysctl.conf config
	if err := firewalldConfig(); err != nil {
		return err
	}

	log.Infof("SystemConfig: Succeed!")
	return nil
}


func selinuxConfig() error {
	selinuxConfigCmd := fmt.Sprintf("setenforce 0 && sed -i 's/SELINUX=enforcing/SELINUX=disabled/g' /etc/sysconfig/selinux")
	if out, err := utils.RunCommand("getenforce"); err == nil {
		if strings.Contains(out, "Disabled") {
			selinuxConfigCmd = fmt.Sprintf("sed -i 's/SELINUX=enforcing/SELINUX=disabled/g' /etc/sysconfig/selinux")
			utils.RunCommand(selinuxConfigCmd)
		}
	} else {
		return err
	}

	return nil
}

func firewalldConfig () error {
	firewallCmd := fmt.Sprintf("systemctl stop firewalld && systemctl disable firewalld")
	if _, err := utils.RunCommand(firewallCmd); err != nil {
		return  err
	}

	//wgetCmd := fmt.Sprintf("wget -P /tmp http://%s/sysctl.conf", utils.CdsOssAddress)
	//if _, err := utils.RunCommand(wgetCmd); err != nil {
	//	return err
	//}

	modifyCmd := fmt.Sprintf("mv /etc/sysctl.conf /etc/bak-sysctl.conf && wget -P /etc http://%s/sysctl.conf", utils.CdsOssAddress)
	if _, err := utils.RunCommand(modifyCmd); err != nil {
		return  err
	}

	return nil
}