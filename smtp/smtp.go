package smtp

import m "sql-to-smtp-service/models"
import "gopkg.in/gomail.v2"

type SMTP interface {
	mailFabric([]*m.Mail)
}

// https://godoc.org/gopkg.in/gomail.v2
// https://stackoverflow.com/questions/24703943/passing-a-slice-into-a-channel
func mailFabric(mails []*m.Mail) <- chan *gomail.Message  {

}
