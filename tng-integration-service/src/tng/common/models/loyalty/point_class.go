
package loyalty

type PointClass struct {
	
        Id int64 `orm:"column(id)"`
                
        Title string `orm:"column(title)"`
        
        MinAccMoney int64 `orm:"column(min_acc_money)"`
        
        DiscountPercent float64 `orm:"column(discount_percent)"`
        
        RequireNumOfTrans int64 `orm:"column(require_num_of_trans)"`
        
        RequireNumOfPoint int64 `orm:"column(require_num_of_point)"`
        
        JsonDetail string `orm:"column(json_detail)"`
        
        Created string `orm:"column(created)"`
        }