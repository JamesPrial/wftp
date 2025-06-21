package torrents_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	torrents "github.com/jamesprial/wftp/pkg/torrents"
)

func TestPOSTTorrents(t *testing.T) {

	store := torrents.StubTorrent{
		map[string]torrents.Torrent{},
	}
	server := &torrents.TorrentServer{&store}
	sampleReqStruct := torrents.Torrent{
		TorrentName: "TestTorrent", Username: "planter", Category: "", ContentPath: "/tmp/test", Size: 10, Tracker: "localhost", TorrentId: "1",
	}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(sampleReqStruct)
	t.Run("POSTs sample msg return 201", func(t *testing.T) {

		request, _ := http.NewRequest(http.MethodPost, "/torrents/1", payloadBuf)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := 201
		if got != want {
			t.Errorf("Code: got %d, want %d", got, want)
		}
		if store.Torrents["1"] != sampleReqStruct {
			t.Errorf("Buffer: got %v, want %v", store.Torrents["1"], sampleReqStruct)
		}

	})
}

func TestGETDownload(t *testing.T) {

}
