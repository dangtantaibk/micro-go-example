package zalo

const (
	EndpointGetAccessToken         = "/v3/access_token"
	EndpointGetUserProfileOA       = "/v2.0/oa/getprofile"
	EndpointGetUserProfileOfficial = "/v2.0/me"
)

type GetAccessTokenRequest struct {
	AppID       string `json:"app_id"`
	AppSecret   string `json:"app_secret"`
	Code        string `json:"code"`
	RedirectUri string `json:"redirect_uri"`
	IsSDK       string `json:"is_sdk"`
}

type GetAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

type DataGetProfileRequest struct {
	UserID string `json:"user_id"`
}
type GetProfileOARequest struct {
	AccessToken string `json:"access_token"`
	UserID      string `json:"user_id"`
}
type GetProfileOAResponse struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
	Data    *struct {
		UserGender       int    `json:"user_gender"`
		UserID           int64  `json:"user_id"`
		UserIDByApp      int64  `json:"user_id_by_app"`
		Avatar           string `json:"avatar"`
		DisplayName      string `json:"display_name"`
		BirthDate        string `json:"birth_date"`
		SharedInfo       string `json:"shared_info"`
		TagsAndNotesInfo *struct {
			TagNames []string `json:"tag_names"`
			Notes    []string `json:"notes"`
		} `json:"tags_and_notes_info"`
	} `json:"data"`
}
type GetProfileOfficialRequest struct {
	AccessToken string `json:"access_token"`
}
type GetProfileOfficialResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Picture  *struct {
		Data *struct {
			URL string `json:"url"`
		} `json:"data"`
	} `json:"picture"`
}
