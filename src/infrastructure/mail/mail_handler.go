package mail

import (
	"errors"
	"fmt"
	"os"

	"github.com/gaku3601/clean-blog/src/usecase"
	"github.com/sendgrid/rest"
	sendgrid "github.com/sendgrid/sendgrid-go"
)

// MailHandler ハンドラー
type MailHandler struct {
	Request rest.Request
}

// NewMailHandler mailハンドラーを作成
func NewMailHandler() usecase.UserMail {
	m := new(MailHandler)
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	endpoint := "/v3/mail/send"
	m.Request = sendgrid.GetRequest(apiKey, endpoint, host)
	m.Request.Method = "POST"

	return m
}

// SendConfirmValidEmail email有効化用のmailを送信する。
func (m *MailHandler) SendConfirmValidEmail(email string, token string) (err error) {
	url, err := createConfirmValidEmailURL(token)
	if err != nil {
		return err
	}
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
	`, email, url)
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

// SendForgotPasswordMail パスワードを忘れた際のメールを発行する。
func (m *MailHandler) SendForgotPasswordMail(email string, token string) (err error) {
	url, err := createForgotPasswordURL(token)
	if err != nil {
		return err
	}
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
	`, email, url)
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

// createConfirmValidEmailURL email有効化用のurlを発行します。
func createConfirmValidEmailURL(token string) (url string, err error) {
	env := os.Getenv("FRONTHOST")
	if env == "" {
		return "", errors.New("$FRONTHOST環境変数を設定してください。")
	}
	url = env + "validemail/" + token
	return url, nil
}

// createForgotPasswordURL forgat password用のURLを発行します。
func createForgotPasswordURL(token string) (url string, err error) {
	env := os.Getenv("FRONTHOST")
	if env == "" {
		return "", errors.New("$FRONTHOST環境変数を設定してください。")
	}
	url = env + "forgotpassword/" + token
	return url, nil
}
