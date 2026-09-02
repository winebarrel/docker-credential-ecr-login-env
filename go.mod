module github.com/winebarrel/docker-credential-ecr-login-env

go 1.21

toolchain go1.27.1

require (
	github.com/docker/docker-credential-helpers v0.9.9
	github.com/stretchr/testify v1.12.1
)

require go.yaml.in/yaml/v3 v3.0.5 // indirect
