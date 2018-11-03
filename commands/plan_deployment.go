package commands

import (
	"github.com/kun-lun/common/errors"
	"github.com/kun-lun/common/storage"
)

type PlanDeployment struct {
	stateStore storage.Store
}

func NewPlanDeployment(
	stateStore storage.Store,
) PlanDeployment {
	return PlanDeployment{
		stateStore: stateStore,
	}
}

func (p PlanDeployment) CheckFastFails(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}

func (p PlanDeployment) Execute(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}
