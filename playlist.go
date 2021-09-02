package ytmusic

import (
	"net/url"
)

func getWatchPlaylist(videoID string) ([]*TrackItem, error) {
	page, err := makeRequest(
		"next",
		map[string]interface{}{
			"videoId":                       videoID,
			"enablePersistentPlaylistPanel": true,
			"isAudioOnly":                   true,
		},
		url.Values{},
	)

	if err != nil {
		return nil, err
	}

	return parsePlaylistPage(page), nil
}
