package apiserver

func (s *server) configureRouter() {
	apiRouter := s.router.PathPrefix("/api").Subrouter()
	apiRouter.Use(s.addJsonHeader)
	apiRouter.Use(s.logRequest)

	bloggerRouter := apiRouter.PathPrefix("/blogger").Subrouter()
	bloggerRouter.HandleFunc("", s.HandleCreateBlogger()).Methods("POST")
	bloggerRouter.HandleFunc("", s.HandleGetBloggers()).Methods("GET")
	bloggerRouter.HandleFunc("", s.HandleDeleteBloggers()).Methods("DELETE")


	generalRouter := apiRouter.PathPrefix("/general").Subrouter()
	generalRouter.HandleFunc("/warm-up", s.HandleWarmUp()).Methods("POST")

	reviewRouter := apiRouter.PathPrefix("/review").Subrouter()
	reviewRouter.HandleFunc("", s.HandleCreateReview()).Methods("POST")
	reviewRouter.HandleFunc("", s.HandleGetReviews()).Methods("GET")
	reviewRouter.HandleFunc("", s.HandleDeleteReviews()).Methods("DELETE")

	pressTourRouter := apiRouter.PathPrefix("/press-tour").Subrouter()
	pressTourRouter.HandleFunc("", s.HandleCreatePressTour()).Methods("POST")
	pressTourRouter.HandleFunc("", s.HandleGetPressTours()).Methods("GET")
	pressTourRouter.HandleFunc("", s.HandleDeletePressTours()).Methods("DELETE")
}