package dtos

type UserSessionInfo struct {
	ID         int64  `json:"id"`
	AppID      int32  `json:"appid"`
	UCode      string `json:"ucode"`
	Title      string `json:"title"`
	FirstName  string `json:"firstname"`
	SurName    string `json:"surname"`
	FullName   string `json:"fullname"`
	LastName   string `json:"lastname"`
	Phone      string `json:"phone"`
	HomePhone  string `json:"home_phone"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	SocialID   string `json:"social_id"`
	LoginType  string `json:"login_type"`
	WardID     string `json:"ward_id"`
	DistrictID string `json:"district_id"`
	ProvinceID string `json:"province_id"`
	CountryID  string `json:"country_id"`
	Created    string `json:"created"`
	CreatedBy  string `json:"created_by"`
	Avatar     string `json:"avatar"`
	Status     string `json:"status"`
	Lat        string `json:"lat"`
	Long       string `json:"long"`
}

type UserSessionToken struct {
	Token string `json:"token"`
}

type GetUserSessionToken struct {
	Token string `json:"token"`
}

type UpdateUserSessionToken struct {
	Token       string           `json:"token"`
	UserSession *UserSessionInfo `json:"user_session"`
}
