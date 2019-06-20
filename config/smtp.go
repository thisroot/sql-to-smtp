package config

import "gopkg.in/gomail.v2"

type SMTPConfig struct {
	Host  string
	Port int
}

func (s SMTPConfig) GetSMTPConnectionConfig() *gomail.Dialer  {
	return &gomail.Dialer{
		Host: s.Host,
		Port: s.Port,
		SSL: false,
	}
}
