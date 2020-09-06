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

const demoKey = "demo"

func updateCacheKey() {
	if _, err := Cache.Delete(demoKey); err != nil {
		logger.Log.Warnf(
			"cron: delete key %s, warning: Gocache has no %s key",
			demoKey,
			demoKey,
		)
	}

	Cache.Add(demoKey, 0, "your specific object")

	//logger.Log.Debugf("cron: %s is updated", demoKey)
}
