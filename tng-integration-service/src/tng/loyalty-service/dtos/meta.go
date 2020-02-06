package dtos

import (
	"github.com/jinzhu/copier"
	"tng/common/httpcode"
)

const ReturnCodeOK = 1
const ReturnMsgOK = "OK"

type MetaRequest struct {
	MC        string `json:"mc";form:"mc"`
	TS        int64  `json:"ts";form:"ts"`
	UserAgent string `json:"user_agent";form:"user_agent"`
	ClientIP  string `json:"client_ip";form:"client_ip"`
	CheckSum  string `json:"check_sum";form:"check_sum"`
	Data      string `json:"data";form:"data"`
}

// Meta is common meta.
type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

// PaginationMeta is meta information with pagination info.
type PaginationMeta struct {
	Meta
	Total int64 `json:"total"`
}

// CommonResponse is common structure of response without Data.
type CommonResponse struct {
	Meta `json:"meta"`
}

// Data when export file
type Data struct {
	URL string `json:"url"`
}

// BuildCommonResponse returns a bad request response.
func BuildCommonResponse(code int, msg string) CommonResponse {
	responseMessage := msg
	if msg == "" {
		responseMessage = httpcode.GetHTTPStatusText(code)
	}
	return CommonResponse{
		Meta: Meta{
			Code:    code,
			Message: responseMessage,
		},
	}
}

// NewMeta returns a new meta with message.
func NewMeta(code int, messages ...string) Meta {
	msg := httpcode.GetHTTPStatusText(code)
	if len(messages) > 0 {
		msg = messages[0]
	}
	return Meta{
		Code:    code,
		Message: msg,
	}
}

func NewMetaOK() Meta {
	return Meta{
		Code:    ReturnCodeOK,
		Message: ReturnMsgOK,
	}
}

func GetMetaRequest(input interface{}) *MetaRequest {
	meta := &MetaRequest{}
	err := copier.Copy(meta, input)
	if err != nil {
		return nil
	}
	return meta
}
