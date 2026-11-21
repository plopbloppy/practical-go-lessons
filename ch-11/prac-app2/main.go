// package-imports/application-refactor/problem/main.go
package main

import (
    prac-app2/email
    prac-app2/invoice
)

func main() {
    // first reservation
    customerName := "Doe"
    customerEmail := "john.doe@example.com"
    var nights uint = 12
    emailContents := email.Contents("M", customerName, nights)
    email.Send(emailContents, customerEmail)
    invoice.Create(customerName, nights, 145.32)
}