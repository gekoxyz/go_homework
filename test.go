package main

import (
	"fmt"
    "time"
)

func main() {
    c1 := make(chan int, 1)
    c1 <- 0

    go increaseNumber(c1)
    go increaseNumber(c1)

    time.Sleep(time.Second)
    num2 := <- c1
    fmt.Print(num2)
}

func increaseNumber(ic chan int) {
    num := <- ic
    num++
    ic <- num
}

