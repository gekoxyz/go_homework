package main

import (
	"fmt"
	"sync"
	"time"
)

// 5 torte
// 3 pasticceri
// cook -> garnish -> decorate
// primo pasticcere cucina in 1 secondo,
// ha 2 spazi per appoggiare le torte.
// secondo pasticcere guarnisce in 4 secondi
// ha 2 spazi per appoggiare le torte
// terzo pasticcere decora le torte in 8 secondi
var (
	cooked_space    = make(chan struct{}, 2)
	garnished_space = make(chan struct{}, 2)
	cooker          = make(chan struct{}, 1)
	garnisher       = make(chan struct{}, 1)
	decorator       = make(chan struct{}, 1)
	cook_time       = 1 * time.Second
	garnish_time    = 4 * time.Second
	decorate_time   = 8 * time.Second
)

func main() {
	var cake_counter int = 5

	var cakes_waitgroup sync.WaitGroup
	cakes_waitgroup.Add(cake_counter)

	for id := 0; id < cake_counter; id++ {
		go produce(id, &cakes_waitgroup)
	}
	cakes_waitgroup.Wait()
	fmt.Println("Tutte le torte sono state cucinate")
}

func produce(cake_id int, wg *sync.WaitGroup) {
	cook(cake_id)
	garnish(cake_id)
	decorate(cake_id)
	wg.Done()
}

// acquire
// channel <- struct{}{}
// release
// <-channel

func cook(cake_id int) {
	// lock the cooker
	cooker <- struct{}{}
	fmt.Printf("cucino la torta %d\n", cake_id)
	time.Sleep(cook_time)
	fmt.Printf("cucinato la torta %d\n", cake_id)
	// occupy the cooked space
	cooked_space <- struct{}{}
	// release the cooker
	<-cooker
}

func garnish(cake_id int) {
	// lock the cooker
	garnisher <- struct{}{}
	fmt.Printf("guarnisco la torta %d\n", cake_id)
	time.Sleep(garnish_time)
	fmt.Printf("guarnito la torta %d\n", cake_id)
	// free the cooked space, occupy garnished space
	<-cooked_space
	garnished_space <- struct{}{}
	// release the cooker
	<-garnisher
}

func decorate(cake_id int) {
	// lock the decorator
	decorator <- struct{}{}
	fmt.Printf("decoro la torta %d\n", cake_id)
	time.Sleep(decorate_time)
	fmt.Printf("decorato la torta %d\n", cake_id)
	// free the garnished space
	<-garnished_space
	// release the decorator
	<-decorator
}
