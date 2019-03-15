package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTrack(t *testing.T) {
	track := NewTrack()
	assert.NotNil(t, track)
}

