package commands

import (
	"github.com/kun-lun/common/errors"
	"github.com/kun-lun/common/storage"
)

type ApplyDeployment struct {
	stateStore storage.Store
}

func NewApplyDeployment(
	stateStore storage.Store,
) ApplyDeployment {
	return ApplyDeployment{
		stateStore: stateStore,
	}
}

func (p ApplyDeployment) CheckFastFails(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}

func (p ApplyDeployment) Execute(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}
