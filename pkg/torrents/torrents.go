package torrents

import (
	"encoding/json"
	"net/http"
	"strings"
)

type TorrentServer struct {
	Store TorrentStore
}

func (t *TorrentServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	torrentName := strings.TrimPrefix(r.URL.Path, "/torrents/")
	var torrent Torrent
	json.NewDecoder(r.Body).Decode(&torrent)
	t.Store.Add(torrentName, torrent)
	w.WriteHeader(http.StatusCreated)
}
