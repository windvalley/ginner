package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/dustin/go-humanize"
	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"

	"ginner/db/es"
	"ginner/logger"
)

type dataSaveEngine struct {
	ESCli          *elasticsearch.Client
	NumWorkers     int
	FlushBytes     int
	IndexName      string
	IndexAliasName string
	OldIndices     []string
	Err            error
}

func newDataSaveEngine(indexAliasName string) *dataSaveEngine {
	return &dataSaveEngine{
		ESCli:          es.Client,
		IndexAliasName: indexAliasName,
	}
}

func (e *dataSaveEngine) GetOldIndices() *dataSaveEngine {
	e.OldIndices, e.Err = es.GetIndexAliases(e.IndexAliasName)
	return e
}

func (e *dataSaveEngine) SaveData(datas []*domainFinalData) *dataSaveEngine {
	if e.Err != nil {
		return e
	}

	if e.IndexName == "" {
		timeStr := time.Now().Format("20060102150405")
		e.IndexName = e.IndexAliasName + "-" + timeStr
	}
	if e.FlushBytes == 0 {
		e.FlushBytes = 100000
	}
	if e.NumWorkers == 0 {
		e.NumWorkers = runtime.NumCPU()
	}

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         e.IndexName,      // The default index name
		Client:        e.ESCli,          // The Elasticsearch client
		NumWorkers:    e.NumWorkers,     // The number of worker goroutines
		FlushBytes:    e.FlushBytes,     // The flush threshold in bytes
		FlushInterval: 10 * time.Second, // The periodic flush interval
	})
	if err != nil {
		e.Err = fmt.Errorf("creating the indexer failed: %s", err)
		return e
	}

	if err := es.DeleteIndices([]string{e.IndexName}); err != nil {
		e.Err = err
		return e
	}

	if err := es.CreateIndex(e.IndexName); err != nil {
		e.Err = err
		return e
	}

	start := time.Now().UTC()

	var countSuccessful uint64
	for _, a := range datas {
		documentID := a.Host + "." + a.Zone
		data, err := json.Marshal(a)
		if err != nil {
			e.Err = fmt.Errorf(
				"encode domain records %s failed, error: %v", documentID, err)
			return e
		}

		err = bi.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				// Action field configures the operation to perform (index, create, delete, update)
				Action: "index",
				// DocumentID is the (optional) document ID
				DocumentID: documentID,
				// Body is an `io.Reader` with the payload
				Body: bytes.NewReader(data),
				// OnSuccess is called for each successful operation
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem,
					res esutil.BulkIndexerResponseItem) {
					atomic.AddUint64(&countSuccessful, 1)
				},
				// OnFailure is called for each failed operation
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem,
					res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						logger.Log.Errorf(
							"add item(%s) to bulk indexer error: %s", documentID, err)
					} else {
						logger.Log.Errorf(
							"add item(%s) to bulk indexer error, error type: %s, error reason: %s",
							documentID, res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
		if err != nil {
			e.Err = fmt.Errorf(
				"add item(%s) to bulk indexer meet up unexpected error: %s",
				documentID, err)
			return e
		}
	}

	if err := bi.Close(context.Background()); err != nil {
		e.Err = fmt.Errorf("sync data to es meet up unexpected error: %s", err)
		return e
	}

	biStats := bi.Stats()
	dur := time.Since(start)

	if biStats.NumFailed > 0 {
		e.Err = fmt.Errorf(
			"indexed [%s] documents with [%s] errors in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			humanize.Comma(int64(biStats.NumFailed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
		return e
	}

	logger.Log.Debugf(
		"successfully indexed [%s] documents in %s (%s docs/sec), other stats: %+v",
		humanize.Comma(int64(biStats.NumFlushed)),
		dur.Truncate(time.Millisecond),
		humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		biStats,
	)

	return e
}

func (e *dataSaveEngine) Clean() *dataSaveEngine {
	if e.Err != nil {
		return e
	}
	if err := es.PutIndexAlias([]string{e.IndexName}, e.IndexAliasName); err != nil {
		e.Err = err
		return e
	}

	if err := es.DeleteIndexAliases(e.OldIndices, []string{e.IndexAliasName}); err != nil {
		e.Err = err
		return e
	}

	if err := es.DeleteIndices(e.OldIndices); err != nil {
		e.Err = err
		return e
	}

	return e
}
