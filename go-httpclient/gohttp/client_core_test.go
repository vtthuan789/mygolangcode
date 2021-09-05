package gohttp

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
	"time"

	gohttpmock "github.com/vtthuan789/mygolangcode/go-httpclient/gohttp_mock"
	"github.com/vtthuan789/mygolangcode/go-httpclient/gomime"
)

func Test_getRequestBody(t *testing.T) {
	// Initialization
	client := httpClient{}
	// 1st case --- nil body
	t.Run("NilBody", func(t *testing.T) {
		// Execution
		actualBody, err := client.getRequestBody(gomime.ContentTypeJson, nil)
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
		actualBody, err := client.getRequestBody(gomime.ContentTypeJson, requestBody)
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
		actualBody, err := client.getRequestBody(gomime.ContentTypeXml, requestBody)
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

func Test_getConnectionTimeout(t *testing.T) {
	client := NewBuilder().DisableTimeouts(true).Build().(*httpClient)

	t.Run("testTimeoutIsDisabled", func(t *testing.T) {
		timeout := client.getConnectionTimeout()

		if timeout != 0 {
			t.Error("testTimeoutIsDisabled returned timeout value other than 0")
		}
	})

	t.Run("testTimeoutIsEnabled", func(t *testing.T) {
		client.builder.DisableTimeouts(false)
		expectedTimeout := 5 * time.Second
		client.builder.SetConnectionTimeout(expectedTimeout)

		actualTimeout := client.getConnectionTimeout()

		if actualTimeout != expectedTimeout {
			t.Errorf("testTimeoutIsEnabled returned wrong timeout value: expected %d, got %d", expectedTimeout, actualTimeout)
		}
	})

	t.Run("testDefaultTimeout", func(t *testing.T) {
		client.builder.SetConnectionTimeout(0)

		actualTimeout := client.getConnectionTimeout()

		if actualTimeout != defaultConnectionTimeout {
			t.Errorf("testDefaultTimeout returned wrong timeout value: expected %d, got %d", defaultConnectionTimeout, actualTimeout)
		}
	})
}

func Test_getResponseTimeout(t *testing.T) {
	client := NewBuilder().DisableTimeouts(true).Build().(*httpClient)

	t.Run("testTimeoutIsDisabled", func(t *testing.T) {
		timeout := client.getResponseTimeout()

		if timeout != 0 {
			t.Error("testTimeoutIsDisabled returned timeout value other than 0")
		}
	})

	t.Run("testTimeoutIsEnabled", func(t *testing.T) {
		client.builder.DisableTimeouts(false)
		expectedTimeout := 5 * time.Second
		client.builder.SetResponseTimeout(expectedTimeout)

		actualTimeout := client.getResponseTimeout()

		if actualTimeout != expectedTimeout {
			t.Errorf("testTimeoutIsEnabled returned wrong timeout value: expected %d, got %d", expectedTimeout, actualTimeout)
		}
	})

	t.Run("testDefaultTimeout", func(t *testing.T) {
		client.builder.SetResponseTimeout(0)

		actualTimeout := client.getResponseTimeout()

		if actualTimeout != defaultResponseTimeout {
			t.Errorf("testDefaultTimeout returned wrong timeout value: expected %d, got %d", defaultResponseTimeout, actualTimeout)
		}
	})
}

func Test_getMaxIdleConnections(t *testing.T) {
	client := NewBuilder().DisableTimeouts(true).Build().(*httpClient)

	t.Run("testMaxIdleConnectionsNotZero", func(t *testing.T) {
		expectedConnections := 5
		client.builder.SetMaxIdleConnections(expectedConnections)

		actualConnections := client.getMaxIdleConnections()

		if actualConnections != expectedConnections {
			t.Errorf("testMaxIdleConnectionsNotZero returned wrong connections value: expected %d, got %d", expectedConnections, actualConnections)
		}
	})

	t.Run("testZeroMaxIdleConnections", func(t *testing.T) {
		client.builder.SetMaxIdleConnections(0)

		actualConnections := client.getMaxIdleConnections()

		if actualConnections != defaultMaxIdleConnections {
			t.Errorf("testDefaultTimeout returned wrong timeout value: expected %d, got %d", defaultMaxIdleConnections, actualConnections)
		}
	})
}

func Test_getHttpClient(t *testing.T) {
	client := NewBuilder().Build().(*httpClient)

	t.Run("testMockServerNotStart", func(t *testing.T) {
		httpClient := client.getHttpClient()

		if reflect.TypeOf(httpClient).String() != "*http.Client" {
			t.Errorf("testMockServerNotStart returned wrong http client type: expected *http.Client, got %s", reflect.TypeOf(httpClient).String())
		}
	})

	t.Run("testMockServerStarted", func(t *testing.T) {
		gohttpmock.MockupServer.Start()
		httpClient := client.getHttpClient()

		if reflect.TypeOf(httpClient).String() != "*gohttpmock.httpClientMock" {
			t.Errorf("testMockServerStarted returned wrong http client type: expected *gohttpmock.httpClientMock, got %s", reflect.TypeOf(httpClient).String())
		}
	})

	t.Run("testCustomHttpServer", func(t *testing.T) {
		gohttpmock.MockupServer.Stop()
		client := NewBuilder().SetHttpClient(&http.Client{}).SetConnectionTimeout(5 * time.Second).Build().(*httpClient)
		httpClient := client.getHttpClient()

		if reflect.TypeOf(httpClient).String() != "*http.Client" {
			t.Errorf("testCustomHttpServer returned wrong http client type: expected *http.Client, got %s", reflect.TypeOf(httpClient).String())
		}

		if httpClient.(*http.Client).Timeout != 0 {
			t.Errorf("testCustomHttpServer no timeout but got %d", httpClient.(*http.Client).Timeout)
		}
	})
}
func Test_Do(t *testing.T) {
	client := NewBuilder().Build().(*httpClient)

	t.Run("testDoNoError", func(t *testing.T) {
		_, err := client.do(http.MethodGet, "https://test.com", http.Header{}, nil)

		if err != nil {
			t.Errorf("testDoNoError expected no error, but got %s", err.Error())
		}
	})

	t.Run("testDoInvalidRequestBody", func(t *testing.T) {
		_, err := client.do(http.MethodGet, "https://test.com", http.Header{}, complex(6, 2))

		if err.Error() != "json: unsupported type: complex128" {
			t.Errorf("testDoInvalidRequestBody return wrong error message, got %s", err.Error())
		}
	})

	t.Run("testDoInvalidRequestMethod", func(t *testing.T) {
		_, err := client.do("bad method", "https://test.com", http.Header{}, nil)

		if err.Error() != `net/http: invalid method "bad method"` {
			t.Errorf("testDoInvalidRequestMethod return wrong error message, got %s", err.Error())
		}
	})

	t.Run("testDoError", func(t *testing.T) {
		gohttpmock.MockupServer.Start()
		gohttpmock.MockupServer.DeleteMocks()
		gohttpmock.MockupServer.AddMock(gohttpmock.Mock{
			Method: http.MethodGet,
			Url:    "https://test.com",
			Error:  errors.New("error sending HTTP request"),
		})
		_, err := client.do(http.MethodGet, "https://test.com", http.Header{}, nil)

		if err.Error() != "error sending HTTP request" {
			t.Errorf("testDoError return wrong error message, got %s", err.Error())
		}
	})

	t.Run("testDoInvalidResponseBody", func(t *testing.T) {
		gohttpmock.MockupServer.DeleteMocks()
		gohttpmock.MockupServer.AddMock(gohttpmock.Mock{
			Method: http.MethodGet,
			Url:    "https://test.com",

			ResponseStatusCode: http.StatusOK,
			ResponseHeaders:    http.Header{"Content-Length": []string{"1"}},
		})

		_, err := client.do(http.MethodGet, "https://test.com", http.Header{}, nil)

		if err.Error() != "error when reading response body" {
			t.Errorf("testDoInvalidResponseBody return wrong error message, got %s", err.Error())
		}
	})
}
