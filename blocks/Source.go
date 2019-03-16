package blocks

import (
	"errors"

	"github.com/patpir/miditf/core"
)

type Source interface {
	Piece() *core.Piece
}

type SourceFactory func([]Argument) (Source, error)

type sourceType struct {
	info     BlockInfo
	factory  SourceFactory
}

var registeredSources = make(map[string]sourceType)
var registeredSourceInfos = []BlockInfo{}

func RegisterSource(info BlockInfo, factory SourceFactory) {
	registeredSources[info.Identifier()] = sourceType{
		info:    info,
		factory: factory,
	}
	registeredSourceInfos = append(registeredSourceInfos, info)
}

func CreateSource(identifier string, argValues []Argument) (Source, error) {
	st, ok := registeredSources[identifier]
	if ok {
		return st.factory(argValues)
	}
	return nil, errors.New("Source type does not exist")
}

func Sources() []BlockInfo {
	return registeredSourceInfos
}

