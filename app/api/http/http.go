package http

import (
	"github.com/gin-gonic/gin"
	"mail-service/app/conf"
	"mail-service/app/service"
	"mail-service/library/token"
)

var (
	serv *service.Service
	jwt  *token.Token
)

func NewEngine(c *conf.Config, basePath string) *gin.Engine {
	engine := gin.Default()
	rootRouter := engine.RouterGroup.Group(basePath)
	//rootRouter.Use(gin.LoggerWithWriter(*log.DefaultWriter().))

	setupOuterRouter(rootRouter)

	var err error
	serv, err = service.New(c)
	if err != nil {
		return nil
	}

	return engine
}

func setupOuterRouter(group *gin.RouterGroup) {
	mail := group.Group("/mail")
	{
		mail.POST("/send", sendMail)
	}
}

func response(c *gin.Context, code int, msg string, data any) {
	resp := gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}

	c.JSON(200, resp)
}
