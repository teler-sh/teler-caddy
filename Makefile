.DEFAULT_GOAL := help
DIST := ./dist
CADDY := ${DIST}/caddy
ifeq ($(shell which ${CADDY}),)
	CADDY = $(shell which caddy)
endif

help: ## Displays this help message
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Builds the Caddy core and plugs in the teler WAF module (Output: ./dist/caddy)
	@CGO_ENABLED=1 xcaddy build \
		--with github.com/teler-sh/teler-caddy@latest --output dist/caddy

build-local: ## Same as `build` but uses the teler WAF module locally
	@CGO_ENABLED=1 xcaddy build \
		--with github.com/teler-sh/teler-caddy=. --output dist/caddy

adapt: ## Converts a Caddyfile to Caddy's native JSON format (Output: ./caddy.example.json)
	@${CADDY} adapt -c example.Caddyfile -p | tee caddy.example.json

run: ## Runs the Caddy server with Caddy's native JSON configuration
	@${CADDY} run -c caddy.example.json

run-httpbin: ## Runs the httpbin server on port 8081
	@go run -v github.com/mccutchen/go-httpbin/v2/cmd/go-httpbin@latest -port 8081

.PHONY: build build-local adapt run run-httpbin