package accountshttp

import (
	"errors"
	"net/http"
	"net/url"
)

//Can be overriden when url is provided
const (
	Form3Url string = "http://localhost:8080/v1/organisation/accounts"
)

//Interface defnition for execution of http request and handling the response
//from the form3 server
type AccountsHttp interface {
	ExecuteRequest(*AccountsHttpRequest) (*http.Response, error)
	HandleResponse(*http.Response) (*AccountsHttpResponse, error)
}

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

//Initializes and returns the account client object
//to the core package with the default url
func GetHttpClient() (AccountsHttp, error) {
	return &AccountsHttpClient{
		url: Form3Url,
	}, nil
}

//Initializes and returns the account client object
//to the core package with the provided  url
func GetHttpClientWithUrl(userurl string) (AccountsHttp, error) {
	if userurl == "" {
		return nil, errors.New("invalid URL")

	}

	u, err := url.ParseRequestURI(userurl)
	_ = u
	if err != nil {
		return nil, errors.New("invalid URL")
	}

	return &AccountsHttpClient{
		url: userurl,
	}, nil
}
