package commands

import (
	"github.com/xplaceholder/common/errors"
	"github.com/xplaceholder/common/storage"
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
