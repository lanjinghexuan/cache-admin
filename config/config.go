package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Redis
	Logx string `yaml:"logx"`
}

type Redis struct {
	Host string
	Port string
}

var RedisDB *redis.Client
var Ctx = context.Background()
var Logx string

func init() {

	dir, _ := os.Getwd()
	fmt.Println("当前运行目录：", dir)

	var config Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("./config")
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

	Logx = config.Logx

	RedisDB = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.Host, config.Port),
	})
	err = RedisDB.Ping(context.Background()).Err()
	if err != nil {
		fmt.Printf("RedisDB.Ping() failed with %s\n", err)
		panic(err)
	}
}
