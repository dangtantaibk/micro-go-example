package h5_zalopay

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"tng/common/httpcli"
	"tng/common/logger"
	"tng/common/utils/hashutil"
)

type Adapter interface {
	GetMBToken(ctx context.Context, appID, maToken string) (*GetGrantedMBTokenResp, error)
	GetOrder(ctx context.Context, req *H5ZaloPayOrderURLReq) (*H5ZaloPayOrderURLResp, error)
}

type adapter struct {
	baseURL   string
	clientID  int64
	clientKey string
}

// NewAdapter provides access a Core Services
func NewAdapter(baseURL, clientKey string, clientID int64) Adapter {
	return &adapter{
		baseURL:   baseURL,
		clientID:  clientID,
		clientKey: clientKey,
	}
}

func (a *adapter) GetMBToken(ctx context.Context, appID, maToken string) (resp *GetGrantedMBTokenResp, err error) {
	var reqDate int64 = time.Now().Unix()
	dataInput := fmt.Sprintf("%s|%s|%s|%d", appID, a.clientKey, maToken, reqDate)
	sig := hashutil.GetSHA256(dataInput)
	header := map[string]string{
		"Content-Type": "application/json",
	}

	fullUrl := fmt.Sprintf("%s/?appid=%s&matoken=%s&reqdate=%d&sig=%s",
		EndpointGetMBToken,
		appID,
		maToken,
		reqDate,
		sig)

	reqBuilder := httpcli.NewRequestBuilder().
		SetBaseURL(a.baseURL).
		SetMethod(http.MethodGet).
		SetHeader(header).
		SetURL(fullUrl)
	err = reqBuilder.GetResponse(&resp)
	if err != nil {
		return nil, err
	}
	logger.Infof(ctx, "Response grantmbtoken: %v", resp)
	return resp, nil
}

func (a *adapter) GetOrder(ctx context.Context, req *H5ZaloPayOrderURLReq) (resp *H5ZaloPayOrderURLResp, err error) {
	var reqDate = time.Now().Unix()
	dataInput := fmt.Sprintf("%s|%d|%d|%d|%s", req.ZPOrderURL, req.AppID, reqDate, a.clientID, a.clientKey)
	sig := hashutil.GetSHA256(dataInput)
	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	bodyRaw := fmt.Sprintf("mbtoken=%s&appid=%d&zporderurl=%s&sig=%s&clientid=%d&reqdate=%d",
		req.MBToken,
		req.AppID,
		req.ZPOrderURL,
		sig,
		a.clientID,
		reqDate)
	reqBuilder := httpcli.NewRequestBuilder().
		SetBaseURL(a.baseURL).
		SetMethod(http.MethodPost).
		SetHeader(header).
		SetURL(EndpointZPOrderURL).
		SetBodyRaw(bodyRaw)
	err = reqBuilder.GetResponse(&resp)
	if err != nil {
		return nil, err
	}
	logger.Infof(ctx, "Response paymenturl: %v", resp)
	return resp, nil
}
