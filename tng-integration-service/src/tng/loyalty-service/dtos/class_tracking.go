
package dtos

type ClassTrackingInfo struct {
        
        Id int64 `json:"id"`
                
        UserId string `json:"user_id"`
        
        Source string `json:"source"`
        
        OldClassId int64 `json:"old_class_id"`
        
        NewClassId int64 `json:"new_class_id"`
        
        NumOfPoint int64 `json:"num_of_point"`
        
        NumOfTrans int64 `json:"num_of_trans"`
        
        OldClassDate string `json:"old_class_date"`
        
        Created string `json:"created"`
        
        JsonDetail string `json:"json_detail"`
        }

type InsertOrUpdateClassTrackingRequest struct {
	
        Id int64 `json:"id"`
                
        UserId string `json:"user_id"`
        
        Source string `json:"source"`
        
        OldClassId int64 `json:"old_class_id"`
        
        NewClassId int64 `json:"new_class_id"`
        
        NumOfPoint int64 `json:"num_of_point"`
        
        NumOfTrans int64 `json:"num_of_trans"`
        
        OldClassDate string `json:"old_class_date"`
        
        Created string `json:"created"`
        
        JsonDetail string `json:"json_detail"`
        }

type InsertOrUpdateClassTrackingResponse struct {
	Meta Meta `json:"meta"`
}

type DeleteClassTrackingRequest struct {
	
        Id int64 `json:"id"`
        }

type DeleteClassTrackingResponse struct {
	Meta Meta `json:"meta"`
}

type ListClassTrackingRequest struct {
	PageIndex int32 `form:"page_index"`
	PageSize  int32 `form:"page_size"`
}

type ListClassTrackingResponse struct {
	Meta Meta           `json:"meta"`
	Data []*ClassTrackingInfo `json:"data"`
}

type GetClassTrackingRequest struct {
	
        Id int64 `form:"id"`
        }

type GetClassTrackingResponse struct {
	Meta Meta         `json:"meta"`
	Data *ClassTrackingInfo `json:"data"`
}
