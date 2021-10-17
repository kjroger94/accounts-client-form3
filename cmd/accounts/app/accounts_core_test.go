package accounts

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/kjroger94/accounts/models"
	"github.com/kjroger94/accounts/pkg/accountshttp"
	"github.com/stretchr/testify/assert"
)

type mockreqres struct {
	executeRequest func(ah *accountshttp.AccountsHttpRequest) (*http.Response, error)
	handleResponse func(r *http.Response) (*accountshttp.AccountsHttpResponse, error)

	// create func(account *models.Data) (models.Data, error)
	// fetch  func(id string) (models.Data, error)
	// delete func(id string, version int) (bool, error)
}

func (m *mockreqres) ExecuteRequest(ah *accountshttp.AccountsHttpRequest) (*http.Response, error) {
	return m.executeRequest(ah)
}
func (m *mockreqres) HandleResponse(r *http.Response) (*accountshttp.AccountsHttpResponse, error) {
	return m.handleResponse(r)
}

func TestCreateAccountCore(t *testing.T) {
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

	account := account{
		httpclient: &m,
	}

	resp, err := account.create(reqDataPtr)
	assert.NoError(err)
	assert.NotNil(resp)

}

func TestCreateAccountCoreInvalidResponseJson(t *testing.T) {
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
		return nil, errors.New("failed to get the response body")
	}

	account := account{
		httpclient: &m,
	}

	resp, err := account.create(reqDataPtr)
	assert.Error(err)
	assert.Nil(resp)

}

func TestCreateAccountCoreNon201(t *testing.T) {
	assert := assert.New(t)

	var m mockreqres

	reqAndResString := `{"data":{"id":"66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2","organisation_id":"a4a01942-2e57-11ec-a3ff-deb7de462b1d","type":"accounts","Attributes":{"country":"GB","base_currency":"GBP","bank_id":"400300","bank_id_code":"GBDSC","bic":"NWBKGB22123","name":["Rames"]}}}`
	reqData := strings.NewReader(reqAndResString)
	reqDataPtr := &models.Data{}
	json.NewDecoder(reqData).Decode(reqDataPtr)

	m.executeRequest = func(ah *accountshttp.AccountsHttpRequest) (*http.Response, error) {
		return &http.Response{
			StatusCode: 400,
			Body:       nil,
		}, nil
	}

	account := account{
		httpclient: &m,
	}

	resp, err := account.create(reqDataPtr)
	assert.Nil(resp)
	assert.Error(err)

}

func TestCreateAccountCoreNon201ErrMsg(t *testing.T) {
	assert := assert.New(t)

	var m mockreqres

	reqAndResString := `{"data":{"id":"66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2","organisation_id":"a4a01942-2e57-11ec-a3ff-deb7de462b1d","type":"accounts","Attributes":{"country":"GB","base_currency":"GBP","bank_id":"400300","bank_id_code":"GBDSC","bic":"NWBKGB22123","name":["Rames"]}}}`
	reqData := strings.NewReader(reqAndResString)
	reqDataPtr := &models.Data{}
	json.NewDecoder(reqData).Decode(reqDataPtr)

	errmsg := `{"error_message":"validation failure list:\nvalidation failure list:\nvalidation failure list:\nbank_id_code in body should match '^[A-Z]{0,16}$'"}`
	r := io.NopCloser(strings.NewReader(errmsg))

	m.executeRequest = func(ah *accountshttp.AccountsHttpRequest) (*http.Response, error) {
		return &http.Response{
			StatusCode: 201,
			Body:       r,
		}, nil
	}

	m.handleResponse = func(r *http.Response) (*accountshttp.AccountsHttpResponse, error) {
		return &accountshttp.AccountsHttpResponse{
			Body:         reqDataPtr,
			Status:       "OK",
			StatusCode:   400,
			ErrorMessage: errmsg,
		}, errors.New("err msg")
	}

	account := account{
		httpclient: &m,
	}

	resp, err := account.create(reqDataPtr)
	assert.Nil(resp)
	assert.Error(err)

}

