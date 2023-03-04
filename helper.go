package ecrenv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/docker/docker-credential-helpers/credentials"
)

const ECRLoginCmd = "docker-credential-ecr-login"

type ECREnvHelper struct {
	ECRLogin string
}

func NewECREnvHelper() *ECREnvHelper {
	return &ECREnvHelper{
		ECRLogin: ECRLoginCmd,
	}
}

func (h *ECREnvHelper) Add(creds *credentials.Credentials) error {
	cmd := exec.Command(h.ECRLogin, "store")
	buf, err := json.Marshal(creds)

	if err != nil {
		return err
	}

	cmd.Stdin = bytes.NewBuffer(buf)
	return cmd.Run()
}

func (h *ECREnvHelper) Delete(serverURL string) error {
	cmd := exec.Command(h.ECRLogin, "erase")
	cmd.Stdin = bytes.NewBufferString(serverURL)
	return cmd.Run()
}

func (h *ECREnvHelper) Get(serverURL string) (string, string, error) {
	cmd := exec.Command(h.ECRLogin, "get")
	cmd.Stdin = bytes.NewBufferString(serverURL)
	out := &bytes.Buffer{}
	cmd.Stdout = out
	err := cmd.Run()

	if err != nil {
		fmt.Println("koko")
		return "", "", err
	}

	var creds credentials.Credentials
	err = json.Unmarshal(out.Bytes(), &creds)

	if err != nil {
		return "", "", err
	}

	return creds.Username, creds.Secret, nil
}

func (h *ECREnvHelper) List() (map[string]string, error) {
	out, err := exec.Command(h.ECRLogin, "list").Output()

	if err != nil {
		return nil, err
	}

	m := map[string]string{}
	err = json.Unmarshal(out, &m)

	if err != nil {
		return nil, err
	}

	return m, nil
}
