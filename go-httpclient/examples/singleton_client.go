package examples

import (
	"time"

	"github.com/vtthuan789/mygolangcode/go-httpclient/gohttp"
)

var httpClient = getHttpClient()

func getHttpClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetMaxIdleConnections(5).
		SetUseAgent("vtthuan-computer").
		Build()

	return client
}
