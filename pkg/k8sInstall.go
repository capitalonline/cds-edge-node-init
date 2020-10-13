package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	"strings"
)

func K8sInstall () error {
	// wget kubernetes.repo
	wgetCmd := fmt.Sprintf("wget -P /etc/yum.repos.d http://%s/kubernetes.repo", utils.CdsOssAddress)
	if _, err := utils.RunCommand(wgetCmd); err != nil {
		return err
	}

	// install kubeadm and kubelet and kubectl v1.17.0
	installCmd := fmt.Sprintf("yum install -y kubeadm-1.17.0-0 kubelet-1.17.0-0 kubectl-1.17.0-0 --disableexcludes=kubernetes")
	if _, err := utils.RunCommand(installCmd); err != nil {
		return err
	}

	// confirm
	confirmCmd := fmt.Sprintf("kubelet --version")
	out, err := utils.RunCommand(confirmCmd)
	if err != nil {
		return err
	}
	if !strings.Contains(out, "Kubernetes v1.17.0") {
		return fmt.Errorf("confirm kubelet install version failed")
	}

	return nil
}
