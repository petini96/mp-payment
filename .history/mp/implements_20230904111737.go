package mp

import (
	"time"
)

type MPRequestInterface interface {
	GetMPAddress() Address
	GetMPPayer() Payer
	GetMPItems() []Items
	GetMPMetadata() Metadata
}

type MPRequestImplementer struct {
	Address  Address
	Payer    Payer
	Items    []Items
	Metadata Metadata
}

func NewMPRequestImplementer() *MPRequestImplementer {
	return &MPRequestImplementer{}
}

func (un *MPRequestImplementer) NewMPItems(id int, title string, description string, pictureUrl string, categoryID string, quantity int, currencyID string, unitPrice float64) []Items {
	item :=
		Items{
			Id:          id,
			Title:       title,
			Description: description,
			PictureUrl:  pictureUrl,
			CategoryId:  categoryID,
			Quantity:    quantity,
			CurrencyId:  currencyID,
			UnitPrice:   unitPrice,
		}
	un.Items = append(un.Items, item)
	return un.Items
}
func (un *MPRequestImplementer) GetMPItems() []Items {
	return un.Items
}

func (un *MPRequestImplementer) NewMetadata(responsible string, register string) Metadata {
	un.Metadata.Responsible = responsible
	un.Metadata.Register = register
	return un.Metadata
}

func (un *MPRequestImplementer) GetMPMetadata() Metadata {
	return un.Metadata
}
func (un *MPRequestImplementer) NewMPAddress(zipCode string, streetName string, streetNumber string) Address {
	un.Address.ZipCode = zipCode
	un.Address.StreetName = streetName
	un.Address.StreetNumber = streetNumber
	return un.Address
}
func (un *MPRequestImplementer) GetMPAddress() Address {
	return un.Address
}

func (un *MPRequestImplementer) NewMPPayer(id string, name string, email string, areaCode string, cellNumber string, identificationNumber string, createdAt time.Time) Payer {
	un.Payer.Responsible = name
	un.Payer.Register = id

	un.Payer.Phone.AreaCode = areaCode
	un.Payer.Phone.Number = cellNumber

	un.Payer.Identification.Type = "CPF"
	un.Payer.Identification.Number = identificationNumber

	un.Payer.Name = name
	un.Payer.Surname = "Sr. " + name
	un.Payer.Email = email
	un.Payer.DateCreated = ToIso8601(createdAt)
	return un.Payer
}
func (un *MPRequestImplementer) GetMPPayer() Payer {
	return un.Payer
}

func (rq *RequestFields) CompleteMPRequest(caller MPRequestInterface) {
	rq.Payer.Responsible = caller.GetMPPayer().Responsible
	rq.Payer.Register = caller.GetMPPayer().Register
	rq.Payer.Name = caller.GetMPPayer().Name
	rq.Payer.Surname = "Sr. " + caller.GetMPPayer().Name
	rq.Payer.Email = caller.GetMPPayer().Email
	rq.Payer.Phone.AreaCode = caller.GetMPPayer().Phone.AreaCode
	rq.Payer.Phone.Number = caller.GetMPPayer().Phone.Number
	rq.Payer.Identification.Type = caller.GetMPPayer().Identification.Type
	rq.Payer.Identification.Number = caller.GetMPPayer().Identification.Number

	rq.Payer.Address.ZipCode = caller.GetMPAddress().ZipCode
	rq.Payer.Address.StreetName = caller.GetMPAddress().StreetName
	rq.Payer.Address.StreetNumber = caller.GetMPAddress().StreetNumber

	rq.Payer.DateCreated = caller.GetMPPayer().DateCreated

	rq.Items = caller.GetMPItems()
	rq.Metadata = caller.GetMPMetadata()
}

func NewDefaultRequestFields(domain string) RequestFields {
	var rq RequestFields
	arrExcludedPaymentMethods := make([]ReferenceStringID, 1)
	arrExcludedPaymentMethods[0] = ReferenceStringID{"decbal"}

	arrExcludedPaymentTypes := make([]ReferenceStringID, 1)
	arrExcludedPaymentTypes[0] = ReferenceStringID{"debit_card"}

	rq.AutoReturn = "all"
	rq.BackUrls.Success = domain + "/payment/notification"
	rq.BackUrls.Pending = domain + "/payment/notification"
	rq.BackUrls.Failure = domain + "/payment/notification"
	rq.DateOfExpiration = ToIso8601(time.Now().AddDate(0, 0, 2))
	rq.ExpirationDateFrom = ToIso8601(time.Now())
	rq.ExpirationDateTo = ToIso8601(time.Now().AddDate(0, 0, 20))
	rq.Expires = true
	rq.ExternalReference = "777"
	rq.NotificationUrl = "https://webhook.site/33a2dc6e-1dd1-4fe3-9193-03181fc296f8"
	//rq.PaymentMethods.ExcludedPaymentMethods = arrExcludedPaymentMethods
	rq.PaymentMethods.ExcludedPaymentTypes = arrExcludedPaymentTypes
	rq.PaymentMethods.DefaultPaymentMethodId = "pix"
	rq.PaymentMethods.Installments = 10
	rq.PaymentMethods.DefaultInstallments = 1
	rq.StatementDescriptor = "ECOTRIP Tour"

	return rq
}
