package commands

import (
	"github.com/xplaceholder/common/errors"
	"github.com/xplaceholder/common/storage"
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
