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

func TestLyrics(t *testing.T) {
	lyrics, err := GetLyrics("GICwp59Hags")

	if err != nil {
		t.Fatal(err)
	}

	if lyrics == "" {
		t.Fatal("lyrics == \"\"")
	}

	lyrics, err = GetLyrics("9Mf6f8TPH_4")

	if err != nil {
		t.Fatal(err)
	}

	if lyrics != "" {
		t.Fatal("lyrics != \"\"")
	}
}

func TestWatchPlaylist(t *testing.T) {
	watchPlaylist, err := GetWatchPlaylist("FM7MFYoylVs")

	if err != nil {
		t.Fatal(err)
	}

	if len(watchPlaylist) == 0 {
		t.Fatal("playlist list is empty")
	}

	// sometimes the playlist will only have 1 song and 1 empty track item
	if len(watchPlaylist) < 3 {
		t.Fatal("len(watchPlaylist) < 3")
	}

	for _, track := range watchPlaylist {
		if track.VideoID == "" {
			t.Fatal("track.VideoID == \"\"")
		}
		if track.Title == "" {
			t.Fatal("track.Title == \"\"")
		}
	}

	t.Log("Hello")

	if err != nil {
		t.Fatal(err)
	}

	if len(watchPlaylist) == 0 {
		t.Fatal("radio list id empty")
	}

	if len(watchPlaylist) < 3 {
		t.Fatal("len(watchPlaylist) < 3")
	}

	for _, track := range watchPlaylist {
		if track.VideoID == "" {
			t.Fatal("track.VideoID == \"\"")
		}
		if track.Title == "" {
			t.Fatal("track.Title == \"\"")
		}
	}
}
