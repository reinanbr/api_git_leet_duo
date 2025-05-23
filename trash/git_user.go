package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"math/rand"

	"api_git_leet_duo/api/git/tools/user"
)

//////////////////////////
// 🔑 Token Management //
//////////////////////////

// GetGitHubTokens retrieves all GitHub tokens from environment variables.
func GetGitHubTokens() []string {
	var tokens []string
	if token, exists := os.LookupEnv("TOKEN"); exists {
		tokens = append(tokens, token)
	}
	for i := 2; ; i++ {
		envVar := fmt.Sprintf("TOKEN%d", i)
		if token, exists := os.LookupEnv(envVar); exists {
			tokens = append(tokens, token)
		} else {
			break
		}
	}
	return tokens
}

// GetGitHubTokenNative retrieves the default token from the environment.
func GetGitHubTokenNative() (string, error) {
	token := os.Getenv("TOKEN")
	if token == "" {
		return "", errors.New("GitHub token não encontrado")
	}
	return token, nil
}

// GetGitHubToken selects a random token from the provided slice.
func GetGitHubToken(tokens []string) (string, error) {
	if len(tokens) == 0 {
		return "", errors.New("no GitHub token available")
	}
	return tokens[rand.Intn(len(tokens))], nil
}

/////////////////////////////
// 🔧 Query Construction //
/////////////////////////////

type GraphQLQuery struct {
	Query string `json:"query"`
}

// Validate that user is not empty.
func validateUser(user string) error {
	if user == "" {
		return errors.New("user parameter cannot be empty")
	}
	return nil
}

// BuildGraphQLQueryUser generates the user info query.
func BuildGraphQLQueryUser(user string) string {
	return fmt.Sprintf(`
	{
		user(login: "%s") {
			name
			login
			bio
			avatarUrl
			createdAt
		}
	}
	`, user)
}


func BuildGraphQLQueryRepos(user string, cursor *string) string {
	after := ""
	if cursor != nil {
		after = fmt.Sprintf(`, after: "%s"`, *cursor)
	}

	return fmt.Sprintf(`
	{
		user(login: "%s") {
			repositories(first: 100, privacy: PUBLIC%s) {
				pageInfo {
					hasNextPage
					endCursor
				}
				nodes {
					name
					createdAt
					languages(first: 100) {
						edges {
							size
							node {
								name
							}
						}
					}
				}
			}
		}
	}
	`, user, after)
}


// BuildGraphQLQueryLangFull generates a detailed language usage query.
func BuildGraphQLQueryLangFull(user string) (string, error) {
	if err := validateUser(user); err != nil {
		return "", err
	}

	query := `
	{
		user(login: "%s") {
			repositories(first: 100, privacy: PUBLIC) {
				nodes {
					name
					createdAt
					languages(first: 100) {
						edges {
							size
							node {
								name
							}
						}
					}
					defaultBranchRef {
						target {
							... on Commit {
								committedDate
							}
						}
					}
				}
			}
		}
	}
	`
	return fmt.Sprintf(query, user), nil
}

