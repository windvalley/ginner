package cron

import (
	"github.com/robfig/cron/v3" // 版本一定要用v3, 低于这个版本用法不兼容.
)

func Init() {
	c := cron.New()
	c.AddFunc("*/2 * * * *", UpdateCacheKey)

	c.Start()
}
