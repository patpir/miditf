package sources

import (
	"errors"
	"strconv"

	"github.com/patpir/miditf/core"
	"github.com/patpir/miditf/blocks"
)

type majorScale struct {
	baseTone  uint8
}


func newMajorScale(config []blocks.Argument) (blocks.Source, error) {
	if len(config) != 1 {
		return nil, errors.New("Invalid number of arguments")
	}

	baseTone, err := strconv.Atoi(config[0].Value())
	if err != nil {
		return nil, err
	}

	scale := majorScale{}
	scale.baseTone = uint8(baseTone)
	return &scale, nil
}

func (scale *majorScale) Piece() *core.Piece {
	notes := []*core.Note{
		core.NewNote(0, scale.baseTone     , 80,   0,  15),
		core.NewNote(0, scale.baseTone +  2, 80,  16,  31),
		core.NewNote(0, scale.baseTone +  4, 80,  32,  47),
		core.NewNote(0, scale.baseTone +  5, 80,  48,  63),
		core.NewNote(0, scale.baseTone +  7, 80,  64,  79),
		core.NewNote(0, scale.baseTone +  9, 80,  80,  95),
		core.NewNote(0, scale.baseTone + 11, 80,  96, 111),
		core.NewNote(0, scale.baseTone + 12, 80, 112, 127),
	}

	track := core.NewTrack()
	for _, note := range notes {
		track.AddNote(note)
	}

	piece := core.NewPiece()
	piece.AddTrack(track)
	return piece
}

func init() {
	arguments := []blocks.ArgumentInfo{
		blocks.NewArgumentInfo("base-tone", "Tone at which the major scale starts - this is the lowest tone of the scale"),
	}
	info := blocks.NewBlockInfo("major-scale", "Creates a major scale", arguments)
	blocks.DefaultRegistrator().RegisterSource(info, newMajorScale)
}

