package services

import (
	"context"
	"fmt"
	"strings"
	"tng/common/location"
	"tng/common/logger"
	"tng/common/models"
	"tng/common/utils/cfgutil"
	"tng/common/utils/db"
	"tng/common/utils/mqttcli"
	"tng/cron-job-service/dtos"
	"tng/cron-job-service/repositories"
)

type ScheduleService interface {
	Sync(ctx context.Context) error
	NotifyUpdateMenu(ctx context.Context) error
	WarmUp(ctx context.Context, request *dtos.WarmUpScheduleRequest) (*dtos.WarmUpScheduleResponse, error)
}

type scheduleService struct {
	BaseService
	scheduleRepository repositories.ScheduleRepository
	itemRepository     repositories.ItemRepository
	mqttCli            mqttcli.MqttCli
}

func NewMenuScheduleService(
	dbFactory db.Factory,
	mqttCli mqttcli.MqttCli,
	scheduleRepository repositories.ScheduleRepository,
	itemRepository repositories.ItemRepository) ScheduleService {
	return &scheduleService{
		BaseService: BaseService{
			dbFactory: dbFactory,
		},
		scheduleRepository: scheduleRepository,
		itemRepository:     itemRepository,
		mqttCli:            mqttCli,
	}
}

func (s *scheduleService) NotifyUpdateMenu(ctx context.Context) error {
	fmt.Println("NotifyUpdateMenu")
	listMerchantCode := cfgutil.Load("MERCHANT_CODES_UPDATE_MENU")
	arrMerchantCode := strings.Split(listMerchantCode, ",")
	for _, it := range arrMerchantCode {
		fmt.Println("merchant_code: ", it)
		preTopic := cfgutil.Load("MQTT_PRE_TOPIC")
		topic := fmt.Sprintf(preTopic + it)
		msg := "{\"type\":\"quantity_total\",\"data\":\"\",\"device_id\":\"\"}"
		s.mqttCli.Pub(ctx, topic, msg)
	}
	return nil
}

func (s *scheduleService) Sync(ctx context.Context) error {
	fmt.Println("Sync")
	listMerchantCode := cfgutil.Load("MERCHANT_CODES")
	arrMerchantCode := strings.Split(listMerchantCode, ",")
	for _, it := range arrMerchantCode {
		err := s.GetAndInsertListSchedule(ctx, it)
		if err != nil {
			logger.Errorf(ctx, "WarmUp Schedule error: %v", err)
		}
	}
	return nil
}

func (s *scheduleService) WarmUp(ctx context.Context, request *dtos.WarmUpScheduleRequest) (*dtos.WarmUpScheduleResponse, error) {
	fmt.Println("WarmUp")
	if request.MerchantCode == "" {
		listMerchantCode := cfgutil.Load("MERCHANT_CODES")
		arrMerchantCode := strings.Split(listMerchantCode, ",")
		for _, it := range arrMerchantCode {
			err := s.GetAndInsertListSchedule(ctx, it)
			if err != nil {
				logger.Errorf(ctx, "WarmUp Schedule error: %v", err)
			}
		}
	} else {
		err := s.GetAndInsertListSchedule(ctx, request.MerchantCode)
		if err != nil {
			logger.Errorf(ctx, "WarmUp Schedule error: %v", err)
		}
	}

	return &dtos.WarmUpScheduleResponse{
		Meta: dtos.Meta{
			Code:    1,
			Message: "OK",
		},
	}, nil
}

func (s *scheduleService) GetAndInsertListSchedule(ctx context.Context, merchantCode string) error {
	var (
		tx = s.dbFactory.GetDBByAlias(fmt.Sprintf(models.PreDatabase+"%s", merchantCode), true)
	)

	err := s.scheduleRepository.Delete(ctx, tx)
	if err != nil {
		logger.Errorf(ctx, "Delete Schedule error: %v", err)
	}

	listScheduleYesterday, err := s.scheduleRepository.GetList(ctx, tx)
	if err != nil {
		logger.Errorf(ctx, "GetList Schedule error: %v", err)
		return err
	}
	// insert
	t := location.GetVNCurrentTime()
	defer s.dbFactory.Rollback(tx)
	for _, it := range listScheduleYesterday {
		if it == nil {
			continue
		}
		itemId := it.ItemID
		item, err := s.itemRepository.GetDetails(ctx, tx, itemId)
		if err != nil {
			logger.Errorf(ctx, "GetList error: %v", err)
			continue
		}
		if item == nil {
			logger.Errorf(ctx, "No have item")
			continue
		}
		if itemId == int(item.ID) {
			it.DayOfSchedule = t.Day()
			it.Modified = t.Format(models.FormatYYYMMDDHHMMSS)
			err = s.scheduleRepository.Insert(ctx, tx, it)
			s.dbFactory.Commit(tx)
			if err != nil {
				logger.Errorf(ctx, "Insert Schedule error: %v", err)
			}
		}
	}

	return nil
}
