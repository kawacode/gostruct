package gostruct

import (
	"net/url"

	tls "github.com/kawacode/utls"
)

// Creating a new type called `BotData` that is a struct.
type BotData struct {
	// A struct that is used to store the data that is sent to the bot.
	HttpRequest struct {
		Response struct {
			Location        url.URL           `json:"location"`
			Status          string            `json:"status"`
			StatusCode      int               `json:"statuscode"`
			Headers         map[string]string `json:"headers"`
			Cookies         map[string]string `json:"cookies"`
			Protocol        string            `json:"protocol"`
			ContentLength   int64             `json:"contentlength"`
			Source          string            `json:"source"`
			ProtoMajor      int               `json:"protomajor"`
			ProtoMinor      int               `json:"protominor"`
			WasUncompressed bool              `json:"isuncompressed"`
		}
		Request struct {
			InsecureSkipVerify bool              `json:"insecureskipverify"`
			HelloClient        tls.ClientHelloID `json:"helloclient"`
			Proxy              string            `json:"proxy"`
			Ja3                string            `json:"ja3"`
			URL                string            `json:"url"`
			ReadResponseBody   bool              `json:"readresponse"`
			ReadResponseHeaders        bool              `json:"readheaders"`
			ReadResponseCookies        bool              `json:"readcookies"`
			Method             string            `json:"method"`
			Headers            map[string]string `json:"headers"`
			Payload            string            `json:"payload"`
			// Default 2.0
			Protocol           string   `json:"protocol"`
			Timeout            string   `json:"timeout"`
			MaxRedirects       string   `json:"maxredirects"`
			HeaderOrderKey     []string `json:"headerorderkey"`
			PHeaderOrderKey    []string `json:"pheaderorderkey"`
			DisableCompression bool     `json:"disablecompression"`
			HTTP1TRANSPORT     struct {
				ForceAttemptHTTP2 bool `json:"forceattempthttp2"`
			}
			HTTP2TRANSPORT struct {
				SettingsOrder []string `json:"settingsorder"`
				Settings      struct {
					Frames []struct {
						Key   int `json:"key"`
						Value int `json:"value"`
					} `json:"frames"`
					Priorities []struct {
						StreamID  int  `json:"streamid"`
						StreamDEP int  `json:"streamdep"`
						Exclusive bool `json:"exclusive"`
						Weight    int  `json:"weight"`
					} `json:"priorities"`
					Windowupdate int `json:"windowupdate"`
				} `json:"settings"`
			}
		}
	}
}
