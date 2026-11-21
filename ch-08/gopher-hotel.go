// 12 Practical Application
// 12.1 Mission
// 12.1.1 Program

/*
Write a program that will :

    Create a string constant named hotelName and with value "Gopher Hotel"

    Create two untyped constants that will contain respectively 24.806078 and -78.243027.
    The names of those two constants are longitude and latitude. (somewhere in the Bahamas)

    Create a variable of type int named occupancy which is initialized with the value 12.

    Print the hotelName the longitude, the latitude and the occupancy on a new line
*/

package main

import "fmt"

func main() {
    const hotelName string = "Gopher Hotel"
    const longitude = 24.806078 // default type : float64, type : no type
    const latitude = -78.243027 // type : no type
    var occupancy int = 12 // 32-bits/64-bits int depending on machine

    fmt.Println(hotelName, longitude, latitude)
    fmt.Println(occupancy)
}