package repository

import (
	"context"
	"database/sql"
	"fullcycle-api-challenge/server/entity"
	"time"
)

type DollarCotationRepository struct {
	Db *sql.DB
}

func NewDollarCotationRepository(db *sql.DB) *DollarCotationRepository {
	return &DollarCotationRepository{Db: db}
}

func (r *DollarCotationRepository) Save(dollarCotation *entity.DollarCotation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err := r.Db.ExecContext(ctx, "insert into cotations values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", dollarCotation.ID,
		dollarCotation.Code,
		dollarCotation.Codein,
		dollarCotation.Name,
		dollarCotation.High,
		dollarCotation.Low,
		dollarCotation.VarBid,
		dollarCotation.PctChange,
		dollarCotation.Bid,
		dollarCotation.Ask,
		dollarCotation.Timestamp,
		dollarCotation.CreateDate)
	return err
}
