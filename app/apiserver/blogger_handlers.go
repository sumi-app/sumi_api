package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sumi/app/models"
	"sumi/app/utils"
	"sumi/app/utils/convertors"
)

func (s *server) HandleCreateBlogger() http.HandlerFunc {
	var b models.Blogger

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

		if err := json.Unmarshal(body, &b); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			fmt.Println(err)
			return
		}

		if err := b.Validate(); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		//_, existBlogger := s.store.Blogger().GetByLogin(b.Login)
		//if existBlogger != nil {
		//	s.respond(w, r, http.StatusOK, existBlogger)
		//	fmt.Println("Exist blogger")
		//	return
		//}

		createdBlogger, err := s.store.Blogger().Create(&b)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusCreated, createdBlogger)
	}
}


func (s *server) HandleGetBloggers() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		bloggers, err := s.store.Blogger().GetAll()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, bloggers)
	}
}

func (s *server) HandleDeleteBloggers() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		err := s.store.Blogger().Delete()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) HandleSelectBloggers() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ids := []string{}
		idsList := r.URL.Query().Get("ids")

		if len(idsList) > 0 {
			ids = convertors.ParseStringsList(idsList)
		}

		err := s.store.Blogger().Select(ids)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}
