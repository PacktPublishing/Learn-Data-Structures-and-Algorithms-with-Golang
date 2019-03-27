//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt,database/sql, net/http, text/template package
import (
	"database/sql"
	"fmt"
	//  "log"
	//  "net/http"
	//  "text/template"
	//  "errors"
	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
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

func GetCustomers() []Customer {
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
	customers = []Customer{}
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

func InsertCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()

	var error error
	var insert *sql.Stmt
	insert, error = database.Prepare("INSERT INTO CUSTOMER(CustomerName,SSN) VALUES(?,?)")
	if error != nil {
		panic(error.Error())
	}
	insert.Exec(customer.CustomerName, customer.SSN)
	//log.Println("INSERT: Customer Name: " + customer.name + " | SSN: " + customer.ssn)

	defer database.Close()

	//return Customer{}
}

func UpdateCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()

	var error error
	var update *sql.Stmt
	update, error = database.Prepare("UPDATE CUSTOMER SET CustomerName=?, SSN=? WHERE CustomerId=?")
	if error != nil {
		panic(error.Error())
	}
	update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)
	//log.Println("INSERT: Customer Name: " + customer.name + " | SSN: " + customer.ssn)

	defer database.Close()

	//return Customer{}
}
func deleteCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()

	var error error
	var delete *sql.Stmt
	delete, error = database.Prepare("DELETE FROM Customer WHERE Customerid=?")
	if error != nil {
		panic(error.Error())
	}
	delete.Exec(customer.CustomerId)
	//log.Println("INSERT: Customer Name: " + customer.name + " | SSN: " + customer.ssn)

	defer database.Close()

	//return Customer{}
}

func main() {

	var customers []Customer
	customers = GetCustomers()
	/*    fmt.Println("Before Insert",customers)

	  var customer Customer
	  customer.CustomerName = "Arnie Smith"
	  customer.SSN = "2386343"

	  InsertCustomer(customer)

	  customers = GetCustomers()
	  fmt.Println("After Insert",customers)


	fmt.Println("Before Update",customers)
	var customer Customer
	  customer.CustomerName = "George Thompson"
	  customer.SSN = "23233432"
	  customer.CustomerId = 5
	  UpdateCustomer(customer)
	  customers = GetCustomers()
	  fmt.Println("After Update",customers)
	*/

	fmt.Println("Before Delete", customers)
	var customer Customer
	customer.CustomerName = "George Thompson"
	customer.SSN = "23233432"
	customer.CustomerId = 5

	deleteCustomer(customer)
	customers = GetCustomers()
	fmt.Println("After Delete", customers)

}
