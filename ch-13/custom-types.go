// 6 How to create new types: type defn
// 6.1 New types based on predeclared types

// a new type Firstname
// underlying type : string
type Firstname string

// underlying type is string
type Currency string

// underlying type is float64
type VATRate float64

type Country string

// Notes:
// underlying type: type from another type

// 6.2 New types based on composite types
// declare a new type with an underlying composite type

// new type "ExchangeRate"
// underlying type is map[string]float64
// map[string]float64 is a type literal
type ExchangeRate map[string]float64

// new type "Birthdate"
// underlying type : time.Time (type Time from the time pkg)
type Birthdate time.Time

// new type "Hotel"
// underlying type : struct
type Hotel struct {
    Name     string
    Capacity uint8
    Rooms    uint8
    Smoking  bool
}

// new type "Country"
// underlying type : struct
type Country struct {
    Name        string
    CapitalCity string
}

// 7 Type struct variable creation
// 7.1 With fields names
france := Country{
    Name:        "France",
    CapitalCity: "Paris",
}

// specify certain fields:
usa := Country {
    Name: "United States of America"
}
// other fields will be equal to the zero value of the type of the field
// CapitalCity will be equal to ""

// We can create a new element of new Country, w/out specifying any fields val:
empty := Country{}

// 7.2 Without fields names
belgium := Country {
    "Belgium",
    "Bruxelles",
}
// We create a value of type Country and we set the field Name with "Belgium"
// and the field CapitalCity with "Bruxelles"
// this syntax should be used carefully

// 7.2.0.1 Field values need to be specified in the same order as in the type struct declaration
/*
type Country struct {
    Name        string
    CapitalCity string
}
*/

japan := Country{
    "Tokyo", //Country.Name will be "Tokyo"
    "Japan",
}

/*
// 7.2.0.2 You cannot skip a field
china := Country{
    "China",
}
will not compile

7.2.0.3 Do not mix the two syntaxes!
greece := Country{
    Name: "Greece",
    "Athens",
}
*/

// 8 How to access a field: selector expression
/*
fr := Country{Name: "France"}
fr.Name (field identifier)
*/

usa.CapitalCity = "Washington DC"

if usa.Name == "France" {
    fmt.Println("we have an error !")
}

// 9 Embedded fields
/*
In a type struct, we can add embedded fields.
Embedded fields are defined implicitly.

type Hotel struct {
    Name     string
    Capacity uint8
    Rooms    uint8
    Smoking  bool
    Country <- embedded field (another type struct)
               Type: Country
               Field Name: Country
}

type Country struct {
    Name        string
    CapitalCity string
}

- Embedded fields have no explicit name. The field name is the type name.
*/

// 9.1 Pointer type
/* embed a pointer type into a type struct:

type Hotel struct {
    Name string
    *Country <- pointer to an element of type Country
                field name is type name: Country
}

type Country struct {
    Name        string
    CapitalCity string
}

hotel := Hotel{
    Name:    "Hotel super luxe",
    Country: &Country{Name: "France"},
}

fmt.Println(hotel.Country.Name)
*/
