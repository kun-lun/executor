package executor

import "github.com/kun-lun/common/fileio"

type fs interface {
	fileio.Stater
	fileio.TempFiler
	fileio.FileReader
	fileio.FileWriter
}
