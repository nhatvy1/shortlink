package initialize

import (
	"fmt"
	"os"
	"shortlink/pkg/globals"

	"github.com/spf13/viper"
)

func LoadConfig() {

	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "dev"
	}

	v := viper.New()

	v.AddConfigPath("./configs")
	v.SetConfigName(env)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read configuration: %w", err))
	}

	if err := v.Unmarshal(&globals.Config); err != nil {
		panic(fmt.Errorf("unable to decode configuration: %w", err))
	}

	fmt.Println("config loaded:", env)
}
