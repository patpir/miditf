package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNote(t *testing.T) {
	note := NewNote()
	assert.NotNil(t, note)
}

