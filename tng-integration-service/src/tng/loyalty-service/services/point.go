package services

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"math"
	"strconv"
	"tng/common/location"
	"tng/common/logger"
	"tng/common/models"
	"tng/common/models/loyalty"
	"tng/common/utils/db"
	"tng/common/utils/redisutil"
	"tng/loyalty-service/dtos"
	"tng/loyalty-service/repositories"
)

type PointService interface {
	InsertOrUpdate(context.Context, *dtos.InsertOrUpdatePointRequest) (*dtos.InsertOrUpdatePointResponse, error)
	Delete(context.Context, *dtos.DeletePointRequest) (*dtos.DeletePointResponse, error)
	GetByID(context.Context, *dtos.GetPointRequest) (*dtos.GetPointResponse, error)
	List(context.Context, *dtos.ListPointRequest) (*dtos.ListPointResponse, error)
	AddPoint(context.Context, *dtos.AddPointRequest) (*dtos.AddPointResponse, error)
	Search(ctx context.Context, request *dtos.SearchPointRequest) (*dtos.SearchPointResponse, error)
	PointHistory(context.Context, *dtos.PointHistoryRequest) (*dtos.PointHistoryResponse, error)
	CheckPoint(context.Context, *dtos.CheckPointRequest) (*dtos.CheckPointResponse, error)
	CheckOldPoint(context.Context, *dtos.CheckOldPointRequest) (*dtos.CheckOldPointResponse, error)
	getMaxPointToGrow(tx *db.DB, classId int64) (int64, float64, error)
	getFinalClass(tx *db.DB, point int64) (int64, error)
}

type pointService struct {
	BaseService
	redisCache           redisutil.Cache
	pointRepository      repositories.PointRepository
	walletRepository     repositories.WalletRepository
	pointClassRepository repositories.PointClassRepository
}

func NewPointService(
	dbFactory db.Factory,
	redisCache redisutil.Cache,
	pointRepository repositories.PointRepository,
	walletRepository repositories.WalletRepository,
	pointClassRepository repositories.PointClassRepository,
) PointService {
	return &pointService{
		BaseService: BaseService{
			dbFactory: dbFactory,
		},
		redisCache:           redisCache,
		pointRepository:      pointRepository,
		walletRepository:     walletRepository,
		pointClassRepository: pointClassRepository,
	}
}

