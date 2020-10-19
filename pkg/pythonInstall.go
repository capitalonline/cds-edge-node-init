package pkg

import (
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
	"strings"
)

func PythonInstall (k8sV17InitData *utils.K8sV17Config) error {
	log.Infof("PythonInstall: Starting")

	// check
	checkCmd := fmt.Sprintf("python --version")
	out, _ := utils.RunCommand(checkCmd)
	if strings.Contains(out, k8sV17InitData.PythonInstall.Version) {
		log.Warnf("PythonInstall: Python %s installed, ignore install again!", k8sV17InitData.PythonInstall.Version)
		return nil
	}

	// install necessary pkgs
	// installPkgs := []string{"zlib-devel", "bzip2-devel", "openssl-devel", "openssl-static", "ncurses-devel", "sqlite-devel", "readline-devel", "gdbm-devel", "db4-devel", "libpcap-devel", "xz-devel", "libffi-devel", "lzma", "gcc", "tk-devel"}
	if out, err := utils.InstallPkgs(k8sV17InitData.PythonInstall.Pkgs, false); err != nil {
		log.Warnf("PythonInstall: some pkgs install failed, retry")
		if out, err = utils.InstallPkgs(out, false); err != nil {
			log.Errorf("PythonInstall: pkgs: %s install failed again, err is: %s", out, err.Error())
			return err
		}
	}

	// groupInstallPkgs := []string{"Development tools"}
	if out, err := utils.InstallPkgs(k8sV17InitData.PythonInstall.Group, true); err != nil {
		if out, err = utils.InstallPkgs(out, true); err != nil {
			log.Errorf("PythonInstall: pkgs: %s install failed again, err is: %s", out, err.Error())
			return err
		}
	}

	//if out, err := utils.InstallPkgs([]string{k8sV17InitData.PythonInstall.Group}, true); err != nil {
	//	log.Warnf("PythonInstall: group pkgs install failed, retry")
	//	if out, err = utils.InstallPkgs(out, false); err != nil {
	//		log.Errorf("PythonInstall: pkgs: %s install failed again, err is: %s", out, err.Error())
	//		return err
	//	}
	//}

	// install python 3.6
	wgetPythonCmd := fmt.Sprintf("wget -P /usr/local %s", k8sV17InitData.PythonInstall.Install)
	if _, err := utils.RunCommand(wgetPythonCmd); err != nil {
		log.Warnf("PythonInstall: wget python failed, retry")
		if _, err := utils.RunCommand(wgetPythonCmd); err != nil {
			log.Errorf("PythonInstall: wget python failed again, err is: %s", err.Error())
			return err
		}
	}

	installPythonCmd := fmt.Sprintf("cd /usr/local && tar Jxvf Python-3.6.3.tar.xz && mv Python-3.6.3 python3 && cd /usr/local/python3 && ./configure --prefix=/usr/local && make && make install")
	if _, err := utils.RunCommand(installPythonCmd); err != nil {
		log.Errorf("PythonInstall: install python failed, err is: %s", err.Error())
		return err
	}

	// config python3
	configPythonCmd := fmt.Sprintf("cd /usr/bin/ && rm -f python && rm -f pip && ln -s /usr/local/bin/python3 /usr/bin/python && ln -s /usr/local/bin/pip3 /usr/bin/pip")
	if _, err := utils.RunCommand(configPythonCmd); err != nil {
		log.Errorf("PythonInstall: config python3 failed, err is: %s", err.Error())
		return err
	}

	// confirm installed version
	confirmCmd := fmt.Sprintf("python --version && pip --version")
	out, err := utils.RunCommand(confirmCmd)
	if err != nil || !(strings.Contains(out, "Python") && strings.Contains(out, "pip")) {
		log.Errorf("PythonInstall: config python3 failed, err is: %s", err.Error())
		return err
	}

	// modify system's python version
	modifyCmd := fmt.Sprintf("sed -i '1s/python/python2/g' /usr/bin/yum && sed -i '1s/python/python2/g' /usr/bin/yum-config-manager && sed -i '1s/python/python2/g' /usr/libexec/urlgrabber-ext-down")
	if _, err := utils.RunCommand(modifyCmd); err != nil {
		log.Errorf("PythonInstall: modify python version, err is: %s", err.Error())
		return err
	}

	log.Infof("PythonInstall: Succeed!")
	return nil
}
