package interfaces

import (
	"encoding/json"
	"tssh/types"
	"tssh/utils"
)

type goteleport struct {
	status types.GoteleportActive
}

type Goteleport interface {
	ListRoles() ([]string, error)
	ListHosts() ([]string, error)
	ListLogins() ([]string, error)
	Connect(string) error
}

func (t *goteleport) getStatus() error {
	output, err := utils.Exec("tsh", "status", "--format=json")
	if err != nil {
		return err
	}

	var response types.GoteleportCMDStatus
	err = json.Unmarshal(output, &response)
	if err != nil {
		return err
	}

	t.status = response.Active

	return nil
}

func (t *goteleport) ListHosts() ([]string, error) {
	output, err := utils.Exec("tsh", "ls", "--format=json")
	if err != nil {
		return nil, err
	}

	var response []types.GoteleportNode
	err = json.Unmarshal(output, &response)
	if err != nil {
		return nil, err
	}

	hostnames := []string{}
	for _, node := range response {
		hostnames = append(hostnames, node.Spec.Hostname)
	}

	return hostnames, nil
}

func (t *goteleport) ListRoles() ([]string, error) {
	return t.status.Roles, nil
}

func (t *goteleport) ListLogins() ([]string, error) {
	return t.status.Logins, nil
}

func (t *goteleport) Connect(connection string) error {
	err := utils.ExecStdout("tsh", "ssh", connection)
	if err != nil {
		return err
	}

	return nil
}

func NewGoteleportInterface() Goteleport {

	i := goteleport{}
	i.getStatus()

	return &i
}
