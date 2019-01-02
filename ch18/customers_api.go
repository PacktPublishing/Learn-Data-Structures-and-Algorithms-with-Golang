
//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt,database/sql, net/http, text/template package
import (
  //  "fmt"
    "net/http"
  //  "text/template"
//    "errors"
    "log"
)




func Index(writer http.ResponseWriter, request *http.Request) {

    log.Println("Running Customer API")


}

func GetCustomers(writer http.ResponseWriter, request *http.Request) {
  var customers []Customer
  customers = GetCustomersFromDB()
  log.Println(customers)

}

func main() {
    log.Println("Server started on: http://localhost:8000")

    http.HandleFunc("/", Index)
    http.HandleFunc("/customers", GetCustomers)

    http.ListenAndServe(":8000", nil)
}
