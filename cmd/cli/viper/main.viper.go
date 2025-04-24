package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Security struct {
		JWT struct {
			Key string `mapstructure:"key"`
		} `mapstructure:"jwt"`
	} `mapstructure:"security"`
	Databases []struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// fmt.Println(viper.Get("server.port"))
	// fmt.Println(viper.Get("security.jwt.key"))

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("Error unmarshalling config:", err)
		return
	}
	fmt.Println("Server Port:", config.Server.Port)
	fmt.Println("Security JWT Key:", config.Security.JWT.Key)
	for _, db := range config.Databases {
		fmt.Println("Database Host:", db.Host)
		fmt.Println("Database Port:", db.Port)
		fmt.Println("Database User:", db.User)
		fmt.Println("Database Password:", db.Password)
		fmt.Println("Database Name:", db.Name)
	}
}
