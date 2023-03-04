package ecrenv_test

import (
	"testing"

	"github.com/docker/docker-credential-helpers/credentials"
	"github.com/stretchr/testify/assert"
	ecrenv "github.com/winebarrel/docker-credential-ecr-login-with-env"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	h := &ecrenv.ECREnvHelper{
		ECRLogin: "./mock-helper",
	}

	creds := &credentials.Credentials{
		ServerURL: "ServerURL",
		Username:  "Username",
		Secret:    "Secret",
	}

	err := h.Add(creds)
	assert.NoError(err)
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)

	h := &ecrenv.ECREnvHelper{
		ECRLogin: "./mock-helper",
	}

	err := h.Delete("ServerURL")
	assert.NoError(err)
}

func TestGet(t *testing.T) {
	assert := assert.New(t)

	h := &ecrenv.ECREnvHelper{
		ECRLogin: "./mock-helper",
	}

	username, secret, err := h.Get("ServerURL")
	assert.Equal("Username", username)
	assert.Equal("Secret", secret)
	assert.NoError(err)
}

func TestList(t *testing.T) {
	assert := assert.New(t)

	h := &ecrenv.ECREnvHelper{
		ECRLogin: "./mock-helper",
	}

	m, err := h.List()
	assert.Equal(map[string]string{"ServerURL": "Username"}, m)
	assert.NoError(err)
}
