package examples

import (
	"errors"
	"fmt"
	"net/http"
)

type GithubError struct {
	StatusCode       int    `json:"-"`
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
}

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

func CreateRepo(repo Repository) (*Repository, error) {
	fmt.Println("repo:", repo)
	response, err := httpClient.Post("https://api.github.com/user/repos", repo)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusCreated {
		var githubError GithubError
		if err := response.UnmarshalJson(&githubError); err != nil {
			return nil, errors.New("error when parsing json response to GithubError type")
		}
		return nil, errors.New(githubError.Message)
	}

	var repository Repository
	if err := response.UnmarshalJson(&repository); err != nil {
		return nil, errors.New("error when parsing json response to Repository type")
	}

	return &repository, nil
}
