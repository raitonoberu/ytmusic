package ytmusic

import (
	"net/http"
)

var Language = "en"
var Region = "US"
var HTTPClient = &http.Client{}

func Search(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		SearchFilter: NoFilter,
	}
}

func TrackSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		SearchFilter: TrackFilter,
	}
}

func AlbumSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		SearchFilter: AlbumFilter,
	}
}

func ArtistSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		SearchFilter: ArtistFilter,
	}
}

func PlaylistSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		SearchFilter: PlaylistFilter,
	}
}

func VideoSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		SearchFilter: VideoFilter,
	}
}

func GetWatchPlaylist(videoID string) ([]*TrackItem, error) {
	return getWatchPlaylist(videoID)
}
