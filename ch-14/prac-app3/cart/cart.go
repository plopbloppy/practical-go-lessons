package cart

import (
    "methods/example-project/product"
    "os/user"
    "time"

    "github.com/Rhymond/go-money"
)

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

type Item struct {
    product.Product
    Quantity uint8
}

// compute the total price of the cart and return it
// the total price of the cart is for each article
// the sum of the quantity * the unit price
func (c *Cart) TotalPrice() (*money.Money, error) {
    total := money.New(0, c.CurrencyCode)
    var err error
    for _, v := range c.Items {
        itemSubtotal := v.Product.Price.Multiply(int64(v.Quantity))
        total, err = total.Add(itemSubtotal)
        if err != nil {
            return nil, err
        }
    }
    return total, nil
}

// lock the cart to avoid modifications after confirmation
// update the field isLocked to true
// update the field lockedAt to the current time
// if the cart is already locked, it should return an error
func (c *Cart) Lock() error {
    if cart.isLocked {
        return errors.New("cart is already locked")
    }
    c.isLocked = true
    c.lockedAt = time.Now()
    return nil
}

func (c *Cart) delete() error {
    // to implement
    return nil
}

// unit test fcn
func TestTotalPrice(t *testing.T) {
    items := []Item{
        {
            Product: product.Product{
                ID:    "p-1254",
                Name:  "Product test",
                Price: money.New(1000, "EUR"),
            },
            Quantity: 2,
        },
        {
            Product: product.Product{
                ID:    "p-1255",
                Name:  "Product test 2",
                Price: money.New(2000, "EUR"),
            },
            Quantity: 1,
        },
    }
    c := Cart{
        ID:           "1254",
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
        User:         user.User{},
        Items:        items,
        CurrencyCode: "EUR",
    }
    actual, err := c.TotalPrice()
    assert.NoError(t, err)
    assert.Equal(t, money.New(4000, "EUR"), actual)
}

/*
- We begin by creating a slice of 2 fake Item elements:
    - The ID and the Name fields are filled with test data
    - To create a price, simply call money.New(1000, "EUR")
    - The two items have different prices
- A new variable c of type Cart is created
- The field Items is set to the value items
- The field CurrencyCode is set to "EUR" (euro currency code)
- The method is then called on c
- Check there is no error with assert.NoError(t, err)
- Verify that the actual value of the first result is equal to
    money.New(4000, "EUR") (40 Euros)
    - 2 * 10 EUR + 1 * 20 EUR = 40 EUR
*/

// unit tests for Lock

// test that when a cart is locked:
// the field isLocked = true
// the field lockedAt is set with a time corresponding to a UNIX epoch > 0
// this is not optimal, a better test would have checked that the time is correctly set
func TestLock(t *testing.T) {
    c := Cart{
        ID: "1254",
    }
    err := c.Lock()
    assert.NoError(t, err)
    assert.True(t, c.isLocked)
    assert.True(t, c.lockedAt.Unix() > 0)
}

// test the lock fcn when the cart is already locked
// set isLocked to true
// then call Lock to check that an error occured
func TestLockAlreadyLocked(t *testing.T) {
    c := Cart{
        ID:       "1254",
        isLocked: true,
    }
    err := c.Lock()
    assert.Error(t, err)
}