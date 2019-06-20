package smtp

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	"sql-to-smtp-service/config"
	"sql-to-smtp-service/models"
	"time"
)

// https://godoc.org/gopkg.in/gomail.v2
// https://stackoverflow.com/questions/24703943/passing-a-slice-into-a-channel

type SMTP interface {
	MailFabric(mails []*models.Mail) <- chan *gomail.Message
	SendEmail( ch <-chan *gomail.Message)
}

type SMTPClient struct {
	*gomail.Dialer
}

func NewSMTPClient(config *config.Configuration) (*SMTPClient, error)  {
	logrus.Info("create SMTP dialer: ")
	return &SMTPClient{ config.SMTP }, nil
}

func (SMTPClient) MailFabric (mails []*models.Mail) <- chan *gomail.Message  {
	c := make(chan *gomail.Message)
	go func(c chan *gomail.Message) {
		defer close(c)
		for _, ma := range mails {
				mail := gomail.NewMessage()
				mail.SetHeader("From", ma.FromEmail)
				mail.SetHeader("To", ma.ToEmail)
				mail.SetHeader("Subject", ma.Subject)
				mail.SetBody("text/html", ma.HTML)
				mail.SetBody("text/plain", ma.Plaintext)
				c <- mail
		}
	}(c)
	return c
}

func (d SMTPClient) SendEmail( ch <-chan *gomail.Message) {
		var s gomail.SendCloser
		var err error
		open := false
		for {
			select {
			case m, ok := <-ch:
				logrus.Println(m)
				if !ok {
					logrus.Fatal("empty message")
					return
				}
				if !open {
					if s, err = d.Dial(); err != nil {
						logrus.WithError(err).Fatal("not dial")
						panic(err)
					}
					open = true
				}
				if err := gomail.Send(s, m); err != nil {
					logrus.WithError(err).Fatal("error send")
				}
				logrus.Println("message sended")
			// Close the connection to the SMTP server if no email was sent in
			// the last 30 seconds.
			case <-time.After(30 * time.Second):
				if open {
					if err := s.Close(); err != nil {
						panic(err)
					}
					open = false
				}
			}
		}
}