package mp

const TokenSandbox string = "TEST-7101449558674620-051914-5c1f567bc1fc3c7314ae4df303bddf98-488839395"

type Credentials struct {
	TokenSandbox string
}
type Endpoints struct {
	Preferences    string
	PaymentMethods string
}

func NewEndpoints() *Endpoints {
	return &Endpoints{
		Preferences:    "https://api.mercadopago.com/checkout/preferences",
		PaymentMethods: "https://api.mercadopago.com/v1/payment_methods",
	}
}

type RequestFields struct {
	AdditionalInfo      string             `json:"additional_info"`
	AutoReturn          string             `json:"auto_return"`
	BackUrls            BackUrls           `json:"back_urls"`
	DateOfExpiration    string             `json:"date_of_expiration"`
	DifferencialPricing ReferenceIntegerID `json:"differencial_pricing"`
	ExpirationDateFrom  string             `json:"expiration_date_from"`
	ExpirationDateTo    string             `json:"expiration_date_to"`
	Expires             bool               `json:"expires"`
	ExternalReference   string             `json:"external_reference"`
	Items               []Items            `json:"items"`
	Metadata            Metadata           `json:"metadata"`
	NotificationUrl     string             `json:"notification_url"`
	Payer               Payer              `json:"payer"`
	PaymentMethods      PaymentMethods     `json:"payment_methods"`
	StatementDescriptor string             `json:"statement_descriptor"`
}

type BackUrls struct {
	Success string `json:"success"`
	Pending string `json:"pending"`
	Failure string `json:"failure"`
}

type ReferenceIntegerID struct {
	ID int `json:"id"`
}
type ReferenceStringID struct {
	ID string `json:"id"`
}
type Items struct {
	Id          int     `json:"id"`
	Title       string  `json:"string"`
	Description string  `json:"description"`
	PictureUrl  string  `json:"picture_url"`
	CategoryId  string  `json:"category_id"`
	Quantity    int     `json:"quantity"`
	CurrencyId  string  `json:"currency_id"`
	UnitPrice   float64 `json:"unit_price"`
}

type Metadata struct {
	Responsible string `json:"responsible"`
	Register    string `json:"register"`
}

type Payer struct {
	Responsible    string         `json:"responsible"`
	Register       string         `json:"register"`
	Name           string         `json:"name"`
	Surname        string         `json:"surname"`
	Email          string         `json:"email"`
	Phone          Phone          `json:"phone"`
	Identification Identification `json:"identification"`
	Address        Address        `json:"address"`
	DateCreated    string         `json:"date_created"`
}

type Phone struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"`
}

type Identification struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Address struct {
	ZipCode      string `json:"zipCode"`
	StreetName   string `json:"street_name"`
	StreetNumber string `json:"street_number"`
}

type PaymentMethods struct {
	ExcludedPaymentMethods []ReferenceStringID `json:"excluded_payment_methods"`
	ExcludedPaymentTypes   []ReferenceStringID `json:"excluded_payment_types"`
	DefaultPaymentMethodId string              `json:"default_payment_method_id"`
	Installments           int                 `json:"installments"`
	DefaultInstallments    int                 `json:"default_installments"`
}
