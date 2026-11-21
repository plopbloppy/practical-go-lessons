// functions/scope2/main.go
package main

import "fmt"

func main() {
    fmt.Println(computePrice(145.90, 3))
}

func computePrice(rate float32, nights int) float32 {
    n := float32(nights)
    return rate * n
}