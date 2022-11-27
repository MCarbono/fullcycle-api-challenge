package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const FILENAME = "cotacao.txt"

func main() {
	dollarCotation, err := request()
	if err != nil {
		panic(err)
	}
	err = writeDollarCotation(dollarCotation.Bid, FILENAME)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Dollar cotation successfully get and saved in the file %v", FILENAME)
}

func writeDollarCotation(text string, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte("Dólar: " + text))
	return err
}

func request() (*DollarCotationHTTPResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(data))
	}
	var dollarCotation DollarCotationHTTPResponse
	err = json.Unmarshal(data, &dollarCotation)
	if err != nil {
		return nil, err
	}
	return &dollarCotation, nil
}

type DollarCotationHTTPResponse struct {
	Bid string `json:"bid"`
}
