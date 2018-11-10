package patching

type ValueGeneratorFactory interface {
	GetGenerator(valueType string) (ValueGenerator, error)
}
