package presenters

import (
	"fmt"
	"tssh/types"
	"tssh/utils"
)

type connection struct{}

type Connection interface {
	Fzf(connections []types.TsshConnection) string
	Text(connections []types.TsshConnection) string
}

func (c *connection) Text(connections []types.TsshConnection) string {
	context := ""

	for _, connection := range connections {
		context += fmt.Sprintf("%s@%s\n", connection.User, connection.Host)
	}

	return context
}

func (c *connection) Fzf(connections []types.TsshConnection) string {
	context := c.Text(connections)

	return utils.ExecFZF(context)
}

func NewConnectionPresenter() Connection {
	return &connection{}
}
