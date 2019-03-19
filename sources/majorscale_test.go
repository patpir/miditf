package sources

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/patpir/miditf/blocks"
)

func TestCreateMajorScaleSourceFromFactory(t *testing.T) {
	// MIDI key 60 stands fo C4
	src, err := blocks.DefaultRegistrator().CreateSource("major-scale", []blocks.Argument{ blocks.NewArgument("base-tone", "60") })
	assert.Nil(t, err)
	assert.NotNil(t, src)
}

func TestCreateMajorScaleSourceWihoutParams(t *testing.T) {
	src, err := blocks.DefaultRegistrator().CreateSource("major-scale", []blocks.Argument{})
	assert.NotNil(t, err)
	assert.Nil(t, src)
}

func TestMajorScaleNotes(t *testing.T) {
	src, err := blocks.DefaultRegistrator().CreateSource("major-scale", []blocks.Argument{ blocks.NewArgument("base-tone", "60") })
	assert.Nil(t, err)
	assert.NotNil(t, src)

	piece := src.Piece()
	assert.NotNil(t, piece)

	tracks := piece.Tracks()
	assert.Equal(t, 1, len(tracks))
	track := tracks[0]
	assert.NotNil(t, track)

	notes := track.Notes()
	assert.Equal(t, 8, len(notes))

	keys := []int{ 60, 62, 64, 65, 67, 69, 71, 72 }
	for index, note := range notes {
		assert.Equal(t, uint8(keys[index]), note.Key())
		assert.Equal(t, uint32(16 * index), note.StartTime())
		assert.Equal(t, uint32(15), note.EndTime() - note.StartTime())
	}
}

