package fusion_rest_go

import (
	"bytes"
	"net/http"
)
import "encoding/json"
import "fmt"

const VmsUrl string = DefaultRestUrl + "/vms"

type Vm struct {
	Id     string `json:"id"`
	Path   string `json:"path"`
	Cpu    string `json:"cpu"`
	Memory string `json:"memory"`
}

type NewVm struct {
	Cpu      string `json:"cpu"`
	Memory   string `json:"memory"`
	ParentId string `json:"parent_id"`
}

func (c *Client) GetVms() ([]Vm, error) {
	req, err := http.NewRequest("GET", VmsUrl, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var vms []Vm

	err = json.Unmarshal(res, &vms)
	if err != nil {
		return nil, err
	}

	return vms, nil
}

func (c *Client) GetVm(id string) (*Vm, error) {
	var vm Vm
	req, err := http.NewRequest("GET", fmt.Sprintf(VmsUrl+"/%s", id), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &vm)
	if err != nil {
		return nil, err
	}

	return &vm, nil
}

func (c *Client) UpdateVm(t *Vm) error {
	j, err := json.Marshal(t)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", VmsUrl+"/"+t.Id, bytes.NewBuffer(j))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) CreateVm(t *NewVm) (*Vm, error) {
	j, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", VmsUrl, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var vm Vm
	err = json.Unmarshal(res, &vm)
	if err != nil {
		return nil, err
	}

	return &vm, nil
}

func (c *Client) DeleteVm(id string) error {
	req, err := http.NewRequest("DELETE", VmsUrl+"/"+id, nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}