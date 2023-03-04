package main

import (
	"github.com/docker/docker-credential-helpers/credentials"
	ecrenv "github.com/winebarrel/docker-credential-ecr-login-with-env"
)

var version string

func init() {
	credentials.Name = "docker-credential-ecr-login-with-env"
	credentials.Package = "github.com/winebarrel/docker-credential-ecr-login-with-env"
	credentials.Version = version
}

func main() {
	credentials.Serve(ecrenv.NewECREnvHelper())
}
