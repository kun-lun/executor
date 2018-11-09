package artifactgen

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"

    "github.com/kun-lun/digester/pkg/common"
)

type NonIaaS struct {
    ProjectPath         string
    ProgrammingLanguage common.ProgrammingLanguage
    Databases           []common.Database
}

type IaaS struct {
    VMGroup common.VMGroup
}

type Blueprint struct {
    NonIaaS NonIaaS
    IaaS    IaaS
}

type blueprintForYaml struct {
    ProjectPath             string `yaml:"project_path,omitempty"`
    ProgrammingLanguage     string `yaml:"programming_language,omitempty"`
    DatabaseDriver          string `yaml:"database_driver,omitempty"`
    DatabaseVersion         string `yaml:"database_version,omitempty"`
    DatabaseStorage         string `yaml:"database_storage,omitempty"`
    DatabaseOriginHost      string `yaml:"database_origin_host,omitempty"`
    DatabaseOriginName      string `yaml:"database_origin_name,omitempty"`
    DatabaseOriginUsername  string `yaml:"database_origin_username,omitempty"`
    DatabaseOriginPassword  string `yaml:"database_origin_password,omitempty"`
    VMGroupCount            int    `yaml:"vmgroup_count,omitempty"`
    VMGroupSize             string `yaml:"vmgroup_size,omitempty"`
}

// TODO check if it fits into one of the artifacts templates
func (b Blueprint) finalValidation() error {
    return nil
}

func (b Blueprint) ExposeYaml(filePath string) error {
    if err := b.finalValidation(); err != nil {
        return err
    }
    // assume at most one database for now
    bpfy := blueprintForYaml{
        ProjectPath: b.NonIaaS.ProjectPath,
        ProgrammingLanguage: string(b.NonIaaS.ProgrammingLanguage),
        VMGroupCount: b.IaaS.VMGroup.Count,
        VMGroupSize: b.IaaS.VMGroup.Size,
    }
    if len(b.NonIaaS.Databases) > 0 {
        bpfy.DatabaseDriver = b.NonIaaS.Databases[0].Driver
        bpfy.DatabaseVersion = b.NonIaaS.Databases[0].Version
        bpfy.DatabaseStorage = b.NonIaaS.Databases[0].Storage
        bpfy.DatabaseOriginHost = b.NonIaaS.Databases[0].OriginHost
        bpfy.DatabaseOriginName = b.NonIaaS.Databases[0].OriginName
        bpfy.DatabaseOriginUsername = b.NonIaaS.Databases[0].OriginUsername
        bpfy.DatabaseOriginPassword = b.NonIaaS.Databases[0].OriginPassword
    }
    bpBytes, _ := yaml.Marshal(bpfy)
    return ioutil.WriteFile(filePath, bpBytes, 0644)
}

func (b Blueprint) ImportYaml(filePath string) error {
    /*
    bpfy := blueprintForYaml{}
    bpBytes, err := ioutil.ReadFile(filePath)
    if err != nil {
        return err
    }
    if err = yaml.Unmarshal(bpBytes, &bpfy); err != nil {
        return err
    }
    */
    return nil
}

func (b Blueprint) ExposeArtifects(filePath string) error {
    if err := b.finalValidation(); err != nil {
        return err
    }
    return nil
}
