package entities

type Customer struct {
	Id        int64  `json:"id"`
	GatewayId int64  `json:"gateway_id"`
	SourceId  int64  `json:"source_id"`
	OrgId     int64  `json:"org_id"`
	AccountId int64  `json:"account_id"`
	ExtId     string `json:"ext_id"`
	Name      string `json:"name,omitempty"`
	Cards     []Card `json:"cards,omitempty"`
	Flags     int    `json:"flags"`
}

type Card struct {
	Id         int64  `json:"-"`
	ExtId      string `json:"id"`
	CustomerId int64  `json:"customer_id"`
	ExpMonth   int    `json:"exp_month"`
	ExpYear    int    `json:"exp_year"`
	LastFour   string `json:"last_four"`
}
