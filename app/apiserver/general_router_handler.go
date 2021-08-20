package apiserver

import (
"net/http"
)

func (s *server) HandleWarmUp() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, nil)
	}
}
