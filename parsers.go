package ytmusic

type path []interface{}

func getValue(source interface{}, path path) interface{} {
	value := source
	for _, element := range path {
		mustBreak := false
		switch element.(type) {
		case string:
			if val, ok := value.(map[string]interface{})[element.(string)]; ok {
				value = val
			} else {
				value = nil
				mustBreak = true
			}
		case int:
			if len(value.([]interface{})) != 0 {
				value = value.([]interface{})[element.(int)]
			} else {
				value = nil
				mustBreak = true
			}
		}
		if mustBreak {
			break
		}
	}
	return value
}

type itemType int

const (
	trackItemType itemType = iota + 1
	artistItemType
	albumItemType
	playlistItemType
	videoItemType
)

func parseSearchPage(page interface{}) (*SearchResult, string) {
	result := &SearchResult{}
	var continuationKey string

	var shelves interface{}
	shelves = getValue(page, path{"contents", "tabbedSearchResultsRenderer", "tabs", 0, "tabRenderer", "content", "sectionListRenderer", "contents"})
	if shelves == nil {
		if continuationContents := getValue(page, path{"continuationContents"}); continuationContents != nil {
			shelves = []interface{}{continuationContents}
		} else {
			return nil, ""
		}
	}

	for _, shelf := range shelves.([]interface{}) {
		content, itype, key := parseShelf(shelf)
		if key != "" {
			continuationKey = key
		}

		switch itype {
		case trackItemType:
			{
				trackItems := make([]*TrackItem, len(content))
				for i, item := range content {
					trackItems[i] = item.(*TrackItem)
				}
				result.Tracks = append(result.Tracks, trackItems...)
			}
		case artistItemType:
			{
				artistItems := make([]*ArtistItem, len(content))
				for i, item := range content {
					artistItems[i] = item.(*ArtistItem)
				}
				result.Artists = append(result.Artists, artistItems...)
			}
		case albumItemType:
			{
				albumItems := make([]*AlbumItem, len(content))
				for i, item := range content {
					albumItems[i] = item.(*AlbumItem)
				}
				result.Albums = append(result.Albums, albumItems...)
			}
		case playlistItemType:
			{
				playlistItems := make([]*PlaylistItem, len(content))
				for i, item := range content {
					playlistItems[i] = item.(*PlaylistItem)
				}
				result.Playlists = append(result.Playlists, playlistItems...)
			}
		case videoItemType:
			{
				videoItems := make([]*VideoItem, len(content))
				for i, item := range content {
					videoItems[i] = item.(*VideoItem)
				}
				result.Videos = append(result.Videos, videoItems...)
			}
		}
	}
	return result, continuationKey
}

func parseShelf(shelf interface{}) ([]interface{}, itemType, string) {
	var shelfContents []interface{}

	if contents := getValue(shelf, path{"musicShelfRenderer", "contents"}); contents != nil {
		shelfContents = contents.([]interface{})
	} else if contents := getValue(shelf, path{"musicShelfContinuation", "contents"}); contents != nil {
		shelfContents = contents.([]interface{})
	} else {
		return nil, 0, ""
	}

	var itype itemType
	if pageType := getValue(shelfContents, path{0, "musicResponsiveListItemRenderer", "navigationEndpoint", "browseEndpoint", "browseEndpointContextSupportedConfigs", "browseEndpointContextMusicConfig", "pageType"}); pageType != nil {
		pageType := pageType.(string)
		switch pageType {
		case "MUSIC_PAGE_TYPE_ARTIST":
			{
				itype = artistItemType
			}
		case "MUSIC_PAGE_TYPE_ALBUM":
			{
				itype = albumItemType
			}
		case "MUSIC_PAGE_TYPE_PLAYLIST":
			{
				itype = playlistItemType
			}
		}
	} else if videoId := getValue(shelfContents, path{0, "musicResponsiveListItemRenderer", "playlistItemData", "videoId"}); videoId != nil {
		if videoType := getValue(shelfContents, path{0, "musicResponsiveListItemRenderer", "overlay", "musicItemThumbnailOverlayRenderer", "content", "musicPlayButtonRenderer", "playNavigationEndpoint", "watchEndpoint", "watchEndpointMusicSupportedConfigs", "watchEndpointMusicConfig", "musicVideoType"}); videoType != nil {
			videoType := videoType.(string)
			switch videoType {
			case "MUSIC_VIDEO_TYPE_ATV":
				{
					itype = trackItemType
				}
			case "MUSIC_VIDEO_TYPE_OMV", "MUSIC_VIDEO_TYPE_UGC":
				{
					itype = videoItemType
				}
			}
		}
	}
	if itype == 0 {
		return nil, 0, ""
	}

	result := []interface{}{}

	for _, item := range shelfContents {
		switch itype {
		case trackItemType:
			{
				result = append(result, parseTrackItem(item))
			}
		case artistItemType:
			{
				result = append(result, parseArtistItem(item))
			}
		case albumItemType:
			{
				result = append(result, parseAlbumItem(item))
			}
		case playlistItemType:
			{
				result = append(result, parsePlaylistItem(item))
			}
		case videoItemType:
			{
				result = append(result, parseVideoItem(item))
			}
		}
	}

	var continuationKey string
	if key := getValue(shelf, path{"musicShelfRenderer", "continuations", 0, "nextContinuationData", "continuation"}); key != nil {
		continuationKey = key.(string)
	} else if key := getValue(shelf, path{"musicShelfContinuation", "continuations", 0, "nextContinuationData", "continuation"}); key != nil {
		continuationKey = key.(string)
	}
	return result, itype, continuationKey
}

