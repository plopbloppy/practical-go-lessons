// functions/declaration-example/main.go
package main

// compute the price of a room based on its rate per night
// and the number of nights
func computePrice(rate float32, nights int) float32 {
    return rate * float32(nights)
}

// A std fcn: 2 params, 1 results
// func <fcn name (identifier)> (<param1 name> <param1 type>, <param2 name> <param2 type>) <result1 type> {
//              fcn body
// }

/*
Fcn details:
identifier: computePrice
signature: (rate float32, nights int) float 32
1st param: rate (type float32)
2nd param: nights (type int)
1 result: (type float32)
fcn body: return rate * float32(nights)
\-> return: indicates what follows after is the result

Fcn purpose:
- identifier, comments and signature gives info
*/