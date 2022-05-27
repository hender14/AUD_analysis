package service

import (
	"encoding/json"
	"log"
	"os"

	"app/model"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
)

// SendGridｻｰﾊﾞよりﾒｰﾙの送付
func Sendmail(config model.Mails, param map[string]string, state string) (response *rest.Response, err error) {
	var mail any
	// SendGridのmail config
	if (state == "reset") {
		mail = RstSGConfig(config, param)
	} else {
		mail = CntSGConfig(config, param)
	}
	// fmt.Printf("mail: %+v\n", mail)
	// GoのｺｰﾄﾞをJSON化
	jsondata, err := json.Marshal(mail)
	if err != nil {
		log.Println(err)
	}

	// SendGridのrequest config
	request := ReqConfig()
	// JSON化したmailの内容をrequest.Bodyに代入
	request.Body = jsondata

	// sendGridのAPIにﾘｸｴｽﾄをｾｯﾄ, 戻り値でresponseが返ってくる
	response, err = sendgrid.API(request)
	if err != nil {
			log.Println(err)
	} else {
			// fmt.Println(response.StatusCode)
			// fmt.Println(response.Body)
			// fmt.Println(response.Headers)
	}
	return response, err
}

func ReqConfig() ( rest.Request ) {
	apikey := os.Getenv("SENDGRID_API_KEY")
	// ﾎｽﾄ
	host := "https://api.sendgrid.com"
	// ｴﾝﾄﾞﾎﾟｲﾝﾄ
	endpoint := "/v3/mail/send"
	// API KEYとｴﾝﾄﾞﾎﾟｲﾝﾄ、ﾎｽﾄからrestﾊﾟｯｹｰｼﾞのRequestを生成
	request := sendgrid.GetRequest(apikey, endpoint, host)
	// requestのMethodをPostに
	request.Method = "POST"

	return request
}

func RstSGConfig(config model.Mails, param map[string]string) ( model.RstMail ) {
	templateId := os.Getenv("SENDGRID_RST_TEMPLATED_ID")
  // ﾒｰﾙの内容をJSONで作成する
	subject := config.Subject

	mail := model.RstMail{
		Subject: subject,
		Personalizations: []model.RstPersonal{
			{To: []model.MailUser{{
				Email: config.To.Address,
				Name:  config.To.Express,
			}},
			Parameter: model.RstmailPara{
				Reseturl: param["reseturl"],
				Email: param["email"],
			},
			},
		},
		From: model.MailUser{
			Email: config.From.Address,
			Name:  config.From.Express,
		},
		TemplateId: templateId,
	}
	return mail
}

func CntSGConfig(config model.Mails, param map[string]string) ( model.CntMail ) {
	templateId := os.Getenv("SENDGRID_CNT_TEMPLATED_ID")
  // ﾒｰﾙの内容をJSONで作成する
	subject := config.Subject

	mail := model.CntMail{
		Subject: subject,
		Personalizations: []model.CntPersonal{
			{To: []model.MailUser{{
				Email: config.To.Address,
				Name:  config.To.Express,
			}},
			Parameter: model.CntmailPara{
				ID: param["ID"],
				Title: param["title"],
				Content: param["content"],
				Email: param["email"],
			},
			},
		},
		From: model.MailUser{
			Email: config.From.Address,
			Name:  config.From.Express,
		},
		TemplateId: templateId,
	}
	return mail
}