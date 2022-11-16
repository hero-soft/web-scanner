package httpservice

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hero-soft/web-scanner/pkg/websocket"
)

func (h *HttpService) NewRouter(hub *websocket.Hub) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	mux.CORSMethodMiddleware(router)
	// router.Use(mux.CORSMethodMiddleware(router))
	router.Use(h.secureHeaders)
	router.Use(h.logRequest)
	//router.Use(debugHeaders)

	router.NotFoundHandler = http.HandlerFunc(h.notFound)
	router.MethodNotAllowedHandler = http.HandlerFunc(h.methodNotAllowed)

	// Index
	// router.Methods("OPTIONS").
	// 	Path("*").

	router.
		Methods("GET", "OPTIONS").
		Path("/ws/client").
		Handler(http.HandlerFunc(hub.ServeWsClient))
	router.
		Methods("GET", "OPTIONS").
		Path("/ws/recorder").
		Handler(http.HandlerFunc(hub.ServeWsRecorder))
	router.
		// Methods("GET", "POST", "OPTIONS").
		PathPrefix("/upload/").
		Handler(http.HandlerFunc(h.audio))
	router.
		// Methods("GET", "POST", "OPTIONS").
		Path("/talkgroups").
		Handler(http.HandlerFunc(h.talkgroups))
	router.
		Methods("GET", "OPTIONS").
		PathPrefix("/audio/").
		Handler(http.StripPrefix("/audio/", http.FileServer(http.Dir("./audio"))))
	router.
		Methods("GET").
		Path("/").
		Handler(http.HandlerFunc(h.index))
	router.
		Methods("GET").
		PathPrefix("/").
		Handler(http.StripPrefix("/", http.FileServer(http.Dir("./client"))))

	// router.
	// 	Methods("POST", "OPTIONS").
	// 	Path("/survey/{id}/verify").
	// 	Handler(http.HandlerFunc(h.PostVerifySurvey))

	// // File Server
	// router.
	// 	PathPrefix("/files/").
	// 	Methods("GET").
	// 	Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("./files"))))

	return router
}
