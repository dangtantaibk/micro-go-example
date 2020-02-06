package dtos

type Merchant struct {
	MerchantID        int64  `json:"merchant_id"`
	MerchantCode      string `json:"merchant_code"`
	MerchantType      int    `json:"merchant_type"`
	MerchantMame      string `json:"merchant_mame"`
	BranchName        string `json:"branch_name"`
	MerchantGroupID   int64  `json:"merchant_group_id"`
	MerchantGroupName string `json:"merchant_group_name"`
	AppUser           string `json:"app_user"`
	AppID             int64  `json:"app_id"`
	Key1              string `json:"key1"`
	Key2              string `json:"key2"`
	ZpsUrl            string `json:"zps_url"`
	Description       string `json:"description"`
	RsaPrivate        string `json:"rsa_private"`
	RsaPublic         string `json:"rsa_public"`
}
