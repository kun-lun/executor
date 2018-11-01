package commands

import (
	"github.com/kun-lun/common/errors"
	"github.com/kun-lun/common/storage"
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
