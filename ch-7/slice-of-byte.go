// /hexadecimal-octal-ascii-utf8-unicode-runes/slice-of-byte/main.go
package main

import "fmt"

func main() {
    b := make([]byte, 0)
    // we created a slice of bytes (with the built-in fcn make)

    // golang byte type is an alias of uint8
    // Defn: uint8 means we can store unsigned (non-negative)
    // integers on 8 bits (a byte) of data
    // min val is 0, max val is 255 (2^7 + ... + 2^0)

    b = append(b, 255)
    b = append(b, 10)
    // max val is y we ap pend to a byte slice numbers 0 to 255
    // if we try numbers greater than 255, we will get an error:
    // ... constant 256 overflows byte

    // %b formatting verb to print binary rep of a num
    n2 := 0x9F4
    fmt.Printf("Decimal : %d\n", n2) //Output: Decimal : 2548
    fmt.Printf("Binary : %b\n", n2) //Binary : 100111110100

    fmt.Println(b)
}