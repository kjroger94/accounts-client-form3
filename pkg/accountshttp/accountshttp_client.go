package accountshttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kjroger94/accounts/models"
)

//The struct for account which configures the actual http client
type AccountsHttpClient struct {
	url    string
	client Client
}

//The struct that the core accounts package uses to send a requst
//for the low level http client to execute
type AccountsHttpRequest struct {
	Method   string
	Headers  http.Header
	Body     []byte
	Resource string
}

//The struct that the core accounts package reads
//to evaluate the response from http client
type AccountsHttpResponse struct {
	Status       string
	StatusCode   int
	Body         *models.Data
	ErrorMessage string
}

//to return the http client
func (a *AccountsHttpClient) httpClient() Client {

	if a.client != nil {
		return a.client
	}
	a.client = &http.Client{}
	return a.client
}

//lowest level function to handle the execution of http request to form3 server
//response is handled by HandleResponse after core function validates for correct status codes
func (a *AccountsHttpClient) ExecuteRequest(ah *AccountsHttpRequest) (*http.Response, error) {
	var requrl string

	if ah.Method == "" {
		return nil, errors.New("invalid request")
	}
	requrl = a.url

	if ah.Resource != "" {
		requrl += ah.Resource
	}

	request, err := http.NewRequest(ah.Method, requrl, bytes.NewBuffer(ah.Body))
	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	response, err := a.httpClient().Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

//lowest level function to handle the response from the form3 server
//if the request had failed it will give an error or an unsatisfied object is provided
//returns the error code from the form3 server when applicable
//if no error gives the data object returned from form3
func (a *AccountsHttpClient) HandleResponse(response *http.Response) (*AccountsHttpResponse, error) {
	if response.StatusCode == 0 {
		return nil, errors.New("invalid response body")
	}
	defer response.Body.Close()

	if response.StatusCode >= 300 {
		var error_message struct {
			Message string `json:"error_message"`
		}
		err := json.NewDecoder(response.Body).Decode(&error_message)
		if err != nil {
			return nil, errors.New("could not parse error message")
		}

		finalResponse := AccountsHttpResponse{
			Status:       response.Status,
			StatusCode:   response.StatusCode,
			Body:         nil,
			ErrorMessage: error_message.Message,
		}

		return &finalResponse, errors.New("request was not fulfilled from api")

	}

	account := &models.Data{}
	json.NewDecoder(response.Body).Decode(account)

	finalResponse := AccountsHttpResponse{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Body:       account,
	}
	return &finalResponse, nil
}
