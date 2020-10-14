package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
)

func YumConfig () error {
	log.Infof("YumConfig: Starting")

	// install common tools
	pkgInstallSlice := []string{"yum-utils", "device-mapper-persistent-data", "lvm2", "vim", "wget", "lrzsz", "lsof", "ntpdate", "sysstat", "net-tools", "deltarpm", "redhat-lsb"}
	if out, err := utils.InstallPkgs(pkgInstallSlice, false); err != nil {
		log.Warnf("YumConfig: some pkgs install failed, retry")
		if out, err = utils.InstallPkgs(out, false); err != nil {
			log.Errorf("YumConfig: pkgs: %s install failed again, err is: %s", out, err.Error())
			return err
		}
	}

	// replace yum repo
	repoBackDir := fmt.Sprintf("/root/repo-bak")
	if !utils.FileExisted(repoBackDir) {
		if err := utils.CreateDir(repoBackDir, 755); err != nil {
			return  err
		}
	}
	yumCmd := fmt.Sprintf("cd /etc/yum.repos.d && mv * %s && wget http://%s/Centos-7.repo && wget http://%s/epel-7.repo", repoBackDir, utils.CdsOssAddress, utils.CdsOssAddress)
	if _, err := utils.RunCommand(yumCmd); err != nil {
		log.Errorf("YumConfig: yumCmd error, err is: %s", err.Error())
		return  err
	}

	log.Infof("YumConfig: Succeed!")
	return nil
}
