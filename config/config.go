package config

import (
	"sync"
	"time"
)

var (
	config *Global
	lock   = new(sync.RWMutex)
)

func init() {
	config = &Global{}
}

// Conf get config instance
func Conf() *Global {
	lock.RLock()
	defer lock.RUnlock()

	return config
}

// Global config, used to map conf/config.toml
type Global struct {
	AppName      string `toml:"app_name"`
	ServerPort   string `toml:"server_port"`
	Runmode      string
	HTTPS        https
	Log          log
	Mail         mail
	Auth         auth
	RDBs         map[string]rdb
	Mongo        mongo
	Kafka        kafka
	ES           es
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
	LogFormat     string `toml:"log_format"`
	LogLevel      string `toml:"log_level"`
	RotationHours int    `toml:"rotation_hours"`
	SaveDays      int    `toml:"save_days"`
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

type mongo struct {
	Address  string
	DBName   string
	Username string
	Password string
}

type kafka struct {
	Brokers       []string
	ProducerTopic string `toml:"producer_topic"`
	ConsumerTopic string `toml:"consumer_topic"`
	ConsumerGroup string `toml:"consumer_group"`
	Keepalive     time.Duration
}

type es struct {
	Nodes    []string
	Username string
	Password string
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
	MaxIdle     int           `toml:"max_idle"`
	MaxActive   int           `toml:"max_active"`
	IdleTimeout time.Duration `toml:"idle_timeout"`
}

type redisCluster struct {
	Nodes       []string
	Secret      string
	PoolSize    int           `toml:"pool_size"`
	PoolTimeout time.Duration `toml:"pool_timeout"`
}

type acl struct {
	AllowURL []string `toml:"allow_url"`
	AllowIP  []string `toml:"allow_ip"`
}
