
package dtos

type PointTypeInfo struct {
        
        Id string `json:"id"`
                
        Description string `json:"description"`
        
        JsonDetail string `json:"json_detail"`
        
        Created string `json:"created"`
        }

type InsertOrUpdatePointTypeRequest struct {
	
        Id string `json:"id"`
                
        Description string `json:"description"`
        
        JsonDetail string `json:"json_detail"`
        
        Created string `json:"created"`
        }

type InsertOrUpdatePointTypeResponse struct {
	Meta Meta `json:"meta"`
}

type DeletePointTypeRequest struct {
	
        Id string `json:"id"`
        }

type DeletePointTypeResponse struct {
	Meta Meta `json:"meta"`
}

type ListPointTypeRequest struct {
	PageIndex int32 `form:"page_index"`
	PageSize  int32 `form:"page_size"`
}

type ListPointTypeResponse struct {
	Meta Meta           `json:"meta"`
	Data []*PointTypeInfo `json:"data"`
}

type GetPointTypeRequest struct {
	
        Id string `form:"id"`
        }

type GetPointTypeResponse struct {
	Meta Meta         `json:"meta"`
	Data *PointTypeInfo `json:"data"`
}
