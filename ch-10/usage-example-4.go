// functions/usage-example-4/main.go
/*
12 0 param, 0 result
func <fcn name> () {
    fcn body
}
*/
package main

import (
    "fmt"
)

func main() {
    printHeader()
}

func printHeader() {
    fmt.Println("Hotel Golang")
    fmt.Println("San Francisco, CA")
}