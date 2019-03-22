package visualize

import (
	"testing"
	"io/ioutil"

	"github.com/stretchr/testify/assert"

	"github.com/patpir/miditf/core"
	"github.com/patpir/miditf/blocks"
)


func TestCreateMidiWriterFromFactory(t *testing.T) {
	writer, err := blocks.DefaultRegistrator().CreateVisualization("midi-file", make(map[string]interface{}))
	assert.NotNil(t, writer)
	assert.Nil(t, err)
}

func TestMidiWriterOutput(t *testing.T) {
	track := core.NewTrack()
	track.AddNote(core.NewNote(0, 60, 80,   0,  16))
	track.AddNote(core.NewNote(0, 62, 80,  16,  32))
	track.AddNote(core.NewNote(0, 64, 80,  32,  48))
	track.AddNote(core.NewNote(0, 65, 80,  48,  64))
	track.AddNote(core.NewNote(0, 67, 80,  64,  80))
	track.AddNote(core.NewNote(0, 69, 80,  80,  96))
	track.AddNote(core.NewNote(0, 71, 80,  96, 112))
	track.AddNote(core.NewNote(0, 72, 80, 112, 128))

	piece := core.NewPiece()
	piece.AddTrack(track)

	writer, err := newMidiFileWriter(make(map[string]interface{}))
	assert.NotNil(t, writer)
	assert.Nil(t, err)

	result, err := writer.Visualize(piece)
	assert.Nil(t, err)

	midiFile, err := ioutil.ReadFile("../testdata/c-major.midi")
	assert.Nil(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, string(midiFile), result)
}

