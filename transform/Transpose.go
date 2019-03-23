package transform

import (
	"strconv"

	"github.com/patpir/miditf/core"
	"github.com/patpir/miditf/blocks"
)

type transposer struct {
	change  uint8
}


func newTransposer(arguments map[string]interface{}) (blocks.Transformation, error) {
	changeArg, exists := arguments["change"]
	if !exists {
		return nil, blocks.MissingArgumentError
	}

	var change uint8
	switch value := changeArg.(type) {
	case int:
		change = uint8(value)
	case string:
		changeInt, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		change = uint8(changeInt)
	default:
		return nil, blocks.InvalidArgumentTypeError
	}

	tp := transposer{}
	tp.change = change
	return &tp, nil
}


func (tp *transposer) Transform(piece *core.Piece) (*core.Piece, error) {
	tracks := piece.Tracks()
	for _, track := range tracks {
		notes := track.Notes()
		for _, note := range notes {
			note.SetKey(note.Key() + tp.change)
		}
	}

	return piece, nil
}

func init() {
	arguments := []blocks.ArgumentInfo{
		blocks.NewArgumentInfo("change", "Number of semitones by which to transpose the input. Positive numbers indicate a transposition to higher notes, negative numbers indicate a transposition to lower notes.", false),
	}
	info := blocks.NewBlockInfo("transpose", "Transpose all notes by the same number of semitones.", arguments)
	blocks.DefaultRegistrator().RegisterTransformation(info, newTransposer)
}

