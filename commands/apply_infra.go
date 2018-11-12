package commands

import (
	"github.com/kun-lun/common/storage"
	"github.com/kun-lun/infra-producer/handler"
	infra "github.com/kun-lun/infra-producer/pkg/apis"
)

type ApplyInfra struct {
	stateStore storage.Store
}

func NewApplyInfra(
	stateStore storage.Store,
) ApplyInfra {
	return ApplyInfra{
		stateStore: stateStore,
	}
}

func (p ApplyInfra) CheckFastFails(args []string, state storage.State) error {
	return nil
}

func (p ApplyInfra) Execute(args []string, state storage.State) error {
	handlerType := handler.TerraformHandlerType // should get from args
	debug := true
	infraProducer, _ := infra.NewInfraProducer(p.stateStore, handlerType, debug)
	return infraProducer.Apply(state)
}
