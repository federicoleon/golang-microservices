package services

import (
	"github.com/federicoleon/golang-microservices/src/api/utils/errors"
	"github.com/federicoleon/golang-microservices/src/api/domain/repositories"
	"github.com/federicoleon/golang-microservices/src/api/domain/github"
	"sync"
	"net/http"
	"github.com/federicoleon/golang-microservices/src/api/log/option_b"
	"github.com/federicoleon/golang-microservices/src/api/providers/github_provider"
	"github.com/federicoleon/golang-microservices/src/api/config"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(clientId string, input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}
	option_b.Info("about to send request to external api",
		option_b.Field("client_id", clientId),
		option_b.Field("status", "pending"),
		option_b.Field("authenticated", clientId != ""))

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		option_b.Error("response obtained from external api", err,
			option_b.Field("client_id", clientId),
			option_b.Field("status", "error"),
			option_b.Field("authenticated", clientId != ""))
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	option_b.Info("response obtained from external api",
		option_b.Field("client_id", clientId),
		option_b.Field("status", "success"),
		option_b.Field("authenticated", clientId != ""))

	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}

func (s *reposService) CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	input := make(chan repositories.CreateRespositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	var wg sync.WaitGroup
	go s.handleRepoResults(&wg, input, output)

	for _, current := range requests {
		wg.Add(1)
		go s.createRepoConcurrent(current, input)
	}

	wg.Wait()
	close(input)

	result := <-output

	successCreations := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCreations ++
		}
	}
	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(requests) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}
	return result, nil
}

func (s *reposService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRespositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse
	for incomingEvent := range input {
		repoResult := repositories.CreateRespositoriesResult{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
	}
	output <- results
}

func (s *reposService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRespositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRespositoriesResult{Error: err}
		return
	}
	result, err := s.CreateRepo("TODO_client_id", input)
	if err != nil {
		output <- repositories.CreateRespositoriesResult{Error: err}
		return
	}
	output <- repositories.CreateRespositoriesResult{Response: result}
}
