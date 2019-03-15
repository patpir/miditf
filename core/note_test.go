package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNote(t *testing.T) {
	note := NewNote(1, 65, 80, 16, 31)
	assert.NotNil(t, note)
	assert.Equal(t, uint8(1), note.Channel())
	assert.Equal(t, uint8(65), note.Key())
	assert.Equal(t, uint8(80), note.Velocity())
	assert.Equal(t, uint32(16), note.StartTime())
	assert.Equal(t, uint32(31), note.EndTime())
}

