package repositories

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"tng/common/models/loyalty"
	"tng/common/utils/db"
	"tng/common/utils/timeutil"
)

type PointRepository interface {
	InsertOrUpdate(context.Context, *db.DB, *loyalty.Point) error
	Delete(context.Context, *db.DB, int64) error
	GetByID(context.Context, *db.DB, int64) (*loyalty.Point, error)
	List(context.Context, *db.DB, int32, int32) ([]*loyalty.Point, error)
	Search(ctx context.Context, ormer *db.DB,
		userId string, pointType string, source string, forTransactionId string, notes string, appId string,
		pointStr string, transactionAmountStr string,	createdFrom string, createdTo string,
		statusStr string, promotionPercentStr string, rateStr string, sortColumn string, sortDirection string,
		pageIndex int32, pageSize int32) ([]*loyalty.Point, int64, error)
	PointHistory(context.Context, *db.DB, int32, int32, string) ([]*loyalty.Point, error)
	GetByTransactionID(context.Context, *db.DB, string, string, string) (*loyalty.Point, error)
}

type pointRepository struct{}

func NewPointRepository(dbFactory db.Factory) PointRepository {
	return &pointRepository{}
}

func (r *pointRepository) InsertOrUpdate(ctx context.Context, ormer *db.DB, dataInput *loyalty.Point) error {
	_, err := ormer.InsertOrUpdate(dataInput)
	return err
}

func (r *pointRepository) Delete(ctx context.Context, ormer *db.DB, id int64) error {
	qs := ormer.QueryTable(new(loyalty.Point))
	cond := orm.NewCondition().
		And("id", id)
	_, err := qs.SetCond(cond).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (r *pointRepository) GetByID(ctx context.Context, ormer *db.DB, id int64) (*loyalty.Point, error) {
	cond := orm.NewCondition().
		And("id", id)
	item := &loyalty.Point{}
	qs := ormer.QueryTable(new(loyalty.Point))
	err := qs.SetCond(cond).One(item)
	if err != nil {
		return nil, err
	}
	return item, nil

}

func (r *pointRepository) List(ctx context.Context, ormer *db.DB, pageIndex int32, pageSize int32) ([]*loyalty.Point, error) {
	var (
		list       []*loyalty.Point
		qs         = ormer.QueryTable(new(loyalty.Point))
		fromRecord int32
	)
	fromRecord = pageIndex*pageSize - pageSize
	qs = qs.Limit(pageSize, fromRecord)
	if _, err := qs.All(&list); err != nil {
		return nil, err
	}

	return list, nil
}

func (r *pointRepository) PointHistory(ctx context.Context, ormer *db.DB, pageIndex int32, pageSize int32, userId string) ([]*loyalty.Point, error) {
	var (
		cond       = orm.NewCondition()
		list       []*loyalty.Point
		qs         = ormer.QueryTable(new(loyalty.Point))
		fromRecord int32
	)
	fmt.Println("userid: ", userId)

	cond = cond.And("user_id", userId)
	fromRecord = pageIndex*pageSize - pageSize
	qs = qs.SetCond(cond)
	qs = qs.Limit(pageSize, fromRecord).OrderBy("-id")
	if _, err := qs.All(&list); err != nil {
		return nil, err
	}

	return list, nil
}

func (r *pointRepository) GetByTransactionID(ctx context.Context, ormer *db.DB, for_transaction_id string, app_id string, user_id string) (*loyalty.Point, error) {
	cond := orm.NewCondition().
		And("for_transaction_id", for_transaction_id).And("app_id", app_id).And("user_id", user_id)

	item := &loyalty.Point{}
	qs := ormer.QueryTable(new(loyalty.Point))
	err := qs.SetCond(cond).One(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}


func (r *pointRepository) Search(ctx context.Context, ormer *db.DB,
	userId string, pointType string, source string, forTransactionId string, notes string, appId string,
	pointStr string, transactionAmountStr string,	createdFrom string, createdTo string,
	statusStr string, promotionPercentStr string, rateStr string, sortColumn string, sortDirection string,
	pageIndex int32, pageSize int32) ([]*loyalty.Point, int64, error) {
	var (
		list       []*loyalty.Point
		qs         = ormer.QueryTable(new(loyalty.Point))
		fromRecord int32
		totalRecord int64
		layoutISO = "2006-1-02"
	)
	cond := orm.NewCondition();

	if userId != "" {
		cond = cond.And("user_id__icontains", userId)
	}
	if pointType != "" {
		cond = cond.And("point_type__icontains", pointType)
	}
	if source != "" {
		cond = cond.And("source__icontains", source)
	}
	if forTransactionId != "" {
		cond = cond.And("for_transaction_id__icontains", forTransactionId)
	}
	if notes != "" {
		cond = cond.And("notes__icontains", notes)
	}
	if appId != "" {
		cond = cond.And("app_id__icontains", appId)
	}


	if pointStr != "" {
		point, err := strconv.ParseInt(pointStr, 10, 64)
		if err == nil {
			cond = cond.And("point", point)
		}
	}
	if transactionAmountStr != "" {
		transactionAmount, err := strconv.ParseInt(transactionAmountStr, 10, 64)
		if err == nil {
			cond = cond.And("transaction_amount", transactionAmount)
		}
	}
	if createdFrom != "" {
		createdFromDate := timeutil.ParseTime(createdFrom, layoutISO)
		if createdFromDate != nil {
			cond = cond.And("created__gte", createdFromDate)
		}
	}
	if createdTo != "" {
		createdToDate := timeutil.ParseTime(createdTo, layoutISO)
		if createdToDate != nil {
			cond = cond.And("created__lte", createdToDate)
		}
	}
	if statusStr != "" {
		status, err := strconv.ParseInt(statusStr, 10, 64)
		if err == nil {
			cond = cond.And("status", status)
		}
	}
	if promotionPercentStr != "" {
		promotionPercent, err := strconv.ParseFloat(statusStr, 64)
		if err == nil {
			cond = cond.And("promotion_percent", promotionPercent)
		}
	}
	if rateStr != "" {
		rate, err := strconv.ParseFloat(rateStr, 64)
		if err == nil {
			cond = cond.And("rate", rate)
		}
	}
	qs1 := qs.SetCond(cond)
	totalRecord, err := qs1.Count()
	if err != nil {
		return nil, 0, err
	}

	fromRecord = pageIndex*pageSize - pageSize
	qs2 := qs.SetCond(cond)
	qs2 = qs2.Limit(pageSize, fromRecord)

	if sortColumn != "" && sortDirection != "" {
		if strings.ToLower(sortDirection) == "asc" {
			qs2 = qs2.OrderBy(sortColumn)
		}
		if strings.ToLower(sortDirection) == "desc" {
			qs2 = qs2.OrderBy("-"+sortColumn)
		}
	}
	if _, err := qs2.All(&list); err != nil {
		return nil, 0, err
	}

	return list, totalRecord, nil
}