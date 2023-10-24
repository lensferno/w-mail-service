package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mail-service/app/model"
	"mail-service/library/ecode"
	"mail-service/library/log"
)

func sendMail(c *gin.Context) {
	req := new(model.MailRequest)
	err := c.BindJSON(req)
	if err != nil {
		log.Error("邮件发送请求参数解析失败", zap.String("err", err.Error()))
		response(c, ecode.ParamWrong.Code(), "param wrong", nil)
		return
	}

	data := model.MailSendData{
		Target:     req.Email,
		TemplateId: req.EmailType,
		TemplateContent: map[string]string{
			"code": req.Code,
			"date": req.Date,
			"sign": req.Sign,
		},
	}

	err = serv.SendMail(data)
	if err != nil {
		log.Error("邮件发送失败", zap.String("err", err.Error()))
		response(c, ecode.ServerErr.Code(), "mail send fail: "+err.Error(), nil)
		return
	}

	response(c, ecode.OK.Code(), "", nil)
}
