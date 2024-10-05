package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/alokmenghrajani/go-moth/utils"
)

type encryptResponse struct {
	Error      string `json:"error,omitempty"`
	Ciphertext []byte `json:"ciphertext,omitempty"`
}

func (s *state) encryptGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "max-age=86400")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, s.encrypt)
}

func (s *state) encryptPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		resp := encryptResponse{Error: err.Error()}
		fmt.Fprint(w, utils.MustMarshal(resp))
		return
	}

	ciphertext := utils.Encrypt(body)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, utils.MustMarshal(encryptResponse{Ciphertext: ciphertext}))
}
