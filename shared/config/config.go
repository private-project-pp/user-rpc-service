package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type ConfigApp struct {
	Service ServiceConfig `mapstructure:"service" yaml:"service"`
	Db      DbConfig      `mapstructure:"db" yaml:"db"`
}

type ServiceConfig struct {
	Name string `mapstructure:"name" yaml:"name"`
	Port string `mapstructure:"port" yaml:"port"`
}

type DbConfig struct {
	Host string `mapstructure:"host" yaml:"host"`
}

var (
	Service ServiceConfig
	DB      DbConfig
)

func SetupConfig() (out ConfigApp) {
	var err error
	viper.AddRemoteProvider("vault", "https://127.0.0.1:8443", "/secret/whatever/config.json")
	viper.SetConfigType("json")
	err = viper.ReadRemoteConfig()
	if err != nil {
		err = readLocalConfig()
	}
	if err != nil {
		panic(err)
	}

	viper.Unmarshal(&out)

	Service = out.Service
	DB = out.Db
	return out
}

func readLocalConfig() (err error) {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(filepath.Join(path, "shared/config"))
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
