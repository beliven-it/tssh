package interfaces

import (
	"encoding/json"
	"tssh/types"
	"tssh/utils"
)

type goteleport struct {
}

type Goteleport interface {
	ListRoles() ([]string, error)
	Connect(string) error
}

func (t *goteleport) ListRoles() ([]string, error) {
	output, err := utils.Exec("tsh", "status", "--format=json")
	if err != nil {
		return nil, err
	}

	var response types.GoteleportCMDStatus
	err = json.Unmarshal(output, &response)
	if err != nil {
		return nil, err
	}

	return response.Active.Roles, nil
}

func (t *goteleport) Connect(connection string) error {
	err := utils.ExecStdout("tsh", "ssh", connection)
	if err != nil {
		return err
	}

	return nil
}

func NewGoteleportInterface() Goteleport {
	return &goteleport{}
}
