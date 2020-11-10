package main

import (
	"flag"
	"github.com/capitalonline/cds-edge-node-init/run"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
)

const (
	k8sV17Version            = "1.17.0"
	k8sV17InitDefaultJsonUrl = "http://cds-edge-node-init.209faf3a84524f9f81d71f2c0be97de3.oss-cnbj01.cdsgss.com/centos7.6/k8sV1.17.0Install.json"
)

var (
	k8sVersion   = flag.String("version", "1.17.0", "k8s init version")
	clusterID    = flag.String("cluster_id", "", "cluster's id")
	rootPassword = flag.String("root_password", "", "node's root password")
	ak           = flag.String("ak", "", "access key")
	sk           = flag.String("sk", "", "secret key")
	userID       = flag.String("user_id", "", "user id")
	customerID   = flag.String("customer_id", "", "customer id")
	gateWay      = flag.String("gateway", "", "cluster's ros gateway")
	privateIP    = flag.String("private_ip", "", "global private ip")
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	// init cds api ak and sk
	utils.AccessKeyID = *ak
	utils.AccessKeySecret = *sk
}

func main() {
	version := *k8sVersion

	var InitInfo utils.InitData
	InitInfo.ClusterID = *clusterID
	InitInfo.Ak = *ak
	InitInfo.Sk = *sk
	InitInfo.RootPassword = *rootPassword
	InitInfo.UserID = *userID
	InitInfo.ClusterID = *customerID
	InitInfo.Gateway = *gateWay
	InitInfo.PrivateIP = *privateIP
	InitInfo.K8sVersion = *k8sVersion

	switch version {
	case k8sV17Version:
		run.K8sV17Run(k8sV17InitDefaultJsonUrl, &InitInfo)
	default:
		log.Fatalf("unsupported k8sVersion: %s", version)
	}

}
