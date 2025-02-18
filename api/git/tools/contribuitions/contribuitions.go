package contribuitions

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"api_git_leet_duo/api/git/tools/graphql"
	"api_git_leet_duo/api/git/tools/auth"
)


func buildContributionGraphQuery(user string, year int) string {
	start := fmt.Sprintf("%d-01-01T00:00:00Z", year)
	end := fmt.Sprintf("%d-12-31T23:59:59Z", year)
	return fmt.Sprintf(`query { user(login: "%s") { createdAt contributionsCollection(from: "%s", to: "%s") { contributionCalendar { weeks { contributionDays { contributionCount date } } } } } }`, user, start, end)
}



func ExecuteContributionGraphRequests(user string, years []int, tokens []string) (map[int]graphql.Response, error) {
	responses := make(map[int] graphql.Response)
	for _, year := range years {
		token, err := auth.GetGitHubToken(tokens)
		if err != nil {
			return nil, err
		}
		query := buildContributionGraphQuery(user, year)
		response, err := graphql.ExecuteGraphQLQuery(query, token)
		if err != nil {
			return nil, err
		}
		responses[year] = response
	}
	return responses, nil
}





func GetContributionGraphs(user string, startingYear int) (map[int]graphql.Response, error) {
	currentYear := time.Now().Year()
	tokens := auth.GetGitHubTokens()

	// Obter o ano de criação do usuário
	responses, err := ExecuteContributionGraphRequests(user, []int{currentYear}, tokens)
	if err != nil {
		return nil, err
	}

	userCreatedYear := 2005
	if response, exists := responses[currentYear]; exists && response.Data.User.CreatedAt != "" {
		createdAt := response.Data.User.CreatedAt
		if year, err := strconv.Atoi(strings.Split(createdAt, "-")[0]); err == nil {
			userCreatedYear = year
		}
	}

	minYear := startingYear
	if minYear < userCreatedYear {
		minYear = userCreatedYear
	}

	yearsToRequest := []int{}
	for y := minYear; y < currentYear; y++ {
		yearsToRequest = append(yearsToRequest, y)
	}

	moreResponses, err := ExecuteContributionGraphRequests(user, yearsToRequest, tokens)
	if err != nil {
		return nil, err
	}

	for year, resp := range moreResponses {
		responses[year] = resp
	}

	return responses, nil
}
