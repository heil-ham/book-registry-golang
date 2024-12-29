package helper

import "github.com/spf13/viper"

type Config struct {
	Username	string
	Host 		string
	Port 		int
}

func GetConfigs() *Config {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("env")
	config.AddConfigPath("/")

	err := config.ReadInConfig()

	if err != nil {
		panic(err)
	}
	
	return &Config{
		Host: config.GetString("DATABASE_HOST"),
		Username:     config.GetString("DATABASE_USERNAME"),
		Port:     config.GetInt("DATABASE_PORT"),
	}
}