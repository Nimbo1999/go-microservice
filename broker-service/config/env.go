package config

import "os"

type EnvVariables struct {
	AuthServiceUrl string
	MailServiceUrl string
	LogServiceURL  string
	AMQPUrl        string
	RPCUrl         string
}

func NewEnvVariables() *EnvVariables {
	return &EnvVariables{
		AuthServiceUrl: os.Getenv("AUTHENTICATION_SERVICE_BASE_URL"),
		MailServiceUrl: os.Getenv("MAIL_SERVICE_BASE_URL"),
		LogServiceURL:  os.Getenv("LOG_SERVICE_BASE_URL"),
		AMQPUrl:        os.Getenv("AMQP_URL"),
		RPCUrl:         os.Getenv("RPC_URL"),
	}
}
