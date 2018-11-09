package detector

import (
    "io/ioutil"
    "log"
    "os"
    "strings"
    "github.com/kun-lun/digester/pkg/common"
    "github.com/kun-lun/digester/pkg/artifactgen"
    nullFramework "github.com/kun-lun/digester/pkg/detector/frameworks/null"
    nullPackageManager "github.com/kun-lun/digester/pkg/detector/packagemanagers/null"
)

type Detector struct {
    projectPath string
    packageManager common.PackageManager
    framework common.Framework
    blueprint artifactgen.Blueprint
}

func New(projectPath string) (*Detector, error) {
    if !strings.HasSuffix(projectPath, string(os.PathSeparator)) {
        projectPath += string(os.PathSeparator)
    }
    _, err := ioutil.ReadDir(projectPath)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    return &Detector{
        projectPath: projectPath,
        blueprint: artifactgen.Blueprint{
            NonIaaS: artifactgen.NonIaaS{
                ProjectPath: projectPath,
            },
        },
    }, nil
}

func (d *Detector) DetectPackageManager() []common.PackageManagerName {
    packageManagers := getPackageManagers()
    possiblePackageManagers := []common.PackageManagerName{}
    for _, pm := range packageManagers {
        if pm.Identify(d.projectPath) {
            possiblePackageManagers = append(possiblePackageManagers, pm.GetName())
        }
    }
    return possiblePackageManagers
}

func (d *Detector) ConfirmPackageManager(pmn string) {
    packageManagers := getPackageManagers()
    pm, ok := packageManagers[pmn]
    if (ok) {
        d.packageManager = pm
    } else {
        d.packageManager = nullPackageManager.New()
    }
}

func (d *Detector) DetectFramework() []common.FrameworkName {
    return d.packageManager.DetectFramework(d.projectPath)
}

func (d *Detector) ConfirmFramework(fwn string) {
    frameworks := getFrameworks()
    fw, ok := frameworks[fwn]
    if (ok) {
        d.framework = fw
    } else {
        d.framework = nullFramework.New()
    }
    d.blueprint.NonIaaS.ProgrammingLanguage = d.framework.GetProgrammingLanguage()
}

// Only one database in an array for now
func (d *Detector) DetectConfig() {
    d.blueprint.NonIaaS.Databases = d.framework.DetectConfig(d.projectPath)
}

func (d *Detector) ExposeKnownInfo() artifactgen.Blueprint {
    return d.blueprint
}
