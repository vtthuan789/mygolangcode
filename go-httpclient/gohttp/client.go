package gohttp

import (
	"net/http"
	"sync"

	"github.comvtthuan789mygolangcodego-httpclient/core"
)

type httpClient struct {
	builder *clientBuilder

	client     *http.Client
	clientOnce sync.Once
}

type Client interface {
	Get(url string, headers http.Header) (*core.Response, error)
	Post(url string, headers http.Header, body interface{}) (*core.Response, error)
	Put(url string, headers http.Header, body interface{}) (*core.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*core.Response, error)
	Delete(url string, headers http.Header) (*core.Response, error)
	Options(url string, headers http.Header) (*core.Response, error)
}

func (c *httpClient) Get(url string, headers http.Header) (*core.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*core.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*core.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*core.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*core.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}

func (c *httpClient) Options(url string, headers http.Header) (*core.Response, error) {
	return c.do(http.MethodOptions, url, headers, nil)
}