func TestCreateAccountCoreExecuteRequestFail(t *testing.T) {
	assert := assert.New(t)

	var m mockreqres

	reqAndResString := `{"data":{"id":"66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2","organisation_id":"a4a01942-2e57-11ec-a3ff-deb7de462b1d","type":"accounts","Attributes":{"country":"GB","base_currency":"GBP","bank_id":"400300","bank_id_code":"GBDSC","bic":"NWBKGB22123","name":["Rames"]}}}`
	reqData := strings.NewReader(reqAndResString)
	reqDataPtr := &models.Data{}
	json.NewDecoder(reqData).Decode(reqDataPtr)

	m.executeRequest = func(ah *accountshttp.AccountsHttpRequest) (*http.Response, error) {
		return nil, errors.New("could not process request")
	}

	account := account{
		httpclient: &m,
	}

	resp, err := account.create(reqDataPtr)
	assert.Nil(resp)
	assert.Error(err)
}

func TestCreateAccountCoreExecuteRequestEmptyStruct(t *testing.T) {
	assert := assert.New(t)

	var m mockreqres

	reqDataPtr := &models.Data{}

	account := account{
		httpclient: &m,
	}

	resp, err := account.create(reqDataPtr)
	assert.Nil(resp)
	assert.Error(err)

}

func TestFetchAccountCore(t *testing.T) {
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

	account := account{
		httpclient: &m,
	}

	resp, err := account.fetch("66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2")
	assert.NoError(err)
	assert.NotNil(resp)

}

func TestFetchAccountCoreStringNotProper(t *testing.T) {
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

	account := account{
		httpclient: &m,
	}

	resp, err := account.fetch("")
	assert.Error(err)
	assert.Nil(resp)
}

func TestFetchAccountCoreRequestError(t *testing.T) {
	assert := assert.New(t)

	var m mockreqres

	resString := `{"data":{"id":"66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2","organisation_id":"a4a01942-2e57-11ec-a3ff-deb7de462b1d","type":"accounts","Attributes":{"country":"GB","base_currency":"GBP","bank_id":"400300","bank_id_code":"GBDSC","bic":"NWBKGB22123","name":["Rames"]}}}`
	reqData := strings.NewReader(resString)
	reqDataPtr := &models.Data{}
	json.NewDecoder(reqData).Decode(reqDataPtr)

	m.executeRequest = func(ah *accountshttp.AccountsHttpRequest) (*http.Response, error) {
		return nil, errors.New("request failed")
	}

	account := account{
		httpclient: &m,
	}

	resp, err := account.fetch("66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2")
	assert.Error(err)
	assert.Nil(resp)

}

func TestFetchAccountCoreNon200Response(t *testing.T) {
	assert := assert.New(t)

	var m mockreqres

	resString := `{"data":{"id":"66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2","organisation_id":"a4a01942-2e57-11ec-a3ff-deb7de462b1d","type":"accounts","Attributes":{"country":"GB","base_currency":"GBP","bank_id":"400300","bank_id_code":"GBDSC","bic":"NWBKGB22123","name":["Rames"]}}}`
	reqData := strings.NewReader(resString)
	reqDataPtr := &models.Data{}
	json.NewDecoder(reqData).Decode(reqDataPtr)

	m.executeRequest = func(ah *accountshttp.AccountsHttpRequest) (*http.Response, error) {
		return &http.Response{
			StatusCode: 404,
			Body:       nil,
		}, nil
	}

	account := account{
		httpclient: &m,
	}

	resp, err := account.fetch("66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2")
	assert.Error(err)
	assert.Nil(resp)
}

func TestFetchAccountCoreNonErrMsg(t *testing.T) {
	assert := assert.New(t)

	var m mockreqres

	resString := `{"data":{"id":"66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2","organisation_id":"a4a01942-2e57-11ec-a3ff-deb7de462b1d","type":"accounts","Attributes":{"country":"GB","base_currency":"GBP","bank_id":"400300","bank_id_code":"GBDSC","bic":"NWBKGB22123","name":["Rames"]}}}`
	reqData := strings.NewReader(resString)
	reqDataPtr := &models.Data{}
	json.NewDecoder(reqData).Decode(reqDataPtr)

	errmsg := `{"error_message":"invalidid"}`
	r := io.NopCloser(strings.NewReader(errmsg))

	m.executeRequest = func(ah *accountshttp.AccountsHttpRequest) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	m.handleResponse = func(r *http.Response) (*accountshttp.AccountsHttpResponse, error) {
		return &accountshttp.AccountsHttpResponse{
			Body:         reqDataPtr,
			Status:       "OK",
			StatusCode:   400,
			ErrorMessage: errmsg,
		}, errors.New("err msg")
	}

	account := account{
		httpclient: &m,
	}

	resp, err := account.fetch("66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2")
	assert.Error(err)
	assert.Nil(resp)
}

