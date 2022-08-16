package service

import (
	"encoding/json"
	// "log"
	"os"
	// "testing"

	"github.com/hender14/app/model"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
)

type RstSGmail struct {
	Mail    *model.RstMail
	Jsndata []byte
	Request rest.Request
}

type CntSGmail struct {
	Mail    *model.CntMail
	Jsndata []byte
	Request rest.Request
}

func (r *RstSGmail) SGConf(config model.Mails, param model.RstmailPara) (err error) {
	// Mailﾃﾞｰﾀを作成
	r.Mail = &model.RstMail{
		Subject: config.Subject,
		Personalizations: []model.RstPersonal{
			{To: []model.MailUser{{
				Email: config.To.Address,
				Name:  config.To.Express,
			}},
				Parameter: model.RstmailPara{
					Reseturl: param.Reseturl,
					Email:    param.Email,
				},
			},
		},
		From: model.MailUser{
			Email: config.From.Address,
			Name:  config.From.Express,
		},
		TemplateId: os.Getenv("SENDGRID_RST_TEMPLATED_ID"),
	}

	// r.Mail = new(model.RstMail)
	// r.Mail.Subject = config.Subject
	// r.Mail.Personalizations[0].To[0].Email = config.To.Address
	// r.Mail.Personalizations[0].To[0].Name = config.To.Express
	// r.Mail.Personalizations[0].Parameter.Reseturl = param.Reseturl
	// r.Mail.Personalizations[0].Parameter.Email = param.Email
	// r.Mail.From.Email = config.From.Address
	// r.Mail.From.Name = config.From.Express
	// r.Mail.TemplateId = os.Getenv("SENDGRID_RST_TEMPLATED_ID")

	// MailデータをJSON化
	r.Jsndata, err = json.Marshal(&(r.Mail))
	if err != nil {
		return err
	}
	return err
}

func (C *CntSGmail) SGConf(config model.Mails, param model.CntmailPara) (err error) {
	// Mailデータを作成
	C.Mail = &model.CntMail{
		Subject: config.Subject,
		Personalizations: []model.CntPersonal{
			{To: []model.MailUser{{
				Email: config.To.Address,
				Name:  config.To.Express,
			}},
				Parameter: model.CntmailPara{
					ID:      param.ID,
					Title:   param.Title,
					Content: param.Content,
					Email:   param.Email,
				},
			},
		},
		From: model.MailUser{
			Email: config.From.Address,
			Name:  config.From.Express,
		},
		TemplateId: os.Getenv("SENDGRID_CNT_TEMPLATED_ID"),
	}

	// C.Mail.Subject = config.Subject
	// C.Mail.Personalizations[0].To[0].Email = config.To.Address
	// C.Mail.Personalizations[0].To[0].Name = config.To.Express
	// C.Mail.Personalizations[0].Parameter.ID = param.ID
	// C.Mail.Personalizations[0].Parameter.Title = param.Title
	// C.Mail.Personalizations[0].Parameter.Content = param.Content
	// C.Mail.Personalizations[0].Parameter.Email = param.Email
	// C.Mail.From.Email = config.From.Address
	// C.Mail.From.Name = config.From.Express
	// C.Mail.TemplateId = os.Getenv("SENDGRID_CNT_TEMPLATED_ID")

	// MailデータをJSON化
	C.Jsndata, err = json.Marshal(&(C.Mail))
	if err != nil {
		return err
	}
	return err
}

func (r *RstSGmail) Reqconf() (response *rest.Response, err error) {
	// requestのMethodをPostに
	r.Request.Method = "POST"
	r.Request.Body = r.Jsndata
	response, err = sendgrid.API(r.Request)
	// fmt.Println(response.StatusCode)
	// fmt.Println(response.Body)
	// fmt.Println(response.Headers)
	return response, err
}

func (c *CntSGmail) Reqconf() (response *rest.Response, err error) {
	// requestのMethodをPostに
	c.Request.Method = "POST"
	c.Request.Body = c.Jsndata
	response, err = sendgrid.API(c.Request)
	// fmt.Println(response.StatusCode)
	// fmt.Println(response.Body)
	// fmt.Println(response.Headers)
	return response, err
}

func Rst_NewRequest() (r *RstSGmail) {
	r = new(RstSGmail)
	apikey := os.Getenv("SENDGRID_API_KEY")
	// ﾎｽﾄ
	host := "https://api.sendgrid.com"
	// ｴﾝﾄﾞﾎﾟｲﾝﾄ
	endpoint := "/v3/mail/send"
	// API KEYとｴﾝﾄﾞﾎﾟｲﾝﾄ、ﾎｽﾄからrestﾊﾟｯｹｰｼﾞのRequestを生成
	r.Request = sendgrid.GetRequest(apikey, endpoint, host)

	return r
}

func Cnt_NewRequest() (c *CntSGmail) {
	c = new(CntSGmail)
	apikey := os.Getenv("SENDGRID_API_KEY")
	// ﾎｽﾄ
	host := "https://api.sendgrid.com"
	// ｴﾝﾄﾞﾎﾟｲﾝﾄ
	endpoint := "/v3/mail/send"
	// API KEYとｴﾝﾄﾞﾎﾟｲﾝﾄ、ﾎｽﾄからrestﾊﾟｯｹｰｼﾞのRequestを生成
	c.Request = sendgrid.GetRequest(apikey, endpoint, host)

	return c
}

// SendGridｻｰﾊﾞよりﾒｰﾙの送付
func Sendmail(config model.Mails, param model.RstmailPara) (response *rest.Response, err error) {
	R := Rst_NewRequest()
	// SendGridのmail config
	err = R.SGConf(config, param)
	// fmt.Printf("mail: %+v\n", mail)
	if err != nil {
		return nil, err
	}

	// sendGridのAPIにﾘｸｴｽﾄをｾｯﾄ, 戻り値でresponseが返ってくる
	response, err = R.Reqconf()
	if err != nil {
		return nil, err
	}
	return response, err
}

func Cnt_Sendmail(config model.Mails, param model.CntmailPara) (response *rest.Response, err error) {
	C := Cnt_NewRequest()
	// SendGridのmail config
	err = C.SGConf(config, param)
	// fmt.Printf("mail: %+v\n", mail)
	if err != nil {
		return nil, err
	}

	// sendGridのAPIにﾘｸｴｽﾄをｾｯﾄ, 戻り値でresponseが返ってくる
	response, err = C.Reqconf()
	if err != nil {
		return nil, err
	}
	return response, err
}
