package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
)

func ImagePullAndLoad(k8sV17InitData *utils.K8sV17Config) error {
	log.Infof("ImagePullAndLoad: Starting")

	wgetCmd := fmt.Sprintf("wget -O /tmp/k8sV1.17.0-DockerImages.tar.gz %s", k8sV17InitData.DockerImages.ImageTar)
	if _, err := utils.RunCommand(wgetCmd); err != nil {
		return err
	}

	loadCmd := fmt.Sprintf("gunzip -c /tmp/k8sV1.17.0-DockerImages.tar.gz | docker load")
	if _, err := utils.RunCommand(loadCmd); err != nil {
		return err
	}

	log.Infof("ImagePullAndLoad: Succeed!")
	return nil
}