func (s *pointService) Search(ctx context.Context, request *dtos.SearchPointRequest) (*dtos.SearchPointResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	list, totalRecord, err := s.pointRepository.Search(ctx, tx, request.UserId, request.PointType, request.Source, request.ForTransactionId,
		request.Notes, request.AppId, request.Point, request.TransactionAmount, request.CreatedFrom,
		request.CreatedTo, request.Status, request.PromotionPercent, request.Rate, request.SortColumn, request.SortDirection,
		request.PageIndex, request.PageSize)
	if err != nil {
		logger.Errorf(ctx, "List point error: %v", err)
		return nil, err
	}
	data := make([]*dtos.PointInfo, 0)
	for _, item := range list {
		var (
			iv dtos.PointInfo
			_  = copier.Copy(&iv, &item)
		)
		data = append(data, &iv)
	}
	resp := &dtos.SearchPointResponse{
		Meta:        dtos.NewMetaOK(),
		Data:        data,
		TotalRecord: totalRecord,
	}
	return resp, nil
}
func (s *pointService) AddPoint(ctx context.Context, request *dtos.AddPointRequest) (*dtos.AddPointResponse, error) {
	var (
		tx               = s.dbFactory.GetDB(true)
		t                = location.GetVNCurrentTime()
		formatted        = t.Format(models.FormatYYYYMMDD1)
		now              = t.Format(models.FormatYYYMMDDHHMMSS)
		maxClassId int64 = 1
	)
	defer s.dbFactory.Rollback(tx)

	//get user from DB
	wallet, err := s.walletRepository.GetByUserID(ctx, tx, request.UserId)

	if err != nil {
		logger.Errorf(ctx, "Get wallet point error: %v", err)
		return nil, err
	}
	//calculate point
	if wallet.Id == 0 {
		wallet.ClassId = 1
	}

	maxPoint, rate, err := s.getMaxPointToGrow(tx, wallet.ClassId)

	point := math.Round((float64(request.TransactionAmount) * rate) / 100000)

	modelPoint := &loyalty.Point{
		Id:               0,
		UserId:           request.UserId,
		PointType:        "in",
		Point:            int64(math.Round(point)),
		Source:           request.Source,
		ForTransactionId: request.ForTransactionId,
		TransactionAmount: request.TransactionAmount,
		Notes: request.Notes,
		Created: now,
		CreatedYmd: formatted,
		Status: 0,
		AppId: request.AppId,
		PromotionPercent: request.PromotionPercent,
		CampaignCode: request.CampaignCode,
		Channel: request.Channel,
		Rate: rate,
		JsonDetail: request.JsonDetail,
	}
	err = s.pointRepository.InsertOrUpdate(ctx, tx, modelPoint)

	if err != nil {
		logger.Errorf(ctx, "Insert point item error: %v", err)
		return nil, err
	}

	wallet.UserId = request.UserId
	wallet.AccMoney += request.TransactionAmount
	wallet.Balance += point
	wallet.TotalIn += point
	if wallet.Id == 0 {
		wallet.Created = now
		wallet.ClassDate = now
		wallet.PointDate = now
	}
	wallet.Modified = now

	//reclass
	if wallet.Balance >= float64(maxPoint) {
		/*if wallet.ClassId < 6 {
			wallet.ClassId = wallet.ClassId + 1
			wallet.ClassDate = now
		}*/
		maxClassId, err = s.getFinalClass(tx, int64(wallet.Balance))
		if err != nil {
			logger.Errorf(ctx, "Error get class id to rerank: %v", err)
			return nil, err
		}
		wallet.ClassId = maxClassId
		wallet.ClassDate = now
	}
	rck := wallet.UserId + fmt.Sprintf("%d%d%d%d%d%d%d", wallet.Balance, wallet.TotalIn, wallet.TotalInPromo, wallet.TotalOut, wallet.TotalOutPromo, wallet.Status, wallet.ClassId)
	sha512Bytes := sha512.Sum512([]byte(rck))
	wallet.Checksum = hex.EncodeToString(sha512Bytes[:])

	err = s.walletRepository.InsertOrUpdate(ctx, tx, wallet)
	if err != nil {
		logger.Errorf(ctx, "InsertOrUpdate wallet error: %v", err)
		return nil, err
	}

	s.dbFactory.Commit(tx)



	data := &dtos.WalletInfo{}
	err = copier.Copy(&data, &wallet)

	key := "zl-" + wallet.UserId
	s.redisCache.Set(key, data, 0)

	if err != nil {
		logger.Errorf(ctx, "Commit point error: %v", err)
	}
	resp := &dtos.AddPointResponse{
		Meta: dtos.NewMetaOK(),
		Data:data,
	}
	return resp, nil
}
func (s *pointService) InsertOrUpdate(ctx context.Context, request *dtos.InsertOrUpdatePointRequest) (*dtos.InsertOrUpdatePointResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	defer s.dbFactory.Rollback(tx)
	modelPoint := &loyalty.Point{
		Id: request.Id,
		UserId: request.UserId,
		PointType: request.PointType,
		Point: request.Point,
		Source: request.Source,
		ForTransactionId: request.ForTransactionId,
		TransactionAmount: request.TransactionAmount,
		Notes: request.Notes,
		Created: request.Created,
		CreatedYmd: request.CreatedYmd,
		Status: request.Status,
		AppId: request.AppId,
		PromotionPercent: request.PromotionPercent,
		CampaignCode: request.CampaignCode,
		Channel: request.Channel,
		Rate: request.Rate,
		JsonDetail: request.JsonDetail,
	}
	err := s.pointRepository.InsertOrUpdate(ctx, tx, modelPoint)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "InsertOrUpdate point error: %v", err)
	}
	resp := &dtos.InsertOrUpdatePointResponse{
		Meta: dtos.NewMetaOK(),
	}
	return resp, nil
}

func (s *pointService) Delete(ctx context.Context, request *dtos.DeletePointRequest) (*dtos.DeletePointResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	defer s.dbFactory.Rollback(tx)
	err := s.pointRepository.Delete(ctx, tx, request.Id)
	s.dbFactory.Commit(tx)
	if err != nil {
		logger.Errorf(ctx, "Delete point error: %v", err)
		return nil, err
	}
	resp := &dtos.DeletePointResponse{
		Meta: dtos.NewMetaOK(),
	}
	return resp, nil
}

func (s *pointService) GetByID(ctx context.Context, request *dtos.GetPointRequest) (*dtos.GetPointResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	item, err := s.pointRepository.GetByID(ctx, tx, request.Id)
	if err != nil {
		logger.Errorf(ctx, "GetByID point error: %v", err)
		return nil, err
	}
	data := &dtos.PointInfo{}
	err = copier.Copy(&data, &item)
	if err != nil {
		logger.Errorf(ctx, "Copy point error: %v", err)
		return nil, err
	}
	resp := &dtos.GetPointResponse{
		Meta: dtos.NewMetaOK(),
		Data: data,
	}
	return resp, nil
}

