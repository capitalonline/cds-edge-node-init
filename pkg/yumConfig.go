package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
)
func YumConfig () error {
	log.Infof("YumConfig: Starting")

	// install some necessary pkgs
	// pkgInstallCmd := fmt.Sprintf("yum install -y yum-utils device-mapper-persistent-data lvm2 vim wget lrzsz lsof ntpdate sysstat net-tools deltarpm redhat-lsb")
	pkgInstallSlice := []string{"yum-utils", "device-mapper-persistent-data", "lvm2", "vim", "wget", "lrzsz", "lsof", "ntpdate", "sysstat", "net-tools", "deltarpm", "redhat-lsb"}
	if out, err := utils.InstallPkgs(pkgInstallSlice, false); err != nil {
		log.Warnf("YumConfig: some pkgs install failed, retry")
		if out, err = utils.InstallPkgs(out, false); err != nil {
			log.Errorf("YumConfig: pkgs: %s install failed again, err is: %s", out, err.Error())
			return err
		}
	}

	// replace yum repo
	yumCmd := fmt.Sprintf("cd /etc/yum.repos.d/ && mkdir bak && mv * bak && wget http://mirrors.aliyun.com/repo/Centos-7.repo && wget http://mirrors.aliyun.com/repo/epel-7.repo")
	if _, err := utils.RunCommand(yumCmd); err != nil {
		log.Errorf("YumConfig: yumCmd error, err is: %s", err.Error())
		return  err
	}

	log.Infof("YumConfig: Succeed!")
	return nil
}
