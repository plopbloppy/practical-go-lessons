// 13.3 Example
// control-statements/for-statement/main.go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    var ageJohn int = rand.Intn(110)
    fmt.Println("John is", ageJohn, "years old")

    for i := 0; i < ageJohn; i++ {
        fmt.Println("interation N", i)
    }
}