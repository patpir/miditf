package blocks

import (
	"github.com/patpir/miditf/core"
)

type Source interface {
	Piece() (*core.Piece, error)
}

type SourceFactory func([]Argument) (Source, error)

type sourceType struct {
	info     BlockInfo
	factory  SourceFactory
}

