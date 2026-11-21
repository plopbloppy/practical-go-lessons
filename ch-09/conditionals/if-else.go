// control-statements/if-else/main.go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    var agePaul, ageJohn int = rand.Intn(110), rand.Intn(110)

    fmt.Println("John is", ageJohn, "years old")
    fmt.Println("Paul is", agePaul, "years old")

    if agePaul > ageJohn {
        fmt.Println("Paul is older than John")
    } else {
        fmt.Println("Paul is younger than John, or both have the same age")
    }
}