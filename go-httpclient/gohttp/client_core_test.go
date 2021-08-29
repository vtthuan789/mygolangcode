package gohttp

import (
	"net/http"
	"testing"
)

func Test_getRequestHeaders(t *testing.T) {
	// Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.Headers = commonHeaders

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	// Execution
	testHeader := client.getRequestHeaders(requestHeaders)

	// Validation
	if len(testHeader) != 3 {
		t.Errorf("getRequestHeaders expected to return header with length 3, but got %d", len(testHeader))
	}
}
