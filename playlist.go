package ytmusic

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

func getWatchPlaylist(videoID string) ([]*TrackItem, error) {
	payload := map[string]interface{}{
		"context": map[string]interface{}{
			"client": map[string]interface{}{
				"clientName":    "WEB_REMIX",
				"clientVersion": "1.20210823.00.00",
			},
			"user": map[string]interface{}{
				"lockedSafetyMode": false,
			},
			"hl": "en",
			"gl": "us",
		},
		"videoId":                       videoID,
		"enablePersistentPlaylistPanel": true,
		"isAudioOnly":                   true,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("key", searchKey)

	request, err := http.NewRequest("POST", "https://music.youtube.com/youtubei/v1/next?key="+searchKey, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	request.Header = requestHeader
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return parsePlaylistPage(result), nil
}
