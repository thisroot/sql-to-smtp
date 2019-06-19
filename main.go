package main

import (
	"github.com/sirupsen/logrus"
	"github.com/thisroot/go_lib/promclient"
	"sql-to-smtp-service/config"
	"sql-to-smtp-service/models"
	"sql-to-smtp-service/smtp"
)
import _ "github.com/go-sql-driver/mysql"

type Env struct {
	db models.Datastore
}

func main() {
	db, err := models.NewDB(config.Config.DB)
	defer func() {
		err = db.Close()
		if err != nil {
			promclient.IncError("sql_close_connection")
			logrus.WithError(err).Error("can't close sql db connection")
		}
	}()
	if err != nil {
		promclient.IncError("sql_set_connection")
		logrus.WithError(err).Fatal("can't connection to mssql server")
	}
	env := &Env{db }
	mails, err := env.db.AllMails()
	if err != nil {
		logrus.WithError(err).Fatal("error db request")
	}

	ch := smtp.MailFabric(mails)
	smtp.SendEmail(ch)
}
