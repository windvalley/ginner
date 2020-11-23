package es

import (
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
	elasticsearch "github.com/elastic/go-elasticsearch/v7"

	"ginner/config"
	"ginner/logger"
)

// Client global elasticsearch client instance
var Client *elasticsearch.Client

// Init elasticsearch connect initialization
func Init() {
	var err error
	retryBackoff := backoff.NewExponentialBackOff()

	cfg := elasticsearch.Config{
		Addresses: config.Conf().ES.Nodes,
		Username:  config.Conf().ES.Username,
		Password:  config.Conf().ES.Password,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 10,
		},
		// Retry on 429 TooManyRequests statuses
		RetryOnStatus: []int{502, 503, 504, 429},
		// Configure the backoff function
		RetryBackoff: func(i int) time.Duration {
			if i == 1 {
				retryBackoff.Reset()
			}
			return retryBackoff.NextBackOff()
		},
		// Retry up to 5 attempts
		MaxRetries: 5,
	}

	Client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		logger.Log.Fatalf("init elasticsearch failed: %v", err)
	}
}
