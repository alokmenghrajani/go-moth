package api

import (
	"fmt"
	"net/http"
)

func (s *state) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "max-age=86400")
	fmt.Fprint(w, string(s.home))
}
