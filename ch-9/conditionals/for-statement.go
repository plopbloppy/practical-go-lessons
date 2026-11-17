// control-statements/for-statement/main.go
package main

import "fmt"

func main() {
    const emailToSend = 3
    emailSent := 0

    for emailSent < emailToSend {
        fmt.Println("sending email...")
        emailSent++ //*\label{forWithSingle2}
    }
    fmt.Println("end of program")
}

// Notes:
// for statements are used to repeat the execution of a block of code several times
// the num of executions depends on a cond
// in this case, the code will repeat itself until emailSent is no longer less than emailToSend

// 12.3.3 Comment
// Double-check bool exp to be sure it will be false at some point in the program
// Otherwise, the program will run indefinitely
// Syntactically, the condition is optional