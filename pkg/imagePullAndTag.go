package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
)

func ImagePullAndTag (node string) error {
	log.Infof("ImagePullAndTag: Starting")

	imageMasterSlice := []string{
		"registry.aliyuncs.com/google_containers/kube-proxy:v1.17.0",
		"registry.aliyuncs.com/google_containers/kube-apiserver:v1.17.0",
		"registry.aliyuncs.com/google_containers/kube-controller-manager:v1.17.0",
		"registry.aliyuncs.com/google_containers/kube-scheduler:v1.17.0",
		"calico/cni:v3.10.1",
		"calico/pod2daemon-flexvol:v3.10.1",
		"registry.aliyuncs.com/google_containers/etcd:3.4.3-0",
		"registry.aliyuncs.com/google_containers/coredns:1.6.5",
		"registry.aliyuncs.com/google_containers/pause:3.1",
	}

	imageNodeSlice := []string{
		"registry.aliyuncs.com/google_containers/kube-proxy:v1.17.0",
		"registry.aliyuncs.com/google_containers/pause:3.1",
	}

	var imagePull []string
	if node == "master" {
		imagePull = imageMasterSlice
	} else if node == "worker" {
		imagePull = imageNodeSlice
	} else {
		return fmt.Errorf("ImagePullAndTag: node must be one of [master|node], input is: %s", node)
	}

	if out, err := utils.ImagePullAndTag(imagePull); err != nil {
		log.Warnf("ImagePullAndTag: docker images pull and tag failed, retry")
		if out, err = utils.ImagePullAndTag(out); err != nil {
			log.Errorf("ImagePullAndTag: docker images: %s pull and tag failed again, err is: %s", out, err)
			return err
		}
	}

	log.Infof("ImagePullAndTag: Succeed!")
	return nil
}
