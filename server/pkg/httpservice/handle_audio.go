package httpservice

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/hero-soft/web-scanner/pkg/call"
	"github.com/hero-soft/web-scanner/pkg/talkgroup"
	"github.com/hero-soft/web-scanner/pkg/websocket"
	uuid "github.com/satori/go.uuid"
)

func (h HttpService) audio(w http.ResponseWriter, r *http.Request) {

	// h.counters["index_hits"].Inc()

	//http.Redirect(w, r, "/app", 301)

	r.ParseMultipartForm(6000000)

	f, header, err := r.FormFile("call")

	if err != nil {
		h.logger.Errorf("Error getting file from form: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filename := filepath.Join("audio", header.Filename)

	if _, err := os.Stat(filename); err == nil {
		// this prevents duplicate files
		h.logger.Debugf("File %s already exists", filename)
		return
	}

	fileWriter, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		h.logger.Errorf("Error opening file: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	io.Copy(fileWriter, f)

	talkgroupID := r.FormValue("talkgroup_num")

	tg, err := talkgroup.Lookup(talkgroupID, "UNKNOWN")

	if err != nil {
		fmt.Println("Error looking up talkgroup", err)
		return
	}

	h.SendChan <- websocket.SendTo{
		To: "*",
		Message: websocket.Message{
			Type: "audio",
			Call: call.Call{
				ID:        uuid.NewV4().String(),
				Talkgroup: *tg,
				File:      filename,
			},
		},
	}

}
