package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/docker/docker-credential-helpers/credentials"
	ecrenv "github.com/winebarrel/docker-credential-ecr-login-env"
)

var version string

const ConfigFile = "ecr-login-env.json"

func init() {
	credentials.Name = "docker-credential-ecr-login-env"
	credentials.Package = "github.com/winebarrel/docker-credential-ecr-login-env"
	credentials.Version = version
}

func main() {
	home, _ := os.UserHomeDir()
	file, err := os.ReadFile(filepath.Join(home, ".docker", ConfigFile))

	if err != nil {
		log.Fatal(err)
	}

	var envs ecrenv.EnvByServerURL
	err = json.Unmarshal(file, &envs)

	if err != nil {
		log.Fatal(err)
	}

	credentials.Serve(ecrenv.NewECREnvHelper(envs))
}
