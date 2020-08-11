package config

import (
	"sync"

	"github.com/BurntSushi/toml"
)

type GlobalConfig struct {
	ServerPort string `toml:"server_port"`
	Runmode    string `toml:"runmode"`
	DBDemo     DBDemo `toml:"db_demo"`
	Kafka      Kafka
	Influxdb   Influxdb
	ACL        ACL
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

type Kafka struct {
	Brokers       []string
	ProducerTopic string `toml:"producer_topic"`
	ConsumerTopic string `toml:"consumer_topic"`
	ConsumerGroup string `toml:"consumer_group"`
}

type Influxdb struct {
	Address     string
	Username    string
	Password    string
	DBName      string
	Measurement string
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
