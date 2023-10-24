package service

import (
	"github.com/lmmfy/goejs/pkg/contract"
	"github.com/lmmfy/goejs/pkg/interpreter/otto"
	"github.com/wneessen/go-mail"
	"go.uber.org/zap"
	"mail-service/app/conf"
	"mail-service/library/log"
	"time"
)

const (
	wusthelperTokenExpiration = time.Hour * 24
)

type Service struct {
	config         *conf.Config
	mail           *mail.Client
	templateEngine contract.Engine
}

func New(c *conf.Config) (service *Service, err error) {
	service = &Service{
		config:         c,
		templateEngine: otto.NewDefaultOttoEngine(),
	}

	client, err := mail.NewClient(
		c.Mail.SendServer,
		mail.WithPort(c.Mail.SendServerPort),
		mail.WithSMTPAuth(mail.SMTPAuthLogin),
		mail.WithUsername(c.Mail.Account),
		mail.WithPassword(c.Mail.Password),
		mail.WithSSL(),
	)

	if err != nil {
		log.Error("邮件服务初始化失败", zap.String("error", err.Error()))
		return nil, err
	}

	service.mail = client

	return
}
