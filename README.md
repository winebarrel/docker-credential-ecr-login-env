# docker-credential-ecr-login-env

[![build](https://github.com/winebarrel/docker-credential-ecr-login-env/actions/workflows/build.yml/badge.svg)](https://github.com/winebarrel/docker-credential-ecr-login-env/actions/workflows/build.yml)

Wrapper command for docker-credential-ecr-login to add environment variables.

cf. https://github.com/awslabs/amazon-ecr-credential-helper

## Installation

```sh
brew install winebarrel/docker-credential-ecr-login-env/docker-credential-ecr-login-env

# NOTE: docker-credential-ecr-login must also be installed
# see https://github.com/awslabs/amazon-ecr-credential-helper#installing
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

Then modify `~/.docker/config.json`.

```json
{
  "credHelpers": {
    "123456789012.dkr.ecr.us-east-1.amazonaws.com": "ecr-login-env"
  }
}
```
