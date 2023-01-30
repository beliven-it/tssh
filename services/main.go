package services

import (
	"strings"
	"tssh/interfaces"
	"tssh/types"
)

type connection struct {
	goteleport interfaces.Goteleport
}

type Connection interface {
	ListConnections() ([]types.TsshConnection, error)
	Connect(string) error
}

func (s *connection) ListConnections() ([]types.TsshConnection, error) {
	roles, err := s.goteleport.ListRoles()
	if err != nil {
		return nil, err
	}

	connections := []types.TsshConnection{}
	for _, role := range roles {
		partials := strings.Split(role, ".")
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

func (s *connection) Connect(connection string) error {
	return s.goteleport.Connect(connection)
}

func NewConnectionService() Connection {
	return &connection{
		goteleport: interfaces.NewGoteleportInterface(),
	}
}
