package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func init() {
	viper.SetConfigName("wagetrak")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./tmp")
	viper.AddConfigPath("/etc")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", ""))
	viper.SetEnvPrefix("wt")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("could not find config file, using default configuration")
		} else {
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}
	}
}
