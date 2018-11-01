package commands

import (
	"github.com/xplaceholder/common/errors"
	"github.com/xplaceholder/common/storage"
)

type PlanDeployment struct {
}

func NewPlanDeployment() PlanDeployment {
	return PlanDeployment{}
}

func (p PlanDeployment) CheckFastFails(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}

func (p PlanDeployment) Execute(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}
