package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestBlock(t *testing.T) {
	block := NewBlock("my-name", "My Comment", []Argument{})
	assert.Equal(t, "my-name", block.TypeId())
	assert.Equal(t, "My Comment", block.Comment())
}

