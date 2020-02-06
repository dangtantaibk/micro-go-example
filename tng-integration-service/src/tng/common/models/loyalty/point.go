
package loyalty

type Point struct {
	
        Id int64 `orm:"column(id)"`
                
        UserId string `orm:"column(user_id)"`
        
        PointType string `orm:"column(point_type)"`
        
        Point int64 `orm:"column(point)"`
        
        Source string `orm:"column(source)"`
        
        ForTransactionId string `orm:"column(for_transaction_id)"`
        
        TransactionAmount int64 `orm:"column(transaction_amount)"`
        
        Notes string `orm:"column(notes)"`
        
        Created string `orm:"column(created)"`
        
        CreatedYmd string `orm:"column(created_ymd)"`
        
        Status int32 `orm:"column(status)"`
        
        AppId string `orm:"column(app_id)"`
        
        PromotionPercent float64 `orm:"column(promotion_percent)"`
        
        CampaignCode string `orm:"column(campaign_code)"`
        
        Channel string `orm:"column(channel)"`
        
        Rate float64 `orm:"column(rate)"`
        
        JsonDetail string `orm:"column(json_detail)"`
        }