package executor

import (
	"fmt"

	"github.com/xplaceholder/common/configuration"
	"github.com/xplaceholder/common/logger"
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
	logger        logger.Logger
}

func New(configuration configuration.Configuration, usage usage, logger *logger.Logger) App {
	commandSet := commands.CommandSet{}
	commandSet["help"] = commands.NewUsage(logger)
	commandSet["digest"] = commands.NewDigest()
	commandSet["plan_lift"] = commands.NewPlanLift()
	commandSet["lift"] = commands.NewShift()
	commandSet["plan_shift"] = commands.NewPlanShift()
	commandSet["shift"] = commands.NewShift()
	return App{
		commands:      commandSet,
		configuration: configuration,
		usage:         usage,
	}
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

	if a.configuration.Command == "help" && len(a.configuration.SubcommandFlags) != 0 {
		commandString := a.configuration.SubcommandFlags[0]
		command, err = a.getCommand(commandString)
		if err != nil {
			return err
		}
		a.usage.PrintCommandUsage(commandString, command.Usage())
		return nil
	}

	err = command.CheckFastFails(a.configuration.SubcommandFlags, a.configuration.State)
	if err != nil {
		return err
	}

	return command.Execute(a.configuration.SubcommandFlags, a.configuration.State)
}
