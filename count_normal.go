package main

import (
    "fmt"
)

func main() {
    var str string = "aaaaaaaaaaaaabbbbbbbbcccccddddccccccfff"
    var char string = "c"

    var count int = 0

    for _, to_check := range str {
        if string(to_check) == char {
            count++
        }
    }
    fmt.Printf("The character %s appears %d times in the string \n", char, count)
}

