package gohttp

import (
	"fmt"
	"net/http"
	"testing"

	gohttpmock "github.com/vtthuan789/mygolangcode/go-httpclient/gohttp_mock"
	"github.com/vtthuan789/mygolangcode/go-httpclient/gomime"
)

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

func Test_REST_API(t *testing.T) {
	gohttpmock.MockupServer.Start()
	client := NewBuilder().Build()
	t.Run("TestGet", func(t *testing.T) {
		// Initializtion
		gohttpmock.MockupServer.DeleteMocks()
		gohttpmock.MockupServer.AddMock(gohttpmock.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",

			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 123}`,
			ResponseHeaders:    http.Header{gomime.HeaderContentType: []string{gomime.ContentTypeXml}},
		})

		// Execution
		res, err := client.Get("https://api.github.com")

		// Validation
		if err != nil {
			t.Errorf("expected no error but got error %s", err)
		}

		if res.String() != `{"current_user_url": 123}` {
			t.Errorf("invalid response body received, got %s", res.String())
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("invalid status code received: expected %d, got %d", http.StatusOK, res.StatusCode)
		}

		if string(res.Headers.Get(gomime.HeaderContentType)) != gomime.ContentTypeXml {
			t.Errorf(`invalid header received: expected %s, got %s`, gomime.ContentTypeXml, res.Headers.Get(gomime.HeaderContentType))
		}

		if res.Status != fmt.Sprintf("%d %s", http.StatusOK, http.StatusText(http.StatusOK)) {
			t.Errorf("invalid status received: expected %s, got %s", fmt.Sprintf("%d %s", http.StatusOK, http.StatusText(http.StatusOK)), res.Status)
		}
	})

	t.Run("TestPost", func(t *testing.T) {
		// Initializtion
		gohttpmock.MockupServer.DeleteMocks()
		gohttpmock.MockupServer.AddMock(gohttpmock.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","description":"","private":true}`,

			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"id":123,"name":"test-repo"}`,
		})
		postBody := Repository{
			Name:    "test-repo",
			Private: true,
		}
		// Execution
		res, err := client.Post("https://api.github.com/user/repos", postBody)

		// Validation
		if err != nil {
			t.Errorf("expected no error but got error %s", err)
		}

		if res.String() != `{"id":123,"name":"test-repo"}` {
			t.Errorf("invalid response body received, got %s", res.String())
		}

		if res.StatusCode != http.StatusCreated {
			t.Errorf("invalid status code received: expected %d, got %d", http.StatusCreated, res.StatusCode)
		}

		if res.Status != fmt.Sprintf("%d %s", http.StatusCreated, http.StatusText(http.StatusCreated)) {
			t.Errorf("invalid status received: expected %s, got %s", fmt.Sprintf("%d %s", http.StatusCreated, http.StatusText(http.StatusCreated)), res.Status)
		}
	})

	t.Run("TestPut", func(t *testing.T) {
		// Initializtion
		gohttpmock.MockupServer.DeleteMocks()
		gohttpmock.MockupServer.AddMock(gohttpmock.Mock{
			Method:      http.MethodPut,
			Url:         "https://api.github.com/user/123",
			RequestBody: `{"name":"updated-repo","description":"","private":false}`,

			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"id":123,"name":"updated-repo"}`,
		})
		putBody := Repository{
			Name:    "updated-repo",
			Private: false,
		}
		// Execution
		res, err := client.Put("https://api.github.com/user/123", putBody)

		// Validation
		if err != nil {
			t.Errorf("expected no error but got error %s", err)
		}

		if res.String() != `{"id":123,"name":"updated-repo"}` {
			t.Errorf("invalid response body received, got %s", res.String())
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("invalid status code received: expected %d, got %d", http.StatusOK, res.StatusCode)
		}

		if res.Status != fmt.Sprintf("%d %s", http.StatusOK, http.StatusText(http.StatusOK)) {
			t.Errorf("invalid status received: expected %s, got %s", fmt.Sprintf("%d %s", http.StatusOK, http.StatusText(http.StatusOK)), res.Status)
		}
	})

	t.Run("TestPatch", func(t *testing.T) {
		// Initializtion
		gohttpmock.MockupServer.DeleteMocks()
		gohttpmock.MockupServer.AddMock(gohttpmock.Mock{
			Method:      http.MethodPatch,
			Url:         "https://api.github.com/user/123",
			RequestBody: `{"name":"updated-repo","description":"","private":true}`,

			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"id":123,"name":"updated-repo"}`,
		})
		patchBody := Repository{
			Name:    "updated-repo",
			Private: true,
		}
		// Execution
		res, err := client.Patch("https://api.github.com/user/123", patchBody)

		// Validation
		if err != nil {
			t.Errorf("expected no error but got error %s", err)
		}

		if res.String() != `{"id":123,"name":"updated-repo"}` {
			t.Errorf("invalid response body received, got %s", res.String())
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("invalid status code received: expected %d, got %d", http.StatusOK, res.StatusCode)
		}

		if res.Status != fmt.Sprintf("%d %s", http.StatusOK, http.StatusText(http.StatusOK)) {
			t.Errorf("invalid status received: expected %s, got %s", fmt.Sprintf("%d %s", http.StatusOK, http.StatusText(http.StatusOK)), res.Status)
		}
	})

	t.Run("TestDelete", func(t *testing.T) {
		// Initializtion
		gohttpmock.MockupServer.DeleteMocks()
		gohttpmock.MockupServer.AddMock(gohttpmock.Mock{
			Method: http.MethodDelete,
			Url:    "https://api.github.com/user/123",

			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"id":123,"name":"updated-repo"}`,
			ResponseHeaders:    http.Header{gomime.HeaderContentType: []string{gomime.ContentTypeJson}},
		})

		// Execution
		res, err := client.Delete("https://api.github.com/user/123")

		// Validation
		if err != nil {
			t.Errorf("expected no error but got error %s", err)
		}

		if res.String() != `{"id":123,"name":"updated-repo"}` {
			t.Errorf("invalid response body received, got %s", res.String())
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("invalid status code received: expected %d, got %d", http.StatusOK, res.StatusCode)
		}

		if string(res.Headers.Get(gomime.HeaderContentType)) != gomime.ContentTypeJson {
			t.Errorf(`invalid header received: expected %s, got %s`, gomime.ContentTypeJson, res.Headers.Get(gomime.HeaderContentType))
		}

		if res.Status != fmt.Sprintf("%d %s", http.StatusOK, http.StatusText(http.StatusOK)) {
			t.Errorf("invalid status received: expected %s, got %s", fmt.Sprintf("%d %s", http.StatusOK, http.StatusText(http.StatusOK)), res.Status)
		}
	})

	t.Run("TestOptions", func(t *testing.T) {
		// Initializtion
		gohttpmock.MockupServer.DeleteMocks()
		gohttpmock.MockupServer.AddMock(gohttpmock.Mock{
			Method: http.MethodOptions,
			Url:    "https://api.github.com",

			ResponseStatusCode: http.StatusNoContent,
			ResponseHeaders:    http.Header{"Access-Control-Allow-Origin": []string{"https://api.github.com"}},
		})

		// Execution
		res, err := client.Options("https://api.github.com")

		// Validation
		if err != nil {
			t.Errorf("expected no error but got error %s", err)
		}

		if res.StatusCode != http.StatusNoContent {
			t.Errorf("invalid status code received: expected %d, got %d", http.StatusNoContent, res.StatusCode)
		}

		if string(res.Headers.Get("Access-Control-Allow-Origin")) != "https://api.github.com" {
			t.Errorf(`invalid header received: expected %s, got %s`, "https://api.github.com", res.Headers.Get("Access-Control-Allow-Origin"))
		}

		if res.Status != fmt.Sprintf("%d %s", http.StatusNoContent, http.StatusText(http.StatusNoContent)) {
			t.Errorf("invalid status received: expected %s, got %s", fmt.Sprintf("%d %s", http.StatusNoContent, http.StatusText(http.StatusNoContent)), res.Status)
		}
	})
}
