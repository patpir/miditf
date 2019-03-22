package sources

import (
	"strconv"

	"github.com/patpir/miditf/core"
	"github.com/patpir/miditf/blocks"
)

type majorScale struct {
	baseTone  uint8
}


func newMajorScale(config []blocks.Argument) (blocks.Source, error) {
	if len(config) != 1 {
		return nil, blocks.MissingArgumentError
	}

	var baseTone uint8
	switch value := config[0].Value().(type) {
	case int:
		baseTone = uint8(value)
	case string:
		baseToneInt, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		baseTone = uint8(baseToneInt)
	default:
		return nil, blocks.InvalidArgumentTypeError
	}

	scale := majorScale{}
	scale.baseTone = uint8(baseTone)
	return &scale, nil
}

func (scale *majorScale) Piece() (*core.Piece, error) {
	notes := []*core.Note{
		core.NewNote(0, scale.baseTone     , 80,   0,  16),
		core.NewNote(0, scale.baseTone +  2, 80,  16,  32),
		core.NewNote(0, scale.baseTone +  4, 80,  32,  48),
		core.NewNote(0, scale.baseTone +  5, 80,  48,  64),
		core.NewNote(0, scale.baseTone +  7, 80,  64,  80),
		core.NewNote(0, scale.baseTone +  9, 80,  80,  96),
		core.NewNote(0, scale.baseTone + 11, 80,  96, 112),
		core.NewNote(0, scale.baseTone + 12, 80, 112, 128),
	}

	track := core.NewTrack()
	for _, note := range notes {
		track.AddNote(note)
	}

	piece := core.NewPiece()
	piece.AddTrack(track)
	return piece, nil
}

func init() {
	arguments := []blocks.ArgumentInfo{
		blocks.NewArgumentInfo("base-tone", "Tone at which the major scale starts - this is the lowest tone of the scale", false),
	}
	info := blocks.NewBlockInfo("major-scale", "Creates a major scale", arguments)
	blocks.DefaultRegistrator().RegisterSource(info, newMajorScale)
}

