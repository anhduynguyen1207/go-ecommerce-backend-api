package initialize

import (
	"fmt"

	"github.com/anhduynguyen1207/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
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

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Println("Unable to decode configuration %v", err)
	}
}
