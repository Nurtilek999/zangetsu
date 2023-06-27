package config

import (
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
)

func GetConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configurations")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("config error: ", err.Error())
	}
}

func SetupConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     "84705837482-i4ica4s76hc5ggiuvu4jmqdr60rnsr4m.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8000/v1/user/callback",

		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}
