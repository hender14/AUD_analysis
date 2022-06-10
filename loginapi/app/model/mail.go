package model

type Mails struct {
	From    Sendaddress
	To      Sendaddress
	Subject string
	// Text    Content
}

type Sendaddress struct {
	Express string
	Address string
}

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
	To        []MailUser  `json:"to"`
	Parameter CntmailPara `json:"dynamic_template_data"`
}

type CntmailPara struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Email   string `json:"email"`
}
