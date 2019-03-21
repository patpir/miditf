package sources

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/patpir/miditf/blocks"
)

func TestCreateMidiSourceFromFactory(t *testing.T) {
	reader, err := blocks.DefaultRegistrator().CreateSource("midi-file", []blocks.Argument{ blocks.NewArgument("filename", "") })
	assert.NotNil(t, reader)
	assert.Nil(t, err)
}

func TestCreateMidiSourceFromFactoryWithoutParams(t *testing.T) {
	reader, err := blocks.DefaultRegistrator().CreateSource("midi-file", []blocks.Argument{})
	assert.Nil(t, reader)
	assert.NotNil(t, err)
}

func TestReadFromFile(t *testing.T) {
	reader, err := blocks.DefaultRegistrator().CreateSource("midi-file", []blocks.Argument{
		blocks.NewArgument("filename", "../testdata/c-major.midi"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, reader)

	piece, err := reader.Piece()
	assert.Nil(t, err)
	if err != nil {
		assert.Equal(t, "", err.Error())
	}
	assert.NotNil(t, piece)

	tracks := piece.Tracks()
	assert.Equal(t, 1, len(tracks))

	notes := tracks[0].Notes()
	assert.Equal(t, 8, len(notes))

	keys := []uint8{ 60, 62, 64, 65, 67, 69, 71, 72 }
	for i, note := range notes {
		assert.Equal(t, keys[i], note.Key())
		assert.Equal(t, uint8(0), note.Channel())
		assert.Equal(t, uint8(80), note.Velocity())
		assert.Equal(t, uint32(16), note.EndTime() - note.StartTime())
	}
}

