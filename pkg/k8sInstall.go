package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
	"strings"
)

func K8sInstall (version string) error {
	log.Infof("K8sInstall: starting")

	// check
	checkCmd := fmt.Sprintf("kubelet --version")
	if out, _:= utils.RunCommand(checkCmd); strings.Contains(out, version) {
		log.Warnf("kubelet %s installed", version)
		return nil
	}

	// wget kubernetes.repo
	wgetCmd := fmt.Sprintf("wget -P /etc/yum.repos.d http://%s/kubernetes.repo", utils.CdsOssAddress)
	if _, err := utils.RunCommand(wgetCmd); err != nil {
		return err
	}

	// install kubeadm and kubelet and kubectl v1.17.0
	installCmd := fmt.Sprintf("yum install -y kubeadm-%s-0 kubelet-%s-0 kubectl-%s-0 --disableexcludes=kubernetes", version, version, version)
	if _, err := utils.RunCommand(installCmd); err != nil {
		return err
	}

	// confirm
	confirmCmd := fmt.Sprintf("kubelet --version")
	out, err := utils.RunCommand(confirmCmd)
	if err != nil {
		return err
	}
	if !strings.Contains(out, version) {
		return fmt.Errorf("confirm kubelet install version %s failed", version)
	}

	log.Infof("K8sInstall: Succeed!")
	return nil
}
