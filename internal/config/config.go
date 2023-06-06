package config

import (
	"github.com/spf13/viper"
	"log"
)

type Configuration struct {
	App    *AppConfig
	DB     *DBConfig
	Stripe *StripeConfig
}

type AppConfig struct {
	AppName string
	Env     string
	Addr    string
}

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type StripeConfig struct {
	PublishableKey string
	SecretKey      string
}

func GetConfig(path string) *Configuration {
	config := viper.New()

	config.SetConfigName("config")
	config.AddConfigPath(".")

	if path != "" {
		config.AddConfigPath(path)
	}

	config.AutomaticEnv()

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("Could not load config file: #{err} \n")
	}

	return &Configuration{
		App: &AppConfig{
			AppName: config.GetString("app.app-name"),
			Env:     config.GetString("app.env"),
			Addr:    config.GetString("server.addr"),
		},
		DB: &DBConfig{
			Host: config.GetString("db.host"),
			Port: config.GetString("db.port"),
			User: config.GetString("db.user"),
			Pass: config.GetString("db.pass"),
			Name: config.GetString("db.name"),
		},
		Stripe: &StripeConfig{
			PublishableKey: config.GetString("stripe.pk"),
			SecretKey:      config.GetString("stripe.sk"),
		},
	}
}
