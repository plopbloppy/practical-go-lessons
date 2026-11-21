// variables-constants-and-basic-types/untyped-const/main.go
package main

import "fmt"

func main() {
    const occupancyLimit = 12
    // occupancyLimit has no particular type, except it has been set to an int val

    var occupancyLimit1 uint8
    var occupancyLimit2 int64
    var occupancyLimit3 float32

    // assign untyped const to an uint8 var
    occupancyLimit1 = occupancyLimit
    // assign untyped const to an int64 var
    occupancyLimit2 = occupancyLimit
    // assign untyped const to a float32 var
    occupancyLimit3 = occupancyLimit

    fmt.Println(occupancyLimit1, occupancyLimit2, occupancyLimit3)
}

// Notes:
// the program compiles!
// Output:
// 12 12 12
// the const val can be put into vars of diff types
