package v1

import (
	"ginner/db/es"
	"ginner/util"
)

// FilterParams filer parameters
type FilterParams struct {
	IDCName    string `form:"idc_name"`
	DomainType string `form:"domain_type"`
	IPStatus   string `form:"ip_status"`
	AutoSwitch string `form:"auto_switch"`
	Keyword    string `form:"keyword"`
	SearchMode string `form:"search_mod"`
}

// FilterRecords filter and search service
func FilterRecords(departIDs []string, params FilterParams,
	offset, limit int) ([]map[string]interface{}, int, error) {

	filter := []map[string]interface{}{}
	should := []map[string]interface{}{}

	departIDItems := map[string]interface{}{
		"terms": map[string]interface{}{
			"department_id": departIDs,
		},
	}
	filter = append(filter, departIDItems)

	if params.IDCName != "" {
		item := map[string]interface{}{
			"term": map[string]interface{}{
				"value.idc_name.keyword": params.IDCName,
			},
		}
		filter = append(filter, item)
	}
	if params.DomainType != "" {
		item := map[string]interface{}{
			"term": map[string]interface{}{
				"value.domain_type.keyword": params.DomainType,
			},
		}
		filter = append(filter, item)
	}
	if params.IPStatus != "" {
		item := map[string]interface{}{
			"term": map[string]interface{}{
				"value.ip_status.keyword": params.IPStatus,
			},
		}
		filter = append(filter, item)
	}
	if params.AutoSwitch != "" {
		item := map[string]interface{}{
			"term": map[string]interface{}{
				"auto_switch.keyword": params.AutoSwitch,
			},
		}
		filter = append(filter, item)
	}

	// Accurate or fuzzy search for IP or domain names
	if params.Keyword != "" {
		// accurate search
		if params.SearchMode == "1" {
			if util.IsIP(params.Keyword) {
				item := map[string]interface{}{
					"term": map[string]interface{}{
						"value.data.keyword": params.Keyword,
					},
				}
				filter = append(filter, item)
			} else {
				item := map[string]interface{}{
					"term": map[string]interface{}{
						"domain.keyword": params.Keyword,
					},
				}
				filter = append(filter, item)
			}

		} else { // fuzzy search
			ipItem := map[string]interface{}{
				"wildcard": map[string]interface{}{
					"value.data.keyword": "*" + params.Keyword + "*",
				},
			}
			domainItem := map[string]interface{}{
				"wildcard": map[string]interface{}{
					"domain.keyword": "*" + params.Keyword + "*",
				},
			}
			should = append(should, ipItem, domainItem)
		}
	}

	var query map[string]interface{}
	if len(should) == 0 {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"filter": filter,
				},
			},
		}
	} else {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					// filter or accurate search
					"filter": filter,
					// fuzzy search
					"should":               should,
					"minimum_should_match": 1,
				},
			},
		}
	}

	return es.Search("an_index_name", query, offset, limit)
}
