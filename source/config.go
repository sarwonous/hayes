package source

import (
	"fmt"

	"github.com/spf13/viper"
)

func configInit() error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(fmt.Sprintf("%s", "./config.yaml"))
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