func parseTrackItem(trackItem interface{}) *TrackItem {
	track := &TrackItem{}

	if info1 := getValue(trackItem, path{"musicResponsiveListItemRenderer", "flexColumns", 0, "musicResponsiveListItemFlexColumnRenderer", "text", "runs", 0}); info1 != nil {
		if title := getValue(info1, path{"text"}); title != nil {
			track.Title = title.(string)
		}
		if videoId := getValue(info1, path{"navigationEndpoint", "watchEndpoint", "videoId"}); videoId != nil {
			track.VideoID = videoId.(string)
		}
		if playlistId := getValue(info1, path{"navigationEndpoint", "watchEndpoint", "playlistId"}); playlistId != nil {
			track.PlaylistID = playlistId.(string)
		}
	}

	if info2 := getValue(trackItem, path{"musicResponsiveListItemRenderer", "flexColumns", 1, "musicResponsiveListItemFlexColumnRenderer", "text", "runs"}); info2 != nil {
		for _, run := range info2.([]interface{}) {
			if pageType := getValue(run, path{"navigationEndpoint", "browseEndpoint", "browseEndpointContextSupportedConfigs", "browseEndpointContextMusicConfig", "pageType"}); pageType != nil {
				switch pageType.(string) {
				case "MUSIC_PAGE_TYPE_ARTIST":
					{
						artist := Artist{}
						if name := getValue(run, path{"text"}); name != nil {
							artist.Name = name.(string)
						}
						if id := getValue(run, path{"navigationEndpoint", "browseEndpoint", "browseId"}); id != nil {
							artist.ID = id.(string)
						}
						track.Artists = append(track.Artists, artist)
					}
				case "MUSIC_PAGE_TYPE_ALBUM":
					{
						if name := getValue(run, path{"text"}); name != nil {
							track.Album.Name = name.(string)
						}
						if id := getValue(run, path{"navigationEndpoint", "browseEndpoint", "browseId"}); id != nil {
							track.Album.ID = id.(string)
						}
					}
				}
			}
		}
		length := len(info2.([]interface{}))
		if length != 0 {
			if duration := getValue(info2.([]interface{})[length-1], path{"text"}); duration != nil {
				track.Duration = durationToInt(duration.(string))
			}
		}
	}
	if explicit := getValue(trackItem, path{"musicResponsiveListItemRenderer", "badges", 0, "musicInlineBadgeRenderer", "icon", "iconType"}); explicit != nil {
		if explicit.(string) == "MUSIC_EXPLICIT_BADGE" {
			track.IsExplicit = true
		}
	}
	if thumbnails := getValue(trackItem, path{"musicResponsiveListItemRenderer", "thumbnail", "musicThumbnailRenderer", "thumbnail", "thumbnails"}); thumbnails != nil {
		for _, thumbnail := range thumbnails.([]interface{}) {
			t := Thumbnail{}
			if turl := getValue(thumbnail, path{"url"}); turl != nil {
				t.URL = turl.(string)
			}
			if tw := getValue(thumbnail, path{"width"}); tw != nil {
				t.Width = int(tw.(float64))
			}
			if th := getValue(thumbnail, path{"width"}); th != nil {
				t.Height = int(th.(float64))
			}
			track.Thumbnails = append(track.Thumbnails, t)
		}
	}

	return track
}

