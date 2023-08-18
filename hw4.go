package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

type CurrencyPair struct {
	ticker string
	price  float64
}

func main() {

	var wg sync.WaitGroup

	currency_pairs := make(chan CurrencyPair)

	go simulateMarketData(currency_pairs, &wg)
	go selectPair(currency_pairs, &wg)

	// Simulare il ciclo di trading per 1 minuto e poi terminare il programma
	time.Sleep(60 * time.Second)
	defer os.Exit(0)
}

// Funzione selectPair che gestisce le operazioni di vendita e acquisto
// Se EUR/USD supera 1.20 vende EUR/USD. 4 secondi di delay prima della vendita
// Se GPB/USD scende sotto 1.35 acquista GBP/USD. 3 secondi di delay
// Se JPY/USD scende sotto 0.0085 acquista JPY/USD. 3 secondi di delay
func selectPair(currency_pairs chan CurrencyPair, wg *sync.WaitGroup) {
	wg.Wait()
	for currency_pair := range currency_pairs {
		switch currency_pair.ticker {
		case "EUR/USD":
			if currency_pair.price > 1.2 {
				fmt.Printf("EUR/USD: buying @ %.2f\n", currency_pair.price)
				time.Sleep(4 * time.Second)
				fmt.Printf("EUR/USD: bought successfully @ %.2f\n", currency_pair.price)
			}
		case "GBP/USD":
			if currency_pair.price < 1.35 {
				fmt.Printf("GBP/USD: buying @ %.2f\n", currency_pair.price)
				time.Sleep(3 * time.Second)
				fmt.Printf("GBP/USD: bought successfully @ %.2f\n", currency_pair.price)
			}
		case "JPY/USD":
			if currency_pair.price < 0.0085 {
				fmt.Printf("JPY/USD: buying @ %.4f\n", currency_pair.price)
				time.Sleep(3 * time.Second)
				fmt.Printf("JPY/USD: bought successfully @ %.4f\n", currency_pair.price)
			}
		}

	}
	defer wg.Done()
}

// Funzione simulateMarketData che simuli il prezzo delle coppie di valute e invii i dati simulati su un canale
// EUR/USD 1.0/1.5
// GBP/USD 1.0/1.5
// JPY/USD 0.006/0.009
// I prezzi vengono generati e inviati ogni secondo
func simulateMarketData(currency_pairs chan CurrencyPair, wg *sync.WaitGroup) {
	for {
		wg.Wait()
		var eur_usd_price float64 = 1.0 + rand.Float64()*0.5
		var gbp_usd_price float64 = 1.0 + rand.Float64()*0.5
		var jpy_usd_price float64 = 0.006 + rand.Float64()*0.003

		currency_pairs <- CurrencyPair{"EUR/USD", eur_usd_price}
		currency_pairs <- CurrencyPair{"GBP/USD", gbp_usd_price}
		currency_pairs <- CurrencyPair{"JPY/USD", jpy_usd_price}

		time.Sleep(1 * time.Second)
		defer wg.Done()
	}
}
