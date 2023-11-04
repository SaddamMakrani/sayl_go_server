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
	Customer User `json:"customer"`
}
type User struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	MobileNumber string `json:"mobile_number"`
}
