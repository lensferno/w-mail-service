package conf

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
)

const (
	DevEnv  = "dev"
	ProdEnv = "prod"
)

var (
	Conf = &Config{}
)

type Config struct {
	Mail    MailConf
	Service ServiceConf
}

type ServiceConf struct {
	Address  string
	Port     int
	BasePath string
	LogFile  string
}

type MailConf struct {
	TemplateDir       string `json:"TemplateDir"`
	Account           string `json:"Account"`
	Password          string `json:"Password"`
	SendServer        string `json:"SendServer"`
	SendServerPort    int    `json:"SendServerPort"`
	ReceiveServer     string `json:"ReceiveServer"`
	ReceiveServerPort int    `json:"ReceiveServerPort"`
}

func Init() (err error) {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/etc/mail-service")
	viper.AddConfigPath("$HOME/.mail-service")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(Conf)
	if err != nil {
		return
	}

	jsonByte, _ := jsoniter.Marshal(Conf)
	fmt.Println(string(jsonByte))

	return
}
