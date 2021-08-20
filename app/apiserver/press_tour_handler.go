package apiserver



import (
"encoding/json"
"fmt"
"io/ioutil"
"net/http"
"sumi/app/models"
"sumi/app/utils"
)

func (s *server) HandleCreatePressTour() http.HandlerFunc {
	var p models.PressTour

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			s.error(w, r, http.StatusBadRequest, utils.ErrNoBodyData)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := json.Unmarshal(body, &p); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			fmt.Println(err)
			return
		}

		createdReview, err := s.store.PressTour().Create(&p)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusCreated, createdReview)
	}
}

func (s *server) HandleGetPressTours() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		bloggers, err := s.store.PressTour().GetAll()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, bloggers)
	}
}

func (s *server) HandleDeletePressTours() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		err := s.store.PressTour().Delete()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}