func parseArtistItem(artistItem interface{}) *ArtistItem {
	artist := &ArtistItem{}

	if browseId := getValue(artistItem, path{"musicResponsiveListItemRenderer", "navigationEndpoint", "browseEndpoint", "browseId"}); browseId != nil {
		artist.BrowseID = browseId.(string)
	}
	if name := getValue(artistItem, path{"musicResponsiveListItemRenderer", "flexColumns", 0, "musicResponsiveListItemFlexColumnRenderer", "text", "runs", 0, "text"}); name != nil {
		artist.Artist = name.(string)
	}
	if menuItems := getValue(artistItem, path{"musicResponsiveListItemRenderer", "menu", "menuRenderer", "items"}); menuItems != nil {
		if shuffleId := getValue(menuItems, path{0, "menuNavigationItemRenderer", "navigationEndpoint", "watchPlaylistEndpoint", "playlistId"}); shuffleId != nil {
			artist.ShuffleID = shuffleId.(string)
		}
		if radioId := getValue(menuItems, path{0, "menuNavigationItemRenderer", "navigationEndpoint", "watchPlaylistEndpoint", "playlistId"}); radioId != nil {
			artist.RadioID = radioId.(string)
		}
	}
	if thumbnails := getValue(artistItem, path{"musicResponsiveListItemRenderer", "thumbnail", "musicThumbnailRenderer", "thumbnail", "thumbnails"}); thumbnails != nil {
		for _, thumbnail := range thumbnails.([]interface{}) {
			t := Thumbnail{}
			if turl := getValue(thumbnail, path{"url"}); turl != nil {
				t.URL = turl.(string)
			}
			if tw := getValue(thumbnail, path{"width"}); tw != nil {
				t.Width = int(tw.(float64))
			}
			if th := getValue(thumbnail, path{"width"}); th != nil {
				t.Height = int(th.(float64))
			}
			artist.Thumbnails = append(artist.Thumbnails, t)
		}
	}

	return artist
}

func parseAlbumItem(albumItem interface{}) *AlbumItem {
	album := &AlbumItem{}

	if columns := getValue(albumItem, path{"musicResponsiveListItemRenderer", "flexColumns"}); columns != nil {
		if title := getValue(columns, path{0, "musicResponsiveListItemFlexColumnRenderer", "text", "runs", 0, "text"}); title != nil {
			album.Title = title.(string)
		}
		if runs := getValue(columns, path{1, "musicResponsiveListItemFlexColumnRenderer", "text", "runs"}); runs != nil {
			if albumType := getValue(runs, path{0, "text"}); albumType != nil {
				album.Type = albumType.(string)
			}
			for _, run := range runs.([]interface{}) {
				if pageType := getValue(run, path{"navigationEndpoint", "browseEndpoint", "browseEndpointContextSupportedConfigs", "browseEndpointContextMusicConfig", "pageType"}); pageType != nil {
					if pageType.(string) == "MUSIC_PAGE_TYPE_ARTIST" {
						artist := Artist{}
						if name := getValue(run, path{"text"}); name != nil {
							artist.Name = name.(string)
						}
						if id := getValue(run, path{"navigationEndpoint", "browseEndpoint", "browseId"}); id != nil {
							artist.ID = id.(string)
						}
						album.Artists = append(album.Artists, artist)
					}
				}
			}
			if len(album.Artists) == 0 {
				// no browseID
				if artistName := getValue(runs, path{2, "text"}); artistName != nil {
					album.Artists = []Artist{{Name: artistName.(string)}}
				}
			}
			length := len(runs.([]interface{}))
			if length != 0 {
				if year := getValue(runs.([]interface{})[length-1], path{"text"}); year != nil {
					album.Year = year.(string)
				}
			}
		}
	}
	if browseId := getValue(albumItem, path{"musicResponsiveListItemRenderer", "navigationEndpoint", "browseEndpoint", "browseId"}); browseId != nil {
		album.BrowseID = browseId.(string)
	}
	if explicit := getValue(albumItem, path{"musicResponsiveListItemRenderer", "badges", 0, "musicInlineBadgeRenderer", "icon", "iconType"}); explicit != nil {
		if explicit.(string) == "MUSIC_EXPLICIT_BADGE" {
			// not tested
			album.IsExplicit = true
		}
	}
	if thumbnails := getValue(albumItem, path{"musicResponsiveListItemRenderer", "thumbnail", "musicThumbnailRenderer", "thumbnail", "thumbnails"}); thumbnails != nil {
		for _, thumbnail := range thumbnails.([]interface{}) {
			t := Thumbnail{}
			if turl := getValue(thumbnail, path{"url"}); turl != nil {
				t.URL = turl.(string)
			}
			if tw := getValue(thumbnail, path{"width"}); tw != nil {
				t.Width = int(tw.(float64))
			}
			if th := getValue(thumbnail, path{"width"}); th != nil {
				t.Height = int(th.(float64))
			}
			album.Thumbnails = append(album.Thumbnails, t)
		}
	}

	return album
}

