package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
	"strings"
)

func K8sInstall (k8sV17InitData *utils.K8sV17Config) error {
	log.Infof("K8sInstall: starting")

	// check
	out, _:= utils.RunCommand("kubelet --version")
	if strings.Contains(out, k8sV17InitData.K8sInstall.Version) {
		log.Warnf("K8sInstall: kubelet %s installed, ignore install again!", k8sV17InitData.K8sInstall.Version)
		return nil
	}

	// wget kubernetes.repo
	wgetCmd := fmt.Sprintf("wget -P /etc/yum.repos.d %s", k8sV17InitData.K8sInstall.RepoAdd)
	if _, err := utils.RunCommand(wgetCmd); err != nil {
		return err
	}

	// install kubeadm and kubelet and kubectl v1.17.0
	var installPkgs string
	for _, value := range k8sV17InitData.K8sInstall.Install {
		installPkgs = installPkgs + " " + value
	}
	installCmd := fmt.Sprintf("yum install -y %s --disableexcludes=kubernetes", installPkgs)
	if _, err := utils.RunCommand(installCmd); err != nil {
		return err
	}

	// confirm
	confirmCmd := fmt.Sprintf("kubelet --version")
	out, err := utils.RunCommand(confirmCmd)
	if !strings.Contains(out, k8sV17InitData.K8sInstall.Version) {
		return fmt.Errorf("confirm kubelet installed version %s failed, err(out) is: %s", k8sV17InitData.K8sInstall.Version, err.Error())
	}

	log.Infof("K8sInstall: Succeed!")
	return nil
}
