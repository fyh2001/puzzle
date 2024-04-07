package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Application Application `mapstructure:"app"`
	Mysql       Mysql       `mapstructure:"mysql"`
	Redis       Redis       `mapstructure:"redis"`
	RabbitMQ    RabbitMQ    `mapstructure:"rabbitmq"`
	Cos         Cos         `mapstructure:"cos"`
}

type Application struct {
	Name string `mapstructure:"name"`
	Port string `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}

type RabbitMQ struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Cos struct {
	SecretID  string `mapstructure:"secret_id"`
	SecretKey string `mapstructure:"secret_key"`
	Url       string `mapstructure:"url"`
}

var Settings Config

func InitConfig() {
	// 设置配置文件名
	viper.SetConfigName("config")
	// 设置配置文件类型
	viper.SetConfigType("yaml")
	// 设置配置文件路径
	viper.AddConfigPath(".")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// 将配置文件内容映射到结构体
	err = viper.Unmarshal(&Settings)
	if err != nil {
		panic(fmt.Errorf("unmarshal config file error: %s", err))
	}
}
