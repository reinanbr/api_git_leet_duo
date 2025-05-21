package query

import "fmt"

func BuildUserQuery(username string) string {
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
	`, username)
}

func BuildRepoQuery(username string, cursor *string) string {
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
	`, username, after)
}
