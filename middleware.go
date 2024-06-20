package teler

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/teler-sh/teler-waf"
)

// Middleware integrates the robust security features of teler WAF into the
// Caddy web server, ensuring your web servers remain secure and resilient
// against web-based attacks.
type Middleware struct {
	// Options holds the settings for teler WAF.
	teler.Options `json:"-"`

	// Format is the type of configuration file, either "json" or "yaml".
	Format string `json:"format"`

	// LoadFrom is the path to the configuration file.
	LoadFrom string `json:"load_from"`

	// Inline is the configuration options written directly as a string.
	Inline string `json:"inline"`

	// t is an instance of teler WAF.
	t *teler.Teler
}

// CaddyModule returns the Caddy module information.
func (Middleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  moduleID,
		New: func() caddy.Module { return new(Middleware) },
	}
}

// Provision implements caddy.Provisioner.
func (m *Middleware) Provision(ctx caddy.Context) error {
	var err error

	m.Options, err = getTelerOptions(m)
	if err != nil {
		return err
	}

	// NOTE(dwisiswant0): force no standard error output
	m.Options.NoStderr = true

	m.t = teler.New(m.Options)

	return nil
}

// Validate implements caddy.Validator.
func (m *Middleware) Validate() error {
	if m.t == nil {
		return ErrNoTelerInstance
	}

	return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	handler := m.t.CaddyHandler(next)

	return handler.ServeHTTP(w, r)
}
