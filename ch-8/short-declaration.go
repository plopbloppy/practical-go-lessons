// variables-constants-and-basic-types/short-declaration/main.go
package main

import "fmt"

func main() {
    roomNumber, floorNumber := 154, 3
    fmt.Println(roomNumber. floorNumber)
}

// Warning : short var declaration cannot be used outside fcns!
//     ""  : cannot use the value nil for short var declaration
//           (the compiler cannot infer the type of the var)