package initializer

import "github.com/spf13/viper"

type Config struct {
	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
	ServerPort   string `mapstructure:"SERVER_PORT"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBDBname     string `mapstructure:"DB_DBNAME"`
	DBSSLMode    string `mapstructure:"DB_SSLMODE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	return
}
