# API Git Leet Duo

A comprehensive API that provides information from Duolingo, LeetCode, and GitHub platforms.

## Overview

This API allows you to fetch user data from three popular platforms:
- **GitHub**: User profiles, repositories, languages, streaks, and commit history
- **LeetCode**: User statistics, submission history, and problem-solving data
- **Duolingo**: User profiles, course progress, and language learning statistics

## Base URL

**Production:**
```
https://api-git-leet-duo.vercel.app
```

**Local Development:**
```
http://localhost:8080
```

## Quick Start

1. **Clone and Setup:**
   ```bash
   git clone <repository-url>
   cd api_git_leet_duo
   go mod download
   ```

2. **Configure GitHub Token:**
   ```bash
   # Copy example environment file
   cp env.example .env
   
   # Edit .env and add your GitHub token
   nano .env
   ```

3. **Run the Server:**
   ```bash
   go run main.go
   ```

4. **Test the API:**
   - Open http://localhost:8080 for the interactive documentation
   - Or test directly: `curl "http://localhost:8080/api/git/user?user=reinanbr"`

**Production Examples:**
```bash
# GitHub User Profile
curl "https://api-git-leet-duo.vercel.app/api/git/user?user=reinanbr"

# GitHub Repositories
curl "https://api-git-leet-duo.vercel.app/api/git/repos?user=reinanbr"

# LeetCode User Stats
curl "https://api-git-leet-duo.vercel.app/api/leet/user?user=reinanbr"

# Duolingo User Profile
curl "https://api-git-leet-duo.vercel.app/api/duo/user?user=reinanbr"
```

## Usage Examples

### JavaScript/Node.js
```javascript
// GitHub User Profile
const response = await fetch('https://api-git-leet-duo.vercel.app/api/git/user?user=reinanbr');
const data = await response.json();
console.log(data);

// LeetCode User Stats
const leetResponse = await fetch('https://api-git-leet-duo.vercel.app/api/leet/user?user=reinanbr');
const leetData = await leetResponse.json();
console.log(leetData);
```

### Python
```python
import requests

# GitHub User Profile
response = requests.get('https://api-git-leet-duo.vercel.app/api/git/user?user=reinanbr')
data = response.json()
print(data)

# Duolingo User Profile
duo_response = requests.get('https://api-git-leet-duo.vercel.app/api/duo/user?user=reinanbr')
duo_data = duo_response.json()
print(duo_data)
```

### cURL
```bash
# GitHub User Profile
curl "https://api-git-leet-duo.vercel.app/api/git/user?user=reinanbr"

# GitHub Repositories with pretty JSON
curl "https://api-git-leet-duo.vercel.app/api/git/repos?user=reinanbr" | jq

# LeetCode User Stats
curl "https://api-git-leet-duo.vercel.app/api/leet/user?user=reinanbr"

# Duolingo User Profile
curl "https://api-git-leet-duo.vercel.app/api/duo/user?user=reinanbr"
```

### PHP
```php
<?php
// GitHub User Profile
$response = file_get_contents('https://api-git-leet-duo.vercel.app/api/git/user?user=reinanbr');
$data = json_decode($response, true);
print_r($data);

// LeetCode User Stats
$leetResponse = file_get_contents('https://api-git-leet-duo.vercel.app/api/leet/user?user=reinanbr');
$leetData = json_decode($leetResponse, true);
print_r($leetData);
?>
```

## API Endpoints

### GitHub API

#### Get User Profile
Retrieves basic user information from GitHub.

**Endpoint:** `GET /git/user`

**Query Parameters:**
- `user` (required): GitHub username

**Example Request:**
```
GET /git/user?user=reinanbr
```

**Example Response:**
```json
{
  "user": {
    "name": "Reinan Bezerra",
    "login": "reinanbr",
    "bio": "Physics Student | Open Source | Data Science | FreeLancer",
    "avatarUrl": "https://avatars.githubusercontent.com/u/44844786?u=8a667b0d67dcc800f4420bf89869b27abd719c78&v=4",
    "createdAt": "2018-11-07T16:31:43Z"
  }
}
```

#### Get User Repositories
Retrieves user's repositories with language information.

**Endpoint:** `GET /git/repos`

**Query Parameters:**
- `user` (required): GitHub username

**Example Request:**
```
GET /git/repos?user=reinanbr
```

**Example Response:**
```json
{
  "repositories": [
    {
      "name": "chat",
      "createdAt": "2019-07-25T22:59:26Z",
      "languages": {
        "edges": [
          {
            "size": 5589,
            "node": {
              "name": "JavaScript"
            }
          },
          {
            "size": 917,
            "node": {
              "name": "HTML"
            }
          },
          {
            "size": 363,
            "node": {
              "name": "CSS"
            }
          },
          {
            "size": 44,
            "node": {
              "name": "Shell"
            }
          }
        ]
      }
    }
  ]
}
```

