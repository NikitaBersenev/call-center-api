package domain

type CallCreateRequest struct {
	ClientName  string `json:"client_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required,phonenumber"`
	Description string `json:"description" binding:"required"`
}
