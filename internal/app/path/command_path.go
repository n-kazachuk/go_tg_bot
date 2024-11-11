package path

import (
	"errors"
	"fmt"
)

type CommandPath struct {
	CommandName string
}

var ErrUnknownCommand = errors.New("unknown command")

func ParseCommand(commandText string) (CommandPath, error) {
	if len(commandText) == 0 {
		return CommandPath{}, ErrUnknownCommand
	}

	return CommandPath{
		CommandName: commandText,
	}, nil
}

func (c CommandPath) WithCommandName(commandName string) CommandPath {
	c.CommandName = commandName

	return c
}

func (c CommandPath) String() string {
	return fmt.Sprintf("/%s", c.CommandName)
}
