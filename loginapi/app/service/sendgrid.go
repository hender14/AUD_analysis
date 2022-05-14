package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"app/model"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
)

// SendGridｻｰﾊﾞよりﾒｰﾙの送付
func Sendmail(config model.Mails, url string) (response *rest.Response, err error) {
	apikey := os.Getenv("SENDGRID_API_KEY")
	// ﾎｽﾄ
	host := "https://api.sendgrid.com"
	// ｴﾝﾄﾞﾎﾟｲﾝﾄ
	endpoint := "/v3/mail/send"
	templateId := os.Getenv("SENDGRID_TEMPLATED_ID")
	reseturl := url
	email := os.Getenv("SENDGRID_FROM_EMAIL")
	// API KEYとｴﾝﾄﾞﾎﾟｲﾝﾄ、ﾎｽﾄからrestﾊﾟｯｹｰｼﾞのRequestを生成
	request := sendgrid.GetRequest(apikey, endpoint, host)
	// requestのMethodをPostに
	request.Method = "POST"
  // ﾒｰﾙの内容をJSONで作成する
	subject := config.Subject

	mail := model.Mail{
		Subject: subject,
		Personalizations: []model.Personalizations{
			{To: []model.MailUser{{
				Email: config.To.Address,
				Name:  config.To.Express,
			}},
			Parameter: model.TemplatePara{
				Reseturl: reseturl,
				Email: email,
			},
			},
		},
		From: model.MailUser{
			Email: config.From.Address,
			Name:  config.From.Express,
		},
		TemplateId: templateId,
		// Content: []model.Contents{{
		// 	// Type:  "text/plain",
		// 	Type:  "text/html",
		// 	Value: config.Text.Htmltext,
		// }},
	}
	// GoのｺｰﾄﾞをJSON化
	data, err := json.Marshal(mail)

	log.Println(string(data))
	if err != nil {
		log.Println(err)
	}
	// JSON化したmailの内容をrequest.Bodyに代入
	request.Body = data

	// sendGridのAPIにﾘｸｴｽﾄをｾｯﾄ, 戻り値でresponseが返ってくる
	response, err = sendgrid.API(request)
	if err != nil {
			log.Println(err)
	} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)
	}

	return response, err
}