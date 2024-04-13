package initializer

import "github.com/spf13/viper"

type Config struct {
	ServerPort string `mapstructure:"PORT"`
	UriAddress string `mapstructure:"URI_ADDRESS"`
	DBName string `mapstructure:"DB_NAME"`
	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
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