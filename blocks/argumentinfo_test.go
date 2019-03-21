package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestArgumentInfo(t *testing.T) {
	arg := NewArgumentInfo("my-arg", "Arg Description", false)
	assert.Equal(t, "my-arg", arg.Name())
	assert.Equal(t, "Arg Description", arg.Description())
	assert.False(t, arg.IsOptional())
}

