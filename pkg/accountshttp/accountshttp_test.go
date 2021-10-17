package accountshttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockTest struct {
	mockResponse func(req *http.Request) (*http.Response, error)
}

func (m *mockTest) Do(req *http.Request) (*http.Response, error) {
	return m.mockResponse(req)
}

func TestExecuteRequestWithValidObject(t *testing.T) {
	assert := assert.New(t)
	var m mockTest

	a := AccountsHttpClient{
		url:    "http://localhost:8080/v1/organisation/accounts",
		client: &m,
	}

	reqdata := `{"request":"json"}`
	resdata := `{"response":"json"}`

	var j struct{ k, v string }

	re := ioutil.NopCloser(bytes.NewReader([]byte(resdata)))

	r, err := json.Marshal(reqdata)
	assert.NoError(err)

	m.mockResponse = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       re,
		}, nil
	}

	req := &AccountsHttpRequest{
		Method: "POST",
		Body:   r,
	}

	resp, err := a.ExecuteRequest(req)
	assert.NoError(err)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.NoError(err)

	err = json.Unmarshal(bodyBytes, &j)
	assert.NoError(err)

	assert.NotNil(resp)
	assert.NotNil(bodyBytes)
	assert.NoError(err)

}

func TestExecuteRequestExecutionFail(t *testing.T) {
	assert := assert.New(t)
	var m mockTest

	a := AccountsHttpClient{
		url:    "http://localhost:8080/v1/organisation/accounts",
		client: &m,
	}

	reqdata := `{"request":"json"}`

	r, err := json.Marshal(reqdata)
	assert.NoError(err)

	m.mockResponse = func(*http.Request) (*http.Response, error) {
		return nil, errors.New("could not finish request, link down")
	}

	req := &AccountsHttpRequest{
		Method: "POST",
		Body:   r,
	}

	resp, err := a.ExecuteRequest(req)
	assert.Error(err)
	assert.Nil(resp)

}

func TestExecuteRequestWithValidObjectAndResource(t *testing.T) {
	assert := assert.New(t)
	var m mockTest

	a := AccountsHttpClient{
		url:    "http://localhost:8080/v1/organisation/accounts",
		client: &m,
	}

	reqdata := `{"request":"json"}`
	resdata := `{"response":"json"}`

	var j struct{ k, v string }

	re := ioutil.NopCloser(bytes.NewReader([]byte(resdata)))

	r, err := json.Marshal(reqdata)
	assert.NoError(err)

	m.mockResponse = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       re,
		}, nil
	}

	req := &AccountsHttpRequest{
		Method:   "POST",
		Body:     r,
		Resource: "/test",
	}

	resp, err := a.ExecuteRequest(req)
	assert.NoError(err)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.NoError(err)

	err = json.Unmarshal(bodyBytes, &j)
	assert.NotNil(resp)
	assert.NotNil(bodyBytes)
	assert.NoError(err)

}

func TestExecuteRequestWithNil(t *testing.T) {
	assert := assert.New(t)
	var m mockTest

	a := AccountsHttpClient{
		url:    "http://localhost:8080/v1/organisation/accounts",
		client: &m,
	}

	req := &AccountsHttpRequest{}

	resp, err := a.ExecuteRequest(req)
	assert.Error(err)
	assert.Nil(resp)

}

func TestExecuteRequestWithInvalidHttpResponse(t *testing.T) {
	assert := assert.New(t)
	var m mockTest

	a := AccountsHttpClient{
		url:    "http://localhost:8080/v1/organisation/accounts",
		client: &m,
	}

	reqdata := `{"request":"json"}`
	r, err := json.Marshal(reqdata)
	assert.NoError(err)

	m.mockResponse = func(*http.Request) (*http.Response, error) {
		return nil, errors.New("Could not satisfy your req")
	}

	req := &AccountsHttpRequest{
		Method: "POST",
		Body:   r,
	}

	resp, err := a.ExecuteRequest(req)
	assert.Nil(resp)
	assert.Error(err)

}

