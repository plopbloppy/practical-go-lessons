// hexadecimal-octal-ascii-utf8-unicode-runes/hexa-lower/main.go
package main

import "fmt"

func main() {
    n := 2548
    fmt.Printf("%x", n) //output: 9f4
    // %x is the formatting verb for hexadecimal (with letters lowercase)
    // Note: n is a number denoted using the deci system

    // %X is used to print hexadecimal num w capitalized letters
    fmt.Printf("%X", n) //output: 9F4

    // represent a number in hexadecimal, add 0x b4 numeral:
    n2 := 0x9F4
    fmt.Printf("%X\n", n) //output: 9F4
    fmt.Printf("%x\n", n2) //output 9f4

    // to print the number in base ten, use the verb "%d"
    fmt.Printf("Decimal : %d\n", n2) //output: 2548
}