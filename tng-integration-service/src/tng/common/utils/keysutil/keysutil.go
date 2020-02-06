package keysutil

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type (
	infos map[string]string
	keys  map[string]infos
)

var keysMap keys

func InitialKeysResource(resourcePath string) error {
	if keysMap != nil {
		keysMap = map[string]infos{}
	}
	buf, err := ioutil.ReadFile(resourcePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buf, &keysMap)
	if err != nil {
		return err
	}
	return nil
}

func GetHashKey(key string) string {
	vKeysMap, ok := keysMap[key]
	if !ok {
		return ""
	}
	if vKeysMap == nil {
		return ""
	}

	vKey, ok := vKeysMap["hash_key"]
	if !ok {
		return ""
	}
	if vKey == "" {
		return ""
	}
	return vKey
}

func GetZaLoAppID(key string) string  {
	vKeysMap, ok := keysMap[key]
	if !ok {
		return ""
	}
	if vKeysMap == nil {
		return ""
	}

	vKey, ok := vKeysMap["zalo_appid"]
	if !ok {
		return ""
	}
	if vKey == "" {
		return ""
	}
	return vKey
}

func GetZaLoAppSecret(key string) string  {
	vKeysMap, ok := keysMap[key]
	if !ok {
		return ""
	}
	if vKeysMap == nil {
		return ""
	}

	vKey, ok := vKeysMap["zalo_app_secret"]
	if !ok {
		return ""
	}
	if vKey == "" {
		return ""
	}
	return vKey
}

func GetZaLoRedirectURI(key string) string  {
	vKeysMap, ok := keysMap[key]
	if !ok {
		return ""
	}
	if vKeysMap == nil {
		return ""
	}

	vKey, ok := vKeysMap["zalo_redirect_uri"]
	if !ok {
		return ""
	}
	if vKey == "" {
		return ""
	}
	return vKey
}

