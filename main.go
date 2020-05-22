//go:generate statik -src=./public -include=*.html,*.css,*.js,*.png,*.ico

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	st "github.com/ad/pwnmothership/static"

	// go get github.com/rakyll/statik
	_ "github.com/ad/pwnmothership/statik"
)

var static *st.FS
var p = make(map[string]*Pwnagotchi)

// Pwnagotchi ...
type Pwnagotchi struct {
	Fingerprint string `json:"fingerprint"`
	Initialised bool   `json:"initialised"`

	Name   string `json:"name,omitempty"`
	Mode   string `json:"mode,omitempty"`
	Status string `json:"status,omitempty"`

	Uptime  string `json:"uptime,omitempty"`
	Version string `json:"version,omitempty"`
	Epoch   int64  `json:"epoch,omitempty"`

	APsOnChannel int64  `json:"aps_on_channel,omitempty"`
	APsName      string `json:"aps_text,omitempty"`
	APsTotal     int64  `json:"apt_tot,omitempty"`
	Channel      int64  `json:"channel,omitempty"`
	ChannelName  string `json:"channel_text,omitempty"`

	Face string `json:"face,omitempty"`

	FriendFace string `json:"friend_face_text,omitempty"`
	FriendName string `json:"friend_name_text,omitempty"`

	CPU         float64 `json:"cpu,omitempty"`
	Memory      float64 `json:"memory,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	Bluetooth   string  `json:"bluetooth,omitempty"`

	// "peers": [],
	PeersCount int64 `json:"num_peers,omitempty"`

	PwnedLast   string `json:"pwnd_last,omitempty"`
	PwnedRun    string `json:"pwnd_run,omitempty"`
	PwnedTotal  int64  `json:"pwnd_tot,omitempty"`
	PwnedDeauth string `json:"pwnd_deauth,omitempty"`

	TotalMessages  int64 `json:"total_messages,omitempty"`
	UnreadMessages int64 `json:"unread_messages,omitempty"`
}

func main() {
	static := st.NewFS()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(static.StatikFS)))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/" {
			w.Header().Set("Content-Type", "text/html")
			path = "html/default.html"
		}
		if path == "/apple-touch-icon.png" || path == "/favicon-32x32.png" || path == "/favicon-16x16.png" || path == "/favicon.ico" {
			path = "favicon" + path
		}

		if content, err := static.ReadFile(path); err == nil {
			w.Write(content)
		} else {
			http.NotFound(w, r)
			log.Println(path, err)
		}
	})

	http.HandleFunc("/api/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}

		hash := r.URL.Query().Get("hash")

		w.Header().Set("Content-Type", "application/json")

		notFound := `{"initialized": false}`

		if d, ok := p[hash]; ok {
			b, err := json.Marshal(d)
			if err != nil {
				w.Write([]byte(notFound))
				return
			}

			w.Write(b)
			return
		}

		w.Write([]byte(notFound))
	})

	http.HandleFunc("/api/set", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}

		var d Pwnagotchi
		if err := json.Unmarshal(body, &d); err != nil {
			log.Println(err)
		} else {
			if d.Fingerprint != "" {
				p[d.Fingerprint] = &d
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"success": true}`))
	})

	http.ListenAndServe("0.0.0.0:9090", nil)
}
