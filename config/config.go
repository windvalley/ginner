package config

import "time"

type GlobalConfig struct {
	AppName    string `toml:"app_name"`
	ServerPort string `toml:"server_port"`
	Runmode    string
	Log        Log
	Mail       Mail
	Auth       Auth
	MySQL      MySQL
	PostgreSQL PostgreSQL
	Kafka      Kafka
	Influxdb   Influxdb
	Redis      Redis
	ACL        ACL
}

type Log struct {
	Dirname       string
	RotationHours int    `toml:"rotation_hours"`
	SaveDays      int    `toml:"save_days"`
	LogFormat     string `toml:"log_format"`
}

type Mail struct {
	SMTPHost string `toml:"smtp_host"`
	Port     int
	User     string
	Password string
	MailTo   []string `toml:"mail_to"`
}

type Auth struct {
	JWTSecret       string `toml:"jwt_secret"`
	JWTLifetime     int64  `toml:"jwt_lifetime"`
	JWTMaxLifetime  int64  `toml:"jwt_max_lifetime"`
	APISignLifetime int64  `toml:"apisign_lifetime"`
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

type Redis struct {
	Address     string
	DB          int
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type ACL struct {
	AllowURL []string `toml:"allow_url"`
	AllowIP  []string `toml:"allow_ip"`
}