#### Get User Languages
Retrieves programming languages used by the user with percentage distribution.

**Endpoint:** `GET /git/langs`

**Query Parameters:**
- `user` (required): GitHub username

**Example Request:**
```
GET /git/langs?user=reinanbr
```

**Example Response:**
```json
{
  "languages": [
    {
      "Lang": "HTML",
      "Percentage": 26.343612318632452
    },
    {
      "Lang": "JavaScript",
      "Percentage": 15.25918117420442
    },
    {
      "Lang": "Python",
      "Percentage": 13.193957220007013
    }
  ],
  "total_bytes": 5448949,
  "user": "reinanbr"
}
```

#### Get User Streak
Retrieves user's contribution streak information.

**Endpoint:** `GET /git/streak`

**Query Parameters:**
- `user` (required): GitHub username

**Example Request:**
```
GET /git/streak?user=reinanbr
```

**Example Response:**
```json
{
  "streak": {
    "current_streak": 0,
    "max_streak": 46
  },
  "user": "reinanbr"
}
```

#### Get User Commits
Retrieves detailed commit history and contribution calendar.

**Endpoint:** `GET /git/commit`

**Query Parameters:**
- `user` (required): GitHub username

**Example Request:**
```
GET /git/commit?user=reinanbr
```

**Example Response:**
```json
{
  "commit": {
    "2018": {
      "data": {
        "user": {
          "createdAt": "2018-11-07T16:31:43Z",
          "contributionsCollection": {
            "contributionCalendar": {
              "weeks": [
                {
                  "contributionDays": [
                    {
                      "date": "2018-01-01",
                      "contributionCount": 0
                    }
                  ]
                }
              ]
            }
          }
        }
      }
    }
  }
}
```

### LeetCode API

#### Get User Profile
Retrieves comprehensive user statistics from LeetCode.

**Endpoint:** `GET /leet/user`

**Query Parameters:**
- `user` (required): LeetCode username

**Example Request:**
```
GET /leet/user?user=reinanbr
```

**Example Response:**
```json
{
  "data": {
    "allQuestionsCount": [
      {
        "difficulty": "All",
        "count": 3601
      },
      {
        "difficulty": "Easy",
        "count": 883
      },
      {
        "difficulty": "Medium",
        "count": 1872
      },
      {
        "difficulty": "Hard",
        "count": 846
      }
    ],
    "matchedUser": {
      "username": "reinanbr",
      "firstName": "Reinan",
      "lastName": "Br.",
      "contributions": {
        "Points": 125
      },
      "profile": {
        "reputation": 1,
        "ranking": 3502512,
        "userAvatar": "https://assets.leetcode.com/users/avatars/avatar_1691689536.png"
      },
      "submissionCalendar": "{\"1736467200\": 3, \"1736553600\": 1}",
      "submitStats": {
        "acSubmissionNum": [
          {
            "difficulty": "All",
            "count": 18,
            "submissions": 24
          },
          {
            "difficulty": "Easy",
            "count": 16,
            "submissions": 22
          },
          {
            "difficulty": "Medium",
            "count": 2,
            "submissions": 2
          },
          {
            "difficulty": "Hard",
            "count": 0,
            "submissions": 0
          }
        ]
      },
      "streak": {
        "totalSubmissions": 17,
        "currentStreak": 0,
        "longestStreak": 4
      }
    },
    "recentSubmissionList": [
      {
        "title": "Add Binary",
        "titleSlug": "add-binary",
        "timestamp": "1747573697",
        "statusDisplay": "Accepted",
        "lang": "python3"
      }
    ]
  }
}
```

### Duolingo API

#### Get User Profile
Retrieves user profile and course progress from Duolingo.

**Endpoint:** `GET /duo/user`

**Query Parameters:**
- `user` (required): Duolingo username

**Example Request:**
```
GET /duo/user?user=reinanbr
```

**Example Response:**
```json
{
  "username": "REInanBR",
  "name": "REInan",
  "firstName": "",
  "lastName": "",
  "bio": "",
  "picture": "//simg-ssl.duolingo.com/avatar/default_2",
  "creationDate": 1454870245,
  "streak": 0,
  "totalXp": 5596,
  "courses": [
    {
      "title": "French",
      "learningLanguage": "fr",
      "fromLanguage": "en",
      "xp": 151,
      "crowns": 9999,
      "id": "DUOLINGO_FR_EN"
    },
    {
      "title": "English",
      "learningLanguage": "en",
      "fromLanguage": "pt",
      "xp": 2375,
      "crowns": 9999,
      "id": "DUOLINGO_EN_PT"
    }
  ],
  "streakData": {
    "currentStreak": {
      "startDate": "",
      "length": 0,
      "endDate": ""
    }
  },
  "xp_by_language": [
    {
      "language_id": "Spanish",
      "percentage": 43.316654753395284
    },
    {
      "language_id": "English",
      "percentage": 42.44102930664761
    }
  ]
}
```

