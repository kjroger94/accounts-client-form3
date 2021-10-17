package testing

import (
	"testing"

	accounts "github.com/kjroger94/accounts/cmd/accounts/app"
	"github.com/kjroger94/accounts/models"
	"github.com/stretchr/testify/assert"
)

var id []string

func set(s string) {
	id = append(id, s)
}

func unsetids() {
	id = nil
}

func TestAccountCreationsWithMinParams(t *testing.T) {
	assert := assert.New(t)

	accountsClinet, err := accounts.NewAccountsClient()
	assert.NoError(err)

	var validate *models.Data

	tests := []*models.Data{
		createBritishAccountData().Done(),
		createAustralianAccountData().Done(),
		createBelgiumAccountData().Done(),
		createCanadaAccountData().Done(),
		createEstoniaAccountData().Done(),
		createFrenchAccountData().Done(),
		createGermanAccountData().Done(),
		createGreeceAccountData().Done(),
		createHKAccountData().Done(),
		createIrishAccountData().Done(),
		createItalianAccountData().Done(),
		createLuxAccountData().Done(),
		createNetherlandsAccountData().Done(),
		createPolandAccountData().Done(),
		createPortugalAccountData().Done(),
		createSpainAccountData().Done(),
		createSwissAccountData().Done(),
		createUSAAccountData().Done(),
	}

	for _, data := range tests {
		validate, err = accountsClinet.Create(data)
		assert.NoError(err)
		assert.Equal(validate.Account.ID, data.Account.ID)
		assert.Equal(validate.Account.OrganisationID, data.Account.OrganisationID)
		assert.Equal(validate.Account.Attributes.Country, data.Account.Attributes.Country)
		assert.Equal(validate.Account.Attributes.BankID, data.Account.Attributes.BankID)
		assert.Equal(validate.Account.Attributes.BankIDCode, data.Account.Attributes.BankIDCode)
		assert.Equal(validate.Account.Type, data.Account.Type)
		assert.Equal(validate.Account.Attributes.Name, data.Account.Attributes.Name)
		assert.Equal(validate.Account.Attributes.BIC, data.Account.Attributes.BIC)

		set(validate.Account.ID)

	}

}

