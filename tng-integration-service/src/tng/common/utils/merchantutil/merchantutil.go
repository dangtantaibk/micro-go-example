package merchantutil

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

const (
	MERCHANT_KEY1        = "key1"
	MERCHANT_KEY2        = "key2"
	MERCHANT_ZPS_URL     = "zps_url"
	MERCHANT_RSA_PRIVATE = "rsa_private"
	MERCHANT_RSA_        = "rsa_public"
	MERCHANT_APP_ID      = "app_id"
	MERCHANT_CODE        = "merchant_code"
	MERCHANT_STATUS      = "status"
)

type (
	merchantInfos     map[string]string
	keysMerchantInfos map[string]merchantInfos
)

var merchantInfosMap keysMerchantInfos

func InitialResource(resourcePath string) error {
	if merchantInfosMap != nil {
		merchantInfosMap = map[string]merchantInfos{}
	}
	buf, err := ioutil.ReadFile(resourcePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buf, &merchantInfosMap)
	if err != nil {
		return err
	}
	return nil
}

func GetValue(merchantCode, key string) string {
	vKeysMap, ok := merchantInfosMap[merchantCode]
	if !ok {
		return ""
	}
	if vKeysMap == nil {
		return ""
	}

	vKey, ok := vKeysMap[key]
	if !ok {
		return ""
	}
	if vKey == "" {
		return ""
	}
	return vKey
}

func GetKeys() []reflect.Value  {
	keys := reflect.ValueOf(merchantInfosMap).MapKeys()
	return keys
}
