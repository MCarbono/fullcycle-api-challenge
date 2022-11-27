package gateway

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type DollarCotationGatewayResponse struct {
	USDBRL USDBRLGatewayResponse `json:"USDBRL"`
}

type USDBRLGatewayResponse struct {
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

func DollarCotation(url string) (*DollarCotationGatewayResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var dollarCotation DollarCotationGatewayResponse
	err = json.Unmarshal(body, &dollarCotation)
	return &dollarCotation, err
}
