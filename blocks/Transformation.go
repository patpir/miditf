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
	info    BlockInfo
	factory TransformationFactory
}

var registeredTransformations = make(map[string]transformationType)
var registeredTransformationInfos = []BlockInfo{}

func RegisterTransformation(info BlockInfo, factory TransformationFactory) {
	registeredTransformations[info.Identifier()] = transformationType{
		info:    info,
		factory: factory,
	}
	registeredTransformationInfos = append(registeredTransformationInfos, info)
}

func CreateTransformation(identifier string, argValues []string) (Transformation, error) {
	tt, ok := registeredTransformations[identifier]
	if ok {
		return tt.factory(argValues)
	}
	return nil, errors.New("Transformation type does not exist")
}

func Transformations() []BlockInfo {
	return registeredTransformationInfos
}

