package main

import (
	"flag"
	"github.com/capitalonline/cds-edge-node-init/pkg"
	log "github.com/sirupsen/logrus"
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Parse()
}

func main() {
	log.Infof("Init main")

	//if err := pkg.SystemConfig(); err != nil {
	//	log.Errorf("SystemConfig: failed, err is: %s", err.Error())
	//}
	//
	//if err:= pkg.YumConfig(); err != nil {
	//	log.Errorf("YumConfig: failed, err is: %s", err.Error())
	//}
	//
	//if err:= pkg.PythonInstall(); err != nil {
	//	log.Errorf("PythonInstall: failed, err is: %s", err.Error())
	//}

	if err:= pkg.DockerInstall(); err != nil {
		log.Errorf("PythonInstall: failed, err is: %s", err.Error())
	}

	log.Infof("Finished init main")

}