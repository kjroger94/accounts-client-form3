package testing

import (
	"strings"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/kjroger94/accounts/cmd/accountsparams"
)

func createBritishAccountData() accountsparams.DataParams {
	return accountsparams.NewAccountsParams().
		OrganisationID(fake.UUID()).
		Country("GB").
		BankID(getBankID(6)).
		BankIDCode("GBDSC").
		Type("accounts").
		Name(fake.Name()).
		BIC("NWBKGB22")

}

func createAustralianAccountData() accountsparams.DataParams {
	return accountsparams.NewAccountsParams().
		OrganisationID(fake.UUID()).
		Country("AU").
		BankID(getBankID(6)).
		BankIDCode("AUBSB").
		Type("accounts").
		Name(fake.Name()).
		BIC("AUMKGB22")

}

func createBelgiumAccountData() accountsparams.DataParams {
	return accountsparams.NewAccountsParams().
		OrganisationID(fake.UUID()).
		Country("BE").
		BankID(getBankID(6)).
		BankIDCode("BE").
		Type("accounts").
		Name(fake.Name()).
		BIC("BELSGB22")
}

func createCanadaAccountData() accountsparams.DataParams {
	return accountsparams.NewAccountsParams().
		OrganisationID(fake.UUID()).
		Country("CA").
		BankID(getBankID(9)).
		BankIDCode("CACPA").
		Type("accounts").
		Name(fake.Name()).
		BIC("BCANAB22")
}

func createEstoniaAccountData() accountsparams.DataParams {
	return accountsparams.NewAccountsParams().
		OrganisationID(fake.UUID()).
		Country("EE").
		BankID(getBankID(4)).
		BankIDCode("EE").
		Type("accounts").
		Name(fake.Name()).
		BIC("RESTAB32")
}

func createFrenchAccountData() accountsparams.DataParams {
	return accountsparams.NewAccountsParams().
		OrganisationID(fake.UUID()).
		Country("FR").
		BankID(getBankID(10)).
		BankIDCode("FR").
		Type("accounts").
		Name(fake.Name()).
		BIC("RFREAB32")
}

func createGermanAccountData() accountsparams.DataParams {
	return accountsparams.NewAccountsParams().
		OrganisationID(fake.UUID()).
		Country("DE").
		BankID(getBankID(8)).
		BankIDCode("DEBLZ").
		Type("accounts").
		Name(fake.Name()).
		BIC("GEREAB32")
}

func createGreeceAccountData() accountsparams.DataParams {
	return accountsparams.NewAccountsParams().
		OrganisationID(fake.UUID()).
		Country("GR").
		BankID(getBankID(7)).
		BankIDCode("GRBIC").
		Type("accounts").
		Name(fake.Name()).
		BIC("GEREAB32")
}

func createHKAccountData() accountsparams.DataParams {
	return accountsparams.NewAccountsParams().
		OrganisationID(fake.UUID()).
		Country("HK").
		BankID(getBankID(3)).
		BankIDCode("HKNCC").
		Type("accounts").
		Name(fake.Name()).
		BIC("KONGAB32")
}

func createIrishAccountData() accountsparams.DataParams {
	return accountsparams.NewAccountsParams().
		OrganisationID(fake.UUID()).
		Country("IE").
		BankID(getBankID(6)).
		BankIDCode("IENCC").
		Type("accounts").
		Name(fake.Name()).
		BIC("IREEAB32")
}

func createItalianAccountData() accountsparams.DataParams {
	return accountsparams.NewAccountsParams().
		OrganisationID(fake.UUID()).
		Country("IT").
		BankID(getBankID(11)).
		BankIDCode("ITNCC").
		Type("accounts").
		Name(fake.Name()).
		BIC("IITAAB32")
}

func createLuxAccountData() accountsparams.DataParams {
	return accountsparams.NewAccountsParams().
		OrganisationID(fake.UUID()).
		Country("LU").
		BankID(getBankID(3)).
		BankIDCode("LULUX").
		Type("accounts").
		Name(fake.Name()).
		BIC("LUXAAB32")
}
func createNetherlandsAccountData() accountsparams.DataParams {
	return accountsparams.NewAccountsParams().
		OrganisationID(fake.UUID()).
		Country("NL").
		Type("accounts").
		Name(fake.Name()).
		BIC("DUTAAB32")
}

func createPolandAccountData() accountsparams.DataParams {
	return accountsparams.NewAccountsParams().
		OrganisationID(fake.UUID()).
		Country("PL").
		BankID(getBankID(8)).
		BankIDCode("PLKNR").
		Type("accounts").
		Name(fake.Name()).
		BIC("DUPOLB32")
}

func createPortugalAccountData() accountsparams.DataParams {
	return accountsparams.NewAccountsParams().
		OrganisationID(fake.UUID()).
		Country("PT").
		BankID(getBankID(8)).
		BankIDCode("PTNCC").
		Type("accounts").
		Name(fake.Name()).
		BIC("POROLB32")
}

func createSpainAccountData() accountsparams.DataParams {
	return accountsparams.NewAccountsParams().
		OrganisationID(fake.UUID()).
		Country("ES").
		BankID(getBankID(8)).
		BankIDCode("ESNCC").
		Type("accounts").
		Name(fake.Name()).
		BIC("ESPOLB32")
}

func createSwissAccountData() accountsparams.DataParams {
	return accountsparams.NewAccountsParams().
		OrganisationID(fake.UUID()).
		Country("CH").
		BankID(getBankID(5)).
		BankIDCode("CHBCC").
		Type("accounts").
		Name(fake.Name()).
		BIC("SWIOLB32")
}

func createUSAAccountData() accountsparams.DataParams {
	return accountsparams.NewAccountsParams().
		OrganisationID(fake.UUID()).
		Country("US").
		BankID(getBankID(5)).
		BankIDCode("USABA").
		Type("accounts").
		Name(fake.Name()).
		BIC("USAOLB32")
}

func getBankID(n uint) string {
	return strings.ToUpper(fake.LetterN(n))
}
