package teler

import (
	"fmt"
	"strings"

	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m Middleware

	err := m.UnmarshalCaddyfile(h.Dispenser)

	return m, err
}

// validateCfgFormat to validates the teler WAF config format in used.
func validateCfgFormat(format, token string) (string, error) {
	validCfgFormats := map[string]bool{json: true, yaml: true}
	formatLowercased := strings.ToLower(format)

	if !validCfgFormats[formatLowercased] {
		return "", fmt.Errorf(errInvalidFormatValue, format, token)
	}

	return formatLowercased, nil
}
