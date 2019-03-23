package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/patpir/miditf/core"
	"github.com/patpir/miditf/blocks"
)

func TestCreateTransposerFromFactory(t *testing.T) {
	tp, err := blocks.DefaultRegistrator().CreateTransformation("transpose", map[string]interface{}{
		"change": 2,
	})
	assert.Nil(t, err)
	assert.NotNil(t, tp)
}

func TestCreateTransposerWithoutParams(t *testing.T) {
	tp, err := blocks.DefaultRegistrator().CreateTransformation("transpose", make(map[string]interface{}))
	assert.NotNil(t, err)
	assert.Nil(t, tp)
}

func TestTransposeOutput(t *testing.T) {
	tp, err := newTransposer(map[string]interface{}{
		"change": 2,
	})
	assert.Nil(t, err)
	assert.NotNil(t, tp)

	track := core.NewTrack()
	track.AddNote(core.NewNote(0, 60, 80,  0, 16))
	track.AddNote(core.NewNote(0, 67, 80, 16, 32))
	piece := core.NewPiece()
	piece.AddTrack(track)

	tpPiece, err := tp.Transform(piece)
	assert.Nil(t, err)
	assert.NotNil(t, tpPiece)

	tracks := tpPiece.Tracks()
	assert.Equal(t, 1, len(tracks))
	tpTrack := tracks[0]
	assert.NotNil(t, tpTrack)

	notes := tpTrack.Notes()
	assert.Equal(t, 2, len(notes))

	keys := []int{ 62, 69 }
	for index, note := range notes {
		assert.Equal(t, uint8(0), note.Channel())
		assert.Equal(t, uint8(keys[index]), note.Key())
		assert.Equal(t, uint8(80), note.Velocity())
		assert.Equal(t, uint32(16 * index), note.StartTime())
		assert.Equal(t, uint32(16), note.EndTime() - note.StartTime())
	}
}

