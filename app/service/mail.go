package service

import (
	"fmt"
	"github.com/flosch/pongo2/v6"
	"github.com/wneessen/go-mail"
	"go.uber.org/zap"
	"mail-service/app/model"
	"mail-service/library/log"
	"os"
)

var (
	templates = map[string]string{}
)

func (s *Service) SendMail(mailSendData model.MailSendData) (err error) {
	m := mail.NewMsg()
	if err = m.To(mailSendData.Target); err != nil {
		return
	}
	if err = m.From(s.config.Mail.Account); err != nil {
		return
	}

	template := templates[mailSendData.TemplateId]
	if template == "" {
		templateData, err := os.ReadFile(
			fmt.Sprintf("%s/%s.tmpl", s.config.Mail.TemplateDir, mailSendData.TemplateId),
		)
		if err != nil {
			log.Error("读取邮件模板时出错", zap.Error(err))
			return err
		}
		template = string(templateData)
	}

	tmpl, err := pongo2.FromString(template)
	if err != nil {
		log.Error("读取邮件模板时出错", zap.Error(err))
		return err
	}

	context := pongo2.Context{}
	for s2 := range mailSendData.TemplateContent {
		context[s2] = mailSendData.TemplateContent[s2]
	}

	mailText, err := tmpl.Execute(context)
	if err != nil {
		log.Error("生成邮件时出错", zap.Error(err), zap.Any("request", mailSendData))
		return err
	}

	m.Subject("请验证您的验证码")
	m.SetBodyString(mail.TypeTextHTML, mailText)
	if err = s.mail.DialAndSend(m); err != nil {
		log.Error("发送邮件时出错", zap.Error(err))
		return err
	}

	log.Info("邮件发送完成", zap.String("content", template))
	return nil
}
