package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
)

func ImagePullAndLoad () error {
	log.Infof("ImagePullAndLoad: Starting")

	//imageMasterSlice := []string{
	//	"registry.aliyuncs.com/google_containers/kube-proxy:v1.17.0",
	//	"registry.aliyuncs.com/google_containers/kube-apiserver:v1.17.0",
	//	"registry.aliyuncs.com/google_containers/kube-controller-manager:v1.17.0",
	//	"registry.aliyuncs.com/google_containers/kube-scheduler:v1.17.0",
	//	"calico/cni:v3.10.1",
	//	"calico/pod2daemon-flexvol:v3.10.1",
	//	"registry.aliyuncs.com/google_containers/etcd:3.4.3-0",
	//	"registry.aliyuncs.com/google_containers/coredns:1.6.5",
	//	"registry.aliyuncs.com/google_containers/pause:3.1",
	//}
	//
	//imageNodeSlice := []string{
	//	"registry.aliyuncs.com/google_containers/kube-proxy:v1.17.0",
	//	"registry.aliyuncs.com/google_containers/pause:3.1",
	//}
	//
	//var imagePull []string
	//if node == "master" {
	//	imagePull = imageMasterSlice
	//} else if node == "worker" {
	//	imagePull = imageNodeSlice
	//} else {
	//	return fmt.Errorf("ImagePullAndTag: node must be one of [master|node], input is: %s", node)
	//}
	//
	//if out, err := utils.ImagePullAndTag(imagePull); err != nil {
	//	log.Warnf("ImagePullAndTag: docker images pull and tag failed, retry")
	//	if out, err = utils.ImagePullAndTag(out); err != nil {
	//		log.Errorf("ImagePullAndTag: docker images: %s pull and tag failed again, err is: %s", out, err)
	//		return err
	//	}
	//}

	wgetCmd := fmt.Sprintf("wget -P /tmp http://cds-edge-node-init.209faf3a84524f9f81d71f2c0be97de3.oss-cnbj01.cdsgss.com/images/k8sV1.17.0-DockerImages.tar.gz")
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
