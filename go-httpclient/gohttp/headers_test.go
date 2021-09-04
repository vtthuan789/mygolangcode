package gohttp

import (
	"net/http"
	"testing"

	"github.com/vtthuan789/mygolangcode/go-httpclient/gomime"
)

func Test_getRequestHeaders(t *testing.T) {
	// Initialization
	client := httpClient{
		builder: &clientBuilder{},
		client:  &http.Client{},
	}
	commonHeaders := make(http.Header)
	commonHeaders.Set(gomime.HeaderContentType, gomime.ContentTypeJson)
	// commonHeaders.Set(gomime.HeaderUserAgent, "cool-http-client")
	client.builder.headers = commonHeaders

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	client.builder.SetUseAgent("vtthuan-computer")

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
}
