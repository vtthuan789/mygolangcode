package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/vtthuan789/mygolangcode/go-httpclient/core"
	gohttpmock "github.com/vtthuan789/mygolangcode/go-httpclient/gohttp_mock"
	"github.com/vtthuan789/mygolangcode/go-httpclient/gomime"
)

const (
	defaultMaxIdleConnections = 5
	defaultResponseTimeout    = 5 * time.Second
	defaultConnectionTimeout  = 1 * time.Second
)

func (c *httpClient) do(method, url string, headers http.Header, body interface{}) (*core.Response, error) {
	fullHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get(gomime.HeaderContentType), body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	client := c.getHttpClient()

	httpResponse, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer httpResponse.Body.Close()

	bytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	response := &core.Response{
		Status:     httpResponse.Status,
		StatusCode: httpResponse.StatusCode,
		Headers:    httpResponse.Header,
		Body:       bytes,
	}

	return response, nil
}

func (c *httpClient) getHttpClient() core.HttpClient {
	if gohttpmock.MockupServer.IsEnabled() {
		return gohttpmock.MockupServer.GetMockedClient()
	}

	c.clientOnce.Do(func() {
		fmt.Println("Creating a new http client!!!!!")
		if c.builder.client != nil {
			c.client = c.builder.client
			return
		}

		c.client = &http.Client{
			Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
			},
		}
	})

	return c.client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.builder.maxIdleConnections > 0 {
		return c.builder.maxIdleConnections
	}

	return defaultMaxIdleConnections
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.builder.disableTimeouts {
		return 0
	}

	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}

	return defaultResponseTimeout
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.disableTimeouts {
		return 0
	}

	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}

	return defaultConnectionTimeout
}

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case gomime.ContentTypeJson:
		return json.Marshal(body)
	case gomime.ContentTypeXml:
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}
