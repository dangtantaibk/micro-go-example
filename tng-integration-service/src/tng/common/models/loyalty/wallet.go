
package loyalty

type Wallet struct {
	
        Id int64 `orm:"column(id)"`
                
        UserId string `orm:"column(user_id)"`
        
        Balance float64 `orm:"column(balance)"`
        
        TotalIn float64 `orm:"column(total_in)"`
        
        TotalOut float64 `orm:"column(total_out)"`
        
        BalancePromo float64 `orm:"column(balance_promo)"`
        
        TotalInPromo float64 `orm:"column(total_in_promo)"`
        
        TotalOutPromo float64 `orm:"column(total_out_promo)"`
        
        Created string `orm:"column(created)"`
        
        Modified string `orm:"column(modified)"`
        
        Checksum string `orm:"column(checksum)"`
        
        ClassId int64 `orm:"column(class_id)"`
        
        ClassDate string `orm:"column(class_date)"`
        
        Status int32 `orm:"column(status)"`
        
        AccMoney int64 `orm:"column(acc_money)"`
        
        PointDate string `orm:"column(point_date)"`
        }