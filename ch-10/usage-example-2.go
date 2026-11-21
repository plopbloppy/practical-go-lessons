// functions/usage-example-2/main.go
/*
10 2 params, 1 named result
func <fcn name> (<param1 name> <param1 type>, <param2 name> <param2 type>) (<result1 name> <result1 type>) {
    fcn body
}

Ex:
func computePrice(rate float32, nights int) (price float32) {
   price = rate * float32(nights)
   return
}

Notes:
- the parentheses around the result name and result type are mandatory
- price exists already, so it does not need to be declared
- it was initialized in the signature with type32 and value 0
- not necessary to write return price as the result was already named;
  when called, the fcn will return current val of price
- usage syntax does not chge
*/

package main

import "fmt"

func main() {
    johnPrice := computePrice(145.90, 3)
    fmt.Printf("TOTAL : %0.2f $", johnPrice)
}

// compute the price with a 200% margin
func computePrice(rate float32, nights int) (price float32) {
    p := rate * float32(nights)
    p = p * 2
    return // returns 0 b/c a value was never assigned to the named result, price
}

