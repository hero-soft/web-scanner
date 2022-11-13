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

	// b, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	r.ParseMultipartForm(6000000)

	f, header, _ := r.FormFile("call")

	// fmt.Printf("%v\n", r.MultipartForm.Value["api_key"])

	// fmt.Println(string(b))

	filename := filepath.Join("audio", header.Filename)

	if _, err := os.Stat(filename); err == nil {
		// this prevents duplicate files
		return
	}

	// fmt.Println(filename)

	fileWriter, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)

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
