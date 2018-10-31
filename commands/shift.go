package commands

import (
	"github.com/xplaceholder/common/errors"
	"github.com/xplaceholder/common/storage"
)

type Shift struct {
}

func NewShift() Shift {
	return Shift{}
}

func (p Shift) CheckFastFails(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}

func (p Shift) Execute(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}
