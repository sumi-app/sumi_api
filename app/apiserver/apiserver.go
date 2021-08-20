package apiserver

import (
	"database/sql"
	"fmt"
	"net/http"
	"sumi/app/store"
	"sumi/app/store/sqlstore"
)

const PORT = 8080

func Start(config *Config) error {
	store, err := configureStore(config)
	if err != nil {
		return err
	}

	srv := NewServer(store)
	fmt.Println("Api server running on port", PORT)
	return http.ListenAndServe(fmt.Sprintf(":%v",PORT), srv)
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
