package config

type GlobalConfig struct {
	ServerPort string `toml:"server_port"`
	Runmode    string
	Log        Log
	MySQL      MySQL
	PostgreSQL PostgreSQL
	Kafka      Kafka
	Influxdb   Influxdb
	ACL        ACL
}

type Log struct {
	Dirname       string
	RotationHours int `toml:"rotation_hours"`
	SaveDays      int `toml:"save_days"`
}

type MySQL struct {
	DBType   string
	Address  string
	DBName   string
	User     string
	Password string
}

type PostgreSQL struct {
	DBType   string
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