type Repo struct {
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

type Language struct {
	Name string `json:"name"`
	Size int    `json:"size"`
}

type RepoNode struct {
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	Languages struct {
		Edges []struct {
			Size int `json:"size"`
			Node struct {
				Name string `json:"name"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"languages"`
}

type RepoResponse struct {
	Data struct {
		User struct {
			Repositories struct {
				PageInfo struct {
					HasNextPage bool   `json:"hasNextPage"`
					EndCursor   string `json:"endCursor"`
				} `json:"pageInfo"`
				Nodes []RepoNode `json:"nodes"`
			} `json:"repositories"`
		} `json:"user"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

func FetchAllRepos(user string, token string, cursor *string) ([]RepoNode, error) {
	query := BuildGraphQLQueryRepos(user, cursor)

	body, _ := json.Marshal(GraphQLQuery{Query: query})
	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response RepoResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message)
	}

	nodes := response.Data.User.Repositories.Nodes

	// Se houver mais páginas, busca recursivamente
	if response.Data.User.Repositories.PageInfo.HasNextPage {
		nextCursor := response.Data.User.Repositories.PageInfo.EndCursor
		nextNodes, err := FetchAllRepos(user, token, &nextCursor)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, nextNodes...)
	}

	return nodes, nil
}






// BuildGraphQLQueryLite generates a lightweight repositories query.
func BuildGraphQLQueryLite(user string) (string, error) {
	if err := validateUser(user); err != nil {
		return "", err
	}

	query := `
	{
		user(login: "%s") {
			repositories(first: 100, privacy: PUBLIC) {
				nodes {
					name
				}
			}
		}
	}
	`
	return fmt.Sprintf(query, user), nil
}

//////////////////////
// 🧠 Data Models  //
//////////////////////

type ContributionGraphQuery struct {
	Query string `json:"query"`
}

type ContributionDay struct {
	Date              string `json:"date"`
	ContributionCount int    `json:"contributionCount"`
}

type Week struct {
	ContributionDays []ContributionDay `json:"contributionDays"`
}

type ContributionCalendar struct {
	Weeks []Week `json:"weeks"`
}

type ContributionsCollection struct {
	ContributionCalendar ContributionCalendar `json:"contributionCalendar"`
}

type User struct {
	CreatedAt               string                  `json:"createdAt"`
	ContributionsCollection ContributionsCollection `json:"contributionsCollection"`
}

type Data struct {
	User User `json:"user"`
}

type Error struct {
	Message string `json:"message"`
}

type Response struct {
	Data   Data    `json:"data"`
	Errors []Error `json:"errors"`
}

///////////////////////////////
// 🚀 Query Execution Logic  //
///////////////////////////////

// ExecuteGraphQLQuery executes a given GraphQL query with a provided token.
func ExecuteGraphQLQuery(query, token string) (Response, error) {
	var response Response

	body, _ := json.Marshal(ContributionGraphQuery{Query: query})
	client := &http.Client{}

	req, _ := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "GitHub-Readme-Streak-Stats")

	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	if len(response.Errors) > 0 {
		return response, errors.New(response.Errors[0].Message)
	}

	return response, nil
}

////////////////////////////////////
// 👤 Fetch User Information API  //
////////////////////////////////////

type UserInfo struct {
	Name      string `json:"name"`
	Login     string `json:"login"`
	Bio       string `json:"bio"`
	AvatarUrl string `json:"avatarUrl"`
	CreatedAt string `json:"createdAt"`
}

type ResponseInfo struct {
	Data struct {
		UserInfo UserInfo `json:"user"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

// FetchUserData retrieves basic user profile information.
func FetchUserData(user string) (UserInfo, error) {
	token, err := GetGitHubTokenNative()
	if err != nil {
		return UserInfo{}, err
	}

	query := BuildGraphQLQueryUser(user)
	body, _ := json.Marshal(GraphQLQuery{Query: query})

	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(body))
	if err != nil {
		return UserInfo{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return UserInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return UserInfo{}, fmt.Errorf("erro na requisição: status %d", resp.StatusCode)
	}

	var response ResponseInfo
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return UserInfo{}, err
	}

	if len(response.Errors) > 0 {
		return UserInfo{}, errors.New(response.Errors[0].Message)
	}

	return response.Data.UserInfo, nil
}

////////////////////////
// 🌐 HTTP Handlers  //
////////////////////////

// GitUser is an HTTP handler that serves user profile information.
func GitUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("user")
	if username == "" {
		http.Error(w, "Missing 'user' parameter", http.StatusBadRequest)
		return
	}

	userInfo, errUser := user.FetchUserData(username)
	if errUser != nil {
		http.Error(w, fmt.Sprintf("Error retrieving user data: %v", errUser), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user": userInfo,
	})
}
