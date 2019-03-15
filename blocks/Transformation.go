package blocks

import (
	"errors"

	"github.com/patpir/miditf/core"
)

type Transformation interface {
	Transform(piece *core.Piece) *core.Piece
}

type TransformationFactory func([]string) (Transformation, error)

type transformationType struct {
	details Details
	factory TransformationFactory
}

var registeredTransformations = make(map[string]transformationType)
var registeredTransformationDetails = []Details{}

func RegisterTransformation(details Details, factory TransformationFactory) {
	registeredTransformations[details.Identifier()] = transformationType{
		details: details,
		factory: factory,
	}
	registeredTransformationDetails = append(registeredTransformationDetails, details)
}

func CreateTransformation(identifier string, argValues []string) (Transformation, error) {
	tt, ok := registeredTransformations[identifier]
	if ok {
		return tt.factory(argValues)
	}
	return nil, errors.New("Transformation type does not exist")
}

func Transformations() []Details {
	return registeredTransformationDetails
}

