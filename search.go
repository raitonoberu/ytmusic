package ytmusic

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type SearchClient struct {
	Query string

	Language string
	Region   string

	SearchFilter SearchFilter
	HTTPClient   *http.Client

	continuationKey string
	newPage         bool
}

func (search *SearchClient) makeRequest() (interface{}, error) {
	payload := map[string]interface{}{
		"context": map[string]map[string]interface{}{
			"client": {
				"clientName":    "WEB_REMIX",
				"clientVersion": "1.20210823.00.00",
			},
			"user": {
				"lockedSafetyMode": false,
			},
		},
	}
	clientData := payload["context"].(map[string]map[string]interface{})["client"]
	clientData["hl"] = search.Language
	clientData["gl"] = search.Region
	if search.continuationKey == "" {
		payload["query"] = search.Query
		if search.SearchFilter != NoFilter {
			payload["params"] = string(search.SearchFilter)
		}
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("key", searchKey)
	if search.continuationKey != "" {
		params.Add("ctoken", search.continuationKey)
		params.Add("continuation", search.continuationKey)
		params.Add("type", "next")
	}

	request, err := http.NewRequest("POST", "https://music.youtube.com/youtubei/v1/search?"+params.Encode(), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	request.Header = requestHeader
	response, err := search.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
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
	result, key := parsePage(page)
	if result == nil {
		return nil, errors.New("couldn't parse page")
	}
	search.continuationKey = key
	search.newPage = true
	return result, nil
}
