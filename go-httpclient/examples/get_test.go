package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	gohttpmock "github.com/vtthuan789/mygolangcode/go-httpclient/gohttp_mock"
)

func TestMain(m *testing.M) {
	fmt.Println("About to start test cases for package 'examples'")

	// Tell the HTTP library to mock any further requests from here.
	gohttpmock.MockupServer.Start()

	os.Exit(m.Run())
}

func Test_GetEndpoints(t *testing.T) {
	t.Run("TestErrorFetchingFromGitHub", func(t *testing.T) {
		// Initializtion
		gohttpmock.MockupServer.DeleteMocks()
		gohttpmock.MockupServer.AddMock(gohttpmock.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("error getting github endpoint"),
		})

		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		if endpoints != nil {
			t.Error("expected endpoints but got no endpoint")
			return
		}

		if err == nil {
			t.Error("expected an error but got no error")
			return
		}

		if err.Error() != "error getting github endpoint" {
			t.Errorf("invalid error message received, got %s", err.Error())
			return
		}
	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		// Initializtion
		gohttpmock.MockupServer.DeleteMocks()
		gohttpmock.MockupServer.AddMock(gohttpmock.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 123}`,
		})

		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		if endpoints != nil {
			t.Error("expected endpoints but got no endpoint")
			return
		}

		if err == nil {
			t.Error("expected an error but got no error")
			return
		}

		if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct field") {
			t.Errorf("invalid error message received, got %s", err.Error())
			return
		}
	})

	t.Run("TestNoError", func(t *testing.T) {
		// Initializtion
		gohttpmock.MockupServer.DeleteMocks()
		gohttpmock.MockupServer.AddMock(gohttpmock.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
		})

		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		if err != nil {
			t.Errorf("got an error while should not, got %s", err.Error())
			return
		}

		if endpoints == nil {
			t.Error("expected endpoints but got no endpoint")
			return
		}

		if endpoints.CurrentUserUrl != "https://api.github.com/user" {
			t.Errorf("invalid endpoints.CurrentUserUrl field received, got %s", endpoints.CurrentUserUrl)
			return
		}
	})

}
