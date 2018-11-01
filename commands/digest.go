package commands

import (
	"github.com/kun-lun/common/errors"
	"github.com/kun-lun/common/storage"
)

type Digest struct {
}

func NewDigest() Digest {
	return Digest{}
}

func (p Digest) CheckFastFails(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}

func (p Digest) Execute(args []string, state storage.State) error {
	return &errors.NotImplementedError{}
}
