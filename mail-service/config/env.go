package config

import "os"

type EnvVariables struct {
	WebPort         string
	AllowedOrigin   string
	MailDomain      string
	MailHost        string
	MailUsername    string
	MailPassword    string
	MailPort        string
	MailEncryption  string
	MailFromName    string
	MailFromAddress string
}

func NewEnvVariables() *EnvVariables {
	var (
		webport         string = os.Getenv("PORT")
		allowedOrigin   string = os.Getenv("ALLOWED_ORIGIN")
		mailDomain      string = os.Getenv("MAIL_DOMAIN")
		mailHost        string = os.Getenv("MAIL_HOST")
		mailUsername    string = os.Getenv("MAIL_USERNAME")
		mailPassword    string = os.Getenv("MAIL_PASSWORD")
		mailPort        string = os.Getenv("MAIL_PORT")
		mailEncryption  string = os.Getenv("MAIL_ENCRYPTION")
		mailFromName    string = os.Getenv("MAIL_FROM_NAME")
		mailFromAddress string = os.Getenv("MAIL_FROM_ADDRESS")
	)
	if webport == "" {
		webport = "80"
	}

	return &EnvVariables{
		WebPort:         webport,
		AllowedOrigin:   allowedOrigin,
		MailDomain:      mailDomain,
		MailHost:        mailHost,
		MailUsername:    mailUsername,
		MailPassword:    mailPassword,
		MailPort:        mailPort,
		MailEncryption:  mailEncryption,
		MailFromName:    mailFromName,
		MailFromAddress: mailFromAddress,
	}
}
