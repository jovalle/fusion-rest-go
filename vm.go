package fusion_rest_go

import "net/http"
import "encoding/json"
import "fmt"

const VmsUrl string = DefaultRestUrl + "/vms"

type VM struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
}

func (c *Client) GetVMs() ([]VM, error) {
	req, err := http.NewRequest("GET", VmsUrl, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var vms []VM

	err = json.Unmarshal(res, &vms)
	if err != nil {
		return nil, err
	}

	return vms, nil
}

func (c *Client) GetVM(id string) (*VM, error) {
	var vm VM
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
