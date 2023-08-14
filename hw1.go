package main

import (
    "fmt"
    "sync"
)

func main() {
    var str string = "aaaaaaaaaaaaabbbbbbbbcccccddddccccccfff"
    var to_search string = "c"
    
    // inizializzo il count in un buffered channel di capienza 1 per evitare la data race
    count := make(chan int, 1)
    count <- 0

    // inizializzo il waitgroup per aspettare che il gruppo di thread finisca l'esecuzione prima di passare oltre
    var wg sync.WaitGroup
    wg.Add(len(str))

    for _, to_check := range str {
        // per ogni carattere da controllare avvio una goroutine
        go checkChar(to_search, to_check, count, &wg)
    }
    
    // metto in pausa il thread finche' tutti i thread non hanno raggiunto il punto di sincronizzazione
    wg.Wait()
    final_number := <- count

    fmt.Printf("The character %s appears %d times in the string \n", to_search, final_number)
}

func checkChar(to_search string, to_check rune, count chan int, wg *sync.WaitGroup) {
	if to_search == string(to_check) {
	    num := <- count
		num++
	    count <- num
	}
    wg.Done() // comunico che il thread e' arrivato al punto di sincronizzazione
}

