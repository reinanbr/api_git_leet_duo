package graphql

import (
	"fmt"
)


func BuildGraphQLQueryLang(user string) string {
	// Esta consulta irá pegar o histórico de contribuições completo desde que o usuário entrou no GitHub
	return fmt.Sprintf(`
{
  user(login: "%s") {
    repositories(first: 100) {
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


	`, user)
}




