package dtos

type Credentials struct {
	SocialId  string `json:"social_id"`
	Ip        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}
