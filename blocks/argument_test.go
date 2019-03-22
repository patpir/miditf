package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestArgument(t *testing.T) {
	arg := NewArgument("my-arg", "arg-value")
	assert.Equal(t, "my-arg", arg.Name())
	_, ok := arg.Value().(string)
	assert.True(t, ok)
	assert.Equal(t, "arg-value", arg.Value())
}

func TestArgumentNonStringValue(t *testing.T) {
	arg := NewArgument("my-int", 1234)
	assert.Equal(t, "my-int", arg.Name())
	_, ok := arg.Value().(int)
	assert.True(t, ok)
	assert.Equal(t, 1234, arg.Value())
}

