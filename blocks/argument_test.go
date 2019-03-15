package blocks

import "testing"


func TestArgument(t *testing.T) {
	arg := NewArgument("my-arg", "arg-description")
	if arg.Name() != "my-arg" {
		t.Fail()
	}
	if arg.Description() != "arg-description" {
		t.Fail()
	}
}

