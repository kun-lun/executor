package commands

import (
	"github.com/kun-lun/common/errors"
	"github.com/kun-lun/common/storage"
)

type ApplyInfra struct {
}

func NewApplyInfra() ApplyInfra {
	return ApplyInfra{}
}

func (p ApplyInfra) CheckFastFails(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}

func (p ApplyInfra) Execute(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}
