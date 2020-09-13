//go:generate statik -src=./public -include=*.html,*.css,*.js,*.png,*.ico,*.woff,*.woff2,*.eot,*.ttf

package main

import (
	"flag"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	st "github.com/ad/pwnmothership/static"

	// go get github.com/rakyll/statik
	_ "github.com/ad/pwnmothership/statik"
)

var (
	addr   *string
	static *st.FS
	p = make(map[string]*Pwnagotchi)
)

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
	UPS         string  `json:"ups,omitempty"`

	// "peers": [],
	PeersCount int64 `json:"num_peers,omitempty"`

	PwnedLast   string `json:"pwnd_last,omitempty"`
	PwnedRun    string `json:"pwnd_run,omitempty"`
	PwnedTotal  int64  `json:"pwnd_tot,omitempty"`
	PwnedDeauth string `json:"pwnd_deauth,omitempty"`

	TotalMessages  int64 `json:"total_messages,omitempty"`
	UnreadMessages int64 `json:"unread_messages,omitempty"`

	Level string `json:"level,omitempty"`
	Exp   string `json:"exp,omitempty"`
}

func main() {
	addr = flag.String("addr", ":8080", "listen address")
	flag.Parse()
	
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

		if hash == "test" {
			d := &Pwnagotchi{
				Fingerprint: "",
				Initialised: true,

				Name:   randomString([]string{"pwnagotchi", "username", "test", "pet", "cat", "dog"}),
				Mode:   randomString([]string{"AUTO", "MANU", "AI"}),
				Status: randomString([]string{"...", "ololo", "blabla", "Связываюсь с 2020", "Эй, 2020 давай дружить!", "Ждем 5c …", "Хррррр.. (29c)", "Дремлет 19с …", "Осматриваюсь вокруг (3с)", "Просто решил, что 00:12:34:56:78:0a не нужен WiFi! Кхе-кхе)"}),

				Uptime:  randomString([]string{"00:00:01", "12:34:56", "23:45:00"}),
				Version: "1.5.3",
				Epoch:   randomInt64([]int64{1, 2, 3, 10, 20, 300}),

				APsOnChannel: randomInt64([]int64{1, 2, 3, 10, 20, 300}),
				APsName:      randomString([]string{"1 (1)", "2 (5)", "10 (100)"}),
				APsTotal:     randomInt64([]int64{1, 2, 3, 10, 20, 300}),
				Channel:      randomInt64([]int64{1, 2, 3, 10, 20, 300}),
				ChannelName:  randomString([]string{"*", "1", "11"}),

				Face: randomString([]string{"(⇀‿‿↼)", "(≖‿‿≖)", "(◕‿‿◕)", "( ⚆_⚆)", "(☉_☉ )", "( ◕‿◕)", "(◕‿◕ )", "(°▃▃°)", "(⌐■_■)", "(•‿‿•)", "(^‿‿^)", "(ᵔ◡◡ᵔ)", "(✜‿‿✜)", "(♥‿‿♥)", "(☼‿‿☼)", "(≖__≖)", "(-__-)", "(╥☁╥ )", "(ب__ب)", "(☓‿‿☓)", "(#__#)"}),

				FriendFace: randomString([]string{"(⇀‿‿↼)", "(≖‿‿≖)", "(◕‿‿◕)", "( ⚆_⚆)", "(☉_☉ )", "( ◕‿◕)", "(◕‿◕ )", "(°▃▃°)", "(⌐■_■)", "(•‿‿•)", "(^‿‿^)", "(ᵔ◡◡ᵔ)", "(✜‿‿✜)", "(♥‿‿♥)", "(☼‿‿☼)", "(≖__≖)", "(-__-)", "(╥☁╥ )", "(ب__ب)", "(☓‿‿☓)", "(#__#)"}),
				FriendName: randomString([]string{"username", "test", "pet"}),

				CPU:         0.5,
				Memory:      0.5,
				Temperature: 35.5,
				Bluetooth:   randomString([]string{"-", "C", "NF"}),
				UPS:         randomString([]string{"", "0%", "1%", "10%", "11%", "100%"}),

				PeersCount: randomInt64([]int64{1, 2, 3, 10, 20, 300}),

				PwnedLast:   randomString([]string{"pwnagotchi", "username", "test", "pet", "cat", "dog"}),
				PwnedRun:    "2",
				PwnedTotal:  randomInt64([]int64{1, 2, 3, 10, 20, 300}),
				PwnedDeauth: "4",

				TotalMessages:  randomInt64([]int64{1, 2, 3, 10, 20, 300}),
				UnreadMessages: randomInt64([]int64{1, 2, 3, 10, 20, 300}),

				Level: randomString([]string{"-", "1", "11", "50", "99"}),
				Exp:   randomString([]string{"╷          ╷", "╷▄         ╷", "╷▄▄        ╷", "╷▄▄▄       ╷", "╷▄▄▄▄      ╷", "╷▄▄▄▄▄     ╷", "╷▄▄▄▄▄▄    ╷", "╷▄▄▄▄▄▄▄   ╷", "╷▄▄▄▄▄▄▄▄  ╷", "╷▄▄▄▄▄▄▄▄▄▄╷"}),
			}

			b, _ := json.Marshal(d)
			w.Write(b)
			return
		}

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

	log.Println("start listening on", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func randomInt64(list []int64) int64 {
	rand.Seed(time.Now().UnixNano())

	return list[rand.Intn(len(list))]
}

func randomString(list []string) string {
	rand.Seed(time.Now().UnixNano())

	return list[rand.Intn(len(list))]
}
