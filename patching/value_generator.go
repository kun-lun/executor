package patching

type ValueGenerator interface {
	Generate(interface{}) (interface{}, error)
}
