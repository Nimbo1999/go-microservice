package main

import (
	"bytes"
	"html/template"
	"log"
	"mail-service/config"
	"strconv"
	"time"

	"github.com/vanng822/go-premailer/premailer"
	mail "github.com/xhit/go-simple-mail/v2"
)

type Mail struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string
}

func NewMail(env *config.EnvVariables) *Mail {
	port, err := strconv.Atoi(env.MailPort)
	if err != nil {
		log.Panicln(err)
	}
	return &Mail{
		Domain:      env.MailDomain,
		Host:        env.MailHost,
		Port:        port,
		Username:    env.MailUsername,
		Password:    env.MailPassword,
		Encryption:  env.MailEncryption,
		FromAddress: env.MailFromAddress,
		FromName:    env.MailFromName,
	}
}

type Message struct {
	From         string
	FromName     string
	To           string
	Subject      string
	Attachements []string
	Data         any
	DataMap      map[string]any
}

func (m *Mail) SendSMTPMessage(msg Message) error {
	if msg.From == "" {
		msg.From = m.FromAddress
	}
	if msg.FromName == "" {
		msg.FromName = m.FromName
	}
	data := map[string]any{
		"message": msg.Data,
		"subject": msg.Subject,
	}
	msg.DataMap = data
	htmlMessage, err := m.BuildHTMLMessage(msg)
	if err != nil {
		log.Println("[ERROR]:", err)
		return err
	}
	plainMessage, err := m.BuildPlainTextMessage(msg)
	if err != nil {
		log.Println("[ERROR]:", err)
		return err
	}

	server := mail.NewSMTPClient()
	server.Host = m.Host
	server.Port = m.Port
	server.Username = m.Username
	server.Password = m.Password
	server.Encryption = m.getEncryption(m.Encryption)
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		log.Println("[ERROR]:", err)
		return err
	}
	email := mail.NewMSG()
	email.SetFrom(msg.From)
	email.SetSubject(msg.Subject)
	email.AddTo(msg.To)
	email.SetBody(mail.TextPlain, plainMessage)
	email.AddAlternative(mail.TextHTML, htmlMessage)
	if len(msg.Attachements) > 0 {
		for _, x := range msg.Attachements {
			email.AddAttachment(x)
		}
	}
	return email.Send(smtpClient)
}

func (m *Mail) BuildHTMLMessage(msg Message) (string, error) {
	t, err := template.New("email-template").ParseFiles("./templates/mail.html.gohtml")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	formattedMessage := tpl.String()
	formattedMessage, err = m.inlineCSS(formattedMessage)
	if err != nil {
		return "", err
	}
	return formattedMessage, nil
}

func (m *Mail) BuildPlainTextMessage(msg Message) (string, error) {
	t, err := template.New("email-plain").ParseFiles("./templates/mail.plain.gohtml")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	plainMessage := tpl.String()
	return plainMessage, nil
}

func (m *Mail) inlineCSS(doc string) (string, error) {
	options := &premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(doc, options)
	if err != nil {
		return "", err
	}
	html, err := prem.Transform()
	if err != nil {
		return "", err
	}
	return html, nil
}

func (m *Mail) getEncryption(encryption string) mail.Encryption {
	switch encryption {
	case "ssl":
		return mail.EncryptionSSLTLS
	case "none", "":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}
