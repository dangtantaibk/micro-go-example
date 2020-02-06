
package loyalty

type ClassTracking struct {
	
        Id int64 `orm:"column(id)"`
                
        UserId string `orm:"column(user_id)"`
        
        Source string `orm:"column(source)"`
        
        OldClassId int64 `orm:"column(old_class_id)"`
        
        NewClassId int64 `orm:"column(new_class_id)"`
        
        NumOfPoint int64 `orm:"column(num_of_point)"`
        
        NumOfTrans int64 `orm:"column(num_of_trans)"`
        
        OldClassDate string `orm:"column(old_class_date)"`
        
        Created string `orm:"column(created)"`
        
        JsonDetail string `orm:"column(json_detail)"`
        }