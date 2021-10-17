package accounts

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/kjroger94/accounts/models"
	"github.com/kjroger94/accounts/pkg/accountshttp"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {

	assert := assert.New(t)

	var m mockreqres

	reqAndResString := `{"data":{"id":"66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2","organisation_id":"a4a01942-2e57-11ec-a3ff-deb7de462b1d","type":"accounts","Attributes":{"country":"GB","base_currency":"GBP","bank_id":"400300","bank_id_code":"GBDSC","bic":"NWBKGB22123","name":["Rames"]}}}`
	reqData := strings.NewReader(reqAndResString)
	reqDataPtr := &models.Data{}
	json.NewDecoder(reqData).Decode(reqDataPtr)

	re := ioutil.NopCloser(bytes.NewReader([]byte(reqAndResString)))

	m.executeRequest = func(ah *accountshttp.AccountsHttpRequest) (*http.Response, error) {
		return &http.Response{
			StatusCode: 201,
			Body:       re,
		}, nil
	}

	m.handleResponse = func(r *http.Response) (*accountshttp.AccountsHttpResponse, error) {
		return &accountshttp.AccountsHttpResponse{
			Body:       reqDataPtr,
			Status:     "OK",
			StatusCode: 201,
		}, nil
	}

	account, err := NewAccountsClient()
	account.httpclient = &m
	assert.NoError(err)

	resp, err := account.Create(reqDataPtr)
	assert.NoError(err)
	assert.NotNil(resp)
}

func TestFetchAccount(t *testing.T) {
	assert := assert.New(t)

	var m mockreqres

	resString := `{"data":{"id":"66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2","organisation_id":"a4a01942-2e57-11ec-a3ff-deb7de462b1d","type":"accounts","Attributes":{"country":"GB","base_currency":"GBP","bank_id":"400300","bank_id_code":"GBDSC","bic":"NWBKGB22123","name":["Rames"]}}}`
	reqData := strings.NewReader(resString)
	reqDataPtr := &models.Data{}
	json.NewDecoder(reqData).Decode(reqDataPtr)

	re := ioutil.NopCloser(bytes.NewReader([]byte(resString)))

	m.executeRequest = func(ah *accountshttp.AccountsHttpRequest) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       re,
		}, nil
	}

	m.handleResponse = func(r *http.Response) (*accountshttp.AccountsHttpResponse, error) {
		return &accountshttp.AccountsHttpResponse{
			Body:       reqDataPtr,
			Status:     "OK",
			StatusCode: 200,
		}, nil
	}

	account, err := NewAccountsClient()
	account.httpclient = &m
	assert.NoError(err)

	resp, err := account.Fetch("66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2")
	assert.NoError(err)
	assert.NotNil(resp)

}
func TestDeleteAccount(t *testing.T) {
	assert := assert.New(t)

	var m mockreqres

	resString := `{}`
	reqData := strings.NewReader(resString)
	reqDataPtr := &models.Data{}
	json.NewDecoder(reqData).Decode(reqDataPtr)

	re := ioutil.NopCloser(bytes.NewReader([]byte(resString)))
	emptyObj := &models.Data{}

	m.executeRequest = func(ah *accountshttp.AccountsHttpRequest) (*http.Response, error) {
		return &http.Response{
			StatusCode: 204,
			Body:       re,
		}, nil
	}

	m.handleResponse = func(r *http.Response) (*accountshttp.AccountsHttpResponse, error) {
		return &accountshttp.AccountsHttpResponse{
			Body:       emptyObj,
			Status:     "OK",
			StatusCode: 204,
		}, nil
	}

	account, err := NewAccountsClient()
	account.httpclient = &m
	assert.NoError(err)

	resp, err := account.Delete("66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2", 0)
	assert.NoError(err)
	assert.NotNil(resp)

}

func TestCreateAccountWithUrl(t *testing.T) {

	assert := assert.New(t)

	var m mockreqres

	reqAndResString := `{"data":{"id":"66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2","organisation_id":"a4a01942-2e57-11ec-a3ff-deb7de462b1d","type":"accounts","Attributes":{"country":"GB","base_currency":"GBP","bank_id":"400300","bank_id_code":"GBDSC","bic":"NWBKGB22123","name":["Rames"]}}}`
	reqData := strings.NewReader(reqAndResString)
	reqDataPtr := &models.Data{}
	json.NewDecoder(reqData).Decode(reqDataPtr)

	re := ioutil.NopCloser(bytes.NewReader([]byte(reqAndResString)))

	m.executeRequest = func(ah *accountshttp.AccountsHttpRequest) (*http.Response, error) {
		return &http.Response{
			StatusCode: 201,
			Body:       re,
		}, nil
	}

	m.handleResponse = func(r *http.Response) (*accountshttp.AccountsHttpResponse, error) {
		return &accountshttp.AccountsHttpResponse{
			Body:       reqDataPtr,
			Status:     "OK",
			StatusCode: 201,
		}, nil
	}

	account, err := NewAccountsClientWithUrl("https://api.form3.com/accounts")
	account.httpclient = &m
	assert.NoError(err)

	resp, err := account.Create(reqDataPtr)
	assert.NoError(err)
	assert.NotNil(resp)
}

func TestCreateAccountWithInvalid(t *testing.T) {

	assert := assert.New(t)

	account, err := NewAccountsClientWithUrl("")
	assert.Error(err)
	assert.Nil(account)

}
