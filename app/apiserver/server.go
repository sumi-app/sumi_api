package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"sumi/app/store"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

// NewServer ...
func NewServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, _ *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}