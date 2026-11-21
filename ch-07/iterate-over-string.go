// /hexadecimal-octal-ascii-utf8-unicode-runes/iterate-over-string/main.go
package main

import "fmt"

func main() {
    b := "hello"
    for i := 0; i < len(b); i++ {
        fmt.Println(b[i])
    }
    // will output :
    // 104
    // 101
    // 108
    // 108
    // 111
    // and NOT :
    // h
    // e
    // l
    // l
    // o
}