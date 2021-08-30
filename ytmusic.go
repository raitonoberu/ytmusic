package ytmusic

import "net/http"

func Search(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		Language:     "en",
		Region:       "US",
		SearchFilter: NoFilter,
		HTTPClient:   &http.Client{},
	}
}

func TrackSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		Language:     "en",
		Region:       "US",
		SearchFilter: TrackFilter,
		HTTPClient:   &http.Client{},
	}
}

func AlbumSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		Language:     "en",
		Region:       "US",
		SearchFilter: AlbumFilter,
		HTTPClient:   &http.Client{},
	}
}

func ArtistSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		Language:     "en",
		Region:       "US",
		SearchFilter: ArtistFilter,
		HTTPClient:   &http.Client{},
	}
}

func PlaylistSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		Language:     "en",
		Region:       "US",
		SearchFilter: PlaylistFilter,
		HTTPClient:   &http.Client{},
	}
}

func VideoSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		Language:     "en",
		Region:       "US",
		SearchFilter: VideoFilter,
		HTTPClient:   &http.Client{},
	}
}
