package model

type MailRequest struct {
	Code      string `json:"code"`
	Email     string `json:"email"`
	EmailType string `json:"emailType"`
	Date      string `json:"date"`
	Sign      string `json:"sign"`
}

type MailSendData struct {
	Target          string `json:"email"`
	TemplateId      string `json:"emailType"`
	TemplateContent map[string]string
}
