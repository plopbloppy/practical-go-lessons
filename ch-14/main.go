package main

import (
    "call/cart"
    "log"
)

func main() {
    // load the cart into variable cart
    newCart := cart.Cart{}

    // the method TotalPrice is called
    // the variable cart is passed to the method TotalPrice
    // bound to type Cart from the current pkg (main)
    totalPrice, err := newCart.TotalPrice()
    if err != nil {
        log.Printf("impossible to compute price of the cart: %s", err)
        return
    }
    // method Display is bound to type Money from the pkg money
    log.Println("Total Price", totalPrice.Display())

    // the method Lock is called
    // the variable cart is passed to the method Lock
    // bound to type Cart from the current pkg (main)
    err = newCart.Lock()
    if err != nil {
        log.Printf("impossible to lock the cart: %s", err)
        return
    }
}