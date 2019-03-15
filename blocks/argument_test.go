package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestArgument(t *testing.T) {
	arg := NewArgument("my-arg", "arg-description")
	assert.Equal(t, "my-arg", arg.Name())
	assert.Equal(t, "arg-description", arg.Description())
}

