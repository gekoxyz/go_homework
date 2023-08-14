package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// struttura cliente con campo nome
type Customer struct {
	name string
}

// struttura veicolo con capo tipo
type Veichle struct {
	category int
}

// Veicoli disponibili: Berlina, SUV, Station Wagon
const (
	berlina       = iota // 0
	suv                  // 1
	station_wagon        // 2
)

func main() {
	var customers [10]Customer

	var wg sync.WaitGroup
	wg.Add(len(customers))

	// Noleggi d'auto che deve gestire le prenotazioni di 10 clienti.
	for c := range customers {
		customers[c].name = "customer" + fmt.Sprintf("%d", c+1)
		fmt.Println("creato: " + customers[c].name)
	}

	for _, c := range customers {
		go noleggia(c, &wg)
	}

	// aspetto che tutto il gruppo di thread finisca l'esecuzione
	wg.Wait()

	// stampo il risultato
	// stampa(result_channel)
}

// function noleggia prende come input un cliente che prenota uno a caso tra i veicoli.
// stampare il cliente x che noleggia il veicolo y
func noleggia(customer Customer, wg *sync.WaitGroup) {
	rented_veichle := rand.Intn(3)
	rented_veichle_str := ""

	switch rented_veichle {
	case berlina:
		rented_veichle_str = "berlina"
	case suv:
		rented_veichle_str = "SUV"
	case station_wagon:
		rented_veichle_str = "station wagon"
	default:
		fmt.Println("Error while generating the random veichle")
	}

	fmt.Println("Il cliente " + customer.name + " ha noleggiato: " + rented_veichle_str)

	wg.Done()
}

// function stampa che stampa alla fine del processo il numero di berline, suv e station wagon noleggiati
// ogni cliente puo' noleggiare un veicolo contemporaneamente ad altri
func stampa(result_channel chan int) {
	rented_veichles := make(map[string]int)

	rented_veichles["berlina"] = 0
	rented_veichles["SUV"] = 0
	rented_veichles["station wagon"] = 0

	for v := range result_channel {
		switch v {
		case berlina:
			rented_veichles["berlina"]++
		case suv:
			rented_veichles["SUV"]++
		case station_wagon:
			rented_veichles["station wagon"]++
		}
	}

	for el := range rented_veichles {
		fmt.Println(el)
	}
}
