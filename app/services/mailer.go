package services

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
)

type email struct {
	SenderEmail   string
	ReceiverEmail string
	Subject       string
	Message       string
}

func (m email) SendDraftEmail() error  {
	from := mail.NewEmail("polymail", m.SenderEmail)
	to := mail.NewEmail("polymail", m.ReceiverEmail)
	content := mail.NewSingleEmail(from, m.Subject, to, " ", m.Message)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	if _, err := client.Send(content); err != nil {
		return err
	}
	return nil
}