package commands

import (
	"github.com/kun-lun/common/errors"
	"github.com/kun-lun/common/storage"
)

type PlanInfra struct {
}

func NewPlanInfra() PlanInfra {
	return PlanInfra{}
}

func (p PlanInfra) CheckFastFails(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}

func (p PlanInfra) Execute(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}
