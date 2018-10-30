package commands

type logger interface {
	Step(string, ...interface{})
	Printf(string, ...interface{})
	Println(string)
	Prompt(string) bool
}
