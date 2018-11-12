package commands

import (
	"github.com/kun-lun/common/fileio"
	"github.com/kun-lun/common/storage"
	deploymentProducer "github.com/kun-lun/deployment-producer/pkg/apis"
	patching "github.com/kun-lun/patching/pkg/apis"
)

type PlanDeployment struct {
	stateStore storage.Store
	fs         fileio.Fs
	logger     logger
}

func NewPlanDeployment(
	stateStore storage.Store,
	fs fileio.Fs,
	logger logger,
) PlanDeployment {
	return PlanDeployment{
		stateStore: stateStore,
		fs:         fs,
		logger:     logger,
	}
}

func (p PlanDeployment) CheckFastFails(args []string, state storage.State) error {
	return nil
}

func (p PlanDeployment) Execute(args []string, state storage.State) error {
	// load the provisioned manifest
	patching := patching.NewPatching(p.stateStore, p.fs)
	manifest, err := patching.ProvisionManifest()
	if err != nil {
		return err
	}
	deploymentProducer := deploymentProducer.NewDeploymentProducer(p.stateStore, p.logger, p.fs)
	err = deploymentProducer.Produce(*manifest)
	return err
}
