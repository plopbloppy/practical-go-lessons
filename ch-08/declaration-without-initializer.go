// variables-constants-and-basic-types/declaration-without-initializer/main.go
package main

import "fmt"

func main() {
    var roomNumber, floorNumber int
    fmt.Println(roomNumber, floorNumber)
    // Zero value of type int is 0

    var password string
    fmt.Println(password)
    // Zero value of type string is "" (empty string)

     // Output:
     // 0 0
}