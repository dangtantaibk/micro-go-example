package zalo

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
)

type Adapter interface {
	GetAccessToken(*GetAccessTokenRequest) (*GetAccessTokenResponse, error)
	GerProfileOA(*GetProfileOARequest) (*GetProfileOAResponse, error)
	GerProfileOfficial(*GetProfileOfficialRequest) (*GetProfileOfficialResponse, error)
}

type adapter struct {
	baseUrl string
}

func NewAdapter(baseUrl string) Adapter {
	return &adapter{
		baseUrl: baseUrl,
	}
}

func (a *adapter) GetAccessToken(request *GetAccessTokenRequest) (response *GetAccessTokenResponse, err error) {
	req := httplib.Get(a.baseUrl + EndpointGetAccessToken)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Param("app_id", request.AppID)
	req.Param("code", request.Code)
	req.Param("redirect_uri", request.RedirectUri)
	req.Param("isSDK", request.IsSDK)

	err = req.ToJSON(&response)
	fmt.Printf(response.AccessToken)

	return response, err
}

func (a *adapter) GerProfileOA(request *GetProfileOARequest) (response *GetProfileOAResponse, err error) {
	req := httplib.Get(a.baseUrl + EndpointGetUserProfileOA)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Param("access_token", request.AccessToken)
	data := DataGetProfileRequest{UserID: request.UserID}
	dataByteArray, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("marshal error")
	}
	req.Param("data", string(dataByteArray))
	err = req.ToJSON(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (a *adapter) GerProfileOfficial(request *GetProfileOfficialRequest) (response *GetProfileOfficialResponse, err error) {
	req := httplib.Get(a.baseUrl + EndpointGetUserProfileOfficial)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Param("access_token", request.AccessToken)
	err = req.ToJSON(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
