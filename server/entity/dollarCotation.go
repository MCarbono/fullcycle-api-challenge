package entity

import "fullcycle-api-challenge/server/gateway"

type DollarCotation struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func NewDollarCotation(id string, data *gateway.DollarCotationGatewayResponse) *DollarCotation {
	return &DollarCotation{
		ID:         id,
		Code:       data.USDBRL.Code,
		Codein:     data.USDBRL.Codein,
		Name:       data.USDBRL.Name,
		High:       data.USDBRL.High,
		Low:        data.USDBRL.Low,
		VarBid:     data.USDBRL.VarBid,
		PctChange:  data.USDBRL.PctChange,
		Bid:        data.USDBRL.Bid,
		Ask:        data.USDBRL.Ask,
		Timestamp:  data.USDBRL.Timestamp,
		CreateDate: data.USDBRL.CreateDate,
	}
}
