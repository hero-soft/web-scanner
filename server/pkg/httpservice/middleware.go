package httpservice

import (
	"fmt"
	"net/http"
	"time"
)

func debugHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)

		fmt.Println(w.Header())
	})
}

func (h *HttpService) secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")

		// if h.permissiveHeaders {
		w.Header().Set("Access-Control-Allow-Headers", "content-type,authorization,cache-control,ngsw-bypass,pragma")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// }

		if r.Method != "OPTIONS" {
			next.ServeHTTP(w, r)
		}
	})
}

func (h *HttpService) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		h.logger.Infow(r.URL.RequestURI(), "method", r.Method, "protocol", r.Proto, "duration", time.Since(start), "remote_address", r.RemoteAddr)

	})
}

/*
func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
*/
