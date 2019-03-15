package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePiece(t *testing.T) {
	piece := NewPiece()
	assert.NotNil(t, piece)
}

