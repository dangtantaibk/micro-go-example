package dtos

type SignUpShipperRequest struct {
	MerchantCode string `json:"merchant_code"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	Status       int    `json:"status"`
}

type SignUpShipperResponse struct {
	SignUpSuccess bool `json:"sign_up_success"`
	Meta          Meta `json:"meta"`
}

type LoginWithPasswordShipperRequest struct {
	MerchantCode string `json:"merchant_code"`
	PhoneNumber  string `json:"phone_number"`
	Password     string `json:"password"`
}

type LoginWithPasswordShipperResponse struct {
	LoginSuccess bool   `json:"login_success"`
	Token        string `json:"token"`
	Meta         Meta   `json:"meta"`
}

type VerifyPhoneNumberShipperRequest struct {
	MerchantCode string `json:"merchant_code"`
	PhoneNumber  string `json:"phone_number"`
}

type VerifyPhoneNumberShipperResponse struct {
	VerifySuccess bool `json:"verify_success"`
	Meta          Meta `json:"meta"`
}

type RefreshTokenShipperRequest struct {
	OldToken string `json:"old_token"`
}

type RefreshTokenShipperResponse struct {
	NewToken string `json:"new_token"`
	Meta     Meta   `json:"meta"`
}

type UpdateShipperRequest struct {
	MerchantCode string `json:"merchant_code"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	Status       int    `json:"status"`
	ID           int64  `json:"id"`
}

type UpdateShipperResponse struct {
	Data ShipperInfo `json:"data"`
	Meta Meta        `json:"meta"`
}

type ListShipperRequest struct {
	MerchantCode      string `form:"merchant_code"`
	CurrentPage       int32  `form:"current_page"`
	TotalTransPerPage int32  `form:"total_trans_per_page"`
}

type ListShipperResponse struct {
	Meta        Meta          `json:"meta"`
	TotalRecord int64         `json:"total_record"`
	Data        []ShipperInfo `json:"data"`
}

type DeleteShipperAccountRequest struct {
	MerchantCode string `json:"merchant_code"`
	ID           int64  `json:"id"`
}

type DeleteShipperAccountResponse struct {
	Meta Meta `json:"meta"`
}

type ShipperInfo struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Status int    `json:"status"`
}
