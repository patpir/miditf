package blocks

import (
	"errors"
)

var MissingArgumentError error = errors.New("Non-optional argument missing for block")
var InvalidArgumentTypeError error = errors.New("Argument for block has invalid type")

