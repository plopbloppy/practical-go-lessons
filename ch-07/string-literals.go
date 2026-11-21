// /hexadecimal-octal-ascii-utf8-unicode-runes/string-literals/main.go
package main

import "fmt"

func main() {

    raw := `spring rain:
browsing under an umbrella
at the picture-book store`
    fmt.Println(raw)

    interpreted := "i love spring"
    fmt.Println(interpreted)

    // Note: string literals are implicitly encoded using UTF-8
}