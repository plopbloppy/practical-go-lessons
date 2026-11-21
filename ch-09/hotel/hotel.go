// control-statements/first-statements/main.go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    hotelName := "The Gopher Hotel"
    fmt.Println("Hotel " + hotelName)
    var roomsAvailable int //we use int here instead of uint8 b/c rand.Intn generates an int
    var rooms, roomsOccupied int = 100, rand.Intn(100)
    roomsAvailable = rooms - roomsOccupied //*\label{cond1diff}
    fmt.Println(roomsAvailable, "rooms available")
}