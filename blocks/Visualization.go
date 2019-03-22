package blocks

import (
	"github.com/patpir/miditf/core"
)

type Visualization interface {
	Visualize(piece *core.Piece) (string, error)
}

type VisualizationFactory func(map[string]interface{}) (Visualization, error)

type visualizationType struct {
	info    BlockInfo
	factory VisualizationFactory
}

