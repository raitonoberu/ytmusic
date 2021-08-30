package ytmusic

var requestHeader = map[string][]string{
	"Content-Type": {"application/json; charset=utf-8"},
	"Referer":      {"https://music.youtube.com/search"},
	"User-Agent": {
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36",
	},
}

const searchKey = "AIzaSyC9XL3ZjWddXya6X74dJoCTL-WEYFDNX30"