func parsePlaylistItem(playlistItem interface{}) *PlaylistItem {
	playlist := &PlaylistItem{}

	if columns := getValue(playlistItem, path{"musicResponsiveListItemRenderer", "flexColumns"}); columns != nil {
		if title := getValue(columns, path{0, "musicResponsiveListItemFlexColumnRenderer", "text", "runs", 0, "text"}); title != nil {
			playlist.Title = title.(string)
		}
		if author := getValue(columns, path{1, "musicResponsiveListItemFlexColumnRenderer", "text", "runs", 2, "text"}); author != nil {
			playlist.Author = author.(string)
		}
		if itemCount := getValue(columns, path{1, "musicResponsiveListItemFlexColumnRenderer", "text", "runs", 4, "text"}); itemCount != nil {
			playlist.ItemCount = itemCount.(string)
		}
	}
	if browseId := getValue(playlistItem, path{"musicResponsiveListItemRenderer", "navigationEndpoint", "browseEndpoint", "browseId"}); browseId != nil {
		playlist.BrowseID = browseId.(string)
	}
	if thumbnails := getValue(playlistItem, path{"musicResponsiveListItemRenderer", "thumbnail", "musicThumbnailRenderer", "thumbnail", "thumbnails"}); thumbnails != nil {
		for _, thumbnail := range thumbnails.([]interface{}) {
			t := Thumbnail{}
			if turl := getValue(thumbnail, path{"url"}); turl != nil {
				t.URL = turl.(string)
			}
			if tw := getValue(thumbnail, path{"width"}); tw != nil {
				t.Width = int(tw.(float64))
			}
			if th := getValue(thumbnail, path{"width"}); th != nil {
				t.Height = int(th.(float64))
			}
			playlist.Thumbnails = append(playlist.Thumbnails, t)
		}
	}

	return playlist
}

func parseVideoItem(videoItem interface{}) *VideoItem {
	video := &VideoItem{}

	if info1 := getValue(videoItem, path{"musicResponsiveListItemRenderer", "flexColumns", 0, "musicResponsiveListItemFlexColumnRenderer", "text", "runs", 0}); info1 != nil {
		if title := getValue(info1, path{"text"}); title != nil {
			video.Title = title.(string)
		}
		if videoId := getValue(info1, path{"navigationEndpoint", "watchEndpoint", "videoId"}); videoId != nil {
			video.VideoID = videoId.(string)
		}
		if playlistId := getValue(info1, path{"navigationEndpoint", "watchEndpoint", "playlistId"}); playlistId != nil {
			video.PlaylistID = playlistId.(string)
		}
	}

	if info2 := getValue(videoItem, path{"musicResponsiveListItemRenderer", "flexColumns", 1, "musicResponsiveListItemFlexColumnRenderer", "text", "runs"}); info2 != nil {
		for _, run := range info2.([]interface{}) {
			if pageType := getValue(run, path{"navigationEndpoint", "browseEndpoint", "browseEndpointContextSupportedConfigs", "browseEndpointContextMusicConfig", "pageType"}); pageType != nil {
				if pageType.(string) == "MUSIC_PAGE_TYPE_ARTIST" || pageType.(string) == "MUSIC_PAGE_TYPE_USER_CHANNEL" {
					artist := Artist{}
					if name := getValue(run, path{"text"}); name != nil {
						artist.Name = name.(string)
					}
					if id := getValue(run, path{"navigationEndpoint", "browseEndpoint", "browseId"}); id != nil {
						artist.ID = id.(string)
					}
					video.Artists = append(video.Artists, artist)
				}
			}
		}
		length := len(info2.([]interface{}))
		if length != 0 {
			if duration := getValue(info2.([]interface{})[length-1], path{"text"}); duration != nil {
				video.Duration = durationToInt(duration.(string))
			}
			if views := getValue(info2.([]interface{})[length-3], path{"text"}); views != nil {
				video.Views = views.(string)
			}
		}
	}
	if thumbnails := getValue(videoItem, path{"musicResponsiveListItemRenderer", "thumbnail", "musicThumbnailRenderer", "thumbnail", "thumbnails"}); thumbnails != nil {
		for _, thumbnail := range thumbnails.([]interface{}) {
			t := Thumbnail{}
			if turl := getValue(thumbnail, path{"url"}); turl != nil {
				t.URL = turl.(string)
			}
			if tw := getValue(thumbnail, path{"width"}); tw != nil {
				t.Width = int(tw.(float64))
			}
			if th := getValue(thumbnail, path{"width"}); th != nil {
				t.Height = int(th.(float64))
			}
			video.Thumbnails = append(video.Thumbnails, t)
		}
	}

	return video
}

