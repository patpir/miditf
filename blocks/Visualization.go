package blocks

import (
	"errors"

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

var registeredVisualizations = make(map[string]visualizationType)
var registeredVisualizationInfos = []BlockInfo{}

func RegisterVisualization(info BlockInfo, factory VisualizationFactory) {
	registeredVisualizations[info.Identifier()] = visualizationType{
		info:    info,
		factory: factory,
	}
	registeredVisualizationInfos = append(registeredVisualizationInfos, info)
}

func CreateVisualization(identifier string, argValues []Argument) (Visualization, error) {
	vt, ok := registeredVisualizations[identifier]
	if ok {
		return vt.factory(argValues)
	}
	return nil, errors.New("Visualization type does not exist")
}

func Visualizations() []BlockInfo {
	return registeredVisualizationInfos
}

