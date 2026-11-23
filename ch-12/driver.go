// https://github.com/go-sql-driver/mysql/blob/master/driver.go
package MySQL

import (
   //...
   "database/sql"
   //...
)
// ...


func init() {
   sql.Register("mysql", &MySQLDriver{}) //Register is an exported fcn from database/sql std lib pkg
}

/*
Register fcn:

// Register makes a database driver available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, driver driver.Driver) {
   //...
}
*/