// 14 Practical Application
// 14.1 Mission
/*
Refactor this code and create functions that will :

    Compute the occupancy level

    Compute the occupancy rate

    Print the details of a specific room.
*/
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    const hotelName = "Gopher Paris Inn"
    const totalRooms = 134
    const firstRoomNumber = 110

    rand.Seed(time.Now().UTC().UnixNano())
    roomsOccupied := rand.Intn(totalRooms)
    roomsAvailable := totalRooms - roomsOccupied

    fmt.Println("Hotel:", hotelName)
    fmt.Println("Number of rooms", totalRooms)
    fmt.Println("Rooms available", roomsAvailable)
    fmt.Println("                  Occupancy Level:", occupancyLevel(occupancyRate(roomsOccupied, totalRooms)))
    fmt.Printf("                  Occupancy Rate: %0.2f %%\n", occupancyRate(roomsOccupied, totalRooms))

    if roomsAvailable > 0 {
        fmt.Println("Rooms:")
        for i := 0; roomsAvailable > i; i++ {
            roomNumber := firstRoomNumber + i
            size := rand.Intn(6) + 1
            nights := rand.Intn(10) + 1
            printRoomDetailsV2(roomNumber, size, nights)
        }
    } else {
        fmt.Println("No rooms available for tonight")
    }
}

// consume an occupancyRate from the fcn occupancyRate
// For 0% ≤ occupancyRate ≤ 20% return "Low"
// For 20% < occupancyRate ≤ 70% return "Medium"
// For 70% < occupancyRate return "High"
func occupancyLevel(occupancyRate float32) string{
    if occupancyRate > 70 {
        return "High"
    } else if occupancyRate > 20 {
        return "Medium"
    } else {
        return "Low"
    }
}

// compute hotel occupancy rate
// return a percentage from 0% to 100%
func occupancyRate(roomsOccupied, totalRooms int) float32 {
    return (float32(roomsOccupied) / float32(totalRooms)) * 100
}

// display info about a room
func printRoomDetails(roomNumber int, size int, nights int) {
    fmt.Println(roomNumber, ":", size, "people /", nights, " nights ")
}

/*
14.3.1.1 Challenge: plural vs. singular
1. Write a second version of fcn printRoomDetails that corrects the plural singular bug
2. The input param types are questionable. Why?
-> printRoomDetails takes in type int for vals that should only be +ive
*/

// display info about a room
// if the room is for one night, display # people / 1 night
// otherwise, display # people / # nights
func printRoomDetailsV2(roomNumber int, size int, nights int) {
    nightText := "nights"
    if (nights == 1) {
        nightText = "night"
    }
    fmt.Println(roomNumber, ":", size, "people /", nights, nightText)
}

/*
14.3.2.1 Challenge: give me your better switch case
Write a second version of the function occupancyLevel which use a switch/case instead of an if / else if / else structure.
*/

// consume an occupancyRate from fcn occupancyRate
// For 0% ≤ occupancyRate ≤ 20% return "Low"
// For 20% < occupancyRate ≤ 70% return "Medium"
// For 70% < occupancyRate return "High"
func occupancyLevelV2(occupancyRate float32) string {
    switch {
    case occupancyRate > 70:
        return "High"
    case occupancyRate > 20:
        return "Medium"
    default:
        return "Low"
    }
}