package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/capitalonline/cds-edge-node-init/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func TunnelSetup(initData *utils.InitData) error {
	log.Infof("TunnelSetup: starting")

	// get setup parameters
	resParams, err := tunnelGetParams(initData)
	if err != nil {
		return err
	}

	// setup tunnel
	setupCmd := fmt.Sprintf("docker run -d --restart always --env SERVER_ADDR=%s --env SERVER_PORT=%s --env AUTH_TOKEN=%s --env REMOTE_PORT=%s --net host --name cck-agent %s/agent:%s", resParams.Data.TunnelAddress, resParams.Data.ServerPort, resParams.Data.Token, resParams.Data.TunnelPort, resParams.Data.ImageUrl, resParams.Data.Version)
	if _, err := utils.RunCommand(setupCmd); err != nil {
		return err
	}

	// inform
	resInit, err := tunnelInit(initData, resParams.Data.NodeID, initData.PrivateIP)
	if err != nil {
		return err
	}

	log.Infof("TunnelSetup: taskId: %s", resInit.Data.TaskId)
	log.Infof("TunnelSetup: succeed!")
	return nil
}

func tunnelGetParams(initData *utils.InitData) (*utils.TunnelGetResponse, error) {
	//log.Infof("TunnelSetup-tunnelGetParams: starting")
	payload := struct {
		UserId     string `json:"user_id"`
		CustomerId string `json:"customer_id"`
		Flag       string `json:"flag"`
		Data       struct {
			ClusterId    string `json:"cluster_id"`
			RootPassword string `json:"root_password"`
		} `json:"data"`
	}{
		initData.UserID,
		initData.CustomerID,
		"tunnel",
		struct {
			ClusterId    string `json:"cluster_id"`
			RootPassword string `json:"root_password"`
		}{ClusterId: initData.ClusterID, RootPassword: initData.RootPassword},
	}

	body, err := utils.MarshalJsonToIOReader(payload)
	if err != nil {
		return nil, err
	}

	req, err := utils.NewCCKRequest("AddExternalNode", http.MethodPost, nil, body)
	response, err := utils.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &utils.TunnelGetResponse{}
	err = json.Unmarshal(content, res)
	if err != nil {
		return nil, err
	}

	//log.Infof("TunnelSetup-tunnelGetParams: finish")
	return res, nil
}

func tunnelInit(initData *utils.InitData, nodeId, ip string) (*utils.TunnelInitResponse, error) {
	//log.Infof("TunnelSetup-tunnelInit: starting")
	payload := struct {
		UserId     string `json:"user_id,omitempty"`
		CustomerId string `json:"customer_id,omitempty"`
		Flag       string `json:"flag,omitempty"`
		Data       struct {
			ClusterId    string `json:"cluster_id"`
			NodeId       string `json:"node_id"`
			RootPassword string `json:"root_password"`
			Ip           string `json:"ip"`
		} `json:"data"`
	}{
		initData.UserID,
		initData.CustomerID,
		"init",
		struct {
			ClusterId    string `json:"cluster_id"`
			NodeId       string `json:"node_id"`
			RootPassword string `json:"root_password"`
			Ip           string `json:"ip"`
		}{ClusterId: initData.ClusterID, NodeId: nodeId, RootPassword: initData.RootPassword, Ip: ip},
	}

	body, err := utils.MarshalJsonToIOReader(payload)
	if err != nil {
		return nil, err
	}

	req, err := utils.NewCCKRequest("AddExternalNode", http.MethodPost, nil, body)
	response, err := utils.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &utils.TunnelInitResponse{}
	err = json.Unmarshal(content, res)
	if err != nil {
		return nil, err
	}

	//log.Infof("TunnelSetup-tunnelInit: finish")
	return res, nil
}
