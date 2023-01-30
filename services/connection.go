package services

import (
	"strings"
	"tssh/defs"
	"tssh/interfaces"
	"tssh/types"
	"tssh/utils"

	"github.com/spf13/viper"
)

type connection struct {
	goteleport interfaces.Goteleport
}

type Connection interface {
	ListConnections() ([]types.TsshConnection, error)
	Connect(string) error
}

func (s *connection) listConnectionsForSysadminUsers(list []string) ([]types.TsshConnection, error) {
	connections := []types.TsshConnection{}
	user := viper.GetString(defs.ConfigKeyAdminUser)
	if user == "" {
		user = "root"
	}

	for _, element := range list {
		connections = append(connections, types.TsshConnection{
			Host: element,
			User: user,
		})
	}

	return connections, nil
}

func (s *connection) listConnectionsForUsers(list []string) ([]types.TsshConnection, error) {
	connections := []types.TsshConnection{}
	for _, element := range list {
		partials := strings.Split(element, ".")
		if len(partials) < 2 {
			continue
		}

		connections = append(connections, types.TsshConnection{
			Host: partials[0],
			User: strings.Join(partials[1:], "."),
		})
	}

	return connections, nil
}

func (s *connection) ListConnections() ([]types.TsshConnection, error) {
	roles, err := s.goteleport.ListRoles()
	if err != nil {
		return nil, err
	}

	if utils.InSlice(viper.GetString(defs.ConfigKeyAdminRole), roles) {
		hosts, err := s.goteleport.ListHosts()
		if err != nil {
			return nil, err
		}

		return s.listConnectionsForSysadminUsers(hosts)
	} else {
		return s.listConnectionsForUsers(roles)
	}
}

func (s *connection) Connect(connection string) error {
	return s.goteleport.Connect(connection)
}

func NewConnectionService() Connection {
	return &connection{
		goteleport: interfaces.NewGoteleportInterface(),
	}
}
