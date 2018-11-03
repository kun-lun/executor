package commands

import (
	"github.com/kun-lun/common/errors"
	"github.com/kun-lun/common/storage"
)

type PlanInfra struct {
	stateStore storage.Store
}

func NewPlanInfra(
	stateStore storage.Store,
) PlanInfra {
	return PlanInfra{
		stateStore: stateStore,
	}
}

func (p PlanInfra) CheckFastFails(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}

func (p PlanInfra) Execute(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}
