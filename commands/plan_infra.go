package commands

import (
	"io/ioutil"

	"github.com/kun-lun/common/storage"
	"github.com/kun-lun/infra-producer/handler"
	infra "github.com/kun-lun/infra-producer/pkg/apis"
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
	return nil
}

func (p PlanInfra) Execute(args []string, state storage.State) error {
	handlerType := handler.TerraformHandlerType // should get from args
	debug := true
	infraProducer, _ := infra.NewInfraProducer(p.stateStore, handlerType, debug)

	artifactFilePath, _ := p.stateStore.GetMainArtifactFilePath()
	b, _ := ioutil.ReadFile(artifactFilePath)

	return infraProducer.Setup(b, state)
}
