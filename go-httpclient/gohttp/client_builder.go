package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers            http.Header
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	maxIdleConnections int
	disableTimeouts    bool
	client             *http.Client
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(i int) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
	SetHttpClient(client *http.Client) ClientBuilder
	Build() Client
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() Client {
	client := &httpClient{
		builder: c,
	}

	return client
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(i int) ClientBuilder {
	c.maxIdleConnections = i
	return c
}

func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable
	return c
}

func (c *clientBuilder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client
	return c
}
