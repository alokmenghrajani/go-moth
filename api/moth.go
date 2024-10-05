package api

import (
	"fmt"
	"net/http"
	"time"

	_ "embed"

	"github.com/alokmenghrajani/go-moth/utils"
)

type mothResponse struct {
	Moth string `json:"moth,omitempty"`
}

func (s *state) mothHandler(w http.ResponseWriter, r *http.Request) {
	hours, _, _ := time.Now().Clock()
	moth := s.plaintextMoths[hours%len(s.plaintextMoths)]
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=3600")
	fmt.Fprint(w, utils.MustMarshal(mothResponse{Moth: moth}))
}
