package resty

import (
	"net/url"

	"github.com/go-resty/resty/v2"
)

var client = resty.New().SetDebug(true)

type HttpClient struct {
	Body    interface{}
	Headers map[string]string
	URL     string
}

type HttpClientParam struct {
	Params  map[string]string
	Headers map[string]string
	URL     string
}
type HttpClientParamEncode struct {
	Params  url.Values
	Headers map[string]string
	URL     string
}

func (httpClient *HttpClient) Post() (*resty.Response, error) {
	return client.R().SetHeaders(httpClient.Headers).SetBody(httpClient.Body).Post(httpClient.URL)
}
func (httpClient *HttpClientParamEncode) PostFormData() (*resty.Response, error) {
	return client.R().SetDebug(false).SetHeaders(httpClient.Headers).SetFormDataFromValues(httpClient.Params).Post(httpClient.URL)
}

func (httpClient *HttpClientParam) Get() (*resty.Response, error) {
	return client.R().SetHeaders(httpClient.Headers).SetQueryParams(httpClient.Params).Get(httpClient.URL)
}
