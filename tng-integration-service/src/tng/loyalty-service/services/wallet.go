
package services

import (
	"context"
	"github.com/jinzhu/copier"
	"tng/common/logger"
	"tng/common/models/loyalty"
	"tng/common/utils/db"
	"tng/common/utils/redisutil"
	"tng/loyalty-service/dtos"
	"tng/loyalty-service/repositories"
)

type WalletService interface {
	InsertOrUpdate(context.Context, *dtos.InsertOrUpdateWalletRequest) (*dtos.InsertOrUpdateWalletResponse, error)
	Delete(context.Context, *dtos.DeleteWalletRequest) (*dtos.DeleteWalletResponse, error)
	GetByID(context.Context, *dtos.GetWalletRequest) (*dtos.GetWalletResponse, error)
	List(context.Context, *dtos.ListWalletRequest) (*dtos.ListWalletResponse, error)
}

type walletService struct {
	BaseService
	redisCache        redisutil.Cache
	walletRepository repositories.WalletRepository
}

func NewWalletService(
	dbFactory db.Factory,
	redisCache redisutil.Cache,
	walletRepository repositories.WalletRepository,
) WalletService {
	return &walletService{
		BaseService: BaseService{
			dbFactory: dbFactory,
		},
		redisCache:        redisCache,
		walletRepository: walletRepository,
	}
}

func (s *walletService) InsertOrUpdate(ctx context.Context, request *dtos.InsertOrUpdateWalletRequest) (*dtos.InsertOrUpdateWalletResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	defer s.dbFactory.Rollback(tx)
	modelWallet := &loyalty.Wallet{

            
            Id: request.Id,

                        
            UserId: request.UserId,

            
            Balance: request.Balance,

            
            TotalIn: request.TotalIn,

            
            TotalOut: request.TotalOut,

            
            BalancePromo: request.BalancePromo,

            
            TotalInPromo: request.TotalInPromo,

            
            TotalOutPromo: request.TotalOutPromo,

            
            Created: request.Created,

            
            Modified: request.Modified,

            
            Checksum: request.Checksum,

            
            ClassId: request.ClassId,

            
            ClassDate: request.ClassDate,

            
            Status: request.Status,

            
            AccMoney: request.AccMoney,

            
            PointDate: request.PointDate,

            	}
	err := s.walletRepository.InsertOrUpdate(ctx, tx, modelWallet)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "InsertOrUpdate wallet error: %v", err)
		return nil, err
	}
	resp := &dtos.InsertOrUpdateWalletResponse{
		Meta: dtos.NewMetaOK(),
	}
	return resp, nil
}

func (s *walletService) Delete(ctx context.Context, request *dtos.DeleteWalletRequest) (*dtos.DeleteWalletResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	defer s.dbFactory.Rollback(tx)
	err := s.walletRepository.Delete(ctx, tx, request.Id)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "Delete wallet error: %v", err)
		return nil, err
	}
	resp := &dtos.DeleteWalletResponse{
		Meta: dtos.NewMetaOK(),
	}
	return resp, nil
}

func (s *walletService) GetByID(ctx context.Context, request *dtos.GetWalletRequest) (*dtos.GetWalletResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	item, err := s.walletRepository.GetByID(ctx, tx, request.Id)
	if err != nil {
		logger.Errorf(ctx, "GetByID wallet error: %v", err)
		return nil, err
	}
	data := &dtos.WalletInfo{}
	err = copier.Copy(&data, &item)
	if err != nil {
		logger.Errorf(ctx, "Copy wallet error: %v", err)
		return nil, err
	}
	resp := &dtos.GetWalletResponse{
		Meta: dtos.NewMetaOK(),
		Data: data,
	}
	return resp, nil
}

func (s *walletService) List(ctx context.Context, request *dtos.ListWalletRequest) (*dtos.ListWalletResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	list, err := s.walletRepository.List(ctx, tx, request.PageIndex, request.PageSize)
	if err != nil {
		logger.Errorf(ctx, "List wallet error: %v", err)
		return nil, err
	}
	data := make([]*dtos.WalletInfo, 0)
	for _, item := range list {
		var (
			iv dtos.WalletInfo
			_  = copier.Copy(&iv, &item)
		)
		data = append(data, &iv)
	}
	resp := &dtos.ListWalletResponse{
		Meta: dtos.NewMetaOK(),
		Data: data,
	}
	return resp, nil
}

