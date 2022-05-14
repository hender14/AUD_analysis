package model

type Mails struct {
	From    Sendaddress
	To      Sendaddress
	Subject string
	Text    Content
}

type Sendaddress struct {
	Express    string
	Address     string
}

type Content struct {
	Plantext    string
	Htmltext    string
}

// requert.Bodyに格納するJSONの元となるメールを表す構造体
type Mail struct {
	Subject          string             `json:"subject"`
	Personalizations []Personalizations `json:"personalizations"`
	From             MailUser           `json:"from"`
	TemplateId       string             `json:"template_id"`
	// Content          []Contents         `json:"content"`
}

// 封筒のようなもの
// メールのメタデータを表す構造体
type Personalizations struct {
	To []MailUser `json:"to"`
	Parameter TemplatePara `json:"dynamic_template_data"`
}

// メールのユーザーを表す構造体
type MailUser struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// メールの変数情報
type TemplatePara struct {
	Reseturl  string `json:"reseturl"`
	Email string `json:"email"`
}

// メールの中身を表す構造体
// type Contents struct {
// 	Type  string `json:"type"`
// 	Value string `json:"value"`
// }