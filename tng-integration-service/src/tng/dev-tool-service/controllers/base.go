package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"tng/common/utils/msgutil"
	dtos2 "tng/dev-tool-service/dtos"

	"github.com/astaxie/beego"
	"gopkg.in/go-playground/validator.v9"
	"tng/common/httpcode"
	"tng/common/logger"
	"tng/common/utils/cfgutil"
)

const controllerPrefix = "Controller"

// These global variables for validator struct or get message translated in JSON file.
var globalValidator *validator.Validate

// BaseController represents basic controller of all controllers.
type BaseController struct {
	beego.Controller
	validator *validator.Validate
}

// Prepare is prepare function of BaseController.
func (c *BaseController) Prepare() {
	c.init()
}

func (c *BaseController) init() {
	c.validator = globalValidator
	c.Ctx.Output.ContentType("application/json; charset=utf-8")
}

// GetInts returns the input string slice by key int or the default value while it's present and input is blank
// it's designed for multi-value input field such as checkbox(input[type=checkbox]), multi-selection.
func (c *BaseController) GetInts(key string) ([]int, error) {
	var defv []int
	if f := c.Input(); f == nil {
		return defv, nil
	} else if vs := f[key]; len(vs) > 0 {
		for _, item := range vs {
			if item == "" {
				return defv, nil
			}
			params := strings.Split(item, ",")
			for _, value := range params {
				iValue, err := strconv.Atoi(value)
				if err != nil {
					return nil, err
				}
				defv = append(defv, iValue)
			}

		}
	}

	return defv, nil
}

// GetInt64s returns the input string slice by key int or the default value while it's present and input is blank
// it's designed for multi-value input field such as checkbox(input[type=checkbox]), multi-selection.
func (c *BaseController) GetInt64s(key string) ([]int64, error) {
	var defv []int64

	if f := c.Input(); f == nil {
		return defv, nil
	} else if vs := f[key]; len(vs) > 0 {
		for _, item := range vs {
			if item == "" {
				return defv, nil
			}
			params := strings.Split(item, ",")
			for _, value := range params {
				iValue, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, err
				}
				defv = append(defv, iValue)
			}

		}
	}
	return defv, nil
}

// GetStrings returns the input string slice by key int or the default value while it's present and input is blank
// it's designed for multi-value input field such as checkbox(input[type=checkbox]), multi-selection.
func (c *BaseController) GetStrings(key string) ([]string, error) {
	var defv []string

	if f := c.Input(); f == nil {
		return defv, nil
	} else if vs := f[key]; len(vs) > 0 {
		for _, item := range vs {
			if item == "" {
				return defv, nil
			}
			params := strings.Split(item, ",")
			defv = append(defv, params...)
		}
	}
	return defv, nil
}

// Validate valids struct and return error.
func (c *BaseController) Validate(s interface{}) error {
	return c.validator.Struct(s)
}

// Validate if string is not contain special character.
func customNotAllowSpecialCharacter(fl validator.FieldLevel) bool {
	var regex = regexp.MustCompile(`^[_A-z0-9]*((-|\s)*[_A-z0-9])*$`)
	return regex.MatchString(fl.Field().String())
}

// For initial translator and validator.
func init() {
	globalValidator = validator.New()
	_ = globalValidator.RegisterValidation("not_allow_special", customNotAllowSpecialCharacter)
}

// Respond responds to HTTP request.
// Note: v is a pointer.
func (c *BaseController) Respond(v interface{}, err error) {
	var (
		moduleName, _ = c.GetModuleAndMethodName()
		lng           = c.GetLang()
		respData      interface{}
	)
	if err != nil {
		appErr, ok := err.(dtos2.AppError)
		if ok {
			respData = appErr
		} else {
			respData = dtos2.NewAppError(dtos2.InternalServerError)
		}
	} else {
		respData = v
	}
	var (
		resp    = reflect.ValueOf(respData)
		respPtr reflect.Value
	)
	if resp.Kind() == reflect.Struct {
		f := reflect.New(reflect.TypeOf(respData))
		f.Elem().Set(reflect.ValueOf(respData))
		respPtr = f
	} else {
		respPtr = resp
	}
	meta := respPtr.Elem().FieldByName("Meta")
	var (
		code = int(meta.FieldByName("Code").Int())
		msg  = msgutil.GetMessage(moduleName, lng, code)
	)
	meta.FieldByName("Message").SetString(msg)

	c.Data["json"] = respPtr.Interface()
	c.Ctx.Output.SetStatus(httpcode.GetHTTPCode(code))
	c.ServeJSON()
}

// GetModuleAndMethodName returns module name.
func (c *BaseController) GetModuleAndMethodName() (mod, action string) {
	controller, action := c.GetControllerAndAction()
	mod = strings.Replace(controller, "Controller", "", -1)
	return mod, action
}

// GetLang returns Lang in header.
func (c *BaseController) GetLang() string {
	langHeader := c.Ctx.Input.Header(dtos2.HeaderLang)
	if langHeader == "" {
		return cfgutil.Load("DEFAULT_LANG")
	}
	return langHeader
}

// GetIDInt64 will return the integer type value of key parameter specified in query string.
func (c *BaseController) GetIDInt64(key string) (int64, error) {

	key = strings.TrimSpace(key)
	if key == "" {
		return 0, dtos2.NewAppError(dtos2.InvalidRequestError, fmt.Sprintf("Invalid Parameter : %s", key))
	}

	// Retrieve the value of parameter
	value, err := c.GetInt64(key)

	if err != nil {
		logger.Errorf(c.Ctx.Request.Context(), "Invalid %s", key)
		return 0, dtos2.NewAppError(dtos2.InvalidRequestError, fmt.Sprintf("Error occured during converting %s to int64.", key))
	}

	return value, nil
}

func (c *BaseController) logRequestBody() {
	if c.Ctx.Input.Method() == http.MethodPost ||
		c.Ctx.Input.Method() == http.MethodDelete ||
		c.Ctx.Input.Method() == http.MethodPut {
		var (
			out = bytes.Buffer{}
			err = json.Compact(&out, c.Ctx.Input.RequestBody)
		)
		if err == nil {
			var (
				logRequest     = out.String()
				exceptionField = cfgutil.Load("LOG_REQUEST_BLACKLIST_FIELD")
			)
			if exceptionField != "" {
				list := strings.Split(exceptionField, ",")
				for _, field := range list {
					var (
						pattern = "\"" + field + "\":\"[\\S\\s]+?\""
						re      = regexp.MustCompile(pattern)
						matches = re.FindStringSubmatch(logRequest)
					)
					for _, match := range matches {
						logRequest = strings.Replace(logRequest, match, "\""+field+"\":\"***\"", -1)
						break
					}
				}
			}
			logger.Infof(c.Ctx.Request.Context(), "Request body: %v", logRequest)
		}
	}
}
