package config

import "os"

type EnvVariables struct {
	AuthServiceUrl string
	MailServiceUrl string
}

func NewEnvVariables() *EnvVariables {
	return &EnvVariables{
		AuthServiceUrl: os.Getenv("AUTHENTICATION_SERVICE_BASE_URL"),
		MailServiceUrl: os.Getenv("MAIL_SERVICE_BASE_URL"),
	}
}
