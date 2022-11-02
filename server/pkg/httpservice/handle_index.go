package httpservice

import (
	"fmt"
	"net/http"
)

func (h HttpService) index(w http.ResponseWriter, r *http.Request) {
	h.counters["index_hits"].Inc()

	//http.Redirect(w, r, "/app", 301)
	w.Header().Set("Content-Type", "text/JSON; charset=UTF-8")
	fmt.Fprintf(w, "Key Engineering Power Survey Backend")
}
