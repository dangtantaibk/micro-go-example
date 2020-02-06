package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/jinzhu/copier"
	"reflect"
	"tng/common/models/pos"
	"tng/common/utils/db"
	"tng/common/utils/redisutil"
	"tng/pos-service/dtos"
	"tng/pos-service/helper"
)

type ItemRepository interface {
	ListOrder(context.Context, []string, redisutil.Cache, *db.DB, interface{}) error
	MigrateItem(context.Context, string, redisutil.Cache, *db.DB) error
}

type itemRepository struct{}

func NewItemRepository() ItemRepository {
	return &itemRepository{}
}

func (r *itemRepository) ListOrder(ctx context.Context, keyItemIds []string, redis redisutil.Cache, ormer *db.DB, listItem interface{}) error {
	// read in redis
	result, err := redis.MGet(keyItemIds)
	if err != nil {
		return err
	}
	dataType := getDataType(listItem)
	values := reflect.MakeSlice(reflect.SliceOf(reflect.PtrTo(dataType)), len(keyItemIds), len(keyItemIds))
	for i, _ := range keyItemIds {
		values.Index(i).Set(reflect.New(dataType))

	}
	for i, _res := range result {
		value := values.Index(i).Interface()
		if str, ok := _res.(string); ok {
			if err := json.Unmarshal([]byte(str), &value); err != nil {
				continue
			}
		}
	}

	copier.Copy(`listItem`, values)
	return nil
}

func (r *itemRepository) MigrateItem(ctx context.Context, merchantCode string, redis redisutil.Cache, ormer *db.DB) error {
	var (
		list []*pos.Item
		qs   = ormer.QueryTable(new(pos.Item))
	)
	if _, err := qs.All(&list); err != nil {
		return err
	}
	for _, item := range list {
		dtosItem := &dtos.Item{
			ItemID:           item.ID,
			ItemName:         item.ItemName,
			ItemCode:         item.ItemCode,
			Price:            int64(item.Price),
			UnitName:         "",
			CategoryName:     "",
			CategoryID:       item.CategoryID,
			ImgPath:          item.ImgPath,
			ImgCrc:           item.ImgCrc,
			Description:      item.Description,
			Inventory:        item.Inventory,
			Status:           item.Status,
			CreateBy:         "",
			CreateDate:       "",
			BarCode:          item.BarCode,
			ModifiedBy:       item.ModifiedBy,
			ModifiedDateTime: item.ModifiedDateTime,
			Order:            item.Order,
			OriginalPrice:    int64(item.OriginalPrice),
			PromotionType:    0,
			PromotionID:      int64(item.PromotionID),
			CateMask:         int64(item.CateMask),
			PrinterMask:      item.PrinterMask,
			KitchenAreaID:    int64(item.KitchenAreaID),
			AreaID:           item.KitchenAreaID,
			DayOfSchedule:    0,
		}
		k := helper.KeyItem(merchantCode, int(item.ID))
		_ = redis.Set(k, dtosItem, 0)
	}
	return nil
}

func protoEncode(msg interface{}) ([]byte, error) {
	if pMsg, ok := msg.(protobuf.Message); ok {
		return protobuf.Marshal(pMsg)
	}
	return nil, fmt.Errorf("%v not proto message", reflect.TypeOf(msg))
}
func protoDecode(data string, msg interface{}) error {
	if pMsg, ok := msg.(protobuf.Message); ok {
		return protobuf.Unmarshal([]byte(data), pMsg)
	}
	return fmt.Errorf("%v not proto message", reflect.TypeOf(msg))
}

func getDataType(data interface{}) reflect.Type {
	value := reflect.Indirect(reflect.ValueOf(data))

	if value.Kind() == reflect.Slice {
		value := value.Interface()
		newSlice := reflect.New(reflect.TypeOf(value))
		if err := json.Unmarshal([]byte("[{}]"), newSlice.Interface()); err != nil {
			return nil
		}
		valueSlice := reflect.Indirect(newSlice)
		protoItem := valueSlice.Index(0)
		item := reflect.Indirect(protoItem).Interface()
		return reflect.TypeOf(item)

	}
	return nil
}
