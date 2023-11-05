package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Cfg struct {
		AppEnv AppEnv
		Port   string `env:"PORT"`
		DB     DB
	}

	DB struct {
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT"`
		Name     string `env:"DB_NAME"`
		Username string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
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
			readFrom(
				"./config/dev.env",
			),
		)
	}

	if appEnv == AppStage {
		return load(
			readFrom(
				"./config/stage.env",
			),
		)
	}

	if appEnv == AppProd {
		return load(
			readFrom(
				"./config/prod.env",
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

func readFrom(paths ...string) option {
	return func(c *Cfg) {
		for _, path := range paths {
			err := cleanenv.ReadConfig(path, c)
			if err != nil {
				panic(err)
			}
		}
	}
}
