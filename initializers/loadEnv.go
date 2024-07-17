package initializers

import (
	"time"

	"github.com/spf13/viper"
)

// Allowed environment variables structure
type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

// Load environment variables from app.env file and make them
// accessible in other files and packages within the applicaton code.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env") // method to tell Viper the type of configuration file we used.
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig() // start reading the values in the config file.
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config) // unmarshal the values into the provided struct.
	return
}
