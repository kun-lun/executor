package commands

import (
	"github.com/xplaceholder/common/errors"
	"github.com/xplaceholder/common/storage"
)

type PlanShift struct {
}

func NewPlanShift() PlanShift {
	return PlanShift{}
}

func (p PlanShift) CheckFastFails(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}

func (p PlanShift) Execute(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}
