package blocks

import (
	"errors"

	"github.com/patpir/miditf/core"
)

type Visualization interface {
	Visualize(piece *core.Piece) string
}

type VisualizationFactory func([]string) (Visualization, error)

type visualizationType struct {
	details Details
	factory VisualizationFactory
}

var registeredVisualizations = make(map[string]visualizationType)
var registeredVisualizationDetails = []Details{}

func RegisterVisualization(details Details, factory VisualizationFactory) {
	registeredVisualizations[details.Identifier()] = visualizationType{
		details: details,
		factory: factory,
	}
	registeredVisualizationDetails = append(registeredVisualizationDetails, details)
}

func CreateVisualization(identifier string, argValues []string) (Visualization, error) {
	vt, ok := registeredVisualizations[identifier]
	if ok {
		return vt.factory(argValues)
	}
	return nil, errors.New("Visualization type does not exist")
}

func Visualizations() []Details {
	return registeredVisualizationDetails
}