func (s *pointService) List(ctx context.Context, request *dtos.ListPointRequest) (*dtos.ListPointResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	list, err := s.pointRepository.List(ctx, tx, request.PageIndex, request.PageSize)
	if err != nil {
		logger.Errorf(ctx, "List point error: %v", err)
		return nil, err
	}
	data := make([]*dtos.PointInfo, 0)
	for _, item := range list {
		var (
			iv dtos.PointInfo
			_  = copier.Copy(&iv, &item)
		)
		data = append(data, &iv)
	}
	resp := &dtos.ListPointResponse{
		Meta: dtos.NewMetaOK(),
		Data: data,
	}
	return resp, nil
}

func (s *pointService) CheckPoint(ctx context.Context, request *dtos.CheckPointRequest) (*dtos.CheckPointResponse, error) {
	var (
		tx         = s.dbFactory.GetDB(true)
		walletInfo = &dtos.WalletInfo{}
	)
	defer s.dbFactory.Rollback(tx)

	key := "zl-" + request.UserId
	rs, err := s.redisCache.Get(key)

	//jsonInput, err := strconv.Unquote(string(rs))
	err = json.Unmarshal([]byte(rs), walletInfo)

	if err == nil {
		resp := &dtos.CheckPointResponse{
			Meta: dtos.NewMetaOK(),
			Data: walletInfo,
		}
		return resp, nil
	}
	wallet, err := s.walletRepository.GetByUserID(ctx, tx, request.UserId)

	if err != nil {
		logger.Errorf(ctx, "GetByUserID point error: %v", err)
		return nil, err
	}

	data := &dtos.WalletInfo{}
	err = copier.Copy(&data, &wallet)
	s.redisCache.Set(key, data, 0)
	resp := &dtos.CheckPointResponse{
		Meta: dtos.NewMetaOK(),
		Data: walletInfo,
	}
	return resp, nil
}

func (s *pointService) PointHistory(ctx context.Context, request *dtos.PointHistoryRequest) (*dtos.PointHistoryResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	list, err := s.pointRepository.PointHistory(ctx, tx, request.PageIndex, request.PageSize, request.UserId)
	if err != nil {
		logger.Errorf(ctx, "PointHistory point error: %v", err)
		return nil, err
	}
	data := make([]*dtos.PointInfo, 0)
	for _, item := range list {
		var (
			iv dtos.PointInfo
			_  = copier.Copy(&iv, &item)
		)
		data = append(data, &iv)
	}
	resp := &dtos.PointHistoryResponse{
		Meta: dtos.NewMetaOK(),
		Data: data,
	}
	return resp, nil
}
func (s *pointService) CheckOldPoint(ctx context.Context, request *dtos.CheckOldPointRequest) (*dtos.CheckOldPointResponse, error) {
	var (
		tx = s.dbFactory.GetDB(true)
	)
	item, err := s.pointRepository.GetByTransactionID(ctx, tx, request.ForTransactionId, request.AppId, request.UserId)
	if err != nil {
		logger.Errorf(ctx, "GetByID point error: %v", err)
		return nil, err
	}
	data := &dtos.PointInfo{}
	err = copier.Copy(&data, &item)
	if err != nil {
		logger.Errorf(ctx, "Copy point error: %v", err)
		return nil, err
	}
	resp := &dtos.CheckOldPointResponse{
		Meta: dtos.NewMetaOK(),
		Data: item,
	}
	return resp, nil
}
func (s *pointService) getMaxPointToGrow(tx *db.DB, classId int64) (int64, float64, error) {
	//get class
	keyClass := "point_class"
	rsClass, err := s.redisCache.Get(keyClass)

	var listClass []*loyalty.PointClass

	jsonInput, err := strconv.Unquote(string(rsClass))
	err = json.Unmarshal([]byte(jsonInput), listClass)

	if err != nil {
		listClass, err = s.pointClassRepository.GetAll(tx)
		s.redisCache.Set(keyClass, listClass, 0)
		if err != nil {
			logger.Errorf(nil, "Error get class from db point error: %v", err)
			return 0, 0, err
		}
	}
	for _, item := range listClass {
		if item.Id == classId {
			return item.RequireNumOfPoint, item.DiscountPercent, nil
		}
	}
	return 0, 0, err
}
func (s *pointService) getFinalClass(tx *db.DB, point int64) (int64, error) {
	var (
		maxClassId int64 = 1
	)
	//get class
	keyClass := "point_class"
	rsClass, err := s.redisCache.Get(keyClass)

	var listClass []*loyalty.PointClass

	jsonInput, err := strconv.Unquote(string(rsClass))
	err = json.Unmarshal([]byte(jsonInput), listClass)

	if err != nil {
		listClass, err = s.pointClassRepository.GetAll(tx)
		s.redisCache.Set(keyClass, listClass, 0)
		if err != nil {
			logger.Errorf(nil, "Error get class from db point error: %v", err)
			return 0, err
		}
	}
	for _, item := range listClass {
		if item.RequireNumOfPoint <= point {
			maxClassId = item.Id
		}
	}
	return maxClassId, err
}
