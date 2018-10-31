package executor

import (
	"fmt"

	"github.com/xplaceholder/common/configuration"
	"github.com/xplaceholder/common/errors"
	"github.com/xplaceholder/executor/commands"
)

type usage interface {
	Print()
	PrintCommandUsage(command, message string)
}

type App struct {
	commands      commands.CommandSet
	configuration configuration.Configuration
	usage         usage
}

func New(configuration configuration.Configuration, usage usage) App {
	return App{
		commands:      getBuildInCommands(),
		configuration: configuration,
		usage:         usage,
	}
}

func getBuildInCommands() commands.CommandSet {
	commandSet := commands.CommandSet{}
	return commandSet
}

func (a App) Run() error {
	err := a.execute()
	return err
}

func (a App) getCommand(commandString string) (commands.Command, error) {
	command, ok := a.commands[commandString]
	if !ok {
		a.usage.Print()
		return nil, fmt.Errorf("unknown command: %s", commandString)
	}
	return command, nil
}

func (a App) execute() error {
	command, err := a.getCommand(a.configuration.Command)
	if err != nil {
		return err
	}

	if a.configuration.ShowCommandHelp {
		a.usage.PrintCommandUsage(a.configuration.Command, command.Usage())
		return nil
	}
	return &errors.NotImplementedError{}
}
