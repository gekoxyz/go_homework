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
	customers := make([]Customer, 10)

	var wg sync.WaitGroup
	wg.Add(len(customers))

	channel := make(chan int, len(customers))

	// Noleggi d'auto che deve gestire le prenotazioni di 10 clienti.
	for c := range customers {
		customers[c].name = "customer" + fmt.Sprintf("%d", c+1)
	}

	for _, c := range customers {
		go noleggia(c, &wg, channel)
	}

	// aspetto che tutto il gruppo di thread finisca l'esecuzione
	wg.Wait()
	close(channel)

	// stampo il risultato
	stampa(channel)
}

// function noleggia prende come input un cliente che prenota uno a caso tra i veicoli.
// stampare il cliente x che noleggia il veicolo y
func noleggia(customer Customer, wg *sync.WaitGroup, rented_channel chan int) {
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
		fmt.Println("Errore durante la generazione del veicolo random")
	}

	fmt.Println("Il cliente " + customer.name + " ha noleggiato: " + rented_veichle_str)
	// maybe has to be in a mutex
	rented_channel <- rented_veichle
	wg.Done()
}

// function stampa che stampa alla fine del processo il numero di berline, suv e station wagon noleggiati
// ogni cliente puo' noleggiare un veicolo contemporaneamente ad altri
func stampa(channel chan int) {
	rented_veichles := map[string]int{
		"berlina":       0,
		"SUV":           0,
		"station wagon": 0,
	}

	for el := range channel {
		switch el {
		case 0:
			rented_veichles["berlina"]++
		case 1:
			rented_veichles["SUV"]++
		case 2:
			rented_veichles["station wagon"]++
		}
	}
	fmt.Println("i veicoli noleggiati sono stati:")
	for veichle_type, veichle_count := range rented_veichles {
		fmt.Printf("Tipo macchina: %s, Numero noleggi: %d\n", veichle_type, veichle_count)
	}

}
