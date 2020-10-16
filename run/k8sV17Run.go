package run

import (
	"encoding/json"
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/pkg"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

func K8sV17Run (k8s17InitJsonUrl string) {
	log.Infof("K8sV17Run: init starting")

	// wget k8s17InitJsonUrl
	wgetInstall := []string{"wget"}
	if _, err := utils.InstallPkgs(wgetInstall, false); err != nil {
		log.Fatalf("K8sV17Run: install wget failed, err is: %s", err)
	}

	wgetCmd := fmt.Sprintf("wget -P /tmp %s", k8s17InitJsonUrl)
	if _, err := utils.RunCommand(wgetCmd); err != nil {
		log.Warnf("K8sV17Run: wget k8s17InitJson failed, retry 2s")
		time.Sleep(time.Second *2)
		if _, err := utils.RunCommand(wgetCmd); err != nil {
			log.Fatalf("K8sV17Run: wget k8s17InitJson failed AGAIN, err is: %s", err)
		}
	}

	// read /tmp/k8sV1.17.0Install.json and unmarshal
	var k8sV17InitData utils.K8sV17Config
	if res, err := ioutil.ReadFile("/tmp/k8sV1.17.0Install.json"); err != nil {
		log.Fatalf(err.Error())
	} else {
		if err = json.Unmarshal(res, &k8sV17InitData); err != nil {
			log.Fatalf(err.Error())
		}
	}

	// init
	switch k8sV17InitData.K8sInstall.Version {
	case utils.K8sV17:
		if err := pkg.SystemConfig(&k8sV17InitData); err != nil {
			log.Fatalf("SystemConfig: failed, err is: %s", err.Error())
		}

		if err:= pkg.YumConfig(&k8sV17InitData); err != nil {
			log.Fatalf("YumConfig: failed, err is: %s", err.Error())
		}

		if err:= pkg.PythonInstall(&k8sV17InitData); err != nil {
			log.Fatalf("PythonInstall: failed, err is: %s", err.Error())
		}

		if err:= pkg.DockerInstall(&k8sV17InitData); err != nil {
			log.Fatalf("DockerInstall: failed, err is: %s", err.Error())
		}

		if err:= pkg.ImagePullAndLoad(&k8sV17InitData); err != nil {
			log.Fatalf("ImagePullAndTag: failed, err is: %s", err.Error())
		}

		if err:= pkg.K8sInstall(&k8sV17InitData); err != nil {
			log.Fatalf("K8sInstall: failed, err is: %s", err)
		}

		if err:= pkg.NetworkConfig(&k8sV17InitData); err != nil {
			log.Fatalf("NetworkConfig: failed, err is: %s", err)
		}
	default:
		log.Fatalf("K8sV17Run: unsupported k8s install version: %s", k8sV17InitData.K8sInstall.Version)
	}

	log.Infof("K8sV17Run: Successfully!")
}
