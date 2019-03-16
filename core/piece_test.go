package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPiece(t *testing.T) {
	piece := NewPiece()
	assert.NotNil(t, piece)

	assert.Equal(t, 0, len(piece.Tracks()))
	track := NewTrack()
	piece.AddTrack(track)
	assert.Equal(t, 1, len(piece.Tracks()))
}

