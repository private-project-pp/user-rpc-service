package config

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
	Config ConfigApp
)

func SetupConfig() (out ConfigApp, err error) {

	Config = out
	return out, nil
}
