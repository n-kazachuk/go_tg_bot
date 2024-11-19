package path

import (
	"errors"
	"fmt"
)

type CommandPath struct {
	CommandName string
}

var ErrUnknownCommand = errors.New("unknown command")

func NewCommandPath(commandName string) *CommandPath {
	return &CommandPath{
		CommandName: commandName,
	}
}

func ParseCommand(commandText string) (*CommandPath, error) {
	if len(commandText) == 0 {
		return nil, ErrUnknownCommand
	}

	return &CommandPath{
		CommandName: commandText,
	}, nil
}

func (c CommandPath) String() string {
	return fmt.Sprintf("/%s", c.CommandName)
}
