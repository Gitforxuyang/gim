package conf

import (
	"fmt"
	"gim/utils"
	"github.com/spf13/viper"
)

type RedisConfig struct {
	Addr     string
	Password string
	Db       int
}

type RabbitConfig struct {
}

type Config struct {
	Redis    *RedisConfig
	Rabbit   *RabbitConfig
	LogLevel string
}

func InitConfig() *Config {
	config := Config{}
	v := viper.New()
	v.SetConfigName("config.default")
	v.AddConfigPath("./conf")
	v.SetConfigType("json")

	err := v.ReadInConfig()
	utils.Must(err)
	v.BindEnv("ENV")
	env := v.GetString("ENV")
	if env == "" {
		env = "default"
	}
	v.SetConfigName(fmt.Sprintf("config.%s", env))
	err = v.MergeInConfig()
	utils.Must(err)
	err = v.UnmarshalKey("redis", &config.Redis)
	utils.Must(err)
	err = v.UnmarshalKey("rabbit", &config.Rabbit)
	utils.Must(err)
	config.LogLevel = v.GetString("logLevel")
	return &config
}
