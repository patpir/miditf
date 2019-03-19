package blocks

import (
	"github.com/patpir/miditf/core"
)

type Visualization interface {
	Visualize(piece *core.Piece) string
}

type VisualizationFactory func([]Argument) (Visualization, error)

type visualizationType struct {
	info    BlockInfo
	factory VisualizationFactory
}

