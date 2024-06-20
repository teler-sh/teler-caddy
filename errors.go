package teler

import "errors"

const (
	errInvalidFormatValue = "invalid %q format value for %q argument"
	errExpectedToken      = "expected %q token"
	errInvalidKey         = "invalid key %q"
)

// exported
var (
	ErrNoTelerInstance = errors.New("no Teler instance initialized")
)
