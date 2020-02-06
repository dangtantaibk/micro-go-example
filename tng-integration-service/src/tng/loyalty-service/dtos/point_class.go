
package dtos

type PointClassInfo struct {
        
        Id int64 `json:"id"`
                
        Title string `json:"title"`
        
        MinAccMoney int64 `json:"min_acc_money"`
        
        DiscountPercent float64 `json:"discount_percent"`
        
        RequireNumOfTrans int64 `json:"require_num_of_trans"`
        
        RequireNumOfPoint int64 `json:"require_num_of_point"`
        
        JsonDetail string `json:"json_detail"`
        
        Created string `json:"created"`
        }

type InsertOrUpdatePointClassRequest struct {
	
        Id int64 `json:"id"`
                
        Title string `json:"title"`
        
        MinAccMoney int64 `json:"min_acc_money"`
        
        DiscountPercent float64 `json:"discount_percent"`
        
        RequireNumOfTrans int64 `json:"require_num_of_trans"`
        
        RequireNumOfPoint int64 `json:"require_num_of_point"`
        
        JsonDetail string `json:"json_detail"`
        
        Created string `json:"created"`
        }

type InsertOrUpdatePointClassResponse struct {
	Meta Meta `json:"meta"`
}

type DeletePointClassRequest struct {
	
        Id int64 `json:"id"`
        }

type DeletePointClassResponse struct {
	Meta Meta `json:"meta"`
}

type ListPointClassRequest struct {
	PageIndex int32 `json:"page_index"`
	PageSize  int32 `json:"page_size"`
}

type ListPointClassResponse struct {
	Meta Meta           `json:"meta"`
	Data []*PointClassInfo `json:"data"`
}

type GetPointClassRequest struct {
	
        Id int64 `form:"id"`
        }

type GetPointClassResponse struct {
	Meta Meta         `json:"meta"`
	Data *PointClassInfo `json:"data"`
}
