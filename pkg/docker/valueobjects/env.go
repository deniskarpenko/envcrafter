package valueobjects

import "fmt"

type Env struct {
	variable string
	value    string
}

func (e *Env) Variable() string {
	return e.variable
}

func (e *Env) Value() string {
	return e.value
}

func (e *Env) ToYaml() string {
	return fmt.Sprintf("%s:%s", e.variable, e.value)
}

func NewEnv(variable string, value string) Env {
	return Env{variable: variable, value: value}
}
