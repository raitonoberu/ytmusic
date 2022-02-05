package ytmusic

import (
	"testing"
)

func TestSearch(t *testing.T) {
	s := Search("ncs")
	r, err := s.Next()
	if err != nil {
		t.Fatal(err)
	}

	if len(r.Tracks) == 0 {
		t.Fatal("len(r.Tracks) == 0")
	}
	if r.Tracks[0].VideoID == "" {
		t.Fatal("r.Tracks[0].VideoID == \"\"")
	}
	if len(r.Tracks[0].Artists) == 0 {
		t.Fatal("len(r.Tracks[0].Artists) == 0")
	}

	if len(r.Videos) == 0 {
		t.Fatal("len(r.Videos) == 0")
	}
	if r.Videos[0].VideoID == "" {
		t.Fatal("r.Videos[0].VideoID == \"\"")
	}

	if len(r.Artists) == 0 {
		t.Fatal("len(r.Artists) == 0")
	}
	if r.Artists[0].BrowseID == "" {
		t.Fatal("r.Artists[0].BrowseID == \"\"")
	}

	if len(r.Playlists) == 0 {
		t.Fatal("len(r.Playlists) == 0")
	}
	if r.Playlists[0].BrowseID == "" {
		t.Fatal("r.Playlists[0].BrowseID == \"\"")
	}

	if len(r.Albums) == 0 {
		t.Fatal("len(r.Albums) == 0")
	}
	if r.Albums[0].BrowseID == "" {
		t.Fatal("r.Albums[0].BrowseID == \"\"")
	}
}
