package accounts

import (
	"errors"

	"github.com/kjroger94/accounts/models"
	"github.com/kjroger94/accounts/pkg/accountshttp"
)

//The Accounts Interface that defines Create,Fetch and Delete
type Accounts interface {
	Create(account *models.Data) (*models.Data, error)
	Fetch(id string) (*models.Data, error)
	Delete(id string, version int) (bool, error)
}

//structure to maintain the client configured as needed to communicate
//with the form3 server with help of accountshttp package
type account struct {
	httpclient accountshttp.AccountsHttp
}

func (a *account) Create(account *models.Data) (*models.Data, error) {
	return a.create(account)
}

func (a *account) Fetch(id string) (*models.Data, error) {
	return a.fetch(id)
}

func (a *account) Delete(id string, version int) (bool, error) {
	return a.delete(id, version)
}

//Entrypoint to return account client that the end user can interact with
//when url not provided, it defaults to link in the client
func NewAccountsClient() (*account, error) {
	acc, err := accountshttp.GetHttpClient()
	_ = err //To keep returning a client consistent, this fuction will never return an error
	return &account{
		httpclient: acc,
	}, nil
}

//Entrypoint to return account client that the end user can interact with
//when url provided, it sets the base url for the client
func NewAccountsClientWithUrl(url string) (*account, error) {
	acc, err := accountshttp.GetHttpClientWithUrl(url)
	if err != nil {
		return nil, errors.New("could not get accounts client")
	}
	return &account{
		httpclient: acc,
	}, nil
}
