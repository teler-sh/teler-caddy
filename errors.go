package teler

import "errors"

const (
	errExpectedToken      = "expected %q token"
	errInvalidFormatValue = "invalid %q format value for %q argument"
	errInvalidKey         = "invalid key %q"
	errUnsupportedFormat  = "unsupported %q format"
)

// exported
var (
	ErrNoTelerInstance = errors.New("no Teler instance initialized")
)
