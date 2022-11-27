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
	db, err := sql.Open("sqlite3", "./dollar_cotation.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := repository.NewDollarCotationRepository(db)
	handler := handler.DollarCotationHandler{DollarCotationUseCase: usecase.DollarCotationUseCase{DollarCotationRepository: *repository}}
	http.HandleFunc("/cotacao", handler.Get)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
