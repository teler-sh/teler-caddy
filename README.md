# teler Caddy

**teler-caddy** integrates the robust security features of teler WAF into the Caddy web server. With the teler Caddy module, you can leverage these comprehensive security measures to ensure your web servers remain secure and resilient against OWASP Top 10 threats, known vulnerabilities, malicious actors, botnets, unwanted crawlers, and brute force attacks.

---

* [Usage](#usage)
  * [Configuration](#configuration)
    * [`load_from` subdirective](#load_from-subdirective)
    * [`inline` subdirective](#inline-subdirective)
  * [Examples](#examples)
* [Development](#development)
* [JSON structure](#json-structure)
* [Demo](#demo)
* [Community](#community)
* [License](#license)

## Usage

To use this module, follow these steps:

* Build the `caddy` core and plug-in this module with [xcaddy](https://github.com/caddyserver/xcaddy).

```bash
CGO_ENABLED=1 xcaddy build \
    --with github.com/teler-sh/teler-caddy@latest --output dist/caddy
```

* Add the **`teler_waf`** directive within your `route` configuration.
* Then, run the Caddy server with the specified configuration: `./dist/caddy run --config /path/to/your/Caddyfile`.

That's it! By following these steps, you will integrate teler WAF into your Caddy server. The teler WAF now will seamlessly apply a default configuration, ensuring that your site remains protected with sensible and reasonable settings.

### Configuration

This module allows for fine-tuning and customization through two subdirectives: **`load_from`** and **`inline`**. These subdirectives enable you to set various options to tailor the behavior of the teler WAF to meet your specific security needs.

Here is the syntax and usage for each subdirective:

#### `load_from` subdirective

Use this subdirective to load teler WAF configuration from a specified file. The configuration file can be in JSON or YAML format.

```caddy
load_from <format> <filepath>
```

> [!NOTE]
> * **format**: Specifies the format of the teler WAF configuration file. Valid values are **`json`** and **`yaml`** *(case-insensitive)*.
> * **filepath**: Specifies the location path of the teler WAF configuration file.

#### `inline` subdirective

Use this subdirective to define teler WAF configuration options directly within the Caddyfile. The configuration can be provided in JSON or YAML format.

```caddy
inline <format> <options>
```

These configuration subdirectives provide flexibility in managing the teler WAF settings, allowing you to either load configurations from an external file or define them directly within your Caddyfile, ensuring that your web servers are adequately protected with tailored security measures.

### Examples

Here are examples of how to configure this module using the **`load_from`** and **`inline`** subdirectives:

* With **load_from** subdirective

This example demonstrates how to load the teler WAF configuration from a YAML or JSON file.

```caddy
example.com {
    route {
        teler_waf {
            load_from YAML /path/to/your/teler-waf.conf.yaml
            # or
            load_from JSON /path/to/your/teler-waf.conf.json
        }
    }
}
```

* With **inline** subdirective

This example demonstrates how to define the teler WAF configuration directly within the Caddyfile.

> [!TIP]
> For better readability and management of options, write your teler WAF options using backticks or heredoc. See [Tokens and quotes](https://caddyserver.com/docs/caddyfile/concepts#tokens-and-quotes).

```caddy
example.com {
    route {
        teler_waf {
            inline YAML <<--
                excludes: []
                whitelists: []
                customs: []
                customs_from_file: ""
                log_file: ""
                no_stderr: false
                no_update_check: false
                development: false
                in_memory: false
                falcosidekick_url: ""
                verbose: false
                --
        }
    }
}
```

> [!TIP]
> To apply this module globally as middleware across all routes, reorder the teler WAF module directive to be the first in the Caddy's HTTP handler chain.

```caddy
{
    order teler_waf first
}

example.com {
    teler_waf {
        # load_from ...
        # inline ...
    }
}
```

These examples illustrate how to effectively configure the teler WAF in Caddy using different methods, providing flexibility to suit your specific setup and requirements.

## Development

Here are the available commands to assist with development:

```console
$ make
help                           Displays this help message
build                          Builds the Caddy core and plug-in teler WAF module (Output: ./dist/caddy)
build-local                    Same as `build` but use teler WAF module locally
adapt                          Converts a Caddyfile to Caddy's native JSON format (Output: ./caddy.example.json)
run                            Runs the Caddy with Caddy's native JSON configuration
run-httpbin                    Runs the httpbin server with port 8081
```

## JSON Structure

Here is how you can configure this module using both the Caddyfile and its equivalent Caddy's native JSON structure.

Caddyfile:

```caddy
:8080 {
    route {
        teler_waf
    }

    reverse_proxy localhost:8081
}
```

The same configuration can be expressed in Caddy's native JSON structure as follows:

```json
{
  "apps": {
    "http": {
      "servers": {
        "srv0": {
          "listen": [
            ":8080"
          ],
          "routes": [
            {
              "handle": [
                {
                  "handler": "subroute",
                  "routes": [
                    {
                      "handle": [
                        {
                          "format": "",
                          "handler": "teler",
                          "inline": "",
                          "load_from": ""
                        }
                      ]
                    }
                  ]
                },
                {
                  "handler": "reverse_proxy",
                  "upstreams": [
                    {
                      "dial": "localhost:8081"
                    }
                  ]
                }
              ]
            }
          ]
        }
      }
    }
  }
}
```

## Demo

To demonstrate the teler Caddy module in action, follow these steps:

```bash
# in: tty1
$ make build-local
$ make run-httpbin

# in: tty2
$ make run

# in: tty3
$ curl localhost:8080

# out: tty2
2024/06/19 23:15:29.580 ERROR   http.log.error  bad crawler {"request": {"remote_ip": "::1", "remote_port": "59510", "client_ip": "::1", "proto": "HTTP/1.1", "method": "GET", "host": "localhost:8080", "uri": "/", "headers": {"User-Agent": ["curl/8.6.0"], "Accept": ["*/*"]}}, "duration": 0.004394569}
```

This demo showcases the module's ability to detect and defend against various forms of cyber threats, providing an example of its protective capabilities in a real-world scenario.

## Community

We use the Google Groups as our dedicated mailing list. Subscribe to [teler-announce](https://groups.google.com/g/teler-announce) via [teler-announce+subscribe@googlegroups.com](mailto:teler-announce+subscribe@googlegroups.com) for important announcements, such as the availability of new releases. This subscription will keep you informed about significant developments related to [teler IDS](https://github.com/teler-sh/teler), [teler WAF](https://github.com/teler-sh/teler-waf), [teler Proxy](https://github.com/teler-sh/teler-proxy), [teler Caddy](https://github.com/teler-sh/teler-caddy), and [teler Resources](https://github.com/teler-sh/teler-resources).

For any [inquiries](https://github.com/teler-sh/teler-caddy/discussions/categories/q-a), [discussions](https://github.com/teler-sh/teler-caddy/discussions), or [issues](https://github.com/teler-sh/teler-caddy/issues) are being tracked here on GitHub. This is where we actively manage and address these aspects of our community engagement.

## License

This module is free software: you can redistribute it and/or modify it under the terms of the [Apache-2.0 license](/LICENSE). teler-caddy and any contributions are copyright © by Dwi Siswanto 2024.