# ytmusic
Go library for searching on YouTube Music. Work in progress.

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/raitonoberu/ytmusic"
)

func main() {
	s := ytmusic.Search("death grips")
	result, err := s.Next()
	if err != nil {
		panic(err)
	}

	jsonstr, _ := json.Marshal(result)
	fmt.Println(string(jsonstr))
}
```

<details>
    <summary>Example result</summary>

```json
{
  "tracks":[
    {
      "videoId":"LdTkeqDLNP4",
      "playlistId":"RDAMVMLdTkeqDLNP4",
      "title":"Streaky",
      "artists":[
        {
          "name":"Death Grips",
          "id":"UCvM75EuDkJE-65aH5rkfWmg"
        }
      ],
      "album":{
        "name":"Year Of The Snitch",
        "id":"MPREb_lVa6ZUh2pzt"
      },
      "duration":"2:57",
      "isExplicit":false,
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/2casvjuOmQ3rtzgWjzQX-JMhB-OZu9nsg_vzqBDCviHg8hYoD_O_P38TjbvwfT2AArgn39pKuL4UqMlQ=w60-h60-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/2casvjuOmQ3rtzgWjzQX-JMhB-OZu9nsg_vzqBDCviHg8hYoD_O_P38TjbvwfT2AArgn39pKuL4UqMlQ=w120-h120-l90-rj",
          "width":120,
          "height":120
        }
      ]
    },
    {
      "videoId":"YsSDJ9lkWSk",
      "playlistId":"RDAMVMYsSDJ9lkWSk",
      "title":"Government Plates",
      "artists":[
        {
          "name":"Death Grips",
          "id":"UCvM75EuDkJE-65aH5rkfWmg"
        }
      ],
      "album":{
        "name":"Government Plates",
        "id":"MPREb_qRwFpD0q0JA"
      },
      "duration":"2:42",
      "isExplicit":true,
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/sx0DgmYSEp8xPldr5rYZkxlFBg89AU9V-yLxKdgZhPkssi4vvvZnkdq6FhWqxo9xB0qrc3Hau2gGDFg=w60-h60-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/sx0DgmYSEp8xPldr5rYZkxlFBg89AU9V-yLxKdgZhPkssi4vvvZnkdq6FhWqxo9xB0qrc3Hau2gGDFg=w120-h120-l90-rj",
          "width":120,
          "height":120
        }
      ]
    },
    {
      "videoId":"KBEA6ZGzhbk",
      "playlistId":"RDAMVMKBEA6ZGzhbk",
      "title":"Up My Sleeves",
      "artists":[
        {
          "name":"Death Grips",
          "id":"UCvM75EuDkJE-65aH5rkfWmg"
        }
      ],
      "album":{
        "name":"The Powers That B",
        "id":"MPREb_6KTedIfvZMt"
      },
      "duration":"5:16",
      "isExplicit":true,
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/hpZ66X3wL9T66abAPjP5x4v-a6W-Nkk6hrJ6HIIIMtoBx9YyW1H26Ax3-peirP84uMCbrIvw-FCoZSGk=w60-h60-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/hpZ66X3wL9T66abAPjP5x4v-a6W-Nkk6hrJ6HIIIMtoBx9YyW1H26Ax3-peirP84uMCbrIvw-FCoZSGk=w120-h120-l90-rj",
          "width":120,
          "height":120
        }
      ]
    }
  ],
  "artists":[
    {
      "browseId":"UCvM75EuDkJE-65aH5rkfWmg",
      "artist":"Death Grips",
      "shuffleId":"RDAOIa469pUbW1R4pxYoy_l_NQ",
      "radioId":"RDAOIa469pUbW1R4pxYoy_l_NQ",
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/boo3K8l_2EoOXj18QzaEMlK7ecI8V6WtSvXpCjTbySGH_mL8kxHWl-h2yQAEcpZEfdfr71VJ-eslcqU=w60-h60-p-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/boo3K8l_2EoOXj18QzaEMlK7ecI8V6WtSvXpCjTbySGH_mL8kxHWl-h2yQAEcpZEfdfr71VJ-eslcqU=w120-h120-p-l90-rj",
          "width":120,
          "height":120
        }
      ]
    },
    {
      "browseId":"UCpFDrI9UAMp6cpwtZwzUg-g",
      "artist":"The Bug",
      "shuffleId":"RDAOCqmlt0lB6uJa2xKMY-ppOw",
      "radioId":"RDAOCqmlt0lB6uJa2xKMY-ppOw",
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/ICmQwBidmunTxvAtsyZeFKUUOmmM9_BL8DecU0zHgBASOR3JumnQKpLFzXH3XqSS0KC-5vXLXPn-TY2T=w60-h60-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/ICmQwBidmunTxvAtsyZeFKUUOmmM9_BL8DecU0zHgBASOR3JumnQKpLFzXH3XqSS0KC-5vXLXPn-TY2T=w120-h120-l90-rj",
          "width":120,
          "height":120
        }
      ]
    },
    {
      "browseId":"UCrg2LfxablyfpYkD75hDxMA",
      "artist":"Kero Kero Bonito",
      "shuffleId":"RDAOWXjw4GeAz3E3rzHqNWS52Q",
      "radioId":"RDAOWXjw4GeAz3E3rzHqNWS52Q",
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/oUJfc9sL7EM1wKEGO07ctB7HEUMLC6cnjXacebnA7cIQYEwpgKS6iaeaJF8hdiDctBgPo0Cinbt_G6Yq=w60-h60-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/oUJfc9sL7EM1wKEGO07ctB7HEUMLC6cnjXacebnA7cIQYEwpgKS6iaeaJF8hdiDctBgPo0Cinbt_G6Yq=w120-h120-l90-rj",
          "width":120,
          "height":120
        }
      ]
    },
    {
      "browseId":"UCIzvaEjamDYR85oxnqAB0eQ",
      "artist":"TNGHT",
      "shuffleId":"RDAOYh5LQzA9VF5UJn5C65__0Q",
      "radioId":"RDAOYh5LQzA9VF5UJn5C65__0Q",
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/b03Ha1rd4C81J0y5-YPFnCF4tEMCpVnRBMGB8BQa2E_dkvy1P9Jm3Gk5mc4uXgGTuQyJ7LD5JPNSibSM=w60-h60-p-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/b03Ha1rd4C81J0y5-YPFnCF4tEMCpVnRBMGB8BQa2E_dkvy1P9Jm3Gk5mc4uXgGTuQyJ7LD5JPNSibSM=w120-h120-p-l90-rj",
          "width":120,
          "height":120
        }
      ]
    }
  ],
  "albums":[
    {
      "browseId":"MPREb_FCRKNFJVb1t",
      "title":"No Love Deep Web",
      "type":"Album",
      "artists":[
        {
          "name":"Death Grips",
          "id":"UCvM75EuDkJE-65aH5rkfWmg"
        }
      ],
      "year":"2012",
      "isExplicit":true,
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/RXF4F3ptT_mOITf70OiHipxZ9MWYwqOXxcVNZzswii928wIbTXloel2ngFTtfckhN79aQ41YN1Y7Orl7=w60-h60-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/RXF4F3ptT_mOITf70OiHipxZ9MWYwqOXxcVNZzswii928wIbTXloel2ngFTtfckhN79aQ41YN1Y7Orl7=w120-h120-l90-rj",
          "width":120,
          "height":120
        },
        {
          "url":"https://lh3.googleusercontent.com/RXF4F3ptT_mOITf70OiHipxZ9MWYwqOXxcVNZzswii928wIbTXloel2ngFTtfckhN79aQ41YN1Y7Orl7=w226-h226-l90-rj",
          "width":226,
          "height":226
        },
        {
          "url":"https://lh3.googleusercontent.com/RXF4F3ptT_mOITf70OiHipxZ9MWYwqOXxcVNZzswii928wIbTXloel2ngFTtfckhN79aQ41YN1Y7Orl7=w544-h544-l90-rj",
          "width":544,
          "height":544
        }
      ]
    },
    {
      "browseId":"MPREb_BL9sWaZWAUE",
      "title":"The Money Store",
      "type":"Album",
      "artists":[
        {
          "name":"Death Grips",
          "id":"UCvM75EuDkJE-65aH5rkfWmg"
        }
      ],
      "year":"2012",
      "isExplicit":true,
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/WmWd-wg3LsIjvefwD2VvfqCeFLMmCTZlj8QlGSL373bonYFy_SCZK_qsWBIfRF7vbl3aU0GTGEAErSs=w60-h60-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/WmWd-wg3LsIjvefwD2VvfqCeFLMmCTZlj8QlGSL373bonYFy_SCZK_qsWBIfRF7vbl3aU0GTGEAErSs=w120-h120-l90-rj",
          "width":120,
          "height":120
        },
        {
          "url":"https://lh3.googleusercontent.com/WmWd-wg3LsIjvefwD2VvfqCeFLMmCTZlj8QlGSL373bonYFy_SCZK_qsWBIfRF7vbl3aU0GTGEAErSs=w226-h226-l90-rj",
          "width":226,
          "height":226
        },
        {
          "url":"https://lh3.googleusercontent.com/WmWd-wg3LsIjvefwD2VvfqCeFLMmCTZlj8QlGSL373bonYFy_SCZK_qsWBIfRF7vbl3aU0GTGEAErSs=w544-h544-l90-rj",
          "width":544,
          "height":544
        }
      ]
    },
    {
      "browseId":"MPREb_D89UwANNyA9",
      "title":"Black Paint",
      "type":"Single",
      "artists":[
        {
          "name":"Death Grips",
          "id":"UCvM75EuDkJE-65aH5rkfWmg"
        }
      ],
      "year":"2018",
      "isExplicit":false,
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/2casvjuOmQ3rtzgWjzQX-JMhB-OZu9nsg_vzqBDCviHg8hYoD_O_P38TjbvwfT2AArgn39pKuL4UqMlQ=w60-h60-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/2casvjuOmQ3rtzgWjzQX-JMhB-OZu9nsg_vzqBDCviHg8hYoD_O_P38TjbvwfT2AArgn39pKuL4UqMlQ=w120-h120-l90-rj",
          "width":120,
          "height":120
        },
        {
          "url":"https://lh3.googleusercontent.com/2casvjuOmQ3rtzgWjzQX-JMhB-OZu9nsg_vzqBDCviHg8hYoD_O_P38TjbvwfT2AArgn39pKuL4UqMlQ=w226-h226-l90-rj",
          "width":226,
          "height":226
        },
        {
          "url":"https://lh3.googleusercontent.com/2casvjuOmQ3rtzgWjzQX-JMhB-OZu9nsg_vzqBDCviHg8hYoD_O_P38TjbvwfT2AArgn39pKuL4UqMlQ=w544-h544-l90-rj",
          "width":544,
          "height":544
        }
      ]
    }
  ],
  "playlists":[
    {
      "browseId":"VLPLI7cAsCDsB3F0qHcVGWWAIQAtsvnyhsl0",
      "title":"Death Grips Discography",
      "author":"flubberfan2006",
      "itemCount":"140 songs",
      "thumbnails":[
        {
          "url":"https://yt3.ggpht.com/lyIJJtfJ83ek2QyqnAazs5mh16Deoip1gqVEfMaFd9hCxVzukesU_1KrfbbD3rpneiqpB6quUzs=s192",
          "width":192,
          "height":192
        },
        {
          "url":"https://yt3.ggpht.com/lyIJJtfJ83ek2QyqnAazs5mh16Deoip1gqVEfMaFd9hCxVzukesU_1KrfbbD3rpneiqpB6quUzs=s576",
          "width":576,
          "height":576
        },
        {
          "url":"https://yt3.ggpht.com/lyIJJtfJ83ek2QyqnAazs5mh16Deoip1gqVEfMaFd9hCxVzukesU_1KrfbbD3rpneiqpB6quUzs=s1200",
          "width":1200,
          "height":1200
        }
      ]
    },
    {
      "browseId":"VLPL_tgtrkztDm-WAEZZoK0d6Q5PZ5mJKw6V",
      "title":"Death Grips essentials",
      "author":"FoxInShadow",
      "itemCount":"43 songs",
      "thumbnails":[
        {
          "url":"https://yt3.ggpht.com/Du-ToeL4gSKH901rjNWbHbOkbdZ3-Rc8VCIlGIdtIfXTVLGqbskazl0XuHKdjJSiV-zEVH-qcIc=s192",
          "width":192,
          "height":192
        },
        {
          "url":"https://yt3.ggpht.com/Du-ToeL4gSKH901rjNWbHbOkbdZ3-Rc8VCIlGIdtIfXTVLGqbskazl0XuHKdjJSiV-zEVH-qcIc=s576",
          "width":576,
          "height":576
        },
        {
          "url":"https://yt3.ggpht.com/Du-ToeL4gSKH901rjNWbHbOkbdZ3-Rc8VCIlGIdtIfXTVLGqbskazl0XuHKdjJSiV-zEVH-qcIc=s1200",
          "width":1200,
          "height":1200
        }
      ]
    },
    {
      "browseId":"VLPLNsIB5OmdsLoCvESpB2k2Cik09uGy3K9I",
      "title":"Death Grips - Steroids (Crouching Tiger Hidden Gabber Megamix)",
      "author":"Tsuaretehn_Espira",
      "itemCount":"7 songs",
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/UMh6ZySqP-o/sddefault.jpg?sqp=-oaymwEWCJADEOEBIAQqCghqEJQEGHgg6AJIWg\u0026rs=AMzJL3lm5L1XyhHYIRNfaMBGFZNXfE9lRw",
          "width":400,
          "height":400
        },
        {
          "url":"https://i.ytimg.com/vi/UMh6ZySqP-o/hq720.jpg?sqp=-oaymwEXCKAGEMIDIAQqCwjVARCqCBh4INgESFo\u0026rs=AMzJL3k2YtTE02YvnVM4q0TLFgIVWXCGxg",
          "width":800,
          "height":800
        },
        {
          "url":"https://i.ytimg.com/vi/UMh6ZySqP-o/hq720.jpg?sqp=-oaymwEXCNUGEOADIAQqCwjVARCqCBh4INgESFo\u0026rs=AMzJL3m9cJM8j-WwtpioXdlv6GmiepLgLg",
          "width":853,
          "height":853
        }
      ]
    }
  ],
  "videos":[
    {
      "videoId":"iPtPo8Sa3NE",
      "playlistId":"RDAMVMiPtPo8Sa3NE",
      "title":"Death Grips - Exmilitary [Full Mixtape]",
      "artists":[
        {
          "name":"An0rhu",
          "id":"UCqktZCXcZQf1irrzkertseQ"
        }
      ],
      "views":"5.5M views",
      "duration":"48:31",
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/iPtPo8Sa3NE/sddefault.jpg?sqp=-oaymwEWCJADEOEBIAQqCghqEJQEGHgg6AJIWg\u0026rs=AMzJL3kKmQiuO_3TTh_9Qu1Ac-HiajpISQ",
          "width":400,
          "height":400
        }
      ]
    },
    {
      "videoId":"eu-MFuqymsc",
      "playlistId":"RDAMVMeu-MFuqymsc",
      "title":"Death Grips - NO LOVE DEEP WEB (Full Album)",
      "artists":[
        {
          "name":"Death Grips",
          "id":"UCuq1H-HXWoW4JL-hX5bWxzw"
        }
      ],
      "views":"1.2M views",
      "duration":"45:48",
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/eu-MFuqymsc/sddefault.jpg?sqp=-oaymwEWCJADEOEBIAQqCghqEJQEGHgg6AJIWg\u0026rs=AMzJL3mFRLwkFtB-cTJ8_mkElBXSSh3MnA",
          "width":400,
          "height":400
        }
      ]
    },
    {
      "videoId":"sticXkHxZC4",
      "playlistId":"RDAMVMsticXkHxZC4",
      "title":"I've Seen Footage",
      "artists":[
        {
          "name":"Death Grips",
          "id":"UCvM75EuDkJE-65aH5rkfWmg"
        }
      ],
      "views":"1.2M views",
      "duration":"3:22",
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/sticXkHxZC4/sddefault.jpg?sqp=-oaymwEWCJADEOEBIAQqCghqEJQEGHgg6AJIWg\u0026rs=AMzJL3lRYBDh7tM2VlK0t7yj5TJqrfS1jg",
          "width":400,
          "height":400
        }
      ]
    }
  ]
}
```
</details>