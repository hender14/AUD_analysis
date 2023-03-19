package gateway

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hender14/app/domain"
	"github.com/sendgrid/rest"
)

// requert.Bodyに格納するJSONの元となるメールを表す構造体
type RstMail struct {
	Subject          string        `json:"subject"`
	Personalizations []RstPersonal `json:"personalizations"`
	From             MailUser      `json:"from"`
	TemplateId       string        `json:"template_id"`
}

// 封筒のようなもの
// メールのメタデータを表す構造体
type RstPersonal struct {
	To        []MailUser  `json:"to"`
	Parameter RstmailPara `json:"dynamic_template_data"`
}

// メールのユーザーを表す構造体
type MailUser struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// メールの変数情報
type RstmailPara struct {
	Reseturl string `json:"reseturl"`
	Email    string `json:"email"`
}

// requert.Bodyに格納するJSONの元となるメールを表す構造体
type CntMail struct {
	Subject          string        `json:"subject"`
	Personalizations []CntPersonal `json:"personalizations"`
	From             MailUser      `json:"from"`
	TemplateId       string        `json:"template_id"`
}

// メールのメタデータを表す構造体
type CntPersonal struct {
	To        []MailUser         `json:"to"`
	Parameter domain.CntmailPara `json:"dynamic_template_data"`
}

// SendGridｻｰﾊﾞよりﾒｰﾙの送付
func (repo *UserRepository) Sendmail(config *domain.Mails, param *domain.RstmailPara) (response *rest.Response, err error) {
	// SendGridのmail config
	data, err := sgConf_Rst(config, param)
	// fmt.Printf("mail: %+v\n", mail)
	if err != nil {
		fmt.Printf("SGconf result: %s/n", data)
		fmt.Printf("err result: %s/n", err)
		return response, err
	}

	// sendGridのAPIにﾘｸｴｽﾄをｾｯﾄ, 戻り値でresponseが返ってくる
	response, err = repo.config.Reqconf(data)
	if err != nil {
		return nil, err
	}
	return response, err
}

func sgConf_Rst(config *domain.Mails, param *domain.RstmailPara) (jsondata []byte, err error) {
	// Mailﾃﾞｰﾀを作成
	mail := &RstMail{
		Subject: config.Subject,
		Personalizations: []RstPersonal{
			{To: []MailUser{{
				Email: config.To.Address,
				Name:  config.To.Express,
			}},
				Parameter: RstmailPara{
					Reseturl: param.Reseturl,
					Email:    param.Email,
				},
			},
		},
		From: MailUser{
			Email: config.From.Address,
			Name:  config.From.Express,
		},
		TemplateId: os.Getenv("SENDGRID_RST_TEMPLATED_ID"),
	}

	// MailデータをJSON化
	jsondata, err = json.Marshal(&(mail))
	if err != nil {
		return jsondata, err
	}
	return jsondata, err
}

func (repo *UserRepository) Sendmail_Cnt(config *domain.Mails, param *domain.CntmailPara) (response *rest.Response, err error) {
	// SendGridのmail config
	data, err := sgConf_Cnt(config, *param)
	// fmt.Printf("mail: %+v\n", mail)
	if err != nil {
		return nil, err
	}

	// sendGridのAPIにﾘｸｴｽﾄをｾｯﾄ, 戻り値でresponseが返ってくる
	response, err = repo.config.Reqconf(data)
	if err != nil {
		return nil, err
	}
	return response, err
}

func sgConf_Cnt(config *domain.Mails, param domain.CntmailPara) (jsondata []byte, err error) {
	// Mailデータを作成
	mail := &CntMail{
		Subject: config.Subject,
		Personalizations: []CntPersonal{
			{To: []MailUser{{
				Email: config.To.Address,
				Name:  config.To.Express,
			}},
				Parameter: param,
			},
		},
		From: MailUser{
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
	jsondata, err = json.Marshal(&mail)
	if err != nil {
		return jsondata, err
	}
	return jsondata, err
}
