package main

import (
	"database/sql"
	"fullcycle-api-challenge/server/handler"
	"fullcycle-api-challenge/server/repository"
	"fullcycle-api-challenge/server/usecase"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := repository.NewDollarCotationRepository(db)
	handler := handler.DollarCotationHandler{DollarCotationUseCase: usecase.DollarCotationUseCase{DollarCotationRepository: *repository}}
	http.HandleFunc("/cotacao", handler.Get)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "dollar_cotations.db")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE cotations (
		id varchar(255) NOT NULL,
		code varchar(255),
		codein varchar(255),
		name varchar(255),
		high varchar(255),
		low varchar(255),
		varBid varchar(255),
		pctChange varchar(255),
		bid varchar(255),
		ask varchar(255),
		timestamp TIMESTAMP,
		create_date DATE(255),
		PRIMARY KEY (id));`)
	if err != nil {
		return nil, err
	}
	return db, nil
}
