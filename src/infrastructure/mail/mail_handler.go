package mail

import (
	"errors"
	"fmt"
	"os"

	"github.com/gaku3601/clean-blog/src/usecase"
	"github.com/sendgrid/rest"
	sendgrid "github.com/sendgrid/sendgrid-go"
)

type MailHandler struct {
	Request rest.Request
}

func NewMailHandler() usecase.UserMail {
	m := new(MailHandler)
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	endpoint := "/v3/mail/send"
	m.Request = sendgrid.GetRequest(apiKey, endpoint, host)
	m.Request.Method = "POST"

	return m
}

func (m *MailHandler) SendConfirmValidEmail(email string, token string) (err error) {
	data := fmt.Sprintf(`
	{
		"personalizations": [
			{
				"to": [
					{
						"email": "%s"
					}
				],
				"dynamic_template_data": {
					"url": "%s",
				}
			}
		],
		"from": {
			"email": "mail@b-body.xyz",
			"name": "b-body"
		},
		"template_id":"d-aba4554d7e4e40a79ad7c62773ee20ef",
	}	
	`, email, token)
	m.Request.Body = []byte(data)

	res, err := sendgrid.API(m.Request)
	if err != nil {
		return
	}
	if res.StatusCode >= 400 {
		return errors.New(res.Body)
	}

	return
}
func (m *MailHandler) SendForgotPasswordMail(email string, token string) (err error) {
	return
}
