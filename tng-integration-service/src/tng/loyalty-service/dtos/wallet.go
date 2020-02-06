
package dtos

type WalletInfo struct {
        
        Id int64 `json:"id"`
                
        UserId string `json:"user_id"`
        
        Balance float64 `json:"balance"`
        
        TotalIn float64 `json:"total_in"`
        
        TotalOut float64 `json:"total_out"`
        
        BalancePromo float64 `json:"balance_promo"`
        
        TotalInPromo float64 `json:"total_in_promo"`
        
        TotalOutPromo float64 `json:"total_out_promo"`
        
        Created string `json:"created"`
        
        Modified string `json:"modified"`
        
        Checksum string `json:"checksum"`
        
        ClassId int64 `json:"class_id"`
        
        ClassDate string `json:"class_date"`
        
        Status int32 `json:"status"`
        
        AccMoney int64 `json:"acc_money"`
        
        PointDate string `json:"point_date"`
        }

type InsertOrUpdateWalletRequest struct {
	
        Id int64 `json:"id"`
                
        UserId string `json:"user_id"`
        
        Balance float64 `json:"balance"`
        
        TotalIn float64 `json:"total_in"`
        
        TotalOut float64 `json:"total_out"`
        
        BalancePromo float64 `json:"balance_promo"`
        
        TotalInPromo float64 `json:"total_in_promo"`
        
        TotalOutPromo float64 `json:"total_out_promo"`
        
        Created  string `json:"created"`
        
        Modified string `json:"modified"`
        
        Checksum string `json:"checksum"`
        
        ClassId int64 `json:"class_id"`
        
        ClassDate string `json:"class_date"`
        
        Status int32 `json:"status"`
        
        AccMoney int64 `json:"acc_money"`
        
        PointDate string `json:"point_date"`
        }

type InsertOrUpdateWalletResponse struct {
	Meta Meta `json:"meta"`
}

type DeleteWalletRequest struct {
	
        Id int64 `json:"id"`
        }

type DeleteWalletResponse struct {
	Meta Meta `json:"meta"`
}

type ListWalletRequest struct {
	PageIndex int32 `form:"page_index"`
	PageSize  int32 `form:"page_size"`
}

type ListWalletResponse struct {
	Meta Meta           `json:"meta"`
	Data []*WalletInfo `json:"data"`
}

type GetWalletRequest struct {
	
        Id int64 `form:"id"`
        }

type GetWalletResponse struct {
	Meta Meta         `json:"meta"`
	Data *WalletInfo `json:"data"`
}
