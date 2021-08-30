package ytmusic

type SearchFilter string

// Credits to github.com/sigma67/ytmusicapi
const (
	NoFilter                SearchFilter = ""
	TrackFilter             SearchFilter = "EgWKAQIIAWoMEA4QChADEAQQCRAF"
	AlbumFilter             SearchFilter = "EgWKAQIYAWoMEA4QChADEAQQCRAF"
	ArtistFilter            SearchFilter = "EgWKAQIgAWoMEA4QChADEAQQCRAF"
	PlaylistFilter          SearchFilter = "Eg-KAQwIABAAGAAgACgBMABqChAEEAMQCRAFEAo%3D"
	VideoFilter             SearchFilter = "EgWKAQIQAWoMEA4QChADEAQQCRAF"
	FeaturedPlaylistFilter  SearchFilter = "EgeKAQQoADgBagwQDhAKEAMQBBAJEAU%3D"
	CommunityPlaylistFilter SearchFilter = "EgeKAQQoAEABagwQDhAKEAMQBBAJEAU%3D"
)
