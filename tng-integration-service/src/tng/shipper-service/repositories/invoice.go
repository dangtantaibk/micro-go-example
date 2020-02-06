package repositories

import (
	"github.com/astaxie/beego/orm"
	"github.com/jinzhu/copier"
	modelShipper "tng/common/models/shiper"
	"tng/common/utils/db"
	"tng/shipper-service/dtos"
)

type InvoiceRepository interface {
	List(ormer *db.DB, paymentDate string, phoneNumber string, current int32, totalPerPage int32) ([] *modelShipper.Invoice, int64, int64, error)
	AllList(ormer *db.DB, paymentDate string, current int32, totalPerPage int32) ([] *modelShipper.Invoice, int64, error)
	Update(ormer *db.DB, request dtos.UpdateInvoiceStatusRequest) error
	GetByVposToken(ormer *db.DB, vposToken string) (*modelShipper.Invoice, [] *dtos.InvoiceDetail, error)
	GetInvoiceDetail(ormer *db.DB, invoiceCode string) (*modelShipper.Invoice, [] *dtos.InvoiceDetail, error)
}

type invoiceRepository struct{

}

func (r *invoiceRepository) GetByVposToken(ormer *db.DB, vposToken string) (*modelShipper.Invoice, [] *dtos.InvoiceDetail, error) {
	cond := orm.NewCondition().
		And("vpostoken", vposToken)
	condGeneral := orm.NewCondition().And("name", "img_host")
	invoiceInfo := &modelShipper.Invoice{}
	general := &modelShipper.General{}

	qs := ormer.QueryTable(new(modelShipper.Invoice))
	qsGeneral := ormer.QueryTable(new(modelShipper.General))
	err := qs.SetCond(cond).One(invoiceInfo)
	if err != nil {
		return nil, nil, err
	}
	errGeneral := qsGeneral.SetCond(condGeneral).One(general)
	if errGeneral != nil {
		return nil, nil, errGeneral
	}

	invoiceId := invoiceInfo.ID
	invoiceIdCond := orm.NewCondition().And("invoice_id", invoiceId)
	var invoiceDetailList []*modelShipper.InvoiceDetail
	qs1 := ormer.QueryTable(new(modelShipper.InvoiceDetail))
	_, err = qs1.SetCond(invoiceIdCond).All(&invoiceDetailList)
	if err != nil {
		return nil, nil, err
	}
	qsItem := ormer.QueryTable(new(modelShipper.Item))

	listInvoiceDetailDto := make([]*dtos.InvoiceDetail, 0)
	for _, item := range invoiceDetailList  {
		var (
			iv dtos.InvoiceDetail
			_  = copier.Copy(&iv, &item)
		)
		itemInfo := &modelShipper.Item{}
		condItem := orm.NewCondition().And("item_id", item.ItemID)
		err = qsItem.SetCond(condItem).One(itemInfo)
		if err != nil {
			return nil, nil, err
		}
		iv.ImgPath = general.Value + itemInfo.ImgPath
		listInvoiceDetailDto = append(listInvoiceDetailDto, &iv)
	}

	return invoiceInfo, listInvoiceDetailDto, nil


}

