package ytmusic

import (
	"net/url"
)

func getWatchPlaylist(videoID string) ([]*TrackItem, error) {
	page, err := makeRequest(
		"next",
		map[string]interface{}{
			"videoId":                       videoID,
			"playlistId":                    "RDAMVM" + videoID,
			"enablePersistentPlaylistPanel": true,
			"isAudioOnly":                   true,
		},
		url.Values{},
	)

	if err != nil {
		return nil, err
	}

	return parseWatchPlaylist(page), nil
}
