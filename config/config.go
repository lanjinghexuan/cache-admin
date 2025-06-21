package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type Config struct {
	Redis
}

type Redis struct {
	Host string
	Port string
}

var RedisDB *redis.Client
var Ctx = context.Background()

func init() {
	var config Config
	viper.SetConfigName("./config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed with %s\n", err)
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("viper.Unmarshal() failed with %s\n", err)
		panic(err)
	}

	RedisDB = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.Host, config.Port),
	})
	err = RedisDB.Ping(context.Background()).Err()
	if err != nil {
		fmt.Printf("RedisDB.Ping() failed with %s\n", err)
		panic(err)
	}
}
