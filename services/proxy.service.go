package services

import (
	"errors"
	"github.com/ariandi/kilat-be-go1/dto"
	util "github.com/ariandi/kilat-be-go1/utils"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type ProxyService struct {
	Config util.Config
	Client *resty.Client
}

var proxyService *ProxyService

type ProxyInterface interface {
	PostJsonService(url string, in []byte) (*resty.Response, error)
	PostFormService(url string, in map[string]string, header *dto.HeaderProxy) (*resty.Response, error)
}

// GetProxyService is
func GetProxyService(config util.Config, client *resty.Client) ProxyInterface {

	if proxyService == nil {

		proxyService = &ProxyService{
			Config: config,
			Client: client,
		}
	}

	return proxyService
}

func (o *ProxyService) PostJsonService(url string, in []byte) (*resty.Response, error) {
	logrus.Println("[ProxyService PostJsonService] start.")
	logrus.Println("[ProxyService PostJsonService] send request to digi url : ", url)
	logrus.Println("[ProxyService PostJsonService] send request to digi params : ", string(in))
	resp, err := o.Client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetBody(string(in)).
		Post(url)

	logrus.Println("[ProxyService PostJsonService] status code : ", resp.StatusCode())
	if errResp := CheckResponse(resp); errResp != nil {
		logrus.Println("[ProxyService PostJsonService] resp body is error : ", err)
		return nil, errors.New(errResp.ErrorMessage)
	}

	// another error not catch
	if err != nil {
		logrus.Println("[ProxyService PostJsonService] API Gateway - Error : ", err)
		return resp, err
	}

	logrus.Println("[ProxyService PostJsonService] status code : ", resp.StatusCode())
	logrus.Println("[ProxyService PostJsonService] response from biller : ", string(resp.Body()))
	return resp, nil
}

func (o *ProxyService) PostFormService(url string, in map[string]string, header *dto.HeaderProxy) (*resty.Response, error) {
	logrus.Println("[ProxyService PostFormService] start.")
	logrus.Println("[ProxyService PostFormService] send request to digi url : ", url)
	logrus.Println("[ProxyService PostFormService] send request to digi params : ", in)

	var resp *resty.Response
	var err error
	if header != nil {
		resp, err = o.Client.R().
			EnableTrace().
			SetHeader("Content-Type", "application/x-www-form-urlencoded").
			SetHeader("Content-Type", "application/x-www-form-urlencoded").
			SetFormData(in).
			Post(url)
	} else {
		resp, err = o.Client.R().
			EnableTrace().
			SetHeader("Content-Type", "application/x-www-form-urlencoded").
			SetFormData(in).
			Post(url)
	}

	logrus.Println("[ProxyService PostReqService] status code : ", resp.StatusCode())
	if errResp := CheckResponse(resp); errResp != nil {
		logrus.Println("[ProxyService InqPrePaid] resp body is error : ", err)
		return nil, errors.New(errResp.ErrorMessage)
	}

	// another error not catch
	if err != nil {
		logrus.Println("[ProxyService InqPrePaid] API Gateway - Error : ", err)
		return resp, err
	}

	logrus.Println("[ProxyService PostReqService] status code : ", resp.StatusCode())
	logrus.Println("[ProxyService PostReqService] response from biller : ", string(resp.Body()))
	return resp, nil
}

func CheckResponse(resp *resty.Response) *dto.ErrClientResp {
	var jsonER dto.ErrClientResp

	if resp.IsError() {
		return &jsonER
	}
	return nil
}
