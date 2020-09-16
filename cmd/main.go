package cmd

import (
	"flag"

	log "github.com/sirupsen/logrus"
)



func init() {
	flag.Set("logtostderr", "true")
	flag.Parse()
}

func main () {
	log.Infof("Init main")
}