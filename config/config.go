package config

import "github.com/tkanos/gonfig"

type Config struct {
	DB_NAME       string
	DB_USERNAME   string
	DB_HOST       string
	DB_PASSWORD   string
	EMAIL_MAIL    string
	PASSWORD_MAIL string
	FROM_MAIL     string
	HOST_MAIL     string
	ANDRESS_MAIL  string

}

func GetConfig() Config{
 gonfig.GetConf("config/config.json", &Config{})
 return Config{}
}
