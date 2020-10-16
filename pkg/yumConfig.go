package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/run"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
)

func YumConfig (k8sV17InitData *run.K8sV17Config) error {
	log.Infof("YumConfig: Starting")

	// install common tools
	//pkgInstallSlice := []string{"yum-utils", "device-mapper-persistent-data", "lvm2", "vim", "wget", "lrzsz", "lsof", "ntpdate", "sysstat", "net-tools", "deltarpm", "redhat-lsb"}
	if out, err := utils.InstallPkgs(k8sV17InitData.YumConfig.Pkgs, false); err != nil {
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

	yumBakCmd := fmt.Sprintf("cd /etc/yum.repos.d && mv * %s", repoBackDir)
	if _, err := utils.RunCommand(yumBakCmd); err != nil {
		log.Errorf("YumConfig: yumBakCmd error, err is: %s", err.Error())
		return  err
	}

	for _, value := range k8sV17InitData.YumConfig.RepoReplace {
		wgetRepo := fmt.Sprintf("wget -P /etc/yum.repos.d %s", value)
		if _, err := utils.RunCommand(wgetRepo); err != nil {
			log.Errorf("YumConfig: wgetRepo error, err is: %s", err.Error())
			return  err
		}
	}

	log.Infof("YumConfig: Succeed!")
	return nil
}
