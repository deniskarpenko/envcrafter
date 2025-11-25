package valueobjects

import (
	"fmt"
	"testing"
)

func TestNewEnv(t *testing.T) {
	variable := "variable"
	value := "value"

	e := NewEnv(variable, value)

	if e.value != value {
		t.Errorf("Expected %s, got %s", value, e.value)
	}

	if e.variable != variable {
		t.Errorf("Expected %s, got %s", variable, e.variable)
	}

	yaml := fmt.Sprintf("%s:%s", variable, value)

	if e.ToYaml() != yaml {
		t.Errorf("Expected %s, got %s", yaml, e.ToYaml())
	}
}
