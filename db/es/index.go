package es

import (
	"fmt"

	"github.com/pquerna/ffjson/ffjson"
)

// CreateIndex creates index
func CreateIndex(indexName string) error {
	res, err := Client.Indices.Create(indexName)
	if err != nil {
		return fmt.Errorf("create index %s failed, error: %v", indexName, err)
	}
	if res.IsError() {
		return fmt.Errorf("create index %s failed, error: %v", indexName, res)
	}
	res.Body.Close()

	return nil
}

// DeleteIndices delete indices
func DeleteIndices(indexNames []string) error {
	res, err := Client.Indices.Delete(
		indexNames,
		Client.Indices.Delete.WithIgnoreUnavailable(true),
	)
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("delete indices error: %s", res)
	}
	res.Body.Close()

	return nil
}

// DeleteIndexAliases delete aliases
func DeleteIndexAliases(indexNames, aliasNames []string) error {
	res, err := Client.Indices.DeleteAlias(indexNames, aliasNames)
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("delete index aliases error: %s", res)
	}
	res.Body.Close()

	return nil
}

// PutIndexAlias creates or updates an alias
func PutIndexAlias(indexNames []string, aliasName string) error {
	res, err := Client.Indices.PutAlias(indexNames, aliasName)
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("put index alias error: %s", res)
	}
	res.Body.Close()

	return nil
}

// GetIndexAliases return alias indices
func GetIndexAliases(indexAliasName string) ([]string, error) {
	var resp map[string]interface{}
	res, err := Client.Indices.GetAlias(
		Client.Indices.GetAlias.WithIndex(indexAliasName))
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, fmt.Errorf("get index aliases error: %s", res)
	}
	defer res.Body.Close()

	decoder := ffjson.NewDecoder()
	if err := decoder.DecodeReader(res.Body, &resp); err != nil {
		return nil, err
	}

	ret := []string{}
	for key := range resp {
		ret = append(ret, key)
	}

	return ret, nil
}