func TestFetchAccountCoreResponseHandlerFailed(t *testing.T) {
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
		return nil, errors.New("unable to fetch response")
	}

	account := account{
		httpclient: &m,
	}

	resp, err := account.fetch("66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2")
	assert.Error(err)
	assert.Nil(resp)

}

func TestDeleteAccountCore(t *testing.T) {
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

	account := account{
		httpclient: &m,
	}

	resp, err := account.delete("66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2", 0)
	assert.NoError(err)
	assert.NotNil(resp)

}

func TestDeleteAccountCoreNon204Response(t *testing.T) {
	assert := assert.New(t)

	var m mockreqres

	resString := `{}`
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

	account := account{
		httpclient: &m,
	}

	resp, err := account.delete("66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2", 0)
	assert.Error(err)
	assert.Equal(resp, false)

}

func TestDeleteAccountCoreNonErrMsg(t *testing.T) {
	assert := assert.New(t)

	var m mockreqres

	resString := `{}`
	reqData := strings.NewReader(resString)
	reqDataPtr := &models.Data{}
	json.NewDecoder(reqData).Decode(reqDataPtr)

	emptyObj := &models.Data{}

	errmsg := `{"error_message":"invalid_id"}`
	r := io.NopCloser(strings.NewReader(errmsg))

	m.executeRequest = func(ah *accountshttp.AccountsHttpRequest) (*http.Response, error) {
		return &http.Response{
			StatusCode: 204,
			Body:       r,
		}, nil
	}

	m.handleResponse = func(r *http.Response) (*accountshttp.AccountsHttpResponse, error) {
		return &accountshttp.AccountsHttpResponse{
			Body:         emptyObj,
			Status:       "OK",
			StatusCode:   404,
			ErrorMessage: errmsg,
		}, errors.New("err msg")
	}

	account := account{
		httpclient: &m,
	}

	resp, err := account.delete("66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2", 0)
	assert.Error(err)
	assert.Equal(resp, false)

}

func TestDeleteAccountCoreNonReqError(t *testing.T) {
	assert := assert.New(t)

	var m mockreqres

	resString := `{}`
	reqData := strings.NewReader(resString)
	reqDataPtr := &models.Data{}
	json.NewDecoder(reqData).Decode(reqDataPtr)

	m.executeRequest = func(ah *accountshttp.AccountsHttpRequest) (*http.Response, error) {
		return nil, errors.New("could not perform task")
	}

	account := account{
		httpclient: &m,
	}

	resp, err := account.delete("66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2", 0)
	assert.Error(err)
	assert.Equal(resp, false)

}

func TestDeleteAccountCoreResponseHandlerError(t *testing.T) {
	assert := assert.New(t)

	var m mockreqres

	resString := `{}`
	reqData := strings.NewReader(resString)
	reqDataPtr := &models.Data{}
	json.NewDecoder(reqData).Decode(reqDataPtr)

	re := ioutil.NopCloser(bytes.NewReader([]byte(resString)))

	m.executeRequest = func(ah *accountshttp.AccountsHttpRequest) (*http.Response, error) {
		return &http.Response{
			StatusCode: 204,
			Body:       re,
		}, nil
	}

	m.handleResponse = func(r *http.Response) (*accountshttp.AccountsHttpResponse, error) {
		return nil, errors.New("could not process response")
	}

	account := account{
		httpclient: &m,
	}

	resp, err := account.delete("66bc6421-2d96-4c4e-9d2b-864d6f4cf2a2", 0)
	assert.Error(err)
	assert.Equal(resp, false)

}

func TestDeleteAccountCoreInvalidID(t *testing.T) {
	assert := assert.New(t)

	var m mockreqres

	account := account{
		httpclient: &m,
	}

	resp, err := account.delete("", 0)
	assert.Error(err)
	assert.Equal(resp, false)

}
