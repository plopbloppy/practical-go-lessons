// /hexadecimal-octal-ascii-utf8-unicode-runes/rune/main.go
package main

import "fmt"

// Create a rune using simple quotes
func main() {
    var aRune rune = 'Z'
    fmt.Printf("Unicode Code point of &#39;%c&#39;: %U\n", aRune, aRune)
}