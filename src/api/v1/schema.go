package v1

// User contains user information
type User struct {
	FirstName string `validate:"required" json:"firstName"`
	LastName  string `validate:"required" json:"lastName"`
	Age       uint8  `validate:"gte=0,lte=130" json:"age"`
	Email     string `validate:"required,email" json:"email"`
}
