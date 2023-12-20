package models

type RegisBody struct {
	Email       string `json:"email" validate:"required,email"`
	Username    string `json:"username" validate:"required,excludesall= !@#?.*"`
	LineId      string `json:"lineId" validate:"required"`
	PhoneNumber string `json:"phone" validate:"required,numeric,len=10"`
	BType       string `json:"btype" validate:"required"`
	WUrl        string `json:"wUrl" validate:"required,min=2,max=30,lowercase,excludesall= !_@#?.*"`
}