func (r *invoiceRepository) Update(ormer *db.DB, request dtos.UpdateInvoiceStatusRequest) error {

	qs := ormer.QueryTable(new(modelShipper.Invoice))
	cond := orm.NewCondition().
		And("invoice_code", request.InvoiceCode)
	var invoiceInfo modelShipper.Invoice
	err := qs.SetCond(cond).One(&invoiceInfo)
	if err == nil {
		invoiceInfo.PaymentStatus = request.PaymentStatus
		_, err := ormer.Update(&invoiceInfo)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *invoiceRepository) List(ormer *db.DB, paymentDate string, phoneNumber string, current int32, totalPerPage int32) ([] *modelShipper.Invoice, int64, int64, error) {
	var (
		list []*modelShipper.Invoice
		qs   = ormer.QueryTable(new(modelShipper.Invoice))
		cond = orm.NewCondition().And("shiper_phone", phoneNumber)
		condAll = orm.NewCondition()
		fromRecord int32
		totalRecord int64
		totalAllRecord int64
	)
	fromRecord = current * totalPerPage - totalPerPage
	qs1 := qs.SetCond(cond).
		Filter("payment_date_time__istartswith", paymentDate).
		Filter("payment_status__in", 11, 12, 13, 14).
		OrderBy("-invoice_code")
	qs2 := qs.SetCond(cond).
		Filter("payment_date_time__istartswith", paymentDate).
		Filter("payment_status__in", 11, 12, 13, 14)
	qs3 := qs.SetCond(condAll).
		Filter("payment_date_time__istartswith", paymentDate).
		Filter("payment_status__in", 11, 12, 13, 14)

	totalRecord, err := qs2.Count()
	if err != nil {
		return nil, 0, 0, err
	}

	totalAllRecord, err1 := qs3.Count()
	if err1 != nil {
		return nil, 0, 0, err1
	}

	qs1 = qs1.Limit(totalPerPage, fromRecord)
	if _, err := qs1.All(&list); err != nil {
		return nil, 0, 0, err
	}

	return list, totalRecord, totalAllRecord, nil
}

func (r *invoiceRepository) AllList(ormer *db.DB, paymentDate string, current int32, totalPerPage int32) ([] *modelShipper.Invoice, int64, error) {
	var (
		list []*modelShipper.Invoice
		qs   = ormer.QueryTable(new(modelShipper.Invoice))
		cond = orm.NewCondition()
		fromRecord int32
		totalRecord int64
	)
	fromRecord = current * totalPerPage - totalPerPage
	qs1 := qs.SetCond(cond).
		Filter("payment_date_time__istartswith", paymentDate).
		Filter("payment_status__in", 11, 12, 13, 14).
		OrderBy("-invoice_code")
	qs2 := qs.SetCond(cond).
		Filter("payment_date_time__istartswith", paymentDate).
		Filter("payment_status__in", 11, 12, 13, 14)

	totalRecord, err := qs2.Count()
	if err != nil {
		return nil, 0, err
	}

	qs1 = qs1.Limit(totalPerPage, fromRecord)
	if _, err := qs1.All(&list); err != nil {
		return nil, 0, err
	}

	return list, totalRecord, nil
}

func (r *invoiceRepository) GetInvoiceDetail(ormer *db.DB, invoiceCode string) (*modelShipper.Invoice, [] *dtos.InvoiceDetail, error) {
	cond := orm.NewCondition().
		And("invoice_code", invoiceCode)
	condGeneral := orm.NewCondition().And("name", "img_host")
	invoiceInfo := &modelShipper.Invoice{}
	general := &modelShipper.General{}

	qs := ormer.QueryTable(new(modelShipper.Invoice))
	qsGeneral := ormer.QueryTable(new(modelShipper.General))
	err := qs.SetCond(cond).One(invoiceInfo)
	if err != nil {
		return nil, nil, err
	}
	errGeneral := qsGeneral.SetCond(condGeneral).One(general)
	if errGeneral != nil {
		return nil, nil, errGeneral
	}

	invoiceId := invoiceInfo.ID
	invoiceIdCond := orm.NewCondition().And("invoice_id", invoiceId)
	var invoiceDetailList []*modelShipper.InvoiceDetail
	qs1 := ormer.QueryTable(new(modelShipper.InvoiceDetail))
	_, err = qs1.SetCond(invoiceIdCond).All(&invoiceDetailList)
	if err != nil {
		return nil, nil, err
	}
	qsItem := ormer.QueryTable(new(modelShipper.Item))

	listInvoiceDetailDto := make([]*dtos.InvoiceDetail, 0)
	for _, item := range invoiceDetailList  {
		var (
			iv dtos.InvoiceDetail
			_  = copier.Copy(&iv, &item)
		)
		itemInfo := &modelShipper.Item{}
		condItem := orm.NewCondition().And("item_id", item.ItemID)
		err = qsItem.SetCond(condItem).One(itemInfo)
		if err != nil {
			return nil, nil, err
		}
		iv.ImgPath = general.Value + itemInfo.ImgPath
		listInvoiceDetailDto = append(listInvoiceDetailDto, &iv)
	}

	return invoiceInfo, listInvoiceDetailDto, nil

}

func NewInvoiceRepository() InvoiceRepository {
	return &invoiceRepository{}
}


