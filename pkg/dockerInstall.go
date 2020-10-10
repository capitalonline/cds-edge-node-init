package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
	"strings"
)

// version 19.03.11
func DockerInstall () error {
	log.Infof("DockerInstall: starting")

	// create docker dir
	if !utils.FileExisted("/data/kubernetes/docker") {
		if err := utils.CreateDir("/data/kubernetes/docker", 755); err != nil {
			log.Errorf("DockerInstall: create docker dir failed, err is: %s", err.Error())
			return err
		}
	}

	// config docker repo
	repoCmd := fmt.Sprintf("cd /data/kubernetes/docker && yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo && yum makecache fast")
	if _, err := utils.RunCommand(repoCmd); err != nil {
		log.Errorf("DockerInstall: config docker repo failed, err is: %s", err.Error())
		return err
	}

	// install docker
	installCmd := fmt.Sprintf(" yum install -y docker-ce-19.03.11 docker-ce-cli-19.03.11 containerd.io")
	if _, err := utils.RunCommand(installCmd); err != nil {
		log.Errorf("DockerInstall: install docker failed, err is: %s", err.Error())
		return err
	}

	// confirm
	confirmCmd := fmt.Sprintf("docker --version")
	if _, err := utils.RunCommand(confirmCmd); err != nil {
		log.Errorf("DockerInstall: install docker failed, err is: %s", err.Error())
		return err
	}

	if !strings.Contains(confirmCmd, "Docker") {
		log.Errorf("DockerInstall: install docker failed")
		return fmt.Errorf("DockerInstall: install docker failed")
	}

	log.Infof("DockerInstall: Succeed!")
	return nil
}
