package commands

import (
	"github.com/xplaceholder/common/errors"
	"github.com/xplaceholder/common/storage"
)

type ApplyInfra struct {
}

func NewApplyInfra() {

}

func (p ApplyInfra) CheckFastFails(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}

func (p ApplyInfra) Execute(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}
