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
		user = defs.DefaultTSHRole
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

	adminRole := viper.GetString(defs.ConfigKeyAdminRole)
	haveAdminRole := utils.InSlice(adminRole, roles)

	sysadminConnections := []types.TsshConnection{}

	if adminRole != "" && haveAdminRole {
		hosts, err := s.goteleport.ListHosts()
		if err != nil {
			return nil, err
		}

		connections, err := s.listConnectionsForSysadminUsers(hosts)
		if err != nil {
			return connections, err
		}

		sysadminConnections = connections
	}

	normalConnections, err := s.listConnectionsForUsers(roles)

	connections := append(sysadminConnections, normalConnections...)

	uniqueConnections := []types.TsshConnection{}
	for _, connection := range connections {
		match := false

		for _, uc := range uniqueConnections {
			if uc.Host == connection.Host && uc.User == connection.User {
				match = true
			}
		}

		if !match {
			uniqueConnections = append(uniqueConnections, connection)
		}
	}

	return uniqueConnections, err
}

func (s *connection) Connect(connection string) error {
	return s.goteleport.Connect(connection)
}

func NewConnectionService(user, proxy string, passwordless bool) (Connection, error) {
	goteleport, err := interfaces.NewGoteleportInterface(
		user,
		proxy,
		passwordless,
	)

	return &connection{
		goteleport: goteleport,
	}, err
}
