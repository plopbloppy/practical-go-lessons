// methods/first-example/main.go
package cart

import (
    "os/user"
    "time"

    "github.com/plopbloppy/go-money"
)

type Item struct {
    ID string
}

type Cart struct {
    ID        string
    CreatedAt time.Time
    UpdatedAt time.Time
    lockedAt  time.Time
    user.User
    Items        []Item
    CurrencyCode string
    isLocked     bool
}

func (c *Cart) TotalPrice() (*money.Money, error) {
    // ...
    return nil, nil
}

func (c *Cart) Lock() error {
    // ...
    return nil
}