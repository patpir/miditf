package blocks

import (
	"github.com/patpir/miditf/core"
)

type Transformation interface {
	Transform(piece *core.Piece) *core.Piece
}

type TransformationFactory func([]Argument) (Transformation, error)

type transformationType struct {
	info    BlockInfo
	factory TransformationFactory
}

