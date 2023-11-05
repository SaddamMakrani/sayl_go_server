package structs

type Products struct {
	Product []Product `json:"products"`
}
type Product struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	Status   string    `json:"status"`
	Variants []Variant `json:"variants"`
}

type Variant struct {
	ID        int64  `json:"id"`
	ProductID int64  `json:"product_id"`
	Title     string `json:"title"`
	Price     string `json:"price"`
}

type Customer struct {
	User `json:"customer"`
}
type User struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	MobileNumber string `json:"mobile_number"`
}

type CreatedUser struct {
	CreateUser CreateUser `json:"customer"`
}

type CreateUser struct {
	ID                    int64  `json:"id"`
	Email                 string `json:"email"`
	AcceptsMarketing      bool   `json:"accepts_marketing"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	OrdersCount           int    `json:"orders_count"`
	State                 string `json:"state"`
	TotalSpent            string `json:"total_spent"`
	LastOrderID           string `json:"last_order_id"`
	Note                  string `json:"note"`
	VerifiedEmail         bool   `json:"verified_email"`
	MultipassIdentifier   string `json:"multipass_identifier"`
	TaxExempt             bool   `json:"tax_exempt"`
	Tags                  string `json:"tags"`
	LastOrderName         string `json:"last_order_name"`
	Currency              string `json:"currency"`
	Phone                 string `json:"phone"`
	EmailMarketingConsent string `json:"email_marketing_consent"`
	SMSMarketingConsent   string `json:"sms_marketing_consent"`
	AdminGraphQLAPIID     string `json:"admin_graphql_api_id"`
}
