package utils

import (
	"log"
)

var (
	K8sV17         = "1.17.0"
	FlagDebugMode  *bool
	Logger         *log.Logger
	MaxWaitingTime = 200 //seconds

	AccessKeyID     string
	AccessKeySecret string
)

const (
	cckProductType   = "cck"
	version          = "2019-08-08"
	signatureVersion = "1.0"
	signatureMethod  = "HMAC-SHA1"
	timeStampFormat  = "2006-01-02T15:04:05Z"
	apiHost          = "http://cdsapi.capitalonline.net"
)

type BaseResponse struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
}

type InitData struct {
	K8sVersion       string
	ClusterID        string
	RootPassword     string
	Ak               string
	Sk               string
	UserID           string
	CustomerID       string
	Gateway          string
	PrivateIP		 string
}

type TunnelGetResponse struct {
	BaseResponse
	Data struct {
		NodeID        string `json:"node_id"`
		TunnelAddress string `json:"tunnel_address"`
		TunnelPort    string `json:"tunnel_port"`
		IdRsaPub      string `json:"id_rsa_pub"`
		Token         string `json:"token"`
		Version       string `json:"version"`
		ServerPort    string `json:"server_port"`
		ImageUrl      string `json:"image_url"`
	} `json:"data"`
}

type TunnelInitResponse struct {
	BaseResponse
	Data struct {
		ClusterID string `json:"cluster_id"`
		TaskId    string `json:"task_id"`
	} `json:"data"`
}
type K8sV17Config struct {
	K8sInstall    k8s
	SystemConfig  config
	YumConfig     yum
	PythonInstall python
	DockerInstall docker
	DockerImages  images
	NetworkConfig network
}

type k8s struct {
	Version string   `json:"version"`
	RepoAdd string   `json:"repoAdd"`
	Install []string `json:"install"`
}

type config struct {
	Version     string `json:"version"`
	Sysctl      string `json:"sysctl"`
	NtpdConfUrl string `json:"ntpdConfUrl"`
}

type yum struct {
	Pkgs        []string `json:"pkgs"`
	RepoReplace []string `json:"repoReplace"`
}

type python struct {
	Version string   `json:"version"`
	Pkgs    []string `json:"pkgs"`
	Group   []string `json:"group"`
	Install string   `json:"install"`
}

type docker struct {
	Version    string `json:"version"`
	RepoAdd    string `json:"repoAdd"`
	DaemonFile string `json:"daemonFile"`
}

type images struct {
	ImageTar string `json:"imageTar"`
}

type network struct {
	Ipvs string   `json:"ipvs"`
	Pkgs []string `json:"pkgs"`
}
