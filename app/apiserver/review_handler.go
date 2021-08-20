package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sumi/app/models"
	"sumi/app/utils"
)

func (s *server) HandleCreateReview() http.HandlerFunc {
	var review models.Review

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

		if err := json.Unmarshal(body, &review); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			fmt.Println(err)
			return
		}

		if err := review.Validate(); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		createdReview, err := s.store.Review().Create(&review)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusCreated, createdReview)
	}
}

func (s *server) HandleGetReviews() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		pressTourId := r.URL.Query().Get("press_tour_id")

		if len(pressTourId) > 0 {
			id, err := strconv.Atoi(pressTourId)
			if err == nil{
				bloggers, err := s.store.Review().GetByPressTourId(id)
				if err != nil {
					s.error(w, r, http.StatusInternalServerError, err)
					return
				}
				s.respond(w, r, http.StatusOK, bloggers)
				return
			}
		}

		bloggers, err := s.store.Review().GetAll()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, bloggers)
	}
}

func (s *server) HandleDeleteReviews() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		err := s.store.Review().Delete()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}
