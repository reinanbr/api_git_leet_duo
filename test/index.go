package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type GraphQLQuery struct {
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

type Language struct {
	Name string `json:"name"`
}

type LanguageEdge struct {
	Size int      `json:"size"`
	Node Language `json:"node"`
}

type Repository struct {
	Languages struct {
		Edges []LanguageEdge `json:"edges"`
	} `json:"languages"`
}

type User struct {
	Name                    string `json:"name"`
	Login                   string `json:"login"`
	Bio                     string `json:"bio"`
	AvatarUrl               string `json:"avatarUrl"`
	CreatedAt               string `json:"createdAt"`
	ContributionsCollection struct {
		ContributionCalendar ContributionCalendar `json:"contributionCalendar"`
	} `json:"contributionsCollection"`
	Repositories struct {
		Nodes []Repository `json:"nodes"`
	} `json:"repositories"`
}

type Response struct {
	Data struct {
		User User `json:"user"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

func getGitHubToken() (string, error) {
	token := os.Getenv("TOKEN")
	if token == "" {
		return "", errors.New("GitHub token não encontrado")
	}
	return token, nil
}

func buildGraphQLQuery(user string) string {
	// Esta consulta irá pegar o histórico de contribuições completo desde que o usuário entrou no GitHub
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
}
	`, user)
}

/*
contributionsCollection(from: "2018-01-01T00:00:00Z", to: "2018-12-31T23:59:59Z") {
      contributionCalendar {
        weeks {
          contributionDays {
            date
            contributionCount
          }
        }
      }
    }
    repositories(first: 100) {
      pageInfo {
        hasNextPage
        endCursor
      }
      nodes {
        languages(first: 100) {
          edges {
            size
            node {
              name
            }
          }
        }
      }
*/


func fetchUserData(user string) (User, error) {
	token, err := getGitHubToken()
	if err != nil {
		return User{}, err
	}

	query := buildGraphQLQuery(user)
	body, _ := json.Marshal(GraphQLQuery{Query: query})

	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(body))
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return User{}, fmt.Errorf("erro na requisição: status %d", resp.StatusCode)
	}

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return User{}, err
	}

	if len(response.Errors) > 0 {
		return User{}, errors.New(response.Errors[0].Message)
	}

	return response.Data.User, nil
}

func Git(w http.ResponseWriter, r *http.Request) {
	// Obtém o nome de usuário do parâmetro de consulta (query parameter)
	user := r.URL.Query().Get("user")
	if user == "" {
		http.Error(w, "Parâmetro 'user' é obrigatório", http.StatusBadRequest)
		return
	}

	// Busca os dados do usuário do GitHub
	data, err := fetchUserData(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao buscar dados do GitHub: %s", err), http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho como JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Codifica os dados do usuário em JSON e retorna na resposta
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Erro ao codificar os dados", http.StatusInternalServerError)
	}
}
