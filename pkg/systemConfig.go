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

	// firewalld config
	firewallCmd := fmt.Sprintf("systemctl stop firewalld && systemctl disable firewalld")
	if _, err := utils.RunCommand(firewallCmd); err != nil {
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