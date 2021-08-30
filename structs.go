package ytmusic

type SearchResult struct {
	Tracks    []*TrackItem    `json:"tracks"`
	Artists   []*ArtistItem   `json:"artists"`
	Albums    []*AlbumItem    `json:"albums"`
	Playlists []*PlaylistItem `json:"playlists"`
	Videos    []*VideoItem    `json:"videos"`
}

type TrackItem struct {
	VideoID    string      `json:"videoId"`
	PlaylistID string      `json:"playlistId"`
	Title      string      `json:"title"`
	Artists    []Artist    `json:"artists"`
	Album      Album       `json:"album"`
	Duration   string      `json:"duration"`
	IsExplicit bool        `json:"isExplicit"`
	Thumbnails []Thumbnail `json:"thumbnails"`
}

type ArtistItem struct {
	BrowseID   string      `json:"browseId"`
	Artist     string      `json:"artist"`
	ShuffleID  string      `json:"shuffleId"`
	RadioID    string      `json:"radioId"`
	Thumbnails []Thumbnail `json:"thumbnails"`
}

type AlbumItem struct {
	BrowseID   string      `json:"browseId"`
	Title      string      `json:"title"`
	Type       string      `json:"type"`
	Artists    []Artist    `json:"artists"`
	Year       string      `json:"year"`
	IsExplicit bool        `json:"isExplicit"`
	Thumbnails []Thumbnail `json:"thumbnails"`
}

type PlaylistItem struct {
	BrowseID   string      `json:"browseId"`
	Title      string      `json:"title"`
	Author     string      `json:"author"`
	ItemCount  string      `json:"itemCount"`
	Thumbnails []Thumbnail `json:"thumbnails"`
}

type VideoItem struct {
	VideoID    string      `json:"videoId"`
	PlaylistID string      `json:"playlistId"`
	Title      string      `json:"title"`
	Artists    []Artist    `json:"artists"`
	Views      string      `json:"views"`
	Duration   string      `json:"duration"`
	Thumbnails []Thumbnail `json:"thumbnails"`
}

type Thumbnail struct {
	URL    string  `json:"url"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type Album struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type Artist struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
