package languages

import (
	"api_git_leet_duo/api/git/tools/auth"
	"api_git_leet_duo/api/git/tools/graphql"
	"fmt"
	"errors"
	"net/http"
	"encoding/json"
	"bytes"
)

type Language struct {
	Name string `json:"name"`
}


type LanguageEdge struct {
	Size int     `json:"size"`
	Node Language `json:"node"`
}

type Repository struct {
	Name	  	string `json:"name"`
	DateCreate	string `json:"createdAt"`
	Languages struct {
		Edges []LanguageEdge `json:"edges"`
	} `json:"languages"`
	DefaultBranchRef struct {
		Target struct {
			CommittedDate string `json:"committedDate"`
		} `json:"target"`
	} `json:"defaultBranchRef"`
}

type Repo struct {
	Repositories struct {
		Nodes []Repository `json:"nodes"`
	} `json:"repositories"`
}

type ResponseLangs struct {
	Data struct {
		Repo Repo `json:"user"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}


type RepositoryLite struct {
	Name	  	string `json:"name"`
}


type RepoLite struct {
	Repositories struct {
		Nodes []RepositoryLite `json:"nodes"`
	} `json:"repositories"`
}

type ResponseLite struct {
	Data struct {
		Repo RepoLite `json:"user"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}



func FetchUserLangsFull(user string) (Repo, error) {
	token, err := auth.GetGitHubTokenNative()
	if err != nil {
		return Repo{}, err
	}

	query := graphql.BuildGraphQLQueryLangFull(user)
	body, _ := json.Marshal(graphql.GraphQLQuery{Query: query})

	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(body))
	if err != nil {
		return Repo{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Repo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Repo{}, fmt.Errorf("erro na requisição: status %d", resp.StatusCode)
	}

	var response ResponseLangs
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return Repo{}, err
	}

	if len(response.Errors) > 0 {
		return Repo{}, errors.New(response.Errors[0].Message)
	}

	return response.Data.Repo, nil
}







func FetchUserLite(user string)  (RepoLite, error) {
	token, err := auth.GetGitHubTokenNative()
	if err != nil {
		return RepoLite{}, err
	}

	query := graphql.BuildGraphQLQueryLite(user)
	body, _ := json.Marshal(graphql.GraphQLQuery{Query: query})

	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(body))
	if err != nil {
		return RepoLite{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return RepoLite{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return RepoLite{}, fmt.Errorf("erro na requisição: status %d", resp.StatusCode)
	}

	var response ResponseLite
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return RepoLite{}, err
	}

	if len(response.Errors) > 0 {
		return RepoLite{}, errors.New(response.Errors[0].Message)
	}

	return response.Data.Repo, nil
}