package gohttp

import (
	"net/http"

	"github.com/vtthuan789/mygolangcode/go-httpclient/gomime"
)

func getHeaders(headers ...http.Header) http.Header {
	result := http.Header{}
	if len(headers) > 0 {
		result = headers[0]
	}
	return result
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	// Add common headers from the HTTP client instance
	for header, value := range c.builder.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// Add custom headers from the current request
	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// Set User-Agent if it is defined and not there yet
	if c.builder.userAgent != "" {
		if result.Get(gomime.HeaderUserAgent) != "" {
			return result
		}
		result.Set(gomime.HeaderUserAgent, c.builder.userAgent)
	}

	return result
}
