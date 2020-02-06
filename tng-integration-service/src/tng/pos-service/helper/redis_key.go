package helper

import (
	"fmt"
	"time"
	"tng/common/models"
	"tng/common/utils/cfgutil"
	"tng/pos-service/dtos"
)

var (
	AppName = cfgutil.Load(dtos.CfgAppName)
)

func KeyGenInvoiceCode(merchantCode string) (string, string) {
	nowDate := time.Now().Local().Format(models.FormatYYMMDD)
	return nowDate, fmt.Sprintf("%s_%s_%s", AppName, merchantCode, nowDate)
}

func KeyItem(merchantCode string, itemID int) string {
	return fmt.Sprintf("%s_%s_%d", AppName, merchantCode, itemID)
}
