package ytmusic

import (
	"net/http"
)

var HTTPClient = &http.Client{}

func Search(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		Language:     "en",
		Region:       "US",
		SearchFilter: NoFilter,
	}
}

func TrackSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		Language:     "en",
		Region:       "US",
		SearchFilter: TrackFilter,
	}
}

func AlbumSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		Language:     "en",
		Region:       "US",
		SearchFilter: AlbumFilter,
	}
}

func ArtistSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		Language:     "en",
		Region:       "US",
		SearchFilter: ArtistFilter,
	}
}

func PlaylistSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		Language:     "en",
		Region:       "US",
		SearchFilter: PlaylistFilter,
	}
}

func VideoSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		Language:     "en",
		Region:       "US",
		SearchFilter: VideoFilter,
	}
}

func GetWatchPlaylist(videoID string) ([]*TrackItem, error) {
	return getWatchPlaylist(videoID)
}
