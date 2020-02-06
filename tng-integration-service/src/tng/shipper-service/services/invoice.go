package services

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"tng/common/logger"
	"tng/common/models"
	"tng/common/utils/db"
	"tng/common/utils/redisutil"
	"tng/shipper-service/dtos"
	"tng/shipper-service/repositories"
)

type invoiceService struct {
	BaseService
	redisCache        redisutil.Cache
	invoiceRepository repositories.InvoiceRepository
}

type InvoiceService interface {
	List(ctx context.Context, request *dtos.ListInvoiceRequest, phoneNumber string) (*dtos.ListInvoiceResponse, error)
	AllList(ctx context.Context, request *dtos.ListInvoiceRequest) (*dtos.ListInvoiceResponse, error)
	UpdateStatus(ctx context.Context, request *dtos.UpdateInvoiceStatusRequest) (*dtos.UpdateInvoiceStatusResponse, error)
	GetByVposToken(ctx context.Context, request *dtos.ScanQRCodeRequest) (*dtos.ScanQRCodeResponse, error)
	GetInvoiceDetail(ctx context.Context, request *dtos.GetInvoiceDetailRequest) (*dtos.GetInvoiceDetailResponse, error)
}

func (s *invoiceService) List(ctx context.Context, request *dtos.ListInvoiceRequest, phoneNumber string) (*dtos.ListInvoiceResponse, error) {
	var (
		tx = s.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
		currentPage int32 = 1
		totalTranPerPage int32 = 50
	)
	if request.CurrentPage != 0 {
		currentPage = request.CurrentPage
	}
	if request.TotalTransPerPage != 0 {
		totalTranPerPage = request.TotalTransPerPage
	}

	list, totalRecord, totalAllRecord, err := s.invoiceRepository.List(tx, request.Date, phoneNumber, currentPage, totalTranPerPage)

	if err != nil {
		logger.Errorf(ctx, "Get list invoice error: %v", err)
		return nil, err
	}

	data := make([]dtos.Invoice, 0)
	for _, item := range list {
		var (
			iv dtos.Invoice
			_  = copier.Copy(&iv, &item)
		)
		data = append(data, iv)
	}
	response := &dtos.ListInvoiceResponse{
		Meta: dtos.Meta{
			Code:    1,
			Message: "ok",
		},
		Data: data,
		TotalRecord: totalRecord,
		TotalAllRecord: totalAllRecord,
	}
	return response, nil
}

func (s *invoiceService) AllList(ctx context.Context, request *dtos.ListInvoiceRequest) (*dtos.ListInvoiceResponse, error) {
	var (
		tx = s.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
		currentPage int32 = 1
		totalTranPerPage int32 = 50
	)
	if request.CurrentPage != 0 {
		currentPage = request.CurrentPage
	}
	if request.TotalTransPerPage != 0 {
		totalTranPerPage = request.TotalTransPerPage
	}

	list, totalRecord, err := s.invoiceRepository.AllList(tx, request.Date, currentPage, totalTranPerPage)

	if err != nil {
		logger.Errorf(ctx, "Get list invoice error: %v", err)
		return nil, err
	}

	data := make([]dtos.Invoice, 0)
	for _, item := range list {
		var (
			iv dtos.Invoice
			_  = copier.Copy(&iv, &item)
		)
		data = append(data, iv)
	}
	response := &dtos.ListInvoiceResponse{
		Meta: dtos.Meta{
			Code:    1,
			Message: "ok",
		},
		Data: data,
		TotalAllRecord: totalRecord,
	}
	return response, nil
}

func (s *invoiceService) UpdateStatus(ctx context.Context, request *dtos.UpdateInvoiceStatusRequest) (*dtos.UpdateInvoiceStatusResponse, error) {
	var (
		tx = s.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
	)
	defer s.dbFactory.Rollback(tx)
	err := s.invoiceRepository.Update(tx, *request)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "Update payment status error: %v", err)
		return nil, err
	}
	response := &dtos.UpdateInvoiceStatusResponse{
			Meta:dtos.Meta{
			Code:    1,
			Message: "ok",
		},
	}
	return response, nil
}

func (s *invoiceService) GetInvoiceDetail(ctx context.Context, request *dtos.GetInvoiceDetailRequest) (*dtos.GetInvoiceDetailResponse, error) {
	var (
		tx = s.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
	)
	invoice, listInvoice, err := s.invoiceRepository.GetInvoiceDetail(tx, request.InvoiceCode)
	if err != nil {
		logger.Errorf(ctx, "Get invoice detail error: %v", err)
		return nil, dtos.NewAppError(dtos.ErrorNotFoundData)
	}

	var ivDto dtos.Invoice
	err = copier.Copy(&ivDto, &invoice)
	if err != nil {
		logger.Errorf(ctx, "parse object error: %v", err)
		return nil, dtos.NewAppError(dtos.ErrorUnknown)
	}

	//listInvoiceDetailDto := make([]dtos.InvoiceDetail, 0)
	//for _, item := range listInvoice {
	//	var (
	//		iv dtos.InvoiceDetail
	//		_  = copier.Copy(&iv, &item)
	//	)
	//	listInvoiceDetailDto = append(listInvoiceDetailDto, iv)
	//}

	response := &dtos.GetInvoiceDetailResponse{
		Meta:          dtos.Meta{
			Code:    1,
			Message: "OK",
		},
		Invoice:       ivDto,
		InvoiceDetail: listInvoice,
	}

	return response, nil
}

func (s *invoiceService) GetByVposToken(ctx context.Context, request *dtos.ScanQRCodeRequest) (*dtos.ScanQRCodeResponse, error) {
	var (
		tx = s.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", request.MerchantCode), true)
	)
	invoice, listInvoice, err := s.invoiceRepository.GetByVposToken(tx, request.VposToken)
	if err != nil {
		logger.Errorf(ctx, "Get list invoice error: %v", err)
		return nil, dtos.NewAppError(dtos.ErrorNotFoundData)
	}

	var ivDto dtos.Invoice
	err = copier.Copy(&ivDto, &invoice)
	if err != nil {
		logger.Errorf(ctx, "parse object error: %v", err)
		return nil, dtos.NewAppError(dtos.ErrorUnknown)
	}

	response := &dtos.ScanQRCodeResponse{
		Meta:          dtos.Meta{
			Code:    1,
			Message: "OK",
		},
		Invoice:       ivDto,
		InvoiceDetail: listInvoice,
	}

	return response, nil
}

func NewInvoiceService(
	dbFactory db.Factory,
	redisCache redisutil.Cache,
	invoiceRepository repositories.InvoiceRepository) InvoiceService {
	return &invoiceService{
		BaseService: BaseService{
			dbFactory: dbFactory,
		},
		redisCache:        redisCache,
		invoiceRepository: invoiceRepository,
	}
}
