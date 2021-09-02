# ytmusic

### Go library for searching on YouTube Music and getting other useful information.

## Installing

```bash
go get github.com/raitonoberu/ytmusic
```

## Usage

### Search for tracks, artists, albums, playlists and videos

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
  "tracks": [
    {
      "videoId": "v54KJ9H2nnc",
      "playlistId": "RDAMVMv54KJ9H2nnc",
      "title": "Fuck That",
      "artists": [{ "name": "Death Grips", "id": "UCvM75EuDkJE-65aH5rkfWmg" }],
      "album": { "name": "The Money Store", "id": "MPREb_BL9sWaZWAUE" },
      "duration": 144,
      "isExplicit": true,
      "thumbnails": [
        {
          "url": "https://lh3.googleusercontent.com/WmWd-wg3LsIjvefwD2VvfqCeFLMmCTZlj8QlGSL373bonYFy_SCZK_qsWBIfRF7vbl3aU0GTGEAErSs=w60-h60-l90-rj",
          "width": 60,
          "height": 60
        },
        {
          "url": "https://lh3.googleusercontent.com/WmWd-wg3LsIjvefwD2VvfqCeFLMmCTZlj8QlGSL373bonYFy_SCZK_qsWBIfRF7vbl3aU0GTGEAErSs=w120-h120-l90-rj",
          "width": 120,
          "height": 120
        }
      ]
    },
    {
      "videoId": "8d7DWVAeiFo",
      "playlistId": "RDAMVM8d7DWVAeiFo",
      "title": "Come Up And Get Me",
      "artists": [{ "name": "Death Grips", "id": "UCvM75EuDkJE-65aH5rkfWmg" }],
      "album": { "name": "No Love Deep Web", "id": "MPREb_FCRKNFJVb1t" },
      "duration": 252,
      "isExplicit": true,
      "thumbnails": [
        {
          "url": "https://lh3.googleusercontent.com/RXF4F3ptT_mOITf70OiHipxZ9MWYwqOXxcVNZzswii928wIbTXloel2ngFTtfckhN79aQ41YN1Y7Orl7=w60-h60-l90-rj",
          "width": 60,
          "height": 60
        },
        {
          "url": "https://lh3.googleusercontent.com/RXF4F3ptT_mOITf70OiHipxZ9MWYwqOXxcVNZzswii928wIbTXloel2ngFTtfckhN79aQ41YN1Y7Orl7=w120-h120-l90-rj",
          "width": 120,
          "height": 120
        }
      ]
    },
    {
      "videoId": "mbZ7HYGs1YE",
      "playlistId": "RDAMVMmbZ7HYGs1YE",
      "title": "Bubbles Buried In This Jungle",
      "artists": [{ "name": "Death Grips", "id": "UCvM75EuDkJE-65aH5rkfWmg" }],
      "album": { "name": "Bottomless Pit", "id": "MPREb_BeLAl3SNrRR" },
      "duration": 159,
      "isExplicit": true,
      "thumbnails": [
        {
          "url": "https://lh3.googleusercontent.com/6d7zQPUhUNMSu1txmKIuzJ-BDCIS_yQYAvYzu__42VI2QeucfoUFinYi-CkU5dOOKUEkTTv_PQcpjhM=w60-h60-l90-rj",
          "width": 60,
          "height": 60
        },
        {
          "url": "https://lh3.googleusercontent.com/6d7zQPUhUNMSu1txmKIuzJ-BDCIS_yQYAvYzu__42VI2QeucfoUFinYi-CkU5dOOKUEkTTv_PQcpjhM=w120-h120-l90-rj",
          "width": 120,
          "height": 120
        }
      ]
    }
  ],
  "artists": [
    {
      "browseId": "UCvM75EuDkJE-65aH5rkfWmg",
      "artist": "Death Grips",
      "shuffleId": "RDAOIa469pUbW1R4pxYoy_l_NQ",
      "radioId": "RDAOIa469pUbW1R4pxYoy_l_NQ",
      "thumbnails": [
        {
          "url": "https://lh3.googleusercontent.com/boo3K8l_2EoOXj18QzaEMlK7ecI8V6WtSvXpCjTbySGH_mL8kxHWl-h2yQAEcpZEfdfr71VJ-eslcqU=w60-h60-p-l90-rj",
          "width": 60,
          "height": 60
        },
        {
          "url": "https://lh3.googleusercontent.com/boo3K8l_2EoOXj18QzaEMlK7ecI8V6WtSvXpCjTbySGH_mL8kxHWl-h2yQAEcpZEfdfr71VJ-eslcqU=w120-h120-p-l90-rj",
          "width": 120,
          "height": 120
        }
      ]
    },
    {
      "browseId": "UCrg2LfxablyfpYkD75hDxMA",
      "artist": "Kero Kero Bonito",
      "shuffleId": "RDAOWXjw4GeAz3E3rzHqNWS52Q",
      "radioId": "RDAOWXjw4GeAz3E3rzHqNWS52Q",
      "thumbnails": [
        {
          "url": "https://lh3.googleusercontent.com/oUJfc9sL7EM1wKEGO07ctB7HEUMLC6cnjXacebnA7cIQYEwpgKS6iaeaJF8hdiDctBgPo0Cinbt_G6Yq=w60-h60-l90-rj",
          "width": 60,
          "height": 60
        },
        {
          "url": "https://lh3.googleusercontent.com/oUJfc9sL7EM1wKEGO07ctB7HEUMLC6cnjXacebnA7cIQYEwpgKS6iaeaJF8hdiDctBgPo0Cinbt_G6Yq=w120-h120-l90-rj",
          "width": 120,
          "height": 120
        }
      ]
    }
  ],
  "albums": [
    {
      "browseId": "MPREb_BL9sWaZWAUE",
      "title": "The Money Store",
      "type": "Album",
      "artists": [{ "name": "Death Grips", "id": "UCvM75EuDkJE-65aH5rkfWmg" }],
      "year": "2012",
      "isExplicit": true,
      "thumbnails": [
        {
          "url": "https://lh3.googleusercontent.com/WmWd-wg3LsIjvefwD2VvfqCeFLMmCTZlj8QlGSL373bonYFy_SCZK_qsWBIfRF7vbl3aU0GTGEAErSs=w60-h60-l90-rj",
          "width": 60,
          "height": 60
        },
        {
          "url": "https://lh3.googleusercontent.com/WmWd-wg3LsIjvefwD2VvfqCeFLMmCTZlj8QlGSL373bonYFy_SCZK_qsWBIfRF7vbl3aU0GTGEAErSs=w120-h120-l90-rj",
          "width": 120,
          "height": 120
        },
        {
          "url": "https://lh3.googleusercontent.com/WmWd-wg3LsIjvefwD2VvfqCeFLMmCTZlj8QlGSL373bonYFy_SCZK_qsWBIfRF7vbl3aU0GTGEAErSs=w226-h226-l90-rj",
          "width": 226,
          "height": 226
        },
        {
          "url": "https://lh3.googleusercontent.com/WmWd-wg3LsIjvefwD2VvfqCeFLMmCTZlj8QlGSL373bonYFy_SCZK_qsWBIfRF7vbl3aU0GTGEAErSs=w544-h544-l90-rj",
          "width": 544,
          "height": 544
        }
      ]
    },
    {
      "browseId": "MPREb_AkwoahACC3J",
      "title": "Flies",
      "type": "Single",
      "artists": [{ "name": "Death Grips", "id": "UCvM75EuDkJE-65aH5rkfWmg" }],
      "year": "2018",
      "isExplicit": true,
      "thumbnails": [
        {
          "url": "https://lh3.googleusercontent.com/2casvjuOmQ3rtzgWjzQX-JMhB-OZu9nsg_vzqBDCviHg8hYoD_O_P38TjbvwfT2AArgn39pKuL4UqMlQ=w60-h60-l90-rj",
          "width": 60,
          "height": 60
        },
        {
          "url": "https://lh3.googleusercontent.com/2casvjuOmQ3rtzgWjzQX-JMhB-OZu9nsg_vzqBDCviHg8hYoD_O_P38TjbvwfT2AArgn39pKuL4UqMlQ=w120-h120-l90-rj",
          "width": 120,
          "height": 120
        },
        {
          "url": "https://lh3.googleusercontent.com/2casvjuOmQ3rtzgWjzQX-JMhB-OZu9nsg_vzqBDCviHg8hYoD_O_P38TjbvwfT2AArgn39pKuL4UqMlQ=w226-h226-l90-rj",
          "width": 226,
          "height": 226
        },
        {
          "url": "https://lh3.googleusercontent.com/2casvjuOmQ3rtzgWjzQX-JMhB-OZu9nsg_vzqBDCviHg8hYoD_O_P38TjbvwfT2AArgn39pKuL4UqMlQ=w544-h544-l90-rj",
          "width": 544,
          "height": 544
        }
      ]
    },
    {
      "browseId": "MPREb_D89UwANNyA9",
      "title": "Black Paint",
      "type": "Single",
      "artists": [{ "name": "Death Grips", "id": "UCvM75EuDkJE-65aH5rkfWmg" }],
      "year": "2018",
      "isExplicit": false,
      "thumbnails": [
        {
          "url": "https://lh3.googleusercontent.com/2casvjuOmQ3rtzgWjzQX-JMhB-OZu9nsg_vzqBDCviHg8hYoD_O_P38TjbvwfT2AArgn39pKuL4UqMlQ=w60-h60-l90-rj",
          "width": 60,
          "height": 60
        },
        {
          "url": "https://lh3.googleusercontent.com/2casvjuOmQ3rtzgWjzQX-JMhB-OZu9nsg_vzqBDCviHg8hYoD_O_P38TjbvwfT2AArgn39pKuL4UqMlQ=w120-h120-l90-rj",
          "width": 120,
          "height": 120
        },
        {
          "url": "https://lh3.googleusercontent.com/2casvjuOmQ3rtzgWjzQX-JMhB-OZu9nsg_vzqBDCviHg8hYoD_O_P38TjbvwfT2AArgn39pKuL4UqMlQ=w226-h226-l90-rj",
          "width": 226,
          "height": 226
        },
        {
          "url": "https://lh3.googleusercontent.com/2casvjuOmQ3rtzgWjzQX-JMhB-OZu9nsg_vzqBDCviHg8hYoD_O_P38TjbvwfT2AArgn39pKuL4UqMlQ=w544-h544-l90-rj",
          "width": 544,
          "height": 544
        }
      ]
    }
  ],
  "playlists": [
    {
      "browseId": "VLPLI7cAsCDsB3F0qHcVGWWAIQAtsvnyhsl0",
      "title": "Death Grips Discography",
      "author": "flubberfan2006",
      "itemCount": "140 songs",
      "thumbnails": [
        {
          "url": "https://yt3.ggpht.com/lyIJJtfJ83ek2QyqnAazs5mh16Deoip1gqVEfMaFd9hCxVzukesU_1KrfbbD3rpneiqpB6quUzs=s192",
          "width": 192,
          "height": 192
        },
        {
          "url": "https://yt3.ggpht.com/lyIJJtfJ83ek2QyqnAazs5mh16Deoip1gqVEfMaFd9hCxVzukesU_1KrfbbD3rpneiqpB6quUzs=s576",
          "width": 576,
          "height": 576
        },
        {
          "url": "https://yt3.ggpht.com/lyIJJtfJ83ek2QyqnAazs5mh16Deoip1gqVEfMaFd9hCxVzukesU_1KrfbbD3rpneiqpB6quUzs=s1200",
          "width": 1200,
          "height": 1200
        }
      ]
    },
    {
      "browseId": "VLPLoZWz2YfH95_v8gvePJ95NL88VpGPJUol",
      "title": "great death grips mashups",
      "author": "nadia",
      "itemCount": "169 songs",
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/cdmhyEqCINM/sddefault.jpg?sqp=-oaymwEWCJADEOEBIAQqCghqEJQEGHgg6AJIWg\u0026rs=AMzJL3laCHXn9eTlVqGdixcZU72hdYUMLw",
          "width": 400,
          "height": 400
        },
        {
          "url": "https://i.ytimg.com/vi/cdmhyEqCINM/hq720.jpg?sqp=-oaymwEXCKAGEMIDIAQqCwjVARCqCBh4INgESFo\u0026rs=AMzJL3kUtOYItC7iZpJVPSbA5b25euRB2g",
          "width": 800,
          "height": 800
        },
        {
          "url": "https://i.ytimg.com/vi/cdmhyEqCINM/hq720.jpg?sqp=-oaymwEXCNUGEOADIAQqCwjVARCqCBh4INgESFo\u0026rs=AMzJL3moAynTolvYbOPi4A5OAs4RLVp1jg",
          "width": 853,
          "height": 853
        }
      ]
    },
    {
      "browseId": "VLPLH5jk9dL-v_w_TakRXVjC5LznsptXYD4U",
      "title": "Death Grips Absolute Discography",
      "author": "Yangjo",
      "itemCount": "117 songs",
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/sFmZ6WDkuj0/hqdefault.jpg?sqp=-oaymwEWCMACELQBIAQqCghQEJADGFogjgJIWg\u0026rs=AMzJL3k8ELkIvqikzlQ-Y8Oky2kfXEAmkg",
          "width": 320,
          "height": 320
        }
      ]
    }
  ],
  "videos": [
    {
      "videoId": "iPtPo8Sa3NE",
      "playlistId": "RDAMVMiPtPo8Sa3NE",
      "title": "Death Grips - Exmilitary [Full Mixtape]",
      "artists": [{ "name": "An0rhu", "id": "UCqktZCXcZQf1irrzkertseQ" }],
      "views": "5.5M views",
      "duration": 2911,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/iPtPo8Sa3NE/sddefault.jpg?sqp=-oaymwEWCJADEOEBIAQqCghqEJQEGHgg6AJIWg\u0026rs=AMzJL3kKmQiuO_3TTh_9Qu1Ac-HiajpISQ",
          "width": 400,
          "height": 400
        }
      ]
    },
    {
      "videoId": "vz2KSKxRRlo",
      "playlistId": "RDAMVMvz2KSKxRRlo",
      "title": "Punk Weight",
      "artists": [{ "name": "Death Grips", "id": "UCvM75EuDkJE-65aH5rkfWmg" }],
      "views": "1M views",
      "duration": 205,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/vz2KSKxRRlo/sddefault.jpg?sqp=-oaymwEWCJADEOEBIAQqCghqEJQEGHgg6AJIWg\u0026rs=AMzJL3kx89mHX7Fcyrw3QI0hbMQdbbosbw",
          "width": 400,
          "height": 400
        }
      ]
    },
    {
      "videoId": "f459zKYpbJ0",
      "playlistId": "RDAMVMf459zKYpbJ0",
      "title": "Death Grips - Fashion Week [Full Album]",
      "artists": [{ "name": "KSS TV", "id": "UCadS9YA-kcPlrN52vy4JLCA" }],
      "views": "47K views",
      "duration": 2881,
      "thumbnails": [
        {
          "url": "https://i.ytimg.com/vi/f459zKYpbJ0/sddefault.jpg?sqp=-oaymwEWCJADEOEBIAQqCghqEJQEGHgg6AJIWg\u0026rs=AMzJL3lQ-c8WFxDXDH0tESXltajsUb9X-A",
          "width": 400,
          "height": 400
        }
      ]
    }
  ]
}
```
</details>

### Get watch playlist

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/raitonoberu/ytmusic"
)

func main() {
    result, err := ytmusic.GetWatchPlaylist("ljUtuoFt-8c")
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
[
  {
    "videoId": "ljUtuoFt-8c",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Smells Like Teen Spirit",
    "artists": [{ "name": "Nirvana", "id": "UCrPe3hLA51968GwxHSZ1llw" }],
    "album": { "name": "Nevermind", "id": "MPREb_jPOYfjGgApr" },
    "duration": 301,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/eyKiPBSqEu556sYTd_IyZhfxun5e_hatZ9tAyu8bnmVRgtbM3aW-SXUvhVX-d7s1oU0Yf3a38JOuYMZK5w=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/eyKiPBSqEu556sYTd_IyZhfxun5e_hatZ9tAyu8bnmVRgtbM3aW-SXUvhVX-d7s1oU0Yf3a38JOuYMZK5w=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/eyKiPBSqEu556sYTd_IyZhfxun5e_hatZ9tAyu8bnmVRgtbM3aW-SXUvhVX-d7s1oU0Yf3a38JOuYMZK5w=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/eyKiPBSqEu556sYTd_IyZhfxun5e_hatZ9tAyu8bnmVRgtbM3aW-SXUvhVX-d7s1oU0Yf3a38JOuYMZK5w=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/eyKiPBSqEu556sYTd_IyZhfxun5e_hatZ9tAyu8bnmVRgtbM3aW-SXUvhVX-d7s1oU0Yf3a38JOuYMZK5w=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/eyKiPBSqEu556sYTd_IyZhfxun5e_hatZ9tAyu8bnmVRgtbM3aW-SXUvhVX-d7s1oU0Yf3a38JOuYMZK5w=w544-h544-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "kaOOfci2YC8",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Nothing Else Matters",
    "artists": [{ "name": "Metallica", "id": "UCGexNm_Kw4rdQjLxmpb2EKw" }],
    "album": { "name": "Metallica", "id": "MPREb_AmmVchA41UA" },
    "duration": 389,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/P6UQgFQD0MduEsjvh0D1CahKTqebTtObTBUkYp_7F9QcDbtgqFsqkHQwTOtJX6XQv5vL8fxEiBgpjj0K_Q=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/P6UQgFQD0MduEsjvh0D1CahKTqebTtObTBUkYp_7F9QcDbtgqFsqkHQwTOtJX6XQv5vL8fxEiBgpjj0K_Q=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/P6UQgFQD0MduEsjvh0D1CahKTqebTtObTBUkYp_7F9QcDbtgqFsqkHQwTOtJX6XQv5vL8fxEiBgpjj0K_Q=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/P6UQgFQD0MduEsjvh0D1CahKTqebTtObTBUkYp_7F9QcDbtgqFsqkHQwTOtJX6XQv5vL8fxEiBgpjj0K_Q=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/P6UQgFQD0MduEsjvh0D1CahKTqebTtObTBUkYp_7F9QcDbtgqFsqkHQwTOtJX6XQv5vL8fxEiBgpjj0K_Q=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/P6UQgFQD0MduEsjvh0D1CahKTqebTtObTBUkYp_7F9QcDbtgqFsqkHQwTOtJX6XQv5vL8fxEiBgpjj0K_Q=w544-h544-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "A__cH65WRvE",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Californication",
    "artists": [
      { "name": "Red Hot Chili Peppers", "id": "UCrSorX845CEWXzU4Z7BojjA" }
    ],
    "album": {
      "name": "Californication (Deluxe Edition)",
      "id": "MPREb_9zDg9WFf7KR"
    },
    "duration": 329,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/vEzFuxLILzCtIKEdZWOwSoIfXN3462-JcZ7AYXgV7i9n06NzazzUQOloJ7ghh2Weswt7OcJQH79BI9kePA=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/vEzFuxLILzCtIKEdZWOwSoIfXN3462-JcZ7AYXgV7i9n06NzazzUQOloJ7ghh2Weswt7OcJQH79BI9kePA=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/vEzFuxLILzCtIKEdZWOwSoIfXN3462-JcZ7AYXgV7i9n06NzazzUQOloJ7ghh2Weswt7OcJQH79BI9kePA=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/vEzFuxLILzCtIKEdZWOwSoIfXN3462-JcZ7AYXgV7i9n06NzazzUQOloJ7ghh2Weswt7OcJQH79BI9kePA=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/vEzFuxLILzCtIKEdZWOwSoIfXN3462-JcZ7AYXgV7i9n06NzazzUQOloJ7ghh2Weswt7OcJQH79BI9kePA=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/vEzFuxLILzCtIKEdZWOwSoIfXN3462-JcZ7AYXgV7i9n06NzazzUQOloJ7ghh2Weswt7OcJQH79BI9kePA=w512-h512-l90-rj",
        "width": 512,
        "height": 512
      }
    ]
  },
  {
    "videoId": "9vWNauaZAgg",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Back In Black",
    "artists": [{ "name": "AC/DC", "id": "UCVm4YdI3hobkwsHTTOMVJKg" }],
    "album": { "name": "Back In Black", "id": "MPREb_epDHrawzNu7" },
    "duration": 256,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/etTz20YiB4ccbsUO2yLrCY9wSS9GybYF5qJh-j5tu8MLTqP2GjgROiBt_4JUC5rjnnd_RiuWa3ndUAeO=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/etTz20YiB4ccbsUO2yLrCY9wSS9GybYF5qJh-j5tu8MLTqP2GjgROiBt_4JUC5rjnnd_RiuWa3ndUAeO=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/etTz20YiB4ccbsUO2yLrCY9wSS9GybYF5qJh-j5tu8MLTqP2GjgROiBt_4JUC5rjnnd_RiuWa3ndUAeO=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/etTz20YiB4ccbsUO2yLrCY9wSS9GybYF5qJh-j5tu8MLTqP2GjgROiBt_4JUC5rjnnd_RiuWa3ndUAeO=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/etTz20YiB4ccbsUO2yLrCY9wSS9GybYF5qJh-j5tu8MLTqP2GjgROiBt_4JUC5rjnnd_RiuWa3ndUAeO=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/etTz20YiB4ccbsUO2yLrCY9wSS9GybYF5qJh-j5tu8MLTqP2GjgROiBt_4JUC5rjnnd_RiuWa3ndUAeO=w544-h544-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "1WTATreAg08",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Bitter Sweet Symphony (Extended Version)",
    "artists": [{ "name": "The Verve", "id": "UCt2KxZpY5D__kapeQ8cauQw" }],
    "album": {
      "name": "Urban Hymns (Super Deluxe / Remastered 2016)",
      "id": "MPREb_KOt6PrA0bWj"
    },
    "duration": 469,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/o8FwEeFfmuLXHpKBUHcldm9joALAVwf35TT3xmVS8-X5B2bKE0o0Xd_YUqlVOXNQQLT6ie4J9X8Lck2Muw=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/o8FwEeFfmuLXHpKBUHcldm9joALAVwf35TT3xmVS8-X5B2bKE0o0Xd_YUqlVOXNQQLT6ie4J9X8Lck2Muw=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/o8FwEeFfmuLXHpKBUHcldm9joALAVwf35TT3xmVS8-X5B2bKE0o0Xd_YUqlVOXNQQLT6ie4J9X8Lck2Muw=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/o8FwEeFfmuLXHpKBUHcldm9joALAVwf35TT3xmVS8-X5B2bKE0o0Xd_YUqlVOXNQQLT6ie4J9X8Lck2Muw=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/o8FwEeFfmuLXHpKBUHcldm9joALAVwf35TT3xmVS8-X5B2bKE0o0Xd_YUqlVOXNQQLT6ie4J9X8Lck2Muw=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/o8FwEeFfmuLXHpKBUHcldm9joALAVwf35TT3xmVS8-X5B2bKE0o0Xd_YUqlVOXNQQLT6ie4J9X8Lck2Muw=w544-h544-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "mmIXwYkDs9k",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "The Unforgiven",
    "artists": [{ "name": "Metallica", "id": "UCGexNm_Kw4rdQjLxmpb2EKw" }],
    "album": { "name": "Metallica", "id": "MPREb_AmmVchA41UA" },
    "duration": 386,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/P6UQgFQD0MduEsjvh0D1CahKTqebTtObTBUkYp_7F9QcDbtgqFsqkHQwTOtJX6XQv5vL8fxEiBgpjj0K_Q=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/P6UQgFQD0MduEsjvh0D1CahKTqebTtObTBUkYp_7F9QcDbtgqFsqkHQwTOtJX6XQv5vL8fxEiBgpjj0K_Q=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/P6UQgFQD0MduEsjvh0D1CahKTqebTtObTBUkYp_7F9QcDbtgqFsqkHQwTOtJX6XQv5vL8fxEiBgpjj0K_Q=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/P6UQgFQD0MduEsjvh0D1CahKTqebTtObTBUkYp_7F9QcDbtgqFsqkHQwTOtJX6XQv5vL8fxEiBgpjj0K_Q=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/P6UQgFQD0MduEsjvh0D1CahKTqebTtObTBUkYp_7F9QcDbtgqFsqkHQwTOtJX6XQv5vL8fxEiBgpjj0K_Q=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/P6UQgFQD0MduEsjvh0D1CahKTqebTtObTBUkYp_7F9QcDbtgqFsqkHQwTOtJX6XQv5vL8fxEiBgpjj0K_Q=w544-h544-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "5qZQEq_C3vc",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Numb",
    "artists": [{ "name": "Linkin Park", "id": "UCxgN32UVVztKAQd2HkXzBtw" }],
    "album": { "name": "Meteora (Bonus Edition)", "id": "MPREb_ebWSmYnO2aa" },
    "duration": 186,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/Rqh2HDAD-Fj1-oHYqVklPGklMHIV3r6SMPMZmFE3FlLiWeBBnP3SCESfgeNdzQ-DmpCIvOGE4J0hzVxa=w60-h60-s-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/Rqh2HDAD-Fj1-oHYqVklPGklMHIV3r6SMPMZmFE3FlLiWeBBnP3SCESfgeNdzQ-DmpCIvOGE4J0hzVxa=w120-h120-s-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/Rqh2HDAD-Fj1-oHYqVklPGklMHIV3r6SMPMZmFE3FlLiWeBBnP3SCESfgeNdzQ-DmpCIvOGE4J0hzVxa=w180-h180-s-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/Rqh2HDAD-Fj1-oHYqVklPGklMHIV3r6SMPMZmFE3FlLiWeBBnP3SCESfgeNdzQ-DmpCIvOGE4J0hzVxa=w226-h226-s-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/Rqh2HDAD-Fj1-oHYqVklPGklMHIV3r6SMPMZmFE3FlLiWeBBnP3SCESfgeNdzQ-DmpCIvOGE4J0hzVxa=w302-h302-s-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/Rqh2HDAD-Fj1-oHYqVklPGklMHIV3r6SMPMZmFE3FlLiWeBBnP3SCESfgeNdzQ-DmpCIvOGE4J0hzVxa=w544-h544-s-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "y6lfK3bH4z8",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "November Rain",
    "artists": [{ "name": "Guns N' Roses", "id": "UCSLbbBoUqpin6BE34whSOvA" }],
    "album": { "name": "Use Your Illusion I", "id": "MPREb_fj1B4Tl3r6C" },
    "duration": 537,
    "isExplicit": true,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/pcrIVPiugsSpXJJPCaPNdp2vpHPo5UsmsoVjqlT7L1p7Lki663u9tONk-S3oqK9GvAsoj2m0AaNgNJj9Zg=w60-h60-s-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/pcrIVPiugsSpXJJPCaPNdp2vpHPo5UsmsoVjqlT7L1p7Lki663u9tONk-S3oqK9GvAsoj2m0AaNgNJj9Zg=w120-h120-s-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/pcrIVPiugsSpXJJPCaPNdp2vpHPo5UsmsoVjqlT7L1p7Lki663u9tONk-S3oqK9GvAsoj2m0AaNgNJj9Zg=w180-h180-s-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/pcrIVPiugsSpXJJPCaPNdp2vpHPo5UsmsoVjqlT7L1p7Lki663u9tONk-S3oqK9GvAsoj2m0AaNgNJj9Zg=w226-h226-s-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/pcrIVPiugsSpXJJPCaPNdp2vpHPo5UsmsoVjqlT7L1p7Lki663u9tONk-S3oqK9GvAsoj2m0AaNgNJj9Zg=w302-h302-s-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/pcrIVPiugsSpXJJPCaPNdp2vpHPo5UsmsoVjqlT7L1p7Lki663u9tONk-S3oqK9GvAsoj2m0AaNgNJj9Zg=w544-h544-s-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "8901V1M5lDk",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Otherside",
    "artists": [
      { "name": "Red Hot Chili Peppers", "id": "UCrSorX845CEWXzU4Z7BojjA" }
    ],
    "album": {
      "name": "Californication (Deluxe Edition)",
      "id": "MPREb_9zDg9WFf7KR"
    },
    "duration": 255,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/vEzFuxLILzCtIKEdZWOwSoIfXN3462-JcZ7AYXgV7i9n06NzazzUQOloJ7ghh2Weswt7OcJQH79BI9kePA=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/vEzFuxLILzCtIKEdZWOwSoIfXN3462-JcZ7AYXgV7i9n06NzazzUQOloJ7ghh2Weswt7OcJQH79BI9kePA=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/vEzFuxLILzCtIKEdZWOwSoIfXN3462-JcZ7AYXgV7i9n06NzazzUQOloJ7ghh2Weswt7OcJQH79BI9kePA=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/vEzFuxLILzCtIKEdZWOwSoIfXN3462-JcZ7AYXgV7i9n06NzazzUQOloJ7ghh2Weswt7OcJQH79BI9kePA=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/vEzFuxLILzCtIKEdZWOwSoIfXN3462-JcZ7AYXgV7i9n06NzazzUQOloJ7ghh2Weswt7OcJQH79BI9kePA=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/vEzFuxLILzCtIKEdZWOwSoIfXN3462-JcZ7AYXgV7i9n06NzazzUQOloJ7ghh2Weswt7OcJQH79BI9kePA=w512-h512-l90-rj",
        "width": 512,
        "height": 512
      }
    ]
  },
  {
    "videoId": "oMfMUfgjiLg",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Sweet Child O' Mine",
    "artists": [{ "name": "Guns N' Roses", "id": "UCSLbbBoUqpin6BE34whSOvA" }],
    "album": { "name": "Appetite For Destruction", "id": "MPREb_fN7uwe7VDV6" },
    "duration": 357,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/Z2zgZT7ICWRL7ACDERnWaGA3nBmLSK3fqOXDg-EQuhSnwvNV_i5pEQ6-z_m_v4S44HuNKyFjwtmSghqJ=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/Z2zgZT7ICWRL7ACDERnWaGA3nBmLSK3fqOXDg-EQuhSnwvNV_i5pEQ6-z_m_v4S44HuNKyFjwtmSghqJ=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/Z2zgZT7ICWRL7ACDERnWaGA3nBmLSK3fqOXDg-EQuhSnwvNV_i5pEQ6-z_m_v4S44HuNKyFjwtmSghqJ=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/Z2zgZT7ICWRL7ACDERnWaGA3nBmLSK3fqOXDg-EQuhSnwvNV_i5pEQ6-z_m_v4S44HuNKyFjwtmSghqJ=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/Z2zgZT7ICWRL7ACDERnWaGA3nBmLSK3fqOXDg-EQuhSnwvNV_i5pEQ6-z_m_v4S44HuNKyFjwtmSghqJ=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/Z2zgZT7ICWRL7ACDERnWaGA3nBmLSK3fqOXDg-EQuhSnwvNV_i5pEQ6-z_m_v4S44HuNKyFjwtmSghqJ=w544-h544-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "YxRxd8aNd6I",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Enter Sandman",
    "artists": [{ "name": "Metallica", "id": "UCGexNm_Kw4rdQjLxmpb2EKw" }],
    "album": { "name": "Metallica", "id": "MPREb_AmmVchA41UA" },
    "duration": 333,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/P6UQgFQD0MduEsjvh0D1CahKTqebTtObTBUkYp_7F9QcDbtgqFsqkHQwTOtJX6XQv5vL8fxEiBgpjj0K_Q=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/P6UQgFQD0MduEsjvh0D1CahKTqebTtObTBUkYp_7F9QcDbtgqFsqkHQwTOtJX6XQv5vL8fxEiBgpjj0K_Q=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/P6UQgFQD0MduEsjvh0D1CahKTqebTtObTBUkYp_7F9QcDbtgqFsqkHQwTOtJX6XQv5vL8fxEiBgpjj0K_Q=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/P6UQgFQD0MduEsjvh0D1CahKTqebTtObTBUkYp_7F9QcDbtgqFsqkHQwTOtJX6XQv5vL8fxEiBgpjj0K_Q=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/P6UQgFQD0MduEsjvh0D1CahKTqebTtObTBUkYp_7F9QcDbtgqFsqkHQwTOtJX6XQv5vL8fxEiBgpjj0K_Q=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/P6UQgFQD0MduEsjvh0D1CahKTqebTtObTBUkYp_7F9QcDbtgqFsqkHQwTOtJX6XQv5vL8fxEiBgpjj0K_Q=w544-h544-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "xBQO0ac41Xo",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Are You Gonna Be My Girl",
    "artists": [{ "name": "Jet", "id": "UCDY3bjtG-e456ADVpfDZ4Tw" }],
    "album": { "name": "Are You Gonna Be My Girl", "id": "MPREb_vHQ0sUxM14M" },
    "duration": 214,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/TZQtONsHrIjjLd_gOyhn1kYkSrY0D8j479vvPGUxcaMh7__6mNzGSxXYW0mj9FWa-y9FPQsoxxI-Xr4=w60-h60-s-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/TZQtONsHrIjjLd_gOyhn1kYkSrY0D8j479vvPGUxcaMh7__6mNzGSxXYW0mj9FWa-y9FPQsoxxI-Xr4=w120-h120-s-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/TZQtONsHrIjjLd_gOyhn1kYkSrY0D8j479vvPGUxcaMh7__6mNzGSxXYW0mj9FWa-y9FPQsoxxI-Xr4=w180-h180-s-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/TZQtONsHrIjjLd_gOyhn1kYkSrY0D8j479vvPGUxcaMh7__6mNzGSxXYW0mj9FWa-y9FPQsoxxI-Xr4=w226-h226-s-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/TZQtONsHrIjjLd_gOyhn1kYkSrY0D8j479vvPGUxcaMh7__6mNzGSxXYW0mj9FWa-y9FPQsoxxI-Xr4=w302-h302-s-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/TZQtONsHrIjjLd_gOyhn1kYkSrY0D8j479vvPGUxcaMh7__6mNzGSxXYW0mj9FWa-y9FPQsoxxI-Xr4=w544-h544-s-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "J0xe5DcnYSA",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Holiday / Boulevard of Broken Dreams",
    "artists": [{ "name": "Green Day", "id": "UC4JNeITH4P7G51C1hJoG6vQ" }],
    "album": { "name": "American Idiot", "id": "MPREb_jnYA50rlu4H" },
    "duration": 494,
    "isExplicit": true,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/UXFGtBhmxZd0cWbwgJCsaUvLH59uZe_T_9E1plQPi_zHBuPAHTuhzz_h8I-uIYaXJMKvVfRB7gcRJFI=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/UXFGtBhmxZd0cWbwgJCsaUvLH59uZe_T_9E1plQPi_zHBuPAHTuhzz_h8I-uIYaXJMKvVfRB7gcRJFI=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/UXFGtBhmxZd0cWbwgJCsaUvLH59uZe_T_9E1plQPi_zHBuPAHTuhzz_h8I-uIYaXJMKvVfRB7gcRJFI=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/UXFGtBhmxZd0cWbwgJCsaUvLH59uZe_T_9E1plQPi_zHBuPAHTuhzz_h8I-uIYaXJMKvVfRB7gcRJFI=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/UXFGtBhmxZd0cWbwgJCsaUvLH59uZe_T_9E1plQPi_zHBuPAHTuhzz_h8I-uIYaXJMKvVfRB7gcRJFI=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/UXFGtBhmxZd0cWbwgJCsaUvLH59uZe_T_9E1plQPi_zHBuPAHTuhzz_h8I-uIYaXJMKvVfRB7gcRJFI=w544-h544-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "-DPomaw9Sl0",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Don't Cry (Alternate Lyrics)",
    "artists": [{ "name": "Guns N' Roses", "id": "UCSLbbBoUqpin6BE34whSOvA" }],
    "album": { "name": "Use Your Illusion II", "id": "MPREb_BWthKEBKL9Y" },
    "duration": 285,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/Q7XFwlxOBgvsK4ZPyK855aeG5piiewmvIWhN_6aoAfe4hn0uqKQ2vzEOLr-CRLuohr0RwuIo7p4rCdE=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/Q7XFwlxOBgvsK4ZPyK855aeG5piiewmvIWhN_6aoAfe4hn0uqKQ2vzEOLr-CRLuohr0RwuIo7p4rCdE=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/Q7XFwlxOBgvsK4ZPyK855aeG5piiewmvIWhN_6aoAfe4hn0uqKQ2vzEOLr-CRLuohr0RwuIo7p4rCdE=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/Q7XFwlxOBgvsK4ZPyK855aeG5piiewmvIWhN_6aoAfe4hn0uqKQ2vzEOLr-CRLuohr0RwuIo7p4rCdE=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/Q7XFwlxOBgvsK4ZPyK855aeG5piiewmvIWhN_6aoAfe4hn0uqKQ2vzEOLr-CRLuohr0RwuIo7p4rCdE=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/Q7XFwlxOBgvsK4ZPyK855aeG5piiewmvIWhN_6aoAfe4hn0uqKQ2vzEOLr-CRLuohr0RwuIo7p4rCdE=w544-h544-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "ZrOKjDZOtkA",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Wonderwall",
    "artists": [{ "name": "Oasis", "id": "UCmMUZbaYdNH0bEd1PAlAqsA" }],
    "album": {
      "name": "(What's The Story) Morning Glory? (Remastered)",
      "id": "MPREb_9nqEki4ZDpp"
    },
    "duration": 259,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/P99EPH7dpg6_t25NsyXiYhPCBQE5hw2vs5IWEf5mjMSpD042hw_W0comDEZzRMeMtLdsCQUPHs8AOgk=w60-h60-s-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/P99EPH7dpg6_t25NsyXiYhPCBQE5hw2vs5IWEf5mjMSpD042hw_W0comDEZzRMeMtLdsCQUPHs8AOgk=w120-h120-s-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/P99EPH7dpg6_t25NsyXiYhPCBQE5hw2vs5IWEf5mjMSpD042hw_W0comDEZzRMeMtLdsCQUPHs8AOgk=w180-h180-s-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/P99EPH7dpg6_t25NsyXiYhPCBQE5hw2vs5IWEf5mjMSpD042hw_W0comDEZzRMeMtLdsCQUPHs8AOgk=w226-h226-s-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/P99EPH7dpg6_t25NsyXiYhPCBQE5hw2vs5IWEf5mjMSpD042hw_W0comDEZzRMeMtLdsCQUPHs8AOgk=w302-h302-s-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/P99EPH7dpg6_t25NsyXiYhPCBQE5hw2vs5IWEf5mjMSpD042hw_W0comDEZzRMeMtLdsCQUPHs8AOgk=w544-h544-s-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "lhg9bYNLvOg",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Thunderstruck",
    "artists": [{ "name": "AC/DC", "id": "UCVm4YdI3hobkwsHTTOMVJKg" }],
    "album": { "name": "The Razors Edge", "id": "MPREb_Dlj3PZaTijk" },
    "duration": 293,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/DV_ebk0MH4HMLfn2CxeH6Thjf9OSs1Q6FhZulHE6KIyoPWYT6rh2FOZIERUFLRKg7yxpPqwIBE31D5ETkg=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/DV_ebk0MH4HMLfn2CxeH6Thjf9OSs1Q6FhZulHE6KIyoPWYT6rh2FOZIERUFLRKg7yxpPqwIBE31D5ETkg=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/DV_ebk0MH4HMLfn2CxeH6Thjf9OSs1Q6FhZulHE6KIyoPWYT6rh2FOZIERUFLRKg7yxpPqwIBE31D5ETkg=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/DV_ebk0MH4HMLfn2CxeH6Thjf9OSs1Q6FhZulHE6KIyoPWYT6rh2FOZIERUFLRKg7yxpPqwIBE31D5ETkg=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/DV_ebk0MH4HMLfn2CxeH6Thjf9OSs1Q6FhZulHE6KIyoPWYT6rh2FOZIERUFLRKg7yxpPqwIBE31D5ETkg=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/DV_ebk0MH4HMLfn2CxeH6Thjf9OSs1Q6FhZulHE6KIyoPWYT6rh2FOZIERUFLRKg7yxpPqwIBE31D5ETkg=w544-h544-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "BMMGwtklEeE",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "The Pretender",
    "artists": [{ "name": "Foo Fighters", "id": "UCwDw2KaS8jRUxAPScL2U8og" }],
    "album": {
      "name": "Echoes, Silence, Patience \u0026 Grace",
      "id": "MPREb_LxuNwq2DCKw"
    },
    "duration": 270,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/0t56z7BHoL-y707OisUF5JfM6fJFKdhRUW9XHoW120v7PdeNGrXRh68djTKXaRJy1oc3fFfLcWBO8ndU_g=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/0t56z7BHoL-y707OisUF5JfM6fJFKdhRUW9XHoW120v7PdeNGrXRh68djTKXaRJy1oc3fFfLcWBO8ndU_g=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/0t56z7BHoL-y707OisUF5JfM6fJFKdhRUW9XHoW120v7PdeNGrXRh68djTKXaRJy1oc3fFfLcWBO8ndU_g=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/0t56z7BHoL-y707OisUF5JfM6fJFKdhRUW9XHoW120v7PdeNGrXRh68djTKXaRJy1oc3fFfLcWBO8ndU_g=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/0t56z7BHoL-y707OisUF5JfM6fJFKdhRUW9XHoW120v7PdeNGrXRh68djTKXaRJy1oc3fFfLcWBO8ndU_g=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/0t56z7BHoL-y707OisUF5JfM6fJFKdhRUW9XHoW120v7PdeNGrXRh68djTKXaRJy1oc3fFfLcWBO8ndU_g=w544-h544-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "tyBiHyazMOc",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "One",
    "artists": [{ "name": "Metallica", "id": "UCGexNm_Kw4rdQjLxmpb2EKw" }],
    "album": { "name": "…And Justice for All", "id": "MPREb_cJGNPCluDG0" },
    "duration": 447,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/5wGgPxjsT2aEKpr29-Jr5EbQThiT5pRdtcbotWqoB3TqFrs7G8RWB0P7HBJycuazFL2XTb9MR-9LIjJj=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/5wGgPxjsT2aEKpr29-Jr5EbQThiT5pRdtcbotWqoB3TqFrs7G8RWB0P7HBJycuazFL2XTb9MR-9LIjJj=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/5wGgPxjsT2aEKpr29-Jr5EbQThiT5pRdtcbotWqoB3TqFrs7G8RWB0P7HBJycuazFL2XTb9MR-9LIjJj=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/5wGgPxjsT2aEKpr29-Jr5EbQThiT5pRdtcbotWqoB3TqFrs7G8RWB0P7HBJycuazFL2XTb9MR-9LIjJj=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/5wGgPxjsT2aEKpr29-Jr5EbQThiT5pRdtcbotWqoB3TqFrs7G8RWB0P7HBJycuazFL2XTb9MR-9LIjJj=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/5wGgPxjsT2aEKpr29-Jr5EbQThiT5pRdtcbotWqoB3TqFrs7G8RWB0P7HBJycuazFL2XTb9MR-9LIjJj=w544-h544-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "j2R6uT7uLFg",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Last Resort",
    "artists": [{ "name": "Papa Roach", "id": "UCRiFekze9Yx5zDDjAVrYvbg" }],
    "album": { "name": "Infest", "id": "MPREb_cv3dgfHNJ1j" },
    "duration": 200,
    "isExplicit": true,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/UOqZS25iQnlpJCJLtweobidchyC_hG3x-abJj0l-ATFzK4-esEaK8WHGKaj4saMFiPE6ndzu98bpZoE=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/UOqZS25iQnlpJCJLtweobidchyC_hG3x-abJj0l-ATFzK4-esEaK8WHGKaj4saMFiPE6ndzu98bpZoE=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/UOqZS25iQnlpJCJLtweobidchyC_hG3x-abJj0l-ATFzK4-esEaK8WHGKaj4saMFiPE6ndzu98bpZoE=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/UOqZS25iQnlpJCJLtweobidchyC_hG3x-abJj0l-ATFzK4-esEaK8WHGKaj4saMFiPE6ndzu98bpZoE=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/UOqZS25iQnlpJCJLtweobidchyC_hG3x-abJj0l-ATFzK4-esEaK8WHGKaj4saMFiPE6ndzu98bpZoE=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/UOqZS25iQnlpJCJLtweobidchyC_hG3x-abJj0l-ATFzK4-esEaK8WHGKaj4saMFiPE6ndzu98bpZoE=w544-h544-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "T0ZmErXkbxE",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Paradise City",
    "artists": [{ "name": "Guns N' Roses", "id": "UCSLbbBoUqpin6BE34whSOvA" }],
    "album": { "name": "Appetite For Destruction", "id": "MPREb_fN7uwe7VDV6" },
    "duration": 406,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/Z2zgZT7ICWRL7ACDERnWaGA3nBmLSK3fqOXDg-EQuhSnwvNV_i5pEQ6-z_m_v4S44HuNKyFjwtmSghqJ=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/Z2zgZT7ICWRL7ACDERnWaGA3nBmLSK3fqOXDg-EQuhSnwvNV_i5pEQ6-z_m_v4S44HuNKyFjwtmSghqJ=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/Z2zgZT7ICWRL7ACDERnWaGA3nBmLSK3fqOXDg-EQuhSnwvNV_i5pEQ6-z_m_v4S44HuNKyFjwtmSghqJ=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/Z2zgZT7ICWRL7ACDERnWaGA3nBmLSK3fqOXDg-EQuhSnwvNV_i5pEQ6-z_m_v4S44HuNKyFjwtmSghqJ=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/Z2zgZT7ICWRL7ACDERnWaGA3nBmLSK3fqOXDg-EQuhSnwvNV_i5pEQ6-z_m_v4S44HuNKyFjwtmSghqJ=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/Z2zgZT7ICWRL7ACDERnWaGA3nBmLSK3fqOXDg-EQuhSnwvNV_i5pEQ6-z_m_v4S44HuNKyFjwtmSghqJ=w544-h544-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "mUEsqQpact0",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Toxicity",
    "artists": [
      { "name": "System Of A Down", "id": "UCDJftX2zx_UT_QSnBGIF96w" }
    ],
    "album": { "name": "Toxicity", "id": "MPREb_3n9jhahGRLv" },
    "duration": 219,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/W7YWcMqTUODZf_S_10reZDmYPIjTIYOdljFidYubSeKAuE_7MX6uP54ufi565Fz0AXmY7fMDby2gv7UT=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/W7YWcMqTUODZf_S_10reZDmYPIjTIYOdljFidYubSeKAuE_7MX6uP54ufi565Fz0AXmY7fMDby2gv7UT=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/W7YWcMqTUODZf_S_10reZDmYPIjTIYOdljFidYubSeKAuE_7MX6uP54ufi565Fz0AXmY7fMDby2gv7UT=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/W7YWcMqTUODZf_S_10reZDmYPIjTIYOdljFidYubSeKAuE_7MX6uP54ufi565Fz0AXmY7fMDby2gv7UT=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/W7YWcMqTUODZf_S_10reZDmYPIjTIYOdljFidYubSeKAuE_7MX6uP54ufi565Fz0AXmY7fMDby2gv7UT=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/W7YWcMqTUODZf_S_10reZDmYPIjTIYOdljFidYubSeKAuE_7MX6uP54ufi565Fz0AXmY7fMDby2gv7UT=w544-h544-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "bY3vXr7fm8k",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "It's My Life",
    "artists": [{ "name": "Bon Jovi", "id": "UCCL-yoaPLR-7bxZM1xlcQcQ" }],
    "album": { "name": "Crush", "id": "MPREb_AYqyYRAQjJc" },
    "duration": 224,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/yoETujnSXc5KVgp6RvFQLkoaUknI0ygPW7o9T8Z9MedeBfM9q6nDY0t2I4LPn90MjQbXeHubqOcGKy4=w60-h60-s-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/yoETujnSXc5KVgp6RvFQLkoaUknI0ygPW7o9T8Z9MedeBfM9q6nDY0t2I4LPn90MjQbXeHubqOcGKy4=w120-h120-s-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/yoETujnSXc5KVgp6RvFQLkoaUknI0ygPW7o9T8Z9MedeBfM9q6nDY0t2I4LPn90MjQbXeHubqOcGKy4=w180-h180-s-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/yoETujnSXc5KVgp6RvFQLkoaUknI0ygPW7o9T8Z9MedeBfM9q6nDY0t2I4LPn90MjQbXeHubqOcGKy4=w226-h226-s-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/yoETujnSXc5KVgp6RvFQLkoaUknI0ygPW7o9T8Z9MedeBfM9q6nDY0t2I4LPn90MjQbXeHubqOcGKy4=w302-h302-s-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/yoETujnSXc5KVgp6RvFQLkoaUknI0ygPW7o9T8Z9MedeBfM9q6nDY0t2I4LPn90MjQbXeHubqOcGKy4=w544-h544-s-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "ikFFVfObwss",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Highway to Hell",
    "artists": [{ "name": "AC/DC", "id": "UCVm4YdI3hobkwsHTTOMVJKg" }],
    "album": { "name": "Highway to Hell", "id": "MPREb_D5nYt9190Tm" },
    "duration": 209,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/PUFomqBydVJyjWf41HjnsoBmEgA2AkCuy3u9x1Xz6jzVjNgpf01SMhUvgz3kgacUQ6rmXnSpc4UVD4tL=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/PUFomqBydVJyjWf41HjnsoBmEgA2AkCuy3u9x1Xz6jzVjNgpf01SMhUvgz3kgacUQ6rmXnSpc4UVD4tL=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/PUFomqBydVJyjWf41HjnsoBmEgA2AkCuy3u9x1Xz6jzVjNgpf01SMhUvgz3kgacUQ6rmXnSpc4UVD4tL=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/PUFomqBydVJyjWf41HjnsoBmEgA2AkCuy3u9x1Xz6jzVjNgpf01SMhUvgz3kgacUQ6rmXnSpc4UVD4tL=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/PUFomqBydVJyjWf41HjnsoBmEgA2AkCuy3u9x1Xz6jzVjNgpf01SMhUvgz3kgacUQ6rmXnSpc4UVD4tL=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/PUFomqBydVJyjWf41HjnsoBmEgA2AkCuy3u9x1Xz6jzVjNgpf01SMhUvgz3kgacUQ6rmXnSpc4UVD4tL=w544-h544-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "yGxHDmHoqXc",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Zombie",
    "artists": [
      { "name": "The Cranberries", "id": "UCE5eDJ9T05bGzDvJ5QYsdJQ" }
    ],
    "album": { "name": "No Need To Argue", "id": "MPREb_6ydEdre7lFP" },
    "duration": 306,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/qo7-NUTRjTGRJgUyVXYG43mLc5vXYaJ6Ms0U7OGQm--G6mv5pt89vyCZHRPLXZctZBrEg-QomnK5n3XtTg=w60-h60-s-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/qo7-NUTRjTGRJgUyVXYG43mLc5vXYaJ6Ms0U7OGQm--G6mv5pt89vyCZHRPLXZctZBrEg-QomnK5n3XtTg=w120-h120-s-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/qo7-NUTRjTGRJgUyVXYG43mLc5vXYaJ6Ms0U7OGQm--G6mv5pt89vyCZHRPLXZctZBrEg-QomnK5n3XtTg=w180-h180-s-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/qo7-NUTRjTGRJgUyVXYG43mLc5vXYaJ6Ms0U7OGQm--G6mv5pt89vyCZHRPLXZctZBrEg-QomnK5n3XtTg=w226-h226-s-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/qo7-NUTRjTGRJgUyVXYG43mLc5vXYaJ6Ms0U7OGQm--G6mv5pt89vyCZHRPLXZctZBrEg-QomnK5n3XtTg=w302-h302-s-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/qo7-NUTRjTGRJgUyVXYG43mLc5vXYaJ6Ms0U7OGQm--G6mv5pt89vyCZHRPLXZctZBrEg-QomnK5n3XtTg=w544-h544-s-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  },
  {
    "videoId": "D4INE2zO9OU",
    "playlistId": "RDAMVMljUtuoFt-8c",
    "title": "Can't Stop",
    "artists": [
      { "name": "Red Hot Chili Peppers", "id": "UCrSorX845CEWXzU4Z7BojjA" }
    ],
    "album": {
      "name": "By The Way (Deluxe Version)",
      "id": "MPREb_0D0kxqmIgo2"
    },
    "duration": 269,
    "isExplicit": false,
    "thumbnails": [
      {
        "url": "https://lh3.googleusercontent.com/wyNOYrZV_oEjqpVwqLQCkL48-gK7Gn1Ouhe3WDaveui18lv0h0dkvYSPI8xjHbb9vNf4_jyRvFGUMoZEKg=w60-h60-l90-rj",
        "width": 60,
        "height": 60
      },
      {
        "url": "https://lh3.googleusercontent.com/wyNOYrZV_oEjqpVwqLQCkL48-gK7Gn1Ouhe3WDaveui18lv0h0dkvYSPI8xjHbb9vNf4_jyRvFGUMoZEKg=w120-h120-l90-rj",
        "width": 120,
        "height": 120
      },
      {
        "url": "https://lh3.googleusercontent.com/wyNOYrZV_oEjqpVwqLQCkL48-gK7Gn1Ouhe3WDaveui18lv0h0dkvYSPI8xjHbb9vNf4_jyRvFGUMoZEKg=w180-h180-l90-rj",
        "width": 180,
        "height": 180
      },
      {
        "url": "https://lh3.googleusercontent.com/wyNOYrZV_oEjqpVwqLQCkL48-gK7Gn1Ouhe3WDaveui18lv0h0dkvYSPI8xjHbb9vNf4_jyRvFGUMoZEKg=w226-h226-l90-rj",
        "width": 226,
        "height": 226
      },
      {
        "url": "https://lh3.googleusercontent.com/wyNOYrZV_oEjqpVwqLQCkL48-gK7Gn1Ouhe3WDaveui18lv0h0dkvYSPI8xjHbb9vNf4_jyRvFGUMoZEKg=w302-h302-l90-rj",
        "width": 302,
        "height": 302
      },
      {
        "url": "https://lh3.googleusercontent.com/wyNOYrZV_oEjqpVwqLQCkL48-gK7Gn1Ouhe3WDaveui18lv0h0dkvYSPI8xjHbb9vNf4_jyRvFGUMoZEKg=w544-h544-l90-rj",
        "width": 544,
        "height": 544
      }
    ]
  }
]

```
</details>

### Get search suggestions

```go
package main

import (
    "fmt"
    "github.com/raitonoberu/ytmusic"
)

func main() {
    result, err := ytmusic.GetSearchSuggestions("fka tw")
    if err != nil {
        panic(err)
    }

    for _, query := range result {
        fmt.Println(query)
    }
}
```
<details>
    <summary>Example result</summary>

```
fka twigs
fka twig
fukk sleep (feat. fka twigs) asap rocky audio
two weeks fka twigs
fka twigs two weeks
fka twigs live
fka twigs sad day
```
</details>

## Information

I have neither the desire nor the time to port the entire [ytmusicapi](https://github.com/sigma67/ytmusicapi) to Go, so I ported only those functions that I needed. Contributions are welcome, feel free to add missing features.

### TODO:
- Add search suggestions
- Add tests
- Add docs & comments

## License

**MIT License**, see [LICENSE](./LICENSE) file for additional information.