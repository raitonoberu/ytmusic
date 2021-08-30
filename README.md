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
	s := ytmusic.Search("summer vibes")
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
      "videoId":"5vidYMzCduk",
      "playlistId":"RDAMVM5vidYMzCduk",
      "title":"Summer Vibes",
      "artists":[
        {
          "name":"The Vibest",
          "id":"UCpcafGwguVVCZXjK7hc5t3g"
        },
        {
          "name":"LZY.",
          "id":"UC8usDEBiYET4giH1f0e3r7w"
        }
      ],
      "album":{
        "name":"Summer Vibes",
        "id":"MPREb_hWNtpxDSsbB"
      },
      "duration":"3:16",
      "isExplicit":false,
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/FcL0_YTDsRez1YxlGKLrRxPCa4SS3eGB-rp-t8aIIQZaIwYwSoUZ9vDT2qj1B0AByNYyrc_rJ4GSbIZzuw=w60-h60-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/FcL0_YTDsRez1YxlGKLrRxPCa4SS3eGB-rp-t8aIIQZaIwYwSoUZ9vDT2qj1B0AByNYyrc_rJ4GSbIZzuw=w120-h120-l90-rj",
          "width":120,
          "height":120
        }
      ]
    },
    {
      "videoId":"3wQjpZiHlwY",
      "playlistId":"RDAMVM3wQjpZiHlwY",
      "title":"Summer Vibes",
      "artists":[
        {
          "name":"Mylène \u0026 Rosanne",
          "id":"UC5G0ghL0bECKwQqr6Q8qQNg"
        }
      ],
      "album":{
        "name":"Summer Vibes",
        "id":"MPREb_vVMWP143xCH"
      },
      "duration":"3:18",
      "isExplicit":false,
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/hKF17YbojU-cTTiPtPTQmrbGJ5tE42m7mClzrUly-OoMc5AyWsjEW14C61ISd9pAEHqAH4jtvnfRmrQ=w60-h60-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/hKF17YbojU-cTTiPtPTQmrbGJ5tE42m7mClzrUly-OoMc5AyWsjEW14C61ISd9pAEHqAH4jtvnfRmrQ=w120-h120-l90-rj",
          "width":120,
          "height":120
        }
      ]
    },
    {
      "videoId":"KH916UU3DhQ",
      "playlistId":"RDAMVMKH916UU3DhQ",
      "title":"Summer Vibes",
      "artists":[
        {
          "name":"Dejan S",
          "id":"UCIN8xHW7GHYL6qw7-SHWUrA"
        }
      ],
      "album":{
        "name":"Chords",
        "id":"MPREb_7ngvy4Oe3gB"
      },
      "duration":"4:24",
      "isExplicit":false,
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/aByfvRbI1AHF9fJLi9vIR16m3EoHk8q_XhzRhFw-I5-fu9HKU05fa8Zu_Pwaofx7X9UDk_lzUXYCtpJRJA=w60-h60-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/aByfvRbI1AHF9fJLi9vIR16m3EoHk8q_XhzRhFw-I5-fu9HKU05fa8Zu_Pwaofx7X9UDk_lzUXYCtpJRJA=w120-h120-l90-rj",
          "width":120,
          "height":120
        }
      ]
    }
  ],
  "artists":[
    {
      "browseId":"UCSU4CYKzBxdemZASnAMumbQ",
      "artist":"Summer vibes",
      "shuffleId":"RDAOBVHasxZj8-yv-IEuQYp2UQ",
      "radioId":"RDAOBVHasxZj8-yv-IEuQYp2UQ",
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/AErY4X5cICvuRrT0M-GEKT3t8QPYTHsSTr0Ng9WlLTw8atvzdoS5woyi0WMVWqrlQ1oyXGmJzXAkXD_O=w60-h60-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/AErY4X5cICvuRrT0M-GEKT3t8QPYTHsSTr0Ng9WlLTw8atvzdoS5woyi0WMVWqrlQ1oyXGmJzXAkXD_O=w120-h120-l90-rj",
          "width":120,
          "height":120
        }
      ]
    },
    {
      "browseId":"UCCWUf6HxTNr7J2UVbt6tRQg",
      "artist":"Mauve",
      "shuffleId":"RDAOtnZvjIDCvguZSA2cHrqHUQ",
      "radioId":"RDAOtnZvjIDCvguZSA2cHrqHUQ",
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/9DFl4mOmbQ72aAGke7VYV51rLniQkbO1BZ3n2vSPa06tN_8vGPx2MmpAfIQCNx0TvIKCMRO-8KqhjA=w60-h60-p-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/9DFl4mOmbQ72aAGke7VYV51rLniQkbO1BZ3n2vSPa06tN_8vGPx2MmpAfIQCNx0TvIKCMRO-8KqhjA=w120-h120-p-l90-rj",
          "width":120,
          "height":120
        }
      ]
    },
    {
      "browseId":"UCRuLdiD8id_rOlYAx6wS4xQ",
      "artist":"Retro Jungle",
      "shuffleId":"RDAOEWNnLW5XDUk8u-1xG27CWQ",
      "radioId":"RDAOEWNnLW5XDUk8u-1xG27CWQ",
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/0eFg4fksKrb0hTG97ni-kdf2mleYjyolNfDQZLCel1KTfs7wHu01kFZmEUuKBDXxjT9v2m-O4E-3Eac=w60-h60-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/0eFg4fksKrb0hTG97ni-kdf2mleYjyolNfDQZLCel1KTfs7wHu01kFZmEUuKBDXxjT9v2m-O4E-3Eac=w120-h120-l90-rj",
          "width":120,
          "height":120
        }
      ]
    }
  ],
  "albums":[
    {
      "browseId":"MPREb_AYKLwswJX9L",
      "title":"Summer Vibes",
      "type":"Single",
      "artists":null,
      "year":"2021",
      "isExplicit":false,
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/Gh4RB90F2gHsK-X_kWNRKgoFCOhnogD2X9Hz2IeD6FkWDAxrs_U_ZfyST7NwiTxyHTDm0Bq64rSA-PRm=w60-h60-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/Gh4RB90F2gHsK-X_kWNRKgoFCOhnogD2X9Hz2IeD6FkWDAxrs_U_ZfyST7NwiTxyHTDm0Bq64rSA-PRm=w120-h120-l90-rj",
          "width":120,
          "height":120
        },
        {
          "url":"https://lh3.googleusercontent.com/Gh4RB90F2gHsK-X_kWNRKgoFCOhnogD2X9Hz2IeD6FkWDAxrs_U_ZfyST7NwiTxyHTDm0Bq64rSA-PRm=w226-h226-l90-rj",
          "width":226,
          "height":226
        },
        {
          "url":"https://lh3.googleusercontent.com/Gh4RB90F2gHsK-X_kWNRKgoFCOhnogD2X9Hz2IeD6FkWDAxrs_U_ZfyST7NwiTxyHTDm0Bq64rSA-PRm=w544-h544-l90-rj",
          "width":544,
          "height":544
        }
      ]
    },
    {
      "browseId":"MPREb_Su87ovU5iGS",
      "title":"Summer Vibes, Vol. 1",
      "type":"Album",
      "artists":null,
      "year":"2021",
      "isExplicit":false,
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/VJPl48E179gQa-UWfqN-vWu3KIbYo195pIXOCN3Xt2hoLtYdkWQ3J7G5uo-9pF0_Db_7IL0vmWjP0Zwi=w60-h60-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/VJPl48E179gQa-UWfqN-vWu3KIbYo195pIXOCN3Xt2hoLtYdkWQ3J7G5uo-9pF0_Db_7IL0vmWjP0Zwi=w120-h120-l90-rj",
          "width":120,
          "height":120
        },
        {
          "url":"https://lh3.googleusercontent.com/VJPl48E179gQa-UWfqN-vWu3KIbYo195pIXOCN3Xt2hoLtYdkWQ3J7G5uo-9pF0_Db_7IL0vmWjP0Zwi=w226-h226-l90-rj",
          "width":226,
          "height":226
        },
        {
          "url":"https://lh3.googleusercontent.com/VJPl48E179gQa-UWfqN-vWu3KIbYo195pIXOCN3Xt2hoLtYdkWQ3J7G5uo-9pF0_Db_7IL0vmWjP0Zwi=w544-h544-l90-rj",
          "width":544,
          "height":544
        }
      ]
    },
    {
      "browseId":"MPREb_KyjqW5Y4IFU",
      "title":"Summer Vibes, Vol. 3",
      "type":"EP",
      "artists":[
        {
          "name":"Sir Henrik",
          "id":"UC9O4GQr91AMhTKB_mDR7AvA"
        }
      ],
      "year":"2021",
      "isExplicit":false,
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/CrAaclEP2m2rCKE2xldx-cDb2jBYOWT72cZs55JgAL8JG34XJCli_9UBVMrM3cyjJexcjuzifyhLq6vHdQ=w60-h60-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/CrAaclEP2m2rCKE2xldx-cDb2jBYOWT72cZs55JgAL8JG34XJCli_9UBVMrM3cyjJexcjuzifyhLq6vHdQ=w120-h120-l90-rj",
          "width":120,
          "height":120
        },
        {
          "url":"https://lh3.googleusercontent.com/CrAaclEP2m2rCKE2xldx-cDb2jBYOWT72cZs55JgAL8JG34XJCli_9UBVMrM3cyjJexcjuzifyhLq6vHdQ=w226-h226-l90-rj",
          "width":226,
          "height":226
        },
        {
          "url":"https://lh3.googleusercontent.com/CrAaclEP2m2rCKE2xldx-cDb2jBYOWT72cZs55JgAL8JG34XJCli_9UBVMrM3cyjJexcjuzifyhLq6vHdQ=w544-h544-l90-rj",
          "width":544,
          "height":544
        }
      ]
    }
  ],
  "playlists":[
    {
      "browseId":"VLRDCLAK5uy_mycqfrKYZscut0P3euskox4nw16DtjUxs",
      "title":"Summer of Rap",
      "author":"YouTube Music",
      "itemCount":"59 songs",
      "thumbnails":[
        {
          "url":"https://lh3.googleusercontent.com/fB_ICXdPVAkmmrJrH1JBy8y6t6jdMxkeYlH9qF-AEkq06PmfTHgeXcd3xXLZD2PnnPnWRuTqjRHU4fI3=w60-h60-l90-rj",
          "width":60,
          "height":60
        },
        {
          "url":"https://lh3.googleusercontent.com/fB_ICXdPVAkmmrJrH1JBy8y6t6jdMxkeYlH9qF-AEkq06PmfTHgeXcd3xXLZD2PnnPnWRuTqjRHU4fI3=w120-h120-l90-rj",
          "width":120,
          "height":120
        },
        {
          "url":"https://lh3.googleusercontent.com/fB_ICXdPVAkmmrJrH1JBy8y6t6jdMxkeYlH9qF-AEkq06PmfTHgeXcd3xXLZD2PnnPnWRuTqjRHU4fI3=w226-h226-l90-rj",
          "width":226,
          "height":226
        },
        {
          "url":"https://lh3.googleusercontent.com/fB_ICXdPVAkmmrJrH1JBy8y6t6jdMxkeYlH9qF-AEkq06PmfTHgeXcd3xXLZD2PnnPnWRuTqjRHU4fI3=w544-h544-l90-rj",
          "width":544,
          "height":544
        }
      ]
    },
    {
      "browseId":"VLPLmOk00V-7RN6_6tsHyVsRLqQ-wj7rqUbq",
      "title":"Hot Summer Songs 2021 - Best Summer Vibes Music 2021 to 2022",
      "author":"Redlist Music",
      "itemCount":"80 songs",
      "thumbnails":[
        {
          "url":"https://yt3.ggpht.com/7eN41ekTvxDg6B5AwloJoRdce69eE6UvYhDPRcsUFqd073I1LvL_X0DLLqhZZKRbwsZxIA2u27c=s192",
          "width":192,
          "height":192
        },
        {
          "url":"https://yt3.ggpht.com/7eN41ekTvxDg6B5AwloJoRdce69eE6UvYhDPRcsUFqd073I1LvL_X0DLLqhZZKRbwsZxIA2u27c=s576",
          "width":576,
          "height":576
        },
        {
          "url":"https://yt3.ggpht.com/7eN41ekTvxDg6B5AwloJoRdce69eE6UvYhDPRcsUFqd073I1LvL_X0DLLqhZZKRbwsZxIA2u27c=s1200",
          "width":1200,
          "height":1200
        }
      ]
    },
    {
      "browseId":"VLPLa2a9FJY91_2BEPSAuRdQMZfXC8OKszgP",
      "title":"☀ Summer Vibes 2021 ☀",
      "author":"The Vibe Guide",
      "itemCount":"296 songs",
      "thumbnails":[
        {
          "url":"https://yt3.ggpht.com/Ja5DzoZ9EcFtJWxm1UsYxmaqytxjByscPYmAwipp7wm530ZWjYhGOHxbtbdmy6Q4WRp3A4p1GA=s192",
          "width":192,
          "height":192
        },
        {
          "url":"https://yt3.ggpht.com/Ja5DzoZ9EcFtJWxm1UsYxmaqytxjByscPYmAwipp7wm530ZWjYhGOHxbtbdmy6Q4WRp3A4p1GA=s576",
          "width":576,
          "height":576
        },
        {
          "url":"https://yt3.ggpht.com/Ja5DzoZ9EcFtJWxm1UsYxmaqytxjByscPYmAwipp7wm530ZWjYhGOHxbtbdmy6Q4WRp3A4p1GA=s1200",
          "width":1200,
          "height":1200
        }
      ]
    },
    {
      "browseId":"VLPLasIJTg1l6iXQXeODV-K8qm1rHdjTP5pd",
      "title":"🌴SUMMER VIBES",
      "author":"Rammor",
      "itemCount":"16 songs",
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/a8vLltFM8kQ/sddefault.jpg?sqp=-oaymwEWCJADEOEBIAQqCghqEJQEGHgg6AJIWg\u0026rs=AMzJL3kRJ4oN3FQ6Q9ujnc1CZYYLlz2MtA",
          "width":400,
          "height":400
        },
        {
          "url":"https://i.ytimg.com/vi/a8vLltFM8kQ/hq720.jpg?sqp=-oaymwEXCKAGEMIDIAQqCwjVARCqCBh4INgESFo\u0026rs=AMzJL3mQs7PeBqYwPqF2PtwnXxoChcSjfQ",
          "width":800,
          "height":800
        },
        {
          "url":"https://i.ytimg.com/vi/a8vLltFM8kQ/hq720.jpg?sqp=-oaymwEXCNUGEOADIAQqCwjVARCqCBh4INgESFo\u0026rs=AMzJL3lkOgMac7iXDUP3yZfZGfUUYuh_2w",
          "width":853,
          "height":853
        }
      ]
    }
  ],
  "videos":[
    {
      "videoId":"YjYjrd_ozQw",
      "playlistId":"RDAMVMYjYjrd_ozQw",
      "title":"Kygo, Avicii, Robin Schulz, Felix Jaehn, Alok, Lost Frequencies - Summer Vibes Deep House Mix",
      "artists":null,
      "views":"12M views",
      "duration":"43:10",
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/YjYjrd_ozQw/sddefault.jpg?sqp=-oaymwEWCJADEOEBIAQqCghqEJQEGHgg6AJIWg\u0026rs=AMzJL3kQZZz_ekikeAwD08C_CJrIO5k9zA",
          "width":400,
          "height":400
        }
      ]
    },
    {
      "videoId":"JS6uhdQJo4A",
      "playlistId":"RDAMVMJS6uhdQJo4A",
      "title":"Avicii, Jonas Blue, Kygo, Calvin Harris, Alok, Robin Schulz - Summer Vibes Deep House Mix",
      "artists":null,
      "views":"27M views",
      "duration":"56:43",
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/JS6uhdQJo4A/sddefault.jpg?sqp=-oaymwEWCJADEOEBIAQqCghqEJQEGHgg6AJIWg\u0026rs=AMzJL3mmmvNk-F5NHTD2Qn5ypXGEBXaMIw",
          "width":400,
          "height":400
        }
      ]
    },
    {
      "videoId":"a8vLltFM8kQ",
      "playlistId":"RDAMVMa8vLltFM8kQ",
      "title":"Kygo, Martin Garrix, Jonas Blue, Robin Schulz, Lost Frequencies, Gryffin - Summer Vibes Mix Rammor",
      "artists":null,
      "views":"69K views",
      "duration":"50:25",
      "thumbnails":[
        {
          "url":"https://i.ytimg.com/vi/a8vLltFM8kQ/sddefault.jpg?sqp=-oaymwEWCJADEOEBIAQqCghqEJQEGHgg6AJIWg\u0026rs=AMzJL3kRJ4oN3FQ6Q9ujnc1CZYYLlz2MtA",
          "width":400,
          "height":400
        }
      ]
    }
  ]
}
```
</details>