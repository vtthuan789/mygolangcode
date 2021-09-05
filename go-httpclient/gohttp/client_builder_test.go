package gohttp

import (
	"net/http"
	"testing"
	"time"

	"github.com/vtthuan789/mygolangcode/go-httpclient/gomime"
)

func Test_ClientBuilder(t *testing.T) {
	client := NewBuilder().
		SetHeaders(http.Header{gomime.HeaderContentType: []string{gomime.ContentTypeJson}}).
		SetConnectionTimeout(5 * time.Second).
		SetResponseTimeout(5 * time.Second).
		SetMaxIdleConnections(5).
		DisableTimeouts(true).
		SetHttpClient(&http.Client{Timeout: 0}).
		SetUseAgent("computer-name").Build().(*httpClient)

	if client.builder.headers[gomime.HeaderContentType][0] != gomime.ContentTypeJson {
		t.Error("Test_ClientBuilder failed because of wrong headers")
	}

	if client.builder.connectionTimeout != 5*time.Second {
		t.Error("Test_ClientBuilder failed because of wrong connectionTimeout")
	}

	if client.builder.responseTimeout != 5*time.Second {
		t.Error("Test_ClientBuilder failed because of wrong responseTimeout")
	}

	if time.Duration(client.builder.maxIdleConnections) != 5 {
		t.Error("Test_ClientBuilder failed because of wrong maxIdleConnections")
	}

	if client.builder.disableTimeouts != true {
		t.Error("Test_ClientBuilder failed because of wrong disableTimeouts")
	}

	if client.builder.userAgent != "computer-name" {
		t.Error("Test_ClientBuilder failed because of wrong userAgent")
	}

	if client.builder.client.Timeout.String() != "0s" {
		t.Error("Test_ClientBuilder failed because of wrong client timeout")
	}
}
