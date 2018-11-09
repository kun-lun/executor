package artifactgen

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"

    "github.com/kun-lun/digester/pkg/common"
)

type NonIaaS struct {
    ProgrammingLanguage common.ProgrammingLanguage `json:ProgrammingLanguageXX`
    Databases []common.Database
}

type IaaS struct {
    VMGroup common.VMGroup
}

type Blueprint struct {
    NonIaaS NonIaaS
    IaaS    IaaS
}

// TODO check if it fits into one of the artifacts templates
func (b Blueprint) finalValidation() error {
    return nil
}

func (b Blueprint) ExposeYaml(filePath string) error {
    if err := b.finalValidation(); err != nil {
        return err
    }
    bpBytes, _ := yaml.Marshal(b)
    return ioutil.WriteFile(filePath, bpBytes, 0644)
}

func (b Blueprint) ImportYaml(filePath string) error {
    bp := Blueprint{}
    bpBytes, err := ioutil.ReadFile(filePath)
    if err != nil {
        return err
    }
    return yaml.Unmarshal(bpBytes, &bp)
}

func (b Blueprint) ExposeArtifects(filePath string) error {
    if err := b.finalValidation(); err != nil {
        return err
    }
    return nil
}
