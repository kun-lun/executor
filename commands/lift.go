package commands

import (
	"github.com/xplaceholder/common/errors"
	"github.com/xplaceholder/common/storage"
)

type Lift struct {
}

func NewLift() {

}

func (p Lift) CheckFastFails(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}

func (p Lift) Execute(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}
