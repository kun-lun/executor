package commands

import (
	"github.com/xplaceholder/common/errors"
	"github.com/xplaceholder/common/storage"
)

type Promote struct {
}

func NewPromote() Promote {
	return Promote{}
}

func (p Promote) CheckFastFails(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}

func (p Promote) Execute(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}
