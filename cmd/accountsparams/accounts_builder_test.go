package accountsparams

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewAccountsParams(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams()
	assert.NotNil(params)
}

func TestOrganisationID(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString())
	assert.NotNil(params)
}

func TestType(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts")
	assert.NotNil(params)
}

func TestVersion(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1)
	assert.NotNil(params)
}

func TestCountry(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB")
	assert.NotNil(params)
}

func TestBaseCurrency(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP")
	assert.NotNil(params)
}

func TestBankID(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS")
	assert.NotNil(params)
}

func TestBankIDCode(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS").
		BankIDCode("373BSJS")
	assert.NotNil(params)
}

func TestAccountNumber(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS").
		BankIDCode("373BSJS").
		AccountNumber("AISHD2992KKS")
	assert.NotNil(params)
}

func TestBIC(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS").
		BankIDCode("373BSJS").
		BIC("AISHD2992KKS")
	assert.NotNil(params)
}

func TestIBAN(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS").
		IBAN("373BSJS").
		AccountNumber("AISHD2992KKS")
	assert.NotNil(params)
}

func TestCustomerID(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS").
		IBAN("373BSJS").
		CustomerID("AISHD2992KKS")
	assert.NotNil(params)
}

func TestName(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS").
		IBAN("373BSJS").
		CustomerID("AISHD2992KKS").
		Name("Kshitij")
	assert.NotNil(params)
}

func TestAlternativeNames(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS").
		IBAN("373BSJS").
		CustomerID("AISHD2992KKS").
		AlternativeNames("Kshitij")
	assert.NotNil(params)
}

func TestAccountClassification(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS").
		IBAN("373BSJS").
		CustomerID("AISHD2992KKS").
		AccountClassification("Shure")
	assert.NotNil(params)
}

func TestJointAccount(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS").
		IBAN("373BSJS").
		CustomerID("AISHD2992KKS").
		JointAccount(false)
	assert.NotNil(params)
}

func TestAccountMatchingOptOut(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS").
		IBAN("373BSJS").
		CustomerID("AISHD2992KKS").
		AccountMatchingOptOut(true)
	assert.NotNil(params)
}

func TestSecondaryIdentification(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS").
		IBAN("373BSJS").
		CustomerID("AISHD2992KKS").
		SecondaryIdentification("NA")
	assert.NotNil(params)
}

func TestSwitched(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS").
		IBAN("373BSJS").
		CustomerID("AISHD2992KKS").
		Switched(false)
	assert.NotNil(params)
}

func TestProcessingService(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS").
		IBAN("373BSJS").
		CustomerID("AISHD2992KKS").
		ProcessingService("Form3")
	assert.NotNil(params)
}

func TestUserDefinedInformation(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS").
		IBAN("373BSJS").
		CustomerID("AISHD2992KKS").
		UserDefinedInformation("Info check")
	assert.NotNil(params)
}

func TestValidationType(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS").
		IBAN("373BSJS").
		CustomerID("AISHD2992KKS").
		ValidationType("Info check")
	assert.NotNil(params)
}

func TestReferenceMask(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS").
		IBAN("373BSJS").
		CustomerID("AISHD2992KKS").
		ReferenceMask("&&&&&&&&&&&")
	assert.NotNil(params)
}

func TestAccepranceQualifier(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS").
		IBAN("373BSJS").
		CustomerID("AISHD2992KKS").
		AccepranceQualifier("myname")
	assert.NotNil(params)
}

func TestDone(t *testing.T) {
	assert := assert.New(t)
	params := NewAccountsParams().
		OrganisationID(uuid.NewString()).
		Type("accounts").
		Version(1).
		Country("GB").
		BaseCurrency("GBP").
		BankID("KJS382NS").
		IBAN("373BSJS").
		CustomerID("AISHD2992KKS").
		AccepranceQualifier("myname").
		Done()
	assert.NotNil(params)
}
