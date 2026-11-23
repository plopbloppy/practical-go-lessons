// types/composite/main.go
package main

import "fmt"

func main() {
    // array constructed with the basic type uint8
    var arr [3]uint8

    // pointer constructed with the basic type uint8
    var myPointer *uint8

    // function  constructed with the basic type string
    var nameDisplayer func(name, firstname string) string

    // slices constructed with the basic type uint8
    var roomNumbers []uint8

    // maps constructed with the basic types uint8 and string
    var score map[string]uint8

    // channel constructed with the basic type bool
    var received chan<- bool

    fmt.Println(arr, myPointer, nameDisplayer, roomNumbers, score, received)
}

// Note
// [3]uint8, *uint8, func(name, firstname string) string,... are type literals
// composite types are constructed with type literals