package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/spf13/pflag"
)

var (
	conf *GlobalConfig
	lock = new(sync.RWMutex)
)

func ParseConfig(f string) {
	if _, err := toml.DecodeFile(f, &conf); err != nil {
		panic(err)
	}
}

func Conf() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return conf
}

// load config from command line parameters
func LoadFromCLIParams() {
	cfg := pflag.StringP("config", "c", "", "Specify your configuration file")
	pflag.Parse()
	if *cfg == "" {
		binName := filepath.Base(os.Args[0])
		fmt.Printf("missing parameter\nUsage of %s:\n  -c, --config string"+
			"   Specify your configuration file\n", binName)
		os.Exit(2)
	}

	ParseConfig(*cfg)
}

// load from system environment variable RUNENV: prod/dev
func LoadFromENV() {
	abPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	confPath := abPath + "/conf/"
	runenv := os.Getenv("RUNENV")

	switch runenv {
	case "dev":
		ParseConfig(confPath + "dev.config.toml")
	case "prod":
		ParseConfig(confPath + "config.toml")
	case "":
		panic("system environment variable RUNENV is not set, optinal value: prod or dev")
	default:
		panic("the value of RUNENV can only be prod or dev")
	}
}
