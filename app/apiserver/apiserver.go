package apiserver

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"sumi/app/store"
	"sumi/app/store/sqlstore"
)

func Start(config *Config) error {

	store, err := configureStore(config)
	if err != nil {
		return err
	}

	var port string
	envPort := os.Getenv("DATABASE_URL")
	if len(envPort) < 0{
		port = "8080"
	} else {
		port = envPort
	}



	srv := NewServer(store)
	fmt.Println("Api server running on port", port)
	return http.ListenAndServe(fmt.Sprintf(":%v", port), srv)
}


func configureStore(config *Config) (store.Store, error) {

	db, err := newDbConnect(config.DatabaseUrl)
	if err != nil {
		return nil, err
	}

	s := sqlstore.New(db)
	return s, nil
}

func newDbConnect(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
