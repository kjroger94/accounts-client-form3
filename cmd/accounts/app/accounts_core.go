package accounts

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/kjroger94/accounts/models"
	"github.com/kjroger94/accounts/pkg/accountshttp"
)

//The low level create function that communicates with the accountshttp package to
//manage the http client request-response to create the account with form3
func (a *account) create(account *models.Data) (*models.Data, error) {
	if account.Account == nil {
		return nil, errors.New("cannot accept empty body")
	}

	body, err := json.Marshal(account)
	if err != nil {
		return nil, errors.New("could not marshal json")
	}

	requestBody := &accountshttp.AccountsHttpRequest{
		Method: "POST",
		Body:   body,
	}

	resp, err := a.httpclient.ExecuteRequest(requestBody)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 {
		return nil, errors.New("resource not created")

	}

	data, err := a.httpclient.HandleResponse(resp)
	if err != nil {
		if data != nil {
			return nil, errors.New(data.ErrorMessage)
		}
		return nil, err
	}
	return data.Body, nil
}

//The low level create function that communicates with the accountshttp package to
//manage the http client request-response to fetch the account with form3

func (a *account) fetch(id string) (*models.Data, error) {
	if id == "" {
		return nil, errors.New("cannot accept empty string")
	}

	requestBody := &accountshttp.AccountsHttpRequest{
		Method:   "GET",
		Resource: "/" + id,
	}

	resp, err := a.httpclient.ExecuteRequest(requestBody)
	if err != nil {
		return nil, errors.New("request error")
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("request failed")
	}

	data, err := a.httpclient.HandleResponse(resp)
	if err != nil {
		if data != nil {
			return nil, errors.New(data.ErrorMessage)
		}
		return nil, errors.New("response error")
	}
	return data.Body, nil
}

//The low level create function that communicates with the accountshttp package to
//manage the http client request-response to delete the account with form3
func (a *account) delete(id string, version int) (bool, error) {
	if id == "" {
		return false, errors.New("cannot accept empty string")

	}
	resource := fmt.Sprintf("/" + id + "?version=" + strconv.Itoa(version))
	requestBody := &accountshttp.AccountsHttpRequest{
		Method:   "DELETE",
		Resource: resource,
	}

	resp, err := a.httpclient.ExecuteRequest(requestBody)
	if err != nil {
		return false, errors.New("request error")
	}

	if resp.StatusCode != 204 {
		return false, errors.New("resource not deleted")
	}

	data, err := a.httpclient.HandleResponse(resp)
	if err != nil {
		if data != nil {
			return false, errors.New(data.ErrorMessage)
		}
		return false, errors.New("response error")
	}

	return true, nil
}
