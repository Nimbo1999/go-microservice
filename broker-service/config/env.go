package config

import "os"

type EnvVariables struct {
	AuthServiceUrl string
	MailServiceUrl string
	LogServiceURL  string
	AMQP_URL       string
}

func NewEnvVariables() *EnvVariables {
	return &EnvVariables{
		AuthServiceUrl: os.Getenv("AUTHENTICATION_SERVICE_BASE_URL"),
		MailServiceUrl: os.Getenv("MAIL_SERVICE_BASE_URL"),
		LogServiceURL:  os.Getenv("LOG_SERVICE_BASE_URL"),
		AMQP_URL:       os.Getenv("AMQP_URL"),
	}
}
