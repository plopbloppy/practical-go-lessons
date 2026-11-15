// 14 Runes
package main

import "fmt"

func main() {
    s := "我爱 Golang" // I love Golang
    for _, v := range s { // iterates over each char of the string
        // v is of type rune
        fmt.Printf("Unicode code point : %U - character '%c' - binary %b - hex %X - Decimal %d\n ", v, v, v, v, v)
    }
}

// Notes:
// rune is a built-in type
// Defn: rune is an alias for int32 and is equiv. to int32 in all ways.
// It is used, by convention, to distinguish char vals from integer vals.
type rune = int32

// %U -> prints code point in conventional format

// Output:
// Unicode code point : U+6211 - character '我' - binary 110001000010001 - hex 6211 - Decimal 25105
//  Unicode code point : U+7231 - character '爱' - binary 111001000110001 - hex 7231 - Decimal 29233
//  Unicode code point : U+0020 - character ' ' - binary 100000 - hex 20 - Decimal 32
//  Unicode code point : U+0047 - character 'G' - binary 1000111 - hex 47 - Decimal 71
//  Unicode code point : U+006F - character 'o' - binary 1101111 - hex 6F - Decimal 111
//  Unicode code point : U+006C - character 'l' - binary 1101100 - hex 6C - Decimal 108
//  Unicode code point : U+0061 - character 'a' - binary 1100001 - hex 61 - Decimal 97
//  Unicode code point : U+006E - character 'n' - binary 1101110 - hex 6E - Decimal 110
//  Unicode code point : U+0067 - character 'g' - binary 1100111 - hex 67 - Decimal 103