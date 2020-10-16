package main

import (
	"flag"
	"github.com/capitalonline/cds-edge-node-init/run"
	log "github.com/sirupsen/logrus"
)

const (
	k8sV17Version           = "1.17.0"
	k8s17InitDefaultJsonUrl = "http://cds-edge-node-init.209faf3a84524f9f81d71f2c0be97de3.oss-cnbj01.cdsgss.com/k8sVersion/k8sV1.17.0Install.json"
)

var (
	k8sVersion = flag.String("version", "1.17.0", "k8s init version")
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Parse()
}

func main() {
	version := *k8sVersion

	switch version {
	case k8sV17Version:
		run.K8sV17Run(k8s17InitDefaultJsonUrl)
	default:
		log.Fatalf("unsupported k8sVersion: %s", version)
	}

}
