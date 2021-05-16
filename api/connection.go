package api

import (
	"fmt"
	"log"
	"net/http"
)

func EstablishConnection() {
	resp, err := http.Get("https://api.binance.com")
	if err != nil {
		log.Fatal("Http request failed", err)
	}
	fmt.Println(resp)
}
