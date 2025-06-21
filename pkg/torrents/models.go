package torrents

type Torrent struct {
	TorrentName string `json:"torrent_name"`
	Username    string `json:"username"`
	Category    string `json:"category"`
	ContentPath string `json:"content_path"`
	Size        int    `json:"size"`
	Tracker     string `json:"tracker"`
	TorrentId   string `json:"torrent_id"`
}
type TorrentStore interface {
	Add(name string, torrent Torrent) error
	//Get(name string) (Torrent, error)
	//Update(name string, torrent Torrent) error
	//List() (map[string]Torrent, error)
	//Remove(name string) error
}

type StubTorrent struct {
	Torrents map[string]Torrent
}

func (s *StubTorrent) Add(name string, torrent Torrent) error {
	s.Torrents[name] = torrent
	return nil
}
