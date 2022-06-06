package configs

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
)

var config *Config

type Config struct {
	LogLevel logrus.Level `yaml:"logLevel"` // 5=debug 4=info 3=warn
	Web      WebConfig    `yaml:"web"`
	Redis    RedisConfig  `yaml:"redis"`
	Mongo    MongoConfig  `yaml:"mongo"`
}

type WebConfig struct {
	Addr   string `json:"addr" yaml:"addr"`
	View   string `yaml:"view"`
	Static string `yaml:"static"`
	Domain string `yaml:"domain"`
}

type RedisConfig struct {
	Addr     string `json:"addr" yaml:"addr"`
	Password string `json:"password" yaml:"password"`
	Db       int    `json:"db" yaml:"db"`
}

type MongoConfig struct {
	Username      string `json:"username" yaml:"username"`
	Password      string `json:"password" yaml:"password"`
	Host          string `json:"host" yaml:"host"`
	AdminDatabase string `json:"adminDatabase" yaml:"adminDatabase"`
	PoolSize      uint64 `json:"poolSize" yaml:"poolSize"`
	MaxPoolSize   uint64 `json:"maxPoolSize" yaml:"maxPoolSize"`
	Timeout       int    `json:"timeout" yaml:"timeout"`
	Database      string `json:"database" yaml:"database"`
}

var lock sync.RWMutex

func ReloadConfig() error {
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return err
	}
	lock.Lock()
	defer lock.Unlock()

	config = &c
	return nil
}

func Get() *Config {
	lock.RLock()
	defer lock.RUnlock()

	var c Config
	c = *config
	return &c
}
