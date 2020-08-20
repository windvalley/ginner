package cron

import (
	"github.com/muesli/cache2go"

	"use-gin/logger"
)

// Cache usage:
//    e.g. res, err := cron.Cache.Value(cron.DemoKey)
//if err != nil {
//cron.UpdateCacheKey()
//res, err = cron.Cache.Value(cron.DemoKey)
//if err != nil {
//return nil, err
//}
//}
//value := res.Data().(yourType)
var Cache = cache2go.Cache("gocache")

const DemoKey = "demo"

func UpdateCacheKey() {
	if _, err := Cache.Delete(DemoKey); err != nil {
		logger.Log.Warnf(
			"cron: delete key %s, warning: Gocache has no %s key",
			DemoKey,
			DemoKey,
		)
	}

	Cache.Add(DemoKey, 0, "your specific object")

	logger.Log.Debugf("cron: %s is updated", DemoKey)
}
