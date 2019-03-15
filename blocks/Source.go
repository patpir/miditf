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
var registeredSourceDetails = []Details{}

func RegisterSource(details Details, factory SourceFactory) {
	registeredSources[details.Identifier()] = sourceType{
		details: details,
		factory: factory,
	}
	registeredSourceDetails = append(registeredSourceDetails, details)
}

func CreateSource(identifier string, argValues []string) (Source, error) {
	st, ok := registeredSources[identifier]
	if ok {
		return st.factory(argValues)
	}
	return nil, errors.New("Source type does not exist")
}

func Sources() []Details {
	return registeredSourceDetails
}

