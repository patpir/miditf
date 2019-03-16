package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestArgumentInfo(t *testing.T) {
	arg := NewArgumentInfo("my-arg", "Arg Description")
	assert.Equal(t, "my-arg", arg.Name())
	assert.Equal(t, "Arg Description", arg.Description())
}

