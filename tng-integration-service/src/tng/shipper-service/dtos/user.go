package dtos

type LoginRequest struct {
	MerchantCode string `json:"merchant_code"`
	UserId       int64  `json:"user_id"`
	Code         string `json:"code"`
	RedirectURI  string `json:"redirect_uri"`
	AppID        string `json:"app_id"`
}

type LoginResponse struct {
	LoginSuccess bool   `json:"login_success"`
	Token        string `json:"token"`
	Meta         Meta   `json:"meta"`
}

type SignUpRequest struct {
	MerchantCode string `json:"merchant_code"`
	Password     string `json:"password"`
	SocialId     string `json:"social_id"`
	SocialName   string `json:"social_name"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phone_number"`
}

type SignUpResponse struct {
	SignUpSuccess bool `json:"sign_up_success"`
	Meta          Meta `json:"meta"`
}

type LoginWithPasswordRequest struct {
	MerchantCode string `json:"merchant_code"`
	PhoneNumber string `json:"phone_number"`
	Password string `json:"password"`
}

type LoginWithPasswordResponse struct {
	LoginSuccess bool `json:"login_success"`
	Token string `json:"token"`
	Meta Meta `json:"meta"`
}

type VerifyPhoneNumberRequest struct {
	MerchantCode string `json:"merchant_code"`
	PhoneNumber string `json:"phone_number"`
}

type VerifyPhoneNumberResponse struct {
	VerifySuccess bool `json:"verify_success"`
	Meta Meta `json:"meta"`
}

type RefreshTokenRequest struct {
	OldToken string `json:"old_token"`
}

type RefreshTokenResponse struct {
	NewToken string `json:"new_token"`
	Meta Meta `json:"meta"`
}