func parseWatchPlaylist(page interface{}) []*TrackItem {
	contents := getValue(page, path{"contents", "singleColumnMusicWatchNextResultsRenderer", "tabbedRenderer", "watchNextTabbedResultsRenderer", "tabs", 0, "tabRenderer", "content", "musicQueueRenderer", "content", "playlistPanelRenderer", "contents"})
	if contents == nil {
		return nil
	}

	result := make([]*TrackItem, len(contents.([]interface{})))

	for index, item := range contents.([]interface{}) {
		result[index] = parsePlaylistTrack(item)
	}
	return result
}

func parsePlaylistTrack(trackItem interface{}) *TrackItem {
	track := &TrackItem{
		Thumbnails: nil,
	}

	if title := getValue(trackItem, path{"playlistPanelVideoRenderer", "title", "runs", 0, "text"}); title != nil {
		track.Title = title.(string)
	}

	if info1 := getValue(trackItem, path{"playlistPanelVideoRenderer", "navigationEndpoint", "watchEndpoint"}); info1 != nil {
		if videoId := getValue(info1, path{"videoId"}); videoId != nil {
			track.VideoID = videoId.(string)
		}
		if playlistId := getValue(info1, path{"playlistId"}); playlistId != nil {
			track.PlaylistID = playlistId.(string)
		}
	}

	if info2 := getValue(trackItem, path{"playlistPanelVideoRenderer", "longBylineText", "runs"}); info2 != nil {
		for _, run := range info2.([]interface{}) {
			if pageType := getValue(run, path{"navigationEndpoint", "browseEndpoint", "browseEndpointContextSupportedConfigs", "browseEndpointContextMusicConfig", "pageType"}); pageType != nil {
				switch pageType.(string) {
				case "MUSIC_PAGE_TYPE_ARTIST":
					{
						artist := Artist{}
						if name := getValue(run, path{"text"}); name != nil {
							artist.Name = name.(string)
						}
						if id := getValue(run, path{"navigationEndpoint", "browseEndpoint", "browseId"}); id != nil {
							artist.ID = id.(string)
						}
						track.Artists = append(track.Artists, artist)
					}
				case "MUSIC_PAGE_TYPE_ALBUM":
					{
						if name := getValue(run, path{"text"}); name != nil {
							track.Album.Name = name.(string)
						}
						if id := getValue(run, path{"navigationEndpoint", "browseEndpoint", "browseId"}); id != nil {
							track.Album.ID = id.(string)
						}
					}
				}
			}
		}
	}
	if duration := getValue(trackItem, path{"playlistPanelVideoRenderer", "lengthText", "runs", 0, "text"}); duration != nil {
		track.Duration = durationToInt(duration.(string))
	}
	if explicit := getValue(trackItem, path{"playlistPanelVideoRenderer", "badges", 0, "musicInlineBadgeRenderer", "icon", "iconType"}); explicit != nil {
		if explicit.(string) == "MUSIC_EXPLICIT_BADGE" {
			track.IsExplicit = true
		}
	}
	if thumbnails := getValue(trackItem, path{"playlistPanelVideoRenderer", "thumbnail", "thumbnails"}); thumbnails != nil {
		for _, thumbnail := range thumbnails.([]interface{}) {
			t := Thumbnail{}
			if turl := getValue(thumbnail, path{"url"}); turl != nil {
				t.URL = turl.(string)
			}
			if tw := getValue(thumbnail, path{"width"}); tw != nil {
				t.Width = int(tw.(float64))
			}
			if th := getValue(thumbnail, path{"width"}); th != nil {
				t.Height = int(int(th.(float64)))
			}
			track.Thumbnails = append(track.Thumbnails, t)
		}
	}

	return track
}

func parseSearchSuggestions(page interface{}) []string {
	contents := getValue(page, path{"contents", 0, "searchSuggestionsSectionRenderer", "contents"})
	if contents == nil {
		return nil
	}

	result := make([]string, len(contents.([]interface{})))

	for index, item := range contents.([]interface{}) {
		if query := getValue(item, path{"searchSuggestionRenderer", "navigationEndpoint", "searchEndpoint", "query"}); query != nil {
			result[index] = query.(string)
		}
	}
	return result
}
