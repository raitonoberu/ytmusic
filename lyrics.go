package ytmusic

import (
	"fmt"
	"net/url"
)

func getLyrics(videoID string) (string, error) {
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
		return "", err
	}

	lyricsBrowseIdIface := getValue(page, path{"contents", "singleColumnMusicWatchNextResultsRenderer", "tabbedRenderer", "watchNextTabbedResultsRenderer", "tabs", 1, "tabRenderer", "endpoint", "browseEndpoint", "browseId"})
	if lyricsBrowseIdIface == nil {
		return "", fmt.Errorf("couldn't extract lyrics browseId")
	}

	lyricsBrowseID := lyricsBrowseIdIface.(string)

	return lyricsBrowseID, nil
}
