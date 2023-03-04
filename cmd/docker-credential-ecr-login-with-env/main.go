package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/docker/docker-credential-helpers/credentials"
	ecrenv "github.com/winebarrel/docker-credential-ecr-login-with-env"
)

var version string

const ConfigFile = "ecr-login-env.toml"

func init() {
	credentials.Name = "docker-credential-ecr-login-with-env"
	credentials.Package = "github.com/winebarrel/docker-credential-ecr-login-with-env"
	credentials.Version = version
}

func main() {
	home, _ := os.UserHomeDir()
	var envs ecrenv.EnvByServerURL
	_, err := toml.DecodeFile(filepath.Join(home, ".docker", ConfigFile), &envs)

	if err != nil {
		log.Fatal(err)
	}

	credentials.Serve(ecrenv.NewECREnvHelper(envs))
}
