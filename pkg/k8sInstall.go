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
	checkCmd := fmt.Sprintf("kubelet --version")
	if out, _:= utils.RunCommand(checkCmd); strings.Contains(out, k8sV17InitData.K8sInstall.Version) {
		log.Warnf("kubelet %s installed", k8sV17InitData.K8sInstall.Version)
		return nil
	}

	// wget kubernetes.repo
	wgetCmd := fmt.Sprintf("wget -P /etc/yum.repos.d %s", k8sV17InitData.K8sInstall.RepoAdd)
	if _, err := utils.RunCommand(wgetCmd); err != nil {
		return err
	}

	// install kubeadm and kubelet and kubectl v1.17.0
	for _, value := range k8sV17InitData.K8sInstall.Install {
		installCmd := fmt.Sprintf("yum install -y %s", value)
		if _, err := utils.RunCommand(installCmd); err != nil {
			return err
		}
	}
	//installCmd := fmt.Sprintf("yum install -y kubeadm-%s-0 kubelet-%s-0 kubectl-%s-0 --disableexcludes=kubernetes", k8sV17InitData.K8sInstall.Version, k8sV17InitData.K8sInstall.Version, k8sV17InitData.K8sInstall.Version)
	//if _, err := utils.RunCommand(installCmd); err != nil {
	//	return err
	//}

	// confirm
	confirmCmd := fmt.Sprintf("kubelet --version")
	out, err := utils.RunCommand(confirmCmd)
	if err != nil {
		return err
	}
	if !strings.Contains(out, k8sV17InitData.K8sInstall.Version) {
		return fmt.Errorf("confirm kubelet install version %s failed", k8sV17InitData.K8sInstall.Version)
	}

	log.Infof("K8sInstall: Succeed!")
	return nil
}
