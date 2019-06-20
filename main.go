package main

import (
	"github.com/sirupsen/logrus"
	"sql-to-smtp-service/config"
	"sql-to-smtp-service/models"
	"sql-to-smtp-service/smtp"
)
import _ "github.com/go-sql-driver/mysql"

type Env struct {
	db models.Datastore
	smtp smtp.SMTP
}

func main() {
	/*
	// Get connections and configurations
	*/
	db, err := models.NewDB(config.Config.DB)
	defer func() {
		err = db.Close()
		if err != nil {
			logrus.WithError(err).Error("can't close sql db connection")
		}
	}()
	if err != nil {
		logrus.WithError(err).Fatal("can't connection to mssql server")
	}
	sd, err := smtp.NewSMTPClient(config.Config)
	if err != nil {
		logrus.WithError(err).Fatal("cannot set dialer")
	}
	/*
	// Processing mails
	*/
	env := &Env{db, sd }

	mails, err := env.db.AllMails()
	if err != nil {
		logrus.WithError(err).Fatal("error db request")
	}

	env.smtp.SendEmail(env.smtp.MailFabric(mails))
}
