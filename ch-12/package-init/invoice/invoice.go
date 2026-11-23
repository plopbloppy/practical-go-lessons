// package-init/rules-illustration/invoice/invoice.go
package invoice

import (
    "fmt"

    "pkgInit/rules/currency"
)

func init() {
    fmt.Println("invoice init")
}

func Print() {
    fmt.Println("INVOICE Number X")
    fmt.Println(54, currency.EuroSymbol())
}