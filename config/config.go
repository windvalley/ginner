package config

import "time"

// GlobalConfig map conf/config.toml
type GlobalConfig struct {
	AppName      string `toml:"app_name"`
	ServerPort   string `toml:"server_port"`
	Runmode      string
	HTTPS        https
	Log          log
	Mail         mail
	Auth         auth
	RDBs         map[string]rdb
	Kafka        kafka
	Influxdb     influxdb
	Redis        redis
	RedisCluster redisCluster `toml:"redis_cluster"`
	ACL          acl
}

type https struct {
	Enable bool
	Cert   string
	Key    string
}

type log struct {
	Dirname       string
	RotationHours int    `toml:"rotation_hours"`
	SaveDays      int    `toml:"save_days"`
	LogFormat     string `toml:"log_format"`
}

type mail struct {
	SMTPHost string `toml:"smtp_host"`
	Port     int
	User     string
	Password string
	MailTo   []string `toml:"mail_to"`
}

type auth struct {
	JWTSecret       string `toml:"jwt_secret"`
	JWTLifetime     int64  `toml:"jwt_lifetime"`
	JWTMaxLifetime  int64  `toml:"jwt_max_lifetime"`
	APISignLifetime int64  `toml:"apisign_lifetime"`
}

type rdb struct {
	DBType       string
	Address      string
	DBName       string
	User         string
	Password     string
	MaxIdleConns int `toml:"max_idle_conns"`
	MaxOpenConns int `toml:"max_open_conns"`
}

type kafka struct {
	Brokers       []string
	ProducerTopic string `toml:"producer_topic"`
	ConsumerTopic string `toml:"consumer_topic"`
	ConsumerGroup string `toml:"consumer_group"`
}

type influxdb struct {
	Address     string
	Username    string
	Password    string
	DBName      string
	Measurement string
}

type redis struct {
	Address     string
	DB          int
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type redisCluster struct {
	Nodes       []string
	Secret      string
	PoolSize    int
	PoolTimeout time.Duration
}

type acl struct {
	AllowURL []string `toml:"allow_url"`
	AllowIP  []string `toml:"allow_ip"`
}
