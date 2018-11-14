package apis

import (
	"github.com/kun-lun/common/storage"
	"github.com/kun-lun/digester/pkg/common"
	"github.com/kun-lun/digester/pkg/questionnaire"
)

func Run(state storage.State, filePath string) error {
	bp := questionnaire.Run(state, filePath)
	return bp.ExposeYaml(filePath)
}

func ImportBlueprintYaml(filePath string) (common.Blueprint, error) {
	return common.ImportBlueprintYaml(filePath)
}
