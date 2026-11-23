// package-init/rules-illustration-advanced/main.go
package main

import (
    "fmt"

    "pkgInit/rules/invoice"
)

func init() {
    fmt.Println("main")
}

func main() {
    fmt.Println("--program start--")
    invoice.Print()
}