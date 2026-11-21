// 13 For statement with a for clause
// Exercise
package main

import "fmt"

func main() {
    const MAX = 0
    var loopCounter = 1

    for loopCounter < MAX {
        loopCounter++
    }
    fmt.Println("end of program")
}

// 13.2 Syntax
// for <init statement>; <cond>; <post statement> {
//              other instructions...
//               (instruction block)
// }
// Note:
// init statemnt, cond and post statemnt are optional