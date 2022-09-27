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

	browseIdIface := getValue(page, path{"contents", "singleColumnMusicWatchNextResultsRenderer", "tabbedRenderer", "watchNextTabbedResultsRenderer", "tabs", 1, "tabRenderer", "endpoint", "browseEndpoint", "browseId"})
	if browseIdIface == nil {
		return "", fmt.Errorf("couldn't extract lyrics browseId")
	}
	browseID := browseIdIface.(string)

	page, err = makeRequest(
		"browse",
		map[string]interface{}{
			"browseId": browseID,
		},
		url.Values{},
	)
	if err != nil {
		return "", err
	}

	lyrics := ""
	if lyricsIface := getValue(page, path{"contents", "sectionListRenderer", "contents", 0, "musicDescriptionShelfRenderer", "description", "runs", 0, "text"}); lyricsIface != nil {
		lyrics = lyricsIface.(string)
	}

	return lyrics, nil
}
