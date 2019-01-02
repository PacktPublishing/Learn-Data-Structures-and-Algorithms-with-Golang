//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt,database/sql, net/http, text/template package
import (
//    "fmt"
    "database/sql"
  //  "log"
  //  "net/http"
  //  "text/template"
  //  "errors"
    _ "github.com/go-sql-driver/mysql"
)

type Customer struct {
    CustomerId    int
    CustomerName  string
    SSN string
}

func GetConnection() (database *sql.DB) {
    databaseDriver := "mysql"
    databaseUser := "newuser"
    databasePass := "newuser"
    databaseName := "crm"
    database, error := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@/"+databaseName)
    if error != nil {
        panic(error.Error())
    }
    return database
}


func GetCustomersFromDB() []Customer {
    var database *sql.DB
    database = GetConnection()

    var error error
    var rows *sql.Rows
    rows, error = database.Query("SELECT * FROM Customer ORDER BY Customerid DESC")
    if error != nil {
        panic(error.Error())
    }
    var customer Customer
    customer = Customer{}

    var customers []Customer
    customers= []Customer{}
    for rows.Next() {
        var customerId int
        var customerName string
        var ssn string
        error = rows.Scan(&customerId, &customerName, &ssn)
        if error != nil {
            panic(error.Error())
        }
        customer.CustomerId = customerId
        customer.CustomerName = customerName
        customer.SSN = ssn
        customers = append(customers, customer)
    }

    defer database.Close()

    return customers
}
