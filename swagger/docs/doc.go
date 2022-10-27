// Package docs awesome.
//
// Documentation of our awesome API.
//
//	 Schemes: http, https
//	 BasePath: /
//	 Version: 0.1.0
//	 Host: some-url.com
//
//	 Consumes:
//	 - application/json
//
//	 Produces:
//	 - application/json
//
//	 Security:
//	 - basic
//
//	SecurityDefinitions:
//	basic:
//	  type: basic
//
// swagger:meta
package docs

//go:generate swagger generate spec -o swagger.yaml
//go:generate swagger serve --no-open -F=swagger --port 36666 swagger.yaml
