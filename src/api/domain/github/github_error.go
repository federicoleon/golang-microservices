package github

type GithubErrorResponse struct {
	StatusCode      int           `json:"status_code"`
	Message         string        `json:"message"`
	DocumentationUr string        `json:"documentation_url"`
	Errors          []GithubError `json:"errors"`
}

func (r GithubErrorResponse) Error() string {
	return r.Message
}

type GithubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
