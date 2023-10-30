package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Cfg struct {
		AppEnv AppEnv
		Port   string `env:"PORT"`
	}

	AppEnv string

	option func(*Cfg)
)

const (
	AppProd  AppEnv = "prod"
	AppStage AppEnv = "stage"
	AppDev   AppEnv = "dev"
)

func (appEnv *AppEnv) String() string {
	return string(*appEnv)
}

func New() *Cfg {
	appEnv := AppEnv(os.Getenv("APP_ENV"))
	if appEnv == AppDev {
		return load(
			readFromFiles(
				"./config/.env",
				"./config/dev.env",
			),
		)
	}
	if appEnv == AppProd {
		return load(
			readFromFiles(
				"./config/.env",
				"./config/prod.env",
			),
		)
	}
	if appEnv == AppStage {
		return load(
			readFromFiles(
				"./config/.env",
				"./config/stage.env",
			),
		)
	}

	return nil
}

func load(opts ...option) *Cfg {
	var cfg Cfg
	for _, option := range opts {
		option(&cfg)
	}

	return &cfg
}

func readFromFiles(paths ...string) option {
	return func(c *Cfg) {
		for _, path := range paths {
			err := cleanenv.ReadConfig(path, c)
			if err != nil {
				panic(err)
			}
		}
	}
}
