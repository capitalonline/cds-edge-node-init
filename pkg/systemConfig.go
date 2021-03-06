package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
	"strings"
)

func SystemConfig(k8sV17InitData *utils.K8sV17Config) error {
	log.Infof("SystemConfig: Starting")

	// selinux config
	if err := selinuxConfig(); err != nil {
		return err
	}

	// firewalld and /etc/sysctl.conf config
	if err := firewalldAndSysctlConfig(k8sV17InitData.SystemConfig.Sysctl); err != nil {
		return err
	}

	// ntpd config
	if err := ntpdConfig(k8sV17InitData.SystemConfig.NtpdConfUrl); err != nil {
		return err
	}

	// switch off swap
	swapCmd := fmt.Sprintf("sed -i '/swap/ s/^/#/' /etc/fstab && swapoff -a")
	if _, err := utils.RunCommand(swapCmd); err != nil {
		return err
	}

	// set files limit to 65535
	limitsCmd := fmt.Sprintf("echo '* - nofile 65535' >> /etc/security/limits.conf")
	if _, err := utils.RunCommand(limitsCmd); err != nil {
		return err
	}

	log.Infof("SystemConfig: Succeed!")
	return nil
}

func selinuxConfig() error {
	selinuxConfigCmd := fmt.Sprintf("setenforce 0 && sed -i 's/SELINUX=enforcing/SELINUX=disabled/g' /etc/selinux/config")
	if out, err := utils.RunCommand("getenforce"); err == nil {
		if strings.Contains(out, "Disabled") {
			selinuxConfigCmd = fmt.Sprintf("sed -i 's/SELINUX=enforcing/SELINUX=disabled/g' /etc/selinux/config")
		}
	} else {
		return err
	}

	if _, err := utils.RunCommand(selinuxConfigCmd); err != nil {
		return err
	}
	return nil
}

func firewalldAndSysctlConfig(sysctlUrl string) error {
	firewallCmd := fmt.Sprintf("systemctl stop firewalld && systemctl disable firewalld")
	if _, err := utils.RunCommand(firewallCmd); err != nil {
		return err
	}

	modifyCmd := fmt.Sprintf("mv /etc/sysctl.conf /etc/bak-sysctl.conf && wget -P /etc %s", sysctlUrl)
	if _, err := utils.RunCommand(modifyCmd); err != nil {
		return err
	}

	return nil
}

func ntpdConfig(confUrl string) error {
	ntpInstallCmd := fmt.Sprintf("yum install -y ntp ntpdate")
	if _, err := utils.RunCommand(ntpInstallCmd); err != nil {
		return err
	}

	wgetCmd := fmt.Sprintf("wget -O /etc/ntp.conf %s", confUrl)
	if _, err := utils.RunCommand(wgetCmd); err != nil {
		return err
	}

	ntpEnableCmd := fmt.Sprintf("systemctl restart ntpd && systemctl enable ntpd")
	if _, err := utils.RunCommand(ntpEnableCmd); err != nil {
		return err
	}

	return nil
}
