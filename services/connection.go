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
	sysadminRole string
	sysadminUser string
	goteleport   interfaces.Goteleport
}

type Connection interface {
	ListConnections() ([]types.TsshConnection, error)
	Connect(string) error
}

func (s *connection) listConnectionsForSysadminUsers(list []string) ([]types.TsshConnection, error) {
	connections := []types.TsshConnection{}

	for _, element := range list {
		connections = append(connections, types.TsshConnection{
			Host: element,
			User: s.sysadminUser,
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

func (s *connection) getUniqueCollections(connections []types.TsshConnection) []types.TsshConnection {
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

	return uniqueConnections
}

func (s *connection) ListConnections() ([]types.TsshConnection, error) {
	roles, err := s.goteleport.ListRoles()
	if err != nil {
		return nil, err
	}

	// Check if the user have admin roles
	haveAdminRole := utils.InSlice(s.sysadminRole, roles)

	// Get normal connetions
	connections, err := s.listConnectionsForUsers(roles)
	if err != nil {
		return nil, err
	}

	// Get sysadmin connections
	if s.sysadminRole != "" && haveAdminRole {
		hosts, err := s.goteleport.ListHosts()
		if err != nil {
			return nil, err
		}

		sysadminConnections, err := s.listConnectionsForSysadminUsers(hosts)
		if err != nil {
			return nil, err
		}

		connections = append(connections, sysadminConnections...)
	}

	return s.getUniqueCollections(connections), err
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

	sysadminUser := viper.GetString(defs.ConfigKeyAdminUser)
	if sysadminUser == "" {
		sysadminUser = defs.DefaultTSHRole
	}

	return &connection{
		goteleport:   goteleport,
		sysadminRole: viper.GetString(defs.ConfigKeyAdminRole),
		sysadminUser: sysadminUser,
	}, err
}
