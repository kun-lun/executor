package commands

import (
	"os/exec"

	"github.com/kun-lun/common/storage"
)

type ApplyDeployment struct {
	stateStore storage.Store
}

func NewApplyDeployment(
	stateStore storage.Store,
) ApplyDeployment {
	return ApplyDeployment{
		stateStore: stateStore,
	}
}

func (p ApplyDeployment) CheckFastFails(args []string, state storage.State) error {
	// return &errors.NotImplementedError{}
	return nil
}

func (p ApplyDeployment) Execute(args []string, state storage.State) error {
	// run

	deploymentScriptFilePath, err := p.stateStore.GetDeploymentScriptFile()
	if err != nil {
		return err
	}

	command := exec.Command(deploymentScriptFilePath)
	// command.Dir = workingDirectory

	// command.Env = os.Environ()
	// command.Env = append(command.Env, extraEnvVars...)

	// command.Stdout = io.MultiWriter(stdout, c.outputBuffer)
	// command.Stderr = c.errorBuffer

	return command.Run()
	// return &errors.NotImplementedError{}
}
