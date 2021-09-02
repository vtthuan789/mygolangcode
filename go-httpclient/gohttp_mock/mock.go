package gohttpmock

import (
	"fmt"
	"net/http"

	"github.comvtthuan789mygolangcodego-httpclient/gohttp"
)

type Mock struct {
	Method      string
	Url         string
	RequestBody string

	Error              error
	ResponseStatusCode int
	ResponseBody       string
	ResponseHeaders    http.Header
}

func (m *Mock) GetResponse() (*gohttp.Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	response := gohttp.Response{
		Status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		StatusCode: m.ResponseStatusCode,
		Body:       []byte(m.ResponseBody),
		Headers:    make(http.Header),
	}

	for header := range m.ResponseHeaders {
		response.Headers.Set(header, m.ResponseHeaders.Get(header))
	}

	return &response, nil
}
