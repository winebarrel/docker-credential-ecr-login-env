# docker-credential-ecr-login-with-env

Wrapper command for docker-credential-ecr-login to add environment variables.

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
