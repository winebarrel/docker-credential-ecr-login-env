# docker-credential-ecr-login-with-env

[![build](https://github.com/winebarrel/docker-credential-ecr-login-with-env/actions/workflows/build.yml/badge.svg)](https://github.com/winebarrel/docker-credential-ecr-login-with-env/actions/workflows/build.yml)

Wrapper command for docker-credential-ecr-login to add environment variables.

## Installation

```sh
brew install winebarrel/docker-credential-ecr-login-with-env/docker-credential-ecr-login-with-env
```

## Usage

First, create `~/.docker/ecr-login-env.json`.

```json
{
  "123456789012.dkr.ecr.us-east-1.amazonaws.com": {
    "AWS_PROFILE": "my-profile"
  }
}
```

Then modify `./.docker/config.json`.

```json
{
  "credHelpers": {
    "123456789012.dkr.ecr.us-east-1.amazonaws.com": "ecr-login-with-env"
  }
}
```
