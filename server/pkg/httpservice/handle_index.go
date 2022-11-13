package httpservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hero-soft/web-scanner/pkg/talkgroup"
)

func (h HttpService) index(w http.ResponseWriter, r *http.Request) {
	h.counters["index_hits"].Inc()

	//http.Redirect(w, r, "/app", 301)
	w.Header().Set("Content-Type", "text/JSON; charset=UTF-8")
	fmt.Fprintf(w, "Hero Web Scanner")
}

func (h HttpService) talkgroups(w http.ResponseWriter, r *http.Request) {
	h.counters["index_hits"].Inc()

	tgs, err := talkgroup.GetAll()

	if err != nil {
		h.logger.Errorf("Error getting talkgroups: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/JSON; charset=UTF-8")

	b, err := json.MarshalIndent(tgs, "", "  ")

	if err != nil {
		h.logger.Errorf("Error marshalling talkgroups: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(b))
}
