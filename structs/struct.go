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
	ID              int64  `json:"id"`
	ProductID       int64  `json:"product_id"`
	Title           string `json:"title"`
	Price           string `json:"price"`
	InventoryItemId int64  `json:"inventory_item_id"`
}

type VariantResponse struct {
	Variant Variant `json:"variant"`
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

type CartItem struct {
	VariantID int64 `json:"variant_id"`
	Quantity  int   `json:"quantity"`
}
type ViewCartItem struct {
	VariantID    int64  `json:"variant_id"`
	VariantTitle string `json:"variant_title"`
}

type UserCart struct {
	Items []CartItem
}

type OrderResponse struct {
	OrderID string `json:"order_id"`
}

type ShopifyOrderRequest struct {
	LineItems []CartItem `json:"line_items"`
}

type ShopifyOrder struct {
	Order Order `json:"order"`
}

type Order struct {
	ID                    int64  `json:"id"`
	AdminGraphQLAPIID     string `json:"admin_graphql_api_id"`
	AppID                 int64  `json:"app_id"`
	BrowserIP             string `json:"browser_ip"`
	BuyerAcceptsMarketing bool   `json:"buyer_accepts_marketing"`
	CancelReason          string `json:"cancel_reason"`
	CancelledAt           string `json:"cancelled_at"`
	CartToken             string `json:"cart_token"`
	CheckoutID            string `json:"checkout_id"`
}
type ShopifyOrderResponseByUtm struct {
	Order []Order `json:"orders"`
}