func TestAccountCreationsWithAdditionalParams(t *testing.T) {
	assert := assert.New(t)

	accountsClinet, err := accounts.NewAccountsClient()
	assert.NoError(err)

	var validate *models.Data

	tests := []*models.Data{
		createBritishAccountData().
			AccountClassification("Personal").
			JointAccount(false).
			ProcessingService("ABCBANK").
			UserDefinedInformation("Info").
			ValidationType("card").
			ReferenceMask("############").
			AccepranceQualifier("JKHASHD").
			Done(),

		createAustralianAccountData().
			AccountClassification("Business").
			JointAccount(true).
			ProcessingService("STRBANK").
			UserDefinedInformation("Info").
			ValidationType("card").
			ReferenceMask("############").
			AccepranceQualifier("GSUSJM").
			Done(),

		createBelgiumAccountData().
			AccountClassification("Business").
			JointAccount(true).
			ProcessingService("UTABANK").
			UserDefinedInformation("Info").
			ValidationType("card").
			ReferenceMask("############").
			AccepranceQualifier("GSUSJM").
			Done(),

		createCanadaAccountData().
			JointAccount(true).
			ProcessingService("BANK").
			UserDefinedInformation("some more info").
			ValidationType("card").
			ReferenceMask("############").
			AccepranceQualifier("GSUSJM").
			Done(),

		createEstoniaAccountData().
			JointAccount(true).
			ProcessingService("BANK").
			UserDefinedInformation("some more info").
			ValidationType("card").
			ReferenceMask("############").
			AccepranceQualifier("GSUSJM").
			Done(),

		createFrenchAccountData().
			AccountClassification("Personal").
			JointAccount(false).
			ProcessingService("ABCBANK").
			UserDefinedInformation("Info").
			ValidationType("card").
			ReferenceMask("############").
			AccepranceQualifier("JKHASHD").
			Done(),

		createGermanAccountData().
			AccountClassification("Business").
			JointAccount(false).
			ProcessingService("ABCBANK").
			UserDefinedInformation("Info").
			ValidationType("card").
			ReferenceMask("############").
			AccepranceQualifier("JKHASHD").
			Done(),

		createGreeceAccountData().
			AccountClassification("Personal").
			JointAccount(false).
			ProcessingService("ABCBANK").
			UserDefinedInformation("Info").
			ValidationType("card").
			ReferenceMask("############").
			AccepranceQualifier("JKHASHD").
			Done(),

		createHKAccountData().
			AccountClassification("Personal").
			JointAccount(false).
			ProcessingService("ABCBANK").
			UserDefinedInformation("Info").
			ValidationType("card").
			ReferenceMask("############").
			AccepranceQualifier("JKHASHD").
			Done(),

		createIrishAccountData().
			AccountClassification("Personal").
			JointAccount(false).
			ProcessingService("ABCBANK").
			UserDefinedInformation("Info").
			ValidationType("card").
			ReferenceMask("############").
			AccepranceQualifier("JKHASHD").
			Done(),

		createItalianAccountData().
			AccountClassification("Personal").
			JointAccount(false).
			ProcessingService("ABCBANK").
			UserDefinedInformation("Info").
			ValidationType("card").
			ReferenceMask("############").
			AccepranceQualifier("ASF2343").
			Done(),
		createLuxAccountData().
			AccountClassification("Personal").
			JointAccount(false).
			ProcessingService("ABCBANK").
			UserDefinedInformation("Info").
			ValidationType("card").
			Done(),

		createNetherlandsAccountData().
			AccountClassification("Personal").
			JointAccount(false).
			ProcessingService("ABCBANK").
			UserDefinedInformation("Info").
			Done(),

		createPolandAccountData().
			AccountClassification("Personal").
			JointAccount(false).
			ReferenceMask("############").
			AccepranceQualifier("JKHASHD").
			Done(),

		createPortugalAccountData().
			AccountClassification("Personal").
			JointAccount(false).
			ProcessingService("ABCBANK").
			UserDefinedInformation("Info").
			ValidationType("card").
			ReferenceMask("############").
			AccepranceQualifier("JKHASHD").
			Done(),
		createSpainAccountData().
			AccountClassification("Personal").
			JointAccount(false).
			ProcessingService("ABCBANK").
			UserDefinedInformation("Info").
			ValidationType("card").
			ReferenceMask("############").
			AccepranceQualifier("JKHASHD").
			Done(),
		createSwissAccountData().
			AccountClassification("Personal").
			JointAccount(false).
			ProcessingService("ABCBANK").
			UserDefinedInformation("Info").
			ValidationType("card").
			ReferenceMask("############").
			AccepranceQualifier("JKHASHD").
			Done(),
		createUSAAccountData().
			AccountClassification("Personal").
			JointAccount(false).
			ProcessingService("ABCBANK").
			UserDefinedInformation("Info").
			ValidationType("card").
			ReferenceMask("############").
			AccepranceQualifier("JKHASHD").
			Done(),
	}

	for _, data := range tests {
		validate, err = accountsClinet.Create(data)
		assert.NoError(err)
		assert.Equal(validate.Account.ID, data.Account.ID)
		assert.Equal(validate.Account.OrganisationID, data.Account.OrganisationID)
		assert.Equal(validate.Account.Attributes.Country, data.Account.Attributes.Country)
		assert.Equal(validate.Account.Attributes.BankID, data.Account.Attributes.BankID)
		assert.Equal(validate.Account.Attributes.BankIDCode, data.Account.Attributes.BankIDCode)
		assert.Equal(validate.Account.Type, data.Account.Type)
		assert.Equal(validate.Account.Attributes.Name, data.Account.Attributes.Name)
		assert.Equal(validate.Account.Attributes.BIC, data.Account.Attributes.BIC)
		set(validate.Account.ID)

	}

}

