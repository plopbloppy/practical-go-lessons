// variables-constants-and-basic-types/declaration-variant/main.go
package main

import "fmt"

func main() {
    var roomNumber, floorNumber int = 154, 3
    fmt.Println(roomNumber, floorNumber)

    var password = "notSecured"
    // type was not initialized, so Go gives password the type of the initialized value
    // "notSecured" is a string; thus type of the var password is string

    fmt.Println(password)

    // Output:
    // 154 3
    // notSecured
}