package gohttpmock

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type httpClientMock struct{}

type HttpClientMock interface {
	Do(request *http.Request) (*http.Response, error)
}

func (c *httpClientMock) Do(request *http.Request) (*http.Response, error) {
	requestBody, err := request.GetBody()
	if err != nil {
		return nil, err
	}
	defer requestBody.Close()

	body, err := ioutil.ReadAll(requestBody)
	if err != nil {
		return nil, err
	}

	var response http.Response

	mock := MockupServer.mocks[MockupServer.getMockKey(request.Method, request.URL.String(), string(body))]
	if mock != nil {
		if mock.Error != nil {
			return nil, mock.Error
		}
		response.Status = fmt.Sprintf("%d %s", mock.ResponseStatusCode, http.StatusText(mock.ResponseStatusCode))
		response.StatusCode = mock.ResponseStatusCode
		response.Body = ioutil.NopCloser(strings.NewReader(mock.ResponseBody))
		response.ContentLength = int64(len(mock.ResponseBody))
		response.Request = request
		response.Header = mock.ResponseHeaders
		return &response, nil
	}
	return nil, fmt.Errorf("no mock matching %s from '%s' with given body", request.Method, request.URL.String())
}
