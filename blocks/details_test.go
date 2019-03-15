package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestDetails(t *testing.T) {
	details := NewDetails("my-name", "my-description", []Argument{})
	assert.Equal(t, "my-name", details.Identifier())
	assert.Equal(t, "my-description", details.Description())
}

