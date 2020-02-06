package dtos

const (
	LOGIN_TYPE_EMAIL         = 1
	LOGIN_TYPE_FB            = 2
	LOGIN_TYPE_GOOGLE        = 3
	LOGIN_TUPE_OPENID        = 4
	LOGIN_TYPE_PHONE         = 5
	LOGIN_TYPE_USERNAME      = 6
	LOGIN_TYPE_ZALO_OA       = 7
	LOGIN_TYPE_ZALO_OFFICIAL = 8
	LOGIN_TYPE_ZALO_ZPI      = 9
)

type User struct {
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

type CreateProfileRequest struct {
	Meta *MetaRequest `json:"meta"`
	Data *User        `json:"data"`
}

type CreateProfileResponse struct {
	Meta Meta `json:"meta"`
}

type GetProfileByIDRequest struct {
	UserID    int64  `form:"user_id"`
	AppID     int32  `form:"app_id"`
	Timestamp int64  `form:"timestamp"`
	Sig       string `form:"sig"`
}

type GetProfileUCodeRequest struct {
	UCode     int64  `form:"u_code"`
	AppID     int32  `form:"app_id"`
	Timestamp int64  `form:"timestamp"`
	Sig       string `form:"sig"`
}

type GetProfileResponse struct {
	Meta Meta  `json:"meta"`
	Data *User `json:"data"`
}

type UpdateProfileRequest struct {
	Meta *MetaRequest `json:"meta"`
	Data *User        `json:"data"`
}

type UpdateProfileResponse struct {
	Meta Meta `json:"meta"`
}

// Sig for SHA256, DataInput: SocialID + | + AppID + | + Platform + | + TS + <Client-Key>
type DataLoginRequest struct {
	SocialID string `json:"social_id"`
	Platform string `json:"platform"`
	AppID    int32  `json:"app_id"` // appid for zalopay
	TS       int64  `json:"ts"`
	Sig      string `json:"sig"`
}

// Data for login -> DataLoginRequest
type LoginRequest struct {
	LoginType int    `json:"login_type"`
	Data      string `json:"data"`
}

type LoginResponseData struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Meta Meta               `json:"meta"`
	Data *LoginResponseData `json:"data"`
}

type CheckLoginRequest struct {
	Token string `json:"token"`
}

type CheckLoginResponse struct {
	Meta Meta `json:"meta"`
}
