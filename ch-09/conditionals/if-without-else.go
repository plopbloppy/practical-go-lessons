// control-statements/if-without-else/main.go
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
    }
    fmt.Println("End of the Program")
}