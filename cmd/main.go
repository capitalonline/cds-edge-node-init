package main

import (
	"flag"
	"github.com/capitalonline/cds-edge-node-init/run"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
)

const (
	k8sV17Version            = "1.17.0"
	k8sV17InitDefaultJsonUrl = "http://cds-edge-node-init.209faf3a84524f9f81d71f2c0be97de3.oss-cnbj01.cdsgss.com/k8sV17Init/k8sV1.17.0Install.json"
)

var (
	k8sVersion   = flag.String("k8s_version", "1.17.0", "k8s init version")
	clusterID    = flag.String("cluster_id", "", "cluster's id")
	rootPassword = flag.String("root_password", "", "node's root password")
	ak           = flag.String("ak", "", "access key")
	sk           = flag.String("sk", "", "secret key")
	userID       = flag.String("user_id", "", "user id")
	customerID   = flag.String("customer_id", "", "customer id")
	gateWay      = flag.String("gateway", "", "cluster's ros gateway")
	privateIP    = flag.String("private_ip", "", "global private ip")
	oversea      = flag.String("oversea", "", "oversea")
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	// init cds api ak and sk
	utils.AccessKeyID = *ak
	utils.AccessKeySecret = *sk

	// init APIHost
	if "true" == *oversea {
		utils.APIHost = "http://cdsapi-us.capitalonline.net"
	} else if "false" == *oversea {
		utils.APIHost = "http://cdsapi.capitalonline.net"
	} else {
		log.Fatalf("unsupported oversea: %s, should be one of [true|false]!", *oversea)
	}
}

func main() {
	version := *k8sVersion

	var InitInfo utils.InitData
	InitInfo.UserID = *userID
	InitInfo.CustomerID = *customerID
	InitInfo.Ak = *ak
	InitInfo.Sk = *sk
	InitInfo.ClusterID = *clusterID
	InitInfo.RootPassword = *rootPassword
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
