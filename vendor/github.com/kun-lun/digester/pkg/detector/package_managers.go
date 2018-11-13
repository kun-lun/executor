package detector

import (
	"github.com/kun-lun/digester/pkg/common"
	"github.com/kun-lun/digester/pkg/detector/packagemanagers/composer"
)

const (
	UnknownPackageManager common.PackageManagerName = "unknown"
)

var Composer common.PackageManagerName = composer.New().GetName()

func getPackageManagers() map[string]common.PackageManager {
	return map[string]common.PackageManager{
		string(Composer): composer.New(),
	}
}