func TestAccountFetch(t *testing.T) {
	assert := assert.New(t)

	accountsClinet, err := accounts.NewAccountsClient()
	assert.NoError(err)

	var validate *models.Data
	for _, data := range id {
		validate, err = accountsClinet.Fetch(data)
		assert.NoError(err)
		assert.NotNil(validate.Account.ID)
		assert.NotNil(validate.Account.OrganisationID)
		assert.NotNil(validate.Account.Attributes.Country)
		assert.NotNil(validate.Account.Attributes.BankID)
		assert.NotNil(validate.Account.Attributes.BankIDCode)
		assert.NotNil(validate.Account.Type)
		assert.NotNil(validate.Account.Attributes.Name)
		assert.NotNil(validate.Account.Attributes.BIC)
	}
}

func TestAccountCreationsWithDuplicateID(t *testing.T) {
	assert := assert.New(t)

	accountsClinet, err := accounts.NewAccountsClient()
	assert.NoError(err)

	var validate *models.Data

	tests := []*models.Data{
		createBritishAccountData().Done(),
		createAustralianAccountData().Done(),
		createBelgiumAccountData().Done(),
		createCanadaAccountData().Done(),
		createEstoniaAccountData().Done(),
		createFrenchAccountData().Done(),
		createGermanAccountData().Done(),
		createGreeceAccountData().Done(),
		createHKAccountData().Done(),
		createIrishAccountData().Done(),
		createItalianAccountData().Done(),
		createLuxAccountData().Done(),
		createNetherlandsAccountData().Done(),
		createPolandAccountData().Done(),
		createPortugalAccountData().Done(),
		createSpainAccountData().Done(),
		createSwissAccountData().Done(),
		createUSAAccountData().Done(),
	}

	for i, data := range tests {
		data.Account.ID = id[i]
		validate, err = accountsClinet.Create(data)
		assert.Error(err)
		assert.Nil(validate)
	}

}

func TestAccountDelete(t *testing.T) {
	assert := assert.New(t)

	accountsClinet, err := accounts.NewAccountsClient()
	assert.NoError(err)

	var validate bool

	for _, data := range id {
		validate, err = accountsClinet.Delete(data, 0)
		assert.NoError(err)
		assert.Equal(validate, true)
	}
	unsetids()
}

func TestAccountFetchAfterDelete(t *testing.T) {
	assert := assert.New(t)

	accountsClinet, err := accounts.NewAccountsClient()
	assert.NoError(err)

	var created *models.Data

	tests := []*models.Data{
		createBritishAccountData().Done(),
		createAustralianAccountData().Done(),
		createBelgiumAccountData().Done(),
		createCanadaAccountData().Done(),
		createEstoniaAccountData().Done(),
		createFrenchAccountData().Done(),
		createGermanAccountData().Done(),
		createGreeceAccountData().Done(),
		createHKAccountData().Done(),
		createIrishAccountData().Done(),
		createItalianAccountData().Done(),
		createLuxAccountData().Done(),
		createNetherlandsAccountData().Done(),
		createPolandAccountData().Done(),
		createPortugalAccountData().Done(),
		createSpainAccountData().Done(),
		createSwissAccountData().Done(),
		createUSAAccountData().Done(),
	}

	for _, data := range tests {
		created, err = accountsClinet.Create(data)
		assert.NoError(err)
		assert.NotNil(created)

		set(created.Account.ID)
	}

	var validate bool
	var fetch *models.Data

	for _, data := range id {
		validate, err = accountsClinet.Delete(data, 0)

		assert.NoError(err)
		assert.Equal(validate, true)

		fetch, err = accountsClinet.Fetch(data)
		assert.Error(err)
		assert.Nil(fetch)

	}
}
