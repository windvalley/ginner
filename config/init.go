package config

import (
	"sync"

	"github.com/BurntSushi/toml"
)

type GlobalConfig struct {
	ServerPort string `toml:"server_port"`
	Runmode    string `toml:"runmode"`
	DBDemo     `toml:"db_demo"`
	ACL        `toml:"acl"`
}

var (
	config *GlobalConfig
	lock   = new(sync.RWMutex)
)

type DBDemo struct {
	Address  string
	DBName   string
	User     string
	Password string
}

type ACL struct {
	AllowURL []string `toml:"allow_url"`
	AllowIP  []string `toml:"allow_ip"`
}

func ParseConfig(f string) {
	if _, err := toml.DecodeFile(f, &config); err != nil {
		panic(err)
	}
	return
}

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}
