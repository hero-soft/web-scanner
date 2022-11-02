package httpservice

import (
	"encoding/json"
	"fmt"

	"net/http"
)

func (h *HttpService) notFound(w http.ResponseWriter, r *http.Request) {
	h.logger.Warnf("404 for URL: %s", r.RequestURI)
	http.Error(w, "404 Not Found", 404)
}

func (h *HttpService) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "405 Method not allowed", 405)
}

func (h *HttpService) returnProblem(problem *Problem, w http.ResponseWriter) {
	h.logger.Errorf("encountered problem: %s", problem.Detail)
	w.Header().Add("Content-Type", "application/problem+json")
	w.WriteHeader(problem.Status)
	err := json.NewEncoder(w).Encode(problem)
	if err != nil {
		h.logger.Errorf("Could not marshall problem: %v", err)
	}
}

type Problem struct {
	Type     string `json:"type,omitempty"`
	Title    string `json:"title,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
	Status   int    `json:"status,omitempty"`
}

func problem500(err error) *Problem {
	return &Problem{
		Type:   "urn:ietf:params:acme:error:serverInternal",
		Title:  "Internal Server Error",
		Detail: err.Error(),
		Status: http.StatusInternalServerError,
	}
}

func problem400(err error) *Problem {
	return &Problem{
		Type:   "urn:ietf:params:acme:error:malformed",
		Title:  "Bad Request",
		Detail: err.Error(),
		Status: http.StatusBadRequest,
	}
}

func problemBadNonce() *Problem {
	return &Problem{
		Type:   "urn:ietf:params:acme:error:badNonce",
		Title:  "Bad Nonce",
		Detail: "The client sent an unacceptable anti-replay nonce",
		Status: http.StatusBadRequest,
	}
}

func problemBadPublicKey() *Problem {
	return &Problem{
		Type:   "urn:ietf:params:acme:error:badPublicKey",
		Title:  "Bad Public Key",
		Detail: "The JWS was signed by an invalid public key",
		Status: http.StatusBadRequest,
	}
}

func problemBadSignature() *Problem {
	return &Problem{
		Type:   "urn:ietf:params:acme:error:malformed",
		Title:  "Invalid JWS Signature",
		Detail: "The JWS signature is invalid",
		Status: http.StatusBadRequest,
	}
}

func problemAccountDNE(kid string) *Problem {
	return &Problem{
		Type:   "urn:ietf:params:acme:error:accountDoesNotExist",
		Title:  "Account Does Not Exist",
		Detail: fmt.Sprintf("The request specified an account that does not exist (%s)", kid),
		Status: http.StatusBadRequest,
	}
}
