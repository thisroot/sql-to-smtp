package models

type Mail struct {
	ID int64 `json:"id"`
	FromEmail string `json:"from_email"`
	ToEmail string `json:"to_email"`
	Subject string `json:"subject"`
	HTML string `json:"html"`
	Plaintext string `json:"html"`
}

func (db *DB) AllMails() ([]*Mail, error) {
	rows, err := db.Query("SELECT  * FROM mail")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	mails := make([]*Mail, 0)
	for rows.Next() {
		mail := new(Mail)
		err := rows.Scan(
			&mail.ID,
			&mail.FromEmail,
			&mail.ToEmail,
			&mail.Subject,
			&mail.HTML,
			&mail.Plaintext)
		if err != nil {
			return nil, err
		}
		mails = append(mails, mail)
	}
	return mails, nil
}
