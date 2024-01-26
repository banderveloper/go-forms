package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"strings"
	"time"
)

type Config struct {
	Environment string `yaml:"environment" env:"ENV" env-required:"true"`
	Server      `yaml:"server"`
	Database    `yaml:"database"`
}

type Server struct {
	Address     string        `yaml:"address" env-default:"localhost:5001"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Database struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-default:"5432"`
	DbName   string `yaml:"db_name" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password"`
}

// MustLoad 'Must' naming is used when function does not return error but use panic
// rare practice, but config initialization moment is ok
func MustLoad() *Config {

	launchArgs := getLaunchArgs()

	cfgPath := launchArgs["config"]
	dbPass := launchArgs["db-pass"]

	if cfgPath == "" {
		log.Fatal("required --config launch argument is not set")
	}
	if dbPass == "" {
		log.Fatal("required --db-pass launch argument is not set")
	}

	// if config file doesn't exists
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatalf("configuration file doesn't exists: %s", cfgPath)
	}

	var cfg Config

	// try read config and parse it into cfg var
	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		log.Fatalf("error during reading configuration file (%s): %s", cfgPath, err)
	}

	// insert db password from launch arguments
	cfg.Database.Password = dbPass

	return &cfg
}

// getLaunchArgs returns map of all launch args written in format --key value
// args example: --config ./config/local.yaml --db-pass qwerty
// map result: res["config"] = "./config/local.yaml" ...
func getLaunchArgs() map[string]string {
	res := make(map[string]string)

	for i, arg := range os.Args {
		if arg[:2] == "--" {
			cleanArg := strings.TrimLeft(arg, "--")
			res[cleanArg] = os.Args[i+1]
		}
	}

	return res
}
