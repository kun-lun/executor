package commands

import (
	"github.com/xplaceholder/common/errors"
	"github.com/xplaceholder/common/storage"
)

type PlanLift struct {
}

func NewPlanLift() PlanLift {
	return PlanLift{}
}

func (p PlanLift) CheckFastFails(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}

func (p PlanLift) Execute(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}
