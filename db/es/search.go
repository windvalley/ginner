package es

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7/esapi"

	"ginner/logger"
)

// Search query elasticsearch
func Search(indexName string, query map[string]interface{}, from, limit int) (
	[]map[string]interface{}, int, error) {
	var (
		buf bytes.Buffer
		res *esapi.Response
		err error
		r   map[string]interface{}
	)

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, 0, fmt.Errorf("encoding query failed: %s", err)
	}

	res, err = Client.Search(
		Client.Search.WithContext(context.Background()),
		Client.Search.WithIndex(indexName),
		Client.Search.WithBody(&buf),
		Client.Search.WithTrackTotalHits(true),
		Client.Search.WithFrom(from),
		Client.Search.WithSize(limit),
	)
	if err != nil {
		return nil, 0, fmt.Errorf("getting response from es failed: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, 0, fmt.Errorf("parsing the response body failed: %s", err)
		}
		return nil, 0, fmt.Errorf(
			"[%s] %s: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"],
		)
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, 0, fmt.Errorf("parsing the response body failed: %s", err)
	}

	count := int(r["hits"].(map[string]interface{})["total"].(float64))

	resp := make([]map[string]interface{}, 0)
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		resp = append(resp,
			hit.(map[string]interface{})["_source"].(map[string]interface{}))
	}

	logger.Log.Debugf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		count,
		int(r["took"].(float64)),
	)

	return resp, count, nil
}
