package initialize

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"pawtopia.com/global"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	env := os.Getenv("ENV")
	if env == "dev" {
		viper.SetConfigName("dev")
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// fmt.Println(viper.Get("server.port"))
	// fmt.Println(viper.Get("security.jwt.key"))

	err = viper.Unmarshal(&global.Config)
	if err != nil {
		fmt.Println("Error unmarshalling config:", err)
		return
	}
}
