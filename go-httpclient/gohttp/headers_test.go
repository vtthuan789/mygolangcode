package gohttp

import (
	"net/http"
	"testing"

	"github.com/vtthuan789/mygolangcode/go-httpclient/gomime"
)

func Test_getRequestHeaders(t *testing.T) {
	t.Run("testUserAgentIsDefined", func(t *testing.T) {
		// Initialization
		client := httpClient{
			builder: &clientBuilder{},
			client:  &http.Client{},
		}
		commonHeaders := make(http.Header)
		commonHeaders.Set(gomime.HeaderContentType, gomime.ContentTypeJson)
		commonHeaders.Set(gomime.HeaderUserAgent, "computer-name")
		client.builder.headers = commonHeaders

		requestHeaders := make(http.Header)
		requestHeaders.Set("X-Request-Id", "ABC-123")

		client.builder.SetUseAgent("invalid-computer")

		// Execution
		actualHeaders := client.getRequestHeaders(requestHeaders)

		// Validation
		if len(actualHeaders) != 3 {
			t.Errorf("getRequestHeaders expected to return 3 headers, but got %d", len(actualHeaders))
		}

		for expectedKey, expectedValue := range commonHeaders {
			if actualHeaders.Get(expectedKey) != expectedValue[0] {
				t.Errorf("getRequestHeaders failed when testing for common header %s", expectedKey)
			}
		}

		for expectedKey, expectedValue := range requestHeaders {
			if actualHeaders.Get(expectedKey) != expectedValue[0] {
				t.Errorf("getRequestHeaders failed when testing for custom header %s", expectedKey)
			}
		}
	})

	t.Run("testUserAgentIsNotDefined", func(t *testing.T) {
		// Initialization
		client := httpClient{
			builder: &clientBuilder{},
			client:  &http.Client{},
		}
		commonHeaders := make(http.Header)
		commonHeaders.Set(gomime.HeaderContentType, gomime.ContentTypeJson)
		client.builder.headers = commonHeaders

		requestHeaders := make(http.Header)
		requestHeaders.Set("X-Request-Id", "ABC-123")

		userAgent := "computer-name"
		client.builder.SetUseAgent(userAgent)

		// Execution
		actualHeaders := client.getRequestHeaders(requestHeaders)

		// Validation
		if len(actualHeaders) != 3 {
			t.Errorf("getRequestHeaders expected to return 3 headers, but got %d", len(actualHeaders))
		}

		for expectedKey, expectedValue := range commonHeaders {
			if actualHeaders.Get(expectedKey) != expectedValue[0] {
				t.Errorf("getRequestHeaders failed when testing for common header %s", expectedKey)
			}
		}

		for expectedKey, expectedValue := range requestHeaders {
			if actualHeaders.Get(expectedKey) != expectedValue[0] {
				t.Errorf("getRequestHeaders failed when testing for custom header %s", expectedKey)
			}
		}

		if actualHeaders.Get(gomime.HeaderUserAgent) != "computer-name" {
			t.Errorf("getRequestHeaders failed when testing for %s header, expected %s, got %s",
				gomime.HeaderUserAgent, userAgent, actualHeaders.Get(gomime.HeaderUserAgent))
		}
	})
}

func Test_getHeaders(t *testing.T) {
	t.Run("testEmptyArgument", func(t *testing.T) {
		headers := getHeaders()

		if len(headers) != 0 {
			t.Error("testEmptyArgument failed to return empty header")
		}
	})

	t.Run("testNonEmptyArgument", func(t *testing.T) {
		headers := getHeaders(http.Header{gomime.HeaderContentType: []string{gomime.ContentTypeJson}},
			http.Header{"X-Request-Id": []string{"ABC-123"}})

		if len(headers) != 1 {
			t.Error("testNonEmptyArgument failed to return exactly one header")
		}

		if value, ok := headers[gomime.HeaderContentType]; !ok {
			t.Errorf("testNonEmptyArgument failed to test for header key: expected %s", gomime.HeaderContentType)

			if value[0] != gomime.ContentTypeJson {
				t.Errorf("testNonEmptyArgument failed to test for header value: expected %s, got %s", gomime.ContentTypeJson, value[0])
			}
		}
	})
}
