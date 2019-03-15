package blocks

import "testing"


func TestDetails(t *testing.T) {
	details := NewDetails("my-name", "my-description", []Argument{})

	if details.Identifier() != "my-name" {
		t.Fail()
	}
	if details.Description() != "my-description" {
		t.Fail()
	}
}

