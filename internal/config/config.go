package configs

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"sync"
)

type Config struct {
	IsDebug *bool   `yaml:"is_debug" env-required:"true"`
	Listen  Listen  `yaml:"listen"`
	Swagger Swagger `yaml:"swagger"`
}

type Swagger struct {
	SVersion    string `yaml:"version" env-default:"1.0"`
	ServiceName string `yaml:"service_name" env-default:"Salam Service"`
	Title       string `yaml:"title" env-default:"Salam Hj"`
	Host        string `yaml:"host" env-required:"true"`
}

type Listen struct {
	Type   string `yaml:"type" env-required:"true"`
	BindIP string `yaml:"bind_ip" env-required:"true"`
	Port   string `yaml:"port" env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		pathConfig := pwd + "/config.yml"

		fmt.Println("read application configuration pwd: ", pwd)

		instance = &Config{}

		if err = cleanenv.ReadConfig(pathConfig, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			fmt.Println(help)
			fmt.Println(err)
		}
	})
	return instance
}
