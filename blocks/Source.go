package blocks

import (
	"errors"

	"github.com/patpir/miditf/core"
)

type Source interface {
	Piece() *core.Piece
}

type SourceFactory func([]string) (Source, error)

type sourceType struct {
	details  Details
	factory  SourceFactory
}

var registeredSources = make(map[string]sourceType)

func RegisterSource(details Details, factory SourceFactory) {
	registeredSources[details.Identifier()] = sourceType{
		details: details,
		factory: factory,
	}
}

func CreateSource(identifier string, argValues []string) (Source, error) {
	st, ok := registeredSources[identifier]
	if ok {
		return st.factory(argValues)
	}
	return nil, errors.New("Source type does not exist")
}

