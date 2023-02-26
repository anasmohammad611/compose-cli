package errors

import "errors"

var ErrInvalidDockerComposeString = errors.New("Invalid docker-compose string")

var ErrInvalidYaml = errors.New("Invalid yaml file")
