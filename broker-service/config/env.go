package config

import "os"

type EnvVariables struct {
	AuthServiceUrl string
}

func NewEnvVariables() *EnvVariables {
	return &EnvVariables{
		AuthServiceUrl: os.Getenv("AUTHENTICATION_SERVICE_BASE_URL"),
	}
}
