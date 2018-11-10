package patching

import (
	"path"

	artifacts "github.com/kun-lun/artifacts/pkg/apis"
	"github.com/kun-lun/common/fileio"
	"github.com/kun-lun/common/storage"
)

type Patching struct {
	stateStore storage.Store
	fs         fileio.Fs
}

func NewPatching(
	stateStore storage.Store,
	fs fileio.Fs,
) Patching {
	return Patching{
		stateStore: stateStore,
		fs:         fs,
	}
}

func (p Patching) ProvisionManifest() (*artifacts.Manifest, error) {
	mainArtifactFilePath, err := p.stateStore.GetMainArtifactFilePath()
	if err != nil {
		return nil, err
	}
	content, err := p.fs.ReadFile(mainArtifactFilePath)
	template := NewTemplate(content)

	// construct the ops
	artifactsPatchDir, err := p.stateStore.GetArtifactsPatchDir()
	fileInfos, err := p.fs.ReadDir(artifactsPatchDir)
	opsFileArgs := []OpsFileArg{}
	for _, fileInfo := range fileInfos {
		fileArg := OpsFileArg{
			fileReader: p.fs,
		}
		fileArg.UnmarshalFlag(path.Join(artifactsPatchDir, fileInfo.Name()))
		opsFileArgs = append(opsFileArgs, fileArg)
	}

	opsFlags := OpsFlags{
		OpsFiles: opsFileArgs,
	}
	content, err = template.Evaluate(opsFlags.AsOp())
	manifest, err := artifacts.NewManifestFromYAML(content)
	return manifest, nil
}
