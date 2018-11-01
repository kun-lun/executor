package commands

import (
	"github.com/xplaceholder/common/errors"
	"github.com/xplaceholder/common/storage"
)

type ApplyDeployment struct {
}

func NewApplyDeployment() ApplyDeployment {
	return ApplyDeployment{}
}

func (p ApplyDeployment) CheckFastFails(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}

func (p ApplyDeployment) Execute(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}