func TestHandleResponseWithValidJSON(t *testing.T) {
	assert := assert.New(t)
	var m mockTest

	a := AccountsHttpClient{
		url:    "http://localhost:8080/v1/organisation/accounts",
		client: &m,
	}

	resdata := `{"data":{"id":"66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2","organisation_id":"a4a01942-2e57-11ec-a3ff-deb7de462b1d","type":"accounts","Attributes":{"country":"GB","base_currency":"GBP","bank_id":"400300","bank_id_code":"GBDSC","bic":"NWBKGB22123","name":["Eminem"]}}}`

	re := ioutil.NopCloser(bytes.NewReader([]byte(resdata)))

	param := &http.Response{
		StatusCode: 201,
		Body:       re,
	}

	resp, err := a.HandleResponse(param)
	assert.NotNil(resp)
	assert.NoError(err)

}

func TestHandleResponseWithInvalidResponseBody(t *testing.T) {
	assert := assert.New(t)
	var m mockTest

	a := AccountsHttpClient{
		url:    "http://localhost:8080/v1/organisation/accounts",
		client: &m,
	}

	param := &http.Response{}

	resp, err := a.HandleResponse(param)
	assert.Nil(resp)
	assert.Error(err)

}

func TestGetHttpClient(t *testing.T) {
	assert := assert.New(t)

	client, err := GetHttpClient()
	assert.NotNil(client)
	assert.NoError(err)

}

func TestGetHttpClientForUrl(t *testing.T) {
	assert := assert.New(t)

	client, err := GetHttpClientWithUrl("http://localhost.com")
	assert.NotNil(client)
	assert.NoError(err)
}

func TestGetHttpClientForUrlWhenInvalid(t *testing.T) {
	assert := assert.New(t)

	client, err := GetHttpClientWithUrl("htp//localhost.com")
	assert.Nil(client)
	assert.Error(err)
}

func TestHandleResponseInvalidResponseJson(t *testing.T) {
	assert := assert.New(t)
	var m mockTest

	a := AccountsHttpClient{
		url:    "http://localhost:8080/v1/organisation/accounts",
		client: &m,
	}

	resdata := `{"data":{"id":"66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2","org":"a4a01942-2e57-11ec-a3ff-deb7de462b1d","type":"accounts","attr":{"country":"GB","base_currency":"GBP","bank_id":"400300","bank_id_code":"GBDSC","bic":"NWBKGB22123","name":["Eminem"]}}}`

	re := ioutil.NopCloser(bytes.NewReader([]byte(resdata)))

	param := &http.Response{
		StatusCode: 201,
		Body:       re,
	}

	resp, err := a.HandleResponse(param)
	assert.NotNil(resp)
	assert.NoError(err)

}

func TestHandleResponseErrorMessage(t *testing.T) {
	assert := assert.New(t)
	var m mockTest

	a := AccountsHttpClient{
		url:    "http://localhost:8080/v1/organisation/accounts",
		client: &m,
	}

	resdata := `{"error_message":"validation failure list:\nvalidation failure list:\nvalidation failure list:\nname in body is required"}`

	re := ioutil.NopCloser(bytes.NewReader([]byte(resdata)))

	param := &http.Response{
		StatusCode: 404,
		Body:       re,
	}

	resp, err := a.HandleResponse(param)
	assert.NotNil(resp)
	assert.Error(err)

}

func TestHandleResponseUnparsableError(t *testing.T) {
	assert := assert.New(t)
	var m mockTest

	a := AccountsHttpClient{
		url:    "http://localhost:8080/v1/organisation/accounts",
		client: &m,
	}

	resdata := ``

	re := ioutil.NopCloser(bytes.NewReader([]byte(resdata)))

	param := &http.Response{
		StatusCode: 404,
		Body:       re,
	}

	resp, err := a.HandleResponse(param)
	assert.Nil(resp)
	assert.Error(err)

}
