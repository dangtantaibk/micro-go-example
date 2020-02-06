
package loyalty

type PointType struct {
	
        Id string `orm:"column(id);pk"`
                
        Description string `orm:"column(description)"`
        
        JsonDetail string `orm:"column(json_detail)"`
        
        Created string `orm:"column(created)"`
        }