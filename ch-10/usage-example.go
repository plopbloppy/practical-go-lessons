// functions/usage-example/main.go
package main

import "fmt"

func main() {
    johnPrice := computePrice(145.90, 3) //*\label{funcEx1Us1}
    paulPrice := computePrice(26.32, 10) //*\label{funcEx1Us2}
    robPrice := computePrice(189.21, 2) //*\label{funcEx1Us3}

    total := johnPrice + paulPrice + robPrice
    fmt.Printf("TOTAL : %0.2f $", total)
}

func computePrice(rate float32, nights int) float32 { //*\label{funcEx1Dec}
    return rate * float32(nights)
}

/*
9 Scope of a fcn
A fcn's scope begins with { and stops with }
A scope is a portion and an extent of the source code.

Ex: rate and nights can be used freely inside the scope of computePrice,
but do not exist outside the curly braces.

If you attempt to use a variable defined within a fcn's scope,
this error occurs when attempting to compile the program:
=> undefined: rate
*/