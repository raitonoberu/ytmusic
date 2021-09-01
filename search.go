package ytmusic

import (
	"errors"
	"net/url"
)

type SearchClient struct {
	Query string

	Language string
	Region   string

	SearchFilter SearchFilter

	continuationKey string
	newPage         bool
}

func (search *SearchClient) makeRequest() (interface{}, error) {
	body := map[string]interface{}{}

	if search.continuationKey == "" {
		body["query"] = search.Query
		if search.SearchFilter != NoFilter {
			body["params"] = string(search.SearchFilter)
		}
	}

	params := url.Values{}
	if search.continuationKey != "" {
		params.Add("ctoken", search.continuationKey)
		params.Add("continuation", search.continuationKey)
		params.Add("type", "next")
	}

	page, err := makeRequest(
		"search",
		search.Language,
		search.Region,
		body,
		params,
	)
	if err != nil {
		return nil, err
	}

	return page, nil
}

func (search *SearchClient) NextExists() bool {
	if !search.newPage {
		return true
	}
	if search.continuationKey != "" {
		return true
	}
	return false
}

func (search *SearchClient) Next() (*SearchResult, error) {
	if !search.NextExists() {
		return nil, errors.New("end reached")
	}

	page, err := search.makeRequest()
	if err != nil {
		return nil, err
	}
	result, key := parseSearchPage(page)
	if result == nil {
		return nil, errors.New("couldn't parse page")
	}
	search.continuationKey = key
	search.newPage = true
	return result, nil
}