## Error Responses

All endpoints may return the following error responses:

### 400 Bad Request
When required parameters are missing or invalid.

```json
{
  "error": "Missing required parameter: user"
}
```

### 404 Not Found
When the requested user is not found.

```json
{
  "error": "User not found"
}
```

### 500 Internal Server Error
When there's an internal server error.

```json
{
  "error": "Internal server error"
}
```

## Rate Limiting

Please be mindful of rate limits when making requests to external APIs:
- GitHub API: 5000 requests per hour for authenticated requests
- LeetCode: No official rate limit documented
- Duolingo: No official rate limit documented

## Configuration

### API Tokens Setup

This API requires authentication tokens for some endpoints to work properly. Follow the instructions below to configure your tokens:

#### GitHub Token
The GitHub API endpoints require a personal access token for authenticated requests.

1. **Create a GitHub Personal Access Token:**
   - Go to [GitHub Settings > Developer settings > Personal access tokens](https://github.com/settings/tokens)
   - Click "Generate new token (classic)"
   - Select the following scopes:
     - `public_repo` (for public repository access)
     - `read:user` (for user profile information)
     - `read:org` (for organization data if needed)
   - Copy the generated token

2. **Configure the Token:**
   
   **For Local Development:**
   ```bash
   export GITHUB_TOKEN="your_github_token_here"
   ```
   
   **For Vercel Deployment:**
   - Go to your Vercel project dashboard
   - Navigate to Settings > Environment Variables
   - Add `GITHUB_TOKEN` with your token value
   
   **Using .env file (optional):**
   ```bash
   # Copy the example file and edit it
   cp env.example .env
   # Edit .env and add your actual token
   nano .env
   ```

#### LeetCode Token (Optional)
LeetCode API doesn't require authentication for public user data, but you may need to handle rate limiting.

#### Duolingo Token (Optional)
Duolingo API doesn't require authentication for public user data.

### Environment Variables

The following environment variables can be configured:

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| `GITHUB_TOKEN` | GitHub Personal Access Token | Yes (for GitHub APIs) | None |
| `PORT` | Server port | No | 8080 |
| `ENVIRONMENT` | Environment (development/production) | No | development |

### Token Security

⚠️ **Important Security Notes:**
- Never commit tokens to version control
- Use environment variables for sensitive data
- Rotate tokens regularly
- Use the minimum required permissions for each token
- Consider using GitHub Apps instead of personal access tokens for production use

## Development

### Prerequisites
- Go 1.19 or higher

### Installation
```bash
git clone <repository-url>
cd api_git_leet_duo
go mod download
```

### Running Locally
```bash
go run main.go
```

The server will start on `http://localhost:8080`

### Building
```bash
go build -o api_git_leet_duo main.go
```

## Deployment

This API can be deployed to various platforms:

### Vercel
The project includes a `vercel.json` configuration file for easy deployment to Vercel.

### Docker
```dockerfile
FROM golang:1.19-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Troubleshooting

### Common Issues

#### "GitHub token não encontrado" Error
This error occurs when the GitHub API endpoints are called without a valid token.

**Solution:**
1. Ensure you have set the `GITHUB_TOKEN` environment variable
2. Verify the token has the correct permissions
3. Check if the token is valid and not expired

```bash
# Test your token
curl -H "Authorization: token YOUR_TOKEN" https://api.github.com/user
```

#### Rate Limiting Issues
If you encounter rate limiting errors:

**For GitHub:**
- Use authenticated requests (with token) for higher limits
- Implement exponential backoff in your requests
- Consider using GitHub Apps for higher rate limits

**For LeetCode/Duolingo:**
- Add delays between requests
- Implement request caching

#### CORS Issues
If you're getting CORS errors when testing from a browser:

**Solution:**
The API includes CORS headers in the `vercel.json` configuration. For local development, you may need to add CORS middleware to your Go server.

### Getting Help

If you encounter any issues or have questions, please:

1. Check the [troubleshooting section](#troubleshooting) above
2. Search existing [issues](https://github.com/your-repo/issues)
3. Open a new issue with:
   - Description of the problem
   - Steps to reproduce
   - Expected vs actual behavior
   - Environment details (OS, Go version, etc.)

## Support

If you encounter any issues or have questions, please open an issue on the GitHub repository.
