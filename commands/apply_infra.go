package commands

import (
	"path/filepath"

	"github.com/kun-lun/common/fileio"
	"github.com/kun-lun/common/storage"
	"github.com/kun-lun/infra-producer/handler"
	infra "github.com/kun-lun/infra-producer/pkg/apis"
)

type ApplyInfra struct {
	stateStore storage.Store
	fs         fileio.Fs
}

func NewApplyInfra(
	stateStore storage.Store,
	fs fileio.Fs,
) ApplyInfra {
	return ApplyInfra{
		stateStore: stateStore,
		fs:         fs,
	}
}

func (p ApplyInfra) CheckFastFails(args []string, state storage.State) error {
	return nil
}

func (p ApplyInfra) Execute(args []string, state storage.State) error {
	handlerType := handler.TerraformHandlerType // should get from args
	debug := true
	infraProducer, _ := infra.NewInfraProducer(p.stateStore, handlerType, debug)

	err := infraProducer.Apply(state)
	if err != nil {
		return err
	}

	contents, err := infraProducer.GetOutputs()
	if err != nil {
		return err
	}

	artifactsPatchDir, err := p.stateStore.GetArtifactsPatchDir()
	if err != nil {
		return err
	}
	outputsOpsFilePath := filepath.Join(artifactsPatchDir, "outputs.yml")

	err = p.fs.WriteFile(outputsOpsFilePath, []byte(contents), 0644)
	if err != nil {
		return err
	}

	return nil
}
