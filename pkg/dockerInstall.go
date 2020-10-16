package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/run"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
	"strings"
)

// version 19.03.11
func DockerInstall (k8sV17InitData *run.K8sV17Config) error {
	log.Infof("DockerInstall: starting")

	// check
	checkCmd := fmt.Sprintf("docker --version")
	if out, _ := utils.RunCommand(checkCmd); strings.Contains(out, k8sV17InitData.DockerInstall.Version) {
		log.Warnf("DockerInstall: installed, ignore install again!")
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
	repoCmd := fmt.Sprintf("cd /data/kubernetes/docker && yum-config-manager --add-repo %s && yum makecache fast", k8sV17InitData.DockerInstall.RepoAdd)
	if _, err := utils.RunCommand(repoCmd); err != nil {
		log.Errorf("DockerInstall: config docker repo failed, err is: %s", err)
		return err
	}

	// install docker
	installDockerSlice := []string{"docker-ce-"+k8sV17InitData.DockerInstall.Version, "docker-ce-cli-"+k8sV17InitData.DockerInstall.Version, "containerd.io"}
	if out, err := utils.InstallPkgs(installDockerSlice, false); err != nil {
		log.Warnf("PythonInstall: some pkgs install failed, retry")
		if _, err := utils.InstallPkgs(out, false); err != nil {
			log.Errorf("DockerInstall: install docker failed, err is: %s", err)
			return err
		}
	}

	// wget docker daemon.json
	wgetCmd := fmt.Sprintf("wget -P /etc/docker %s", k8sV17InitData.DockerInstall.DaemonFile)
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

	if !strings.Contains(out, k8sV17InitData.DockerInstall.Version) {
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
