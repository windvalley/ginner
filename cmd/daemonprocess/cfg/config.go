package cfg

import "sync"

var (
	// Config instance
	Config *Global
	lock   = new(sync.RWMutex)
)

// Conf use config
func Conf() *Global {
	lock.RLock()
	defer lock.RUnlock()

	return Config
}

// Global config
type Global struct {
	Runmode string
	Log     log
}

type log struct {
	Dirname       string
	LogFormat     string `toml:"log_format"`
	RotationHours int    `toml:"rotation_hours"`
	SaveDays      int    `toml:"save_days"`
}
