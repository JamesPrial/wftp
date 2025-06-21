package main

import (
	"log"
	"net/http"

	torrents "github.com/jamesprial/wftp/pkg/torrents"
)

func main() {
	server := &torrents.TorrentServer{}
	log.Fatal(http.ListenAndServe(":5000", server))
}
