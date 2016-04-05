package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"gitlab.innoq.com/daniel/go-workshop-utils/search/images/scraper"
)

func main() {
	addr := flag.String("addr", "localhost:4242", "where to listen for connection")
	timeout := flag.Duration("timeout", 2*time.Second, "longest time to wait for requests to finish")

	handleRequest := func(w http.ResponseWriter, req *http.Request) {
		if path := req.URL.Path; path == "/" {
			t := *timeout

			if req.FormValue("no-timeout") == "1" {
				t = 1 * time.Hour
			}

			images, stats, err := MetaImageSearch(req.FormValue("q"), t)

			data := struct {
				Images []scraper.Image
				Stats  Stats
				Err    error
			}{Err: err, Images: images, Stats: stats}

			if req.FormValue("format") == "json" {
				w.Header().Add("Content-Type", "application/json")

				b, err := json.Marshal(data)
				if err != nil {
					log.Fatal(err)
				}

				var out bytes.Buffer
				json.Indent(&out, b, " ", "\t")
				out.WriteTo(w)
				return
			}

			indexTemplate.Execute(w, data)

			return
		}
		http.FileServer(http.Dir("static")).ServeHTTP(w, req)
	}

	flag.Parse()

	fmt.Printf("open http://%s", *addr)

	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(*addr, nil)
}
