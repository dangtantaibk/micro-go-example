package httpcli

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// requestBuilder
type requestBuilder struct {
	baseURL           string
	URL               string
	method            string
	headers           map[string]string
	query             url.Values
	connectionTimeout int
	readWriteTimeout  int
	body              string
	username          string
	password          string
}

// RequestBuilder builder method for http client
type RequestBuilder interface {
	SetBaseURL(baseURL string) RequestBuilder
	SetURL(url string) RequestBuilder
	SetMethod(method string) RequestBuilder
	SetHeader(headers map[string]string) RequestBuilder
	SetQueryParams(v interface{}) RequestBuilder
	SetBody(v interface{}) RequestBuilder
	SetBodyRaw(v string) RequestBuilder
	// ToJSON(v interface{}) (err error)
	// ToString() (string, error)
	LogRequest() string
	GetRawResponse() ([]byte, error)
	GetResponse(result interface{}) error
	SetBasicAuth(username, password string) RequestBuilder
}

// NewRequestBuilder returns a new instance of MerchantRepository.
func NewRequestBuilder() RequestBuilder {
	return &requestBuilder{
		connectionTimeout: 5,
		readWriteTimeout:  30,
	}
}

// SetBaseUrl return base url http request
func (r *requestBuilder) SetBaseURL(baseURL string) RequestBuilder {
	r.baseURL = baseURL
	return r
}

// SetConnectionTimeout to modify timeout for connection
func (r *requestBuilder) SetConnectionTimeout(timeout int) RequestBuilder {
	r.connectionTimeout = timeout
	return r
}

// SetReadWriteTimeout to modify timeout for read write data
func (r *requestBuilder) SetReadWriteTimeout(timeout int) RequestBuilder {
	r.readWriteTimeout = timeout
	return r
}

// SetUrl to set url after base url
func (r *requestBuilder) SetURL(currentURL string) RequestBuilder {
	r.URL = currentURL
	return r
}

// SetMethod to set request method GET - POST - PUT - DELETE
func (r *requestBuilder) SetMethod(method string) RequestBuilder {
	r.method = method
	return r
}

// SetBasicAuth to set basic auth
func (r *requestBuilder) SetBasicAuth(username, password string) RequestBuilder {
	r.username = username
	r.password = password
	return r
}

// SetHeader to set request header for authentication
func (r *requestBuilder) SetHeader(headers map[string]string) RequestBuilder {
	r.headers = headers
	return r
}

// SetQueryParams to pass parameter to url
func (r *requestBuilder) SetQueryParams(v interface{}) RequestBuilder {
	params, _ := query.Values(v)
	r.query = params
	return r
}

// SetBody to set body json in post method
func (r *requestBuilder) SetBody(v interface{}) RequestBuilder {
	jsonValue, _ := json.Marshal(v)
	r.body = fmt.Sprintf("%s", jsonValue)
	return r
}

func (r *requestBuilder) SetBodyRaw(v string) RequestBuilder {
	r.body = v
	return r
}

// Build to response data from http request
func (r *requestBuilder) GetResponse(result interface{}) error {
	if len(r.query) > 0 {
		if strings.Contains(r.URL, "?") {
			r.URL += r.query.Encode()
		} else {
			r.URL += "?" + r.query.Encode()
		}
	}
	r.URL = r.baseURL + r.URL
	request := httplib.NewBeegoRequest(r.URL, r.method)

	for key, value := range r.headers {
		request.Header(key, value)
	}
	if r.username != "" && r.password != "" {
		request.SetBasicAuth(r.username, r.password)
	}

	request.SetTimeout(time.Duration(r.connectionTimeout)*time.Second, time.Duration(r.readWriteTimeout)*time.Second)
	response, err := request.Body(r.body).Response()

	if err != nil {
		return err
	}

	data, err := r.toBytes(response)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
}

// Build to response data from http request
func (r *requestBuilder) GetRawResponse() ([]byte, error) {
	if len(r.query) > 0 {
		if strings.Contains(r.URL, "?") {
			r.URL += r.query.Encode()
		} else {
			r.URL += "?" + r.query.Encode()
		}
	}
	r.URL = r.baseURL + r.URL
	request := httplib.NewBeegoRequest(r.URL, r.method)

	for key, value := range r.headers {
		request.Header(key, value)
	}
	if r.username != "" && r.password != "" {
		request.SetBasicAuth(r.username, r.password)
	}

	request.SetTimeout(time.Duration(r.connectionTimeout)*time.Second, time.Duration(r.readWriteTimeout)*time.Second)
	response, err := request.Body(r.body).Response()

	if err != nil {
		return nil, err
	}

	data, err := r.toBytes(response)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *requestBuilder) toBytes(response *http.Response) ([]byte, error) {
	defer func() {
		_ = response.Body.Close()
	}()

	if response.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(response.Body)
		if err != nil {
			return nil, err
		}
		return ioutil.ReadAll(reader)
	}

	return ioutil.ReadAll(response.Body)
}

// formatRequest generates ascii representation of a request
func (r *requestBuilder) LogRequest() string {
	// Create return string
	var request []string
	// Add the request string
	urlPath := fmt.Sprintf("%v %v %v", r.method, r.baseURL, r.URL)
	request = append(request, urlPath)
	// Loop through headers
	for name, header := range r.headers {
		request = append(request, fmt.Sprintf("%v: %v", name, header))
	}

	// If this is a POST, add post data
	request = append(append(request, "\n"), r.body)

	// Return the request as a string
	return strings.Join(request, "\n")
}
