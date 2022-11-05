package httpservice

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func (h HttpService) index(w http.ResponseWriter, r *http.Request) {
	h.counters["index_hits"].Inc()

	//http.Redirect(w, r, "/app", 301)
	w.Header().Set("Content-Type", "text/JSON; charset=UTF-8")
	fmt.Fprintf(w, "Key Engineering Power Survey Backend")
}

func (h HttpService) audio(w http.ResponseWriter, r *http.Request) {

}

func (h HttpService) test(w http.ResponseWriter, r *http.Request) {
	// h.counters["index_hits"].Inc()

	//http.Redirect(w, r, "/app", 301)

	fmt.Println("got upload request")

	// b, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	r.ParseMultipartForm(6000000)

	f, header, _ := r.FormFile("call")

	// fmt.Printf("%v", r.MultipartForm.Value)
	// fmt.Println(string(b))

	filename := filepath.Join("audio", header.Filename)

	fmt.Println(filename)

	fileWriter, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)

	io.Copy(fileWriter, f)

}
