package run

var (
	K8sV17   = "1.17.0"
)

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
	Version string `json:"version"`
	Sysctl  string `json:"sysctl"`
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
