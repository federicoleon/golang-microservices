package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "SECRET_GITHUB_ACCESS_TOKEN", secretGithubAccessToken)
}

func TestGetGithubAccessToken(t *testing.T) {
	assert.EqualValues(t, "", GetGithubAccessToken())
}