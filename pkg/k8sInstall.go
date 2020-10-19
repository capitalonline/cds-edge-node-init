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
	out, err:= utils.RunCommand("kubelet --version")
	if err != nil && strings.Contains(err.Error(), "kubelet: command not found") {
		// kubelet not installed, start do it
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

		// confirm
		confirmCmd := fmt.Sprintf("kubelet --version")
		out, err := utils.RunCommand(confirmCmd)
		if err != nil {
			return err
		}
		if !strings.Contains(out, k8sV17InitData.K8sInstall.Version) {
			return fmt.Errorf("confirm kubelet install version %s failed", k8sV17InitData.K8sInstall.Version)
		}

	}

	if strings.Contains(out, k8sV17InitData.K8sInstall.Version) {
		log.Warnf("kubelet %s installed", k8sV17InitData.K8sInstall.Version)
		return nil
	} else {
		log.Errorf("kubelet is installed, but out is: %s is different from expect version: %s", out, k8sV17InitData.K8sInstall.Version)
		return fmt.Errorf("kubelet is installed, but out is: %s is different from expect version: %s", out, k8sV17InitData.K8sInstall.Version)
	}

	log.Infof("K8sInstall: Succeed!")
	return nil
}
