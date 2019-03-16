package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTrack(t *testing.T) {
	track := NewTrack()
	assert.NotNil(t, track)

	assert.Equal(t, 0, len(track.Notes()))
	note := NewNote(0, 60, 80, 0, 15)
	track.AddNote(note)
	assert.Equal(t, 1, len(track.Notes()))
}

