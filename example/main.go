package main

import (
	"fmt"

	accounts "github.com/kjroger94/accounts/cmd/accounts/app"
	"github.com/kjroger94/accounts/cmd/accountsparams"
)

func main() {

	accounts, err := accounts.NewAccountsClient()
	_ = err //This function returns a nil error alwasys, kept to be consistent with NewAccountsClientWithUrl(s string)

	params := accountsparams.NewAccountsParams().
		OrganisationID("9100d649-57a8-4ed4-b296-9b222ee072ca").
		Type("accounts").
		Name("Laura Grey").
		Country("GB").
		BaseCurrency("GBP").
		BankID("400300").
		BankIDCode("GBDSC").
		BIC("NWBKGB22").
		ProcessingService("ABCBANK").
		UserDefinedInformation("Info").
		AlternativeNames("Salim Shady", "Your Neighbour").
		ValidationType("card").
		ReferenceMask("############").
		AccepranceQualifier("SameDay").
		Done() //Done is to be called once the paramters are set by the user

	//Account Creation
	data, err := accounts.Create(params)
	if err != nil {
		fmt.Println("Error in account creation", err)
	}
	fmt.Printf("Created account: %v\n", data.Account.ID)

	//Fetch Account
	fetchedaccount, err := accounts.Fetch(data.Account.ID) //Any valid UUID, I am fetching the account just created
	if err != nil {
		fmt.Println("Error in fetching account", err)
	}
	fmt.Printf("Account details: %v\n", fetchedaccount.Account)

	//Delete account
	isDeleted, err := accounts.Delete(data.Account.ID, 0) //Any valid uuid and a int version of the account you created
	if err != nil {
		fmt.Println("Error in deleting account", err)
	}
	fmt.Printf("Deleted account %v, function returned %t\n", data.Account.ID, isDeleted)

}
