package commands

import (
	"github.com/xplaceholder/common/errors"
	"github.com/xplaceholder/common/storage"
)

type Usage struct {
	logger logger
}

func NewUsage(logger logger) Usage {
	return Usage{
		logger: logger,
	}
}

func (u Usage) CheckFastFails(subcommandFlags []string, state storage.State) error {
	return &errors.NotImplementedError{}
}

func (u Usage) Execute(subcommandFlags []string, state storage.State) error {
	return &errors.NotImplementedError{}
}
