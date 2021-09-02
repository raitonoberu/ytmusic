package ytmusic

import "net/url"

func getSearchSuggestions(input string) ([]string, error) {
	page, err := makeRequest(
		"music/get_search_suggestions",
		map[string]interface{}{
			"input": input,
		},
		url.Values{},
	)

	if err != nil {
		return nil, err
	}

	return parseSearchSuggestions(page), err
}
