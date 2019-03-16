package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestBlockInfo(t *testing.T) {
	info := NewBlockInfo("my-name", "My Description", []ArgumentInfo{})
	assert.Equal(t, "my-name", info.Identifier())
	assert.Equal(t, "My Description", info.Description())
}

