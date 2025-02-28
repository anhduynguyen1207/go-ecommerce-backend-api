package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/") //path to config file
	viper.SetConfigName("local")     // name of config file
	viper.SetConfigType("yaml")

	// read config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failded to red configuration %w \n", err))
	}

	// read value from configuration
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Server Security::", viper.GetString("security.jwt.key"))

	// configure structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}

	fmt.Println("Config Port::", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("DB User: %s, password: %s, host: %s \n", db.User, db.Password, db.Host)

	}

}
