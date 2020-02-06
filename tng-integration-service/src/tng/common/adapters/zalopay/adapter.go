package zalopay

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"tng/common/httpcli"
	"tng/common/utils/hashutil"
	"tng/common/utils/merchantutil"
)

type Adapter interface {
	CreateInvoice(context.Context, *ZlpCreateOrderRequest) (*ZlpCreateOrderResponse, error)
	GetInvoiceStatus(context.Context, *ZlpGetInvoiceStatusRequest) (*ZlpGetInvoiceStatusResponse, error)
	Refund(context.Context, *ZlpRefundInvoiceRequest) (*ZlpRefundInvoiceResponse, error)
	GetRefundStatus(context.Context, *ZlpGetRefundStatusRequest) (*ZlpGetRefundStatusResponse, error)
}
type adapter struct {
	baseURL string
}

func NewAdapter(baseURL string) Adapter {
	return &adapter{
		baseURL: baseURL,
	}
}

func (a *adapter) CreateInvoice(ctx context.Context, request *ZlpCreateOrderRequest) (response *ZlpCreateOrderResponse, err error) {
	params := make(url.Values)
	params.Add("appid", request.AppID)
	params.Add("amount", request.Amount)
	params.Add("appuser", request.AppUser)
	params.Add("embeddata", request.EmbedData)
	params.Add("item", request.Item)
	params.Add("description", request.Description)
	params.Add("bankcode", ZlpBankCode)

	now := time.Now()
	params.Add("apptime", strconv.FormatInt(now.UnixNano()/int64(time.Millisecond), 10))
	params.Add("apptransid", request.AppTransID)
	dataInput := fmt.Sprintf("%v|%v|%v|%v|%v|%v|%v",
		params.Get("appid"),
		params.Get("apptransid"),
		params.Get("appuser"),
		params.Get("amount"),
		params.Get("apptime"),
		params.Get("embeddata"),
		params.Get("item"),
	)
	key1 := merchantutil.GetValue(request.MerchantCode, merchantutil.MERCHANT_KEY1)
	mac := hashutil.GetHmacSHA256(key1, dataInput)
	params.Add("mac", mac)
	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	reqBuilder := httpcli.NewRequestBuilder().
		SetBaseURL(a.baseURL).
		SetMethod(http.MethodPost).
		SetHeader(header).
		SetURL(EndpointCreateOrder).
		SetQueryParams(params)
	err = reqBuilder.GetResponse(&response)
	return response, err
}

func (a *adapter) GetInvoiceStatus(ctx context.Context, request *ZlpGetInvoiceStatusRequest) (response *ZlpGetInvoiceStatusResponse, err error) {
	params := make(url.Values)
	params.Add("appid", request.AppID)
	params.Add("apptransid", request.AppTransID)
	key1 := merchantutil.GetValue(request.MerchantCode, merchantutil.MERCHANT_KEY1)
	dataInput := fmt.Sprintf("%s|%s|%s",
		request.AppID,
		params.Get("apptransid"),
		key1,
	)
	mac := hashutil.GetHmacSHA256(key1, dataInput)
	header := map[string]string{
		"Content-Type": "application/json",
	}
	params.Add("mac", mac)
	reqBuilder := httpcli.NewRequestBuilder().
		SetBaseURL(a.baseURL).
		SetMethod(http.MethodGet).
		SetHeader(header).
		SetURL(EndpointGetOrderStatus).
		SetQueryParams(params)
	err = reqBuilder.GetResponse(&response)
	return response, err
}

func (a *adapter) Refund(ctx context.Context, request *ZlpRefundInvoiceRequest) (response *ZlpRefundInvoiceResponse, err error) {
	params := make(url.Values)
	params.Add("appid", request.AppID)
	params.Add("zptransid", request.ZpTransID)
	params.Add("amount", request.Amount)
	params.Add("description", request.Description)

	now := time.Now()
	timestamp := now.UnixNano() / int64(time.Millisecond)
	params.Add("timestamp", strconv.FormatInt(timestamp, 10))

	uid := fmt.Sprintf("%d%d", timestamp, 111+rand.Intn(888))
	params.Add("mrefundid", fmt.Sprintf("%02d%02d%02d_%v_%v",
		now.Year()%100,
		int(now.Month()),
		now.Day(),
		request.AppID,
		uid),
	)

	dataInput := fmt.Sprintf("%v|%v|%v|%v|%v",
		request.AppID,
		params.Get("zptransid"),
		params.Get("amount"),
		params.Get("description"),
		params.Get("timestamp"),
	)
	key1 := merchantutil.GetValue(request.MerchantCode, merchantutil.MERCHANT_KEY1)
	mac := hashutil.GetHmacSHA256(key1, dataInput)
	params.Add("mac", mac)
	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	reqBuilder := httpcli.NewRequestBuilder().
		SetBaseURL(a.baseURL).
		SetMethod(http.MethodPost).
		SetHeader(header).
		SetURL(EndpointRefund).
		SetQueryParams(params)
	err = reqBuilder.GetResponse(&response)
	return response, err
}

func (a *adapter) GetRefundStatus(ctx context.Context, request *ZlpGetRefundStatusRequest) (response *ZlpGetRefundStatusResponse, err error) {
	params := make(url.Values)
	params.Add("appid", request.AppID)
	params.Add("mrefundid", request.MRefundID)
	now := time.Now()
	timestamp := now.UnixNano() / int64(time.Millisecond)
	params.Add("timestamp", strconv.FormatInt(timestamp, 10))

	dataInput := fmt.Sprintf("%v|%v|%v",
		request.AppID,
		params.Get("mrefundid"),
		params.Get("timestamp"),
	)
	key1 := merchantutil.GetValue(request.MerchantCode, merchantutil.MERCHANT_KEY1)
	mac := hashutil.GetHmacSHA256(key1, dataInput)
	params.Add("mac", mac)
	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	reqBuilder := httpcli.NewRequestBuilder().
		SetBaseURL(a.baseURL).
		SetMethod(http.MethodGet).
		SetHeader(header).
		SetURL(EndpointGetRefundStatus).
		SetQueryParams(params)
	err = reqBuilder.GetResponse(&response)
	return response, err
}
