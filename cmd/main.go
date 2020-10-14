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

func main () {
	log.Infof("Init main")

	//err := pkg.SystemConfig()
	//fmt.Println(err)

	if err := pkg.DockerInstall(); err != nil {
		fmt.Errorf("DockerInstall: failed, err is: %s\n", err.Error())
	}
	
	fmt.Printf("DockerInstall: Succeed! \n")

}