package services

import (
	"tssh/defs"
	"tssh/interfaces"
	"tssh/types"
	"tssh/utils"

	"github.com/spf13/viper"
)

type connection struct {
	sysadminRole string
	goteleport   interfaces.Goteleport
}

type Connection interface {
	ListConnections() ([]types.TsshConnection, error)
	Connect(string) error
}

func (s *connection) listConnectionsForUsers() ([]types.TsshConnection, error) {
	connections := []types.TsshConnection{}

	roles, err := s.goteleport.ListRoles()
	if err != nil {
		return connections, err
	}

	for _, role := range roles {
		for _, login := range role.Spec.Allow.Logins {
			if role.Spec.Allow.NodeLabels.Hostname == "" {
				continue
			}

			connections = append(connections, types.TsshConnection{
				Host:  role.Spec.Allow.NodeLabels.Hostname,
				User:  login,
				Refer: role.Metadata.Name,
			})
		}
	}

	return connections, nil
}

func (s *connection) listConnectionsForSysadminUsers() ([]types.TsshConnection, error) {
	connections := []types.TsshConnection{}

	hosts, err := s.goteleport.ListHosts()
	if err != nil {
		return nil, err
	}

	role, err := s.goteleport.FindRole(s.sysadminRole)
	if err != nil {
		return connections, err
	}

	for _, element := range hosts {
		for _, login := range role.Spec.Allow.Logins {
			connections = append(connections, types.TsshConnection{
				Host: element,
				User: login,
			})
		}
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
	// Take roles from auth output
	authRoles, err := s.goteleport.ListRolesForUser()
	if err != nil {
		return nil, err
	}

	// Get all connections
	connections, err := s.listConnectionsForUsers()
	if err != nil {
		return nil, err
	}

	// Exlude connections not owned by user
	var ownedConnections []types.TsshConnection
	for _, role := range connections {
		match := false
		for _, authRole := range authRoles {
			if authRole == role.Refer {
				match = true
			}
		}

		if match {
			ownedConnections = append(ownedConnections, role)
		}
	}

	// Check if the user have admin roles
	haveAdminRole := utils.InSlice(s.sysadminRole, authRoles)

	// If the user has admin permission the system must concatenate the additional connections
	if s.sysadminRole != "" && haveAdminRole {
		sysadminConnections, err := s.listConnectionsForSysadminUsers()
		if err != nil {
			return nil, err
		}

		ownedConnections = append(ownedConnections, sysadminConnections...)
	}

	return s.getUniqueCollections(ownedConnections), err
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
		goteleport:   goteleport,
		sysadminRole: viper.GetString(defs.ConfigKeyAdminRole),
	}, err
}
