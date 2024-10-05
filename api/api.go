package api

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alokmenghrajani/go-moth/utils"
	"github.com/gorilla/mux"
)

//go:embed index.html
//go:embed encrypt.html
//go:embed moth.json
var staticContent embed.FS

type state struct {
	home           string
	encrypt        string
	encryptedMoths [][]byte
	plaintextMoths []string
}

func Start() {
	s := state{}
	s.home = string(utils.MustRead(staticContent, "index.html"))
	s.encrypt = string(utils.MustRead(staticContent, "encrypt.html"))

	mothRaw := utils.MustRead(staticContent, "moth.json")
	json.Unmarshal(mothRaw, &s.encryptedMoths)

	for i := 0; i < len(s.encryptedMoths); i++ {
		plaintext := utils.Decrypt(s.encryptedMoths[i])
		s.plaintextMoths = append(s.plaintextMoths, string(plaintext))
	}

	r := mux.NewRouter()
	r.HandleFunc("/", s.homeHandler).Methods("GET")
	r.HandleFunc("/moth", s.mothHandler).Methods("GET")
	r.HandleFunc("/encrypt", s.encryptGetHandler).Methods("GET")
	r.HandleFunc("/encrypt", s.encryptPostHandler).Methods("POST")
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("0.0.0.0:%d", utils.DEFAULT_PORT),
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
