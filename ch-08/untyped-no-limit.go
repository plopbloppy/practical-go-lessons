// variables-constants-and-basic-types/untyped-no-limit/main.go
package main

func main() {
    // max val of an int is 9223372036854775807
    // 9223372036854775808 (max + 1 ) overflows int
    const profit = 9223372036854775808
    // the program compiles
    // even though it overflows 64-bits!
    // (b/c it has no type it does not depend on any type restriction)

    // this will not compile if we try to assign this const to a typed var
    // var profit2 int64 = profit
    // fmt.Println(profit2)
    // compilation error!
}