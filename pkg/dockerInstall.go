package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
)

// version 19.03.11
func DockerInstall(k8sV17InitData *utils.K8sV17Config) error {
	log.Infof("DockerInstall: %s starting", k8sV17InitData.DockerInstall.Version)

	// check
	checkCmd := fmt.Sprintf("docker --version")
	if out, err := utils.RunCommand(checkCmd); err == nil {
		log.Warnf("DockerInstall: docker %s installed, ignore install again!", out)
		// make sure docker running
		startCmd := fmt.Sprintf("systemctl start docker && systemctl enable docker")
		if _, err := utils.RunCommand(startCmd); err != nil {
			log.Errorf("DockerInstall: start docker failed, err is: %s", err)
			return err
		}

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
	installDockerSlice := []string{"docker-ce-" + k8sV17InitData.DockerInstall.Version, "docker-ce-cli-" + k8sV17InitData.DockerInstall.Version, "containerd.io"}
	if out, err := utils.InstallPkgs(installDockerSlice, false); err != nil {
		log.Warnf("PythonInstall: pkgs: %s install failed, retry", out)
		if _, err := utils.InstallPkgs(out, false); err != nil {
			log.Errorf("DockerInstall: pkgs: %s install failed again, err is: %s", out, err)
			return err
		}
	}

	// create docker dir
	if !utils.FileExisted("/etc/docker") {
		if err := utils.CreateDir("/etc/docker", 755); err != nil {
			log.Errorf("DockerInstall: create /etc/docker dir failed, err is: %s", err)
			return err
		}
	}

	// wget docker daemon.json
	wgetCmd := fmt.Sprintf("wget -O /etc/docker/daemon.json %s", k8sV17InitData.DockerInstall.DaemonFile)
	if _, err := utils.RunCommand(wgetCmd); err != nil {
		log.Errorf("DockerInstall: wget daemon.json failed, err is: %s", err)
		return err
	}

	// confirm
	confirmCmd := fmt.Sprintf("docker --version")
	if _, err := utils.RunCommand(confirmCmd); err != nil {
		log.Errorf("DockerInstall: confirm docker version failed, err is: %s", err)
		return err
	}

	// start docker
	startCmd := fmt.Sprintf("systemctl start docker && systemctl enable docker")
	if _, err := utils.RunCommand(startCmd); err != nil {
		log.Errorf("DockerInstall: start docker failed, err is: %s", err)
		return err
	}

	log.Infof("DockerInstall: Succeed!")
	return nil
}
