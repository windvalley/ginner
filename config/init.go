package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	flag "github.com/spf13/pflag"
)

// Init config from command line params
func Init() {
	cfg := getConfigFileFromCli()
	ParseConfig(*cfg, config)
}

// InitCmd config for subproject cmd
func InitCmd(conf interface{}) {
	cfg := getConfigFileFromCli()
	ParseConfig(*cfg, conf)
}

// ParseConfig parse config from path string
func ParseConfig(f string, conf interface{}) {
	metaData, err := toml.DecodeFile(f, conf)
	if err != nil {
		fmt.Printf("Parse config file failed: %s\n", err)
		usage()
	}

	if len(metaData.Keys()) == 0 {
		fmt.Println("Parse config file failed: has no valid config items")
		usage()
	}

	undecodeItems := metaData.Undecoded()
	if len(undecodeItems) != 0 {
		fmt.Printf(
			"Parse config file failed: follow iterms failed to resolve:\n  %s\n\n",
			undecodeItems,
		)
		usage()
	}
}

func getConfigFileFromCli() *string {
	cfg := flag.StringP("config", "c", "", "Specify your configuration file")
	flag.Parse()

	if *cfg == "" {
		usage()
	}

	return cfg
}

func usage() {
	flag.Usage()
	os.Exit(2)
}
