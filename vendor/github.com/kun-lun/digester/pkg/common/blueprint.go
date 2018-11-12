package common

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "reflect"
)

type NonInfra struct {
    ProjectPath         string
    ProgrammingLanguage ProgrammingLanguage
    Databases           []Database
}

type Infra struct {
    Size InfraSize
}

type Blueprint struct {
    NonInfra NonInfra
    Infra    Infra
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
    bpfy := blueprintForYaml{
        ProjectPath: b.NonInfra.ProjectPath,
        ProgrammingLanguage: string(b.NonInfra.ProgrammingLanguage),
        VMGroupSize: string(b.Infra.Size),
    }
    // TODO support more. Assume at most one database for now.
    if len(b.NonInfra.Databases) > 0 {
        bpfy.DatabaseDriver = b.NonInfra.Databases[0].Driver
        bpfy.DatabaseVersion = b.NonInfra.Databases[0].Version
        bpfy.DatabaseStorage = b.NonInfra.Databases[0].Storage
        bpfy.DatabaseOriginHost = b.NonInfra.Databases[0].OriginHost
        bpfy.DatabaseOriginName = b.NonInfra.Databases[0].OriginName
        bpfy.DatabaseOriginUsername = b.NonInfra.Databases[0].OriginUsername
        bpfy.DatabaseOriginPassword = b.NonInfra.Databases[0].OriginPassword
    }
    bpBytes, _ := yaml.Marshal(bpfy)
    return ioutil.WriteFile(filePath, bpBytes, 0644)
}

func ImportBlueprintYaml(filePath string) (Blueprint, error) {
    bp := Blueprint{}
    bpfy := blueprintForYaml{}
    bpBytes, err := ioutil.ReadFile(filePath)
    if err != nil {
        return bp, err
    }
    if err = yaml.Unmarshal(bpBytes, &bpfy); err != nil {
        return bp, err
    }

    bp.Infra = Infra{
    }
    bp.Infra.Size, err = ParseInfraSize(bpfy.VMGroupSize)
    if err != nil {
        return bp, err
    }

    bp.NonInfra = NonInfra{
        ProjectPath: bpfy.ProjectPath,
    }

    bp.NonInfra.ProgrammingLanguage, err =
        ParseProgrammingLanguage(bpfy.ProgrammingLanguage)
    if err != nil {
        return bp, err
    }

    // TODO support more. Assume at most one database for now.
    db := Database{
        Driver: bpfy.DatabaseDriver,
        Version: bpfy.DatabaseVersion,
        Storage: bpfy.DatabaseStorage,
        OriginHost: bpfy.DatabaseOriginHost,
        OriginName: bpfy.DatabaseOriginName,
        OriginUsername: bpfy.DatabaseOriginUsername,
        OriginPassword: bpfy.DatabaseOriginPassword,
    }
    allEmpty := true
    s := reflect.ValueOf(&db).Elem()
    for i := 0; i < s.NumField(); i++ {
        valField := s.Field(i)
        val := valField.Interface().(string)
        if (val != "") {
            allEmpty = false
        }
    }
    if !allEmpty {
        bp.NonInfra.Databases = append(bp.NonInfra.Databases, db)
    }

    if err = bp.finalValidation(); err != nil {
        return bp, err
    }

    return bp, nil
}
