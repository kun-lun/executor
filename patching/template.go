package patching

import (
	"github.com/cppforlife/go-patch/patch"
	yaml "gopkg.in/yaml.v2"
)

type Template struct {
	bytes []byte
}

func NewTemplate(bytes []byte) Template {
	return Template{bytes: bytes}
}

func (t Template) Evaluate(op patch.Op) ([]byte, error) {
	var obj interface{}

	err := yaml.Unmarshal(t.bytes, &obj)
	if err != nil {
		return []byte{}, err
	}

	if op != nil {
		obj, err = op.Apply(obj)
		if err != nil {
			return []byte{}, err
		}
	}

	bytes, err := yaml.Marshal(obj)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}
