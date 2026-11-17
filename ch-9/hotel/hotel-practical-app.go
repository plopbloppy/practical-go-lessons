// 14 Practical applications
// 14.1 Mission
/*
We will build a hotel application.
We will begin with the front desk program used by receptionists.
At the front desk, the receptionist will welcome customers.
Clients will ask the receptionist for an available room.

He needs to know :
- How many rooms are available now
- For how long each room is available

14.1.1 Remarks
- The total number of rooms is a constant equal to 134.
- We will generate the number of rooms available randomly.
- For simplicity: the available rooms have a fake number that starts at 110
    - If three rooms are available, then the rooms are numbered: 110, 111, 112
- The number of people that you can put in a room and the number of nights available will be generated randomly
    - Minimum is 1
    - Maximum is 10
- The occupancy rate is defined as follow :
    occupancy rate=Number of rooms occupied/Total number of rooms
- The occupancy level is defined as follow :
    - From 0% to 30% occupancy rate : Low
    - From 30% to 60% occupancy rate: Medium
    - From 60% to 100% occupancy rate: High
*/

package main

import (
    "fmt"
    "math/rand"
)

func main() {
    const hotelName = "Gopher Paris Inn"
    const totalRooms = 132
    const firstRoomNumber = 110

    roomsAvail := rand.Intn(135)
    roomsOccupied := totalRooms - roomsAvail
    occupancyRate := (float32(roomsOccupied)/float32(totalRooms)) * 100

    var occupancyLevel string
    if occupancyRate > 60 {
            occupancyLevel = "High"
        } else if occupancyRate > 30 {
            occupancyLevel = "Medium"
        } else {
            occupancyLevel = "Low"
    }

    fmt.Println("Hotel :", hotelName)
    fmt.Println("Number of rooms :", totalRooms)
    fmt.Println("Rooms available :", roomsAvail)
    fmt.Println("                    Occupancy level :", occupancyLevel)
    fmt.Printf("                     Occupancy rate : %0.2f %%\n", occupancyRate)

    if roomsAvail == 0 {
        fmt.Println("No rooms available for tonight")
    } else {
        fmt.Println("Rooms :")
        for i := 0; i < roomsAvail; i++ {
            roomNumber := 110 + i
            people := rand.Intn(10) + 1
            nights := rand.Intn(10) + 1

            if nights == 1 {
                fmt.Println("-", roomNumber, ":", people, "people / 1 night")
            } else {
                fmt.Println("-", roomNumber, ":", people, "people /", nights, "nights")
            }
        }
    }
}