// package-imports/import-declaration/room/room.go
package room

import "fmt"

// display information about a room
func PrintDetails(roomNumber, size, nights int) {
    fmt.Println(roomNumber, ":", size, "people /", nights, " nights ")
}

// 5 File organization
/*
Notes:
- src files of a pkg must be grouped into a single dir
- the dir must have the same name as the pkg
- pkg names are in camelCase and should be ≤ 10 letters in length

e.g. src files of baz pkg must be stored into baz folder
foo (directory)
|
<package foo> <package foo> <package foo> (src files)

bar (directory)
<package bar> <package bar> <package bar> (src files)

13 Linking pkg tgt: a trial and error approach
- to export a fcn, we must capitalize the first letter of the fcn
- a fcn, a var, and a const if not exported is private to the pkg where it is defined

// package-imports/package-example/price/price.go
package price

Ex:
// this fcn is not exported and cannot be used in another pkg
func compute(){
    //
}

// this fcn is exported and  can be used in another pjg
func Update(){
    //...
}

- public: exported
- private: unexported
- visibility: the exported/unexported quality of a method
*/