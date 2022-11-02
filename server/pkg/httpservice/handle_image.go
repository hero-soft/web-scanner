package httpservice

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"strings"

// 	"github.com/Key-Engineering/PowerSurvey/pkg/document"
// 	"github.com/Key-Engineering/PowerSurvey/pkg/keyimage"
// 	"github.com/gorilla/mux"
// )

// func (h *HttpService) PostAddImage(w http.ResponseWriter, r *http.Request) {

// 	body, err := io.ReadAll(r.Body)

// 	if err != nil {
// 		h.logger.Errorf("Could not read request body: %v", err)

// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprint(w, "Invalid request body")

// 		return
// 	}

// 	defer r.Body.Close()

// 	i := keyimage.Image{}

// 	err = json.Unmarshal(body, &i)

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		h.logger.Errorf("Could not unmarshall request body: %v", err)

// 		fmt.Fprint(w, "Invalid request body")

// 		return
// 	}

// 	err = keyimage.AddImage(i)

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		h.logger.Errorf("Could not add image: %v", err)

// 		fmt.Fprint(w, "Could not add image")

// 		return
// 	}

// }

// func (h *HttpService) GetImageByCode(w http.ResponseWriter, r *http.Request) {
// 	// h.counters["directory_hits"].Inc()

// 	vars := mux.Vars(r)

// 	s, err := document.GetDocumentByCode(strings.ToUpper(vars["code"]), false)

// 	if err != nil {
// 		w.WriteHeader(404)
// 		h.logger.Errorf("could not get document: %v", err)
// 		fmt.Fprint(w, "Could not get document")
// 		return
// 	}

// 	w.Header().Add("Content-Type", "application/json")

// 	_ = json.NewEncoder(w).Encode(s)

// 	// Save a copy of this request for debugging.
// 	/*requestDump, err := httputil.DumpRequest(r, true)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(requestDump))*/

// }
