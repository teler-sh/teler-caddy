package teler

import "github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"

// UnmarshalCaddyfile implements caddyfile.Unmarshaler.
func (m *Middleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	var cfgFormat, cfgPath, cfgInline string

	if !d.Next() { // consume directive name
		return d.Errf(errExpectedToken, dir)
	}

	for d.NextBlock(0) {
		key := d.Val()
		var target *string

		switch key {
		case keyLoadFrom:
			target = &cfgPath
		case keyInline:
			target = &cfgInline
		default:
			return d.Errf(errInvalidKey, key)
		}

		if !d.Args(&cfgFormat, target) {
			return d.ArgErr()
		}

		format, err := validateCfgFormat(cfgFormat, key)
		if err != nil {
			return d.WrapErr(err)
		}

		if d.NextArg() {
			// too many args
			return d.ArgErr()
		}

		m.Format = format

		switch key {
		case keyLoadFrom:
			m.LoadFrom = cfgPath
		case keyInline:
			m.Inline = cfgInline
		}
	}

	return nil
}
