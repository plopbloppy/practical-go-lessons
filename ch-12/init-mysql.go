// package-init/init-mysql/main.go
package main

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql" //blank import, the pkg init fcn will be run
)

func main() {
    db, err := sql.Open("mysql", "user:password@/dbname") //Open opens a connection to the db
    if err != nil {
        panic(err)
    }
    log.Println(db)
    //...
}

// Note:
// the author of the driver pkg could also remove the init fcn and
// instruct the user to call the Register fcn manually