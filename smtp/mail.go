package smtp

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	"log"
	"sql-to-smtp-service/models"
	"time"
)

// https://godoc.org/gopkg.in/gomail.v2
// https://stackoverflow.com/questions/24703943/passing-a-slice-into-a-channel

func MailFabric(mails []*models.Mail) <- chan *gomail.Message  {
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

func SendEmail( ch <-chan *gomail.Message) {
		d := gomail.Dialer{
			Host:"5.79.119.5",
			Port: 25,
			SSL: false,
		}

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
					log.Print(err)
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