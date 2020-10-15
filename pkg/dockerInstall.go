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

	// check
	checkCmd := fmt.Sprintf("docker --version")
	if out, _ := utils.RunCommand(checkCmd); strings.Contains(out, "Docker") {
		log.Infof("DockerInstall: installed, ignore install again!")
		return nil
	}

	// create docker dir
	if !utils.FileExisted("/data/kubernetes/docker") {
		if err := utils.CreateDir("/data/kubernetes/docker", 755); err != nil {
			log.Errorf("DockerInstall: create docker dir failed, err is: %s", err)
			return err
		}
	}

	// config docker repo
	repoCmd := fmt.Sprintf("cd /data/kubernetes/docker && yum-config-manager --add-repo http://%s/docker-ce.repo && yum makecache fast", utils.CdsOssAddress)
	if _, err := utils.RunCommand(repoCmd); err != nil {
		log.Errorf("DockerInstall: config docker repo failed, err is: %s", err)
		return err
	}

	// install docker
	installDockerSlice := []string{"docker-ce-19.03.11", "docker-ce-cli-19.03.11", "containerd.io"}
	if out, err := utils.InstallPkgs(installDockerSlice, false); err != nil {
		log.Warnf("PythonInstall: some pkgs install failed, retry")
		if _, err := utils.InstallPkgs(out, false); err != nil {
			log.Errorf("DockerInstall: install docker failed, err is: %s", err)
			return err
		}
	}

	// wget docker daemon.json
	wgetCmd := fmt.Sprintf("wget -P /etc/docker http://%s/daemon.json", utils.CdsOssAddress)
	if _, err := utils.RunCommand(wgetCmd); err != nil {
		log.Errorf("DockerInstall: wget daemon.json failed, err is: %s", err)
		return err
	}

	// confirm
	confirmCmd := fmt.Sprintf("docker --version")
	out, err := utils.RunCommand(confirmCmd)
	if  err != nil {
		log.Errorf("DockerInstall: install docker failed, err is: %s", err)
		return err
	}

	if !strings.Contains(out, "Docker") {
		//log.Errorf("DockerInstall: install docker failed")
		return fmt.Errorf(out)
	}

	// start docker
	startCmd := fmt.Sprintf("systemctl start docker")
	if _, err := utils.RunCommand(startCmd); err != nil {
		log.Errorf("DockerInstall: start docker failed, err is: %s", err)
		return err
	}

	log.Infof("DockerInstall: Succeed!")
	return nil
}
