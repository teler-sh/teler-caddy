package teler

import (
	"fmt"

	"github.com/teler-sh/teler-waf"
	"github.com/teler-sh/teler-waf/option"
)

// Loader is a function type that loads teler WAF options from a file path or
// string.
type Loader func(string) (teler.Options, error)

// getTelerOptions based on the Caddyfile definition.
func getTelerOptions(m *Middleware) (teler.Options, error) {
	var loader Loader
	var opt teler.Options

	if m.LoadFrom != "" {
		switch m.Format {
		case json:
			loader = option.LoadFromJSONFile
		case yaml:
			loader = option.LoadFromYAMLFile
		default:
			return opt, fmt.Errorf(errUnsupportedFormat, m.Format)
		}

		return loader(m.LoadFrom)
	}

	if m.Inline != "" {
		switch m.Format {
		case json:
			loader = option.LoadFromJSONString
		case yaml:
			loader = option.LoadFromYAMLString
		default:
			return opt, fmt.Errorf(errUnsupportedFormat, m.Format)
		}

		return loader(m.Inline)
	}

	return opt, nil
}
