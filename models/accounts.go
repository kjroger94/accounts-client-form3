package models

type Data struct {
	Account *Account `json:"data,omitempty"`
}

type Account struct {
	ID             string `json:"id,omitempty"`
	OrganisationID string `json:"organisation_id,omitempty"`
	Type           string `json:"type,omitempty"`
	Version        int    `json:"version,omitempty"`
	Attributes     *AccountAttributes
}

type AccountAttributes struct {
	Country                   string                     `json:"country,omitempty"`
	BaseCurrency              string                     `json:"base_currency,omitempty"`
	BankID                    string                     `json:"bank_id,omitempty"`
	BankIDCode                string                     `json:"bank_id_code,omitempty"`
	AccountNumber             string                     `json:"account_number,omitempty"`
	BIC                       string                     `json:"bic,omitempty"`
	IBAN                      string                     `json:"iban,omitempty"`
	CustomerID                string                     `json:"customer_id,omitempty"`
	Name                      []string                   `json:"name,omitempty"`
	AlternativeNames          []string                   `json:"alternative_names,omitempty"`
	AccountClassification     string                     `json:"account_classification,omitempty"`
	JointAccount              bool                       `json:"joint_account,omitempty"`
	AccountMatchingOptOut     bool                       `json:"account_matching_opt_out,omitempty"`
	SecondaryIdentification   string                     `json:"secondary_identification,omitempty"`
	Switched                  bool                       `json:"switched,omitempty"`
	ProcessingService         string                     `json:"processing_service,omitempty"`
	UserDefinedInformation    string                     `json:"user_defined_information,omitempty"`
	ValidationType            string                     `json:"validation_type,omitempty"`
	ReferenceMask             string                     `json:"reference_mask,omitempty"`
	AccepranceQualifier       string                     `json:"acceptance_qualifier,omitempty"`
	PrivateIdentification     *PrivateIdentification     `json:"private_identification,omitempty"`     //TBC
	OrganisationIDentifcation *OrganisationIDentifcation `json:"organization_identifcation,omitempty"` //TBC
}

type PrivateIdentification struct { //TBC
	BirthDate    string `json:"birth_date,omitempty"`
	BirthCountry string `json:"birth_country,omitempty"`
	Identification
}

type OrganisationIDentifcation struct { //TBC
	Actors
	Identification
}

type Identification struct { //TBC
	Identification string   `json:"identification,omitempty"`
	Address        []string `json:"address,omitempty"`
	City           string   `json:"city,omitempty"`
	Country        string   `json:"country,omitempty"`
}

type Actors struct {
	Name      []string `json:"name,omitempty"`
	BirthDate string   `json:"birth_date,omitempty"`
	Residency string   `json:"residency,omitempty"`
}
