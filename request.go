package ytmusic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func makeRequest(endpoint string, body map[string]interface{}, params url.Values) (interface{}, error) {
	payload := map[string]interface{}{
		"context": map[string]interface{}{
			"client": map[string]interface{}{
				"clientName":    "WEB_REMIX",
				"clientVersion": "1.20220715.04.00",
				"hl":            Language,
				"gl":            Region,
			},
			"user": map[string]interface{}{
				"lockedSafetyMode": false,
			},
		},
	}
	for key, item := range body {
		payload[key] = item
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	params.Add("key", searchKey)

	request, err := http.NewRequest("POST", fmt.Sprintf("https://music.youtube.com/youtubei/v1/%s?%s", endpoint, params.Encode()), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	request.Header = requestHeader
	response, err := HTTPClient.Do(request)
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
