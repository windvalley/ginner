package cron

import (
	"github.com/robfig/cron/v3"
)

// Init cron initialization
func Init() {
	c := cron.New()
	if _, err := c.AddFunc("*/2 * * * *", updateCacheKey); err != nil {
		panic(err)
	}

	c.Start()
}
