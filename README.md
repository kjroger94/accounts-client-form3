Submitted by: Kshitij Sinha (cartersinha@gmail.com)
#### Note
I am new to Go in this context to what I have just developed. I do not have deep or extensive experience in the language. I had worked within limited boundations with Go while working with smartcontracts in blockchain framework Hyperledger but nothing this comprehensive and detailed.

I am primarily and Node.js developer with fair bit of exeperinece in C as well.

# Form3 Accounts Client
This is a go client library that interacts with the Form3 fake accounts API. It provides functionality for:

- Create: Creation of an account
- Fetch: Getting the details for a given account 
- Delete: Deleting an account

### Creating the client

```go
    //Creating with default base url
	accounts,err := accounts.NewAccountsClient()
    //to override configured url call NewAccountsClientWithUrl(url string)
	accounts,err := accounts.NewAccountsClientWithUrl("http://api.staging-form3.tech/v1/organisation/accounts")
```

### Creating parameters for creation of account
The parameters can be given in any order, 
call `NewAccountsParams()` to start initializing the params and finally `Done()` to get an object ready to be sent for creation
```go
params := accountsparams.NewAccountsParams(). //This generates a new account ID ingrained
		OrganisationID("cf4d943a-d535-4812-86a4-e07d7633bb60"). //to be set in the client directly as config for an org
		Type("accounts").
		Name("Michael Scott").
		Country("GB").
		BaseCurrency("GBP").
		BankID("400300").
		BankIDCode("GBDSC").
		BIC("NWBKGB22").
		ProcessingService("ABCBANK").
		AlternativeNames("Michael Klump", "Michael Scarn").
		UserDefinedInformation("Info").
		ValidationType("card").
		ReferenceMask("############").
		Done()  //Done is to be called once the paramters are set by the user
```
now, these params can be passed to the create function

### Account Creation
```go
    data, err := accounts.Create(accountParameters.Account)
    fmt.Println("Created the account",data.Account)
```
### Fetch Account
```go
    data, err := accounts.Fetch("7a1f0ef9-f26b-4aa7-91dd-0a6f3f3c755b")
    fmt.Println("Created the account",data.Account)
```
### Delete Account
```go
    data, err := accounts.Delete("953548c1-e8e2-4ad2-bf10-43c48ab7d63e",0)
    //Delete returns true if the account was sucessfully deleted
```
### Test

The integration tests are in the `testing` folder 
Unit tests are in each folder as the code itself

Run tests 
```bash
go -v -cover ./...
```
With the project - inspect logs, verbose enabled tests
```
docker-compose up
```

### Improvements and considerations
1. The client initialization `NewAccountsClient()` can be extended take options for time outs, connection idle times, response time.
2. The organisation id could be fixed for the client or configurable, I wasn't sure how this ID is used so chose to keep it parameterized for now.
3. The code can be more modular and new interfaces can be introduced so that user of the client library can plug in more features as they seem fit.
4. Introduction of constants and other generators so that for creation of accounts for specific contries the user of the library have to spend minimal time. The rules given in the form3 documentation can be adhered to by default. Specifically `cmd/accountsparams` can be extended
5. Concurrency protection and support (I do not fully understand this in full detail as I am new to Go)