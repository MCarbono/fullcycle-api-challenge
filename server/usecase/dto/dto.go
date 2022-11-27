package dto

import "fullcycle-api-challenge/server/entity"

type DollarCotationDTO struct {
	Bid string `json:"bid"`
}

func NewDollarCotationDTO(dollarCotation *entity.DollarCotation) *DollarCotationDTO {
	return &DollarCotationDTO{
		Bid: dollarCotation.Bid,
	}
}
