package main

import (
	"fmt"
	"github.com/yitter/idgenerator-go/idgen"
	"gopkg.in/natefinch/lumberjack.v2"
	v1 "mail-service/app/api/http"
	"mail-service/app/conf"
	"mail-service/library/ecode"
	"mail-service/library/log"
	"os"
)

func main() {
	// config加载必须在最先位置
	loadConfig()

	setupIdGenerator()
	setupLogger()

	ecode.InitEcodeText()

	// server启动必须在最后
	startServer()
}

func loadConfig() {
	if confErr := conf.Init(); confErr != nil {
		panic(confErr)
	}
}

func setupIdGenerator() {
	// 暂时写死workerId为8，要改以后机器扩容了再说
	options := idgen.NewIdGeneratorOptions(8)
	idgen.SetIdGenerator(options)
}

func setupLogger() {
	stdoutLogLevel := log.WarnLevel

	stdoutLogLevel = log.InfoLevel

	tees := []log.TeeOption{
		{
			Out: os.Stdout,
			LevelEnablerFunc: func(level log.Level) bool {
				return level >= stdoutLogLevel
			},
		},
	}

	logFileLocation := conf.Conf.Service.LogFile
	if logFileLocation != "" {
		tees = append(tees, log.TeeOption{
			Out: &lumberjack.Logger{
				Filename: logFileLocation,
				MaxSize:  128,
				MaxAge:   30,
				Compress: true,
			},
			LevelEnablerFunc: func(level log.Level) bool {
				return level >= log.InfoLevel
			},
		})
	}

	logger := log.NewTee(tees)
	log.SetDefault(logger)
}

func startServer() {
	//setupMiddleware()

	engine := v1.NewEngine(conf.Conf, conf.Conf.Service.BasePath)
	if engine == nil {
		log.Error("邮件服务初始化失败")
		return
	}
	addr := fmt.Sprintf("%s:%d", conf.Conf.Service.Address, conf.Conf.Service.Port)
	err := engine.Run(addr)
	if err != nil {
		panic(err)
	}
}
