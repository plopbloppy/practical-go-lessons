// package-imports/booking/booking.go
/*
9 Practical application 1
9.1 Problem

Add a package named booking to the previous Go program. In this package :
- Define a constant named vatRate with the value 20.0
- Define a function named printVatRate that will print to the standard output
the constant vatRate followed by the percentage sign.
*/

package booking

import "fmt"

const vatRate = 20.0

func printVatRate() {
    fmt.Println("%.2f %%", vatRate)
}