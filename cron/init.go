package cron

import (
	"github.com/robfig/cron/v3"
)

// Init cron initialization
func Init() {
	c := cron.New()
	c.AddFunc("*/2 * * * *", updateCacheKey)

	c.Start()
}
