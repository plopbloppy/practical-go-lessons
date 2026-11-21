// functions/usage-example-3/main.go
/*
11 0 params, 1 result
func <fcn name> () <result1 type> {
    fcn body
}
*/
package main

import (
    "fmt"
    "math/rand"
)

func main() {

    vacant := vacantRooms()
    fmt.Println("Vacant rooms:", vacant)

}

func vacantRooms() int {
    return rand.Intn(100)
}

