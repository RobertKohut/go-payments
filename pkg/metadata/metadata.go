package metadata

const (
	PaymentSourceStripe = 1
)

const (
	FlagsCardActive = 1 << iota
	FagsCardDefault
)

const (
	FlagsCustomerActive = 1 << iota
)

// HashId Service Constants
const (
	HDInvoiceId = 50
	HDChargeId  = 51
	HDCardId    = 52
)
