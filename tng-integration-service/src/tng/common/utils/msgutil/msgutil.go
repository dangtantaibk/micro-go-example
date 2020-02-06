package msgutil

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"tng/common/httpcode"
)

// ModuleCommon is common module.
const ModuleCommon = "Common"

// GetMessage returns error message with code.
func GetMessage(module, lang string, code int) string {
	msg := httpcode.GetHTTPStatusText(code)
	if errorsMap == nil {
		return msg
	}
	err, ok := errorsMap[module]
	if err == nil || !ok {
		return tryGetCommonMessage(lang, code)
	}
	messages, ok := err[code]
	if messages == nil || !ok {
		return tryGetCommonMessage(lang, code)
	}
	msgLang, ok := messages[lang]
	if msgLang == "" || !ok {
		return msg
	}
	return msgLang
}

type (
	messages map[string]string
	appError map[int]messages
	errors   map[string]appError
)

var errorsMap errors

// InitialErrorMessageResource initials error message map from file.
func InitialErrorMessageResource(resourcePath string) error {
	if errorsMap != nil {
		errorsMap = map[string]appError{}
	}
	buf, err := ioutil.ReadFile(resourcePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buf, &errorsMap)
	if err != nil {
		return err
	}
	return nil
}

func tryGetCommonMessage(lang string, code int) string {
	msg := httpcode.GetHTTPStatusText(code)
	if errorsMap == nil {
		return msg
	}
	err, ok := errorsMap[ModuleCommon]
	if err == nil || !ok {
		return msg
	}
	messages, ok := err[code]
	if messages == nil || !ok {
		return msg
	}
	msgLang, ok := messages[lang]
	if msgLang == "" || !ok {
		return msg
	}
	return msgLang
}
