package commands

import (
	"github.com/kun-lun/common/errors"
	"github.com/kun-lun/common/storage"
)

type ApplyInfra struct {
	stateStore storage.Store
}

func NewApplyInfra(
	stateStore storage.Store,
) ApplyInfra {
	return ApplyInfra{
		stateStore: stateStore,
	}
}

func (p ApplyInfra) CheckFastFails(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}

func (p ApplyInfra) Execute(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}
