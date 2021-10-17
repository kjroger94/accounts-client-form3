package accountsparams

import (
	"github.com/google/uuid"
	"github.com/kjroger94/accounts/models"
)

type DataParams interface {
	OrganisationID(string) DataParams
	Type(string) DataParams
	Version(int) DataParams
	Country(string) DataParams
	BaseCurrency(string) DataParams
	BankID(string) DataParams
	BankIDCode(string) DataParams
	AccountNumber(string) DataParams
	BIC(string) DataParams
	IBAN(string) DataParams
	CustomerID(string) DataParams
	Name(...string) DataParams
	AlternativeNames(...string) DataParams
	AccountClassification(string) DataParams
	JointAccount(bool) DataParams
	AccountMatchingOptOut(bool) DataParams
	SecondaryIdentification(string) DataParams
	Switched(bool) DataParams
	ProcessingService(string) DataParams
	UserDefinedInformation(string) DataParams
	ValidationType(string) DataParams
	ReferenceMask(string) DataParams
	AccepranceQualifier(string) DataParams

	Done() *models.Data
}

func NewAccountsParams() DataParams {
	return &params{
		Account: &models.Data{
			Account: &models.Account{
				ID:         uuid.NewString(),
				Attributes: &models.AccountAttributes{},
			},
		},
	}
}

type params struct {
	Account *models.Data
}

func (p *params) OrganisationID(s string) DataParams {
	p.Account.Account.OrganisationID = s
	return p
}

func (p *params) Type(s string) DataParams {
	p.Account.Account.Type = s
	return p
}

func (p *params) Version(i int) DataParams {
	p.Account.Account.Version = i
	return p
}

func (p *params) Country(s string) DataParams {
	p.Account.Account.Attributes.Country = s
	return p
}

func (p *params) BaseCurrency(s string) DataParams {
	p.Account.Account.Attributes.BaseCurrency = s
	return p
}

func (p *params) BankID(s string) DataParams {
	p.Account.Account.Attributes.BankID = s
	return p
}

func (p *params) BankIDCode(s string) DataParams {
	p.Account.Account.Attributes.BankIDCode = s
	return p
}

func (p *params) AccountNumber(s string) DataParams {
	p.Account.Account.Attributes.AccountNumber = s
	return p
}

func (p *params) BIC(s string) DataParams {
	p.Account.Account.Attributes.BIC = s
	return p
}

func (p *params) IBAN(s string) DataParams {
	p.Account.Account.Attributes.IBAN = s
	return p
}
func (p *params) CustomerID(s string) DataParams {
	p.Account.Account.Attributes.CustomerID = s
	return p
}

func (p *params) Name(s ...string) DataParams {
	p.Account.Account.Attributes.Name = s
	return p
}

func (p *params) AlternativeNames(s ...string) DataParams {
	p.Account.Account.Attributes.AlternativeNames = s
	return p
}

func (p *params) AccountClassification(s string) DataParams {
	p.Account.Account.Attributes.AccountClassification = s
	return p
}

func (p *params) JointAccount(b bool) DataParams {
	p.Account.Account.Attributes.JointAccount = b
	return p
}

func (p *params) AccountMatchingOptOut(b bool) DataParams {
	p.Account.Account.Attributes.AccountMatchingOptOut = b
	return p
}

func (p *params) SecondaryIdentification(s string) DataParams {
	p.Account.Account.Attributes.SecondaryIdentification = s
	return p
}

func (p *params) Switched(b bool) DataParams {
	p.Account.Account.Attributes.Switched = b
	return p
}

func (p *params) ProcessingService(s string) DataParams {
	p.Account.Account.Attributes.ProcessingService = s
	return p
}

func (p *params) UserDefinedInformation(s string) DataParams {
	p.Account.Account.Attributes.UserDefinedInformation = s
	return p
}

func (p *params) ValidationType(s string) DataParams {
	p.Account.Account.Attributes.ValidationType = s
	return p
}

func (p *params) ReferenceMask(s string) DataParams {
	p.Account.Account.Attributes.ReferenceMask = s
	return p
}

func (p *params) AccepranceQualifier(s string) DataParams {
	p.Account.Account.Attributes.AccepranceQualifier = s
	return p
}

func (p *params) Done() *models.Data {
	return p.Account
}
