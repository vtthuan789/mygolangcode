package gohttp

import (
	"testing"
)

func Test_getRequestBody(t *testing.T) {
	// Initialization
	client := httpClient{}
	// 1st case --- nil body
	t.Run("NilBody", func(t *testing.T) {
		// Execution
		actualBody, err := client.getRequestBody("application/json", nil)
		if err != nil {
			t.Errorf("test NilBody expected no error, but got %s", err.Error())
		}

		if actualBody != nil {
			t.Errorf("test NilBody expected nil body, but got %s", string(actualBody))
		}
	})
	// 2nd case --- JSON body
	t.Run("JsonBody", func(t *testing.T) {
		// Execution
		requestBody := []string{"one", "two"}
		actualBody, err := client.getRequestBody("application/json", requestBody)
		if err != nil {
			t.Errorf("test JsonBody expected no error, but got %s", err.Error())
		}

		if string(actualBody) != `["one","two"]` {
			t.Errorf("test JsonBody did not return expected body, got %s", string(actualBody))
		}
	})
	// 3rd case --- XML body
	t.Run("XmlBody", func(t *testing.T) {
		// Execution
		requestBody := []string{"one", "two"}
		actualBody, err := client.getRequestBody("application/xml", requestBody)
		if err != nil {
			t.Errorf("test XmlBody expected no error, but got %s", err.Error())
		}

		if string(actualBody) != `<string>one</string><string>two</string>` {
			t.Errorf("test XmlBody did not return expected body, got %s", string(actualBody))
		}
	})
	// 4th case --- JSON body by default
	t.Run("JsonBodyByDefault", func(t *testing.T) {
		// Execution
		requestBody := []string{"one", "two"}
		actualBody, err := client.getRequestBody("", requestBody)
		if err != nil {
			t.Errorf("test JsonBody expected no error, but got %s", err.Error())
		}

		if string(actualBody) != `["one","two"]` {
			t.Errorf("test JsonBody did not return expected body, got %s", string(actualBody))
		}
	})
}
