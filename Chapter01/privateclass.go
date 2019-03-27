//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and encoding/json pacakges
import (
	"encoding/json"
	"fmt"
)

//AccountDetails struct
type AccountDetails struct {
	id          string
	accountType string
}

//Account struct
type Account struct {
	details      *AccountDetails
	CustomerName string
}

// Account class method setDetails
func (account *Account) setDetails(id string, accountType string) {

	account.details = &AccountDetails{id, accountType}
}

//Account class method getId
func (account *Account) getId() string {

	return account.details.id
}

//Account class method getAccountType
func (account *Account) getAccountType() string {

	return account.details.accountType
}

// main method
func main() {

	var account *Account = &Account{CustomerName: "John Smith"}
	account.setDetails("4532", "current")

	jsonAccount, _ := json.Marshal(account)
	fmt.Println("Private Class hidden", string(jsonAccount))

	fmt.Println("Account Id", account.getId())

	fmt.Println("Account Type", account.getAccountType())

}
