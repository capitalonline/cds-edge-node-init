package cmd

import (
	"flag"
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/pkg"
	log "github.com/sirupsen/logrus"
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Parse()
}

func main() {
	log.Infof("Init main")

	//err := pkg.SystemConfig()
	//fmt.Println(err)

	//if err := pkg.YumConfig(); err != nil {
	//	fmt.Errorf("DockerInstall: failed, err is: %s\n", err.Error())
	//}

	if err := pkg.SystemConfig(); err != nil {
		fmt.Errorf("SystemConfig: failed, err is: %s\n", err.Error())
	}

	fmt.Printf("Finished init main\n")

}