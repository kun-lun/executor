package apis

import (
    "github.com/kun-lun/digester/pkg/questionnaire"
)

func Run(filePath string) error {
    bp := questionnaire.Run()
    return bp.ExposeYaml(filePath)
}
