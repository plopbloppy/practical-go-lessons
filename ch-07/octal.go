// /hexadecimal-octal-ascii-utf8-unicode-runes/octal/main.go
package main

import "fmt"

func main() {
    n2 := 0x9F4
    fmt.Printf("Decimal : %d\n", n2) //Output : Decimal : 2548

    // n3 is represented using the octal numeral system
    n3 := 02454
    // alternative : n3 := 0o2454

    // convert in decimal
    fmt.Printf("decimal: %d\n", n3) //Output : decimal: 1324

    // n4 is represented using the deci system
    n4 := 1324
    // output n4 (deci) in octal
    fmt.Printf("octal: %o\n", n4) //Output : octal: 2454
    // output n4 (deci) in octal (with a 0o prefix)
    fmt.Printf("octal with prefix: %0\n", n4) //Output : octal with prefix : 00o2454
}