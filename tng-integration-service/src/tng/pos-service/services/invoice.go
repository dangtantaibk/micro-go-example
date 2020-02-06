package services

import (
	"context"
	"fmt"
	"strconv"
	"time"
	_zlpAdapter "tng/common/adapters/zalopay"
	"tng/common/models"
	"tng/common/models/pos"
	"tng/common/utils/cfgutil"
	"tng/common/utils/db"
	"tng/common/utils/hashutil"
	"tng/common/utils/redisutil"
	"tng/pos-service/dtos"
	"tng/pos-service/helper"
	"tng/pos-service/repositories"
)

type InvoiceService interface {
	Create(ctx context.Context, request *dtos.CreateInvoiceRequest) (*dtos.CreateInvoiceResponse, error)
	Cancel(ctx context.Context, request *dtos.CancelInvoiceRequest) (*dtos.CancelInvoiceResponse, error)
	Refund(ctx context.Context, request *dtos.RefundInvoiceRequest) (*dtos.RefundInvoiceResponse, error)
}

type invoiceService struct {
	BaseService
	redisCache        redisutil.Cache
	zlpAdapter        _zlpAdapter.Adapter
	invoiceRepository repositories.InvoiceRepository
	itemRepository    repositories.ItemRepository
}

func NewInvoiceService(
	dbFactory db.Factory,
	redisCache redisutil.Cache,
	invoiceRepository repositories.InvoiceRepository,
	itemRepository repositories.ItemRepository) InvoiceService {
	return &invoiceService{
		BaseService: BaseService{
			dbFactory: dbFactory,
		},
		redisCache:        redisCache,
		invoiceRepository: invoiceRepository,
		itemRepository:    itemRepository,
	}
}

func (s *invoiceService) Create(ctx context.Context, request *dtos.CreateInvoiceRequest) (*dtos.CreateInvoiceResponse, error) {
	var (
		tx = s.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
	)
	ivNo, ivCode := s.genInvoiceCode(request.MerchantCode)
	timeStamp := time.Now().Unix()
	rawCK := fmt.Sprintf("%s||%d|%s", request.MerchantCode, timeStamp, cfgutil.Load(dtos.CfgMd5KeyInvoice))
	vposToken := hashutil.GetMD5Hash(rawCK)
	totalAmount := s.CalculateAmount(ctx, request.MerchantCode, request.Items)
	fmt.Println("totalAmount: ", totalAmount)
	invoice := &pos.Invoice{
		InvoiceCode:     ivCode,
		InvoiceNo:       int(ivNo),
		Amount:          10000,
		CreatedDateTime: time.Now().Local().Format(models.FormatYYYMMDDHHMMSS),
		PaymentMethod:   request.PaymentMethod,
		PaymentDateTime: models.FormatDefaultDateTime,
		AuditDateTime:   models.FormatDefaultDateTime,
		MachineName:     request.MachineName,
		Note:            request.Description,
		VposToken:       vposToken,
	}
	defer s.dbFactory.Rollback(tx)
	err := s.invoiceRepository.Create(tx, invoice)
	s.dbFactory.Commit(tx)
	if err != nil {
		return nil, err
	}
	return &dtos.CreateInvoiceResponse{
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
	}, nil
}

func (s *invoiceService) Cancel(ctx context.Context, request *dtos.CancelInvoiceRequest) (*dtos.CancelInvoiceResponse, error) {
	panic("implement me")
}

func (s *invoiceService) Refund(ctx context.Context, request *dtos.RefundInvoiceRequest) (*dtos.RefundInvoiceResponse, error) {
	panic("implement me")
}

func (s *invoiceService) genInvoiceCode(merchantCode string) (int64, string) {
	nowDate, key := helper.KeyGenInvoiceCode(merchantCode)
	v, e := s.redisCache.Incr(key)
	if e != nil {
		return 0, ""
	}
	if v <= 0 {
		return 0, ""
	}
	return v, fmt.Sprintf("%s%04d", nowDate, v)
}

func (s *invoiceService) CalculateAmount(ctx context.Context, merchantCode string, items []*dtos.ItemInvoice) int64 {
	var (
		tx          = s.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", merchantCode), true)
		totalAmount int64
	)
	keyItemIds := []string{}
	mapItems := map[int]*dtos.ItemInvoice{}
	for _, item := range items {
		if item == nil {
			continue
		}
		k := helper.KeyItem(merchantCode, item.ItemID)
		keyItemIds = append(keyItemIds, k)
		mapItems[item.ItemID] = item
	}
	listItem := []pos.Item{}
	err := s.itemRepository.ListOrder(ctx, keyItemIds, s.redisCache, tx, listItem)
	fmt.Println("listItem: ", listItem)
	if err != nil {
		return 0
	}
	for _, item := range listItem {
		//if item == nil {
		//	continue
		//}
		if result, ok := mapItems[int(item.ID)]; ok {
			s := fmt.Sprintf("%.0f", item.Price)
			p, _ := strconv.Atoi(s)
			totalAmount += int64(result.Quantity * p)
		}
	}

	return totalAmount
}
