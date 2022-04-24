package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func getLongUrl(shortUrl string) string {
	return "long url of " + shortUrl
}

func createShortUrl(longUrl string) string {
	return "short_url of " + longUrl
}

func main() {
	fmt.Println("URL shortener")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		surl := strings.TrimPrefix(r.URL.Path, "/")
		lurl := getLongUrl(surl)
		fmt.Fprintf(w, "Retrieved long url: %q", lurl)
	})

	http.HandleFunc("/create/", func(w http.ResponseWriter, r *http.Request) {
		lurl := strings.TrimPrefix(r.URL.Path, "/create/")
		surl := createShortUrl(lurl)
		fmt.Fprintf(w, "Created short url: %q", surl)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
