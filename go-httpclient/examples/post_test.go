package examples

import (
	"errors"
	"net/http"
	"testing"

	gohttpmock "github.com/vtthuan789/mygolangcode/go-httpclient/gohttp_mock"
)

func Test_CreateRepo(t *testing.T) {
	t.Run("testErrorCreatingRepoOnGitHub", func(t *testing.T) {
		// Initializtion
		gohttpmock.MockupServer.DeleteMocks()
		gohttpmock.MockupServer.AddMock(gohttpmock.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","description":"","private":true}`,
			Error:       errors.New("error creating github repository"),
		})

		// Execution
		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}
		repo, err := CreateRepo(repository)

		// Validation
		if repo != nil {
			t.Error("testErrorCreatingRepoOnGitHub expected getting no repo but got one")
			return
		}

		if err == nil {
			t.Error("testErrorCreatingRepoOnGitHub expected an error but got no error")
			return
		}

		if err.Error() != "error creating github repository" {
			t.Errorf("testErrorCreatingRepoOnGitHub received invalid error message, got \"%s\"", err.Error())
			return
		}
	})

	t.Run("testWrongStatusCode", func(t *testing.T) {
		// Initializtion
		gohttpmock.MockupServer.DeleteMocks()
		gohttpmock.MockupServer.AddMock(gohttpmock.Mock{
			Method:             http.MethodPost,
			Url:                "https://api.github.com/user/repos",
			RequestBody:        `{"name":"test-repo","description":"","private":true}`,
			ResponseStatusCode: http.StatusConflict,
			ResponseBody:       `{"message":"repo already existed","documentation_url":"https://some-url"}`,
		})

		// Execution
		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}
		repo, err := CreateRepo(repository)

		// Validation
		if repo != nil {
			t.Error("testWrongStatusCode expected getting no repo but got one")
			return
		}

		if err == nil {
			t.Error("testWrongStatusCode expected an error but got no error")
			return
		}

		if err.Error() != "repo already existed" {
			t.Errorf("testWrongStatusCode received invalid error message, got \"%s\"", err.Error())
			return
		}
	})

	t.Run("testErrorUnmarshalGithubError", func(t *testing.T) {
		// Initializtion
		gohttpmock.MockupServer.DeleteMocks()
		gohttpmock.MockupServer.AddMock(gohttpmock.Mock{
			Method:             http.MethodPost,
			Url:                "https://api.github.com/user/repos",
			RequestBody:        `{"name":"test-repo","description":"","private":true}`,
			ResponseStatusCode: http.StatusConflict,
			ResponseBody:       `{"message":"repo already existed","documentation_url":123}`,
		})

		// Execution
		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}
		repo, err := CreateRepo(repository)

		// Validation
		if repo != nil {
			t.Error("testErrorUnmarshalGithubError expected getting no repo but got one")
			return
		}

		if err == nil {
			t.Error("testErrorUnmarshalGithubError expected an error but got no error")
			return
		}

		if err.Error() != "error when parsing json response to GithubError type" {
			t.Errorf("testErrorUnmarshalGithubError received invalid error message, got \"%s\"", err.Error())
			return
		}
	})

	t.Run("testErrorUnmarshalRepository", func(t *testing.T) {
		// Initializtion
		gohttpmock.MockupServer.DeleteMocks()
		gohttpmock.MockupServer.AddMock(gohttpmock.Mock{
			Method:             http.MethodPost,
			Url:                "https://api.github.com/user/repos",
			RequestBody:        `{"name":"test-repo","description":"","private":true}`,
			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"id":wrongid,"name":"test-repo"}`,
		})

		// Execution
		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}
		repo, err := CreateRepo(repository)

		// Validation
		if repo != nil {
			t.Error("testErrorUnmarshalRepository expected getting no repo but got one")
			return
		}

		if err == nil {
			t.Error("testErrorUnmarshalRepository expected an error but got no error")
			return
		}

		if err.Error() != "error when parsing json response to Repository type" {
			t.Errorf("testErrorUnmarshalRepository received invalid error message, got \"%s\"", err.Error())
			return
		}
	})

	t.Run("testNoError", func(t *testing.T) {
		// Initializtion
		gohttpmock.MockupServer.DeleteMocks()
		gohttpmock.MockupServer.AddMock(gohttpmock.Mock{
			Method:             http.MethodPost,
			Url:                "https://api.github.com/user/repos",
			RequestBody:        `{"name":"test-repo","description":"","private":true}`,
			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"id":123,"name":"test-repo"}`,
		})

		// Execution
		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}
		repo, err := CreateRepo(repository)

		// Validation
		if err != nil {
			t.Error("testNoError expected no error but got one")
			return
		}

		if repo == nil {
			t.Error("testNoError expected getting a repo but did not")
			return
		}

		if repo.Name != repository.Name {
			t.Errorf("testNoError received invalid repository name, got %s", repo.Name)
			return
		}
	})

}
