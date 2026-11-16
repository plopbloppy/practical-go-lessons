// variables-constants-and-basic-types/default-type/main.go
package main

import "fmt"

func main() {
    const occupancyLimit = -12

    // Case 1 :
    // var occupancyLimit4 string
    // this will cause an error as we attempt to convert an int -> string

    var occupancyLimit4 int8
    // this is ok because an int can be converted into int8!

    // Case 2 :
    // var occupancyLimit4 uint8 (attempt to assign unsigned int to an int)
    // Error: cannot use occupancyLimit (untyped int constant -12) as uint8 value in assignment (overflows)

    occupancyLimit4 = occupancyLimit

    fmt.Println(occupancyLimit4)
}